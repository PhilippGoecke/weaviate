//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2024 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/jessevdk/go-flags"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"github.com/weaviate/weaviate/adapters/handlers/rest"
	"github.com/weaviate/weaviate/adapters/repos/db/inverted/stopwords"
	enterrors "github.com/weaviate/weaviate/entities/errors"
	"github.com/weaviate/weaviate/exp/query"
	"github.com/weaviate/weaviate/exp/queryschema"
	"github.com/weaviate/weaviate/grpc/generated/protocol/v1"
	modsloads3 "github.com/weaviate/weaviate/modules/offload-s3"
	"github.com/weaviate/weaviate/modules/text2vec-contextionary/client"
	"github.com/weaviate/weaviate/modules/text2vec-contextionary/vectorizer"
	"github.com/weaviate/weaviate/usecases/build"
	"github.com/weaviate/weaviate/usecases/monitoring"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	TargetQuerier = "querier"
)

// TODO: We want this to be part of original `cmd/weaviate-server`.
// But for some reason that binary is auto-generated and I couldn't modify as I need. Hence separate binary for now
func main() {
	var (
		opts Options
		log  logrus.FieldLogger
	)

	log = logrus.WithFields(logrus.Fields{"app": "weaviate"})

	_, err := flags.Parse(&opts)
	if err != nil {
		if err.(*flags.Error).Type == flags.ErrHelp {
			os.Exit(1)
		}
		log.WithField("err", err).Fatal("failed to parse command line args")
	}

	// Set version from swagger spec.
	build.Version = rest.ParseVersionFromSwaggerSpec()

	switch opts.Target {
	case "querier":
		log = log.WithField("target", "querier")
		s3module := modsloads3.New()
		s3module.DataPath = opts.Query.DataPath
		s3module.Bucket = opts.Query.S3URL
		s3module.Endpoint = opts.Query.S3Endpoint

		// This functionality is already in `go-client` of weaviate.
		// TODO(kavi): Find a way to share this functionality in both go-client and in querytenant.
		schemaInfo := queryschema.NewSchemaInfo(opts.Query.SchemaAddr, queryschema.DefaultSchemaPrefix)

		vclient, err := client.NewClient(opts.Query.VectorizerAddr, log)
		if err != nil {
			log.WithFields(logrus.Fields{
				"err":   err,
				"addrs": opts.Query.VectorizerAddr,
			}).Fatal("failed to talk to vectorizer")
		}

		detectStopwords, err := stopwords.NewDetectorFromPreset(stopwords.EnglishPreset)
		if err != nil {
			log.WithFields(logrus.Fields{
				"err": err,
			}).Fatal("failed to create stopwords detector for querier")
		}

		a := query.NewAPI(
			schemaInfo,
			s3module,
			vectorizer.New(vclient),
			detectStopwords,
			&opts.Query,
			log,
		)

		grpcQuerier := query.NewGRPC(a, schemaInfo, log)
		listener, err := net.Listen("tcp", opts.Query.GRPCListenAddr)
		if err != nil {
			log.WithFields(logrus.Fields{
				"err":   err,
				"addrs": opts.Query.GRPCListenAddr,
			}).Fatal("failed to bind grpc server addr")
		}
		svrMetrics := monitoring.NewServerMetrics(opts.Monitoring, prometheus.DefaultRegisterer)
		listener = monitoring.CountingListener(listener, svrMetrics.TCPActiveConnections.WithLabelValues("grpc"))
		grpcServer := grpc.NewServer(GrpcOptions(*svrMetrics)...)
		reflection.Register(grpcServer)
		protocol.RegisterWeaviateServer(grpcServer, grpcQuerier)

		enterrors.GoWrapper(func() {
			metadataSubscription := query.NewMetadataSubscription(
				a,
				opts.Query.MetadataGRPCAddress,
				log)
			if err = metadataSubscription.Start(); err != nil {
				log.WithError(err).Warnf("Failed to start metadata subscription")
			}
		}, log)

		log.WithField("addr", opts.Query.GRPCListenAddr).Info("starting querier over grpc")
		enterrors.GoWrapper(func() {
			if err := grpcServer.Serve(listener); err != nil {
				log.Fatal("failed to start grpc server", err)
			}
		}, log)

		mc := memcache.New("mymemcached:11211")
		bucket := opts.Query.S3URL
		// Initialize AWS S3 client
		awsCfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-central-1"))
		if err != nil {
			log.Fatalf("unable to load SDK config, %v", err)
		}
		s3Client := s3.NewFromConfig(awsCfg)
		for _, key := range query.AllS3Paths {
			// Download the object
			objBytes, err := downloadS3Object(s3Client, bucket, key)
			if err != nil {
				log.Printf("failed to download object %s: %v", key, err)
				continue
			}

			// Store object in Memcached
			err = mc.Set(&memcache.Item{
				Key:   key,
				Value: objBytes,
			})
			if err != nil {
				log.Printf("failed to save object to Memcached %s: %v", key, err)
				continue
			}

			fmt.Printf("Successfully saved %s to Memcached\n", key)
		}
		err = mc.Set(&memcache.Item{Key: "foo", Value: []byte("my value")})
		if err != nil {
			log.WithError(err).Fatal("failed to set item in memcache")
		}

		// serve /metrics
		mux := http.NewServeMux()
		mux.Handle("/metrics", promhttp.Handler())
		log.WithField("addr", opts.Monitoring.Port).Info("starting /metrics server over http")
		http.ListenAndServe(fmt.Sprintf(":%d", opts.Monitoring.Port), mux)
	default:
		log.Fatal("--target empty or unknown")
	}
}

func downloadS3Object(client *s3.Client, bucket, key string) ([]byte, error) {
	// Get the object from S3
	output, err := client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}
	defer output.Body.Close()

	// Read the object data
	data, err := io.ReadAll(output.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Options represents Command line options passed to weaviate binary
type Options struct {
	Target     string            `long:"target" description:"how should weaviate-server be running as e.g: querier, ingester, etc"`
	Query      query.Config      `group:"query" namespace:"query"`
	Monitoring monitoring.Config `group:"monitoring" namespace:"monitoring"`
}

func GrpcOptions(svrMetrics monitoring.ServerMetrics) []grpc.ServerOption {
	grpcOptions := []grpc.ServerOption{
		grpc.StatsHandler(monitoring.NewGrpcStatsHandler(
			svrMetrics.InflightRequests,
			svrMetrics.RequestBodySize,
			svrMetrics.ResponseBodySize,
		)),
	}

	grpcInterceptUnary := grpc.ChainUnaryInterceptor(
		monitoring.UnaryServerInstrument(svrMetrics.RequestDuration),
	)
	grpcOptions = append(grpcOptions, grpcInterceptUnary)

	grpcInterceptStream := grpc.ChainStreamInterceptor(
		monitoring.StreamServerInstrument(svrMetrics.RequestDuration),
	)
	grpcOptions = append(grpcOptions, grpcInterceptStream)

	return grpcOptions
}
