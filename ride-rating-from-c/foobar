	.section	__TEXT,__text,regular,pure_instructions
	.globl	_ride_ratings_update_all
	.align	4, 0x90
_ride_ratings_update_all:               ## @ride_ratings_update_all
## BB#0:
	push	ebp
	mov	ebp, esp
	sub	esp, 24
	call	L0$pb
L0$pb:
	pop	eax
	mov	ecx, 10349160
	movzx	ecx, byte ptr [ecx]
	and	ecx, 2
	cmp	ecx, 0
	mov	dword ptr [ebp - 4], eax ## 4-byte Spill
	je	LBB0_2
## BB#1:
	jmp	LBB0_9
LBB0_2:
	movzx	eax, byte ptr [20493713]
	mov	ecx, eax
	sub	ecx, 5
	mov	dword ptr [ebp - 8], eax ## 4-byte Spill
	mov	dword ptr [ebp - 12], ecx ## 4-byte Spill
	ja	LBB0_9
## BB#10:
	mov	eax, dword ptr [ebp - 4] ## 4-byte Reload
	mov	ecx, dword ptr [ebp - 8] ## 4-byte Reload
Ltmp0 = LJTI0_0-L0$pb
	mov	edx, dword ptr [eax + 4*ecx + Ltmp0]
	add	edx, eax
	jmp	edx
LBB0_3:
	call	_ride_ratings_update_state_0
	jmp	LBB0_9
LBB0_4:
	call	_ride_ratings_update_state_1
	jmp	LBB0_9
LBB0_5:
	call	_ride_ratings_update_state_2
	jmp	LBB0_9
LBB0_6:
	call	_ride_ratings_update_state_3
	jmp	LBB0_9
LBB0_7:
	call	_ride_ratings_update_state_4
	jmp	LBB0_9
LBB0_8:
	call	_ride_ratings_update_state_5
LBB0_9:
	add	esp, 24
	pop	ebp
	ret
	.align	2, 0x90
L0_0_set_3 = LBB0_3-L0$pb
L0_0_set_4 = LBB0_4-L0$pb
L0_0_set_5 = LBB0_5-L0$pb
L0_0_set_6 = LBB0_6-L0$pb
L0_0_set_7 = LBB0_7-L0$pb
L0_0_set_8 = LBB0_8-L0$pb
LJTI0_0:
	.long	L0_0_set_3
	.long	L0_0_set_4
	.long	L0_0_set_5
	.long	L0_0_set_6
	.long	L0_0_set_7
	.long	L0_0_set_8

	.globl	_main
	.align	4, 0x90
_main:                                  ## @main
## BB#0:
	push	ebp
	mov	ebp, esp
	sub	esp, 8
	call	_ride_ratings_update_state_0
	mov	eax, 0
	add	esp, 8
	pop	ebp
	ret

	.align	4, 0x90
_ride_ratings_update_state_5:           ## @ride_ratings_update_state_5
## BB#0:
	push	ebp
	mov	ebp, esp
	sub	esp, 40
	mov	eax, 7036274
	mov	dword ptr [esp], 7036274
	mov	dword ptr [ebp - 28], eax ## 4-byte Spill
	call	_RCT2_CALLPROC_EBPSAFE
	mov	dword ptr [ebp - 32], eax ## 4-byte Spill
	add	esp, 40
	pop	ebp
	ret

	.align	4, 0x90
_RCT2_CALLPROC_EBPSAFE:                 ## @RCT2_CALLPROC_EBPSAFE
## BB#0:
	push	ebp
	mov	ebp, esp
	sub	esp, 40
	mov	eax, dword ptr [ebp + 8]
	mov	ecx, 3149642683
	mov	dword ptr [ebp - 4], eax
	mov	eax, dword ptr [ebp - 4]
	mov	dword ptr [esp], eax
	mov	dword ptr [esp + 4], -1145324613
	mov	dword ptr [esp + 8], -1145324613
	mov	dword ptr [esp + 12], -1145324613
	mov	dword ptr [esp + 16], -1145324613
	mov	dword ptr [esp + 20], -1145324613
	mov	dword ptr [esp + 24], -1145324613
	mov	dword ptr [esp + 28], -1145324613
	mov	dword ptr [ebp - 8], ecx ## 4-byte Spill
	call	_RCT2_CALLPROC_X
	add	esp, 40
	pop	ebp
	ret

	.align	4, 0x90
_RCT2_CALLPROC_X:                       ## @RCT2_CALLPROC_X
## BB#0:
	push	ebp
	mov	ebp, esp
	push	ebx
	push	edi
	push	esi
	sub	esp, 44
	mov	eax, dword ptr [ebp + 36]
	mov	ecx, dword ptr [ebp + 32]
	mov	edx, dword ptr [ebp + 28]
	mov	esi, dword ptr [ebp + 24]
	mov	edi, dword ptr [ebp + 20]
	mov	ebx, dword ptr [ebp + 16]
	mov	dword ptr [ebp - 52], eax ## 4-byte Spill
	mov	eax, dword ptr [ebp + 12]
	mov	dword ptr [ebp - 56], eax ## 4-byte Spill
	mov	eax, dword ptr [ebp + 8]
	mov	dword ptr [ebp - 20], eax
	mov	eax, dword ptr [ebp - 56] ## 4-byte Reload
	mov	dword ptr [ebp - 24], eax
	mov	dword ptr [ebp - 28], ebx
	mov	dword ptr [ebp - 32], edi
	mov	dword ptr [ebp - 36], esi
	mov	dword ptr [ebp - 40], edx
	mov	dword ptr [ebp - 44], ecx
	mov	ecx, dword ptr [ebp - 52] ## 4-byte Reload
	mov	dword ptr [ebp - 48], ecx
	## InlineAsm Start
		
		push ebx 
		push ebp 
		push -20(%ebp) 	
		mov eax, -24(%ebp) 	
		mov ebx, -28(%ebp) 	
		mov ecx, -32(%ebp) 	
		mov edx, -36(%ebp) 	
		mov esi, -40(%ebp) 	
		mov edi, -44(%ebp) 	
		mov ebp, -48(%ebp) 	
		call [esp] 	
		lahf 
		add esp, 4 	
		pop ebp 
		pop ebx 
	 
	## InlineAsm End
	mov	eax, dword ptr [ebp - 16]
	add	esp, 44
	pop	esi
	pop	edi
	pop	ebx
	pop	ebp
	ret

	.align	4, 0x90
_ride_ratings_update_state_4:           ## @ride_ratings_update_state_4
## BB#0:
	push	ebp
	mov	ebp, esp
	sub	esp, 8
	mov	eax, 20493713
	mov	byte ptr [eax], 5
	call	_loc_6B5BB2
	add	esp, 8
	pop	ebp
	ret

	.align	4, 0x90
_loc_6B5BB2:                            ## @loc_6B5BB2
## BB#0:
	push	ebp
	mov	ebp, esp
	push	ebx
	push	edi
	push	esi
	sub	esp, 36
	call	L6$pb
L6$pb:
	pop	eax
	mov	eax, dword ptr [eax + L_g_ride_list$non_lazy_ptr-L6$pb]
	mov	ecx, 20493712
	movzx	ecx, byte ptr [ecx]
	mov	eax, dword ptr [eax]
	imul	ecx, ecx, 608
	add	eax, ecx
	mov	dword ptr [ebp - 16], eax
	mov	eax, dword ptr [ebp - 16]
	movzx	eax, byte ptr [eax]
	cmp	eax, 255
	je	LBB6_2
## BB#1:
	mov	eax, dword ptr [ebp - 16]
	movzx	eax, byte ptr [eax + 73]
	cmp	eax, 0
	jne	LBB6_3
LBB6_2:
	mov	eax, 20493713
	mov	byte ptr [eax], 0
	jmp	LBB6_14
LBB6_3:
	mov	eax, dword ptr [ebp - 16]
	movzx	eax, byte ptr [eax]
	cmp	eax, 20
	jne	LBB6_5
## BB#4:
	mov	eax, 20493713
	mov	byte ptr [eax], 3
	jmp	LBB6_14
LBB6_5:
	mov	dword ptr [ebp - 20], 0
LBB6_6:                                 ## =>This Inner Loop Header: Depth=1
	cmp	dword ptr [ebp - 20], 4
	jge	LBB6_13
## BB#7:                                ##   in Loop: Header=BB6_6 Depth=1
	mov	eax, dword ptr [ebp - 20]
	mov	ecx, dword ptr [ebp - 16]
	movzx	eax, word ptr [ecx + 2*eax + 82]
	cmp	eax, 65535
	je	LBB6_11
## BB#8:
	mov	eax, 20493774
	movzx	ecx, word ptr [eax]
	and	ecx, 4294967294
	mov	dx, cx
	mov	word ptr [eax], dx
	mov	eax, dword ptr [ebp - 20]
	mov	ecx, dword ptr [ebp - 16]
	movzx	eax, word ptr [ecx + 2*eax + 106]
	cmp	eax, 65535
	jne	LBB6_10
## BB#9:
	mov	eax, 20493774
	movzx	ecx, word ptr [eax]
	or	ecx, 1
	mov	dx, cx
	mov	word ptr [eax], dx
LBB6_10:
	mov	eax, 20493710
	mov	ecx, 20493708
	mov	edx, 20493706
	mov	esi, 20493714
	mov	edi, 20493704
	mov	ebx, 20493702
	mov	dword ptr [ebp - 36], eax ## 4-byte Spill
	mov	eax, 20493700
	mov	dword ptr [ebp - 40], eax ## 4-byte Spill
	mov	eax, dword ptr [ebp - 20]
	mov	dword ptr [ebp - 44], eax ## 4-byte Spill
	mov	eax, dword ptr [ebp - 16]
	mov	dword ptr [ebp - 48], ecx ## 4-byte Spill
	mov	ecx, dword ptr [ebp - 44] ## 4-byte Reload
	movzx	eax, word ptr [eax + 2*ecx + 82]
	and	eax, 255
	shl	eax, 5
	mov	dword ptr [ebp - 24], eax
	mov	eax, dword ptr [ebp - 20]
	mov	ecx, dword ptr [ebp - 16]
	movzx	eax, word ptr [ecx + 2*eax + 82]
	sar	eax, 8
	shl	eax, 5
	mov	dword ptr [ebp - 28], eax
	mov	eax, dword ptr [ebp - 20]
	mov	ecx, dword ptr [ebp - 16]
	movzx	eax, byte ptr [ecx + eax + 90]
	shl	eax, 3
	mov	dword ptr [ebp - 32], eax
	mov	eax, dword ptr [ebp - 24]
                                        ## kill: AX<def> AX<kill> EAX<kill>
	mov	ecx, dword ptr [ebp - 40] ## 4-byte Reload
	mov	word ptr [ecx], ax
	mov	ecx, dword ptr [ebp - 28]
	mov	ax, cx
	mov	word ptr [ebx], ax
	mov	ecx, dword ptr [ebp - 32]
	mov	ax, cx
	mov	word ptr [edi], ax
	mov	byte ptr [esi], -1
	mov	ecx, dword ptr [ebp - 24]
	mov	ax, cx
	mov	word ptr [edx], ax
	mov	ecx, dword ptr [ebp - 28]
	mov	ax, cx
	mov	ecx, dword ptr [ebp - 48] ## 4-byte Reload
	mov	word ptr [ecx], ax
	mov	edx, dword ptr [ebp - 32]
	mov	ax, dx
	mov	edx, dword ptr [ebp - 36] ## 4-byte Reload
	mov	word ptr [edx], ax
	jmp	LBB6_14
LBB6_11:                                ##   in Loop: Header=BB6_6 Depth=1
	jmp	LBB6_12
LBB6_12:                                ##   in Loop: Header=BB6_6 Depth=1
	mov	eax, dword ptr [ebp - 20]
	add	eax, 1
	mov	dword ptr [ebp - 20], eax
	jmp	LBB6_6
LBB6_13:
	mov	eax, 20493713
	mov	byte ptr [eax], 0
LBB6_14:
	add	esp, 36
	pop	esi
	pop	edi
	pop	ebx
	pop	ebp
	ret

	.align	4, 0x90
_ride_ratings_update_state_3:           ## @ride_ratings_update_state_3
## BB#0:
	push	ebp
	mov	ebp, esp
	push	ebx
	push	esi
	sub	esp, 48
	call	L7$pb
L7$pb:
	pop	eax
	mov	eax, dword ptr [eax + L_g_ride_list$non_lazy_ptr-L7$pb]
	mov	ecx, 20493712
	movzx	ecx, byte ptr [ecx]
	mov	eax, dword ptr [eax]
	imul	ecx, ecx, 608
	add	eax, ecx
	mov	dword ptr [ebp - 12], eax
	mov	eax, dword ptr [ebp - 12]
	movzx	eax, byte ptr [eax]
	cmp	eax, 255
	je	LBB7_2
## BB#1:
	mov	eax, dword ptr [ebp - 12]
	movzx	eax, byte ptr [eax + 73]
	cmp	eax, 0
	jne	LBB7_3
LBB7_2:
	mov	eax, 20493713
	mov	byte ptr [eax], 0
	jmp	LBB7_4
LBB7_3:
	mov	eax, 12
	mov	ecx, dword ptr [ebp - 12]
	mov	edx, esp
	mov	dword ptr [edx], ecx
	mov	dword ptr [ebp - 16], eax ## 4-byte Spill
	call	_ride_ratings_calculate
	mov	eax, dword ptr [ebp - 12]
	mov	ecx, esp
	mov	dword ptr [ecx + 24], eax
	mov	dword ptr [ecx + 28], 0
	mov	dword ptr [ecx + 20], 0
	mov	dword ptr [ecx + 16], 0
	mov	dword ptr [ecx + 12], 0
	mov	dword ptr [ecx + 8], 0
	mov	dword ptr [ecx + 4], 0
	mov	dword ptr [ecx], 6643556
	call	_RCT2_CALLPROC_X
	mov	ecx, dword ptr [ebp - 12]
	mov	edx, esp
	mov	dword ptr [edx], ecx
	mov	dword ptr [ebp - 20], eax ## 4-byte Spill
	call	_ride_ratings_calculate_value
	mov	bl, byte ptr [20493712]
	movzx	eax, bl
	mov	si, ax
	mov	dword ptr [esp], 12
	movzx	eax, si
	mov	dword ptr [esp + 4], eax
	call	_window_invalidate_by_number
	mov	eax, 20493713
	mov	byte ptr [eax], 0
LBB7_4:
	add	esp, 48
	pop	esi
	pop	ebx
	pop	ebp
	ret

	.align	4, 0x90
_ride_ratings_calculate:                ## @ride_ratings_calculate
## BB#0:
	push	ebp
	mov	ebp, esp
	sub	esp, 56
	call	L8$pb
L8$pb:
	pop	eax
	mov	ecx, dword ptr [ebp + 8]
	mov	dword ptr [ebp - 4], ecx
	mov	ecx, dword ptr [ebp - 4]
	movzx	ecx, byte ptr [ecx]
	mov	eax, dword ptr [eax + 4*ecx + _ride_ratings_calculate_func_table-L8$pb]
	mov	dword ptr [ebp - 8], eax
	cmp	dword ptr [ebp - 8], 0
	jne	LBB8_2
## BB#1:
	mov	eax, 0
	mov	ecx, 9953360
	mov	edx, dword ptr [ebp - 4]
	movzx	edx, byte ptr [edx]
	mov	ecx, dword ptr [ecx + 4*edx]
	mov	dword ptr [ebp - 8], ecx
	mov	ecx, dword ptr [ebp - 8]
	mov	edx, dword ptr [ebp - 4]
	mov	dword ptr [esp], ecx
	mov	dword ptr [esp + 4], 0
	mov	dword ptr [esp + 8], 0
	mov	dword ptr [esp + 12], 0
	mov	dword ptr [esp + 16], 0
	mov	dword ptr [esp + 20], 0
	mov	dword ptr [esp + 24], edx
	mov	dword ptr [esp + 28], 0
	mov	dword ptr [ebp - 12], eax ## 4-byte Spill
	call	_RCT2_CALLPROC_X
	mov	dword ptr [ebp - 16], eax ## 4-byte Spill
	jmp	LBB8_3
LBB8_2:
	mov	eax, dword ptr [ebp - 8]
	mov	ecx, dword ptr [ebp - 4]
	mov	dword ptr [esp], ecx
	call	eax
LBB8_3:
	add	esp, 56
	pop	ebp
	ret

	.align	4, 0x90
_ride_ratings_calculate_value:          ## @ride_ratings_calculate_value
## BB#0:
	push	ebp
	mov	ebp, esp
	push	esi
	sub	esp, 60
	call	L9$pb
L9$pb:
	pop	eax
	mov	ecx, dword ptr [ebp + 8]
	mov	dword ptr [ebp - 8], ecx
	mov	ecx, dword ptr [ebp - 8]
	movsx	ecx, word ptr [ecx + 320]
	cmp	ecx, -1
	mov	dword ptr [ebp - 32], eax ## 4-byte Spill
	jne	LBB9_2
## BB#1:
	jmp	LBB9_34
LBB9_2:
	mov	eax, 16147368
	mov	ecx, dword ptr [ebp - 8]
	movsx	ecx, word ptr [ecx + 320]
	mov	edx, dword ptr [ebp - 8]
	movzx	edx, byte ptr [edx]
	imul	edx, edx, 6
	add	edx, 9948446
	movsx	edx, word ptr [edx]
	imul	ecx, edx
	shl	ecx, 5
	sar	ecx, 15
	mov	edx, dword ptr [ebp - 8]
	movsx	edx, word ptr [edx + 322]
	mov	esi, dword ptr [ebp - 8]
	movzx	esi, byte ptr [esi]
	imul	esi, esi, 6
	add	esi, 9948448
	movsx	esi, word ptr [esi]
	imul	edx, esi
	shl	edx, 5
	sar	edx, 15
	add	ecx, edx
	mov	edx, dword ptr [ebp - 8]
	movsx	edx, word ptr [edx + 324]
	mov	esi, dword ptr [ebp - 8]
	movzx	esi, byte ptr [esi]
	imul	esi, esi, 6
	add	esi, 9948450
	movsx	esi, word ptr [esi]
	imul	edx, esi
	shl	edx, 5
	sar	edx, 15
	add	ecx, edx
	mov	dword ptr [ebp - 24], ecx
	movzx	eax, word ptr [eax]
	mov	ecx, dword ptr [ebp - 8]
	movsx	ecx, word ptr [ecx + 384]
	sub	eax, ecx
	mov	dword ptr [ebp - 28], eax
	cmp	dword ptr [ebp - 28], 12
	jg	LBB9_6
## BB#3:
	mov	eax, dword ptr [ebp - 24]
	add	eax, 10
	mov	dword ptr [ebp - 24], eax
	cmp	dword ptr [ebp - 28], 4
	jg	LBB9_5
## BB#4:
	mov	eax, dword ptr [ebp - 24]
	add	eax, 20
	mov	dword ptr [ebp - 24], eax
LBB9_5:
	jmp	LBB9_6
LBB9_6:
	cmp	dword ptr [ebp - 28], 40
	jl	LBB9_8
## BB#7:
	mov	eax, 4
	mov	ecx, dword ptr [ebp - 24]
	mov	dword ptr [ebp - 36], eax ## 4-byte Spill
	mov	eax, ecx
	cdq
	mov	ecx, dword ptr [ebp - 36] ## 4-byte Reload
	idiv	ecx
	mov	edx, dword ptr [ebp - 24]
	sub	edx, eax
	mov	dword ptr [ebp - 24], edx
LBB9_8:
	cmp	dword ptr [ebp - 28], 64
	jl	LBB9_10
## BB#9:
	mov	eax, 4
	mov	ecx, dword ptr [ebp - 24]
	mov	dword ptr [ebp - 40], eax ## 4-byte Spill
	mov	eax, ecx
	cdq
	mov	ecx, dword ptr [ebp - 40] ## 4-byte Reload
	idiv	ecx
	mov	edx, dword ptr [ebp - 24]
	sub	edx, eax
	mov	dword ptr [ebp - 24], edx
LBB9_10:
	cmp	dword ptr [ebp - 28], 200
	jge	LBB9_20
## BB#11:
	cmp	dword ptr [ebp - 28], 88
	jl	LBB9_13
## BB#12:
	mov	eax, 4
	mov	ecx, dword ptr [ebp - 24]
	mov	dword ptr [ebp - 44], eax ## 4-byte Spill
	mov	eax, ecx
	cdq
	mov	ecx, dword ptr [ebp - 44] ## 4-byte Reload
	idiv	ecx
	mov	edx, dword ptr [ebp - 24]
	sub	edx, eax
	mov	dword ptr [ebp - 24], edx
LBB9_13:
	cmp	dword ptr [ebp - 28], 104
	jl	LBB9_15
## BB#14:
	mov	eax, 4
	mov	ecx, dword ptr [ebp - 24]
	mov	dword ptr [ebp - 48], eax ## 4-byte Spill
	mov	eax, ecx
	cdq
	mov	ecx, dword ptr [ebp - 48] ## 4-byte Reload
	idiv	ecx
	mov	edx, dword ptr [ebp - 24]
	sub	edx, eax
	mov	dword ptr [ebp - 24], edx
LBB9_15:
	cmp	dword ptr [ebp - 28], 120
	jl	LBB9_17
## BB#16:
	mov	eax, 2
	mov	ecx, dword ptr [ebp - 24]
	mov	dword ptr [ebp - 52], eax ## 4-byte Spill
	mov	eax, ecx
	cdq
	mov	ecx, dword ptr [ebp - 52] ## 4-byte Reload
	idiv	ecx
	mov	edx, dword ptr [ebp - 24]
	sub	edx, eax
	mov	dword ptr [ebp - 24], edx
LBB9_17:
	cmp	dword ptr [ebp - 28], 128
	jl	LBB9_19
## BB#18:
	mov	eax, 2
	mov	ecx, dword ptr [ebp - 24]
	mov	dword ptr [ebp - 56], eax ## 4-byte Spill
	mov	eax, ecx
	cdq
	mov	ecx, dword ptr [ebp - 56] ## 4-byte Reload
	idiv	ecx
	mov	edx, dword ptr [ebp - 24]
	sub	edx, eax
	mov	dword ptr [ebp - 24], edx
LBB9_19:
	jmp	LBB9_20
LBB9_20:
	mov	dword ptr [ebp - 20], 0
	mov	dword ptr [ebp - 16], 0
LBB9_21:                                ## =>This Inner Loop Header: Depth=1
	cmp	dword ptr [ebp - 16], 255
	jge	LBB9_28
## BB#22:                               ##   in Loop: Header=BB9_21 Depth=1
	mov	eax, dword ptr [ebp - 32] ## 4-byte Reload
	mov	ecx, dword ptr [eax + L_g_ride_list$non_lazy_ptr-L9$pb]
	mov	edx, dword ptr [ebp - 16]
	mov	ecx, dword ptr [ecx]
	imul	edx, edx, 608
	add	ecx, edx
	mov	dword ptr [ebp - 12], ecx
	movzx	ecx, byte ptr [ecx]
	cmp	ecx, 255
	je	LBB9_26
## BB#23:                               ##   in Loop: Header=BB9_21 Depth=1
	mov	eax, dword ptr [ebp - 12]
	movzx	eax, byte ptr [eax]
	mov	ecx, dword ptr [ebp - 8]
	movzx	ecx, byte ptr [ecx]
	cmp	eax, ecx
	jne	LBB9_25
## BB#24:                               ##   in Loop: Header=BB9_21 Depth=1
	mov	eax, dword ptr [ebp - 20]
	add	eax, 1
	mov	dword ptr [ebp - 20], eax
LBB9_25:                                ##   in Loop: Header=BB9_21 Depth=1
	jmp	LBB9_26
LBB9_26:                                ##   in Loop: Header=BB9_21 Depth=1
	jmp	LBB9_27
LBB9_27:                                ##   in Loop: Header=BB9_21 Depth=1
	mov	eax, dword ptr [ebp - 16]
	add	eax, 1
	mov	dword ptr [ebp - 16], eax
	jmp	LBB9_21
LBB9_28:
	cmp	dword ptr [ebp - 20], 1
	jle	LBB9_30
## BB#29:
	mov	eax, 4
	mov	ecx, dword ptr [ebp - 24]
	mov	dword ptr [ebp - 60], eax ## 4-byte Spill
	mov	eax, ecx
	cdq
	mov	ecx, dword ptr [ebp - 60] ## 4-byte Reload
	idiv	ecx
	mov	edx, dword ptr [ebp - 24]
	sub	edx, eax
	mov	dword ptr [ebp - 24], edx
LBB9_30:
	mov	eax, 0
	cmp	eax, dword ptr [ebp - 24]
	jle	LBB9_32
## BB#31:
	mov	eax, 0
	mov	dword ptr [ebp - 64], eax ## 4-byte Spill
	jmp	LBB9_33
LBB9_32:
	mov	eax, dword ptr [ebp - 24]
	mov	dword ptr [ebp - 64], eax ## 4-byte Spill
LBB9_33:
	mov	eax, dword ptr [ebp - 64] ## 4-byte Reload
	mov	cx, ax
	mov	eax, dword ptr [ebp - 8]
	mov	word ptr [eax + 326], cx
LBB9_34:
	add	esp, 60
	pop	esi
	pop	ebp
	ret

	.align	4, 0x90
_ride_ratings_calculate_mine_train_coaster: ## @ride_ratings_calculate_mine_train_coaster
## BB#0:
	push	ebp
	mov	ebp, esp
	push	ebx
	push	edi
	push	esi
	sub	esp, 124
	mov	eax, dword ptr [ebp + 8]
	mov	dword ptr [ebp - 16], eax
	mov	eax, dword ptr [ebp - 16]
	mov	eax, dword ptr [eax + 464]
	and	eax, 2
	cmp	eax, 0
	jne	LBB10_2
## BB#1:
	jmp	LBB10_21
LBB10_2:
	mov	eax, dword ptr [ebp - 16]
	mov	byte ptr [eax + 408], 16
	mov	eax, dword ptr [ebp - 16]
	mov	dword ptr [esp], eax
	call	_set_unreliability_factor
	mov	word ptr [ebp - 24], 290
	mov	word ptr [ebp - 22], 230
	mov	word ptr [ebp - 20], 210
	mov	eax, dword ptr [ebp - 16]
	mov	dword ptr [esp], eax
	call	_ride_get_total_length
	mov	ecx, 6000
	sar	eax, 16
	mov	dword ptr [ebp - 36], eax
	cmp	ecx, dword ptr [ebp - 36]
	jge	LBB10_4
## BB#3:
	mov	eax, 6000
	mov	dword ptr [ebp - 76], eax ## 4-byte Spill
	jmp	LBB10_5
LBB10_4:
	mov	eax, dword ptr [ebp - 36]
	mov	dword ptr [ebp - 76], eax ## 4-byte Spill
LBB10_5:
	mov	eax, dword ptr [ebp - 76] ## 4-byte Reload
	imul	eax, eax, 764
	sar	eax, 16
	movsx	ecx, word ptr [ebp - 24]
	add	ecx, eax
	mov	dx, cx
	mov	word ptr [ebp - 24], dx
	mov	eax, dword ptr [ebp - 16]
	movzx	eax, byte ptr [eax + 198]
	and	eax, 32
	cmp	eax, 0
	je	LBB10_7
## BB#6:
	movsx	eax, word ptr [ebp - 24]
	add	eax, 40
	mov	cx, ax
	mov	word ptr [ebp - 24], cx
	movsx	eax, word ptr [ebp - 22]
	add	eax, 5
	mov	cx, ax
	mov	word ptr [ebp - 22], cx
LBB10_7:
	mov	eax, 150
	mov	ecx, dword ptr [ebp - 16]
	movzx	ecx, byte ptr [ecx + 201]
	sub	ecx, 1
	imul	ecx, ecx, 187245
	sar	ecx, 16
	movsx	edx, word ptr [ebp - 24]
	add	edx, ecx
	mov	si, dx
	mov	word ptr [ebp - 24], si
	mov	ecx, dword ptr [ebp - 16]
	mov	ecx, dword ptr [ecx + 216]
	sar	ecx, 16
	imul	ecx, ecx, 44281
	sar	ecx, 16
	movsx	edx, word ptr [ebp - 24]
	add	edx, ecx
	mov	si, dx
	mov	word ptr [ebp - 24], si
	mov	ecx, dword ptr [ebp - 16]
	mov	ecx, dword ptr [ecx + 216]
	sar	ecx, 16
	imul	ecx, ecx, 88562
	sar	ecx, 16
	movsx	edx, word ptr [ebp - 22]
	add	edx, ecx
	mov	si, dx
	mov	word ptr [ebp - 22], si
	mov	ecx, dword ptr [ebp - 16]
	mov	ecx, dword ptr [ecx + 216]
	sar	ecx, 16
	imul	ecx, ecx, 35424
	sar	ecx, 16
	movsx	edx, word ptr [ebp - 20]
	add	edx, ecx
	mov	si, dx
	mov	word ptr [ebp - 20], si
	mov	ecx, dword ptr [ebp - 16]
	mov	ecx, dword ptr [ecx + 220]
	sar	ecx, 16
	imul	ecx, ecx, 291271
	sar	ecx, 16
	movsx	edx, word ptr [ebp - 24]
	add	edx, ecx
	mov	si, dx
	mov	word ptr [ebp - 24], si
	mov	ecx, dword ptr [ebp - 16]
	mov	ecx, dword ptr [ecx + 220]
	sar	ecx, 16
	imul	ecx, ecx, 436906
	sar	ecx, 16
	movsx	edx, word ptr [ebp - 22]
	add	edx, ecx
	mov	si, dx
	mov	word ptr [ebp - 22], si
	mov	ecx, dword ptr [ebp - 16]
	movzx	ecx, word ptr [ecx + 244]
	mov	edx, dword ptr [ebp - 16]
	movzx	edx, word ptr [edx + 246]
	add	ecx, edx
	mov	edx, dword ptr [ebp - 16]
	movzx	edx, word ptr [edx + 248]
	add	ecx, edx
	mov	edx, dword ptr [ebp - 16]
	movzx	edx, word ptr [edx + 250]
	add	ecx, edx
	mov	dword ptr [ebp - 40], ecx
	cmp	eax, dword ptr [ebp - 40]
	jge	LBB10_9
## BB#8:
	mov	eax, 150
	mov	dword ptr [ebp - 80], eax ## 4-byte Spill
	jmp	LBB10_10
LBB10_9:
	mov	eax, dword ptr [ebp - 40]
	mov	dword ptr [ebp - 80], eax ## 4-byte Spill
LBB10_10:
	mov	eax, dword ptr [ebp - 80] ## 4-byte Reload
	lea	ecx, dword ptr [ebp - 48]
	imul	eax, eax, 26214
	sar	eax, 16
	movsx	edx, word ptr [ebp - 24]
	add	edx, eax
	mov	si, dx
	mov	word ptr [ebp - 24], si
	mov	eax, dword ptr [ebp - 16]
	mov	dword ptr [esp], ecx
	mov	dword ptr [esp + 4], eax
	call	_ride_ratings_get_gforce_ratings
	sub	esp, 4
	lea	eax, dword ptr [ebp - 56]
	mov	ecx, dword ptr [ebp - 48]
	mov	dword ptr [ebp - 32], ecx
	mov	si, word ptr [ebp - 44]
	mov	word ptr [ebp - 28], si
	movsx	ecx, word ptr [ebp - 32]
	imul	ecx, ecx, 40960
	sar	ecx, 16
	movsx	edx, word ptr [ebp - 24]
	add	edx, ecx
	mov	si, dx
	mov	word ptr [ebp - 24], si
	movsx	ecx, word ptr [ebp - 30]
	imul	ecx, ecx, 35746
	sar	ecx, 16
	movsx	edx, word ptr [ebp - 22]
	add	edx, ecx
	mov	si, dx
	mov	word ptr [ebp - 22], si
	movsx	ecx, word ptr [ebp - 28]
	imul	ecx, ecx, 49648
	sar	ecx, 16
	movsx	edx, word ptr [ebp - 20]
	add	edx, ecx
	mov	si, dx
	mov	word ptr [ebp - 20], si
	mov	ecx, dword ptr [ebp - 16]
	mov	dword ptr [esp], eax
	mov	dword ptr [esp + 4], ecx
	call	_sub_65DDD1
	sub	esp, 4
	lea	eax, dword ptr [ebp - 64]
	mov	ecx, dword ptr [ebp - 56]
	mov	dword ptr [ebp - 32], ecx
	mov	si, word ptr [ebp - 52]
	mov	word ptr [ebp - 28], si
	movsx	ecx, word ptr [ebp - 32]
	imul	ecx, ecx, 29721
	sar	ecx, 16
	movsx	edx, word ptr [ebp - 24]
	add	edx, ecx
	mov	si, dx
	mov	word ptr [ebp - 24], si
	movsx	ecx, word ptr [ebp - 30]
	imul	ecx, ecx, 34767
	sar	ecx, 16
	movsx	edx, word ptr [ebp - 22]
	add	edx, ecx
	mov	si, dx
	mov	word ptr [ebp - 22], si
	movsx	ecx, word ptr [ebp - 28]
	imul	ecx, ecx, 45749
	sar	ecx, 16
	movsx	edx, word ptr [ebp - 20]
	add	edx, ecx
	mov	si, dx
	mov	word ptr [ebp - 20], si
	mov	ecx, dword ptr [ebp - 16]
	mov	dword ptr [esp], eax
	mov	dword ptr [esp + 4], ecx
	call	_ride_ratings_get_drop_ratings
	sub	esp, 4
	lea	eax, dword ptr [ebp - 72]
	mov	ecx, dword ptr [ebp - 64]
	mov	dword ptr [ebp - 32], ecx
	mov	si, word ptr [ebp - 60]
	mov	word ptr [ebp - 28], si
	movsx	ecx, word ptr [ebp - 32]
	imul	ecx, ecx, 29127
	sar	ecx, 16
	movsx	edx, word ptr [ebp - 24]
	add	edx, ecx
	mov	si, dx
	mov	word ptr [ebp - 24], si
	movsx	ecx, word ptr [ebp - 30]
	imul	ecx, ecx, 46811
	sar	ecx, 16
	movsx	edx, word ptr [ebp - 22]
	add	edx, ecx
	mov	si, dx
	mov	word ptr [ebp - 22], si
	movsx	ecx, word ptr [ebp - 28]
	imul	ecx, ecx, 49152
	sar	ecx, 16
	movsx	edx, word ptr [ebp - 20]
	add	edx, ecx
	mov	si, dx
	mov	word ptr [ebp - 20], si
	mov	ecx, dword ptr [ebp - 16]
	mov	dword ptr [esp], eax
	mov	dword ptr [esp + 4], ecx
	call	_sub_65E1C2
	sub	esp, 4
	mov	eax, dword ptr [ebp - 72]
	mov	dword ptr [ebp - 32], eax
	mov	si, word ptr [ebp - 68]
	mov	word ptr [ebp - 28], si
	movsx	eax, word ptr [ebp - 32]
	imul	eax, eax, 19275
	sar	eax, 16
	movsx	ecx, word ptr [ebp - 24]
	add	ecx, eax
	mov	si, cx
	mov	word ptr [ebp - 24], si
	movsx	eax, word ptr [ebp - 30]
	shl	eax, 15
	sar	eax, 16
	movsx	ecx, word ptr [ebp - 22]
	add	ecx, eax
	mov	si, cx
	mov	word ptr [ebp - 22], si
	movsx	eax, word ptr [ebp - 28]
	imul	eax, eax, 35108
	sar	eax, 16
	movsx	ecx, word ptr [ebp - 20]
	add	ecx, eax
	mov	si, cx
	mov	word ptr [ebp - 20], si
	call	_sub_65E277
	imul	eax, eax, 21472
	sar	eax, 16
	movsx	ecx, word ptr [ebp - 24]
	add	ecx, eax
	mov	si, cx
	mov	word ptr [ebp - 24], si
	mov	eax, dword ptr [ebp - 16]
	mov	dword ptr [esp], eax
	call	_ride_ratings_get_scenery_score
	imul	eax, eax, 16732
	sar	eax, 16
	movsx	ecx, word ptr [ebp - 24]
	add	ecx, eax
	mov	si, cx
	mov	word ptr [ebp - 24], si
	mov	eax, dword ptr [ebp - 16]
	movzx	eax, byte ptr [eax + 279]
	cmp	eax, 8
	jge	LBB10_12
## BB#11:
	mov	eax, 2
	mov	ecx, dword ptr [ebp - 16]
	movsx	edx, word ptr [ecx + 320]
	mov	dword ptr [ebp - 84], eax ## 4-byte Spill
	mov	eax, edx
	cdq
	mov	esi, dword ptr [ebp - 84] ## 4-byte Reload
	idiv	esi
	mov	di, ax
	mov	word ptr [ecx + 320], di
	mov	eax, dword ptr [ebp - 16]
	movsx	ecx, word ptr [eax + 322]
	mov	dword ptr [ebp - 88], eax ## 4-byte Spill
	mov	eax, ecx
	cdq
	idiv	esi
	mov	di, ax
	mov	eax, dword ptr [ebp - 88] ## 4-byte Reload
	mov	word ptr [eax + 322], di
	mov	ecx, dword ptr [ebp - 16]
	movsx	eax, word ptr [ecx + 324]
	cdq
	idiv	esi
	mov	di, ax
	mov	word ptr [ecx + 324], di
LBB10_12:
	mov	eax, dword ptr [ebp - 16]
	cmp	dword ptr [eax + 216], 655360
	jge	LBB10_14
## BB#13:
	mov	eax, 2
	mov	ecx, dword ptr [ebp - 16]
	movsx	edx, word ptr [ecx + 320]
	mov	dword ptr [ebp - 92], eax ## 4-byte Spill
	mov	eax, edx
	cdq
	mov	esi, dword ptr [ebp - 92] ## 4-byte Reload
	idiv	esi
	mov	di, ax
	mov	word ptr [ecx + 320], di
	mov	eax, dword ptr [ebp - 16]
	movsx	ecx, word ptr [eax + 322]
	mov	dword ptr [ebp - 96], eax ## 4-byte Spill
	mov	eax, ecx
	cdq
	idiv	esi
	mov	di, ax
	mov	eax, dword ptr [ebp - 96] ## 4-byte Reload
	mov	word ptr [eax + 322], di
	mov	ecx, dword ptr [ebp - 16]
	movsx	eax, word ptr [ecx + 324]
	cdq
	idiv	esi
	mov	di, ax
	mov	word ptr [ecx + 324], di
LBB10_14:
	mov	eax, dword ptr [ebp - 16]
	movsx	eax, word ptr [eax + 254]
	cmp	eax, 10
	jle	LBB10_16
## BB#15:
	mov	eax, 2
	mov	ecx, dword ptr [ebp - 16]
	movsx	edx, word ptr [ecx + 320]
	mov	dword ptr [ebp - 100], eax ## 4-byte Spill
	mov	eax, edx
	cdq
	mov	esi, dword ptr [ebp - 100] ## 4-byte Reload
	idiv	esi
	mov	di, ax
	mov	word ptr [ecx + 320], di
	mov	eax, dword ptr [ebp - 16]
	movsx	ecx, word ptr [eax + 322]
	mov	dword ptr [ebp - 104], eax ## 4-byte Spill
	mov	eax, ecx
	cdq
	idiv	esi
	mov	di, ax
	mov	eax, dword ptr [ebp - 104] ## 4-byte Reload
	mov	word ptr [eax + 322], di
	mov	ecx, dword ptr [ebp - 16]
	movsx	eax, word ptr [ecx + 324]
	cdq
	idiv	esi
	mov	di, ax
	mov	word ptr [ecx + 324], di
LBB10_16:
	mov	eax, dword ptr [ebp - 16]
	cmp	dword ptr [eax + 228], 24248320
	jge	LBB10_18
## BB#17:
	mov	eax, 2
	mov	ecx, dword ptr [ebp - 16]
	movsx	edx, word ptr [ecx + 320]
	mov	dword ptr [ebp - 108], eax ## 4-byte Spill
	mov	eax, edx
	cdq
	mov	esi, dword ptr [ebp - 108] ## 4-byte Reload
	idiv	esi
	mov	di, ax
	mov	word ptr [ecx + 320], di
	mov	eax, dword ptr [ebp - 16]
	movsx	ecx, word ptr [eax + 322]
	mov	dword ptr [ebp - 112], eax ## 4-byte Spill
	mov	eax, ecx
	cdq
	idiv	esi
	mov	di, ax
	mov	eax, dword ptr [ebp - 112] ## 4-byte Reload
	mov	word ptr [eax + 322], di
	mov	ecx, dword ptr [ebp - 16]
	movsx	eax, word ptr [ecx + 324]
	cdq
	idiv	esi
	mov	di, ax
	mov	word ptr [ecx + 324], di
LBB10_18:
	mov	eax, dword ptr [ebp - 16]
	movzx	eax, byte ptr [eax + 277]
	and	eax, 63
	cmp	eax, 2
	jge	LBB10_20
## BB#19:
	mov	eax, 2
	mov	ecx, dword ptr [ebp - 16]
	movsx	edx, word ptr [ecx + 320]
	mov	dword ptr [ebp - 116], eax ## 4-byte Spill
	mov	eax, edx
	cdq
	mov	esi, dword ptr [ebp - 116] ## 4-byte Reload
	idiv	esi
	mov	di, ax
	mov	word ptr [ecx + 320], di
	mov	eax, dword ptr [ebp - 16]
	movsx	ecx, word ptr [eax + 322]
	mov	dword ptr [ebp - 120], eax ## 4-byte Spill
	mov	eax, ecx
	cdq
	idiv	esi
	mov	di, ax
	mov	eax, dword ptr [ebp - 120] ## 4-byte Reload
	mov	word ptr [eax + 322], di
	mov	ecx, dword ptr [ebp - 16]
	movsx	eax, word ptr [ecx + 324]
	cdq
	idiv	esi
	mov	di, ax
	mov	word ptr [ecx + 324], di
LBB10_20:
	lea	eax, dword ptr [ebp - 24]
	mov	dword ptr [esp], eax
	call	_ride_ratings_apply_intensity_penalty
	lea	eax, dword ptr [ebp - 24]
	mov	ecx, dword ptr [ebp - 16]
	mov	dword ptr [esp], ecx
	mov	dword ptr [esp + 4], eax
	call	_ride_ratings_apply_adjustments
	mov	eax, dword ptr [ebp - 16]
	mov	ecx, dword ptr [ebp - 24]
	mov	dword ptr [eax + 320], ecx
	mov	dx, word ptr [ebp - 20]
	mov	word ptr [eax + 324], dx
	mov	eax, dword ptr [ebp - 16]
	mov	dword ptr [esp], eax
	call	_ride_compute_upkeep
	mov	ecx, dword ptr [ebp - 16]
	mov	word ptr [ecx + 386], ax
	mov	ecx, dword ptr [ebp - 16]
	movzx	esi, byte ptr [ecx + 333]
	or	esi, 2
	mov	ebx, esi
	mov	al, bl
	mov	byte ptr [ecx + 333], al
	mov	ecx, dword ptr [ebp - 16]
	movzx	esi, byte ptr [ecx + 276]
	and	esi, 31
	mov	ebx, esi
	mov	al, bl
	mov	byte ptr [ecx + 276], al
	mov	ecx, dword ptr [ebp - 16]
	mov	dword ptr [esp], ecx
	call	_sub_65E72D
	shl	eax, 5
	mov	ecx, dword ptr [ebp - 16]
	movzx	esi, byte ptr [ecx + 276]
	or	esi, eax
	mov	eax, esi
                                        ## kill: AL<def> AL<kill> EAX<kill>
	mov	byte ptr [ecx + 276], al
LBB10_21:
	add	esp, 124
	pop	esi
	pop	edi
	pop	ebx
	pop	ebp
	ret

	.align	4, 0x90
_ride_ratings_calculate_maze:           ## @ride_ratings_calculate_maze
## BB#0:
	push	ebp
	mov	ebp, esp
	push	ebx
	push	esi
	sub	esp, 32
	mov	eax, dword ptr [ebp + 8]
	mov	dword ptr [ebp - 12], eax
	mov	eax, dword ptr [ebp - 12]
	mov	ecx, dword ptr [eax + 464]
	or	ecx, 2
	mov	dword ptr [eax + 464], ecx
	mov	eax, dword ptr [ebp - 12]
	mov	ecx, dword ptr [eax + 464]
	or	ecx, 8
	mov	dword ptr [eax + 464], ecx
	mov	eax, dword ptr [ebp - 12]
	mov	byte ptr [eax + 408], 8
	mov	eax, dword ptr [ebp - 12]
	mov	dword ptr [esp], eax
	call	_set_unreliability_factor
	mov	word ptr [ebp - 24], 130
	mov	word ptr [ebp - 22], 50
	mov	word ptr [ebp - 20], 0
	mov	eax, dword ptr [ebp - 12]
	movzx	eax, word ptr [eax + 350]
	cmp	eax, 100
	jge	LBB11_2
## BB#1:
	mov	eax, dword ptr [ebp - 12]
	movzx	eax, word ptr [eax + 350]
	mov	dword ptr [ebp - 32], eax ## 4-byte Spill
	jmp	LBB11_3
LBB11_2:
	mov	eax, 100
	mov	dword ptr [ebp - 32], eax ## 4-byte Spill
	jmp	LBB11_3
LBB11_3:
	mov	eax, dword ptr [ebp - 32] ## 4-byte Reload
	mov	dword ptr [ebp - 28], eax
	mov	eax, dword ptr [ebp - 28]
	movsx	ecx, word ptr [ebp - 24]
	add	ecx, eax
	mov	dx, cx
	mov	word ptr [ebp - 24], dx
	mov	eax, dword ptr [ebp - 28]
	shl	eax, 1
	movsx	ecx, word ptr [ebp - 22]
	add	ecx, eax
	mov	dx, cx
	mov	word ptr [ebp - 22], dx
	mov	eax, dword ptr [ebp - 12]
	mov	dword ptr [esp], eax
	call	_ride_ratings_get_scenery_score
	lea	ecx, dword ptr [ebp - 24]
	imul	eax, eax, 22310
	sar	eax, 16
	movsx	esi, word ptr [ebp - 24]
	add	esi, eax
	mov	dx, si
	mov	word ptr [ebp - 24], dx
	mov	dword ptr [esp], ecx
	call	_ride_ratings_apply_intensity_penalty
	lea	eax, dword ptr [ebp - 24]
	mov	ecx, dword ptr [ebp - 12]
	mov	dword ptr [esp], ecx
	mov	dword ptr [esp + 4], eax
	call	_ride_ratings_apply_adjustments
	mov	eax, dword ptr [ebp - 12]
	mov	ecx, dword ptr [ebp - 24]
	mov	dword ptr [eax + 320], ecx
	mov	dx, word ptr [ebp - 20]
	mov	word ptr [eax + 324], dx
	mov	eax, dword ptr [ebp - 12]
	mov	dword ptr [esp], eax
	call	_ride_compute_upkeep
	mov	ecx, dword ptr [ebp - 12]
	mov	word ptr [ecx + 386], ax
	mov	ecx, dword ptr [ebp - 12]
	movzx	esi, byte ptr [ecx + 333]
	or	esi, 2
	mov	ebx, esi
	mov	al, bl
	mov	byte ptr [ecx + 333], al
	mov	ecx, dword ptr [ebp - 12]
	movzx	esi, byte ptr [ecx + 276]
	and	esi, 31
	mov	ebx, esi
	mov	al, bl
	mov	byte ptr [ecx + 276], al
	mov	ecx, dword ptr [ebp - 12]
	movzx	esi, byte ptr [ecx + 276]
	mov	ebx, esi
	mov	al, bl
	mov	byte ptr [ecx + 276], al
	add	esp, 32
	pop	esi
	pop	ebx
	pop	ebp
	ret

	.align	4, 0x90
_ride_ratings_calculate_spiral_slide:   ## @ride_ratings_calculate_spiral_slide
## BB#0:
	push	ebp
	mov	ebp, esp
	push	ebx
	push	esi
	sub	esp, 32
	mov	eax, dword ptr [ebp + 8]
	mov	dword ptr [ebp - 12], eax
	mov	eax, dword ptr [ebp - 12]
	mov	ecx, dword ptr [eax + 464]
	or	ecx, 2
	mov	dword ptr [eax + 464], ecx
	mov	eax, dword ptr [ebp - 12]
	mov	ecx, dword ptr [eax + 464]
	or	ecx, 8
	mov	dword ptr [eax + 464], ecx
	mov	eax, dword ptr [ebp - 12]
	mov	byte ptr [eax + 408], 8
	mov	eax, dword ptr [ebp - 12]
	mov	dword ptr [esp], eax
	call	_set_unreliability_factor
	mov	word ptr [ebp - 24], 150
	mov	word ptr [ebp - 22], 140
	mov	word ptr [ebp - 20], 90
	mov	eax, dword ptr [ebp - 12]
	movzx	eax, byte ptr [eax + 4]
	cmp	eax, 10
	jne	LBB12_2
## BB#1:
	movsx	eax, word ptr [ebp - 24]
	add	eax, 40
	mov	cx, ax
	mov	word ptr [ebp - 24], cx
	movsx	eax, word ptr [ebp - 22]
	add	eax, 20
	mov	cx, ax
	mov	word ptr [ebp - 22], cx
	movsx	eax, word ptr [ebp - 20]
	add	eax, 25
	mov	cx, ax
	mov	word ptr [ebp - 20], cx
LBB12_2:
	mov	eax, dword ptr [ebp - 12]
	mov	dword ptr [esp], eax
	call	_ride_ratings_get_scenery_score
	lea	ecx, dword ptr [ebp - 24]
	imul	eax, eax, 25098
	sar	eax, 16
	movsx	edx, word ptr [ebp - 24]
	add	edx, eax
	mov	si, dx
	mov	word ptr [ebp - 24], si
	mov	dword ptr [esp], ecx
	call	_ride_ratings_apply_intensity_penalty
	lea	eax, dword ptr [ebp - 24]
	mov	ecx, dword ptr [ebp - 12]
	mov	dword ptr [esp], ecx
	mov	dword ptr [esp + 4], eax
	call	_ride_ratings_apply_adjustments
	mov	eax, dword ptr [ebp - 12]
	mov	ecx, dword ptr [ebp - 24]
	mov	dword ptr [eax + 320], ecx
	mov	si, word ptr [ebp - 20]
	mov	word ptr [eax + 324], si
	mov	eax, dword ptr [ebp - 12]
	mov	dword ptr [esp], eax
	call	_ride_compute_upkeep
	mov	ecx, dword ptr [ebp - 12]
	mov	word ptr [ecx + 386], ax
	mov	ecx, dword ptr [ebp - 12]
	movzx	edx, byte ptr [ecx + 333]
	or	edx, 2
	mov	bl, dl
	mov	byte ptr [ecx + 333], bl
	mov	ecx, dword ptr [ebp - 12]
	movzx	edx, byte ptr [ecx + 276]
	and	edx, 31
	mov	bl, dl
	mov	byte ptr [ecx + 276], bl
	mov	ecx, dword ptr [ebp - 12]
	movzx	edx, byte ptr [ecx + 276]
	or	edx, 64
	mov	bl, dl
	mov	byte ptr [ecx + 276], bl
	add	esp, 32
	pop	esi
	pop	ebx
	pop	ebp
	ret

	.align	4, 0x90
_ride_ratings_calculate_pirate_ship:    ## @ride_ratings_calculate_pirate_ship
## BB#0:
	push	ebp
	mov	ebp, esp
	push	ebx
	push	esi
	sub	esp, 32
	mov	eax, dword ptr [ebp + 8]
	mov	dword ptr [ebp - 12], eax
	mov	eax, dword ptr [ebp - 12]
	mov	ecx, dword ptr [eax + 464]
	or	ecx, 2
	mov	dword ptr [eax + 464], ecx
	mov	eax, dword ptr [ebp - 12]
	mov	ecx, dword ptr [eax + 464]
	or	ecx, 8
	mov	dword ptr [eax + 464], ecx
	mov	eax, dword ptr [ebp - 12]
	mov	byte ptr [eax + 408], 10
	mov	eax, dword ptr [ebp - 12]
	mov	dword ptr [esp], eax
	call	_set_unreliability_factor
	mov	word ptr [ebp - 24], 150
	mov	word ptr [ebp - 22], 190
	mov	word ptr [ebp - 20], 141
	mov	eax, dword ptr [ebp - 12]
	movzx	eax, byte ptr [eax + 208]
	imul	eax, eax, 5
	movsx	ecx, word ptr [ebp - 24]
	add	ecx, eax
	mov	dx, cx
	mov	word ptr [ebp - 24], dx
	mov	eax, dword ptr [ebp - 12]
	movzx	eax, byte ptr [eax + 208]
	imul	eax, eax, 5
	movsx	ecx, word ptr [ebp - 22]
	add	ecx, eax
	mov	dx, cx
	mov	word ptr [ebp - 22], dx
	mov	eax, dword ptr [ebp - 12]
	movzx	eax, byte ptr [eax + 208]
	imul	eax, eax, 10
	movsx	ecx, word ptr [ebp - 20]
	add	ecx, eax
	mov	dx, cx
	mov	word ptr [ebp - 20], dx
	mov	eax, dword ptr [ebp - 12]
	mov	dword ptr [esp], eax
	call	_ride_ratings_get_scenery_score
	lea	ecx, dword ptr [ebp - 24]
	imul	eax, eax, 16732
	sar	eax, 16
	movsx	esi, word ptr [ebp - 24]
	add	esi, eax
	mov	dx, si
	mov	word ptr [ebp - 24], dx
	mov	dword ptr [esp], ecx
	call	_ride_ratings_apply_intensity_penalty
	lea	eax, dword ptr [ebp - 24]
	mov	ecx, dword ptr [ebp - 12]
	mov	dword ptr [esp], ecx
	mov	dword ptr [esp + 4], eax
	call	_ride_ratings_apply_adjustments
	mov	eax, dword ptr [ebp - 12]
	mov	ecx, dword ptr [ebp - 24]
	mov	dword ptr [eax + 320], ecx
	mov	dx, word ptr [ebp - 20]
	mov	word ptr [eax + 324], dx
	mov	eax, dword ptr [ebp - 12]
	mov	dword ptr [esp], eax
	call	_ride_compute_upkeep
	mov	ecx, dword ptr [ebp - 12]
	mov	word ptr [ecx + 386], ax
	mov	ecx, dword ptr [ebp - 12]
	movzx	esi, byte ptr [ecx + 333]
	or	esi, 2
	mov	ebx, esi
	mov	al, bl
	mov	byte ptr [ecx + 333], al
	mov	ecx, dword ptr [ebp - 12]
	movzx	esi, byte ptr [ecx + 276]
	and	esi, 31
	mov	ebx, esi
	mov	al, bl
	mov	byte ptr [ecx + 276], al
	mov	ecx, dword ptr [ebp - 12]
	movzx	esi, byte ptr [ecx + 276]
	mov	ebx, esi
	mov	al, bl
	mov	byte ptr [ecx + 276], al
	add	esp, 32
	pop	esi
	pop	ebx
	pop	ebp
	ret

	.align	4, 0x90
_ride_ratings_calculate_inverter_ship:  ## @ride_ratings_calculate_inverter_ship
## BB#0:
	push	ebp
	mov	ebp, esp
	push	ebx
	push	esi
	sub	esp, 32
	mov	eax, dword ptr [ebp + 8]
	mov	dword ptr [ebp - 12], eax
	mov	eax, dword ptr [ebp - 12]
	mov	ecx, dword ptr [eax + 464]
	or	ecx, 2
	mov	dword ptr [eax + 464], ecx
	mov	eax, dword ptr [ebp - 12]
	mov	ecx, dword ptr [eax + 464]
	or	ecx, 8
	mov	dword ptr [eax + 464], ecx
	mov	eax, dword ptr [ebp - 12]
	mov	byte ptr [eax + 408], 16
	mov	eax, dword ptr [ebp - 12]
	mov	dword ptr [esp], eax
	call	_set_unreliability_factor
	mov	word ptr [ebp - 24], 250
	mov	word ptr [ebp - 22], 270
	mov	word ptr [ebp - 20], 274
	mov	eax, dword ptr [ebp - 12]
	movzx	eax, byte ptr [eax + 208]
	imul	eax, eax, 11
	movsx	ecx, word ptr [ebp - 24]
	add	ecx, eax
	mov	dx, cx
	mov	word ptr [ebp - 24], dx
	mov	eax, dword ptr [ebp - 12]
	movzx	eax, byte ptr [eax + 208]
	imul	eax, eax, 22
	movsx	ecx, word ptr [ebp - 22]
	add	ecx, eax
	mov	dx, cx
	mov	word ptr [ebp - 22], dx
	mov	eax, dword ptr [ebp - 12]
	movzx	eax, byte ptr [eax + 208]
	imul	eax, eax, 22
	movsx	ecx, word ptr [ebp - 20]
	add	ecx, eax
	mov	dx, cx
	mov	word ptr [ebp - 20], dx
	mov	eax, dword ptr [ebp - 12]
	mov	dword ptr [esp], eax
	call	_ride_ratings_get_scenery_score
	lea	ecx, dword ptr [ebp - 24]
	imul	eax, eax, 11155
	sar	eax, 16
	movsx	esi, word ptr [ebp - 24]
	add	esi, eax
	mov	dx, si
	mov	word ptr [ebp - 24], dx
	mov	dword ptr [esp], ecx
	call	_ride_ratings_apply_intensity_penalty
	lea	eax, dword ptr [ebp - 24]
	mov	ecx, dword ptr [ebp - 12]
	mov	dword ptr [esp], ecx
	mov	dword ptr [esp + 4], eax
	call	_ride_ratings_apply_adjustments
	mov	eax, dword ptr [ebp - 12]
	mov	ecx, dword ptr [ebp - 24]
	mov	dword ptr [eax + 320], ecx
	mov	dx, word ptr [ebp - 20]
	mov	word ptr [eax + 324], dx
	mov	eax, dword ptr [ebp - 12]
	mov	dword ptr [esp], eax
	call	_ride_compute_upkeep
	mov	ecx, dword ptr [ebp - 12]
	mov	word ptr [ecx + 386], ax
	mov	ecx, dword ptr [ebp - 12]
	movzx	esi, byte ptr [ecx + 333]
	or	esi, 2
	mov	ebx, esi
	mov	al, bl
	mov	byte ptr [ecx + 333], al
	mov	ecx, dword ptr [ebp - 12]
	movzx	esi, byte ptr [ecx + 276]
	and	esi, 31
	mov	ebx, esi
	mov	al, bl
	mov	byte ptr [ecx + 276], al
	mov	ecx, dword ptr [ebp - 12]
	movzx	esi, byte ptr [ecx + 276]
	mov	ebx, esi
	mov	al, bl
	mov	byte ptr [ecx + 276], al
	add	esp, 32
	pop	esi
	pop	ebx
	pop	ebp
	ret

	.align	4, 0x90
_ride_ratings_calculate_food_stall:     ## @ride_ratings_calculate_food_stall
## BB#0:
	push	ebp
	mov	ebp, esp
	push	ebx
	sub	esp, 20
	mov	eax, dword ptr [ebp + 8]
	mov	dword ptr [ebp - 8], eax
	mov	eax, dword ptr [ebp - 8]
	mov	dword ptr [esp], eax
	call	_ride_compute_upkeep
	mov	ecx, dword ptr [ebp - 8]
	mov	word ptr [ecx + 386], ax
	mov	ecx, dword ptr [ebp - 8]
	movzx	edx, byte ptr [ecx + 333]
	or	edx, 2
	mov	bl, dl
	mov	byte ptr [ecx + 333], bl
	add	esp, 20
	pop	ebx
	pop	ebp
	ret

	.align	4, 0x90
_ride_ratings_calculate_drink_stall:    ## @ride_ratings_calculate_drink_stall
## BB#0:
	push	ebp
	mov	ebp, esp
	push	ebx
	sub	esp, 20
	mov	eax, dword ptr [ebp + 8]
	mov	dword ptr [ebp - 8], eax
	mov	eax, dword ptr [ebp - 8]
	mov	dword ptr [esp], eax
	call	_ride_compute_upkeep
	mov	ecx, dword ptr [ebp - 8]
	mov	word ptr [ecx + 386], ax
	mov	ecx, dword ptr [ebp - 8]
	movzx	edx, byte ptr [ecx + 333]
	or	edx, 2
	mov	bl, dl
	mov	byte ptr [ecx + 333], bl
	add	esp, 20
	pop	ebx
	pop	ebp
	ret

	.align	4, 0x90
_ride_ratings_calculate_shop:           ## @ride_ratings_calculate_shop
## BB#0:
	push	ebp
	mov	ebp, esp
	push	ebx
	sub	esp, 20
	mov	eax, dword ptr [ebp + 8]
	mov	dword ptr [ebp - 8], eax
	mov	eax, dword ptr [ebp - 8]
	mov	dword ptr [esp], eax
	call	_ride_compute_upkeep
	mov	ecx, dword ptr [ebp - 8]
	mov	word ptr [ecx + 386], ax
	mov	ecx, dword ptr [ebp - 8]
	movzx	edx, byte ptr [ecx + 333]
	or	edx, 2
	mov	bl, dl
	mov	byte ptr [ecx + 333], bl
	add	esp, 20
	pop	ebx
	pop	ebp
	ret

	.align	4, 0x90
_ride_ratings_calculate_merry_go_round: ## @ride_ratings_calculate_merry_go_round
## BB#0:
	push	ebp
	mov	ebp, esp
	push	ebx
	push	esi
	sub	esp, 32
	mov	eax, dword ptr [ebp + 8]
	mov	dword ptr [ebp - 12], eax
	mov	eax, dword ptr [ebp - 12]
	mov	ecx, dword ptr [eax + 464]
	or	ecx, 2
	mov	dword ptr [eax + 464], ecx
	mov	eax, dword ptr [ebp - 12]
	mov	ecx, dword ptr [eax + 464]
	or	ecx, 8
	mov	dword ptr [eax + 464], ecx
	mov	eax, dword ptr [ebp - 12]
	mov	byte ptr [eax + 408], 16
	mov	eax, dword ptr [ebp - 12]
	mov	dword ptr [esp], eax
	call	_set_unreliability_factor
	mov	eax, dword ptr [ebp - 12]
	movzx	eax, byte ptr [eax + 208]
	imul	eax, eax, 5
	mov	dword ptr [ebp - 28], eax
	mov	eax, dword ptr [ebp - 28]
	add	eax, 60
	mov	ecx, dword ptr [ebp - 12]
	mov	dword ptr [esp], ecx
	mov	dword ptr [ebp - 32], eax ## 4-byte Spill
	call	_ride_ratings_get_scenery_score
	lea	ecx, dword ptr [ebp - 24]
	imul	eax, eax, 19521
	sar	eax, 16
	mov	edx, dword ptr [ebp - 32] ## 4-byte Reload
	add	edx, eax
	mov	si, dx
	mov	word ptr [ebp - 24], si
	mov	eax, dword ptr [ebp - 28]
	add	eax, 15
	mov	si, ax
	mov	word ptr [ebp - 22], si
	mov	eax, dword ptr [ebp - 28]
	add	eax, 30
	mov	si, ax
	mov	word ptr [ebp - 20], si
	mov	dword ptr [esp], ecx
	call	_ride_ratings_apply_intensity_penalty
	lea	eax, dword ptr [ebp - 24]
	mov	ecx, dword ptr [ebp - 12]
	mov	dword ptr [esp], ecx
	mov	dword ptr [esp + 4], eax
	call	_ride_ratings_apply_adjustments
	mov	eax, dword ptr [ebp - 12]
	mov	ecx, dword ptr [ebp - 24]
	mov	dword ptr [eax + 320], ecx
	mov	si, word ptr [ebp - 20]
	mov	word ptr [eax + 324], si
	mov	eax, dword ptr [ebp - 12]
	mov	dword ptr [esp], eax
	call	_ride_compute_upkeep
	mov	ecx, dword ptr [ebp - 12]
	mov	word ptr [ecx + 386], ax
	mov	ecx, dword ptr [ebp - 12]
	movzx	edx, byte ptr [ecx + 333]
	or	edx, 2
	mov	bl, dl
	mov	byte ptr [ecx + 333], bl
	mov	ecx, dword ptr [ebp - 12]
	movzx	edx, byte ptr [ecx + 276]
	and	edx, 31
	mov	bl, dl
	mov	byte ptr [ecx + 276], bl
	mov	ecx, dword ptr [ebp - 12]
	movzx	edx, byte ptr [ecx + 276]
	or	edx, 224
	mov	bl, dl
	mov	byte ptr [ecx + 276], bl
	add	esp, 32
	pop	esi
	pop	ebx
	pop	ebp
	ret

	.align	4, 0x90
_ride_ratings_calculate_information_kiosk: ## @ride_ratings_calculate_information_kiosk
## BB#0:
	push	ebp
	mov	ebp, esp
	push	ebx
	sub	esp, 20
	mov	eax, dword ptr [ebp + 8]
	mov	dword ptr [ebp - 8], eax
	mov	eax, dword ptr [ebp - 8]
	mov	dword ptr [esp], eax
	call	_ride_compute_upkeep
	mov	ecx, dword ptr [ebp - 8]
	mov	word ptr [ecx + 386], ax
	mov	ecx, dword ptr [ebp - 8]
	movzx	edx, byte ptr [ecx + 333]
	or	edx, 2
	mov	bl, dl
	mov	byte ptr [ecx + 333], bl
	add	esp, 20
	pop	ebx
	pop	ebp
	ret

	.align	4, 0x90
_ride_ratings_calculate_bathroom:       ## @ride_ratings_calculate_bathroom
## BB#0:
	push	ebp
	mov	ebp, esp
	push	ebx
	sub	esp, 20
	mov	eax, dword ptr [ebp + 8]
	mov	dword ptr [ebp - 8], eax
	mov	eax, dword ptr [ebp - 8]
	mov	dword ptr [esp], eax
	call	_ride_compute_upkeep
	mov	ecx, dword ptr [ebp - 8]
	mov	word ptr [ecx + 386], ax
	mov	ecx, dword ptr [ebp - 8]
	movzx	edx, byte ptr [ecx + 333]
	or	edx, 2
	mov	bl, dl
	mov	byte ptr [ecx + 333], bl
	add	esp, 20
	pop	ebx
	pop	ebp
	ret

	.align	4, 0x90
_ride_ratings_calculate_ferris_wheel:   ## @ride_ratings_calculate_ferris_wheel
## BB#0:
	push	ebp
	mov	ebp, esp
	push	ebx
	push	esi
	sub	esp, 32
	mov	eax, dword ptr [ebp + 8]
	mov	dword ptr [ebp - 12], eax
	mov	eax, dword ptr [ebp - 12]
	mov	ecx, dword ptr [eax + 464]
	or	ecx, 2
	mov	dword ptr [eax + 464], ecx
	mov	eax, dword ptr [ebp - 12]
	mov	ecx, dword ptr [eax + 464]
	or	ecx, 8
	mov	dword ptr [eax + 464], ecx
	mov	eax, dword ptr [ebp - 12]
	mov	byte ptr [eax + 408], 16
	mov	eax, dword ptr [ebp - 12]
	mov	dword ptr [esp], eax
	call	_set_unreliability_factor
	mov	eax, dword ptr [ebp - 12]
	movzx	eax, byte ptr [eax + 208]
	imul	eax, eax, 25
	mov	dword ptr [ebp - 28], eax
	mov	eax, dword ptr [ebp - 28]
	add	eax, 60
	mov	ecx, dword ptr [ebp - 12]
	mov	dword ptr [esp], ecx
	mov	dword ptr [ebp - 32], eax ## 4-byte Spill
	call	_ride_ratings_get_scenery_score
	lea	ecx, dword ptr [ebp - 24]
	imul	eax, eax, 41831
	sar	eax, 16
	mov	edx, dword ptr [ebp - 32] ## 4-byte Reload
	add	edx, eax
	mov	si, dx
	mov	word ptr [ebp - 24], si
	mov	eax, dword ptr [ebp - 28]
	add	eax, 25
	mov	si, ax
	mov	word ptr [ebp - 22], si
	mov	eax, dword ptr [ebp - 28]
	add	eax, 30
	mov	si, ax
	mov	word ptr [ebp - 20], si
	mov	dword ptr [esp], ecx
	call	_ride_ratings_apply_intensity_penalty
	lea	eax, dword ptr [ebp - 24]
	mov	ecx, dword ptr [ebp - 12]
	mov	dword ptr [esp], ecx
	mov	dword ptr [esp + 4], eax
	call	_ride_ratings_apply_adjustments
	mov	eax, dword ptr [ebp - 12]
	mov	ecx, dword ptr [ebp - 24]
	mov	dword ptr [eax + 320], ecx
	mov	si, word ptr [ebp - 20]
	mov	word ptr [eax + 324], si
	mov	eax, dword ptr [ebp - 12]
	mov	dword ptr [esp], eax
	call	_ride_compute_upkeep
	mov	ecx, dword ptr [ebp - 12]
	mov	word ptr [ecx + 386], ax
	mov	ecx, dword ptr [ebp - 12]
	movzx	edx, byte ptr [ecx + 333]
	or	edx, 2
	mov	bl, dl
	mov	byte ptr [ecx + 333], bl
	mov	ecx, dword ptr [ebp - 12]
	movzx	edx, byte ptr [ecx + 276]
	and	edx, 31
	mov	bl, dl
	mov	byte ptr [ecx + 276], bl
	mov	ecx, dword ptr [ebp - 12]
	movzx	edx, byte ptr [ecx + 276]
	mov	bl, dl
	mov	byte ptr [ecx + 276], bl
	add	esp, 32
	pop	esi
	pop	ebx
	pop	ebp
	ret

	.align	4, 0x90
_ride_ratings_calculate_motion_simulator: ## @ride_ratings_calculate_motion_simulator
## BB#0:
	push	ebp
	mov	ebp, esp
	push	ebx
	push	esi
	sub	esp, 32
	mov	eax, dword ptr [ebp + 8]
	mov	dword ptr [ebp - 12], eax
	mov	eax, dword ptr [ebp - 12]
	mov	ecx, dword ptr [eax + 464]
	or	ecx, 2
	mov	dword ptr [eax + 464], ecx
	mov	eax, dword ptr [ebp - 12]
	mov	ecx, dword ptr [eax + 464]
	or	ecx, 8
	mov	dword ptr [eax + 464], ecx
	mov	eax, dword ptr [ebp - 12]
	mov	byte ptr [eax + 408], 21
	mov	eax, dword ptr [ebp - 12]
	mov	dword ptr [esp], eax
	call	_set_unreliability_factor
	mov	eax, dword ptr [ebp - 12]
	movzx	eax, byte ptr [eax + 4]
	cmp	eax, 24
	jne	LBB22_2
## BB#1:
	mov	word ptr [ebp - 24], 290
	mov	word ptr [ebp - 22], 350
	mov	word ptr [ebp - 20], 300
	jmp	LBB22_3
LBB22_2:
	mov	word ptr [ebp - 24], 325
	mov	word ptr [ebp - 22], 410
	mov	word ptr [ebp - 20], 330
LBB22_3:
	lea	eax, dword ptr [ebp - 24]
	mov	dword ptr [esp], eax
	call	_ride_ratings_apply_intensity_penalty
	lea	eax, dword ptr [ebp - 24]
	mov	ecx, dword ptr [ebp - 12]
	mov	dword ptr [esp], ecx
	mov	dword ptr [esp + 4], eax
	call	_ride_ratings_apply_adjustments
	mov	eax, dword ptr [ebp - 12]
	mov	ecx, dword ptr [ebp - 24]
	mov	dword ptr [eax + 320], ecx
	mov	dx, word ptr [ebp - 20]
	mov	word ptr [eax + 324], dx
	mov	eax, dword ptr [ebp - 12]
	mov	dword ptr [esp], eax
	call	_ride_compute_upkeep
	mov	ecx, dword ptr [ebp - 12]
	mov	word ptr [ecx + 386], ax
	mov	ecx, dword ptr [ebp - 12]
	movzx	esi, byte ptr [ecx + 333]
	or	esi, 2
	mov	ebx, esi
	mov	al, bl
	mov	byte ptr [ecx + 333], al
	mov	ecx, dword ptr [ebp - 12]
	movzx	esi, byte ptr [ecx + 276]
	and	esi, 31
	mov	ebx, esi
	mov	al, bl
	mov	byte ptr [ecx + 276], al
	mov	ecx, dword ptr [ebp - 12]
	movzx	esi, byte ptr [ecx + 276]
	or	esi, 224
	mov	ebx, esi
	mov	al, bl
	mov	byte ptr [ecx + 276], al
	add	esp, 32
	pop	esi
	pop	ebx
	pop	ebp
	ret

	.align	4, 0x90
_ride_ratings_calculate_3d_cinema:      ## @ride_ratings_calculate_3d_cinema
## BB#0:
	push	ebp
	mov	ebp, esp
	push	ebx
	push	esi
	sub	esp, 48
	mov	eax, dword ptr [ebp + 8]
	mov	dword ptr [ebp - 12], eax
	mov	cl, byte ptr [eax + 464]
	or	cl, 2
	mov	byte ptr [eax + 464], cl
	mov	eax, dword ptr [ebp - 12]
	mov	edx, dword ptr [eax + 464]
	or	edx, 8
	mov	dword ptr [eax + 464], edx
	mov	eax, dword ptr [ebp - 12]
	mov	byte ptr [eax + 408], 21
	mov	eax, dword ptr [ebp - 12]
	mov	edx, esp
	mov	dword ptr [edx], eax
	call	_set_unreliability_factor
	mov	eax, dword ptr [ebp - 12]
	movzx	eax, byte ptr [eax + 4]
	mov	edx, eax
	sub	edx, 20
	mov	dword ptr [ebp - 28], eax ## 4-byte Spill
	mov	dword ptr [ebp - 32], edx ## 4-byte Spill
	je	LBB23_2
	jmp	LBB23_6
LBB23_6:
	mov	eax, dword ptr [ebp - 28] ## 4-byte Reload
	sub	eax, 25
	mov	dword ptr [ebp - 36], eax ## 4-byte Spill
	je	LBB23_3
	jmp	LBB23_7
LBB23_7:
	mov	eax, dword ptr [ebp - 28] ## 4-byte Reload
	sub	eax, 26
	mov	dword ptr [ebp - 40], eax ## 4-byte Spill
	je	LBB23_4
	jmp	LBB23_1
LBB23_1:
	jmp	LBB23_2
LBB23_2:
	mov	word ptr [ebp - 24], 350
	mov	word ptr [ebp - 22], 240
	mov	word ptr [ebp - 20], 140
	jmp	LBB23_5
LBB23_3:
	mov	word ptr [ebp - 24], 400
	mov	word ptr [ebp - 22], 265
	mov	word ptr [ebp - 20], 155
	jmp	LBB23_5
LBB23_4:
	mov	word ptr [ebp - 24], 420
	mov	word ptr [ebp - 22], 260
	mov	word ptr [ebp - 20], 148
LBB23_5:
	lea	eax, dword ptr [ebp - 24]
	mov	dword ptr [esp], eax
	call	_ride_ratings_apply_intensity_penalty
	lea	eax, dword ptr [ebp - 24]
	mov	ecx, dword ptr [ebp - 12]
	mov	dword ptr [esp], ecx
	mov	dword ptr [esp + 4], eax
	call	_ride_ratings_apply_adjustments
	mov	eax, dword ptr [ebp - 12]
	mov	ecx, dword ptr [ebp - 24]
	mov	dword ptr [eax + 320], ecx
	mov	dx, word ptr [ebp - 20]
	mov	word ptr [eax + 324], dx
	mov	eax, dword ptr [ebp - 12]
	mov	dword ptr [esp], eax
	call	_ride_compute_upkeep
	mov	ecx, dword ptr [ebp - 12]
	mov	word ptr [ecx + 386], ax
	mov	ecx, dword ptr [ebp - 12]
	movzx	esi, byte ptr [ecx + 333]
	or	esi, 2
	mov	ebx, esi
	mov	al, bl
	mov	byte ptr [ecx + 333], al
	mov	ecx, dword ptr [ebp - 12]
	movzx	esi, byte ptr [ecx + 276]
	and	esi, 31
	mov	ebx, esi
	mov	al, bl
	mov	byte ptr [ecx + 276], al
	mov	ecx, dword ptr [ebp - 12]
	movzx	esi, byte ptr [ecx + 276]
	or	esi, 224
	mov	ebx, esi
	mov	al, bl
	mov	byte ptr [ecx + 276], al
	add	esp, 48
	pop	esi
	pop	ebx
	pop	ebp
	ret

	.align	4, 0x90
_ride_ratings_calculate_top_spin:       ## @ride_ratings_calculate_top_spin
## BB#0:
	push	ebp
	mov	ebp, esp
	push	ebx
	push	esi
	sub	esp, 48
	mov	eax, dword ptr [ebp + 8]
	mov	dword ptr [ebp - 12], eax
	mov	cl, byte ptr [eax + 464]
	or	cl, 2
	mov	byte ptr [eax + 464], cl
	mov	eax, dword ptr [ebp - 12]
	mov	edx, dword ptr [eax + 464]
	or	edx, 8
	mov	dword ptr [eax + 464], edx
	mov	eax, dword ptr [ebp - 12]
	mov	byte ptr [eax + 408], 19
	mov	eax, dword ptr [ebp - 12]
	mov	edx, esp
	mov	dword ptr [edx], eax
	call	_set_unreliability_factor
	mov	eax, dword ptr [ebp - 12]
	movzx	eax, byte ptr [eax + 4]
	mov	edx, eax
	sub	edx, 22
	mov	dword ptr [ebp - 28], eax ## 4-byte Spill
	mov	dword ptr [ebp - 32], edx ## 4-byte Spill
	je	LBB24_2
	jmp	LBB24_6
LBB24_6:
	mov	eax, dword ptr [ebp - 28] ## 4-byte Reload
	sub	eax, 27
	mov	dword ptr [ebp - 36], eax ## 4-byte Spill
	je	LBB24_3
	jmp	LBB24_7
LBB24_7:
	mov	eax, dword ptr [ebp - 28] ## 4-byte Reload
	sub	eax, 28
	mov	dword ptr [ebp - 40], eax ## 4-byte Spill
	je	LBB24_4
	jmp	LBB24_1
LBB24_1:
	jmp	LBB24_2
LBB24_2:
	mov	word ptr [ebp - 24], 200
	mov	word ptr [ebp - 22], 480
	mov	word ptr [ebp - 20], 574
	jmp	LBB24_5
LBB24_3:
	mov	word ptr [ebp - 24], 300
	mov	word ptr [ebp - 22], 575
	mov	word ptr [ebp - 20], 664
	jmp	LBB24_5
LBB24_4:
	mov	word ptr [ebp - 24], 320
	mov	word ptr [ebp - 22], 680
	mov	word ptr [ebp - 20], 794
LBB24_5:
	mov	eax, dword ptr [ebp - 12]
	mov	dword ptr [esp], eax
	call	_ride_ratings_get_scenery_score
	lea	ecx, dword ptr [ebp - 24]
	imul	eax, eax, 11155
	sar	eax, 16
	movsx	edx, word ptr [ebp - 24]
	add	edx, eax
	mov	si, dx
	mov	word ptr [ebp - 24], si
	mov	dword ptr [esp], ecx
	call	_ride_ratings_apply_intensity_penalty
	lea	eax, dword ptr [ebp - 24]
	mov	ecx, dword ptr [ebp - 12]
	mov	dword ptr [esp], ecx
	mov	dword ptr [esp + 4], eax
	call	_ride_ratings_apply_adjustments
	mov	eax, dword ptr [ebp - 12]
	mov	ecx, dword ptr [ebp - 24]
	mov	dword ptr [eax + 320], ecx
	mov	si, word ptr [ebp - 20]
	mov	word ptr [eax + 324], si
	mov	eax, dword ptr [ebp - 12]
	mov	dword ptr [esp], eax
	call	_ride_compute_upkeep
	mov	ecx, dword ptr [ebp - 12]
	mov	word ptr [ecx + 386], ax
	mov	ecx, dword ptr [ebp - 12]
	movzx	edx, byte ptr [ecx + 333]
	or	edx, 2
	mov	bl, dl
	mov	byte ptr [ecx + 333], bl
	mov	ecx, dword ptr [ebp - 12]
	movzx	edx, byte ptr [ecx + 276]
	and	edx, 31
	mov	bl, dl
	mov	byte ptr [ecx + 276], bl
	mov	ecx, dword ptr [ebp - 12]
	movzx	edx, byte ptr [ecx + 276]
	mov	bl, dl
	mov	byte ptr [ecx + 276], bl
	add	esp, 48
	pop	esi
	pop	ebx
	pop	ebp
	ret

	.align	4, 0x90
_ride_ratings_calculate_space_rings:    ## @ride_ratings_calculate_space_rings
## BB#0:
	push	ebp
	mov	ebp, esp
	push	ebx
	push	esi
	sub	esp, 32
	mov	eax, dword ptr [ebp + 8]
	mov	dword ptr [ebp - 12], eax
	mov	eax, dword ptr [ebp - 12]
	mov	ecx, dword ptr [eax + 464]
	or	ecx, 2
	mov	dword ptr [eax + 464], ecx
	mov	eax, dword ptr [ebp - 12]
	mov	ecx, dword ptr [eax + 464]
	or	ecx, 8
	mov	dword ptr [eax + 464], ecx
	mov	eax, dword ptr [ebp - 12]
	mov	byte ptr [eax + 408], 7
	mov	eax, dword ptr [ebp - 12]
	mov	dword ptr [esp], eax
	call	_set_unreliability_factor
	mov	word ptr [ebp - 24], 150
	mov	word ptr [ebp - 22], 210
	mov	word ptr [ebp - 20], 650
	mov	eax, dword ptr [ebp - 12]
	mov	dword ptr [esp], eax
	call	_ride_ratings_get_scenery_score
	lea	ecx, dword ptr [ebp - 24]
	imul	eax, eax, 25098
	sar	eax, 16
	movsx	edx, word ptr [ebp - 24]
	add	edx, eax
	mov	si, dx
	mov	word ptr [ebp - 24], si
	mov	dword ptr [esp], ecx
	call	_ride_ratings_apply_intensity_penalty
	lea	eax, dword ptr [ebp - 24]
	mov	ecx, dword ptr [ebp - 12]
	mov	dword ptr [esp], ecx
	mov	dword ptr [esp + 4], eax
	call	_ride_ratings_apply_adjustments
	mov	eax, dword ptr [ebp - 12]
	mov	ecx, dword ptr [ebp - 24]
	mov	dword ptr [eax + 320], ecx
	mov	si, word ptr [ebp - 20]
	mov	word ptr [eax + 324], si
	mov	eax, dword ptr [ebp - 12]
	mov	dword ptr [esp], eax
	call	_ride_compute_upkeep
	mov	ecx, dword ptr [ebp - 12]
	mov	word ptr [ecx + 386], ax
	mov	ecx, dword ptr [ebp - 12]
	movzx	edx, byte ptr [ecx + 333]
	or	edx, 2
	mov	bl, dl
	mov	byte ptr [ecx + 333], bl
	mov	ecx, dword ptr [ebp - 12]
	movzx	edx, byte ptr [ecx + 276]
	and	edx, 31
	mov	bl, dl
	mov	byte ptr [ecx + 276], bl
	mov	ecx, dword ptr [ebp - 12]
	movzx	edx, byte ptr [ecx + 276]
	mov	bl, dl
	mov	byte ptr [ecx + 276], bl
	add	esp, 32
	pop	esi
	pop	ebx
	pop	ebp
	ret

	.align	4, 0x90
_ride_ratings_calculate_elevator:       ## @ride_ratings_calculate_elevator
## BB#0:
	push	ebp
	mov	ebp, esp
	push	ebx
	push	edi
	push	esi
	sub	esp, 28
	mov	eax, dword ptr [ebp + 8]
	mov	dword ptr [ebp - 16], eax
	mov	eax, dword ptr [ebp - 16]
	mov	eax, dword ptr [eax + 464]
	and	eax, 2
	cmp	eax, 0
	jne	LBB26_2
## BB#1:
	jmp	LBB26_4
LBB26_2:
	mov	eax, dword ptr [ebp - 16]
	mov	byte ptr [eax + 408], 15
	mov	eax, dword ptr [ebp - 16]
	mov	dword ptr [esp], eax
	call	_set_unreliability_factor
	mov	word ptr [ebp - 24], 111
	mov	word ptr [ebp - 22], 35
	mov	word ptr [ebp - 20], 30
	mov	eax, dword ptr [ebp - 16]
	mov	dword ptr [esp], eax
	call	_ride_get_total_length
	sar	eax, 16
	mov	dword ptr [ebp - 28], eax
	imul	eax, dword ptr [ebp - 28], 45875
	sar	eax, 16
	movsx	ecx, word ptr [ebp - 24]
	add	ecx, eax
	mov	dx, cx
	mov	word ptr [ebp - 24], dx
	call	_sub_65E277
	imul	eax, eax, 11183
	sar	eax, 16
	movsx	ecx, word ptr [ebp - 24]
	add	ecx, eax
	mov	dx, cx
	mov	word ptr [ebp - 24], dx
	mov	eax, dword ptr [ebp - 16]
	mov	dword ptr [esp], eax
	call	_ride_ratings_get_scenery_score
	lea	ecx, dword ptr [ebp - 24]
	imul	eax, eax, 83662
	sar	eax, 16
	movsx	esi, word ptr [ebp - 24]
	add	esi, eax
	mov	dx, si
	mov	word ptr [ebp - 24], dx
	imul	eax, dword ptr [ebp - 28], 26214
	sar	eax, 16
	movsx	esi, word ptr [ebp - 20]
	add	esi, eax
	mov	dx, si
	mov	word ptr [ebp - 20], dx
	mov	dword ptr [esp], ecx
	call	_ride_ratings_apply_intensity_penalty
	lea	eax, dword ptr [ebp - 24]
	mov	ecx, dword ptr [ebp - 16]
	mov	dword ptr [esp], ecx
	mov	dword ptr [esp + 4], eax
	call	_ride_ratings_apply_adjustments
	mov	eax, dword ptr [ebp - 16]
	mov	ecx, dword ptr [ebp - 24]
	mov	dword ptr [eax + 320], ecx
	mov	dx, word ptr [ebp - 20]
	mov	word ptr [eax + 324], dx
	mov	eax, dword ptr [ebp - 16]
	mov	dword ptr [esp], eax
	call	_ride_compute_upkeep
	mov	ecx, dword ptr [ebp - 16]
	mov	word ptr [ecx + 386], ax
	mov	ecx, dword ptr [ebp - 16]
	movzx	esi, byte ptr [ecx + 333]
	or	esi, 2
	mov	ebx, esi
	mov	al, bl
	mov	byte ptr [ecx + 333], al
	mov	ecx, dword ptr [ebp - 16]
	movzx	esi, byte ptr [ecx + 276]
	and	esi, 31
	mov	ebx, esi
	mov	al, bl
	mov	byte ptr [ecx + 276], al
	mov	ecx, dword ptr [ebp - 16]
	movzx	esi, byte ptr [ecx + 276]
	or	esi, 224
	mov	ebx, esi
	mov	al, bl
	mov	byte ptr [ecx + 276], al
	mov	ecx, dword ptr [ebp - 16]
	mov	dword ptr [esp], ecx
	call	_sub_65E72D
	sar	eax, 8
	cmp	eax, 5
	jl	LBB26_4
## BB#3:
	mov	eax, 4
	mov	ecx, dword ptr [ebp - 16]
	movsx	edx, word ptr [ecx + 320]
	mov	dword ptr [ebp - 32], eax ## 4-byte Spill
	mov	eax, edx
	cdq
	mov	esi, dword ptr [ebp - 32] ## 4-byte Reload
	idiv	esi
	mov	di, ax
	mov	word ptr [ecx + 320], di
LBB26_4:
	add	esp, 28
	pop	esi
	pop	edi
	pop	ebx
	pop	ebp
	ret

	.align	4, 0x90
_ride_ratings_calculate_atm:            ## @ride_ratings_calculate_atm
## BB#0:
	push	ebp
	mov	ebp, esp
	push	ebx
	sub	esp, 20
	mov	eax, dword ptr [ebp + 8]
	mov	dword ptr [ebp - 8], eax
	mov	eax, dword ptr [ebp - 8]
	mov	dword ptr [esp], eax
	call	_ride_compute_upkeep
	mov	ecx, dword ptr [ebp - 8]
	mov	word ptr [ecx + 386], ax
	mov	ecx, dword ptr [ebp - 8]
	movzx	edx, byte ptr [ecx + 333]
	or	edx, 2
	mov	bl, dl
	mov	byte ptr [ecx + 333], bl
	add	esp, 20
	pop	ebx
	pop	ebp
	ret

	.align	4, 0x90
_ride_ratings_calculate_twist:          ## @ride_ratings_calculate_twist
## BB#0:
	push	ebp
	mov	ebp, esp
	push	ebx
	push	esi
	sub	esp, 32
	mov	eax, dword ptr [ebp + 8]
	mov	dword ptr [ebp - 12], eax
	mov	eax, dword ptr [ebp - 12]
	mov	ecx, dword ptr [eax + 464]
	or	ecx, 2
	mov	dword ptr [eax + 464], ecx
	mov	eax, dword ptr [ebp - 12]
	mov	ecx, dword ptr [eax + 464]
	or	ecx, 8
	mov	dword ptr [eax + 464], ecx
	mov	eax, dword ptr [ebp - 12]
	mov	byte ptr [eax + 408], 16
	mov	eax, dword ptr [ebp - 12]
	mov	dword ptr [esp], eax
	call	_set_unreliability_factor
	mov	word ptr [ebp - 24], 113
	mov	word ptr [ebp - 22], 97
	mov	word ptr [ebp - 20], 190
	mov	eax, dword ptr [ebp - 12]
	movzx	eax, byte ptr [eax + 208]
	imul	eax, eax, 20
	movsx	ecx, word ptr [ebp - 24]
	add	ecx, eax
	mov	dx, cx
	mov	word ptr [ebp - 24], dx
	mov	eax, dword ptr [ebp - 12]
	movzx	eax, byte ptr [eax + 208]
	imul	eax, eax, 20
	movsx	ecx, word ptr [ebp - 22]
	add	ecx, eax
	mov	dx, cx
	mov	word ptr [ebp - 22], dx
	mov	eax, dword ptr [ebp - 12]
	movzx	eax, byte ptr [eax + 208]
	imul	eax, eax, 20
	movsx	ecx, word ptr [ebp - 20]
	add	ecx, eax
	mov	dx, cx
	mov	word ptr [ebp - 20], dx
	mov	eax, dword ptr [ebp - 12]
	mov	dword ptr [esp], eax
	call	_ride_ratings_get_scenery_score
	lea	ecx, dword ptr [ebp - 24]
	imul	eax, eax, 13943
	sar	eax, 16
	movsx	esi, word ptr [ebp - 24]
	add	esi, eax
	mov	dx, si
	mov	word ptr [ebp - 24], dx
	mov	dword ptr [esp], ecx
	call	_ride_ratings_apply_intensity_penalty
	lea	eax, dword ptr [ebp - 24]
	mov	ecx, dword ptr [ebp - 12]
	mov	dword ptr [esp], ecx
	mov	dword ptr [esp + 4], eax
	call	_ride_ratings_apply_adjustments
	mov	eax, dword ptr [ebp - 12]
	mov	ecx, dword ptr [ebp - 24]
	mov	dword ptr [eax + 320], ecx
	mov	dx, word ptr [ebp - 20]
	mov	word ptr [eax + 324], dx
	mov	eax, dword ptr [ebp - 12]
	mov	dword ptr [esp], eax
	call	_ride_compute_upkeep
	mov	ecx, dword ptr [ebp - 12]
	mov	word ptr [ecx + 386], ax
	mov	ecx, dword ptr [ebp - 12]
	movzx	esi, byte ptr [ecx + 333]
	or	esi, 2
	mov	ebx, esi
	mov	al, bl
	mov	byte ptr [ecx + 333], al
	mov	ecx, dword ptr [ebp - 12]
	movzx	esi, byte ptr [ecx + 276]
	and	esi, 31
	mov	ebx, esi
	mov	al, bl
	mov	byte ptr [ecx + 276], al
	mov	ecx, dword ptr [ebp - 12]
	movzx	esi, byte ptr [ecx + 276]
	mov	ebx, esi
	mov	al, bl
	mov	byte ptr [ecx + 276], al
	add	esp, 32
	pop	esi
	pop	ebx
	pop	ebp
	ret

	.align	4, 0x90
_ride_ratings_calculate_haunted_house:  ## @ride_ratings_calculate_haunted_house
## BB#0:
	push	ebp
	mov	ebp, esp
	push	ebx
	push	esi
	sub	esp, 32
	mov	eax, dword ptr [ebp + 8]
	mov	dword ptr [ebp - 12], eax
	mov	eax, dword ptr [ebp - 12]
	mov	ecx, dword ptr [eax + 464]
	or	ecx, 2
	mov	dword ptr [eax + 464], ecx
	mov	eax, dword ptr [ebp - 12]
	mov	ecx, dword ptr [eax + 464]
	or	ecx, 8
	mov	dword ptr [eax + 464], ecx
	mov	eax, dword ptr [ebp - 12]
	mov	byte ptr [eax + 408], 8
	mov	eax, dword ptr [ebp - 12]
	mov	dword ptr [esp], eax
	call	_set_unreliability_factor
	lea	eax, dword ptr [ebp - 24]
	mov	word ptr [ebp - 24], 341
	mov	word ptr [ebp - 22], 153
	mov	word ptr [ebp - 20], 10
	mov	dword ptr [esp], eax
	call	_ride_ratings_apply_intensity_penalty
	lea	eax, dword ptr [ebp - 24]
	mov	ecx, dword ptr [ebp - 12]
	mov	dword ptr [esp], ecx
	mov	dword ptr [esp + 4], eax
	call	_ride_ratings_apply_adjustments
	mov	eax, dword ptr [ebp - 12]
	mov	ecx, dword ptr [ebp - 24]
	mov	dword ptr [eax + 320], ecx
	mov	dx, word ptr [ebp - 20]
	mov	word ptr [eax + 324], dx
	mov	eax, dword ptr [ebp - 12]
	mov	dword ptr [esp], eax
	call	_ride_compute_upkeep
	mov	ecx, dword ptr [ebp - 12]
	mov	word ptr [ecx + 386], ax
	mov	ecx, dword ptr [ebp - 12]
	movzx	esi, byte ptr [ecx + 333]
	or	esi, 2
	mov	ebx, esi
	mov	al, bl
	mov	byte ptr [ecx + 333], al
	mov	ecx, dword ptr [ebp - 12]
	movzx	esi, byte ptr [ecx + 276]
	and	esi, 31
	mov	ebx, esi
	mov	al, bl
	mov	byte ptr [ecx + 276], al
	mov	ecx, dword ptr [ebp - 12]
	movzx	esi, byte ptr [ecx + 276]
	or	esi, 224
	mov	ebx, esi
	mov	al, bl
	mov	byte ptr [ecx + 276], al
	add	esp, 32
	pop	esi
	pop	ebx
	pop	ebp
	ret

	.align	4, 0x90
_ride_ratings_calculate_first_aid:      ## @ride_ratings_calculate_first_aid
## BB#0:
	push	ebp
	mov	ebp, esp
	push	ebx
	sub	esp, 20
	mov	eax, dword ptr [ebp + 8]
	mov	dword ptr [ebp - 8], eax
	mov	eax, dword ptr [ebp - 8]
	mov	dword ptr [esp], eax
	call	_ride_compute_upkeep
	mov	ecx, dword ptr [ebp - 8]
	mov	word ptr [ecx + 386], ax
	mov	ecx, dword ptr [ebp - 8]
	movzx	edx, byte ptr [ecx + 333]
	or	edx, 2
	mov	bl, dl
	mov	byte ptr [ecx + 333], bl
	add	esp, 20
	pop	ebx
	pop	ebp
	ret

	.align	4, 0x90
_ride_ratings_calculate_circus_show:    ## @ride_ratings_calculate_circus_show
## BB#0:
	push	ebp
	mov	ebp, esp
	push	ebx
	push	esi
	sub	esp, 32
	mov	eax, dword ptr [ebp + 8]
	mov	dword ptr [ebp - 12], eax
	mov	eax, dword ptr [ebp - 12]
	mov	ecx, dword ptr [eax + 464]
	or	ecx, 2
	mov	dword ptr [eax + 464], ecx
	mov	eax, dword ptr [ebp - 12]
	mov	ecx, dword ptr [eax + 464]
	or	ecx, 8
	mov	dword ptr [eax + 464], ecx
	mov	eax, dword ptr [ebp - 12]
	mov	byte ptr [eax + 408], 9
	mov	eax, dword ptr [ebp - 12]
	mov	dword ptr [esp], eax
	call	_set_unreliability_factor
	lea	eax, dword ptr [ebp - 24]
	mov	word ptr [ebp - 24], 210
	mov	word ptr [ebp - 22], 30
	mov	word ptr [ebp - 20], 0
	mov	dword ptr [esp], eax
	call	_ride_ratings_apply_intensity_penalty
	lea	eax, dword ptr [ebp - 24]
	mov	ecx, dword ptr [ebp - 12]
	mov	dword ptr [esp], ecx
	mov	dword ptr [esp + 4], eax
	call	_ride_ratings_apply_adjustments
	mov	eax, dword ptr [ebp - 12]
	mov	ecx, dword ptr [ebp - 24]
	mov	dword ptr [eax + 320], ecx
	mov	dx, word ptr [ebp - 20]
	mov	word ptr [eax + 324], dx
	mov	eax, dword ptr [ebp - 12]
	mov	dword ptr [esp], eax
	call	_ride_compute_upkeep
	mov	ecx, dword ptr [ebp - 12]
	mov	word ptr [ecx + 386], ax
	mov	ecx, dword ptr [ebp - 12]
	movzx	esi, byte ptr [ecx + 333]
	or	esi, 2
	mov	ebx, esi
	mov	al, bl
	mov	byte ptr [ecx + 333], al
	mov	ecx, dword ptr [ebp - 12]
	movzx	esi, byte ptr [ecx + 276]
	and	esi, 31
	mov	ebx, esi
	mov	al, bl
	mov	byte ptr [ecx + 276], al
	mov	ecx, dword ptr [ebp - 12]
	movzx	esi, byte ptr [ecx + 276]
	or	esi, 224
	mov	ebx, esi
	mov	al, bl
	mov	byte ptr [ecx + 276], al
	add	esp, 32
	pop	esi
	pop	ebx
	pop	ebp
	ret

	.align	4, 0x90
_ride_ratings_calculate_mini_golf:      ## @ride_ratings_calculate_mini_golf
## BB#0:
	push	ebp
	mov	ebp, esp
	push	ebx
	push	esi
	sub	esp, 64
	mov	eax, dword ptr [ebp + 8]
	mov	dword ptr [ebp - 12], eax
	mov	eax, dword ptr [ebp - 12]
	mov	eax, dword ptr [eax + 464]
	and	eax, 2
	cmp	eax, 0
	jne	LBB32_2
## BB#1:
	jmp	LBB32_8
LBB32_2:
	mov	eax, dword ptr [ebp - 12]
	mov	byte ptr [eax + 408], 0
	mov	eax, dword ptr [ebp - 12]
	mov	dword ptr [esp], eax
	call	_set_unreliability_factor
	mov	word ptr [ebp - 24], 150
	mov	word ptr [ebp - 22], 90
	mov	word ptr [ebp - 20], 0
	mov	eax, dword ptr [ebp - 12]
	mov	dword ptr [esp], eax
	call	_ride_get_total_length
	mov	ecx, 6000
	sar	eax, 16
	mov	dword ptr [ebp - 36], eax
	cmp	ecx, dword ptr [ebp - 36]
	jge	LBB32_4
## BB#3:
	mov	eax, 6000
	mov	dword ptr [ebp - 60], eax ## 4-byte Spill
	jmp	LBB32_5
LBB32_4:
	mov	eax, dword ptr [ebp - 36]
	mov	dword ptr [ebp - 60], eax ## 4-byte Spill
LBB32_5:
	mov	eax, dword ptr [ebp - 60] ## 4-byte Reload
	lea	ecx, dword ptr [ebp - 48]
	imul	eax, eax, 873
	sar	eax, 16
	movsx	edx, word ptr [ebp - 24]
	add	edx, eax
	mov	si, dx
	mov	word ptr [ebp - 24], si
	mov	eax, dword ptr [ebp - 12]
	mov	dword ptr [esp], ecx
	mov	dword ptr [esp + 4], eax
	call	_sub_65DDD1
	sub	esp, 4
	lea	eax, dword ptr [ebp - 56]
	mov	ecx, dword ptr [ebp - 48]
	mov	dword ptr [ebp - 32], ecx
	mov	si, word ptr [ebp - 44]
	mov	word ptr [ebp - 28], si
	movsx	ecx, word ptr [ebp - 32]
	imul	ecx, ecx, 14860
	sar	ecx, 16
	movsx	edx, word ptr [ebp - 24]
	add	edx, ecx
	mov	si, dx
	mov	word ptr [ebp - 24], si
	mov	ecx, dword ptr [ebp - 12]
	mov	dword ptr [esp], eax
	mov	dword ptr [esp + 4], ecx
	call	_sub_65E1C2
	sub	esp, 4
	mov	eax, dword ptr [ebp - 56]
	mov	dword ptr [ebp - 32], eax
	mov	si, word ptr [ebp - 52]
	mov	word ptr [ebp - 28], si
	movsx	eax, word ptr [ebp - 32]
	imul	eax, eax, 5140
	sar	eax, 16
	movsx	ecx, word ptr [ebp - 24]
	add	ecx, eax
	mov	si, cx
	mov	word ptr [ebp - 24], si
	movsx	eax, word ptr [ebp - 30]
	imul	eax, eax, 6553
	sar	eax, 16
	movsx	ecx, word ptr [ebp - 22]
	add	ecx, eax
	mov	si, cx
	mov	word ptr [ebp - 22], si
	movsx	eax, word ptr [ebp - 28]
	imul	eax, eax, 4681
	sar	eax, 16
	movsx	ecx, word ptr [ebp - 20]
	add	ecx, eax
	mov	si, cx
	mov	word ptr [ebp - 20], si
	call	_sub_65E277
	imul	eax, eax, 15657
	sar	eax, 16
	movsx	ecx, word ptr [ebp - 24]
	add	ecx, eax
	mov	si, cx
	mov	word ptr [ebp - 24], si
	mov	eax, dword ptr [ebp - 12]
	mov	dword ptr [esp], eax
	call	_ride_ratings_get_scenery_score
	imul	eax, eax, 27887
	sar	eax, 16
	movsx	ecx, word ptr [ebp - 24]
	add	ecx, eax
	mov	si, cx
	mov	word ptr [ebp - 24], si
	mov	eax, dword ptr [ebp - 12]
	movzx	eax, byte ptr [eax + 276]
	and	eax, 31
	imul	eax, eax, 5
	movsx	ecx, word ptr [ebp - 24]
	add	ecx, eax
	mov	si, cx
	mov	word ptr [ebp - 24], si
	mov	eax, dword ptr [ebp - 12]
	movzx	eax, byte ptr [eax + 276]
	and	eax, 31
	cmp	eax, 0
	jne	LBB32_7
## BB#6:
	mov	eax, 2
	mov	ecx, 8
	movsx	edx, word ptr [ebp - 24]
	mov	dword ptr [ebp - 64], eax ## 4-byte Spill
	mov	eax, edx
	cdq
	idiv	ecx
	mov	si, ax
	mov	word ptr [ebp - 24], si
	movsx	eax, word ptr [ebp - 22]
	cdq
	mov	ecx, dword ptr [ebp - 64] ## 4-byte Reload
	idiv	ecx
	mov	si, ax
	mov	word ptr [ebp - 22], si
	movsx	eax, word ptr [ebp - 20]
	cdq
	idiv	ecx
	mov	si, ax
	mov	word ptr [ebp - 20], si
LBB32_7:
	lea	eax, dword ptr [ebp - 24]
	mov	dword ptr [esp], eax
	call	_ride_ratings_apply_intensity_penalty
	lea	eax, dword ptr [ebp - 24]
	mov	ecx, dword ptr [ebp - 12]
	mov	dword ptr [esp], ecx
	mov	dword ptr [esp + 4], eax
	call	_ride_ratings_apply_adjustments
	mov	eax, dword ptr [ebp - 12]
	mov	ecx, dword ptr [ebp - 24]
	mov	dword ptr [eax + 320], ecx
	mov	dx, word ptr [ebp - 20]
	mov	word ptr [eax + 324], dx
	mov	eax, dword ptr [ebp - 12]
	mov	dword ptr [esp], eax
	call	_ride_compute_upkeep
	mov	ecx, dword ptr [ebp - 12]
	mov	word ptr [ecx + 386], ax
	mov	ecx, dword ptr [ebp - 12]
	movzx	esi, byte ptr [ecx + 333]
	or	esi, 2
	mov	ebx, esi
	mov	al, bl
	mov	byte ptr [ecx + 333], al
	mov	ecx, dword ptr [ebp - 12]
	movzx	esi, byte ptr [ecx + 276]
	and	esi, 31
	mov	ebx, esi
	mov	al, bl
	mov	byte ptr [ecx + 276], al
	mov	ecx, dword ptr [ebp - 12]
	mov	dword ptr [esp], ecx
	call	_sub_65E72D
	shl	eax, 5
	mov	ecx, dword ptr [ebp - 12]
	movzx	esi, byte ptr [ecx + 276]
	or	esi, eax
	mov	eax, esi
                                        ## kill: AL<def> AL<kill> EAX<kill>
	mov	byte ptr [ecx + 276], al
LBB32_8:
	add	esp, 64
	pop	esi
	pop	ebx
	pop	ebp
	ret

	.align	4, 0x90
_ride_ratings_calculate_crooked_house:  ## @ride_ratings_calculate_crooked_house
## BB#0:
	push	ebp
	mov	ebp, esp
	push	ebx
	push	esi
	sub	esp, 32
	mov	eax, dword ptr [ebp + 8]
	mov	dword ptr [ebp - 12], eax
	mov	eax, dword ptr [ebp - 12]
	mov	ecx, dword ptr [eax + 464]
	or	ecx, 2
	mov	dword ptr [eax + 464], ecx
	mov	eax, dword ptr [ebp - 12]
	mov	ecx, dword ptr [eax + 464]
	or	ecx, 8
	mov	dword ptr [eax + 464], ecx
	mov	eax, dword ptr [ebp - 12]
	mov	byte ptr [eax + 408], 5
	mov	eax, dword ptr [ebp - 12]
	mov	dword ptr [esp], eax
	call	_set_unreliability_factor
	lea	eax, dword ptr [ebp - 24]
	mov	word ptr [ebp - 24], 215
	mov	word ptr [ebp - 22], 62
	mov	word ptr [ebp - 20], 34
	mov	dword ptr [esp], eax
	call	_ride_ratings_apply_intensity_penalty
	lea	eax, dword ptr [ebp - 24]
	mov	ecx, dword ptr [ebp - 12]
	mov	dword ptr [esp], ecx
	mov	dword ptr [esp + 4], eax
	call	_ride_ratings_apply_adjustments
	mov	eax, dword ptr [ebp - 12]
	mov	ecx, dword ptr [ebp - 24]
	mov	dword ptr [eax + 320], ecx
	mov	dx, word ptr [ebp - 20]
	mov	word ptr [eax + 324], dx
	mov	eax, dword ptr [ebp - 12]
	mov	dword ptr [esp], eax
	call	_ride_compute_upkeep
	mov	ecx, dword ptr [ebp - 12]
	mov	word ptr [ecx + 386], ax
	mov	ecx, dword ptr [ebp - 12]
	movzx	esi, byte ptr [ecx + 333]
	or	esi, 2
	mov	ebx, esi
	mov	al, bl
	mov	byte ptr [ecx + 333], al
	mov	ecx, dword ptr [ebp - 12]
	movzx	esi, byte ptr [ecx + 276]
	and	esi, 31
	mov	ebx, esi
	mov	al, bl
	mov	byte ptr [ecx + 276], al
	mov	ecx, dword ptr [ebp - 12]
	movzx	esi, byte ptr [ecx + 276]
	or	esi, 224
	mov	ebx, esi
	mov	al, bl
	mov	byte ptr [ecx + 276], al
	add	esp, 32
	pop	esi
	pop	ebx
	pop	ebp
	ret

	.align	4, 0x90
_ride_ratings_calculate_magic_carpet:   ## @ride_ratings_calculate_magic_carpet
## BB#0:
	push	ebp
	mov	ebp, esp
	push	ebx
	push	esi
	sub	esp, 32
	mov	eax, dword ptr [ebp + 8]
	mov	dword ptr [ebp - 12], eax
	mov	eax, dword ptr [ebp - 12]
	mov	ecx, dword ptr [eax + 464]
	or	ecx, 2
	mov	dword ptr [eax + 464], ecx
	mov	eax, dword ptr [ebp - 12]
	mov	ecx, dword ptr [eax + 464]
	or	ecx, 8
	mov	dword ptr [eax + 464], ecx
	mov	eax, dword ptr [ebp - 12]
	mov	byte ptr [eax + 408], 16
	mov	eax, dword ptr [ebp - 12]
	mov	dword ptr [esp], eax
	call	_set_unreliability_factor
	mov	word ptr [ebp - 24], 245
	mov	word ptr [ebp - 22], 160
	mov	word ptr [ebp - 20], 260
	mov	eax, dword ptr [ebp - 12]
	movzx	eax, byte ptr [eax + 208]
	imul	eax, eax, 10
	movsx	ecx, word ptr [ebp - 24]
	add	ecx, eax
	mov	dx, cx
	mov	word ptr [ebp - 24], dx
	mov	eax, dword ptr [ebp - 12]
	movzx	eax, byte ptr [eax + 208]
	imul	eax, eax, 20
	movsx	ecx, word ptr [ebp - 22]
	add	ecx, eax
	mov	dx, cx
	mov	word ptr [ebp - 22], dx
	mov	eax, dword ptr [ebp - 12]
	movzx	eax, byte ptr [eax + 208]
	imul	eax, eax, 20
	movsx	ecx, word ptr [ebp - 20]
	add	ecx, eax
	mov	dx, cx
	mov	word ptr [ebp - 20], dx
	mov	eax, dword ptr [ebp - 12]
	mov	dword ptr [esp], eax
	call	_ride_ratings_get_scenery_score
	lea	ecx, dword ptr [ebp - 24]
	imul	eax, eax, 11155
	sar	eax, 16
	movsx	esi, word ptr [ebp - 24]
	add	esi, eax
	mov	dx, si
	mov	word ptr [ebp - 24], dx
	mov	dword ptr [esp], ecx
	call	_ride_ratings_apply_intensity_penalty
	lea	eax, dword ptr [ebp - 24]
	mov	ecx, dword ptr [ebp - 12]
	mov	dword ptr [esp], ecx
	mov	dword ptr [esp + 4], eax
	call	_ride_ratings_apply_adjustments
	mov	eax, dword ptr [ebp - 12]
	mov	ecx, dword ptr [ebp - 24]
	mov	dword ptr [eax + 320], ecx
	mov	dx, word ptr [ebp - 20]
	mov	word ptr [eax + 324], dx
	mov	eax, dword ptr [ebp - 12]
	mov	dword ptr [esp], eax
	call	_ride_compute_upkeep
	mov	ecx, dword ptr [ebp - 12]
	mov	word ptr [ecx + 386], ax
	mov	ecx, dword ptr [ebp - 12]
	movzx	esi, byte ptr [ecx + 333]
	or	esi, 2
	mov	ebx, esi
	mov	al, bl
	mov	byte ptr [ecx + 333], al
	mov	ecx, dword ptr [ebp - 12]
	movzx	esi, byte ptr [ecx + 276]
	and	esi, 31
	mov	ebx, esi
	mov	al, bl
	mov	byte ptr [ecx + 276], al
	mov	ecx, dword ptr [ebp - 12]
	movzx	esi, byte ptr [ecx + 276]
	mov	ebx, esi
	mov	al, bl
	mov	byte ptr [ecx + 276], al
	add	esp, 32
	pop	esi
	pop	ebx
	pop	ebp
	ret

	.align	4, 0x90
_ride_ratings_calculate_enterprise:     ## @ride_ratings_calculate_enterprise
## BB#0:
	push	ebp
	mov	ebp, esp
	push	ebx
	push	esi
	sub	esp, 32
	mov	eax, dword ptr [ebp + 8]
	mov	dword ptr [ebp - 12], eax
	mov	eax, dword ptr [ebp - 12]
	mov	ecx, dword ptr [eax + 464]
	or	ecx, 2
	mov	dword ptr [eax + 464], ecx
	mov	eax, dword ptr [ebp - 12]
	mov	ecx, dword ptr [eax + 464]
	or	ecx, 8
	mov	dword ptr [eax + 464], ecx
	mov	eax, dword ptr [ebp - 12]
	mov	byte ptr [eax + 408], 22
	mov	eax, dword ptr [ebp - 12]
	mov	dword ptr [esp], eax
	call	_set_unreliability_factor
	mov	word ptr [ebp - 24], 360
	mov	word ptr [ebp - 22], 455
	mov	word ptr [ebp - 20], 572
	mov	eax, dword ptr [ebp - 12]
	movzx	eax, byte ptr [eax + 208]
	movsx	ecx, word ptr [ebp - 24]
	add	ecx, eax
	mov	dx, cx
	mov	word ptr [ebp - 24], dx
	mov	eax, dword ptr [ebp - 12]
	movzx	eax, byte ptr [eax + 208]
	shl	eax, 4
	movsx	ecx, word ptr [ebp - 22]
	add	ecx, eax
	mov	dx, cx
	mov	word ptr [ebp - 22], dx
	mov	eax, dword ptr [ebp - 12]
	movzx	eax, byte ptr [eax + 208]
	shl	eax, 4
	movsx	ecx, word ptr [ebp - 20]
	add	ecx, eax
	mov	dx, cx
	mov	word ptr [ebp - 20], dx
	mov	eax, dword ptr [ebp - 12]
	mov	dword ptr [esp], eax
	call	_ride_ratings_get_scenery_score
	lea	ecx, dword ptr [ebp - 24]
	imul	eax, eax, 19521
	sar	eax, 16
	movsx	esi, word ptr [ebp - 24]
	add	esi, eax
	mov	dx, si
	mov	word ptr [ebp - 24], dx
	mov	dword ptr [esp], ecx
	call	_ride_ratings_apply_intensity_penalty
	lea	eax, dword ptr [ebp - 24]
	mov	ecx, dword ptr [ebp - 12]
	mov	dword ptr [esp], ecx
	mov	dword ptr [esp + 4], eax
	call	_ride_ratings_apply_adjustments
	mov	eax, dword ptr [ebp - 12]
	mov	ecx, dword ptr [ebp - 24]
	mov	dword ptr [eax + 320], ecx
	mov	dx, word ptr [ebp - 20]
	mov	word ptr [eax + 324], dx
	mov	eax, dword ptr [ebp - 12]
	mov	dword ptr [esp], eax
	call	_ride_compute_upkeep
	mov	ecx, dword ptr [ebp - 12]
	mov	word ptr [ecx + 386], ax
	mov	ecx, dword ptr [ebp - 12]
	movzx	esi, byte ptr [ecx + 333]
	or	esi, 2
	mov	ebx, esi
	mov	al, bl
	mov	byte ptr [ecx + 333], al
	mov	ecx, dword ptr [ebp - 12]
	movzx	esi, byte ptr [ecx + 276]
	and	esi, 31
	mov	ebx, esi
	mov	al, bl
	mov	byte ptr [ecx + 276], al
	mov	ecx, dword ptr [ebp - 12]
	movzx	esi, byte ptr [ecx + 276]
	or	esi, 96
	mov	ebx, esi
	mov	al, bl
	mov	byte ptr [ecx + 276], al
	add	esp, 32
	pop	esi
	pop	ebx
	pop	ebp
	ret

	.align	4, 0x90
_set_unreliability_factor:              ## @set_unreliability_factor
## BB#0:
	push	ebp
	mov	ebp, esp
	push	esi
	sub	esp, 8
	mov	eax, dword ptr [ebp + 8]
	mov	ecx, 9951177
	mov	dword ptr [ebp - 8], eax
	mov	eax, dword ptr [ebp - 8]
	movzx	eax, byte ptr [eax]
	shl	eax, 2
	mov	dl, byte ptr [ecx + eax]
	mov	byte ptr [ebp - 9], dl
	mov	eax, dword ptr [ebp - 8]
	movzx	eax, byte ptr [eax + 461]
	movzx	ecx, byte ptr [ebp - 9]
	sub	eax, ecx
	shl	eax, 1
	mov	ecx, dword ptr [ebp - 8]
	movzx	esi, byte ptr [ecx + 408]
	add	esi, eax
	mov	eax, esi
	mov	dl, al
	mov	byte ptr [ecx + 408], dl
	add	esp, 8
	pop	esi
	pop	ebp
	ret

	.align	4, 0x90
_ride_ratings_get_scenery_score:        ## @ride_ratings_get_scenery_score
## BB#0:
	push	ebp
	mov	ebp, esp
	sub	esp, 72
	mov	eax, dword ptr [ebp + 8]
	mov	dword ptr [ebp - 8], eax
	mov	dword ptr [ebp - 12], 0
LBB37_1:                                ## =>This Inner Loop Header: Depth=1
	cmp	dword ptr [ebp - 12], 4
	jge	LBB37_6
## BB#2:                                ##   in Loop: Header=BB37_1 Depth=1
	mov	eax, dword ptr [ebp - 12]
	mov	ecx, dword ptr [ebp - 8]
	mov	dx, word ptr [ecx + 2*eax + 82]
	mov	word ptr [ebp - 42], dx
	movzx	eax, word ptr [ebp - 42]
	cmp	eax, 65535
	je	LBB37_4
## BB#3:
	jmp	LBB37_6
LBB37_4:                                ##   in Loop: Header=BB37_1 Depth=1
	jmp	LBB37_5
LBB37_5:                                ##   in Loop: Header=BB37_1 Depth=1
	mov	eax, dword ptr [ebp - 12]
	add	eax, 1
	mov	dword ptr [ebp - 12], eax
	jmp	LBB37_1
LBB37_6:
	cmp	dword ptr [ebp - 12], 4
	jne	LBB37_8
## BB#7:
	mov	dword ptr [ebp - 4], 0
	jmp	LBB37_38
LBB37_8:
	mov	eax, dword ptr [ebp - 8]
	movzx	eax, byte ptr [eax]
	cmp	eax, 20
	jne	LBB37_10
## BB#9:
	mov	eax, dword ptr [ebp - 8]
	mov	cx, word ptr [eax + 106]
	mov	word ptr [ebp - 42], cx
LBB37_10:
	movzx	eax, word ptr [ebp - 42]
	and	eax, 255
	mov	dword ptr [ebp - 16], eax
	movzx	eax, word ptr [ebp - 42]
	sar	eax, 8
	mov	dword ptr [ebp - 20], eax
	mov	eax, dword ptr [ebp - 16]
	shl	eax, 5
	mov	ecx, dword ptr [ebp - 20]
	shl	ecx, 5
	mov	dword ptr [esp], eax
	mov	dword ptr [esp + 4], ecx
	call	_map_element_height
	mov	dword ptr [ebp - 24], eax
	mov	eax, dword ptr [ebp - 24]
	mov	ecx, dword ptr [ebp - 12]
	mov	edx, dword ptr [ebp - 8]
	movzx	ecx, byte ptr [edx + ecx + 90]
	shl	ecx, 3
	cmp	eax, ecx
	jle	LBB37_12
## BB#11:
	mov	dword ptr [ebp - 4], 40
	jmp	LBB37_38
LBB37_12:
	mov	dword ptr [ebp - 40], 0
	mov	eax, dword ptr [ebp - 20]
	sub	eax, 5
	cmp	eax, 0
	jle	LBB37_14
## BB#13:
	mov	eax, dword ptr [ebp - 20]
	sub	eax, 5
	mov	dword ptr [ebp - 52], eax ## 4-byte Spill
	jmp	LBB37_15
LBB37_14:
	mov	eax, 0
	mov	dword ptr [ebp - 52], eax ## 4-byte Spill
	jmp	LBB37_15
LBB37_15:
	mov	eax, dword ptr [ebp - 52] ## 4-byte Reload
	mov	dword ptr [ebp - 32], eax
LBB37_16:                               ## =>This Loop Header: Depth=1
                                        ##     Child Loop BB37_21 Depth 2
                                        ##       Child Loop BB37_23 Depth 3
	mov	eax, dword ptr [ebp - 32]
	mov	ecx, dword ptr [ebp - 20]
	add	ecx, 5
	cmp	eax, ecx
	jg	LBB37_34
## BB#17:                               ##   in Loop: Header=BB37_16 Depth=1
	mov	eax, dword ptr [ebp - 16]
	sub	eax, 5
	cmp	eax, 0
	jle	LBB37_19
## BB#18:                               ##   in Loop: Header=BB37_16 Depth=1
	mov	eax, dword ptr [ebp - 16]
	sub	eax, 5
	mov	dword ptr [ebp - 56], eax ## 4-byte Spill
	jmp	LBB37_20
LBB37_19:                               ##   in Loop: Header=BB37_16 Depth=1
	mov	eax, 0
	mov	dword ptr [ebp - 56], eax ## 4-byte Spill
	jmp	LBB37_20
LBB37_20:                               ##   in Loop: Header=BB37_16 Depth=1
	mov	eax, dword ptr [ebp - 56] ## 4-byte Reload
	mov	dword ptr [ebp - 28], eax
LBB37_21:                               ##   Parent Loop BB37_16 Depth=1
                                        ## =>  This Loop Header: Depth=2
                                        ##       Child Loop BB37_23 Depth 3
	mov	eax, dword ptr [ebp - 28]
	mov	ecx, dword ptr [ebp - 16]
	add	ecx, 5
	cmp	eax, ecx
	jg	LBB37_32
## BB#22:                               ##   in Loop: Header=BB37_21 Depth=2
	mov	eax, dword ptr [ebp - 28]
	mov	ecx, dword ptr [ebp - 32]
	mov	dword ptr [esp], eax
	mov	dword ptr [esp + 4], ecx
	call	_map_get_first_element_at
	mov	dword ptr [ebp - 48], eax
LBB37_23:                               ##   Parent Loop BB37_16 Depth=1
                                        ##     Parent Loop BB37_21 Depth=2
                                        ## =>    This Inner Loop Header: Depth=3
	mov	eax, dword ptr [ebp - 48]
	movzx	eax, byte ptr [eax + 1]
	and	eax, 16
	cmp	eax, 0
	je	LBB37_25
## BB#24:                               ##   in Loop: Header=BB37_23 Depth=3
	jmp	LBB37_29
LBB37_25:                               ##   in Loop: Header=BB37_23 Depth=3
	mov	eax, dword ptr [ebp - 48]
	mov	dword ptr [esp], eax
	call	_map_element_get_type
	mov	dword ptr [ebp - 36], eax
	cmp	dword ptr [ebp - 36], 12
	je	LBB37_27
## BB#26:                               ##   in Loop: Header=BB37_23 Depth=3
	cmp	dword ptr [ebp - 36], 24
	jne	LBB37_28
LBB37_27:                               ##   in Loop: Header=BB37_23 Depth=3
	mov	eax, dword ptr [ebp - 40]
	add	eax, 1
	mov	dword ptr [ebp - 40], eax
LBB37_28:                               ##   in Loop: Header=BB37_23 Depth=3
	jmp	LBB37_29
LBB37_29:                               ##   in Loop: Header=BB37_23 Depth=3
	mov	eax, dword ptr [ebp - 48]
	mov	ecx, eax
	add	ecx, 8
	mov	dword ptr [ebp - 48], ecx
	mov	dword ptr [esp], eax
	call	_map_element_is_last_for_tile
	cmp	eax, 0
	setne	dl
	xor	dl, 1
	test	dl, 1
	jne	LBB37_23
## BB#30:                               ##   in Loop: Header=BB37_21 Depth=2
	jmp	LBB37_31
LBB37_31:                               ##   in Loop: Header=BB37_21 Depth=2
	mov	eax, dword ptr [ebp - 28]
	add	eax, 1
	mov	dword ptr [ebp - 28], eax
	jmp	LBB37_21
LBB37_32:                               ##   in Loop: Header=BB37_16 Depth=1
	jmp	LBB37_33
LBB37_33:                               ##   in Loop: Header=BB37_16 Depth=1
	mov	eax, dword ptr [ebp - 32]
	add	eax, 1
	mov	dword ptr [ebp - 32], eax
	jmp	LBB37_16
LBB37_34:
	cmp	dword ptr [ebp - 40], 47
	jge	LBB37_36
## BB#35:
	mov	eax, dword ptr [ebp - 40]
	mov	dword ptr [ebp - 60], eax ## 4-byte Spill
	jmp	LBB37_37
LBB37_36:
	mov	eax, 47
	mov	dword ptr [ebp - 60], eax ## 4-byte Spill
	jmp	LBB37_37
LBB37_37:
	mov	eax, dword ptr [ebp - 60] ## 4-byte Reload
	imul	eax, eax, 5
	mov	dword ptr [ebp - 4], eax
LBB37_38:
	mov	eax, dword ptr [ebp - 4]
	add	esp, 72
	pop	ebp
	ret

	.align	4, 0x90
_ride_ratings_apply_intensity_penalty:  ## @ride_ratings_apply_intensity_penalty
## BB#0:
	push	ebp
	mov	ebp, esp
	push	esi
	sub	esp, 20
	call	L38$pb
L38$pb:
	pop	eax
	mov	ecx, dword ptr [ebp + 8]
	mov	dword ptr [ebp - 8], ecx
	mov	ecx, dword ptr [ebp - 8]
	mov	dx, word ptr [ecx]
	mov	word ptr [ebp - 14], dx
	mov	dword ptr [ebp - 12], 0
	mov	dword ptr [ebp - 20], eax ## 4-byte Spill
LBB38_1:                                ## =>This Inner Loop Header: Depth=1
	cmp	dword ptr [ebp - 12], 5
	jae	LBB38_6
## BB#2:                                ##   in Loop: Header=BB38_1 Depth=1
	mov	eax, dword ptr [ebp - 8]
	movsx	eax, word ptr [eax + 2]
	mov	ecx, dword ptr [ebp - 12]
	mov	edx, dword ptr [ebp - 20] ## 4-byte Reload
	movsx	ecx, word ptr [edx + 2*ecx + _ride_ratings_apply_intensity_penalty.intensityBounds-L38$pb]
	cmp	eax, ecx
	jl	LBB38_4
## BB#3:                                ##   in Loop: Header=BB38_1 Depth=1
	mov	eax, 4
	movsx	ecx, word ptr [ebp - 14]
	mov	dword ptr [ebp - 24], eax ## 4-byte Spill
	mov	eax, ecx
	cdq
	mov	ecx, dword ptr [ebp - 24] ## 4-byte Reload
	idiv	ecx
	movsx	edx, word ptr [ebp - 14]
	sub	edx, eax
	mov	si, dx
	mov	word ptr [ebp - 14], si
LBB38_4:                                ##   in Loop: Header=BB38_1 Depth=1
	jmp	LBB38_5
LBB38_5:                                ##   in Loop: Header=BB38_1 Depth=1
	mov	eax, dword ptr [ebp - 12]
	add	eax, 1
	mov	dword ptr [ebp - 12], eax
	jmp	LBB38_1
LBB38_6:
	mov	ax, word ptr [ebp - 14]
	mov	ecx, dword ptr [ebp - 8]
	mov	word ptr [ecx], ax
	add	esp, 20
	pop	esi
	pop	ebp
	ret

	.align	4, 0x90
_ride_ratings_apply_adjustments:        ## @ride_ratings_apply_adjustments
## BB#0:
	push	ebp
	mov	ebp, esp
	push	edi
	push	esi
	sub	esp, 24
	call	L39$pb
L39$pb:
	pop	eax
	mov	ecx, dword ptr [ebp + 12]
	mov	edx, dword ptr [ebp + 8]
	mov	eax, dword ptr [eax + L_gRideTypeList$non_lazy_ptr-L39$pb]
	mov	dword ptr [ebp - 12], edx
	mov	dword ptr [ebp - 16], ecx
	mov	ecx, dword ptr [ebp - 12]
	movzx	ecx, byte ptr [ecx + 1]
	mov	eax, dword ptr [eax]
	mov	eax, dword ptr [eax + 4*ecx]
	mov	dword ptr [ebp - 20], eax
	mov	eax, dword ptr [ebp - 16]
	movsx	eax, word ptr [eax]
	mov	ecx, dword ptr [ebp - 20]
	movsx	ecx, byte ptr [ecx + 434]
	imul	eax, ecx
	sar	eax, 7
	mov	ecx, dword ptr [ebp - 16]
	movsx	edx, word ptr [ecx]
	add	edx, eax
	mov	si, dx
	mov	word ptr [ecx], si
	mov	eax, dword ptr [ebp - 16]
	movsx	eax, word ptr [eax + 2]
	mov	ecx, dword ptr [ebp - 20]
	movsx	ecx, byte ptr [ecx + 435]
	imul	eax, ecx
	sar	eax, 7
	mov	ecx, dword ptr [ebp - 16]
	movsx	edx, word ptr [ecx + 2]
	add	edx, eax
	mov	si, dx
	mov	word ptr [ecx + 2], si
	mov	eax, dword ptr [ebp - 16]
	movsx	eax, word ptr [eax + 4]
	mov	ecx, dword ptr [ebp - 20]
	movsx	ecx, byte ptr [ecx + 436]
	imul	eax, ecx
	sar	eax, 7
	mov	ecx, dword ptr [ebp - 16]
	movsx	edx, word ptr [ecx + 4]
	add	edx, eax
	mov	si, dx
	mov	word ptr [ecx + 4], si
	mov	eax, dword ptr [ebp - 12]
	movzx	eax, byte ptr [eax]
	shl	eax, 3
	add	eax, 9950450
	mov	si, word ptr [eax]
	mov	word ptr [ebp - 22], si
	movzx	eax, word ptr [ebp - 22]
	and	eax, 128
	cmp	eax, 0
	je	LBB39_7
## BB#1:
	mov	eax, dword ptr [ebp - 12]
	mov	cx, word ptr [eax + 500]
	mov	word ptr [ebp - 24], cx
	mov	eax, dword ptr [ebp - 20]
	mov	eax, dword ptr [eax + 8]
	and	eax, 2048
	cmp	eax, 0
	je	LBB39_5
## BB#2:
	movzx	eax, word ptr [ebp - 24]
	sub	eax, 96
	mov	cx, ax
	mov	word ptr [ebp - 24], cx
	movzx	eax, word ptr [ebp - 24]
	cmp	eax, 0
	jl	LBB39_4
## BB#3:
	mov	eax, 16
	mov	ecx, 8
	movzx	edx, word ptr [ebp - 24]
	mov	dword ptr [ebp - 28], eax ## 4-byte Spill
	mov	eax, edx
	cdq
	idiv	ecx
	mov	ecx, dword ptr [ebp - 16]
	movsx	edx, word ptr [ecx]
	sub	edx, eax
	mov	si, dx
	mov	word ptr [ecx], si
	movzx	eax, word ptr [ebp - 24]
	cdq
	mov	ecx, dword ptr [ebp - 28] ## 4-byte Reload
	idiv	ecx
	mov	edx, dword ptr [ebp - 16]
	movsx	edi, word ptr [edx + 4]
	sub	edi, eax
	mov	si, di
	mov	word ptr [edx + 4], si
LBB39_4:
	jmp	LBB39_6
LBB39_5:
	mov	eax, 16
	mov	ecx, 8
	movzx	edx, word ptr [ebp - 24]
	mov	dword ptr [ebp - 32], eax ## 4-byte Spill
	mov	eax, edx
	cdq
	idiv	ecx
	mov	ecx, dword ptr [ebp - 16]
	movsx	edx, word ptr [ecx]
	add	edx, eax
	mov	si, dx
	mov	word ptr [ecx], si
	movzx	eax, word ptr [ebp - 24]
	cdq
	mov	ecx, dword ptr [ebp - 32] ## 4-byte Reload
	idiv	ecx
	mov	edx, dword ptr [ebp - 16]
	movsx	edi, word ptr [edx + 4]
	add	edi, eax
	mov	si, di
	mov	word ptr [edx + 4], si
LBB39_6:
	jmp	LBB39_7
LBB39_7:
	add	esp, 24
	pop	esi
	pop	edi
	pop	ebp
	ret

	.align	4, 0x90
_ride_compute_upkeep:                   ## @ride_compute_upkeep
## BB#0:
	push	ebp
	mov	ebp, esp
	push	ebx
	push	edi
	push	esi
	sub	esp, 28
	call	L40$pb
L40$pb:
	pop	eax
	mov	ecx, dword ptr [ebp + 8]
	mov	dword ptr [ebp - 16], ecx
	movzx	ecx, byte ptr [ecx]
	mov	edx, dword ptr [eax + L_initialUpkeepCosts$non_lazy_ptr-L40$pb]
	movzx	ecx, byte ptr [edx + ecx]
	mov	si, cx
	mov	word ptr [ebp - 18], si
	mov	ecx, dword ptr [ebp - 16]
	movzx	ecx, byte ptr [ecx]
	mov	edx, dword ptr [eax + L_costPerTrackPiece$non_lazy_ptr-L40$pb]
	mov	bl, byte ptr [edx + ecx]
	movzx	ecx, bl
	mov	si, cx
	mov	word ptr [ebp - 20], si
	mov	ecx, dword ptr [ebp - 16]
	mov	bl, byte ptr [ecx + 277]
	mov	byte ptr [ebp - 21], bl
	movzx	ecx, byte ptr [ebp - 21]
	sar	ecx, 6
	mov	bl, cl
	mov	byte ptr [ebp - 21], bl
	movzx	ecx, byte ptr [ebp - 21]
	and	ecx, 3
	mov	bl, cl
	mov	byte ptr [ebp - 21], bl
	movzx	ecx, word ptr [ebp - 20]
	movzx	edx, byte ptr [ebp - 21]
	imul	ecx, edx
	movzx	edx, word ptr [ebp - 18]
	add	edx, ecx
	mov	si, dx
	mov	word ptr [ebp - 18], si
	mov	ecx, dword ptr [ebp - 16]
	mov	dword ptr [esp], ecx
	mov	dword ptr [ebp - 36], eax ## 4-byte Spill
	call	_ride_get_total_length
	mov	ecx, dword ptr [ebp - 36] ## 4-byte Reload
	mov	edx, dword ptr [ecx + L_hasRunningTrack$non_lazy_ptr-L40$pb]
	sar	eax, 16
	mov	dword ptr [ebp - 28], eax
	mov	eax, dword ptr [ebp - 16]
	movzx	eax, byte ptr [eax]
	test	byte ptr [edx + eax], 1
	je	LBB40_2
## BB#1:
	imul	eax, dword ptr [ebp - 28], 20
	mov	dword ptr [ebp - 28], eax
LBB40_2:
	mov	eax, dword ptr [ebp - 28]
	shr	eax, 10
	mov	cx, ax
	movzx	eax, cx
	movzx	edx, word ptr [ebp - 18]
	add	edx, eax
	mov	cx, dx
	mov	word ptr [ebp - 18], cx
	mov	eax, dword ptr [ebp - 16]
	mov	eax, dword ptr [eax + 464]
	and	eax, 32
	cmp	eax, 0
	je	LBB40_4
## BB#3:
	movzx	eax, word ptr [ebp - 18]
	add	eax, 40
	mov	cx, ax
	mov	word ptr [ebp - 18], cx
LBB40_4:
	mov	eax, dword ptr [ebp - 16]
	movzx	eax, byte ptr [eax]
	cmp	eax, 65
	jne	LBB40_6
## BB#5:
	mov	word ptr [ebp - 30], 10
	jmp	LBB40_7
LBB40_6:
	mov	word ptr [ebp - 30], 80
LBB40_7:
	mov	eax, dword ptr [ebp - 36] ## 4-byte Reload
	mov	ecx, dword ptr [eax + L_rideUnknownData2$non_lazy_ptr-L40$pb]
	mov	edx, dword ptr [eax + L_rideUnknownData1$non_lazy_ptr-L40$pb]
	mov	esi, 20493770
	mov	edi, 20493772
	mov	bx, word ptr [edi]
	mov	word ptr [ebp - 32], bx
	movzx	edi, word ptr [ebp - 30]
	movsx	eax, word ptr [ebp - 32]
	imul	edi, eax
	movzx	eax, word ptr [ebp - 18]
	add	eax, edi
	mov	bx, ax
	mov	word ptr [ebp - 18], bx
	mov	bx, word ptr [esi]
	mov	word ptr [ebp - 32], bx
	movsx	eax, word ptr [ebp - 32]
	imul	eax, eax, 20
	movzx	esi, word ptr [ebp - 18]
	add	esi, eax
	mov	bx, si
	mov	word ptr [ebp - 18], bx
	mov	eax, dword ptr [ebp - 16]
	movzx	eax, byte ptr [eax]
	movzx	eax, byte ptr [edx + eax]
	mov	edx, dword ptr [ebp - 16]
	movzx	edx, byte ptr [edx + 200]
	imul	eax, edx
	movzx	edx, word ptr [ebp - 18]
	add	edx, eax
	mov	bx, dx
	mov	word ptr [ebp - 18], bx
	mov	eax, dword ptr [ebp - 16]
	movzx	eax, byte ptr [eax]
	test	byte ptr [ecx + eax], 1
	je	LBB40_9
## BB#8:
	mov	eax, dword ptr [ebp - 16]
	movzx	eax, byte ptr [eax + 201]
	imul	eax, eax, 3
	movzx	ecx, word ptr [ebp - 18]
	add	ecx, eax
	mov	dx, cx
	mov	word ptr [ebp - 18], dx
LBB40_9:
	mov	eax, dword ptr [ebp - 36] ## 4-byte Reload
	mov	ecx, dword ptr [eax + L_rideUnknownData3$non_lazy_ptr-L40$pb]
	mov	edx, dword ptr [ebp - 16]
	movzx	edx, byte ptr [edx]
	movzx	ecx, byte ptr [ecx + edx]
	mov	edx, dword ptr [ebp - 16]
	movzx	edx, byte ptr [edx + 199]
	imul	ecx, edx
	movzx	edx, word ptr [ebp - 18]
	add	edx, ecx
	mov	si, dx
	mov	word ptr [ebp - 18], si
	mov	ecx, dword ptr [ebp - 16]
	movzx	ecx, byte ptr [ecx + 4]
	cmp	ecx, 2
	jne	LBB40_11
## BB#10:
	movzx	eax, word ptr [ebp - 18]
	add	eax, 30
	mov	cx, ax
	mov	word ptr [ebp - 18], cx
	jmp	LBB40_21
LBB40_11:
	mov	eax, dword ptr [ebp - 16]
	movzx	eax, byte ptr [eax + 4]
	cmp	eax, 3
	jne	LBB40_13
## BB#12:
	movzx	eax, word ptr [ebp - 18]
	add	eax, 160
	mov	cx, ax
	mov	word ptr [ebp - 18], cx
	jmp	LBB40_20
LBB40_13:
	mov	eax, dword ptr [ebp - 16]
	movzx	eax, byte ptr [eax + 4]
	cmp	eax, 23
	jne	LBB40_15
## BB#14:
	movzx	eax, word ptr [ebp - 18]
	add	eax, 320
	mov	cx, ax
	mov	word ptr [ebp - 18], cx
	jmp	LBB40_19
LBB40_15:
	mov	eax, dword ptr [ebp - 16]
	movzx	eax, byte ptr [eax + 4]
	cmp	eax, 35
	je	LBB40_17
## BB#16:
	mov	eax, dword ptr [ebp - 16]
	movzx	eax, byte ptr [eax + 4]
	cmp	eax, 36
	jne	LBB40_18
LBB40_17:
	movzx	eax, word ptr [ebp - 18]
	add	eax, 220
	mov	cx, ax
	mov	word ptr [ebp - 18], cx
LBB40_18:
	jmp	LBB40_19
LBB40_19:
	jmp	LBB40_20
LBB40_20:
	jmp	LBB40_21
LBB40_21:
	movzx	eax, word ptr [ebp - 18]
	imul	eax, eax, 10
	mov	cx, ax
	mov	word ptr [ebp - 18], cx
	movzx	eax, word ptr [ebp - 18]
	sar	eax, 4
	mov	cx, ax
	mov	word ptr [ebp - 18], cx
	movzx	eax, word ptr [ebp - 18]
	add	esp, 28
	pop	esi
	pop	edi
	pop	ebx
	pop	ebp
	ret

	.align	4, 0x90
_sub_65DDD1:                            ## @sub_65DDD1
## BB#0:
	push	ebp
	mov	ebp, esp
	push	edi
	push	esi
	sub	esp, 80
	mov	eax, dword ptr [ebp + 12]
	mov	ecx, dword ptr [ebp + 8]
	mov	dword ptr [ebp - 12], eax
	mov	dword ptr [ebp - 16], 0
	mov	dword ptr [ebp - 20], 0
	mov	dword ptr [ebp - 24], 0
	mov	eax, dword ptr [ebp - 12]
	movzx	edx, byte ptr [eax]
	mov	esi, esp
	mov	dword ptr [esi + 8], eax
	mov	dword ptr [esi + 4], edx
	lea	eax, dword ptr [ebp - 32]
	mov	dword ptr [esi], eax
	mov	dword ptr [ebp - 76], ecx ## 4-byte Spill
	call	_get_special_track_elements_rating
	sub	esp, 4
	movsx	eax, word ptr [ebp - 32]
	mov	ecx, dword ptr [ebp - 16]
	add	ecx, eax
	mov	dword ptr [ebp - 16], ecx
	movsx	eax, word ptr [ebp - 30]
	mov	ecx, dword ptr [ebp - 20]
	add	ecx, eax
	mov	dword ptr [ebp - 20], ecx
	movsx	eax, word ptr [ebp - 28]
	mov	ecx, dword ptr [ebp - 24]
	add	ecx, eax
	mov	dword ptr [ebp - 24], ecx
	mov	eax, dword ptr [ebp - 12]
	mov	ecx, esp
	mov	dword ptr [ecx + 4], eax
	lea	eax, dword ptr [ebp - 40]
	mov	dword ptr [ecx], eax
	call	_get_var_10E_rating
	sub	esp, 4
	movsx	eax, word ptr [ebp - 40]
	mov	ecx, dword ptr [ebp - 16]
	add	ecx, eax
	mov	dword ptr [ebp - 16], ecx
	movsx	eax, word ptr [ebp - 38]
	mov	ecx, dword ptr [ebp - 20]
	add	ecx, eax
	mov	dword ptr [ebp - 20], ecx
	movsx	eax, word ptr [ebp - 36]
	mov	ecx, dword ptr [ebp - 24]
	add	ecx, eax
	mov	dword ptr [ebp - 24], ecx
	mov	eax, dword ptr [ebp - 12]
	mov	ecx, esp
	mov	dword ptr [ecx + 4], eax
	lea	eax, dword ptr [ebp - 48]
	mov	dword ptr [ecx], eax
	call	_get_var_110_rating
	sub	esp, 4
	movsx	eax, word ptr [ebp - 48]
	mov	ecx, dword ptr [ebp - 16]
	add	ecx, eax
	mov	dword ptr [ebp - 16], ecx
	movsx	eax, word ptr [ebp - 46]
	mov	ecx, dword ptr [ebp - 20]
	add	ecx, eax
	mov	dword ptr [ebp - 20], ecx
	movsx	eax, word ptr [ebp - 44]
	mov	ecx, dword ptr [ebp - 24]
	add	ecx, eax
	mov	dword ptr [ebp - 24], ecx
	mov	eax, dword ptr [ebp - 12]
	movzx	eax, word ptr [eax + 274]
	mov	ecx, esp
	mov	dword ptr [ecx + 4], eax
	lea	eax, dword ptr [ebp - 56]
	mov	dword ptr [ecx], eax
	call	_get_var_112_rating
	sub	esp, 4
	movsx	eax, word ptr [ebp - 56]
	mov	ecx, dword ptr [ebp - 16]
	add	ecx, eax
	mov	dword ptr [ebp - 16], ecx
	movsx	eax, word ptr [ebp - 54]
	mov	ecx, dword ptr [ebp - 20]
	add	ecx, eax
	mov	dword ptr [ebp - 20], ecx
	movsx	eax, word ptr [ebp - 52]
	mov	ecx, dword ptr [ebp - 24]
	add	ecx, eax
	mov	dword ptr [ebp - 24], ecx
	mov	eax, dword ptr [ebp - 12]
	movzx	eax, byte ptr [eax + 276]
	mov	ecx, esp
	mov	dword ptr [ecx + 4], eax
	lea	eax, dword ptr [ebp - 64]
	mov	dword ptr [ecx], eax
	call	_get_inversions_ratings
	sub	esp, 4
	movsx	eax, word ptr [ebp - 64]
	mov	ecx, dword ptr [ebp - 16]
	add	ecx, eax
	mov	dword ptr [ebp - 16], ecx
	movsx	eax, word ptr [ebp - 62]
	mov	ecx, dword ptr [ebp - 20]
	add	ecx, eax
	mov	dword ptr [ebp - 20], ecx
	movsx	eax, word ptr [ebp - 60]
	mov	ecx, dword ptr [ebp - 24]
	add	ecx, eax
	mov	dword ptr [ebp - 24], ecx
	mov	di, word ptr [ebp - 16]
	mov	word ptr [ebp - 72], di
	mov	di, word ptr [ebp - 20]
	mov	word ptr [ebp - 70], di
	mov	di, word ptr [ebp - 24]
	mov	word ptr [ebp - 68], di
	mov	eax, dword ptr [ebp - 76] ## 4-byte Reload
	mov	word ptr [eax + 4], di
	mov	ecx, dword ptr [ebp - 72]
	mov	dword ptr [eax], ecx
	add	esp, 80
	pop	esi
	pop	edi
	pop	ebp
	ret	4

	.align	4, 0x90
_sub_65E1C2:                            ## @sub_65E1C2
## BB#0:
	push	ebp
	mov	ebp, esp
	push	esi
	sub	esp, 84
	mov	eax, dword ptr [ebp + 12]
	mov	ecx, dword ptr [ebp + 8]
	mov	dword ptr [ebp - 8], eax
	mov	dword ptr [ebp - 32], eax
	mov	eax, esp
	lea	edx, dword ptr [ebp - 36]
	mov	dword ptr [eax + 28], edx
	lea	edx, dword ptr [ebp - 32]
	mov	dword ptr [eax + 24], edx
	lea	edx, dword ptr [ebp - 28]
	mov	dword ptr [eax + 20], edx
	lea	edx, dword ptr [ebp - 24]
	mov	dword ptr [eax + 16], edx
	lea	edx, dword ptr [ebp - 20]
	mov	dword ptr [eax + 12], edx
	lea	edx, dword ptr [ebp - 16]
	mov	dword ptr [eax + 8], edx
	lea	edx, dword ptr [ebp - 12]
	mov	dword ptr [eax + 4], edx
	mov	dword ptr [eax], 6676930
	mov	dword ptr [ebp - 52], ecx ## 4-byte Spill
	call	_RCT2_CALLFUNC_X
	mov	si, word ptr [ebp - 16]
	mov	word ptr [ebp - 48], si
	mov	si, word ptr [ebp - 20]
	mov	word ptr [ebp - 46], si
	mov	si, word ptr [ebp - 36]
	mov	word ptr [ebp - 44], si
	mov	ecx, dword ptr [ebp - 52] ## 4-byte Reload
	mov	word ptr [ecx + 4], si
	mov	edx, dword ptr [ebp - 48]
	mov	dword ptr [ecx], edx
	mov	dword ptr [ebp - 56], eax ## 4-byte Spill
	add	esp, 84
	pop	esi
	pop	ebp
	ret	4

	.align	4, 0x90
_sub_65E72D:                            ## @sub_65E72D
## BB#0:
	push	ebp
	mov	ebp, esp
	push	ebx
	push	edi
	push	esi
	sub	esp, 92
	mov	eax, dword ptr [ebp + 8]
	mov	ecx, 6677111
	lea	edx, dword ptr [ebp - 20]
	lea	esi, dword ptr [ebp - 24]
	lea	edi, dword ptr [ebp - 28]
	lea	ebx, dword ptr [ebp - 32]
	mov	dword ptr [ebp - 48], eax ## 4-byte Spill
	lea	eax, dword ptr [ebp - 36]
	mov	dword ptr [ebp - 52], eax ## 4-byte Spill
	lea	eax, dword ptr [ebp - 40]
	mov	dword ptr [ebp - 56], eax ## 4-byte Spill
	lea	eax, dword ptr [ebp - 44]
	mov	dword ptr [ebp - 60], eax ## 4-byte Spill
	mov	eax, dword ptr [ebp - 48] ## 4-byte Reload
	mov	dword ptr [ebp - 16], eax
	mov	eax, dword ptr [ebp - 16]
	mov	dword ptr [ebp - 40], eax
	mov	dword ptr [esp], 6677111
	mov	dword ptr [esp + 4], edx
	mov	dword ptr [esp + 8], esi
	mov	dword ptr [esp + 12], edi
	mov	dword ptr [esp + 16], ebx
	mov	eax, dword ptr [ebp - 52] ## 4-byte Reload
	mov	dword ptr [esp + 20], eax
	mov	edx, dword ptr [ebp - 56] ## 4-byte Reload
	mov	dword ptr [esp + 24], edx
	mov	esi, dword ptr [ebp - 60] ## 4-byte Reload
	mov	dword ptr [esp + 28], esi
	mov	dword ptr [ebp - 64], ecx ## 4-byte Spill
	call	_RCT2_CALLFUNC_X
	mov	ecx, dword ptr [ebp - 32]
	and	ecx, 65535
	mov	dword ptr [ebp - 68], eax ## 4-byte Spill
	mov	eax, ecx
	add	esp, 92
	pop	esi
	pop	edi
	pop	ebx
	pop	ebp
	ret

	.align	4, 0x90
_RCT2_CALLFUNC_X:                       ## @RCT2_CALLFUNC_X
## BB#0:
	push	ebp
	mov	ebp, esp
	push	ebx
	push	edi
	push	esi
	sub	esp, 44
	mov	eax, dword ptr [ebp + 36]
	mov	ecx, dword ptr [ebp + 32]
	mov	edx, dword ptr [ebp + 28]
	mov	esi, dword ptr [ebp + 24]
	mov	edi, dword ptr [ebp + 20]
	mov	ebx, dword ptr [ebp + 16]
	mov	dword ptr [ebp - 52], eax ## 4-byte Spill
	mov	eax, dword ptr [ebp + 12]
	mov	dword ptr [ebp - 56], eax ## 4-byte Spill
	mov	eax, dword ptr [ebp + 8]
	mov	dword ptr [ebp - 20], eax
	mov	eax, dword ptr [ebp - 56] ## 4-byte Reload
	mov	dword ptr [ebp - 24], eax
	mov	dword ptr [ebp - 28], ebx
	mov	dword ptr [ebp - 32], edi
	mov	dword ptr [ebp - 36], esi
	mov	dword ptr [ebp - 40], edx
	mov	dword ptr [ebp - 44], ecx
	mov	ecx, dword ptr [ebp - 52] ## 4-byte Reload
	mov	dword ptr [ebp - 48], ecx
	## InlineAsm Start
		       
                /* Store C's base pointer*/     
                push ebp        
                push ebx        
        
                /* Store -20(%ebp) to call*/   
                push -20(%ebp)         
        
                /* Set all registers to the input values*/      
                mov eax, [-24(%ebp)]      
                mov eax, [eax]  
                mov ebx, [-28(%ebp)]      
                mov ebx, [ebx]  
                mov ecx, [-32(%ebp)]      
                mov ecx, [ecx]  
                mov edx, [-36(%ebp)]      
                mov edx, [edx]  
                mov esi, [-40(%ebp)]      
                mov esi, [esi]  
                mov edi, [-44(%ebp)]      
                mov edi, [edi]  
                mov ebp, [-48(%ebp)]      
                mov ebp, [ebp]  
        
                /* Call function*/      
                call [esp]      
        
				/* Store output eax */ 
				push eax 
				push ebp 
				push ebx 
				mov ebp, [esp + 20] 
				mov ebx, [esp + 16] 
                /* Get resulting ecx, edx, esi, edi registers*/       
                mov eax, [-44(%ebp)]      
                mov [eax], edi  
                mov eax, [-40(%ebp)]      
                mov [eax], esi  
                mov eax, [-36(%ebp)]      
                mov [eax], edx  
                mov eax, [-32(%ebp)]      
                mov [eax], ecx  
				/* Pop ebx reg into ecx*/ 
				pop ecx		
				mov eax, [-28(%ebp)] 
				mov [eax], ecx 
				
				/* Pop ebp reg into ecx */
				pop ecx 
				mov eax, [-48(%ebp)] 
				mov [eax], ecx 
				
				pop eax 
				/* Get resulting eax register*/ 
				mov ecx, [-24(%ebp)] 
				mov [ecx], eax 
				
				/* Save flags as return in eax*/  
				lahf 
				/* Pop address*/ 
				pop ebp 
				
				pop ebx 
				pop ebp 
	 
	## InlineAsm End
	mov	eax, dword ptr [ebp - 16]
	add	esp, 44
	pop	esi
	pop	edi
	pop	ebx
	pop	ebp
	ret

	.align	4, 0x90
_sub_65E277:                            ## @sub_65E277
## BB#0:
	push	ebp
	mov	ebp, esp
	push	ebx
	push	edi
	push	esi
	sub	esp, 76
	mov	eax, 6677111
	lea	ecx, dword ptr [ebp - 16]
	lea	edx, dword ptr [ebp - 20]
	lea	esi, dword ptr [ebp - 24]
	lea	edi, dword ptr [ebp - 28]
	lea	ebx, dword ptr [ebp - 32]
	mov	dword ptr [ebp - 44], eax ## 4-byte Spill
	lea	eax, dword ptr [ebp - 36]
	mov	dword ptr [ebp - 48], eax ## 4-byte Spill
	lea	eax, dword ptr [ebp - 40]
	mov	dword ptr [esp], 6677111
	mov	dword ptr [esp + 4], ecx
	mov	dword ptr [esp + 8], edx
	mov	dword ptr [esp + 12], esi
	mov	dword ptr [esp + 16], edi
	mov	dword ptr [esp + 20], ebx
	mov	ecx, dword ptr [ebp - 48] ## 4-byte Reload
	mov	dword ptr [esp + 24], ecx
	mov	dword ptr [esp + 28], eax
	call	_RCT2_CALLFUNC_X
	mov	ecx, dword ptr [ebp - 20]
	mov	dword ptr [ebp - 52], eax ## 4-byte Spill
	mov	eax, ecx
	add	esp, 76
	pop	esi
	pop	edi
	pop	ebx
	pop	ebp
	ret

	.align	4, 0x90
_get_special_track_elements_rating:     ## @get_special_track_elements_rating
## BB#0:
	push	ebp
	mov	ebp, esp
	sub	esp, 72
	mov	eax, dword ptr [ebp + 16]
	mov	cl, byte ptr [ebp + 12]
	mov	edx, dword ptr [ebp + 8]
	mov	byte ptr [ebp - 1], cl
	mov	dword ptr [ebp - 8], eax
	mov	dword ptr [ebp - 12], 0
	mov	dword ptr [ebp - 16], 0
	mov	dword ptr [ebp - 20], 0
	movzx	eax, byte ptr [ebp - 1]
	cmp	eax, 50
	mov	dword ptr [ebp - 44], edx ## 4-byte Spill
	jne	LBB46_4
## BB#1:
	mov	eax, dword ptr [ebp - 8]
	mov	dword ptr [esp], eax
	call	_ride_has_spinning_tunnel
	test	al, 1
	jne	LBB46_2
	jmp	LBB46_3
LBB46_2:
	mov	eax, dword ptr [ebp - 12]
	add	eax, 40
	mov	dword ptr [ebp - 12], eax
	mov	eax, dword ptr [ebp - 16]
	add	eax, 25
	mov	dword ptr [ebp - 16], eax
	mov	eax, dword ptr [ebp - 20]
	add	eax, 55
	mov	dword ptr [ebp - 20], eax
LBB46_3:
	jmp	LBB46_16
LBB46_4:
	movzx	eax, byte ptr [ebp - 1]
	cmp	eax, 23
	jne	LBB46_8
## BB#5:
	mov	eax, dword ptr [ebp - 8]
	mov	dword ptr [esp], eax
	call	_ride_has_log_reverser
	test	al, 1
	jne	LBB46_6
	jmp	LBB46_7
LBB46_6:
	mov	eax, dword ptr [ebp - 12]
	add	eax, 48
	mov	dword ptr [ebp - 12], eax
	mov	eax, dword ptr [ebp - 16]
	add	eax, 55
	mov	dword ptr [ebp - 16], eax
	mov	eax, dword ptr [ebp - 20]
	add	eax, 65
	mov	dword ptr [ebp - 20], eax
LBB46_7:
	jmp	LBB46_15
LBB46_8:
	mov	eax, dword ptr [ebp - 8]
	mov	dword ptr [esp], eax
	call	_ride_has_water_splash
	test	al, 1
	jne	LBB46_9
	jmp	LBB46_10
LBB46_9:
	mov	eax, dword ptr [ebp - 12]
	add	eax, 50
	mov	dword ptr [ebp - 12], eax
	mov	eax, dword ptr [ebp - 16]
	add	eax, 30
	mov	dword ptr [ebp - 16], eax
	mov	eax, dword ptr [ebp - 20]
	add	eax, 20
	mov	dword ptr [ebp - 20], eax
LBB46_10:
	mov	eax, dword ptr [ebp - 8]
	mov	dword ptr [esp], eax
	call	_ride_has_waterfall
	test	al, 1
	jne	LBB46_11
	jmp	LBB46_12
LBB46_11:
	mov	eax, dword ptr [ebp - 12]
	add	eax, 55
	mov	dword ptr [ebp - 12], eax
	mov	eax, dword ptr [ebp - 16]
	add	eax, 30
	mov	dword ptr [ebp - 16], eax
LBB46_12:
	mov	eax, dword ptr [ebp - 8]
	mov	dword ptr [esp], eax
	call	_ride_has_whirlpool
	test	al, 1
	jne	LBB46_13
	jmp	LBB46_14
LBB46_13:
	mov	eax, dword ptr [ebp - 12]
	add	eax, 35
	mov	dword ptr [ebp - 12], eax
	mov	eax, dword ptr [ebp - 16]
	add	eax, 20
	mov	dword ptr [ebp - 16], eax
	mov	eax, dword ptr [ebp - 20]
	add	eax, 23
	mov	dword ptr [ebp - 20], eax
LBB46_14:
	jmp	LBB46_15
LBB46_15:
	jmp	LBB46_16
LBB46_16:
	mov	eax, dword ptr [ebp - 8]
	mov	dword ptr [esp], eax
	call	_ride_get_helix_sections
	mov	byte ptr [ebp - 21], al
	movzx	ecx, byte ptr [ebp - 21]
	cmp	ecx, 9
	jge	LBB46_18
## BB#17:
	movzx	eax, byte ptr [ebp - 21]
	mov	dword ptr [ebp - 48], eax ## 4-byte Spill
	jmp	LBB46_19
LBB46_18:
	mov	eax, 9
	mov	dword ptr [ebp - 48], eax ## 4-byte Spill
	jmp	LBB46_19
LBB46_19:
	mov	eax, dword ptr [ebp - 48] ## 4-byte Reload
	mov	dword ptr [ebp - 28], eax
	imul	eax, dword ptr [ebp - 28], 254862
	sar	eax, 16
	mov	ecx, dword ptr [ebp - 12]
	add	ecx, eax
	mov	dword ptr [ebp - 12], ecx
	movzx	eax, byte ptr [ebp - 21]
	cmp	eax, 11
	jge	LBB46_21
## BB#20:
	movzx	eax, byte ptr [ebp - 21]
	mov	dword ptr [ebp - 52], eax ## 4-byte Spill
	jmp	LBB46_22
LBB46_21:
	mov	eax, 11
	mov	dword ptr [ebp - 52], eax ## 4-byte Spill
	jmp	LBB46_22
LBB46_22:
	mov	eax, dword ptr [ebp - 52] ## 4-byte Reload
	mov	dword ptr [ebp - 28], eax
	imul	eax, dword ptr [ebp - 28], 148945
	sar	eax, 16
	mov	ecx, dword ptr [ebp - 16]
	add	ecx, eax
	mov	dword ptr [ebp - 16], ecx
	movzx	eax, byte ptr [ebp - 21]
	sub	eax, 5
	cmp	eax, 0
	jle	LBB46_24
## BB#23:
	movzx	eax, byte ptr [ebp - 21]
	sub	eax, 5
	mov	dword ptr [ebp - 56], eax ## 4-byte Spill
	jmp	LBB46_25
LBB46_24:
	mov	eax, 0
	mov	dword ptr [ebp - 56], eax ## 4-byte Spill
	jmp	LBB46_25
LBB46_25:
	mov	eax, dword ptr [ebp - 56] ## 4-byte Reload
	mov	dword ptr [ebp - 28], eax
	cmp	dword ptr [ebp - 28], 10
	jge	LBB46_27
## BB#26:
	mov	eax, dword ptr [ebp - 28]
	mov	dword ptr [ebp - 60], eax ## 4-byte Spill
	jmp	LBB46_28
LBB46_27:
	mov	eax, 10
	mov	dword ptr [ebp - 60], eax ## 4-byte Spill
	jmp	LBB46_28
LBB46_28:
	mov	eax, dword ptr [ebp - 60] ## 4-byte Reload
	mov	dword ptr [ebp - 28], eax
	imul	eax, eax, 1310720
	sar	eax, 16
	mov	ecx, dword ptr [ebp - 20]
	add	ecx, eax
	mov	dword ptr [ebp - 20], ecx
	mov	dx, word ptr [ebp - 12]
	mov	word ptr [ebp - 40], dx
	mov	dx, word ptr [ebp - 16]
	mov	word ptr [ebp - 38], dx
	mov	dx, word ptr [ebp - 20]
	mov	word ptr [ebp - 36], dx
	mov	eax, dword ptr [ebp - 44] ## 4-byte Reload
	mov	word ptr [eax + 4], dx
	mov	ecx, dword ptr [ebp - 40]
	mov	dword ptr [eax], ecx
	add	esp, 72
	pop	ebp
	ret	4

	.align	4, 0x90
_get_var_10E_rating:                    ## @get_var_10E_rating
## BB#0:
	push	ebp
	mov	ebp, esp
	push	esi
	sub	esp, 52
	mov	eax, dword ptr [ebp + 12]
	mov	ecx, dword ptr [ebp + 8]
	mov	dword ptr [ebp - 8], eax
	mov	edx, esp
	mov	dword ptr [edx], eax
	mov	dword ptr [ebp - 44], ecx ## 4-byte Spill
	call	_get_var_10E_unk_1
	mov	dword ptr [ebp - 12], eax
	mov	eax, dword ptr [ebp - 8]
	mov	ecx, esp
	mov	dword ptr [ecx], eax
	call	_get_var_10E_unk_2
	mov	dword ptr [ebp - 16], eax
	mov	eax, dword ptr [ebp - 8]
	mov	ecx, esp
	mov	dword ptr [ecx], eax
	call	_get_var_10E_unk_3
	mov	dword ptr [ebp - 20], eax
	mov	eax, dword ptr [ebp - 12]
	imul	eax, eax, 163840
	sar	eax, 16
	mov	dword ptr [ebp - 24], eax
	mov	ecx, dword ptr [ebp - 16]
	lea	ecx, dword ptr [ecx + 2*ecx]
	add	eax, ecx
	mov	dword ptr [ebp - 24], eax
	mov	ecx, dword ptr [ebp - 20]
	imul	ecx, ecx, 63421
	sar	ecx, 16
	add	eax, ecx
	mov	dword ptr [ebp - 24], eax
	mov	eax, dword ptr [ebp - 12]
	imul	eax, eax, 81920
	sar	eax, 16
	mov	dword ptr [ebp - 28], eax
	mov	ecx, dword ptr [ebp - 16]
	imul	ecx, ecx, 49152
	sar	ecx, 16
	add	eax, ecx
	mov	dword ptr [ebp - 28], eax
	mov	ecx, dword ptr [ebp - 20]
	imul	ecx, ecx, 21140
	sar	ecx, 16
	add	eax, ecx
	mov	dword ptr [ebp - 28], eax
	mov	eax, dword ptr [ebp - 12]
	lea	eax, dword ptr [eax + 4*eax]
	mov	dword ptr [ebp - 32], eax
	mov	ecx, dword ptr [ebp - 16]
	imul	ecx, ecx, 12800
	sar	ecx, 16
	add	eax, ecx
	mov	dword ptr [ebp - 32], eax
	mov	ecx, dword ptr [ebp - 20]
	imul	ecx, ecx, 42281
	sar	ecx, 16
	add	eax, ecx
	mov	dword ptr [ebp - 32], eax
	mov	si, word ptr [ebp - 24]
	mov	word ptr [ebp - 40], si
	mov	si, word ptr [ebp - 28]
	mov	word ptr [ebp - 38], si
	mov	si, word ptr [ebp - 32]
	mov	word ptr [ebp - 36], si
	mov	eax, dword ptr [ebp - 44] ## 4-byte Reload
	mov	word ptr [eax + 4], si
	mov	ecx, dword ptr [ebp - 40]
	mov	dword ptr [eax], ecx
	add	esp, 52
	pop	esi
	pop	ebp
	ret	4

	.align	4, 0x90
_get_var_110_rating:                    ## @get_var_110_rating
## BB#0:
	push	ebp
	mov	ebp, esp
	push	esi
	sub	esp, 52
	mov	eax, dword ptr [ebp + 12]
	mov	ecx, dword ptr [ebp + 8]
	mov	dword ptr [ebp - 8], eax
	mov	edx, esp
	mov	dword ptr [edx], eax
	mov	dword ptr [ebp - 44], ecx ## 4-byte Spill
	call	_get_var_10E_unk_1
	mov	dword ptr [ebp - 12], eax
	mov	eax, dword ptr [ebp - 8]
	mov	ecx, esp
	mov	dword ptr [ecx], eax
	call	_get_var_10E_unk_2
	mov	dword ptr [ebp - 16], eax
	mov	eax, dword ptr [ebp - 8]
	mov	ecx, esp
	mov	dword ptr [ecx], eax
	call	_get_var_10E_unk_3
	mov	dword ptr [ebp - 20], eax
	mov	eax, dword ptr [ebp - 12]
	imul	eax, eax, 245760
	sar	eax, 16
	mov	dword ptr [ebp - 24], eax
	mov	ecx, dword ptr [ebp - 16]
	imul	ecx, ecx, 245760
	sar	ecx, 16
	add	eax, ecx
	mov	dword ptr [ebp - 24], eax
	mov	ecx, dword ptr [ebp - 20]
	imul	ecx, ecx, 73992
	sar	ecx, 16
	add	eax, ecx
	mov	dword ptr [ebp - 24], eax
	mov	eax, dword ptr [ebp - 12]
	imul	eax, eax, 81920
	sar	eax, 16
	mov	dword ptr [ebp - 28], eax
	mov	ecx, dword ptr [ebp - 16]
	imul	ecx, ecx, 49152
	sar	ecx, 16
	add	eax, ecx
	mov	dword ptr [ebp - 28], eax
	mov	ecx, dword ptr [ebp - 20]
	imul	ecx, ecx, 21140
	sar	ecx, 16
	add	eax, ecx
	mov	dword ptr [ebp - 28], eax
	mov	eax, dword ptr [ebp - 12]
	lea	eax, dword ptr [eax + 4*eax]
	mov	dword ptr [ebp - 32], eax
	mov	ecx, dword ptr [ebp - 16]
	imul	ecx, ecx, 204800
	sar	ecx, 16
	add	eax, ecx
	mov	dword ptr [ebp - 32], eax
	mov	ecx, dword ptr [ebp - 20]
	imul	ecx, ecx, 48623
	sar	ecx, 16
	add	eax, ecx
	mov	dword ptr [ebp - 32], eax
	mov	si, word ptr [ebp - 24]
	mov	word ptr [ebp - 40], si
	mov	si, word ptr [ebp - 28]
	mov	word ptr [ebp - 38], si
	mov	si, word ptr [ebp - 32]
	mov	word ptr [ebp - 36], si
	mov	eax, dword ptr [ebp - 44] ## 4-byte Reload
	mov	word ptr [eax + 4], si
	mov	ecx, dword ptr [ebp - 40]
	mov	dword ptr [eax], ecx
	add	esp, 52
	pop	esi
	pop	ebp
	ret	4

	.align	4, 0x90
_get_var_112_rating:                    ## @get_var_112_rating
## BB#0:
	push	ebp
	mov	ebp, esp
	sub	esp, 48
	movzx	eax, word ptr [ebp + 12]
	mov	cx, ax
	mov	eax, dword ptr [ebp + 8]
	mov	word ptr [ebp - 2], cx
	movzx	edx, word ptr [ebp - 2]
	sar	edx, 11
	mov	dword ptr [ebp - 8], edx
	mov	edx, dword ptr [ebp - 8]
	and	edx, 63
	cmp	edx, 4
	mov	dword ptr [ebp - 28], eax ## 4-byte Spill
	jge	LBB49_2
## BB#1:
	mov	eax, dword ptr [ebp - 8]
	and	eax, 63
	mov	dword ptr [ebp - 32], eax ## 4-byte Spill
	jmp	LBB49_3
LBB49_2:
	mov	eax, 4
	mov	dword ptr [ebp - 32], eax ## 4-byte Spill
	jmp	LBB49_3
LBB49_3:
	mov	eax, dword ptr [ebp - 32] ## 4-byte Reload
	mov	dword ptr [ebp - 8], eax
	imul	eax, dword ptr [ebp - 8], 491520
	sar	eax, 16
	mov	dword ptr [ebp - 12], eax
	movzx	eax, word ptr [ebp - 2]
	sar	eax, 11
	mov	dword ptr [ebp - 8], eax
	mov	eax, dword ptr [ebp - 8]
	and	eax, 63
	cmp	eax, 8
	jge	LBB49_5
## BB#4:
	mov	eax, dword ptr [ebp - 8]
	and	eax, 63
	mov	dword ptr [ebp - 36], eax ## 4-byte Spill
	jmp	LBB49_6
LBB49_5:
	mov	eax, 8
	mov	dword ptr [ebp - 36], eax ## 4-byte Spill
	jmp	LBB49_6
LBB49_6:
	mov	eax, dword ptr [ebp - 36] ## 4-byte Reload
	mov	dword ptr [ebp - 8], eax
	imul	eax, dword ptr [ebp - 8], 491520
	sar	eax, 16
	mov	dword ptr [ebp - 16], eax
	movzx	eax, word ptr [ebp - 2]
	sar	eax, 8
	mov	dword ptr [ebp - 8], eax
	mov	eax, dword ptr [ebp - 8]
	and	eax, 7
	cmp	eax, 6
	jge	LBB49_8
## BB#7:
	mov	eax, dword ptr [ebp - 8]
	and	eax, 7
	mov	dword ptr [ebp - 40], eax ## 4-byte Spill
	jmp	LBB49_9
LBB49_8:
	mov	eax, 6
	mov	dword ptr [ebp - 40], eax ## 4-byte Spill
	jmp	LBB49_9
LBB49_9:
	mov	eax, dword ptr [ebp - 40] ## 4-byte Reload
	mov	dword ptr [ebp - 8], eax
	imul	eax, dword ptr [ebp - 8], 273066
	sar	eax, 16
	mov	ecx, dword ptr [ebp - 12]
	add	ecx, eax
	mov	dword ptr [ebp - 12], ecx
	movzx	eax, word ptr [ebp - 2]
	sar	eax, 5
	mov	dword ptr [ebp - 8], eax
	mov	eax, dword ptr [ebp - 8]
	and	eax, 7
	cmp	eax, 6
	jge	LBB49_11
## BB#10:
	mov	eax, dword ptr [ebp - 8]
	and	eax, 7
	mov	dword ptr [ebp - 44], eax ## 4-byte Spill
	jmp	LBB49_12
LBB49_11:
	mov	eax, 6
	mov	dword ptr [ebp - 44], eax ## 4-byte Spill
	jmp	LBB49_12
LBB49_12:
	mov	eax, dword ptr [ebp - 44] ## 4-byte Reload
	mov	dword ptr [ebp - 8], eax
	imul	eax, dword ptr [ebp - 8], 240298
	sar	eax, 16
	mov	ecx, dword ptr [ebp - 12]
	add	ecx, eax
	mov	dword ptr [ebp - 12], ecx
	movzx	eax, word ptr [ebp - 2]
	and	eax, 31
	cmp	eax, 7
	jge	LBB49_14
## BB#13:
	movzx	eax, word ptr [ebp - 2]
	and	eax, 31
	mov	dword ptr [ebp - 48], eax ## 4-byte Spill
	jmp	LBB49_15
LBB49_14:
	mov	eax, 7
	mov	dword ptr [ebp - 48], eax ## 4-byte Spill
	jmp	LBB49_15
LBB49_15:
	mov	eax, dword ptr [ebp - 48] ## 4-byte Reload
	mov	dword ptr [ebp - 8], eax
	imul	eax, eax, 187245
	sar	eax, 16
	mov	ecx, dword ptr [ebp - 12]
	add	ecx, eax
	mov	dword ptr [ebp - 12], ecx
	mov	dx, word ptr [ebp - 12]
	mov	word ptr [ebp - 24], dx
	mov	word ptr [ebp - 22], 0
	mov	dx, word ptr [ebp - 16]
	mov	word ptr [ebp - 20], dx
	mov	eax, dword ptr [ebp - 28] ## 4-byte Reload
	mov	word ptr [eax + 4], dx
	mov	ecx, dword ptr [ebp - 24]
	mov	dword ptr [eax], ecx
	add	esp, 48
	pop	ebp
	ret	4

	.align	4, 0x90
_get_inversions_ratings:                ## @get_inversions_ratings
## BB#0:
	push	ebp
	mov	ebp, esp
	sub	esp, 40
	mov	al, byte ptr [ebp + 12]
	mov	ecx, dword ptr [ebp + 8]
	mov	byte ptr [ebp - 1], al
	movzx	edx, byte ptr [ebp - 1]
	and	edx, 31
	mov	al, dl
	mov	byte ptr [ebp - 1], al
	movzx	edx, byte ptr [ebp - 1]
	cmp	edx, 6
	mov	dword ptr [ebp - 36], ecx ## 4-byte Spill
	jge	LBB50_2
## BB#1:
	movzx	eax, byte ptr [ebp - 1]
	mov	dword ptr [ebp - 40], eax ## 4-byte Spill
	jmp	LBB50_3
LBB50_2:
	mov	eax, 6
	mov	dword ptr [ebp - 40], eax ## 4-byte Spill
	jmp	LBB50_3
LBB50_3:
	mov	eax, dword ptr [ebp - 40] ## 4-byte Reload
	mov	dword ptr [ebp - 8], eax
	imul	eax, eax, 1747626
	sar	eax, 16
	mov	dword ptr [ebp - 12], eax
	movzx	eax, byte ptr [ebp - 1]
	lea	eax, dword ptr [eax + 4*eax]
	mov	dword ptr [ebp - 16], eax
	movzx	eax, byte ptr [ebp - 1]
	imul	eax, eax, 1419946
	shr	eax, 16
	mov	dword ptr [ebp - 20], eax
	mov	cx, word ptr [ebp - 12]
	mov	word ptr [ebp - 32], cx
	mov	cx, word ptr [ebp - 16]
	mov	word ptr [ebp - 30], cx
	mov	cx, word ptr [ebp - 20]
	mov	word ptr [ebp - 28], cx
	mov	eax, dword ptr [ebp - 36] ## 4-byte Reload
	mov	word ptr [eax + 4], cx
	mov	edx, dword ptr [ebp - 32]
	mov	dword ptr [eax], edx
	add	esp, 40
	pop	ebp
	ret	4

	.align	4, 0x90
_ride_ratings_get_gforce_ratings:       ## @ride_ratings_get_gforce_ratings
## BB#0:
	push	ebp
	mov	ebp, esp
	push	edi
	push	esi
	sub	esp, 48
	mov	eax, dword ptr [ebp + 12]
	mov	ecx, dword ptr [ebp + 8]
	mov	edx, 4294967046
	mov	dword ptr [ebp - 12], eax
	mov	word ptr [ebp - 24], 0
	mov	word ptr [ebp - 22], 0
	mov	word ptr [ebp - 20], 0
	mov	eax, dword ptr [ebp - 12]
	movsx	eax, word ptr [eax + 252]
	imul	eax, eax, 5242
	sar	eax, 16
	movsx	esi, word ptr [ebp - 24]
	add	esi, eax
	mov	di, si
	mov	word ptr [ebp - 24], di
	mov	eax, dword ptr [ebp - 12]
	movsx	eax, word ptr [eax + 252]
	imul	eax, eax, 52428
	sar	eax, 16
	movsx	esi, word ptr [ebp - 22]
	add	esi, eax
	mov	di, si
	mov	word ptr [ebp - 22], di
	mov	eax, dword ptr [ebp - 12]
	movsx	eax, word ptr [eax + 252]
	imul	eax, eax, 17039
	sar	eax, 16
	movsx	esi, word ptr [ebp - 20]
	add	esi, eax
	mov	di, si
	mov	word ptr [ebp - 20], di
	mov	eax, dword ptr [ebp - 12]
	mov	di, word ptr [eax + 254]
	mov	word ptr [ebp - 26], di
	movsx	eax, word ptr [ebp - 26]
	cmp	edx, eax
	mov	dword ptr [ebp - 32], ecx ## 4-byte Spill
	jle	LBB51_2
## BB#1:
	mov	eax, 4294967046
	mov	dword ptr [ebp - 36], eax ## 4-byte Spill
	jmp	LBB51_3
LBB51_2:
	movsx	eax, word ptr [ebp - 26]
	mov	dword ptr [ebp - 36], eax ## 4-byte Spill
LBB51_3:
	mov	eax, dword ptr [ebp - 36] ## 4-byte Reload
	mov	ecx, 0
	cmp	ecx, eax
	jge	LBB51_5
## BB#4:
	mov	eax, 0
	mov	dword ptr [ebp - 40], eax ## 4-byte Spill
	jmp	LBB51_9
LBB51_5:
	mov	eax, 4294967046
	movsx	ecx, word ptr [ebp - 26]
	cmp	eax, ecx
	jle	LBB51_7
## BB#6:
	mov	eax, 4294967046
	mov	dword ptr [ebp - 44], eax ## 4-byte Spill
	jmp	LBB51_8
LBB51_7:
	movsx	eax, word ptr [ebp - 26]
	mov	dword ptr [ebp - 44], eax ## 4-byte Spill
LBB51_8:
	mov	eax, dword ptr [ebp - 44] ## 4-byte Reload
	mov	dword ptr [ebp - 40], eax ## 4-byte Spill
LBB51_9:
	mov	eax, dword ptr [ebp - 40] ## 4-byte Reload
	mov	ecx, 150
	imul	eax, eax, 4294951568
	sar	eax, 16
	movsx	edx, word ptr [ebp - 24]
	add	edx, eax
	mov	si, dx
	mov	word ptr [ebp - 24], si
	movsx	eax, word ptr [ebp - 26]
	sub	eax, 100
	imul	eax, eax, 4294914868
	sar	eax, 16
	movsx	edx, word ptr [ebp - 22]
	add	edx, eax
	mov	si, dx
	mov	word ptr [ebp - 22], si
	movsx	eax, word ptr [ebp - 26]
	sub	eax, 100
	imul	eax, eax, 4294952733
	sar	eax, 16
	movsx	edx, word ptr [ebp - 20]
	add	edx, eax
	mov	si, dx
	mov	word ptr [ebp - 20], si
	mov	eax, dword ptr [ebp - 12]
	movsx	eax, word ptr [eax + 256]
	cmp	ecx, eax
	jge	LBB51_11
## BB#10:
	mov	eax, 150
	mov	dword ptr [ebp - 48], eax ## 4-byte Spill
	jmp	LBB51_12
LBB51_11:
	mov	eax, dword ptr [ebp - 12]
	movsx	eax, word ptr [eax + 256]
	mov	dword ptr [ebp - 48], eax ## 4-byte Spill
LBB51_12:
	mov	eax, dword ptr [ebp - 48] ## 4-byte Reload
	imul	eax, eax, 26214
	sar	eax, 16
	movsx	ecx, word ptr [ebp - 24]
	add	ecx, eax
	mov	dx, cx
	mov	word ptr [ebp - 24], dx
	mov	eax, dword ptr [ebp - 12]
	movsx	eax, word ptr [eax + 256]
	movsx	ecx, word ptr [ebp - 22]
	add	ecx, eax
	mov	dx, cx
	mov	word ptr [ebp - 22], dx
	mov	eax, dword ptr [ebp - 12]
	movsx	eax, word ptr [eax + 256]
	imul	eax, eax, 21845
	sar	eax, 16
	movsx	ecx, word ptr [ebp - 20]
	add	ecx, eax
	mov	dx, cx
	mov	word ptr [ebp - 20], dx
	mov	eax, dword ptr [ebp - 12]
	movsx	eax, word ptr [eax + 256]
	cmp	eax, 280
	jle	LBB51_14
## BB#13:
	movsx	eax, word ptr [ebp - 22]
	add	eax, 375
	mov	cx, ax
	mov	word ptr [ebp - 22], cx
	movsx	eax, word ptr [ebp - 20]
	add	eax, 200
	mov	cx, ax
	mov	word ptr [ebp - 20], cx
LBB51_14:
	mov	eax, dword ptr [ebp - 12]
	movsx	eax, word ptr [eax + 256]
	cmp	eax, 310
	jle	LBB51_16
## BB#15:
	mov	eax, 2
	movsx	ecx, word ptr [ebp - 24]
	mov	dword ptr [ebp - 52], eax ## 4-byte Spill
	mov	eax, ecx
	cdq
	mov	ecx, dword ptr [ebp - 52] ## 4-byte Reload
	idiv	ecx
	mov	si, ax
	mov	word ptr [ebp - 24], si
	movsx	eax, word ptr [ebp - 22]
	add	eax, 850
	mov	si, ax
	mov	word ptr [ebp - 22], si
	movsx	eax, word ptr [ebp - 20]
	add	eax, 400
	mov	si, ax
	mov	word ptr [ebp - 20], si
LBB51_16:
	mov	ax, word ptr [ebp - 20]
	mov	ecx, dword ptr [ebp - 32] ## 4-byte Reload
	mov	word ptr [ecx + 4], ax
	mov	edx, dword ptr [ebp - 24]
	mov	dword ptr [ecx], edx
	add	esp, 48
	pop	esi
	pop	edi
	pop	ebp
	ret	4

	.align	4, 0x90
_ride_ratings_get_drop_ratings:         ## @ride_ratings_get_drop_ratings
## BB#0:
	push	ebp
	mov	ebp, esp
	sub	esp, 32
	mov	eax, dword ptr [ebp + 12]
	mov	ecx, dword ptr [ebp + 8]
	mov	edx, 9
	mov	dword ptr [ebp - 4], eax
	mov	word ptr [ebp - 16], 0
	mov	word ptr [ebp - 14], 0
	mov	word ptr [ebp - 12], 0
	mov	eax, dword ptr [ebp - 4]
	movzx	eax, byte ptr [eax + 277]
	and	eax, 63
	mov	dword ptr [ebp - 20], eax
	cmp	edx, dword ptr [ebp - 20]
	mov	dword ptr [ebp - 24], ecx ## 4-byte Spill
	jge	LBB52_2
## BB#1:
	mov	eax, 9
	mov	dword ptr [ebp - 28], eax ## 4-byte Spill
	jmp	LBB52_3
LBB52_2:
	mov	eax, dword ptr [ebp - 20]
	mov	dword ptr [ebp - 28], eax ## 4-byte Spill
LBB52_3:
	mov	eax, dword ptr [ebp - 28] ## 4-byte Reload
	imul	eax, eax, 728177
	shr	eax, 16
	movzx	ecx, word ptr [ebp - 16]
	add	ecx, eax
	mov	dx, cx
	mov	word ptr [ebp - 16], dx
	mov	eax, dword ptr [ebp - 20]
	imul	eax, eax, 928426
	shr	eax, 16
	movzx	ecx, word ptr [ebp - 14]
	add	ecx, eax
	mov	dx, cx
	mov	word ptr [ebp - 14], dx
	mov	eax, dword ptr [ebp - 20]
	imul	eax, eax, 655360
	shr	eax, 16
	movzx	ecx, word ptr [ebp - 12]
	add	ecx, eax
	mov	dx, cx
	mov	word ptr [ebp - 12], dx
	mov	eax, dword ptr [ebp - 4]
	movzx	eax, byte ptr [eax + 279]
	imul	eax, eax, 32000
	shr	eax, 16
	movzx	ecx, word ptr [ebp - 16]
	add	ecx, eax
	mov	dx, cx
	mov	word ptr [ebp - 16], dx
	mov	eax, dword ptr [ebp - 4]
	movzx	eax, byte ptr [eax + 279]
	imul	eax, eax, 64000
	shr	eax, 16
	movzx	ecx, word ptr [ebp - 14]
	add	ecx, eax
	mov	dx, cx
	mov	word ptr [ebp - 14], dx
	mov	eax, dword ptr [ebp - 4]
	movzx	eax, byte ptr [eax + 279]
	imul	eax, eax, 20480
	shr	eax, 16
	movzx	ecx, word ptr [ebp - 12]
	add	ecx, eax
	mov	dx, cx
	mov	word ptr [ebp - 12], dx
	mov	eax, dword ptr [ebp - 24] ## 4-byte Reload
	mov	word ptr [eax + 4], dx
	mov	ecx, dword ptr [ebp - 16]
	mov	dword ptr [eax], ecx
	add	esp, 32
	pop	ebp
	ret	4

	.align	4, 0x90
_ride_ratings_update_state_2:           ## @ride_ratings_update_state_2
## BB#0:
	push	ebp
	mov	ebp, esp
	sub	esp, 72
	mov	eax, 7036006
	mov	dword ptr [esp], 7036006
	mov	dword ptr [ebp - 64], eax ## 4-byte Spill
	call	_RCT2_CALLPROC_EBPSAFE
	mov	dword ptr [ebp - 68], eax ## 4-byte Spill
	add	esp, 72
	pop	ebp
	ret

	.align	4, 0x90
_ride_ratings_update_state_1:           ## @ride_ratings_update_state_1
## BB#0:
	push	ebp
	mov	ebp, esp
	push	ebx
	push	edi
	push	esi
	sub	esp, 108
	mov	eax, 20493774
	mov	ecx, 20493713
	mov	edx, 20493772
	mov	esi, 20493770
	mov	edi, 20493768
	mov	ebx, 20493766
	mov	dword ptr [ebp - 16], eax ## 4-byte Spill
	mov	eax, 20493764
	mov	dword ptr [ebp - 20], eax ## 4-byte Spill
	mov	eax, 20493762
	mov	dword ptr [ebp - 24], eax ## 4-byte Spill
	mov	eax, 20493760
	mov	dword ptr [ebp - 28], eax ## 4-byte Spill
	mov	eax, 20493758
	mov	dword ptr [ebp - 32], eax ## 4-byte Spill
	mov	eax, 20493756
	mov	dword ptr [ebp - 36], eax ## 4-byte Spill
	mov	eax, 20493754
	mov	dword ptr [ebp - 40], eax ## 4-byte Spill
	mov	eax, 20493752
	mov	dword ptr [ebp - 44], eax ## 4-byte Spill
	mov	eax, 20493750
	mov	dword ptr [ebp - 48], eax ## 4-byte Spill
	mov	eax, 20493748
	mov	dword ptr [ebp - 52], eax ## 4-byte Spill
	mov	eax, 20493746
	mov	dword ptr [ebp - 56], eax ## 4-byte Spill
	mov	eax, 20493744
	mov	dword ptr [ebp - 60], eax ## 4-byte Spill
	mov	eax, 20493742
	mov	dword ptr [ebp - 64], eax ## 4-byte Spill
	mov	eax, 20493740
	mov	dword ptr [ebp - 68], eax ## 4-byte Spill
	mov	eax, 20493738
	mov	dword ptr [ebp - 72], eax ## 4-byte Spill
	mov	eax, 20493736
	mov	dword ptr [ebp - 76], eax ## 4-byte Spill
	mov	eax, 20493734
	mov	dword ptr [ebp - 80], eax ## 4-byte Spill
	mov	eax, 20493732
	mov	dword ptr [ebp - 84], eax ## 4-byte Spill
	mov	eax, 20493730
	mov	dword ptr [ebp - 88], eax ## 4-byte Spill
	mov	eax, 20493728
	mov	dword ptr [ebp - 92], eax ## 4-byte Spill
	mov	eax, 20493726
	mov	dword ptr [ebp - 96], eax ## 4-byte Spill
	mov	eax, 20493724
	mov	dword ptr [ebp - 100], eax ## 4-byte Spill
	mov	eax, 20493722
	mov	dword ptr [ebp - 104], eax ## 4-byte Spill
	mov	eax, 20493720
	mov	dword ptr [ebp - 108], eax ## 4-byte Spill
	mov	eax, 20493718
	mov	dword ptr [ebp - 112], eax ## 4-byte Spill
	mov	eax, 20493716
	mov	word ptr [eax], 0
	mov	eax, dword ptr [ebp - 112] ## 4-byte Reload
	mov	word ptr [eax], 0
	mov	eax, dword ptr [ebp - 108] ## 4-byte Reload
	mov	word ptr [eax], 0
	mov	eax, dword ptr [ebp - 104] ## 4-byte Reload
	mov	word ptr [eax], 0
	mov	eax, dword ptr [ebp - 100] ## 4-byte Reload
	mov	word ptr [eax], 0
	mov	eax, dword ptr [ebp - 96] ## 4-byte Reload
	mov	word ptr [eax], 0
	mov	eax, dword ptr [ebp - 92] ## 4-byte Reload
	mov	word ptr [eax], 0
	mov	eax, dword ptr [ebp - 88] ## 4-byte Reload
	mov	word ptr [eax], 0
	mov	eax, dword ptr [ebp - 84] ## 4-byte Reload
	mov	word ptr [eax], 0
	mov	eax, dword ptr [ebp - 80] ## 4-byte Reload
	mov	word ptr [eax], 0
	mov	eax, dword ptr [ebp - 76] ## 4-byte Reload
	mov	word ptr [eax], 0
	mov	eax, dword ptr [ebp - 72] ## 4-byte Reload
	mov	word ptr [eax], 0
	mov	eax, dword ptr [ebp - 68] ## 4-byte Reload
	mov	word ptr [eax], 0
	mov	eax, dword ptr [ebp - 64] ## 4-byte Reload
	mov	word ptr [eax], 0
	mov	eax, dword ptr [ebp - 60] ## 4-byte Reload
	mov	word ptr [eax], 0
	mov	eax, dword ptr [ebp - 56] ## 4-byte Reload
	mov	word ptr [eax], 0
	mov	eax, dword ptr [ebp - 52] ## 4-byte Reload
	mov	word ptr [eax], 0
	mov	eax, dword ptr [ebp - 48] ## 4-byte Reload
	mov	word ptr [eax], 0
	mov	eax, dword ptr [ebp - 44] ## 4-byte Reload
	mov	word ptr [eax], 0
	mov	eax, dword ptr [ebp - 40] ## 4-byte Reload
	mov	word ptr [eax], 0
	mov	eax, dword ptr [ebp - 36] ## 4-byte Reload
	mov	word ptr [eax], 0
	mov	eax, dword ptr [ebp - 32] ## 4-byte Reload
	mov	word ptr [eax], 0
	mov	eax, dword ptr [ebp - 28] ## 4-byte Reload
	mov	word ptr [eax], 0
	mov	eax, dword ptr [ebp - 24] ## 4-byte Reload
	mov	word ptr [eax], 0
	mov	eax, dword ptr [ebp - 20] ## 4-byte Reload
	mov	word ptr [eax], 0
	mov	word ptr [ebx], 0
	mov	word ptr [edi], 0
	mov	word ptr [esi], 0
	mov	word ptr [edx], 0
	mov	byte ptr [ecx], 2
	mov	ecx, dword ptr [ebp - 16] ## 4-byte Reload
	mov	word ptr [ecx], 0
	call	_loc_6B5BB2
	add	esp, 108
	pop	esi
	pop	edi
	pop	ebx
	pop	ebp
	ret

	.align	4, 0x90
_ride_ratings_update_state_0:           ## @ride_ratings_update_state_0
## BB#0:
	push	ebp
	mov	ebp, esp
	push	ebx
	sub	esp, 8
	call	L55$pb
L55$pb:
	pop	eax
	mov	ecx, 20493712
	movzx	edx, byte ptr [ecx]
	add	edx, 1
	mov	bl, dl
	mov	byte ptr [ecx], bl
	movzx	ecx, byte ptr [ecx]
	cmp	ecx, 255
	mov	dword ptr [ebp - 12], eax ## 4-byte Spill
	jne	LBB55_2
## BB#1:
	mov	eax, 20493712
	mov	byte ptr [eax], 0
LBB55_2:
	mov	eax, dword ptr [ebp - 12] ## 4-byte Reload
	mov	ecx, dword ptr [eax + L_g_ride_list$non_lazy_ptr-L55$pb]
	mov	edx, 20493712
	movzx	edx, byte ptr [edx]
	mov	ecx, dword ptr [ecx]
	imul	edx, edx, 608
	add	ecx, edx
	mov	dword ptr [ebp - 8], ecx
	mov	ecx, dword ptr [ebp - 8]
	movzx	ecx, byte ptr [ecx]
	cmp	ecx, 255
	je	LBB55_5
## BB#3:
	mov	eax, dword ptr [ebp - 8]
	movzx	eax, byte ptr [eax + 73]
	cmp	eax, 0
	je	LBB55_5
## BB#4:
	mov	eax, 20493713
	mov	byte ptr [eax], 1
LBB55_5:
	add	esp, 8
	pop	ebx
	pop	ebp
	ret

	.comm	_gRideTypeList,4,2      ## @gRideTypeList
	.comm	_WINDOW_CLASS,4,2       ## @WINDOW_CLASS
	.section	__DATA,__const
	.align	2                       ## @ride_ratings_calculate_func_table
_ride_ratings_calculate_func_table:
	.long	0
	.long	0
	.long	0
	.long	0
	.long	0
	.long	0
	.long	0
	.long	0
	.long	0
	.long	0
	.long	0
	.long	0
	.long	0
	.long	0
	.long	0
	.long	0
	.long	0
	.long	_ride_ratings_calculate_mine_train_coaster
	.long	0
	.long	0
	.long	_ride_ratings_calculate_maze
	.long	_ride_ratings_calculate_spiral_slide
	.long	0
	.long	0
	.long	0
	.long	0
	.long	_ride_ratings_calculate_pirate_ship
	.long	_ride_ratings_calculate_inverter_ship
	.long	_ride_ratings_calculate_food_stall
	.long	0
	.long	_ride_ratings_calculate_drink_stall
	.long	0
	.long	_ride_ratings_calculate_shop
	.long	_ride_ratings_calculate_merry_go_round
	.long	0
	.long	_ride_ratings_calculate_information_kiosk
	.long	_ride_ratings_calculate_bathroom
	.long	_ride_ratings_calculate_ferris_wheel
	.long	_ride_ratings_calculate_motion_simulator
	.long	_ride_ratings_calculate_3d_cinema
	.long	_ride_ratings_calculate_top_spin
	.long	_ride_ratings_calculate_space_rings
	.long	0
	.long	_ride_ratings_calculate_elevator
	.long	0
	.long	_ride_ratings_calculate_atm
	.long	_ride_ratings_calculate_twist
	.long	_ride_ratings_calculate_haunted_house
	.long	_ride_ratings_calculate_first_aid
	.long	_ride_ratings_calculate_circus_show
	.long	0
	.long	0
	.long	0
	.long	0
	.long	0
	.long	0
	.long	0
	.long	0
	.long	0
	.long	0
	.long	0
	.long	0
	.long	0
	.long	0
	.long	0
	.long	0
	.long	0
	.long	_ride_ratings_calculate_mini_golf
	.long	0
	.long	0
	.long	0
	.long	_ride_ratings_calculate_crooked_house
	.long	0
	.long	0
	.long	0
	.long	0
	.long	0
	.long	_ride_ratings_calculate_magic_carpet
	.long	0
	.long	0
	.long	0
	.long	_ride_ratings_calculate_enterprise
	.long	0
	.long	0
	.long	0
	.long	0
	.long	0
	.long	0
	.long	0
	.long	0
	.long	0

	.section	__TEXT,__const
	.align	1                       ## @ride_ratings_apply_intensity_penalty.intensityBounds
_ride_ratings_apply_intensity_penalty.intensityBounds:
	.short	1000                    ## 0x3e8
	.short	1100                    ## 0x44c
	.short	1200                    ## 0x4b0
	.short	1320                    ## 0x528
	.short	1450                    ## 0x5aa


	.section	__IMPORT,__pointers,non_lazy_symbol_pointers
L_costPerTrackPiece$non_lazy_ptr:
	.indirect_symbol	_costPerTrackPiece
	.long	0
L_gRideTypeList$non_lazy_ptr:
	.indirect_symbol	_gRideTypeList
	.long	0
L_g_ride_list$non_lazy_ptr:
	.indirect_symbol	_g_ride_list
	.long	0
L_hasRunningTrack$non_lazy_ptr:
	.indirect_symbol	_hasRunningTrack
	.long	0
L_initialUpkeepCosts$non_lazy_ptr:
	.indirect_symbol	_initialUpkeepCosts
	.long	0
L_rideUnknownData1$non_lazy_ptr:
	.indirect_symbol	_rideUnknownData1
	.long	0
L_rideUnknownData2$non_lazy_ptr:
	.indirect_symbol	_rideUnknownData2
	.long	0
L_rideUnknownData3$non_lazy_ptr:
	.indirect_symbol	_rideUnknownData3
	.long	0

.subsections_via_symbols
