/*
 * Copyright (C) 2023 The GDLang Team.
 *
 * This file is part of GDLang.
 *
 * GDLang is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * GDLang is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with GDLang.  If not, see <http://www.gnu.org/licenses/>.
 */

// Code generated by goyacc -l -o src/gd/ast/gd.y.go src/gd/ast/gd.y. DO NOT EDIT.
package ast

import __yyfmt__ "fmt"

import (
	"gdlang/lib/runtime"
)

type yySymType struct {
	yys   int
	token *NodeTokenInfo

	node      Node
	node_list []Node

	flag bool

	gd_type      runtime.GDTypable
	gd_type_list []runtime.GDTypable
}

const LAS = 57346
const LLSHIFT = 57347
const LRSHIFT = 57348
const LIDENT = 57349
const LINT = 57350
const LFLOAT = 57351
const LSTRING = 57352
const LIMAG = 57353
const LCHAR = 57354
const LCOMMENT = 57355
const LMUL = 57356
const LADD = 57357
const LSUB = 57358
const LQUO = 57359
const LREM = 57360
const LQMARK = 57361
const LNSAFE = 57362
const LADD_ASSIGN = 57363
const LSUB_ASSIGN = 57364
const LMUL_ASSIGN = 57365
const LQUO_ASSIGN = 57366
const LREM_ASSIGN = 57367
const LARROW = 57368
const LINC = 57369
const LDEC = 57370
const LLAND = 57371
const LOR = 57372
const LLOR = 57373
const LNOT = 57374
const LEQL = 57375
const LLSS = 57376
const LGTR = 57377
const LASSIGN = 57378
const LNEQ = 57379
const LLEQ = 57380
const LGEQ = 57381
const LELLIPSIS = 57382
const LLPAREN = 57383
const LLBRACK = 57384
const LLBRACE = 57385
const LCOMMA = 57386
const LPERIOD = 57387
const LRPAREN = 57388
const LRBRACK = 57389
const LRBRACE = 57390
const LSEMICOLON = 57391
const LCOLON = 57392
const LCOLONCOLON = 57393
const LUSE = 57394
const LTYPEALIAS = 57395
const LSET = 57396
const LPUB = 57397
const LCONST = 57398
const LELSE = 57399
const LFOR = 57400
const LIN = 57401
const LFUNC = 57402
const LIF = 57403
const LBREAK = 57404
const LRETURN = 57405
const LTANY = 57406
const LTBOOL = 57407
const LTINT = 57408
const LTFLOAT = 57409
const LTCOMPLEX = 57410
const LTSTRING = 57411
const LTCHAR = 57412
const LTRUE = 57413
const LFALSE = 57414
const LNIL = 57415

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"LAS",
	"LLSHIFT",
	"LRSHIFT",
	"LIDENT",
	"LINT",
	"LFLOAT",
	"LSTRING",
	"LIMAG",
	"LCHAR",
	"LCOMMENT",
	"LMUL",
	"LADD",
	"LSUB",
	"LQUO",
	"LREM",
	"LQMARK",
	"LNSAFE",
	"LADD_ASSIGN",
	"LSUB_ASSIGN",
	"LMUL_ASSIGN",
	"LQUO_ASSIGN",
	"LREM_ASSIGN",
	"LARROW",
	"LINC",
	"LDEC",
	"LLAND",
	"LOR",
	"LLOR",
	"LNOT",
	"LEQL",
	"LLSS",
	"LGTR",
	"LASSIGN",
	"LNEQ",
	"LLEQ",
	"LGEQ",
	"LELLIPSIS",
	"LLPAREN",
	"LLBRACK",
	"LLBRACE",
	"LCOMMA",
	"LPERIOD",
	"LRPAREN",
	"LRBRACK",
	"LRBRACE",
	"LSEMICOLON",
	"LCOLON",
	"LCOLONCOLON",
	"LUSE",
	"LTYPEALIAS",
	"LSET",
	"LPUB",
	"LCONST",
	"LELSE",
	"LFOR",
	"LIN",
	"LFUNC",
	"LIF",
	"LBREAK",
	"LRETURN",
	"LTANY",
	"LTBOOL",
	"LTINT",
	"LTFLOAT",
	"LTCOMPLEX",
	"LTSTRING",
	"LTCHAR",
	"LTRUE",
	"LFALSE",
	"LNIL",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

var yyExca = [...]int16{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	1, 13,
	-2, 20,
	-1, 15,
	1, 12,
	-2, 20,
	-1, 153,
	49, 24,
	-2, 108,
	-1, 157,
	49, 28,
	-2, 135,
	-1, 159,
	49, 30,
	-2, 139,
}

const yyPrivate = 57344

const yyLast = 663

var yyAct = [...]int16{
	80, 136, 225, 139, 177, 135, 205, 84, 151, 62,
	28, 27, 50, 63, 35, 147, 81, 277, 16, 134,
	47, 81, 58, 43, 253, 29, 254, 96, 13, 141,
	21, 19, 19, 5, 279, 10, 237, 20, 46, 222,
	208, 169, 79, 30, 31, 213, 33, 127, 77, 76,
	15, 39, 14, 11, 86, 44, 263, 243, 131, 212,
	40, 88, 101, 102, 85, 234, 203, 122, 125, 124,
	259, 240, 126, 128, 129, 130, 143, 231, 142, 230,
	100, 202, 167, 163, 93, 92, 89, 90, 91, 94,
	95, 157, 103, 138, 153, 159, 140, 81, 256, 272,
	39, 23, 256, 236, 207, 204, 168, 180, 104, 182,
	183, 184, 185, 186, 187, 188, 189, 190, 191, 192,
	193, 194, 195, 196, 232, 178, 198, 44, 179, 34,
	174, 81, 14, 41, 144, 36, 14, 14, 233, 14,
	42, 226, 209, 172, 206, 199, 197, 117, 8, 26,
	118, 119, 158, 14, 77, 211, 175, 210, 171, 261,
	176, 181, 18, 214, 24, 228, 223, 106, 120, 121,
	41, 37, 224, 32, 221, 154, 227, 117, 115, 116,
	118, 119, 105, 235, 45, 17, 78, 30, 99, 98,
	87, 85, 108, 4, 107, 97, 109, 111, 112, 170,
	110, 113, 114, 9, 123, 22, 242, 241, 49, 75,
	245, 244, 25, 96, 82, 163, 83, 247, 248, 249,
	250, 251, 252, 157, 229, 257, 153, 159, 255, 246,
	140, 117, 115, 116, 118, 119, 12, 3, 2, 270,
	266, 264, 137, 38, 1, 55, 275, 88, 101, 102,
	173, 276, 160, 48, 155, 267, 269, 156, 242, 178,
	268, 61, 271, 146, 145, 132, 100, 7, 273, 274,
	93, 92, 89, 90, 91, 94, 95, 6, 280, 60,
	59, 258, 282, 152, 57, 281, 51, 260, 262, 148,
	149, 150, 265, 0, 14, 64, 65, 66, 70, 71,
	0, 0, 52, 53, 0, 0, 0, 0, 0, 0,
	0, 0, 117, 115, 116, 118, 119, 0, 0, 54,
	0, 0, 0, 0, 0, 0, 0, 278, 56, 72,
	73, 109, 111, 112, 0, 110, 113, 114, 0, 0,
	21, 19, 0, 0, 0, 165, 0, 164, 166, 162,
	161, 0, 0, 0, 0, 0, 0, 0, 68, 69,
	67, 14, 64, 65, 66, 70, 71, 0, 0, 52,
	53, 0, 0, 0, 0, 14, 64, 65, 66, 70,
	71, 0, 0, 52, 53, 0, 54, 0, 0, 0,
	0, 0, 0, 0, 0, 56, 72, 73, 133, 0,
	54, 0, 0, 0, 0, 0, 0, 0, 0, 56,
	72, 73, 0, 0, 74, 0, 0, 0, 0, 106,
	120, 121, 0, 0, 0, 68, 69, 67, 74, 117,
	115, 116, 118, 119, 105, 0, 0, 0, 0, 68,
	69, 67, 0, 0, 108, 0, 107, 0, 109, 111,
	112, 0, 110, 113, 114, 0, 106, 120, 121, 0,
	0, 0, 0, 0, 0, 238, 117, 115, 116, 118,
	119, 105, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 108, 0, 107, 0, 109, 111, 112, 0, 110,
	113, 114, 106, 120, 121, 0, 0, 0, 0, 239,
	0, 0, 117, 115, 116, 118, 119, 105, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 108, 0, 107,
	0, 109, 111, 112, 0, 110, 113, 114, 106, 120,
	121, 0, 201, 0, 200, 0, 0, 0, 117, 115,
	116, 118, 119, 105, 0, 216, 217, 218, 219, 220,
	0, 0, 0, 108, 0, 107, 0, 109, 111, 112,
	215, 110, 113, 114, 106, 120, 121, 0, 0, 0,
	0, 0, 0, 0, 117, 115, 116, 118, 119, 105,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 108,
	0, 107, 0, 109, 111, 112, 0, 110, 113, 114,
	0, 0, 0, 81, 117, 115, 116, 118, 119, 105,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 108,
	0, 107, 0, 109, 111, 112, 0, 110, 113, 114,
	117, 115, 116, 118, 119, 0, 0, 117, 115, 116,
	118, 119, 0, 0, 0, 108, 0, 107, 0, 109,
	111, 112, 108, 110, 113, 114, 109, 111, 112, 0,
	110, 113, 114,
}

var yyPact = [...]int16{
	-19, -1000, -20, 4, -1000, 146, -1000, 1, -1000, -23,
	-1000, -19, 56, -1000, -1000, -20, -1000, -1000, -1000, -31,
	146, 146, -1000, 132, -1000, 85, -1000, 99, 130, -1000,
	92, 104, 146, -1000, -31, -1000, 368, -31, -1000, -8,
	88, 146, 20, 46, 64, -1000, 163, -1000, -1000, -1000,
	-1000, 27, 368, 368, 368, -1000, 354, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 368, 45, 92, 32, -1000, 146, -1000, 20,
	-1000, 287, 36, 62, -1000, -9, -1000, -1000, 206, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	92, 20, 146, -1000, 146, 368, 20, 368, 368, 368,
	368, 368, 368, 368, 368, 368, 368, 368, 368, 368,
	368, 368, -1000, 146, 368, 368, -1000, -1000, -1000, -1000,
	-1000, 488, 35, -1000, 19, 61, 163, 60, -1000, -1000,
	-10, 88, 99, -31, -1000, 11, -4, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 368, -1000, 524, 129, -22, 368, 115, 125, 20,
	33, 31, 94, -1000, -1000, 18, 59, -1000, -14, -1000,
	415, -1000, 623, 298, 217, 217, 217, 217, 217, 217,
	133, 133, -1000, -1000, -1000, 590, 590, -1000, 452, 25,
	-1000, 368, -1000, -1000, 368, -1000, 9, 146, 368, -1000,
	-1000, -1000, -1000, 287, 163, 368, 368, 368, 368, 368,
	368, -35, 368, -1000, 54, -1000, 20, -1000, 24, -1000,
	-1000, -1000, 20, 20, -1000, 8, 146, 20, 368, -1000,
	-1000, 58, 163, -1000, -1000, 163, -1000, 163, 163, 163,
	163, 163, 163, 368, 368, 54, 368, -1000, -1000, 115,
	-1000, 55, -1000, -1000, -1000, -1000, 616, 560, 54, -1000,
	-40, -1000, 20, -1000, -1000, -1000, -1000, -27, -1000, 368,
	-1000, 54, -1000,
}

var yyPgo = [...]int16{
	0, 148, 291, 290, 289, 1, 13, 20, 286, 8,
	12, 284, 283, 0, 15, 175, 9, 280, 279, 19,
	277, 267, 5, 265, 264, 263, 261, 3, 257, 254,
	253, 252, 251, 246, 245, 22, 244, 193, 7, 243,
	14, 11, 242, 239, 238, 237, 236, 23, 216, 214,
	212, 209, 149, 152, 208, 204, 10, 203, 6, 2,
	29, 54, 199, 195, 190, 189, 188, 4, 186, 160,
	159, 158,
}

var yyR1 = [...]int8{
	0, 36, 44, 44, 45, 45, 37, 37, 46, 46,
	47, 47, 20, 20, 21, 21, 1, 1, 1, 57,
	57, 53, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 58, 58, 9, 50, 50, 52, 52, 51,
	51, 41, 40, 40, 56, 56, 12, 12, 12, 12,
	12, 12, 39, 68, 68, 38, 64, 64, 64, 64,
	64, 64, 64, 64, 64, 64, 64, 64, 62, 61,
	61, 63, 71, 71, 71, 65, 66, 48, 48, 49,
	49, 59, 59, 60, 60, 69, 69, 67, 70, 70,
	13, 14, 14, 14, 3, 3, 2, 24, 24, 25,
	25, 16, 15, 30, 54, 5, 5, 5, 5, 5,
	5, 5, 5, 5, 5, 5, 5, 5, 5, 5,
	5, 5, 7, 7, 7, 7, 8, 8, 10, 10,
	34, 34, 34, 34, 34, 34, 34, 34, 34, 34,
	6, 55, 55, 22, 22, 19, 19, 11, 11, 11,
	11, 11, 11, 11, 11, 35, 17, 26, 26, 42,
	42, 27, 23, 23, 23, 18, 29, 28, 28, 28,
	31, 43, 43, 32, 33, 33,
}

var yyR2 = [...]int8{
	0, 2, 2, 0, 3, 1, 6, 2, 3, 1,
	3, 1, 2, 0, 3, 1, 2, 2, 2, 1,
	0, 4, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 0, 2, 3, 1, 2, 5, 3,
	1, 2, 2, 0, 1, 0, 3, 3, 3, 3,
	3, 3, 2, 2, 0, 3, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 2, 3, 1,
	3, 3, 1, 2, 3, 3, 4, 3, 1, 1,
	0, 2, 0, 4, 6, 3, 1, 3, 3, 1,
	3, 1, 1, 1, 1, 2, 1, 2, 0, 3,
	1, 3, 4, 5, 3, 1, 1, 1, 1, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 1, 2, 2, 2, 1, 3, 3, 3,
	1, 1, 1, 1, 1, 1, 2, 3, 4, 1,
	4, 1, 1, 3, 1, 2, 0, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 3, 4, 2, 3,
	1, 3, 1, 2, 3, 3, 5, 5, 4, 2,
	5, 2, 0, 4, 2, 0,
}

var yyChk = [...]int16{
	-1000, -36, -44, -45, -37, 52, -20, -21, -1, -57,
	55, 49, -46, -35, 7, 49, -9, -15, -53, 54,
	60, 53, -37, 45, -1, -50, -52, -41, -56, 56,
	-35, -35, 41, -35, 44, -40, 36, 41, -39, -35,
	-60, 41, 36, -47, -35, -52, -5, -7, -30, -54,
	-10, -8, 15, 16, 32, -34, 41, -11, -35, -17,
	-18, -26, -16, -6, 8, 9, 10, 73, 71, 72,
	11, 12, 42, 43, 60, -51, -41, -56, -68, 50,
	-13, 43, -49, -48, -38, -35, -61, -64, 41, 66,
	67, 68, 65, 64, 69, 70, 7, -63, -65, -66,
	60, 42, 43, 46, 44, 19, 4, 31, 29, 33,
	37, 34, 35, 38, 39, 15, 16, 14, 17, 18,
	5, 6, 40, -55, 42, 41, 45, 20, -7, -7,
	-7, -5, -23, 44, -19, -22, -5, -42, 48, -27,
	-35, -60, 46, 44, -61, -24, -25, -14, -4, -3,
	-2, -9, -12, -10, -15, -29, -28, -16, -53, -6,
	-31, 63, 62, -5, 60, 58, 61, 46, 44, 50,
	-62, -71, -61, 44, -60, -61, -69, -67, -35, -47,
	-5, -61, -5, -5, -5, -5, -5, -5, -5, -5,
	-5, -5, -5, -5, -5, -5, -5, -35, -5, -19,
	46, 44, 46, 47, 44, -58, -58, 44, 50, -13,
	-40, -41, 48, 49, -5, 36, 21, 22, 23, 24,
	25, -9, 61, -13, -22, -59, 26, -38, 40, -61,
	46, 46, 30, 44, 47, -58, 44, 50, 50, 47,
	46, -22, -5, 48, -27, -5, -14, -5, -5, -5,
	-5, -5, -5, 59, 61, -22, 44, -13, -61, 46,
	-61, -70, -61, 48, -67, -61, -5, -5, -22, -13,
	-43, -59, 44, -13, -13, -33, -32, 57, -61, 61,
	-13, -22, -13,
}

var yyDef = [...]int16{
	3, -2, -2, 0, 5, 0, 1, 0, 15, 0,
	19, 2, 7, 9, 155, -2, 16, 17, 18, 45,
	0, 0, 4, 0, 14, 34, 36, 43, 0, 44,
	0, 0, 0, 8, 45, 37, 0, 45, 41, 54,
	0, 80, 0, 0, 11, 35, 42, 105, 106, 107,
	108, 122, 0, 0, 0, 126, 0, 130, 131, 132,
	133, 134, 135, 139, 147, 148, 149, 150, 151, 152,
	153, 154, 146, 0, 0, 0, 40, 0, 52, 0,
	102, 98, 0, 79, 78, 0, 21, 69, 0, 56,
	57, 58, 59, 60, 61, 62, 63, 64, 65, 66,
	0, 0, 0, 6, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 136, 0, 0, 146, 141, 142, 123, 124,
	125, 0, 0, 162, 0, 33, 144, 33, 158, 160,
	0, 0, 43, 45, 53, 0, 0, 100, 91, 92,
	93, 22, 23, -2, 25, 26, 27, -2, 29, -2,
	31, 94, 96, 0, 0, 0, 0, 82, 0, 0,
	0, 0, 0, 72, 67, 0, 33, 86, 0, 10,
	0, 104, 109, 110, 111, 112, 113, 114, 115, 116,
	117, 118, 119, 120, 121, 128, 129, 137, 0, 0,
	127, 163, 156, 165, 32, 145, 0, 32, 0, 101,
	38, 39, 90, 97, 95, 0, 0, 0, 0, 0,
	0, 0, 0, 169, 0, 83, 0, 77, 0, 55,
	70, 71, 0, 73, 75, 0, 32, 0, 0, 138,
	140, 164, 143, 157, 159, 161, 99, 46, 47, 48,
	49, 50, 51, 0, 0, 0, 0, 172, 81, 82,
	68, 74, 89, 76, 85, 87, 103, 0, 0, 168,
	175, 84, 0, 166, 167, 170, 171, 0, 88, 0,
	174, 0, 173,
}

var yyTok1 = [...]int8{
	1,
}

var yyTok2 = [...]int8{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 66, 67, 68, 69, 70, 71,
	72, 73,
}

var yyTok3 = [...]int8{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{
	{79, 73, "NIL_AS_A_TYPE_ERR"},
	{1, 52, "USE_ONLY_AT_HEADER_ERR"},
}

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := int(yyPact[state])
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && int(yyChk[int(yyAct[n])]) == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || int(yyExca[i+1]) != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := int(yyExca[i])
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = int(yyTok1[0])
		goto out
	}
	if char < len(yyTok1) {
		token = int(yyTok1[char])
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = int(yyTok2[char-yyPrivate])
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = int(yyTok3[i+0])
		if token == char {
			token = int(yyTok3[i+1])
			goto out
		}
	}

out:
	if token == 0 {
		token = int(yyTok2[1]) /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = int(yyPact[yystate])
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = int(yyAct[yyn])
	if int(yyChk[yyn]) == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = int(yyDef[yystate])
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && int(yyExca[xi+1]) == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = int(yyExca[xi+0])
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = int(yyExca[xi+1])
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = int(yyPact[yyS[yyp].yys]) + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = int(yyAct[yyn]) /* simulate a shift of "error" */
					if int(yyChk[yystate]) == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= int(yyR2[yyn])
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = int(yyR1[yyn])
	yyg := int(yyPgo[yyn])
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = int(yyAct[yyg])
	} else {
		yystate = int(yyAct[yyj])
		if int(yyChk[yystate]) != -yyn {
			yystate = int(yyAct[yyg])
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			file := NewNodeFile(yyDollar[1].node_list, yyDollar[2].node_list)
			yylex.(*Ast).Root = file
			yyVAL.node = file
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.node_list = yyDollar[1].node_list
		}
	case 3:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.node_list = make([]Node, 0)
		}
	case 4:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyDollar[1].node_list = append(yyDollar[1].node_list, yyDollar[3].node)
			yyVAL.node_list = yyDollar[1].node_list
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.node_list = make([]Node, 1)
			yyVAL.node_list[0] = yyDollar[1].node
		}
	case 6:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			// Package with list of names
			yyVAL.node = NewNodePackage(yyDollar[2].node_list, yyDollar[5].node_list)
		}
	case 7:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			lastIdent := yyDollar[2].node_list[len(yyDollar[2].node_list)-1]
			yyDollar[2].node_list = yyDollar[2].node_list[:len(yyDollar[2].node_list)-1]
			yyVAL.node = NewNodePackage(yyDollar[2].node_list, []Node{lastIdent})
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyDollar[1].node_list = append(yyDollar[1].node_list, yyDollar[3].node)
			yyVAL.node_list = yyDollar[1].node_list
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.node_list = []Node{yyDollar[1].node}
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyDollar[3].node_list = append(yyDollar[3].node_list, yyDollar[1].node)
			yyVAL.node_list = yyDollar[3].node_list
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.node_list = []Node{yyDollar[1].node}
		}
	case 13:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.node_list = make([]Node, 0)
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyDollar[1].node_list = append(yyDollar[1].node_list, yyDollar[3].node)
			yyVAL.node_list = yyDollar[1].node_list
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.node_list = make([]Node, 1)
			yyVAL.node_list[0] = yyDollar[1].node
		}
	case 16:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			sets, ok := yyDollar[2].node.(*NodeSets)
			if !ok {
				panic("file_body_stmt: Invalid `*NodeSets` object")
			}

			for _, node := range sets.Nodes {
				if set, ok := node.(*NodeSet); ok {
					set.IsPub = yyDollar[1].flag
				} else {
					panic("file_body_stmt: Invalid `*NodeSet` object")
				}
			}

			yyVAL.node = yyDollar[2].node
		}
	case 17:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyDollar[2].node.(*NodeFunc).IsPub = yyDollar[1].flag
			yyVAL.node = yyDollar[2].node
		}
	case 18:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyDollar[2].node.(*NodeTypeAlias).IsPub = yyDollar[1].flag
			yyVAL.node = yyDollar[2].node
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.flag = true
		}
	case 20:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.flag = false
		}
	case 21:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.node = NewNodeTypeAlias(false, yyDollar[2].node.(*NodeIdent), yyDollar[4].gd_type)
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.flag = true
		}
	case 33:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.flag = false
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.node = NewNodeSets(yyDollar[2].node_list)
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyDollar[1].node_list = append(yyDollar[1].node_list, yyDollar[3].node_list...)
			yyVAL.node_list = yyDollar[1].node_list
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			nodeSet, ok := yyDollar[1].node.(*NodeSet)
			if !ok {
				panic("set_expr_option_list: Invalid `*NodeSet` object")
			}
			nodeSet.Expr = yyDollar[2].node
			yyVAL.node_list = []Node{nodeSet}
		}
	case 38:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			sharedExpr := NewNodeSharedExpr(yyDollar[5].node)
			for i, node := range yyDollar[3].node_list {
				if set, ok := node.(*NodeSet); ok {
					set.Index = byte(i)
					set.Expr = sharedExpr
				} else {
					panic("set: Invalid `*NodeSet` object")
				}
			}
			yyVAL.node_list = yyDollar[3].node_list
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyDollar[1].node_list = append(yyDollar[1].node_list, yyDollar[3].node)
			yyVAL.node_list = yyDollar[1].node_list
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.node_list = make([]Node, 1)
			yyVAL.node_list[0] = yyDollar[1].node
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			identWithType, ok := yyDollar[2].node.(*NodeIdentWithType)
			if !ok {
				panic("const_ident_with_optional_type: Invalid `*NodeIdentWithType` object")
			}
			yyVAL.node = NewNodeSet(false, yyDollar[1].flag, identWithType, nil)
		}
	case 42:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.node = yyDollar[2].node
		}
	case 43:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.node = nil
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.flag = true
		}
	case 45:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.flag = false
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.node = NewNodeUpdateSet(yyDollar[1].node, yyDollar[3].node)
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.node = NewNodeUpdateSet(yyDollar[1].node, NewNodeExprOperation(runtime.ExprOperationAdd, yyDollar[1].node, yyDollar[3].node))
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.node = NewNodeUpdateSet(yyDollar[1].node, NewNodeExprOperation(runtime.ExprOperationSubtract, yyDollar[1].node, yyDollar[3].node))
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.node = NewNodeUpdateSet(yyDollar[1].node, NewNodeExprOperation(runtime.ExprOperationMultiply, yyDollar[1].node, yyDollar[3].node))
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.node = NewNodeUpdateSet(yyDollar[1].node, NewNodeExprOperation(runtime.ExprOperationQuo, yyDollar[1].node, yyDollar[3].node))
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.node = NewNodeUpdateSet(yyDollar[1].node, NewNodeExprOperation(runtime.ExprOperationRem, yyDollar[1].node, yyDollar[3].node))
		}
	case 52:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.node = NewNodeIdentWithType(yyDollar[1].node.(*NodeIdent), yyDollar[2].gd_type)
		}
	case 53:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.gd_type = yyDollar[2].gd_type
		}
	case 54:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.gd_type = runtime.GDUntypedType
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.node = NewNodeIdentWithType(yyDollar[1].node.(*NodeIdent), yyDollar[3].gd_type)
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.gd_type = runtime.GDIntType
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.gd_type = runtime.GDFloatType
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.gd_type = runtime.GDComplexType
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.gd_type = runtime.GDBoolType
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.gd_type = runtime.GDAnyType
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.gd_type = runtime.GDStringType
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.gd_type = runtime.GDCharType
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.gd_type = runtime.GDStringIdentType(yyDollar[1].token.Lit)
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.gd_type = yyDollar[1].gd_type
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.gd_type = yyDollar[1].gd_type
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.gd_type = yyDollar[1].gd_type
		}
	case 67:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.gd_type = yyDollar[2].gd_type
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			if cT, isCT := yyDollar[1].gd_type.(runtime.GDUnionType); isCT {
				yyVAL.gd_type = runtime.NewGDUnionType(append(cT, yyDollar[3].gd_type)...)
			} else {
				yyVAL.gd_type = runtime.NewGDUnionType(yyDollar[1].gd_type, yyDollar[3].gd_type)
			}
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.gd_type = yyDollar[2].gd_type
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.gd_type = runtime.NewGDTupleType(yyDollar[2].gd_type_list...)
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.gd_type_list = make([]runtime.GDTypable, 0)
		}
	case 73:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.gd_type_list = make([]runtime.GDTypable, 1)
			yyVAL.gd_type_list[0] = yyDollar[1].gd_type
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyDollar[3].gd_type_list = append([]runtime.GDTypable{yyDollar[1].gd_type}, yyDollar[3].gd_type_list...)
			yyVAL.gd_type_list = yyDollar[3].gd_type_list
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.gd_type = runtime.NewGDArrayType(yyDollar[2].gd_type)
		}
	case 76:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			attrTypes := make([]runtime.GDStructAttrType, len(yyDollar[2].gd_type_list))
			for i, attr := range yyDollar[2].gd_type_list {
				attrTypes[i] = attr.(runtime.GDStructAttrType)
			}
			yyVAL.gd_type = runtime.NewGDStructType(attrTypes...)
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyDollar[1].node_list = append(yyDollar[1].node_list, yyDollar[3].node)
			yyVAL.node_list = yyDollar[1].node_list
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.node_list = make([]Node, 1)
			yyVAL.node_list[0] = yyDollar[1].node
		}
	case 80:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.node_list = make([]Node, 0)
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.gd_type = yyDollar[2].gd_type
		}
	case 82:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.gd_type = runtime.GDNilType
		}
	case 83:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.gd_type = buildFuncType(yyDollar[2].node_list, false, yyDollar[4].gd_type)
		}
	case 84:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			yyVAL.gd_type = buildFuncType(yyDollar[2].node_list, true, yyDollar[6].gd_type)
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyDollar[1].gd_type_list = append(yyDollar[1].gd_type_list, yyDollar[3].gd_type)
			yyVAL.gd_type_list = yyDollar[1].gd_type_list
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.gd_type_list = make([]runtime.GDTypable, 1)
			yyVAL.gd_type_list[0] = yyDollar[1].gd_type
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			ident := runtime.GDStringIdentType(yyDollar[1].node.(*NodeIdent).Lit)
			yyVAL.gd_type = runtime.GDStructAttrType{Ident: ident, Type: yyDollar[3].gd_type}
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyDollar[1].gd_type_list = append(yyDollar[1].gd_type_list, yyDollar[3].gd_type)
			yyVAL.gd_type_list = yyDollar[1].gd_type_list
		}
	case 89:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.gd_type_list = make([]runtime.GDTypable, 1)
			yyVAL.gd_type_list[0] = yyDollar[1].gd_type
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.node = NewNodeBlock(yyDollar[2].node_list)
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.node = NewNodeReturn(yyDollar[1].token, nil)
		}
	case 95:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.node = NewNodeReturn(yyDollar[1].token, yyDollar[2].node)
		}
	case 96:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.node = NewNodeBreak(yyDollar[1].token)
		}
	case 98:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.node_list = make([]Node, 0)
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyDollar[1].node_list = append(yyDollar[1].node_list, yyDollar[3].node)
			yyVAL.node_list = yyDollar[1].node_list
		}
	case 100:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.node_list = make([]Node, 1)
			yyVAL.node_list[0] = yyDollar[1].node
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.node = NewNodeLambda(yyDollar[2].gd_type.(*runtime.GDLambdaType), yyDollar[3].node.(*NodeBlock))
		}
	case 102:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.node = NewNodeFunc(true, yyDollar[2].node.(*NodeIdent), yyDollar[3].gd_type.(*runtime.GDLambdaType), yyDollar[4].node.(*NodeBlock))
		}
	case 103:
		yyDollar = yyS[yypt-5 : yypt+1]
		{ // cond ? expr : expr
			yyVAL.node = NewNodeTernaryIf(yyDollar[1].node, yyDollar[3].node, yyDollar[5].node)
		}
	case 104:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.node = NewNodeCastExpr(yyDollar[1].node, yyDollar[3].gd_type)
		}
	case 109:
		yyDollar = yyS[yypt-3 : yypt+1]
		{ // ||
			yyVAL.node = NewNodeExprOperation(runtime.ExprOperationOr, yyDollar[1].node, yyDollar[3].node)
		}
	case 110:
		yyDollar = yyS[yypt-3 : yypt+1]
		{ // &&
			yyVAL.node = NewNodeExprOperation(runtime.ExprOperationAnd, yyDollar[1].node, yyDollar[3].node)
		}
	case 111:
		yyDollar = yyS[yypt-3 : yypt+1]
		{ // ==
			yyVAL.node = NewNodeExprOperation(runtime.ExprOperationEqual, yyDollar[1].node, yyDollar[3].node)
		}
	case 112:
		yyDollar = yyS[yypt-3 : yypt+1]
		{ // !=
			yyVAL.node = NewNodeExprOperation(runtime.ExprOperationNotEqual, yyDollar[1].node, yyDollar[3].node)
		}
	case 113:
		yyDollar = yyS[yypt-3 : yypt+1]
		{ // <
			yyVAL.node = NewNodeExprOperation(runtime.ExprOperationLess, yyDollar[1].node, yyDollar[3].node)
		}
	case 114:
		yyDollar = yyS[yypt-3 : yypt+1]
		{ // >
			yyVAL.node = NewNodeExprOperation(runtime.ExprOperationGreater, yyDollar[1].node, yyDollar[3].node)
		}
	case 115:
		yyDollar = yyS[yypt-3 : yypt+1]
		{ // <=
			yyVAL.node = NewNodeExprOperation(runtime.ExprOperationLessEqual, yyDollar[1].node, yyDollar[3].node)
		}
	case 116:
		yyDollar = yyS[yypt-3 : yypt+1]
		{ // >=
			yyVAL.node = NewNodeExprOperation(runtime.ExprOperationGreaterEqual, yyDollar[1].node, yyDollar[3].node)
		}
	case 117:
		yyDollar = yyS[yypt-3 : yypt+1]
		{ // +
			yyVAL.node = NewNodeExprOperation(runtime.ExprOperationAdd, yyDollar[1].node, yyDollar[3].node)
		}
	case 118:
		yyDollar = yyS[yypt-3 : yypt+1]
		{ // -
			yyVAL.node = NewNodeExprOperation(runtime.ExprOperationSubtract, yyDollar[1].node, yyDollar[3].node)
		}
	case 119:
		yyDollar = yyS[yypt-3 : yypt+1]
		{ // *
			yyVAL.node = NewNodeExprOperation(runtime.ExprOperationMultiply, yyDollar[1].node, yyDollar[3].node)
		}
	case 120:
		yyDollar = yyS[yypt-3 : yypt+1]
		{ // /
			yyVAL.node = NewNodeExprOperation(runtime.ExprOperationQuo, yyDollar[1].node, yyDollar[3].node)
		}
	case 121:
		yyDollar = yyS[yypt-3 : yypt+1]
		{ // %
			yyVAL.node = NewNodeExprOperation(runtime.ExprOperationRem, yyDollar[1].node, yyDollar[3].node)
		}
	case 123:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.node = NewNodeExprOperation(runtime.ExprOperationUnaryPlus, yyDollar[2].node, nil)
		}
	case 124:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.node = NewNodeExprOperation(runtime.ExprOperationUnaryMinus, yyDollar[2].node, nil)
		}
	case 125:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.node = NewNodeExprOperation(runtime.ExprOperationNot, yyDollar[2].node, nil)
		}
	case 127:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.node = yyDollar[2].node
		}
	case 128:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.node = NewNodeMutCollectionOp(MutableCollectionAddOp, yyDollar[1].node, yyDollar[3].node)
		}
	case 129:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.node = NewNodeMutCollectionOp(MutableCollectionRemoveOp, yyDollar[1].node, yyDollar[3].node)
		}
	case 136:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.node = NewNodeEllipsisExpr(yyDollar[1].node)
		}
	case 137:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.node = NewNodeSafeDotExpr(yyDollar[1].node, yyDollar[2].flag, yyDollar[3].node)
		}
	case 138:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.node = NewNodeIterIdxExpr(false, yyDollar[1].node, yyDollar[3].node)
		}
	case 140:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.node = NewNodeCallExpr(yyDollar[1].node, yyDollar[3].node_list)
		}
	case 141:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.flag = false
		}
	case 142:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.flag = true
		}
	case 143:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyDollar[1].node_list = append(yyDollar[1].node_list, yyDollar[3].node)
			yyVAL.node_list = yyDollar[1].node_list
		}
	case 144:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.node_list = make([]Node, 1)
			yyVAL.node_list[0] = yyDollar[1].node
		}
	case 146:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.node_list = make([]Node, 0)
		}
	case 147:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.node = NewNodeLiteral(yyDollar[1].token)
		}
	case 148:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.node = NewNodeLiteral(yyDollar[1].token)
		}
	case 149:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.node = NewNodeLiteral(yyDollar[1].token)
		}
	case 150:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.node = NewNodeLiteral(yyDollar[1].token)
		}
	case 151:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.node = NewNodeLiteral(yyDollar[1].token)
		}
	case 152:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.node = NewNodeLiteral(yyDollar[1].token)
		}
	case 153:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.node = NewNodeLiteral(yyDollar[1].token)
		}
	case 154:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.node = NewNodeLiteral(yyDollar[1].token)
		}
	case 155:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.node = NewNodeIdent(yyDollar[1].token)
		}
	case 156:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.node = NewNodeTuple(yyDollar[2].node_list...)
		}
	case 157:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.node = NewNodeStruct(yyDollar[2].node_list...)
		}
	case 158:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.node = NewNodeStruct()
		}
	case 159:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyDollar[1].node_list = append(yyDollar[1].node_list, yyDollar[3].node)
			yyVAL.node_list = yyDollar[1].node_list
		}
	case 160:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.node_list = make([]Node, 1)
			yyVAL.node_list[0] = yyDollar[1].node
		}
	case 161:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.node = NewNodeStructAttr(yyDollar[1].node.(*NodeIdent), yyDollar[3].node)
		}
	case 162:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.node_list = make([]Node, 0)
		}
	case 163:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.node_list = make([]Node, 1)
			yyVAL.node_list[0] = yyDollar[1].node
		}
	case 164:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyDollar[3].node_list = append([]Node{yyDollar[1].node}, yyDollar[3].node_list...)
			yyVAL.node_list = yyDollar[3].node_list
		}
	case 165:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.node = NewNodeArray(yyDollar[1].token, yyDollar[3].token, yyDollar[2].node_list)
		}
	case 166:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.node = NewNodeForIn(yyDollar[2].node, yyDollar[4].node, yyDollar[5].node.(*NodeBlock))
		}
	case 167:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.node = NewNodeForIf(yyDollar[2].node, yyDollar[4].node_list, yyDollar[5].node.(*NodeBlock))
		}
	case 168:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.node = NewNodeForIf(nil, yyDollar[3].node_list, yyDollar[4].node.(*NodeBlock))
		}
	case 169:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.node = NewNodeForIf(nil, nil, yyDollar[2].node.(*NodeBlock))
		}
	case 170:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			nIf := NewNodeIf(yyDollar[2].node_list, yyDollar[3].node.(*NodeBlock))
			yyVAL.node = NewNodeIfElse(nIf, yyDollar[4].node_list, yyDollar[5].node)
		}
	case 171:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyDollar[1].node_list = append(yyDollar[1].node_list, yyDollar[2].node)
			yyVAL.node_list = yyDollar[1].node_list
		}
	case 172:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.node_list = make([]Node, 0)
		}
	case 173:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.node = NewNodeIf(yyDollar[3].node_list, yyDollar[4].node.(*NodeBlock))
		}
	case 174:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.node = NewNodeIf(nil, yyDollar[2].node.(*NodeBlock))
		}
	case 175:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.node = nil
		}
	}
	goto yystack /* stack new state and value */
}
