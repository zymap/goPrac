"".main STEXT size=692 args=0x0 locals=0x208
	0x0000 00000 (test.go:8)	TEXT	"".main(SB), $520-0
	0x0000 00000 (test.go:8)	MOVQ	(TLS), CX
	0x0009 00009 (test.go:8)	LEAQ	-392(SP), AX
	0x0011 00017 (test.go:8)	CMPQ	AX, 16(CX)
	0x0015 00021 (test.go:8)	JLS	682
	0x001b 00027 (test.go:8)	SUBQ	$520, SP
	0x0022 00034 (test.go:8)	MOVQ	BP, 512(SP)
	0x002a 00042 (test.go:8)	LEAQ	512(SP), BP
	0x0032 00050 (test.go:8)	FUNCDATA	$0, gclocals·f0a67958015464e4cc8847ce0df60843(SB)
	0x0032 00050 (test.go:8)	FUNCDATA	$1, gclocals·f74a0b8caecc0f0df89956b0bfefa0fe(SB)
	0x0032 00050 (test.go:8)	FUNCDATA	$3, gclocals·7f910c4fecabcdd009afbf1e2d5c5c17(SB)
	0x0032 00050 (test.go:9)	PCDATA	$2, $0
	0x0032 00050 (test.go:9)	PCDATA	$0, $1
	0x0032 00050 (test.go:9)	XORPS	X1, X1
	0x0035 00053 (test.go:9)	MOVUPS	X1, ""..autotmp_6+184(SP)
	0x003d 00061 (test.go:9)	XORPS	X1, X1
	0x0040 00064 (test.go:9)	MOVUPS	X1, ""..autotmp_6+200(SP)
	0x0048 00072 (test.go:9)	XORPS	X1, X1
	0x004b 00075 (test.go:9)	MOVUPS	X1, ""..autotmp_6+216(SP)
	0x0053 00083 (test.go:9)	PCDATA	$2, $1
	0x0053 00083 (test.go:9)	PCDATA	$0, $2
	0x0053 00083 (test.go:9)	LEAQ	""..autotmp_7+304(SP), DI
	0x005b 00091 (test.go:9)	XORPS	X0, X0
	0x005e 00094 (test.go:9)	PCDATA	$2, $0
	0x005e 00094 (test.go:9)	LEAQ	-48(DI), DI
	0x0062 00098 (test.go:9)	DUFFZERO	$239
	0x0075 00117 (test.go:9)	PCDATA	$2, $2
	0x0075 00117 (test.go:9)	LEAQ	""..autotmp_6+184(SP), AX
	0x007d 00125 (test.go:9)	PCDATA	$2, $0
	0x007d 00125 (test.go:9)	TESTB	AL, (AX)
	0x007f 00127 (test.go:9)	PCDATA	$2, $2
	0x007f 00127 (test.go:9)	LEAQ	""..autotmp_7+304(SP), AX
	0x0087 00135 (test.go:9)	PCDATA	$2, $0
	0x0087 00135 (test.go:9)	MOVQ	AX, ""..autotmp_6+200(SP)
	0x008f 00143 (test.go:9)	PCDATA	$2, $2
	0x008f 00143 (test.go:9)	LEAQ	""..autotmp_6+184(SP), AX
	0x0097 00151 (test.go:9)	PCDATA	$2, $0
	0x0097 00151 (test.go:9)	PCDATA	$0, $3
	0x0097 00151 (test.go:9)	MOVQ	AX, ""..autotmp_8+72(SP)
	0x009c 00156 (test.go:9)	CALL	runtime.fastrand(SB)
	0x00a1 00161 (test.go:9)	PCDATA	$2, $2
	0x00a1 00161 (test.go:9)	PCDATA	$0, $2
	0x00a1 00161 (test.go:9)	MOVQ	""..autotmp_8+72(SP), AX
	0x00a6 00166 (test.go:9)	TESTB	AL, (AX)
	0x00a8 00168 (test.go:9)	MOVL	(SP), CX
	0x00ab 00171 (test.go:9)	PCDATA	$2, $0
	0x00ab 00171 (test.go:9)	MOVL	CX, 12(AX)
	0x00ae 00174 (test.go:9)	PCDATA	$2, $2
	0x00ae 00174 (test.go:9)	LEAQ	""..autotmp_6+184(SP), AX
	0x00b6 00182 (test.go:9)	PCDATA	$2, $0
	0x00b6 00182 (test.go:9)	PCDATA	$0, $4
	0x00b6 00182 (test.go:9)	MOVQ	AX, "".m+56(SP)
	0x00bb 00187 (test.go:13)	PCDATA	$2, $2
	0x00bb 00187 (test.go:13)	PCDATA	$0, $5
	0x00bb 00187 (test.go:13)	LEAQ	""..autotmp_10+232(SP), AX
	0x00c3 00195 (test.go:13)	PCDATA	$0, $6
	0x00c3 00195 (test.go:13)	MOVQ	AX, ""..autotmp_9+64(SP)
	0x00c8 00200 (test.go:13)	PCDATA	$2, $0
	0x00c8 00200 (test.go:13)	TESTB	AL, (AX)
	0x00ca 00202 (test.go:13)	MOVQ	"".statictmp_0(SB), AX
	0x00d1 00209 (test.go:13)	MOVQ	AX, ""..autotmp_10+232(SP)
	0x00d9 00217 (test.go:13)	MOVUPS	"".statictmp_0+8(SB), X1
	0x00e0 00224 (test.go:13)	MOVUPS	X1, ""..autotmp_10+240(SP)
	0x00e8 00232 (test.go:13)	MOVUPS	"".statictmp_0+24(SB), X1
	0x00ef 00239 (test.go:13)	MOVUPS	X1, ""..autotmp_10+256(SP)
	0x00f7 00247 (test.go:13)	MOVUPS	"".statictmp_0+40(SB), X1
	0x00fe 00254 (test.go:13)	MOVUPS	X1, ""..autotmp_10+272(SP)
	0x0106 00262 (test.go:13)	MOVUPS	"".statictmp_0+56(SB), X1
	0x010d 00269 (test.go:13)	MOVUPS	X1, ""..autotmp_10+288(SP)
	0x0115 00277 (test.go:13)	PCDATA	$2, $2
	0x0115 00277 (test.go:13)	PCDATA	$0, $5
	0x0115 00277 (test.go:13)	MOVQ	""..autotmp_9+64(SP), AX
	0x011a 00282 (test.go:13)	TESTB	AL, (AX)
	0x011c 00284 (test.go:13)	JMP	286
	0x011e 00286 (test.go:13)	PCDATA	$2, $0
	0x011e 00286 (test.go:13)	PCDATA	$0, $7
	0x011e 00286 (test.go:13)	MOVQ	AX, "".stus+112(SP)
	0x0123 00291 (test.go:13)	MOVQ	$3, "".stus+120(SP)
	0x012c 00300 (test.go:13)	MOVQ	$3, "".stus+128(SP)
	0x0138 00312 (test.go:16)	PCDATA	$2, $2
	0x0138 00312 (test.go:16)	LEAQ	type."".stu(SB), AX
	0x013f 00319 (test.go:16)	PCDATA	$2, $0
	0x013f 00319 (test.go:16)	MOVQ	AX, (SP)
	0x0143 00323 (test.go:16)	CALL	runtime.newobject(SB)
	0x0148 00328 (test.go:16)	PCDATA	$2, $2
	0x0148 00328 (test.go:16)	MOVQ	8(SP), AX
	0x014d 00333 (test.go:16)	PCDATA	$2, $0
	0x014d 00333 (test.go:16)	PCDATA	$0, $8
	0x014d 00333 (test.go:16)	MOVQ	AX, "".&stu+104(SP)
	0x0152 00338 (test.go:16)	MOVQ	"".stus+128(SP), AX
	0x015a 00346 (test.go:16)	MOVQ	"".stus+120(SP), CX
	0x015f 00351 (test.go:16)	PCDATA	$2, $3
	0x015f 00351 (test.go:16)	PCDATA	$0, $9
	0x015f 00351 (test.go:16)	MOVQ	"".stus+112(SP), DX
	0x0164 00356 (test.go:16)	PCDATA	$2, $0
	0x0164 00356 (test.go:16)	PCDATA	$0, $10
	0x0164 00356 (test.go:16)	MOVQ	DX, ""..autotmp_4+136(SP)
	0x016c 00364 (test.go:16)	MOVQ	CX, ""..autotmp_4+144(SP)
	0x0174 00372 (test.go:16)	MOVQ	AX, ""..autotmp_4+152(SP)
	0x017c 00380 (test.go:16)	MOVQ	$0, ""..autotmp_11+48(SP)
	0x0185 00389 (test.go:16)	MOVQ	""..autotmp_4+144(SP), AX
	0x018d 00397 (test.go:16)	MOVQ	AX, ""..autotmp_12+40(SP)
	0x0192 00402 (test.go:16)	PCDATA	$2, $2
	0x0192 00402 (test.go:16)	PCDATA	$0, $9
	0x0192 00402 (test.go:16)	MOVQ	""..autotmp_4+136(SP), AX
	0x019a 00410 (test.go:16)	PCDATA	$2, $0
	0x019a 00410 (test.go:16)	PCDATA	$0, $11
	0x019a 00410 (test.go:16)	MOVQ	AX, ""..autotmp_13+96(SP)
	0x019f 00415 (test.go:16)	MOVQ	""..autotmp_12+40(SP), AX
	0x01a4 00420 (test.go:16)	CMPQ	""..autotmp_11+48(SP), AX
	0x01a9 00425 (test.go:16)	JLT	432
	0x01ab 00427 (test.go:16)	JMP	680
	0x01b0 00432 (test.go:16)	PCDATA	$2, $-2
	0x01b0 00432 (test.go:16)	PCDATA	$0, $-2
	0x01b0 00432 (test.go:16)	JMP	434
	0x01b2 00434 (test.go:16)	PCDATA	$2, $4
	0x01b2 00434 (test.go:16)	PCDATA	$0, $11
	0x01b2 00434 (test.go:16)	MOVQ	""..autotmp_13+96(SP), CX
	0x01b7 00439 (test.go:16)	TESTB	AL, (CX)
	0x01b9 00441 (test.go:16)	MOVQ	16(CX), DX
	0x01bd 00445 (test.go:16)	PCDATA	$2, $5
	0x01bd 00445 (test.go:16)	MOVQ	(CX), AX
	0x01c0 00448 (test.go:16)	PCDATA	$2, $2
	0x01c0 00448 (test.go:16)	MOVQ	8(CX), CX
	0x01c4 00452 (test.go:16)	MOVQ	AX, ""..autotmp_14+160(SP)
	0x01cc 00460 (test.go:16)	MOVQ	CX, ""..autotmp_14+168(SP)
	0x01d4 00468 (test.go:16)	MOVQ	DX, ""..autotmp_14+176(SP)
	0x01dc 00476 (test.go:16)	PCDATA	$2, $6
	0x01dc 00476 (test.go:16)	MOVQ	"".&stu+104(SP), DI
	0x01e1 00481 (test.go:16)	MOVQ	CX, 8(DI)
	0x01e5 00485 (test.go:16)	MOVQ	DX, 16(DI)
	0x01e9 00489 (test.go:16)	PCDATA	$2, $-2
	0x01e9 00489 (test.go:16)	PCDATA	$0, $-2
	0x01e9 00489 (test.go:16)	CMPL	runtime.writeBarrier(SB), $0
	0x01f0 00496 (test.go:16)	JEQ	503
	0x01f2 00498 (test.go:16)	JMP	670
	0x01f7 00503 (test.go:16)	MOVQ	AX, (DI)
	0x01fa 00506 (test.go:16)	JMP	508
	0x01fc 00508 (test.go:17)	PCDATA	$2, $2
	0x01fc 00508 (test.go:17)	PCDATA	$0, $11
	0x01fc 00508 (test.go:17)	MOVQ	"".&stu+104(SP), AX
	0x0201 00513 (test.go:17)	PCDATA	$2, $0
	0x0201 00513 (test.go:17)	PCDATA	$0, $12
	0x0201 00513 (test.go:17)	MOVQ	AX, ""..autotmp_5+80(SP)
	0x0206 00518 (test.go:17)	PCDATA	$2, $2
	0x0206 00518 (test.go:17)	MOVQ	"".&stu+104(SP), AX
	0x020b 00523 (test.go:17)	PCDATA	$2, $5
	0x020b 00523 (test.go:17)	MOVQ	(AX), CX
	0x020e 00526 (test.go:17)	PCDATA	$2, $4
	0x020e 00526 (test.go:17)	MOVQ	8(AX), AX
	0x0212 00530 (test.go:17)	PCDATA	$2, $0
	0x0212 00530 (test.go:17)	MOVQ	CX, 16(SP)
	0x0217 00535 (test.go:17)	MOVQ	AX, 24(SP)
	0x021c 00540 (test.go:17)	PCDATA	$2, $2
	0x021c 00540 (test.go:17)	LEAQ	type.map[string]*"".stu(SB), AX
	0x0223 00547 (test.go:17)	PCDATA	$2, $0
	0x0223 00547 (test.go:17)	MOVQ	AX, (SP)
	0x0227 00551 (test.go:17)	PCDATA	$2, $4
	0x0227 00551 (test.go:17)	MOVQ	"".m+56(SP), CX
	0x022c 00556 (test.go:17)	PCDATA	$2, $0
	0x022c 00556 (test.go:17)	MOVQ	CX, 8(SP)
	0x0231 00561 (test.go:17)	CALL	runtime.mapassign_faststr(SB)
	0x0236 00566 (test.go:17)	PCDATA	$2, $1
	0x0236 00566 (test.go:17)	MOVQ	32(SP), DI
	0x023b 00571 (test.go:17)	MOVQ	DI, ""..autotmp_15+88(SP)
	0x0240 00576 (test.go:17)	TESTB	AL, (DI)
	0x0242 00578 (test.go:17)	PCDATA	$2, $6
	0x0242 00578 (test.go:17)	PCDATA	$0, $11
	0x0242 00578 (test.go:17)	MOVQ	""..autotmp_5+80(SP), AX
	0x0247 00583 (test.go:17)	PCDATA	$2, $-2
	0x0247 00583 (test.go:17)	PCDATA	$0, $-2
	0x0247 00583 (test.go:17)	CMPL	runtime.writeBarrier(SB), $0
	0x024e 00590 (test.go:17)	JEQ	594
	0x0250 00592 (test.go:17)	JMP	663
	0x0252 00594 (test.go:17)	MOVQ	AX, (DI)
	0x0255 00597 (test.go:17)	JMP	599
	0x0257 00599 (test.go:17)	PCDATA	$2, $0
	0x0257 00599 (test.go:17)	PCDATA	$0, $11
	0x0257 00599 (test.go:17)	JMP	601
	0x0259 00601 (test.go:16)	MOVQ	""..autotmp_11+48(SP), CX
	0x025e 00606 (test.go:16)	INCQ	CX
	0x0261 00609 (test.go:16)	MOVQ	CX, ""..autotmp_11+48(SP)
	0x0266 00614 (test.go:16)	MOVQ	""..autotmp_12+40(SP), DX
	0x026b 00619 (test.go:16)	CMPQ	CX, DX
	0x026e 00622 (test.go:16)	JLT	626
	0x0270 00624 (test.go:16)	JMP	645
	0x0272 00626 (test.go:16)	PCDATA	$2, $4
	0x0272 00626 (test.go:16)	PCDATA	$0, $9
	0x0272 00626 (test.go:16)	MOVQ	""..autotmp_13+96(SP), CX
	0x0277 00631 (test.go:16)	ADDQ	$24, CX
	0x027b 00635 (test.go:16)	PCDATA	$2, $0
	0x027b 00635 (test.go:16)	PCDATA	$0, $11
	0x027b 00635 (test.go:16)	MOVQ	CX, ""..autotmp_13+96(SP)
	0x0280 00640 (test.go:16)	JMP	434
	0x0285 00645 (<unknown line number>)	PCDATA	$2, $-2
	0x0285 00645 (<unknown line number>)	PCDATA	$0, $-2
	0x0285 00645 (<unknown line number>)	JMP	647
	0x0287 00647 (<unknown line number>)	PCDATA	$2, $0
	0x0287 00647 (<unknown line number>)	PCDATA	$0, $13
	0x0287 00647 (<unknown line number>)	MOVQ	512(SP), BP
	0x028f 00655 (<unknown line number>)	ADDQ	$520, SP
	0x0296 00662 (<unknown line number>)	RET
	0x0297 00663 (test.go:17)	PCDATA	$2, $-2
	0x0297 00663 (test.go:17)	PCDATA	$0, $-2
	0x0297 00663 (test.go:17)	CALL	runtime.gcWriteBarrier(SB)
	0x029c 00668 (test.go:17)	JMP	599
	0x029e 00670 (test.go:16)	CALL	runtime.gcWriteBarrier(SB)
	0x02a3 00675 (test.go:16)	JMP	508
	0x02a8 00680 (test.go:16)	JMP	647
	0x02aa 00682 (test.go:16)	NOP
	0x02aa 00682 (test.go:8)	PCDATA	$0, $-1
	0x02aa 00682 (test.go:8)	PCDATA	$2, $-1
	0x02aa 00682 (test.go:8)	CALL	runtime.morestack_noctxt(SB)
	0x02af 00687 (test.go:8)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 8d 84 24 78 fe ff  eH..%....H..$x..
	0x0010 ff 48 3b 41 10 0f 86 8f 02 00 00 48 81 ec 08 02  .H;A.......H....
	0x0020 00 00 48 89 ac 24 00 02 00 00 48 8d ac 24 00 02  ..H..$....H..$..
	0x0030 00 00 0f 57 c9 0f 11 8c 24 b8 00 00 00 0f 57 c9  ...W....$.....W.
	0x0040 0f 11 8c 24 c8 00 00 00 0f 57 c9 0f 11 8c 24 d8  ...$.....W....$.
	0x0050 00 00 00 48 8d bc 24 30 01 00 00 0f 57 c0 48 8d  ...H..$0....W.H.
	0x0060 7f d0 48 89 6c 24 f0 48 8d 6c 24 f0 e8 00 00 00  ..H.l$.H.l$.....
	0x0070 00 48 8b 6d 00 48 8d 84 24 b8 00 00 00 84 00 48  .H.m.H..$......H
	0x0080 8d 84 24 30 01 00 00 48 89 84 24 c8 00 00 00 48  ..$0...H..$....H
	0x0090 8d 84 24 b8 00 00 00 48 89 44 24 48 e8 00 00 00  ..$....H.D$H....
	0x00a0 00 48 8b 44 24 48 84 00 8b 0c 24 89 48 0c 48 8d  .H.D$H....$.H.H.
	0x00b0 84 24 b8 00 00 00 48 89 44 24 38 48 8d 84 24 e8  .$....H.D$8H..$.
	0x00c0 00 00 00 48 89 44 24 40 84 00 48 8b 05 00 00 00  ...H.D$@..H.....
	0x00d0 00 48 89 84 24 e8 00 00 00 0f 10 0d 00 00 00 00  .H..$...........
	0x00e0 0f 11 8c 24 f0 00 00 00 0f 10 0d 00 00 00 00 0f  ...$............
	0x00f0 11 8c 24 00 01 00 00 0f 10 0d 00 00 00 00 0f 11  ..$.............
	0x0100 8c 24 10 01 00 00 0f 10 0d 00 00 00 00 0f 11 8c  .$..............
	0x0110 24 20 01 00 00 48 8b 44 24 40 84 00 eb 00 48 89  $ ...H.D$@....H.
	0x0120 44 24 70 48 c7 44 24 78 03 00 00 00 48 c7 84 24  D$pH.D$x....H..$
	0x0130 80 00 00 00 03 00 00 00 48 8d 05 00 00 00 00 48  ........H......H
	0x0140 89 04 24 e8 00 00 00 00 48 8b 44 24 08 48 89 44  ..$.....H.D$.H.D
	0x0150 24 68 48 8b 84 24 80 00 00 00 48 8b 4c 24 78 48  $hH..$....H.L$xH
	0x0160 8b 54 24 70 48 89 94 24 88 00 00 00 48 89 8c 24  .T$pH..$....H..$
	0x0170 90 00 00 00 48 89 84 24 98 00 00 00 48 c7 44 24  ....H..$....H.D$
	0x0180 30 00 00 00 00 48 8b 84 24 90 00 00 00 48 89 44  0....H..$....H.D
	0x0190 24 28 48 8b 84 24 88 00 00 00 48 89 44 24 60 48  $(H..$....H.D$`H
	0x01a0 8b 44 24 28 48 39 44 24 30 7c 05 e9 f8 00 00 00  .D$(H9D$0|......
	0x01b0 eb 00 48 8b 4c 24 60 84 01 48 8b 51 10 48 8b 01  ..H.L$`..H.Q.H..
	0x01c0 48 8b 49 08 48 89 84 24 a0 00 00 00 48 89 8c 24  H.I.H..$....H..$
	0x01d0 a8 00 00 00 48 89 94 24 b0 00 00 00 48 8b 7c 24  ....H..$....H.|$
	0x01e0 68 48 89 4f 08 48 89 57 10 83 3d 00 00 00 00 00  hH.O.H.W..=.....
	0x01f0 74 05 e9 a7 00 00 00 48 89 07 eb 00 48 8b 44 24  t......H....H.D$
	0x0200 68 48 89 44 24 50 48 8b 44 24 68 48 8b 08 48 8b  hH.D$PH.D$hH..H.
	0x0210 40 08 48 89 4c 24 10 48 89 44 24 18 48 8d 05 00  @.H.L$.H.D$.H...
	0x0220 00 00 00 48 89 04 24 48 8b 4c 24 38 48 89 4c 24  ...H..$H.L$8H.L$
	0x0230 08 e8 00 00 00 00 48 8b 7c 24 20 48 89 7c 24 58  ......H.|$ H.|$X
	0x0240 84 07 48 8b 44 24 50 83 3d 00 00 00 00 00 74 02  ..H.D$P.=.....t.
	0x0250 eb 45 48 89 07 eb 00 eb 00 48 8b 4c 24 30 48 ff  .EH......H.L$0H.
	0x0260 c1 48 89 4c 24 30 48 8b 54 24 28 48 39 d1 7c 02  .H.L$0H.T$(H9.|.
	0x0270 eb 13 48 8b 4c 24 60 48 83 c1 18 48 89 4c 24 60  ..H.L$`H...H.L$`
	0x0280 e9 2d ff ff ff eb 00 48 8b ac 24 00 02 00 00 48  .-.....H..$....H
	0x0290 81 c4 08 02 00 00 c3 e8 00 00 00 00 eb b9 e8 00  ................
	0x02a0 00 00 00 e9 54 ff ff ff eb dd e8 00 00 00 00 e9  ....T...........
	0x02b0 4c fd ff ff                                      L...
	rel 5+4 t=16 TLS+0
	rel 109+4 t=8 runtime.duffzero+239
	rel 157+4 t=8 runtime.fastrand+0
	rel 205+4 t=15 "".statictmp_0+0
	rel 220+4 t=15 "".statictmp_0+8
	rel 235+4 t=15 "".statictmp_0+24
	rel 250+4 t=15 "".statictmp_0+40
	rel 265+4 t=15 "".statictmp_0+56
	rel 315+4 t=15 type."".stu+0
	rel 324+4 t=8 runtime.newobject+0
	rel 491+4 t=15 runtime.writeBarrier+-1
	rel 543+4 t=15 type.map[string]*"".stu+0
	rel 562+4 t=8 runtime.mapassign_faststr+0
	rel 585+4 t=15 runtime.writeBarrier+-1
	rel 664+4 t=8 runtime.gcWriteBarrier+0
	rel 671+4 t=8 runtime.gcWriteBarrier+0
	rel 683+4 t=8 runtime.morestack_noctxt+0
"".init STEXT size=95 args=0x0 locals=0x8
	0x0000 00000 (<autogenerated>:1)	TEXT	"".init(SB), $8-0
	0x0000 00000 (<autogenerated>:1)	MOVQ	(TLS), CX
	0x0009 00009 (<autogenerated>:1)	CMPQ	SP, 16(CX)
	0x000d 00013 (<autogenerated>:1)	JLS	88
	0x000f 00015 (<autogenerated>:1)	SUBQ	$8, SP
	0x0013 00019 (<autogenerated>:1)	MOVQ	BP, (SP)
	0x0017 00023 (<autogenerated>:1)	LEAQ	(SP), BP
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$3, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (<autogenerated>:1)	PCDATA	$2, $0
	0x001b 00027 (<autogenerated>:1)	PCDATA	$0, $0
	0x001b 00027 (<autogenerated>:1)	CMPB	"".initdone·(SB), $1
	0x0022 00034 (<autogenerated>:1)	JHI	38
	0x0024 00036 (<autogenerated>:1)	JMP	47
	0x0026 00038 (<autogenerated>:1)	PCDATA	$2, $-2
	0x0026 00038 (<autogenerated>:1)	PCDATA	$0, $-2
	0x0026 00038 (<autogenerated>:1)	MOVQ	(SP), BP
	0x002a 00042 (<autogenerated>:1)	ADDQ	$8, SP
	0x002e 00046 (<autogenerated>:1)	RET
	0x002f 00047 (<autogenerated>:1)	PCDATA	$2, $0
	0x002f 00047 (<autogenerated>:1)	PCDATA	$0, $0
	0x002f 00047 (<autogenerated>:1)	CMPB	"".initdone·(SB), $1
	0x0036 00054 (<autogenerated>:1)	JEQ	58
	0x0038 00056 (<autogenerated>:1)	JMP	65
	0x003a 00058 (<autogenerated>:1)	CALL	runtime.throwinit(SB)
	0x003f 00063 (<autogenerated>:1)	UNDEF
	0x0041 00065 (<autogenerated>:1)	MOVB	$1, "".initdone·(SB)
	0x0048 00072 (<autogenerated>:1)	MOVB	$2, "".initdone·(SB)
	0x004f 00079 (<autogenerated>:1)	MOVQ	(SP), BP
	0x0053 00083 (<autogenerated>:1)	ADDQ	$8, SP
	0x0057 00087 (<autogenerated>:1)	RET
	0x0058 00088 (<autogenerated>:1)	NOP
	0x0058 00088 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0058 00088 (<autogenerated>:1)	PCDATA	$2, $-1
	0x0058 00088 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x005d 00093 (<autogenerated>:1)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 49 48  eH..%....H;a.vIH
	0x0010 83 ec 08 48 89 2c 24 48 8d 2c 24 80 3d 00 00 00  ...H.,$H.,$.=...
	0x0020 00 01 77 02 eb 09 48 8b 2c 24 48 83 c4 08 c3 80  ..w...H.,$H.....
	0x0030 3d 00 00 00 00 01 74 02 eb 07 e8 00 00 00 00 0f  =.....t.........
	0x0040 0b c6 05 00 00 00 00 01 c6 05 00 00 00 00 02 48  ...............H
	0x0050 8b 2c 24 48 83 c4 08 c3 e8 00 00 00 00 eb a1     .,$H...........
	rel 5+4 t=16 TLS+0
	rel 29+4 t=15 "".initdone·+-1
	rel 49+4 t=15 "".initdone·+-1
	rel 59+4 t=8 runtime.throwinit+0
	rel 67+4 t=15 "".initdone·+-1
	rel 74+4 t=15 "".initdone·+-1
	rel 89+4 t=8 runtime.morestack_noctxt+0
type..hash."".stu STEXT dupok size=139 args=0x18 locals=0x28
	0x0000 00000 (<autogenerated>:1)	TEXT	type..hash."".stu(SB), DUPOK, $40-24
	0x0000 00000 (<autogenerated>:1)	MOVQ	(TLS), CX
	0x0009 00009 (<autogenerated>:1)	CMPQ	SP, 16(CX)
	0x000d 00013 (<autogenerated>:1)	JLS	129
	0x000f 00015 (<autogenerated>:1)	SUBQ	$40, SP
	0x0013 00019 (<autogenerated>:1)	MOVQ	BP, 32(SP)
	0x0018 00024 (<autogenerated>:1)	LEAQ	32(SP), BP
	0x001d 00029 (<autogenerated>:1)	FUNCDATA	$0, gclocals·1a65e721a2ccc325b382662e7ffee780(SB)
	0x001d 00029 (<autogenerated>:1)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x001d 00029 (<autogenerated>:1)	FUNCDATA	$3, gclocals·1cf923758aae2e428391d1783fe59973(SB)
	0x001d 00029 (<autogenerated>:1)	PCDATA	$2, $0
	0x001d 00029 (<autogenerated>:1)	PCDATA	$0, $0
	0x001d 00029 (<autogenerated>:1)	MOVQ	$0, "".~r2+64(SP)
	0x0026 00038 (<autogenerated>:1)	PCDATA	$2, $1
	0x0026 00038 (<autogenerated>:1)	MOVQ	"".p+48(SP), AX
	0x002b 00043 (<autogenerated>:1)	PCDATA	$2, $0
	0x002b 00043 (<autogenerated>:1)	MOVQ	AX, (SP)
	0x002f 00047 (<autogenerated>:1)	MOVQ	"".h+56(SP), AX
	0x0034 00052 (<autogenerated>:1)	MOVQ	AX, 8(SP)
	0x0039 00057 (<autogenerated>:1)	CALL	runtime.strhash(SB)
	0x003e 00062 (<autogenerated>:1)	MOVQ	16(SP), AX
	0x0043 00067 (<autogenerated>:1)	MOVQ	AX, "".h+56(SP)
	0x0048 00072 (<autogenerated>:1)	PCDATA	$2, $2
	0x0048 00072 (<autogenerated>:1)	PCDATA	$0, $1
	0x0048 00072 (<autogenerated>:1)	MOVQ	"".p+48(SP), CX
	0x004d 00077 (<autogenerated>:1)	ADDQ	$16, CX
	0x0051 00081 (<autogenerated>:1)	PCDATA	$2, $0
	0x0051 00081 (<autogenerated>:1)	MOVQ	CX, (SP)
	0x0055 00085 (<autogenerated>:1)	MOVQ	AX, 8(SP)
	0x005a 00090 (<autogenerated>:1)	MOVQ	$8, 16(SP)
	0x0063 00099 (<autogenerated>:1)	CALL	runtime.memhash(SB)
	0x0068 00104 (<autogenerated>:1)	MOVQ	24(SP), AX
	0x006d 00109 (<autogenerated>:1)	MOVQ	AX, "".h+56(SP)
	0x0072 00114 (<autogenerated>:1)	MOVQ	AX, "".~r2+64(SP)
	0x0077 00119 (<autogenerated>:1)	MOVQ	32(SP), BP
	0x007c 00124 (<autogenerated>:1)	ADDQ	$40, SP
	0x0080 00128 (<autogenerated>:1)	RET
	0x0081 00129 (<autogenerated>:1)	NOP
	0x0081 00129 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0081 00129 (<autogenerated>:1)	PCDATA	$2, $-1
	0x0081 00129 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x0086 00134 (<autogenerated>:1)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 72 48  eH..%....H;a.vrH
	0x0010 83 ec 28 48 89 6c 24 20 48 8d 6c 24 20 48 c7 44  ..(H.l$ H.l$ H.D
	0x0020 24 40 00 00 00 00 48 8b 44 24 30 48 89 04 24 48  $@....H.D$0H..$H
	0x0030 8b 44 24 38 48 89 44 24 08 e8 00 00 00 00 48 8b  .D$8H.D$......H.
	0x0040 44 24 10 48 89 44 24 38 48 8b 4c 24 30 48 83 c1  D$.H.D$8H.L$0H..
	0x0050 10 48 89 0c 24 48 89 44 24 08 48 c7 44 24 10 08  .H..$H.D$.H.D$..
	0x0060 00 00 00 e8 00 00 00 00 48 8b 44 24 18 48 89 44  ........H.D$.H.D
	0x0070 24 38 48 89 44 24 40 48 8b 6c 24 20 48 83 c4 28  $8H.D$@H.l$ H..(
	0x0080 c3 e8 00 00 00 00 e9 75 ff ff ff                 .......u...
	rel 5+4 t=16 TLS+0
	rel 58+4 t=8 runtime.strhash+0
	rel 100+4 t=8 runtime.memhash+0
	rel 130+4 t=8 runtime.morestack_noctxt+0
type..eq."".stu STEXT dupok size=192 args=0x18 locals=0x48
	0x0000 00000 (<autogenerated>:1)	TEXT	type..eq."".stu(SB), DUPOK, $72-24
	0x0000 00000 (<autogenerated>:1)	MOVQ	(TLS), CX
	0x0009 00009 (<autogenerated>:1)	CMPQ	SP, 16(CX)
	0x000d 00013 (<autogenerated>:1)	JLS	182
	0x0013 00019 (<autogenerated>:1)	SUBQ	$72, SP
	0x0017 00023 (<autogenerated>:1)	MOVQ	BP, 64(SP)
	0x001c 00028 (<autogenerated>:1)	LEAQ	64(SP), BP
	0x0021 00033 (<autogenerated>:1)	FUNCDATA	$0, gclocals·0e66f3f2491221b77362069379e5c720(SB)
	0x0021 00033 (<autogenerated>:1)	FUNCDATA	$1, gclocals·91d432edea9e3c468c5aec7a805d99d2(SB)
	0x0021 00033 (<autogenerated>:1)	FUNCDATA	$3, gclocals·6e8d7ea4abad763909b26991048ee1fe(SB)
	0x0021 00033 (<autogenerated>:1)	PCDATA	$2, $0
	0x0021 00033 (<autogenerated>:1)	PCDATA	$0, $0
	0x0021 00033 (<autogenerated>:1)	MOVB	$0, "".~r2+96(SP)
	0x0026 00038 (<autogenerated>:1)	PCDATA	$2, $1
	0x0026 00038 (<autogenerated>:1)	MOVQ	"".p+80(SP), AX
	0x002b 00043 (<autogenerated>:1)	MOVQ	8(AX), CX
	0x002f 00047 (<autogenerated>:1)	MOVQ	(AX), AX
	0x0032 00050 (<autogenerated>:1)	PCDATA	$2, $0
	0x0032 00050 (<autogenerated>:1)	PCDATA	$0, $1
	0x0032 00050 (<autogenerated>:1)	MOVQ	AX, ""..autotmp_3+48(SP)
	0x0037 00055 (<autogenerated>:1)	MOVQ	CX, ""..autotmp_3+56(SP)
	0x003c 00060 (<autogenerated>:1)	PCDATA	$2, $1
	0x003c 00060 (<autogenerated>:1)	MOVQ	"".q+88(SP), AX
	0x0041 00065 (<autogenerated>:1)	MOVQ	8(AX), CX
	0x0045 00069 (<autogenerated>:1)	MOVQ	(AX), AX
	0x0048 00072 (<autogenerated>:1)	MOVQ	AX, ""..autotmp_4+32(SP)
	0x004d 00077 (<autogenerated>:1)	MOVQ	CX, ""..autotmp_4+40(SP)
	0x0052 00082 (<autogenerated>:1)	MOVQ	""..autotmp_3+56(SP), DX
	0x0057 00087 (<autogenerated>:1)	CMPQ	DX, CX
	0x005a 00090 (<autogenerated>:1)	SETEQ	CL
	0x005d 00093 (<autogenerated>:1)	JEQ	97
	0x005f 00095 (<autogenerated>:1)	JMP	178
	0x0061 00097 (<autogenerated>:1)	PCDATA	$2, $2
	0x0061 00097 (<autogenerated>:1)	MOVQ	""..autotmp_3+48(SP), CX
	0x0066 00102 (<autogenerated>:1)	PCDATA	$2, $1
	0x0066 00102 (<autogenerated>:1)	MOVQ	CX, (SP)
	0x006a 00106 (<autogenerated>:1)	PCDATA	$2, $0
	0x006a 00106 (<autogenerated>:1)	MOVQ	AX, 8(SP)
	0x006f 00111 (<autogenerated>:1)	PCDATA	$0, $0
	0x006f 00111 (<autogenerated>:1)	MOVQ	""..autotmp_3+56(SP), AX
	0x0074 00116 (<autogenerated>:1)	MOVQ	AX, 16(SP)
	0x0079 00121 (<autogenerated>:1)	CALL	runtime.memequal(SB)
	0x007e 00126 (<autogenerated>:1)	MOVBLZX	24(SP), AX
	0x0083 00131 (<autogenerated>:1)	JMP	133
	0x0085 00133 (<autogenerated>:1)	TESTB	AL, AL
	0x0087 00135 (<autogenerated>:1)	JNE	139
	0x0089 00137 (<autogenerated>:1)	JMP	176
	0x008b 00139 (<autogenerated>:1)	PCDATA	$2, $1
	0x008b 00139 (<autogenerated>:1)	PCDATA	$0, $2
	0x008b 00139 (<autogenerated>:1)	MOVQ	"".q+88(SP), AX
	0x0090 00144 (<autogenerated>:1)	PCDATA	$2, $0
	0x0090 00144 (<autogenerated>:1)	MOVQ	16(AX), AX
	0x0094 00148 (<autogenerated>:1)	PCDATA	$2, $3
	0x0094 00148 (<autogenerated>:1)	PCDATA	$0, $3
	0x0094 00148 (<autogenerated>:1)	MOVQ	"".p+80(SP), CX
	0x0099 00153 (<autogenerated>:1)	PCDATA	$2, $0
	0x0099 00153 (<autogenerated>:1)	CMPQ	16(CX), AX
	0x009d 00157 (<autogenerated>:1)	SETEQ	AL
	0x00a0 00160 (<autogenerated>:1)	JMP	162
	0x00a2 00162 (<autogenerated>:1)	MOVB	AL, "".~r2+96(SP)
	0x00a6 00166 (<autogenerated>:1)	MOVQ	64(SP), BP
	0x00ab 00171 (<autogenerated>:1)	ADDQ	$72, SP
	0x00af 00175 (<autogenerated>:1)	RET
	0x00b0 00176 (<autogenerated>:1)	PCDATA	$2, $-2
	0x00b0 00176 (<autogenerated>:1)	PCDATA	$0, $-2
	0x00b0 00176 (<autogenerated>:1)	JMP	162
	0x00b2 00178 (<autogenerated>:1)	PCDATA	$2, $0
	0x00b2 00178 (<autogenerated>:1)	PCDATA	$0, $0
	0x00b2 00178 (<autogenerated>:1)	MOVL	CX, AX
	0x00b4 00180 (<autogenerated>:1)	JMP	133
	0x00b6 00182 (<autogenerated>:1)	NOP
	0x00b6 00182 (<autogenerated>:1)	PCDATA	$0, $-1
	0x00b6 00182 (<autogenerated>:1)	PCDATA	$2, $-1
	0x00b6 00182 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x00bb 00187 (<autogenerated>:1)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 0f 86 a3  eH..%....H;a....
	0x0010 00 00 00 48 83 ec 48 48 89 6c 24 40 48 8d 6c 24  ...H..HH.l$@H.l$
	0x0020 40 c6 44 24 60 00 48 8b 44 24 50 48 8b 48 08 48  @.D$`.H.D$PH.H.H
	0x0030 8b 00 48 89 44 24 30 48 89 4c 24 38 48 8b 44 24  ..H.D$0H.L$8H.D$
	0x0040 58 48 8b 48 08 48 8b 00 48 89 44 24 20 48 89 4c  XH.H.H..H.D$ H.L
	0x0050 24 28 48 8b 54 24 38 48 39 ca 0f 94 c1 74 02 eb  $(H.T$8H9....t..
	0x0060 51 48 8b 4c 24 30 48 89 0c 24 48 89 44 24 08 48  QH.L$0H..$H.D$.H
	0x0070 8b 44 24 38 48 89 44 24 10 e8 00 00 00 00 0f b6  .D$8H.D$........
	0x0080 44 24 18 eb 00 84 c0 75 02 eb 25 48 8b 44 24 58  D$.....u..%H.D$X
	0x0090 48 8b 40 10 48 8b 4c 24 50 48 39 41 10 0f 94 c0  H.@.H.L$PH9A....
	0x00a0 eb 00 88 44 24 60 48 8b 6c 24 40 48 83 c4 48 c3  ...D$`H.l$@H..H.
	0x00b0 eb f0 89 c8 eb cf e8 00 00 00 00 e9 40 ff ff ff  ............@...
	rel 5+4 t=16 TLS+0
	rel 122+4 t=8 runtime.memequal+0
	rel 183+4 t=8 runtime.morestack_noctxt+0
type..hash.[3]"".stu STEXT dupok size=175 args=0x18 locals=0x38
	0x0000 00000 (<autogenerated>:1)	TEXT	type..hash.[3]"".stu(SB), DUPOK, $56-24
	0x0000 00000 (<autogenerated>:1)	MOVQ	(TLS), CX
	0x0009 00009 (<autogenerated>:1)	CMPQ	SP, 16(CX)
	0x000d 00013 (<autogenerated>:1)	JLS	165
	0x0013 00019 (<autogenerated>:1)	SUBQ	$56, SP
	0x0017 00023 (<autogenerated>:1)	MOVQ	BP, 48(SP)
	0x001c 00028 (<autogenerated>:1)	LEAQ	48(SP), BP
	0x0021 00033 (<autogenerated>:1)	FUNCDATA	$0, gclocals·1a65e721a2ccc325b382662e7ffee780(SB)
	0x0021 00033 (<autogenerated>:1)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x0021 00033 (<autogenerated>:1)	FUNCDATA	$3, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
	0x0021 00033 (<autogenerated>:1)	PCDATA	$2, $0
	0x0021 00033 (<autogenerated>:1)	PCDATA	$0, $0
	0x0021 00033 (<autogenerated>:1)	MOVQ	$0, "".~r2+80(SP)
	0x002a 00042 (<autogenerated>:1)	MOVQ	$0, ""..autotmp_4+40(SP)
	0x0033 00051 (<autogenerated>:1)	MOVQ	$3, ""..autotmp_5+32(SP)
	0x003c 00060 (<autogenerated>:1)	JMP	62
	0x003e 00062 (<autogenerated>:1)	MOVQ	""..autotmp_5+32(SP), AX
	0x0043 00067 (<autogenerated>:1)	CMPQ	""..autotmp_4+40(SP), AX
	0x0048 00072 (<autogenerated>:1)	JLT	76
	0x004a 00074 (<autogenerated>:1)	JMP	145
	0x004c 00076 (<autogenerated>:1)	MOVQ	""..autotmp_4+40(SP), AX
	0x0051 00081 (<autogenerated>:1)	MOVQ	AX, "".i+24(SP)
	0x0056 00086 (<autogenerated>:1)	LEAQ	(AX)(AX*2), AX
	0x005a 00090 (<autogenerated>:1)	SHLQ	$3, AX
	0x005e 00094 (<autogenerated>:1)	PCDATA	$2, $1
	0x005e 00094 (<autogenerated>:1)	ADDQ	"".p+64(SP), AX
	0x0063 00099 (<autogenerated>:1)	PCDATA	$2, $0
	0x0063 00099 (<autogenerated>:1)	MOVQ	AX, (SP)
	0x0067 00103 (<autogenerated>:1)	MOVQ	"".h+72(SP), AX
	0x006c 00108 (<autogenerated>:1)	MOVQ	AX, 8(SP)
	0x0071 00113 (<autogenerated>:1)	CALL	type..hash."".stu(SB)
	0x0076 00118 (<autogenerated>:1)	MOVQ	16(SP), AX
	0x007b 00123 (<autogenerated>:1)	MOVQ	AX, "".h+72(SP)
	0x0080 00128 (<autogenerated>:1)	JMP	130
	0x0082 00130 (<autogenerated>:1)	MOVQ	""..autotmp_4+40(SP), AX
	0x0087 00135 (<autogenerated>:1)	INCQ	AX
	0x008a 00138 (<autogenerated>:1)	MOVQ	AX, ""..autotmp_4+40(SP)
	0x008f 00143 (<autogenerated>:1)	JMP	62
	0x0091 00145 (<autogenerated>:1)	PCDATA	$0, $1
	0x0091 00145 (<autogenerated>:1)	MOVQ	"".h+72(SP), AX
	0x0096 00150 (<autogenerated>:1)	MOVQ	AX, "".~r2+80(SP)
	0x009b 00155 (<autogenerated>:1)	MOVQ	48(SP), BP
	0x00a0 00160 (<autogenerated>:1)	ADDQ	$56, SP
	0x00a4 00164 (<autogenerated>:1)	RET
	0x00a5 00165 (<autogenerated>:1)	NOP
	0x00a5 00165 (<autogenerated>:1)	PCDATA	$0, $-1
	0x00a5 00165 (<autogenerated>:1)	PCDATA	$2, $-1
	0x00a5 00165 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x00aa 00170 (<autogenerated>:1)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 0f 86 92  eH..%....H;a....
	0x0010 00 00 00 48 83 ec 38 48 89 6c 24 30 48 8d 6c 24  ...H..8H.l$0H.l$
	0x0020 30 48 c7 44 24 50 00 00 00 00 48 c7 44 24 28 00  0H.D$P....H.D$(.
	0x0030 00 00 00 48 c7 44 24 20 03 00 00 00 eb 00 48 8b  ...H.D$ ......H.
	0x0040 44 24 20 48 39 44 24 28 7c 02 eb 45 48 8b 44 24  D$ H9D$(|..EH.D$
	0x0050 28 48 89 44 24 18 48 8d 04 40 48 c1 e0 03 48 03  (H.D$.H..@H...H.
	0x0060 44 24 40 48 89 04 24 48 8b 44 24 48 48 89 44 24  D$@H..$H.D$HH.D$
	0x0070 08 e8 00 00 00 00 48 8b 44 24 10 48 89 44 24 48  ......H.D$.H.D$H
	0x0080 eb 00 48 8b 44 24 28 48 ff c0 48 89 44 24 28 eb  ..H.D$(H..H.D$(.
	0x0090 ad 48 8b 44 24 48 48 89 44 24 50 48 8b 6c 24 30  .H.D$HH.D$PH.l$0
	0x00a0 48 83 c4 38 c3 e8 00 00 00 00 e9 51 ff ff ff     H..8.......Q...
	rel 5+4 t=16 TLS+0
	rel 114+4 t=8 type..hash."".stu+0
	rel 166+4 t=8 runtime.morestack_noctxt+0
type..eq.[3]"".stu STEXT dupok size=332 args=0x18 locals=0x60
	0x0000 00000 (<autogenerated>:1)	TEXT	type..eq.[3]"".stu(SB), DUPOK, $96-24
	0x0000 00000 (<autogenerated>:1)	MOVQ	(TLS), CX
	0x0009 00009 (<autogenerated>:1)	CMPQ	SP, 16(CX)
	0x000d 00013 (<autogenerated>:1)	JLS	322
	0x0013 00019 (<autogenerated>:1)	SUBQ	$96, SP
	0x0017 00023 (<autogenerated>:1)	MOVQ	BP, 88(SP)
	0x001c 00028 (<autogenerated>:1)	LEAQ	88(SP), BP
	0x0021 00033 (<autogenerated>:1)	FUNCDATA	$0, gclocals·5e326a2f5498e528ce8fbe7bd86e7d21(SB)
	0x0021 00033 (<autogenerated>:1)	FUNCDATA	$1, gclocals·02e4573554fd483091afa0dd5654ac56(SB)
	0x0021 00033 (<autogenerated>:1)	FUNCDATA	$3, gclocals·39c35ed93fc31b84b0c48f70fc1139f4(SB)
	0x0021 00033 (<autogenerated>:1)	PCDATA	$2, $0
	0x0021 00033 (<autogenerated>:1)	PCDATA	$0, $0
	0x0021 00033 (<autogenerated>:1)	MOVB	$0, "".~r2+120(SP)
	0x0026 00038 (<autogenerated>:1)	MOVQ	$0, ""..autotmp_4+48(SP)
	0x002f 00047 (<autogenerated>:1)	MOVQ	$3, ""..autotmp_5+40(SP)
	0x0038 00056 (<autogenerated>:1)	JMP	58
	0x003a 00058 (<autogenerated>:1)	MOVQ	""..autotmp_5+40(SP), AX
	0x003f 00063 (<autogenerated>:1)	CMPQ	""..autotmp_4+48(SP), AX
	0x0044 00068 (<autogenerated>:1)	JLT	75
	0x0046 00070 (<autogenerated>:1)	JMP	307
	0x004b 00075 (<autogenerated>:1)	MOVQ	""..autotmp_4+48(SP), AX
	0x0050 00080 (<autogenerated>:1)	MOVQ	AX, "".i+32(SP)
	0x0055 00085 (<autogenerated>:1)	LEAQ	(AX)(AX*2), AX
	0x0059 00089 (<autogenerated>:1)	SHLQ	$3, AX
	0x005d 00093 (<autogenerated>:1)	PCDATA	$2, $1
	0x005d 00093 (<autogenerated>:1)	ADDQ	"".p+104(SP), AX
	0x0062 00098 (<autogenerated>:1)	PCDATA	$2, $2
	0x0062 00098 (<autogenerated>:1)	MOVQ	(AX), CX
	0x0065 00101 (<autogenerated>:1)	PCDATA	$2, $3
	0x0065 00101 (<autogenerated>:1)	MOVQ	8(AX), AX
	0x0069 00105 (<autogenerated>:1)	PCDATA	$2, $0
	0x0069 00105 (<autogenerated>:1)	PCDATA	$0, $1
	0x0069 00105 (<autogenerated>:1)	MOVQ	CX, ""..autotmp_6+72(SP)
	0x006e 00110 (<autogenerated>:1)	MOVQ	AX, ""..autotmp_6+80(SP)
	0x0073 00115 (<autogenerated>:1)	MOVQ	"".i+32(SP), AX
	0x0078 00120 (<autogenerated>:1)	LEAQ	(AX)(AX*2), AX
	0x007c 00124 (<autogenerated>:1)	SHLQ	$3, AX
	0x0080 00128 (<autogenerated>:1)	PCDATA	$2, $1
	0x0080 00128 (<autogenerated>:1)	ADDQ	"".q+112(SP), AX
	0x0085 00133 (<autogenerated>:1)	MOVQ	8(AX), CX
	0x0089 00137 (<autogenerated>:1)	MOVQ	(AX), AX
	0x008c 00140 (<autogenerated>:1)	MOVQ	AX, ""..autotmp_7+56(SP)
	0x0091 00145 (<autogenerated>:1)	MOVQ	CX, ""..autotmp_7+64(SP)
	0x0096 00150 (<autogenerated>:1)	MOVQ	""..autotmp_6+80(SP), DX
	0x009b 00155 (<autogenerated>:1)	CMPQ	DX, CX
	0x009e 00158 (<autogenerated>:1)	SETNE	CL
	0x00a1 00161 (<autogenerated>:1)	JNE	165
	0x00a3 00163 (<autogenerated>:1)	JMP	265
	0x00a5 00165 (<autogenerated>:1)	PCDATA	$2, $-2
	0x00a5 00165 (<autogenerated>:1)	PCDATA	$0, $-2
	0x00a5 00165 (<autogenerated>:1)	JMP	167
	0x00a7 00167 (<autogenerated>:1)	PCDATA	$2, $0
	0x00a7 00167 (<autogenerated>:1)	PCDATA	$0, $0
	0x00a7 00167 (<autogenerated>:1)	TESTB	CL, CL
	0x00a9 00169 (<autogenerated>:1)	JNE	173
	0x00ab 00171 (<autogenerated>:1)	JMP	216
	0x00ad 00173 (<autogenerated>:1)	PCDATA	$2, $-2
	0x00ad 00173 (<autogenerated>:1)	PCDATA	$0, $-2
	0x00ad 00173 (<autogenerated>:1)	JMP	175
	0x00af 00175 (<autogenerated>:1)	PCDATA	$2, $0
	0x00af 00175 (<autogenerated>:1)	PCDATA	$0, $0
	0x00af 00175 (<autogenerated>:1)	TESTB	CL, CL
	0x00b1 00177 (<autogenerated>:1)	JNE	181
	0x00b3 00179 (<autogenerated>:1)	JMP	196
	0x00b5 00181 (<autogenerated>:1)	PCDATA	$0, $2
	0x00b5 00181 (<autogenerated>:1)	MOVB	$0, "".~r2+120(SP)
	0x00ba 00186 (<autogenerated>:1)	MOVQ	88(SP), BP
	0x00bf 00191 (<autogenerated>:1)	ADDQ	$96, SP
	0x00c3 00195 (<autogenerated>:1)	RET
	0x00c4 00196 (<autogenerated>:1)	PCDATA	$2, $-2
	0x00c4 00196 (<autogenerated>:1)	PCDATA	$0, $-2
	0x00c4 00196 (<autogenerated>:1)	JMP	198
	0x00c6 00198 (<autogenerated>:1)	PCDATA	$2, $0
	0x00c6 00198 (<autogenerated>:1)	PCDATA	$0, $0
	0x00c6 00198 (<autogenerated>:1)	MOVQ	""..autotmp_4+48(SP), AX
	0x00cb 00203 (<autogenerated>:1)	INCQ	AX
	0x00ce 00206 (<autogenerated>:1)	MOVQ	AX, ""..autotmp_4+48(SP)
	0x00d3 00211 (<autogenerated>:1)	JMP	58
	0x00d8 00216 (<autogenerated>:1)	MOVQ	"".i+32(SP), AX
	0x00dd 00221 (<autogenerated>:1)	LEAQ	(AX)(AX*2), AX
	0x00e1 00225 (<autogenerated>:1)	SHLQ	$3, AX
	0x00e5 00229 (<autogenerated>:1)	PCDATA	$2, $1
	0x00e5 00229 (<autogenerated>:1)	ADDQ	"".p+104(SP), AX
	0x00ea 00234 (<autogenerated>:1)	MOVQ	"".i+32(SP), DX
	0x00ef 00239 (<autogenerated>:1)	LEAQ	(DX)(DX*2), DX
	0x00f3 00243 (<autogenerated>:1)	SHLQ	$3, DX
	0x00f7 00247 (<autogenerated>:1)	PCDATA	$2, $4
	0x00f7 00247 (<autogenerated>:1)	ADDQ	"".q+112(SP), DX
	0x00fc 00252 (<autogenerated>:1)	PCDATA	$2, $1
	0x00fc 00252 (<autogenerated>:1)	MOVQ	16(DX), DX
	0x0100 00256 (<autogenerated>:1)	PCDATA	$2, $0
	0x0100 00256 (<autogenerated>:1)	CMPQ	16(AX), DX
	0x0104 00260 (<autogenerated>:1)	SETNE	CL
	0x0107 00263 (<autogenerated>:1)	JMP	175
	0x0109 00265 (<autogenerated>:1)	PCDATA	$2, $2
	0x0109 00265 (<autogenerated>:1)	PCDATA	$0, $1
	0x0109 00265 (<autogenerated>:1)	MOVQ	""..autotmp_6+72(SP), CX
	0x010e 00270 (<autogenerated>:1)	PCDATA	$2, $1
	0x010e 00270 (<autogenerated>:1)	MOVQ	CX, (SP)
	0x0112 00274 (<autogenerated>:1)	PCDATA	$2, $0
	0x0112 00274 (<autogenerated>:1)	MOVQ	AX, 8(SP)
	0x0117 00279 (<autogenerated>:1)	PCDATA	$0, $0
	0x0117 00279 (<autogenerated>:1)	MOVQ	""..autotmp_6+80(SP), AX
	0x011c 00284 (<autogenerated>:1)	MOVQ	AX, 16(SP)
	0x0121 00289 (<autogenerated>:1)	CALL	runtime.memequal(SB)
	0x0126 00294 (<autogenerated>:1)	MOVBLZX	24(SP), CX
	0x012b 00299 (<autogenerated>:1)	XORL	$1, CX
	0x012e 00302 (<autogenerated>:1)	JMP	167
	0x0133 00307 (<autogenerated>:1)	PCDATA	$0, $2
	0x0133 00307 (<autogenerated>:1)	MOVB	$1, "".~r2+120(SP)
	0x0138 00312 (<autogenerated>:1)	MOVQ	88(SP), BP
	0x013d 00317 (<autogenerated>:1)	ADDQ	$96, SP
	0x0141 00321 (<autogenerated>:1)	RET
	0x0142 00322 (<autogenerated>:1)	NOP
	0x0142 00322 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0142 00322 (<autogenerated>:1)	PCDATA	$2, $-1
	0x0142 00322 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x0147 00327 (<autogenerated>:1)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 0f 86 2f  eH..%....H;a.../
	0x0010 01 00 00 48 83 ec 60 48 89 6c 24 58 48 8d 6c 24  ...H..`H.l$XH.l$
	0x0020 58 c6 44 24 78 00 48 c7 44 24 30 00 00 00 00 48  X.D$x.H.D$0....H
	0x0030 c7 44 24 28 03 00 00 00 eb 00 48 8b 44 24 28 48  .D$(......H.D$(H
	0x0040 39 44 24 30 7c 05 e9 e8 00 00 00 48 8b 44 24 30  9D$0|......H.D$0
	0x0050 48 89 44 24 20 48 8d 04 40 48 c1 e0 03 48 03 44  H.D$ H..@H...H.D
	0x0060 24 68 48 8b 08 48 8b 40 08 48 89 4c 24 48 48 89  $hH..H.@.H.L$HH.
	0x0070 44 24 50 48 8b 44 24 20 48 8d 04 40 48 c1 e0 03  D$PH.D$ H..@H...
	0x0080 48 03 44 24 70 48 8b 48 08 48 8b 00 48 89 44 24  H.D$pH.H.H..H.D$
	0x0090 38 48 89 4c 24 40 48 8b 54 24 50 48 39 ca 0f 95  8H.L$@H.T$PH9...
	0x00a0 c1 75 02 eb 64 eb 00 84 c9 75 02 eb 2b eb 00 84  .u..d....u..+...
	0x00b0 c9 75 02 eb 0f c6 44 24 78 00 48 8b 6c 24 58 48  .u....D$x.H.l$XH
	0x00c0 83 c4 60 c3 eb 00 48 8b 44 24 30 48 ff c0 48 89  ..`...H.D$0H..H.
	0x00d0 44 24 30 e9 62 ff ff ff 48 8b 44 24 20 48 8d 04  D$0.b...H.D$ H..
	0x00e0 40 48 c1 e0 03 48 03 44 24 68 48 8b 54 24 20 48  @H...H.D$hH.T$ H
	0x00f0 8d 14 52 48 c1 e2 03 48 03 54 24 70 48 8b 52 10  ..RH...H.T$pH.R.
	0x0100 48 39 50 10 0f 95 c1 eb a6 48 8b 4c 24 48 48 89  H9P......H.L$HH.
	0x0110 0c 24 48 89 44 24 08 48 8b 44 24 50 48 89 44 24  .$H.D$.H.D$PH.D$
	0x0120 10 e8 00 00 00 00 0f b6 4c 24 18 83 f1 01 e9 74  ........L$.....t
	0x0130 ff ff ff c6 44 24 78 01 48 8b 6c 24 58 48 83 c4  ....D$x.H.l$XH..
	0x0140 60 c3 e8 00 00 00 00 e9 b4 fe ff ff              `...........
	rel 5+4 t=16 TLS+0
	rel 290+4 t=8 runtime.memequal+0
	rel 323+4 t=8 runtime.morestack_noctxt+0
go.string."zhou" SRODATA dupok size=4
	0x0000 7a 68 6f 75                                      zhou
go.string."li" SRODATA dupok size=2
	0x0000 6c 69                                            li
go.string."wang" SRODATA dupok size=4
	0x0000 77 61 6e 67                                      wang
go.loc."".main SDWARFLOC size=0
go.info."".main SDWARFINFO size=81
	0x0000 02 22 22 2e 6d 61 69 6e 00 00 00 00 00 00 00 00  ."".main........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01  ................
	0x0020 09 6d 00 09 00 00 00 00 03 91 a8 7c 09 73 74 75  .m.........|.stu
	0x0030 73 00 0a 00 00 00 00 03 91 e0 7c 13 00 00 00 00  s.........|.....
	0x0040 09 26 73 74 75 00 10 00 00 00 00 03 91 d8 7c 00  .&stu.........|.
	0x0050 00                                               .
	rel 9+8 t=1 "".main+0
	rel 17+8 t=1 "".main+692
	rel 27+4 t=29 gofile../Users/zhangyong/GoProjects/goPrac/test/test.go+0
	rel 36+4 t=28 go.info.map[string]*"".stu+0
	rel 51+4 t=28 go.info.[]"".stu+0
	rel 60+4 t=28 go.range."".main+0
	rel 71+4 t=28 go.info.*"".stu+0
go.range."".main SDWARFRANGE size=64
	0x0000 ff ff ff ff ff ff ff ff 00 00 00 00 00 00 00 00  ................
	0x0010 38 01 00 00 00 00 00 00 85 02 00 00 00 00 00 00  8...............
	0x0020 97 02 00 00 00 00 00 00 aa 02 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 8+8 t=1 "".main+0
go.isstmt."".main SDWARFMISC size=0
	0x0000 04 1b 04 17 03 03 01 86 01 02 08 01 75 02 07 01  ............u...
	0x0010 73 02 05 01 45 02 05 01 58 02 05 01 4c 02 0a 00  s...E...X...L...
go.loc."".init SDWARFLOC size=0
go.info."".init SDWARFINFO size=33
	0x0000 02 22 22 2e 69 6e 69 74 00 00 00 00 00 00 00 00  ."".init........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01  ................
	0x0020 00                                               .
	rel 9+8 t=1 "".init+0
	rel 17+8 t=1 "".init+95
	rel 27+4 t=29 gofile..<autogenerated>+0
go.range."".init SDWARFRANGE size=0
go.isstmt."".init SDWARFMISC size=0
	0x0000 04 0f 04 0c 03 07 01 04 02 09 01 10 02 09 01 10  ................
	0x0010 02 07 00                                         ...
"".statictmp_0 SRODATA size=72
	0x0000 00 00 00 00 00 00 00 00 04 00 00 00 00 00 00 00  ................
	0x0010 18 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 02 00 00 00 00 00 00 00 17 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 04 00 00 00 00 00 00 00  ................
	0x0040 16 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 go.string."zhou"+0
	rel 24+8 t=1 go.string."li"+0
	rel 48+8 t=1 go.string."wang"+0
"".initdone· SNOPTRBSS size=1
go.loc.type..hash."".stu SDWARFLOC dupok size=0
go.info.type..hash."".stu SDWARFINFO dupok size=80
	0x0000 02 74 79 70 65 2e 2e 68 61 73 68 2e 22 22 2e 73  .type..hash."".s
	0x0010 74 75 00 00 00 00 00 00 00 00 00 00 00 00 00 00  tu..............
	0x0020 00 00 00 01 9c 00 00 00 00 01 0e 70 00 00 01 00  ...........p....
	0x0030 00 00 00 01 9c 0e 68 00 00 01 00 00 00 00 02 91  ......h.........
	0x0040 08 0e 7e 72 32 00 01 01 00 00 00 00 02 91 10 00  ..~r2...........
	rel 19+8 t=1 type..hash."".stu+0
	rel 27+8 t=1 type..hash."".stu+139
	rel 37+4 t=29 gofile..<autogenerated>+0
	rel 47+4 t=28 go.info.*"".stu+0
	rel 58+4 t=28 go.info.uintptr+0
	rel 72+4 t=28 go.info.uintptr+0
go.range.type..hash."".stu SDWARFRANGE dupok size=0
go.isstmt.type..hash."".stu SDWARFMISC dupok size=0
	0x0000 04 0f 04 0e 03 09 01 5b 02 0a 00                 .......[...
go.loc.type..eq."".stu SDWARFLOC dupok size=0
go.info.type..eq."".stu SDWARFINFO dupok size=78
	0x0000 02 74 79 70 65 2e 2e 65 71 2e 22 22 2e 73 74 75  .type..eq."".stu
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 01 9c 00 00 00 00 01 0e 70 00 00 01 00 00 00  .........p......
	0x0030 00 01 9c 0e 71 00 00 01 00 00 00 00 02 91 08 0e  ....q...........
	0x0040 7e 72 32 00 01 01 00 00 00 00 02 91 10 00        ~r2...........
	rel 17+8 t=1 type..eq."".stu+0
	rel 25+8 t=1 type..eq."".stu+192
	rel 35+4 t=29 gofile..<autogenerated>+0
	rel 45+4 t=28 go.info.*"".stu+0
	rel 56+4 t=28 go.info.*"".stu+0
	rel 70+4 t=28 go.info.bool+0
go.range.type..eq."".stu SDWARFRANGE dupok size=0
go.isstmt.type..eq."".stu SDWARFMISC dupok size=0
	0x0000 04 13 04 0e 03 05 01 5f 02 02 01 04 02 05 01 12  ......._........
	0x0010 02 04 01 0a 02 02 01 04 02 0a 00                 ...........
type..hashfunc."".stu SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 type..hash."".stu+0
type..eqfunc."".stu SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 type..eq."".stu+0
type..alg."".stu SRODATA dupok size=16
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 0+8 t=1 type..hashfunc."".stu+0
	rel 8+8 t=1 type..eqfunc."".stu+0
runtime.gcbits.01 SRODATA dupok size=1
	0x0000 01                                               .
type..namedata.*main.stu- SRODATA dupok size=12
	0x0000 00 00 09 2a 6d 61 69 6e 2e 73 74 75              ...*main.stu
type.*"".stu SRODATA size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 80 5f 30 0a 00 08 08 36 00 00 00 00 00 00 00 00  ._0....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*main.stu-+0
	rel 48+8 t=1 type."".stu+0
type..namedata.Name. SRODATA dupok size=7
	0x0000 01 00 04 4e 61 6d 65                             ...Name
type..namedata.Age. SRODATA dupok size=6
	0x0000 01 00 03 41 67 65                                ...Age
type."".stu SRODATA size=144
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 0d ea c4 87 07 08 08 19 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 02 00 00 00 00 00 00 00 02 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 40 00 00 00 00 00 00 00  ........@.......
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0070 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0080 00 00 00 00 00 00 00 00 20 00 00 00 00 00 00 00  ........ .......
	rel 24+8 t=1 type..alg."".stu+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*main.stu-+0
	rel 44+4 t=5 type.*"".stu+0
	rel 56+8 t=1 type."".stu+96
	rel 80+4 t=5 type..importpath."".+0
	rel 96+8 t=1 type..namedata.Name.+0
	rel 104+8 t=1 type.string+0
	rel 120+8 t=1 type..namedata.Age.+0
	rel 128+8 t=1 type.int+0
type..namedata.**main.stu- SRODATA dupok size=13
	0x0000 00 00 0a 2a 2a 6d 61 69 6e 2e 73 74 75           ...**main.stu
type.**"".stu SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 71 7d b4 00 00 08 08 36 00 00 00 00 00 00 00 00  q}.....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.**main.stu-+0
	rel 48+8 t=1 type.*"".stu+0
type..namedata.*[]main.stu- SRODATA dupok size=14
	0x0000 00 00 0b 2a 5b 5d 6d 61 69 6e 2e 73 74 75        ...*[]main.stu
type.*[]"".stu SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 4b 79 1b 91 00 08 08 36 00 00 00 00 00 00 00 00  Ky.....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]main.stu-+0
	rel 48+8 t=1 type.[]"".stu+0
type.[]"".stu SRODATA dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 47 f2 0d 7e 02 08 08 17 00 00 00 00 00 00 00 00  G..~............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]main.stu-+0
	rel 44+4 t=6 type.*[]"".stu+0
	rel 48+8 t=1 type."".stu+0
go.loc.type..hash.[3]"".stu SDWARFLOC dupok size=0
go.info.type..hash.[3]"".stu SDWARFINFO dupok size=94
	0x0000 02 74 79 70 65 2e 2e 68 61 73 68 2e 5b 33 5d 22  .type..hash.[3]"
	0x0010 22 2e 73 74 75 00 00 00 00 00 00 00 00 00 00 00  ".stu...........
	0x0020 00 00 00 00 00 00 01 9c 00 00 00 00 01 09 69 00  ..............i.
	0x0030 01 00 00 00 00 02 91 58 0e 70 00 00 01 00 00 00  .......X.p......
	0x0040 00 01 9c 0e 68 00 00 01 00 00 00 00 02 91 08 0e  ....h...........
	0x0050 7e 72 32 00 01 01 00 00 00 00 02 91 10 00        ~r2...........
	rel 22+8 t=1 type..hash.[3]"".stu+0
	rel 30+8 t=1 type..hash.[3]"".stu+175
	rel 40+4 t=29 gofile..<autogenerated>+0
	rel 49+4 t=28 go.info.int+0
	rel 61+4 t=28 go.info.*[3]"".stu+0
	rel 72+4 t=28 go.info.uintptr+0
	rel 86+4 t=28 go.info.uintptr+0
go.range.type..hash.[3]"".stu SDWARFRANGE dupok size=0
go.isstmt.type..hash.[3]"".stu SDWARFMISC dupok size=0
	0x0000 04 13 04 0e 03 09 01 14 02 05 01 4e 02 05 01 0f  ...........N....
	0x0010 02 0a 00                                         ...
go.loc.type..eq.[3]"".stu SDWARFLOC dupok size=0
go.info.type..eq.[3]"".stu SDWARFINFO dupok size=93
	0x0000 02 74 79 70 65 2e 2e 65 71 2e 5b 33 5d 22 22 2e  .type..eq.[3]"".
	0x0010 73 74 75 00 00 00 00 00 00 00 00 00 00 00 00 00  stu.............
	0x0020 00 00 00 00 01 9c 00 00 00 00 01 09 69 00 01 00  ............i...
	0x0030 00 00 00 03 91 b8 7f 0e 70 00 00 01 00 00 00 00  ........p.......
	0x0040 01 9c 0e 71 00 00 01 00 00 00 00 02 91 08 0e 7e  ...q...........~
	0x0050 72 32 00 01 01 00 00 00 00 02 91 10 00           r2...........
	rel 20+8 t=1 type..eq.[3]"".stu+0
	rel 28+8 t=1 type..eq.[3]"".stu+332
	rel 38+4 t=29 gofile..<autogenerated>+0
	rel 47+4 t=28 go.info.int+0
	rel 60+4 t=28 go.info.*[3]"".stu+0
	rel 71+4 t=28 go.info.*[3]"".stu+0
	rel 85+4 t=28 go.info.bool+0
go.range.type..eq.[3]"".stu SDWARFRANGE dupok size=0
go.isstmt.type..eq.[3]"".stu SDWARFMISC dupok size=0
	0x0000 04 13 04 0e 03 05 01 14 02 05 01 68 02 02 01 04  ...........h....
	0x0010 02 04 01 04 02 05 01 0c 02 05 01 0d 02 05 01 56  ...............V
	0x0020 02 05 01 0a 02 0a 00                             .......
type..hashfunc.[3]"".stu SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 type..hash.[3]"".stu+0
type..eqfunc.[3]"".stu SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 type..eq.[3]"".stu+0
type..alg.[3]"".stu SRODATA dupok size=16
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 0+8 t=1 type..hashfunc.[3]"".stu+0
	rel 8+8 t=1 type..eqfunc.[3]"".stu+0
runtime.gcbits.49 SRODATA dupok size=1
	0x0000 49                                               I
type..namedata.*[3]main.stu- SRODATA dupok size=15
	0x0000 00 00 0c 2a 5b 33 5d 6d 61 69 6e 2e 73 74 75     ...*[3]main.stu
type.[3]"".stu SRODATA dupok size=72
	0x0000 48 00 00 00 00 00 00 00 38 00 00 00 00 00 00 00  H.......8.......
	0x0010 f9 98 b4 04 02 08 08 11 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 03 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 type..alg.[3]"".stu+0
	rel 32+8 t=1 runtime.gcbits.49+0
	rel 40+4 t=5 type..namedata.*[3]main.stu-+0
	rel 44+4 t=6 type.*[3]"".stu+0
	rel 48+8 t=1 type."".stu+0
	rel 56+8 t=1 type.[]"".stu+0
type.*[3]"".stu SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 28 f5 17 51 00 08 08 36 00 00 00 00 00 00 00 00  (..Q...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[3]main.stu-+0
	rel 48+8 t=1 type.[3]"".stu+0
type..namedata.*[]uint8- SRODATA dupok size=11
	0x0000 00 00 08 2a 5b 5d 75 69 6e 74 38                 ...*[]uint8
type.*[]uint8 SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 a5 8e d0 69 00 08 08 36 00 00 00 00 00 00 00 00  ...i...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]uint8-+0
	rel 48+8 t=1 type.[]uint8+0
type.[]uint8 SRODATA dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 df 7e 2e 38 02 08 08 17 00 00 00 00 00 00 00 00  .~.8............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]uint8-+0
	rel 44+4 t=6 type.*[]uint8+0
	rel 48+8 t=1 type.uint8+0
type..namedata.*[8]uint8- SRODATA dupok size=12
	0x0000 00 00 09 2a 5b 38 5d 75 69 6e 74 38              ...*[8]uint8
type.*[8]uint8 SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 a9 89 a5 7a 00 08 08 36 00 00 00 00 00 00 00 00  ...z...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[8]uint8-+0
	rel 48+8 t=1 type.[8]uint8+0
runtime.gcbits. SRODATA dupok size=0
type.[8]uint8 SRODATA dupok size=72
	0x0000 08 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 3e f9 30 b4 02 01 01 91 00 00 00 00 00 00 00 00  >.0.............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 08 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.+0
	rel 40+4 t=5 type..namedata.*[8]uint8-+0
	rel 44+4 t=6 type.*[8]uint8+0
	rel 48+8 t=1 type.uint8+0
	rel 56+8 t=1 type.[]uint8+0
type..namedata.*[]string- SRODATA dupok size=12
	0x0000 00 00 09 2a 5b 5d 73 74 72 69 6e 67              ...*[]string
type.*[]string SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 92 22 76 84 00 08 08 36 00 00 00 00 00 00 00 00  ."v....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]string-+0
	rel 48+8 t=1 type.[]string+0
type.[]string SRODATA dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 d3 a8 f3 0a 02 08 08 17 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]string-+0
	rel 44+4 t=6 type.*[]string+0
	rel 48+8 t=1 type.string+0
type..namedata.*[8]string- SRODATA dupok size=13
	0x0000 00 00 0a 2a 5b 38 5d 73 74 72 69 6e 67           ...*[8]string
type.*[8]string SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 ad 94 14 6f 00 08 08 36 00 00 00 00 00 00 00 00  ...o...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[8]string-+0
	rel 48+8 t=1 type.noalg.[8]string+0
runtime.gcbits.5555 SRODATA dupok size=2
	0x0000 55 55                                            UU
type.noalg.[8]string SRODATA dupok size=72
	0x0000 80 00 00 00 00 00 00 00 78 00 00 00 00 00 00 00  ........x.......
	0x0010 55 53 8c 3e 02 08 08 11 00 00 00 00 00 00 00 00  US.>............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 08 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.5555+0
	rel 40+4 t=5 type..namedata.*[8]string-+0
	rel 44+4 t=6 type.*[8]string+0
	rel 48+8 t=1 type.string+0
	rel 56+8 t=1 type.[]string+0
type..namedata.*[]*main.stu- SRODATA dupok size=15
	0x0000 00 00 0c 2a 5b 5d 2a 6d 61 69 6e 2e 73 74 75     ...*[]*main.stu
type.*[]*"".stu SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 2a c4 4d aa 00 08 08 36 00 00 00 00 00 00 00 00  *.M....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]*main.stu-+0
	rel 48+8 t=1 type.[]*"".stu+0
type.[]*"".stu SRODATA dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 12 47 93 a8 02 08 08 17 00 00 00 00 00 00 00 00  .G..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]*main.stu-+0
	rel 44+4 t=6 type.*[]*"".stu+0
	rel 48+8 t=1 type.*"".stu+0
type..namedata.*[8]*main.stu- SRODATA dupok size=16
	0x0000 00 00 0d 2a 5b 38 5d 2a 6d 61 69 6e 2e 73 74 75  ...*[8]*main.stu
type.*[8]*"".stu SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 31 83 bc 59 00 08 08 36 00 00 00 00 00 00 00 00  1..Y...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[8]*main.stu-+0
	rel 48+8 t=1 type.noalg.[8]*"".stu+0
runtime.gcbits.ff SRODATA dupok size=1
	0x0000 ff                                               .
type.noalg.[8]*"".stu SRODATA dupok size=72
	0x0000 40 00 00 00 00 00 00 00 40 00 00 00 00 00 00 00  @.......@.......
	0x0010 af 8c 5f 6a 02 08 08 11 00 00 00 00 00 00 00 00  .._j............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 08 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.ff+0
	rel 40+4 t=5 type..namedata.*[8]*main.stu-+0
	rel 44+4 t=6 type.*[8]*"".stu+0
	rel 48+8 t=1 type.*"".stu+0
	rel 56+8 t=1 type.[]*"".stu+0
runtime.gcbits.aaaafe03 SRODATA dupok size=4
	0x0000 aa aa fe 03                                      ....
type..namedata.*map.bucket[string]*main.stu- SRODATA dupok size=31
	0x0000 00 00 1c 2a 6d 61 70 2e 62 75 63 6b 65 74 5b 73  ...*map.bucket[s
	0x0010 74 72 69 6e 67 5d 2a 6d 61 69 6e 2e 73 74 75     tring]*main.stu
type..importpath.. SRODATA dupok size=3
	0x0000 00 00 00                                         ...
type..namedata.topbits- SRODATA dupok size=10
	0x0000 00 00 07 74 6f 70 62 69 74 73                    ...topbits
type..namedata.keys- SRODATA dupok size=7
	0x0000 00 00 04 6b 65 79 73                             ...keys
type..namedata.values- SRODATA dupok size=9
	0x0000 00 00 06 76 61 6c 75 65 73                       ...values
type..namedata.overflow- SRODATA dupok size=11
	0x0000 00 00 08 6f 76 65 72 66 6c 6f 77                 ...overflow
type.noalg.map.bucket[string]*"".stu SRODATA dupok size=176
	0x0000 d0 00 00 00 00 00 00 00 d0 00 00 00 00 00 00 00  ................
	0x0010 9b 6c 59 f4 02 08 08 19 00 00 00 00 00 00 00 00  .lY.............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 04 00 00 00 00 00 00 00 04 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0070 00 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0080 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0090 10 01 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x00a0 00 00 00 00 00 00 00 00 90 01 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.aaaafe03+0
	rel 40+4 t=5 type..namedata.*map.bucket[string]*main.stu-+0
	rel 44+4 t=6 type.*map.bucket[string]*"".stu+0
	rel 48+8 t=1 type..importpath..+0
	rel 56+8 t=1 type.noalg.map.bucket[string]*"".stu+80
	rel 80+8 t=1 type..namedata.topbits-+0
	rel 88+8 t=1 type.[8]uint8+0
	rel 104+8 t=1 type..namedata.keys-+0
	rel 112+8 t=1 type.noalg.[8]string+0
	rel 128+8 t=1 type..namedata.values-+0
	rel 136+8 t=1 type.noalg.[8]*"".stu+0
	rel 152+8 t=1 type..namedata.overflow-+0
	rel 160+8 t=1 type.*map.bucket[string]*"".stu+0
type.*map.bucket[string]*"".stu SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 64 4f 5b e1 00 08 08 36 00 00 00 00 00 00 00 00  dO[....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*map.bucket[string]*main.stu-+0
	rel 48+8 t=1 type.noalg.map.bucket[string]*"".stu+0
runtime.gcbits.2c SRODATA dupok size=1
	0x0000 2c                                               ,
type..namedata.*map.hdr[string]*main.stu- SRODATA dupok size=28
	0x0000 00 00 19 2a 6d 61 70 2e 68 64 72 5b 73 74 72 69  ...*map.hdr[stri
	0x0010 6e 67 5d 2a 6d 61 69 6e 2e 73 74 75              ng]*main.stu
type..namedata.count- SRODATA dupok size=8
	0x0000 00 00 05 63 6f 75 6e 74                          ...count
type..namedata.flags- SRODATA dupok size=8
	0x0000 00 00 05 66 6c 61 67 73                          ...flags
type..namedata.B. SRODATA dupok size=4
	0x0000 01 00 01 42                                      ...B
type..namedata.noverflow- SRODATA dupok size=12
	0x0000 00 00 09 6e 6f 76 65 72 66 6c 6f 77              ...noverflow
type..namedata.hash0- SRODATA dupok size=8
	0x0000 00 00 05 68 61 73 68 30                          ...hash0
type..namedata.buckets- SRODATA dupok size=10
	0x0000 00 00 07 62 75 63 6b 65 74 73                    ...buckets
type..namedata.oldbuckets- SRODATA dupok size=13
	0x0000 00 00 0a 6f 6c 64 62 75 63 6b 65 74 73           ...oldbuckets
type..namedata.nevacuate- SRODATA dupok size=12
	0x0000 00 00 09 6e 65 76 61 63 75 61 74 65              ...nevacuate
type..namedata.extra- SRODATA dupok size=8
	0x0000 00 00 05 65 78 74 72 61                          ...extra
type.noalg.map.hdr[string]*"".stu SRODATA dupok size=296
	0x0000 30 00 00 00 00 00 00 00 30 00 00 00 00 00 00 00  0.......0.......
	0x0010 17 f6 f7 d3 02 08 08 19 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 09 00 00 00 00 00 00 00 09 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0070 00 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0080 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0090 12 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x00a0 00 00 00 00 00 00 00 00 14 00 00 00 00 00 00 00  ................
	0x00b0 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x00c0 18 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x00d0 00 00 00 00 00 00 00 00 20 00 00 00 00 00 00 00  ........ .......
	0x00e0 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x00f0 30 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  0...............
	0x0100 00 00 00 00 00 00 00 00 40 00 00 00 00 00 00 00  ........@.......
	0x0110 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0120 50 00 00 00 00 00 00 00                          P.......
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.2c+0
	rel 40+4 t=5 type..namedata.*map.hdr[string]*main.stu-+0
	rel 44+4 t=6 type.*map.hdr[string]*"".stu+0
	rel 48+8 t=1 type..importpath..+0
	rel 56+8 t=1 type.noalg.map.hdr[string]*"".stu+80
	rel 80+8 t=1 type..namedata.count-+0
	rel 88+8 t=1 type.int+0
	rel 104+8 t=1 type..namedata.flags-+0
	rel 112+8 t=1 type.uint8+0
	rel 128+8 t=1 type..namedata.B.+0
	rel 136+8 t=1 type.uint8+0
	rel 152+8 t=1 type..namedata.noverflow-+0
	rel 160+8 t=1 type.uint16+0
	rel 176+8 t=1 type..namedata.hash0-+0
	rel 184+8 t=1 type.uint32+0
	rel 200+8 t=1 type..namedata.buckets-+0
	rel 208+8 t=1 type.*map.bucket[string]*"".stu+0
	rel 224+8 t=1 type..namedata.oldbuckets-+0
	rel 232+8 t=1 type.*map.bucket[string]*"".stu+0
	rel 248+8 t=1 type..namedata.nevacuate-+0
	rel 256+8 t=1 type.uintptr+0
	rel 272+8 t=1 type..namedata.extra-+0
	rel 280+8 t=1 type.unsafe.Pointer+0
type.*map.hdr[string]*"".stu SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 01 35 d4 9a 00 08 08 36 00 00 00 00 00 00 00 00  .5.....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*map.hdr[string]*main.stu-+0
	rel 48+8 t=1 type.noalg.map.hdr[string]*"".stu+0
type..namedata.*map[string]*main.stu- SRODATA dupok size=24
	0x0000 00 00 15 2a 6d 61 70 5b 73 74 72 69 6e 67 5d 2a  ...*map[string]*
	0x0010 6d 61 69 6e 2e 73 74 75                          main.stu
type.*map[string]*"".stu SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 cf 44 b3 60 00 08 08 36 00 00 00 00 00 00 00 00  .D.`...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*map[string]*main.stu-+0
	rel 48+8 t=1 type.map[string]*"".stu+0
type.map[string]*"".stu SRODATA dupok size=80
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 6e b3 ec 03 02 08 08 35 00 00 00 00 00 00 00 00  n......5........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 10 00 08 00 d0 00 01 01  ................
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*map[string]*main.stu-+0
	rel 44+4 t=6 type.*map[string]*"".stu+0
	rel 48+8 t=1 type.string+0
	rel 56+8 t=1 type.*"".stu+0
	rel 64+8 t=1 type.noalg.map.bucket[string]*"".stu+0
gclocals·f0a67958015464e4cc8847ce0df60843 SRODATA dupok size=8
	0x0000 0e 00 00 00 00 00 00 00                          ........
gclocals·f74a0b8caecc0f0df89956b0bfefa0fe SRODATA dupok size=120
	0x0000 0e 00 00 00 39 00 00 00 00 00 00 00 00 00 00 00  ....9...........
	0x0010 00 00 2c 00 00 00 00 00 00 00 2c 00 55 55 ff 01  ..,.......,.UU..
	0x0020 04 00 2c 00 55 55 ff 01 01 00 2c 00 55 55 ff 01  ..,.UU....,.UU..
	0x0030 01 00 6c 12 55 55 ff 01 03 00 6c 12 55 55 ff 01  ..l.UU....l.UU..
	0x0040 81 00 6c 12 55 55 ff 01 c1 00 6c 12 55 55 ff 01  ..l.UU....l.UU..
	0x0050 41 00 6c 12 55 55 ff 01 41 04 6c 12 55 55 ff 01  A.l.UU..A.l.UU..
	0x0060 61 00 6c 12 55 55 ff 01 69 00 6c 12 55 55 ff 01  a.l.UU..i.l.UU..
	0x0070 00 00 6c 12 55 55 ff 01                          ..l.UU..
gclocals·7f910c4fecabcdd009afbf1e2d5c5c17 SRODATA dupok size=15
	0x0000 07 00 00 00 07 00 00 00 00 40 01 04 02 03 41     .........@....A
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
gclocals·1a65e721a2ccc325b382662e7ffee780 SRODATA dupok size=10
	0x0000 02 00 00 00 01 00 00 00 01 00                    ..........
gclocals·69c1753bd5f81501d95132d08af04464 SRODATA dupok size=8
	0x0000 02 00 00 00 00 00 00 00                          ........
gclocals·1cf923758aae2e428391d1783fe59973 SRODATA dupok size=11
	0x0000 03 00 00 00 02 00 00 00 00 01 02                 ...........
gclocals·0e66f3f2491221b77362069379e5c720 SRODATA dupok size=12
	0x0000 04 00 00 00 02 00 00 00 03 03 01 00              ............
gclocals·91d432edea9e3c468c5aec7a805d99d2 SRODATA dupok size=12
	0x0000 04 00 00 00 04 00 00 00 00 04 00 00              ............
gclocals·6e8d7ea4abad763909b26991048ee1fe SRODATA dupok size=12
	0x0000 04 00 00 00 02 00 00 00 00 01 03 02              ............
gclocals·9fb7f0986f647f17cb53dda1484e0f7a SRODATA dupok size=10
	0x0000 02 00 00 00 01 00 00 00 00 01                    ..........
gclocals·5e326a2f5498e528ce8fbe7bd86e7d21 SRODATA dupok size=11
	0x0000 03 00 00 00 02 00 00 00 03 03 00                 ...........
gclocals·02e4573554fd483091afa0dd5654ac56 SRODATA dupok size=11
	0x0000 03 00 00 00 04 00 00 00 00 04 00                 ...........
gclocals·39c35ed93fc31b84b0c48f70fc1139f4 SRODATA dupok size=13
	0x0000 05 00 00 00 03 00 00 00 00 01 03 02 05           .............
