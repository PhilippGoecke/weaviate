//go:build !noasm && amd64
// AUTO-GENERATED BY GOAT -- DO NOT EDIT

TEXT ·l2(SB), $0-32
	MOVQ a+0(FP), DI
	MOVQ b+8(FP), SI
	MOVQ res+16(FP), DX
	MOVQ len+24(FP), CX
	BYTE $0x55               // pushq	%rbp
	WORD $0x8948; BYTE $0xe5 // movq	%rsp, %rbp
	LONG $0xf8e48348         // andq	$-8, %rsp
	WORD $0x018b             // movl	(%rcx), %eax
	WORD $0xf883; BYTE $0x07 // cmpl	$7, %eax
	JG   LBB0_9
	LONG $0xff408d44         // leal	-1(%rax), %r8d
	WORD $0x03a8             // testb	$3, %al
	JE   LBB0_2
	WORD $0x8941; BYTE $0xc1 // movl	%eax, %r9d
	LONG $0x03e18341         // andl	$3, %r9d
	LONG $0xc057f8c5         // vxorps	%xmm0, %xmm0, %xmm0
	WORD $0xc931             // xorl	%ecx, %ecx

LBB0_4:
	LONG $0x0f10fac5         // vmovss	(%rdi), %xmm1
	LONG $0x0e5cf2c5         // vsubss	(%rsi), %xmm1, %xmm1
	LONG $0xc959f2c5         // vmulss	%xmm1, %xmm1, %xmm1
	LONG $0xc158fac5         // vaddss	%xmm1, %xmm0, %xmm0
	LONG $0x04c78348         // addq	$4, %rdi
	LONG $0x04c68348         // addq	$4, %rsi
	LONG $0x01c18348         // addq	$1, %rcx
	WORD $0x3941; BYTE $0xc9 // cmpl	%ecx, %r9d
	JNE  LBB0_4
	WORD $0xc829             // subl	%ecx, %eax
	LONG $0x03f88341         // cmpl	$3, %r8d
	JAE  LBB0_7

LBB0_26:
	LONG $0x0211fac5         // vmovss	%xmm0, (%rdx)
	WORD $0x8948; BYTE $0xec // movq	%rbp, %rsp
	BYTE $0x5d               // popq	%rbp
	BYTE $0xc3               // retq

LBB0_9:
	LONG $0xc057f8c5         // vxorps	%xmm0, %xmm0, %xmm0
	WORD $0xf883; BYTE $0x20 // cmpl	$32, %eax
	JB   LBB0_10
	LONG $0xc057f8c5         // vxorps	%xmm0, %xmm0, %xmm0
	LONG $0xc957f0c5         // vxorps	%xmm1, %xmm1, %xmm1
	LONG $0xd257e8c5         // vxorps	%xmm2, %xmm2, %xmm2
	LONG $0xdb57e0c5         // vxorps	%xmm3, %xmm3, %xmm3

LBB0_16:
	LONG $0x2710fcc5             // vmovups	(%rdi), %ymm4
	LONG $0x6f10fcc5; BYTE $0x20 // vmovups	32(%rdi), %ymm5
	LONG $0x7710fcc5; BYTE $0x40 // vmovups	64(%rdi), %ymm6
	LONG $0x7f10fcc5; BYTE $0x60 // vmovups	96(%rdi), %ymm7
	LONG $0x265cdcc5             // vsubps	(%rsi), %ymm4, %ymm4
	LONG $0x6e5cd4c5; BYTE $0x20 // vsubps	32(%rsi), %ymm5, %ymm5
	LONG $0x765cccc5; BYTE $0x40 // vsubps	64(%rsi), %ymm6, %ymm6
	LONG $0x7e5cc4c5; BYTE $0x60 // vsubps	96(%rsi), %ymm7, %ymm7
	LONG $0xb85de2c4; BYTE $0xdc // vfmadd231ps	%ymm4, %ymm4, %ymm3
	LONG $0xb855e2c4; BYTE $0xd5 // vfmadd231ps	%ymm5, %ymm5, %ymm2
	LONG $0xb84de2c4; BYTE $0xce // vfmadd231ps	%ymm6, %ymm6, %ymm1
	LONG $0xb845e2c4; BYTE $0xc7 // vfmadd231ps	%ymm7, %ymm7, %ymm0
	WORD $0xc083; BYTE $0xe0     // addl	$-32, %eax
	LONG $0x80ef8348             // subq	$-128, %rdi
	LONG $0x80ee8348             // subq	$-128, %rsi
	WORD $0xf883; BYTE $0x1f     // cmpl	$31, %eax
	JA   LBB0_16
	WORD $0xf883; BYTE $0x08     // cmpl	$8, %eax
	JAE  LBB0_11
	JMP  LBB0_13

LBB0_10:
	LONG $0xc957f0c5 // vxorps	%xmm1, %xmm1, %xmm1
	LONG $0xd257e8c5 // vxorps	%xmm2, %xmm2, %xmm2
	LONG $0xdb57e0c5 // vxorps	%xmm3, %xmm3, %xmm3

LBB0_11:
	LONG $0x2710fcc5             // vmovups	(%rdi), %ymm4
	LONG $0x265cdcc5             // vsubps	(%rsi), %ymm4, %ymm4
	LONG $0xb85de2c4; BYTE $0xdc // vfmadd231ps	%ymm4, %ymm4, %ymm3
	WORD $0xc083; BYTE $0xf8     // addl	$-8, %eax
	LONG $0x20c78348             // addq	$32, %rdi
	LONG $0x20c68348             // addq	$32, %rsi
	WORD $0xf883; BYTE $0x07     // cmpl	$7, %eax
	JA   LBB0_11

LBB0_13:
	WORD $0xc085             // testl	%eax, %eax
	JE   LBB0_14
	LONG $0xff408d44         // leal	-1(%rax), %r8d
	WORD $0x03a8             // testb	$3, %al
	JE   LBB0_18
	WORD $0x8941; BYTE $0xc1 // movl	%eax, %r9d
	LONG $0x03e18341         // andl	$3, %r9d
	LONG $0xe457d8c5         // vxorps	%xmm4, %xmm4, %xmm4
	WORD $0xc931             // xorl	%ecx, %ecx

LBB0_20:
	LONG $0x2f10fac5         // vmovss	(%rdi), %xmm5
	LONG $0x2e5cd2c5         // vsubss	(%rsi), %xmm5, %xmm5
	LONG $0xed59d2c5         // vmulss	%xmm5, %xmm5, %xmm5
	LONG $0xe558dac5         // vaddss	%xmm5, %xmm4, %xmm4
	LONG $0x04c78348         // addq	$4, %rdi
	LONG $0x04c68348         // addq	$4, %rsi
	LONG $0x01c18348         // addq	$1, %rcx
	WORD $0x3941; BYTE $0xc9 // cmpl	%ecx, %r9d
	JNE  LBB0_20
	WORD $0xc829             // subl	%ecx, %eax
	LONG $0x03f88341         // cmpl	$3, %r8d
	JAE  LBB0_23
	JMP  LBB0_25

LBB0_2:
	LONG $0xc057f8c5 // vxorps	%xmm0, %xmm0, %xmm0
	LONG $0x03f88341 // cmpl	$3, %r8d
	JB   LBB0_26

LBB0_7:
	WORD $0xc089 // movl	%eax, %eax
	WORD $0xc931 // xorl	%ecx, %ecx

LBB0_8:
	LONG $0x0c10fac5; BYTE $0x8f   // vmovss	(%rdi,%rcx,4), %xmm1
	LONG $0x5410fac5; WORD $0x048f // vmovss	4(%rdi,%rcx,4), %xmm2
	LONG $0x0c5cf2c5; BYTE $0x8e   // vsubss	(%rsi,%rcx,4), %xmm1, %xmm1
	LONG $0xc959f2c5               // vmulss	%xmm1, %xmm1, %xmm1
	LONG $0xc158fac5               // vaddss	%xmm1, %xmm0, %xmm0
	LONG $0x4c5ceac5; WORD $0x048e // vsubss	4(%rsi,%rcx,4), %xmm2, %xmm1
	LONG $0xc959f2c5               // vmulss	%xmm1, %xmm1, %xmm1
	LONG $0xc158fac5               // vaddss	%xmm1, %xmm0, %xmm0
	LONG $0x4c10fac5; WORD $0x088f // vmovss	8(%rdi,%rcx,4), %xmm1
	LONG $0x4c5cf2c5; WORD $0x088e // vsubss	8(%rsi,%rcx,4), %xmm1, %xmm1
	LONG $0xc959f2c5               // vmulss	%xmm1, %xmm1, %xmm1
	LONG $0xc158fac5               // vaddss	%xmm1, %xmm0, %xmm0
	LONG $0x4c10fac5; WORD $0x0c8f // vmovss	12(%rdi,%rcx,4), %xmm1
	LONG $0x4c5cf2c5; WORD $0x0c8e // vsubss	12(%rsi,%rcx,4), %xmm1, %xmm1
	LONG $0xc959f2c5               // vmulss	%xmm1, %xmm1, %xmm1
	LONG $0xc158fac5               // vaddss	%xmm1, %xmm0, %xmm0
	LONG $0x04c18348               // addq	$4, %rcx
	WORD $0xc839                   // cmpl	%ecx, %eax
	JNE  LBB0_8
	JMP  LBB0_26

LBB0_14:
	LONG $0xe457d8c5 // vxorps	%xmm4, %xmm4, %xmm4
	JMP  LBB0_25

LBB0_18:
	LONG $0xe457d8c5 // vxorps	%xmm4, %xmm4, %xmm4
	LONG $0x03f88341 // cmpl	$3, %r8d
	JB   LBB0_25

LBB0_23:
	WORD $0xc089 // movl	%eax, %eax
	WORD $0xc931 // xorl	%ecx, %ecx

LBB0_24:
	LONG $0x2c10fac5; BYTE $0x8f   // vmovss	(%rdi,%rcx,4), %xmm5
	LONG $0x7410fac5; WORD $0x048f // vmovss	4(%rdi,%rcx,4), %xmm6
	LONG $0x2c5cd2c5; BYTE $0x8e   // vsubss	(%rsi,%rcx,4), %xmm5, %xmm5
	LONG $0xed59d2c5               // vmulss	%xmm5, %xmm5, %xmm5
	LONG $0xe558dac5               // vaddss	%xmm5, %xmm4, %xmm4
	LONG $0x6c5ccac5; WORD $0x048e // vsubss	4(%rsi,%rcx,4), %xmm6, %xmm5
	LONG $0xed59d2c5               // vmulss	%xmm5, %xmm5, %xmm5
	LONG $0xe558dac5               // vaddss	%xmm5, %xmm4, %xmm4
	LONG $0x6c10fac5; WORD $0x088f // vmovss	8(%rdi,%rcx,4), %xmm5
	LONG $0x6c5cd2c5; WORD $0x088e // vsubss	8(%rsi,%rcx,4), %xmm5, %xmm5
	LONG $0xed59d2c5               // vmulss	%xmm5, %xmm5, %xmm5
	LONG $0xe558dac5               // vaddss	%xmm5, %xmm4, %xmm4
	LONG $0x6c10fac5; WORD $0x0c8f // vmovss	12(%rdi,%rcx,4), %xmm5
	LONG $0x6c5cd2c5; WORD $0x0c8e // vsubss	12(%rsi,%rcx,4), %xmm5, %xmm5
	LONG $0xed59d2c5               // vmulss	%xmm5, %xmm5, %xmm5
	LONG $0xe558dac5               // vaddss	%xmm5, %xmm4, %xmm4
	LONG $0x04c18348               // addq	$4, %rcx
	WORD $0xc839                   // cmpl	%ecx, %eax
	JNE  LBB0_24

LBB0_25:
	LONG $0xd358ecc5               // vaddps	%ymm3, %ymm2, %ymm2
	LONG $0xc058f4c5               // vaddps	%ymm0, %ymm1, %ymm0
	LONG $0xc258fcc5               // vaddps	%ymm2, %ymm0, %ymm0
	LONG $0xc07cffc5               // vhaddps	%ymm0, %ymm0, %ymm0
	LONG $0xc07cffc5               // vhaddps	%ymm0, %ymm0, %ymm0
	LONG $0x197de3c4; WORD $0x01c1 // vextractf128	$1, %ymm0, %xmm1
	LONG $0xc158fac5               // vaddss	%xmm1, %xmm0, %xmm0
	LONG $0xc058dac5               // vaddss	%xmm0, %xmm4, %xmm0
	LONG $0x0211fac5               // vmovss	%xmm0, (%rdx)
	WORD $0x8948; BYTE $0xec       // movq	%rbp, %rsp
	BYTE $0x5d                     // popq	%rbp
	WORD $0xf8c5; BYTE $0x77       // vzeroupper
	BYTE $0xc3                     // retq
