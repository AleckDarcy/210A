// Code generated from zz/grammar/ZZ.g4 by ANTLR 4.7.2. DO NOT EDIT.

package grammar // ZZ
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 42, 380,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 3, 2, 3, 2, 5, 2,
	67, 10, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 75, 10, 3, 12, 3,
	14, 3, 78, 11, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6,
	88, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7,
	99, 10, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 107, 10, 7, 12, 7,
	14, 7, 110, 11, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 7, 8, 118, 10, 8,
	12, 8, 14, 8, 121, 11, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 5, 9, 134, 10, 9, 3, 9, 3, 9, 3, 9, 7, 9, 139, 10,
	9, 12, 9, 14, 9, 142, 11, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 155, 10, 12, 12, 12, 14, 12,
	158, 11, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 7, 14, 169, 10, 14, 12, 14, 14, 14, 172, 11, 14, 3, 15, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 196,
	10, 15, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 202, 10, 16, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 5, 16, 209, 10, 16, 3, 16, 3, 16, 7, 16, 213, 10, 16,
	12, 16, 14, 16, 216, 11, 16, 3, 16, 3, 16, 3, 16, 5, 16, 221, 10, 16, 3,
	16, 5, 16, 224, 10, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16,
	232, 10, 16, 3, 17, 3, 17, 5, 17, 236, 10, 17, 3, 17, 3, 17, 3, 17, 3,
	17, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 246, 10, 17, 3, 17, 3, 17, 3, 17,
	3, 17, 3, 17, 3, 17, 5, 17, 254, 10, 17, 3, 18, 3, 18, 5, 18, 258, 10,
	18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 7, 19, 265, 10, 19, 12, 19, 14,
	19, 268, 11, 19, 3, 20, 3, 20, 3, 20, 5, 20, 273, 10, 20, 3, 21, 3, 21,
	3, 21, 3, 21, 3, 21, 3, 21, 7, 21, 281, 10, 21, 12, 21, 14, 21, 284, 11,
	21, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 7, 23,
	295, 10, 23, 12, 23, 14, 23, 298, 11, 23, 3, 24, 3, 24, 3, 24, 5, 24, 303,
	10, 24, 3, 24, 3, 24, 3, 24, 5, 24, 308, 10, 24, 3, 24, 5, 24, 311, 10,
	24, 3, 25, 3, 25, 3, 25, 3, 25, 5, 25, 317, 10, 25, 3, 25, 3, 25, 3, 25,
	5, 25, 322, 10, 25, 3, 25, 5, 25, 325, 10, 25, 3, 26, 3, 26, 3, 26, 3,
	26, 5, 26, 331, 10, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 7, 27,
	339, 10, 27, 12, 27, 14, 27, 342, 11, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3,
	28, 5, 28, 349, 10, 28, 5, 28, 351, 10, 28, 3, 29, 3, 29, 3, 29, 3, 29,
	3, 29, 7, 29, 358, 10, 29, 12, 29, 14, 29, 361, 11, 29, 3, 30, 5, 30, 364,
	10, 30, 3, 31, 3, 31, 3, 31, 5, 31, 369, 10, 31, 3, 31, 3, 31, 3, 32, 3,
	32, 3, 32, 5, 32, 376, 10, 32, 3, 32, 3, 32, 3, 32, 2, 13, 4, 12, 14, 16,
	22, 26, 36, 40, 44, 52, 56, 33, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
	24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58,
	60, 62, 2, 8, 4, 2, 4, 4, 34, 35, 3, 2, 7, 8, 3, 2, 9, 10, 3, 2, 13, 18,
	4, 2, 13, 13, 18, 20, 3, 2, 39, 40, 2, 399, 2, 66, 3, 2, 2, 2, 4, 68, 3,
	2, 2, 2, 6, 79, 3, 2, 2, 2, 8, 81, 3, 2, 2, 2, 10, 87, 3, 2, 2, 2, 12,
	98, 3, 2, 2, 2, 14, 111, 3, 2, 2, 2, 16, 133, 3, 2, 2, 2, 18, 143, 3, 2,
	2, 2, 20, 145, 3, 2, 2, 2, 22, 149, 3, 2, 2, 2, 24, 159, 3, 2, 2, 2, 26,
	162, 3, 2, 2, 2, 28, 195, 3, 2, 2, 2, 30, 231, 3, 2, 2, 2, 32, 253, 3,
	2, 2, 2, 34, 257, 3, 2, 2, 2, 36, 259, 3, 2, 2, 2, 38, 272, 3, 2, 2, 2,
	40, 274, 3, 2, 2, 2, 42, 285, 3, 2, 2, 2, 44, 288, 3, 2, 2, 2, 46, 299,
	3, 2, 2, 2, 48, 312, 3, 2, 2, 2, 50, 330, 3, 2, 2, 2, 52, 332, 3, 2, 2,
	2, 54, 350, 3, 2, 2, 2, 56, 352, 3, 2, 2, 2, 58, 363, 3, 2, 2, 2, 60, 365,
	3, 2, 2, 2, 62, 372, 3, 2, 2, 2, 64, 67, 7, 39, 2, 2, 65, 67, 5, 24, 13,
	2, 66, 64, 3, 2, 2, 2, 66, 65, 3, 2, 2, 2, 67, 3, 3, 2, 2, 2, 68, 69, 8,
	3, 1, 2, 69, 70, 5, 2, 2, 2, 70, 76, 3, 2, 2, 2, 71, 72, 12, 3, 2, 2, 72,
	73, 7, 3, 2, 2, 73, 75, 5, 2, 2, 2, 74, 71, 3, 2, 2, 2, 75, 78, 3, 2, 2,
	2, 76, 74, 3, 2, 2, 2, 76, 77, 3, 2, 2, 2, 77, 5, 3, 2, 2, 2, 78, 76, 3,
	2, 2, 2, 79, 80, 9, 2, 2, 2, 80, 7, 3, 2, 2, 2, 81, 82, 5, 6, 4, 2, 82,
	9, 3, 2, 2, 2, 83, 88, 5, 8, 5, 2, 84, 85, 7, 5, 2, 2, 85, 86, 7, 6, 2,
	2, 86, 88, 5, 10, 6, 2, 87, 83, 3, 2, 2, 2, 87, 84, 3, 2, 2, 2, 88, 11,
	3, 2, 2, 2, 89, 90, 8, 7, 1, 2, 90, 99, 7, 40, 2, 2, 91, 99, 7, 41, 2,
	2, 92, 99, 7, 39, 2, 2, 93, 99, 5, 24, 13, 2, 94, 95, 7, 11, 2, 2, 95,
	96, 5, 12, 7, 2, 96, 97, 7, 12, 2, 2, 97, 99, 3, 2, 2, 2, 98, 89, 3, 2,
	2, 2, 98, 91, 3, 2, 2, 2, 98, 92, 3, 2, 2, 2, 98, 93, 3, 2, 2, 2, 98, 94,
	3, 2, 2, 2, 99, 108, 3, 2, 2, 2, 100, 101, 12, 5, 2, 2, 101, 102, 9, 3,
	2, 2, 102, 107, 5, 12, 7, 6, 103, 104, 12, 4, 2, 2, 104, 105, 9, 4, 2,
	2, 105, 107, 5, 12, 7, 5, 106, 100, 3, 2, 2, 2, 106, 103, 3, 2, 2, 2, 107,
	110, 3, 2, 2, 2, 108, 106, 3, 2, 2, 2, 108, 109, 3, 2, 2, 2, 109, 13, 3,
	2, 2, 2, 110, 108, 3, 2, 2, 2, 111, 112, 8, 8, 1, 2, 112, 113, 5, 12, 7,
	2, 113, 119, 3, 2, 2, 2, 114, 115, 12, 3, 2, 2, 115, 116, 7, 3, 2, 2, 116,
	118, 5, 12, 7, 2, 117, 114, 3, 2, 2, 2, 118, 121, 3, 2, 2, 2, 119, 117,
	3, 2, 2, 2, 119, 120, 3, 2, 2, 2, 120, 15, 3, 2, 2, 2, 121, 119, 3, 2,
	2, 2, 122, 123, 8, 9, 1, 2, 123, 124, 5, 12, 7, 2, 124, 125, 9, 5, 2, 2,
	125, 126, 5, 12, 7, 2, 126, 134, 3, 2, 2, 2, 127, 128, 7, 21, 2, 2, 128,
	134, 5, 16, 9, 4, 129, 130, 7, 11, 2, 2, 130, 131, 5, 16, 9, 2, 131, 132,
	7, 12, 2, 2, 132, 134, 3, 2, 2, 2, 133, 122, 3, 2, 2, 2, 133, 127, 3, 2,
	2, 2, 133, 129, 3, 2, 2, 2, 134, 140, 3, 2, 2, 2, 135, 136, 12, 5, 2, 2,
	136, 137, 9, 6, 2, 2, 137, 139, 5, 16, 9, 6, 138, 135, 3, 2, 2, 2, 139,
	142, 3, 2, 2, 2, 140, 138, 3, 2, 2, 2, 140, 141, 3, 2, 2, 2, 141, 17, 3,
	2, 2, 2, 142, 140, 3, 2, 2, 2, 143, 144, 9, 7, 2, 2, 144, 19, 3, 2, 2,
	2, 145, 146, 7, 5, 2, 2, 146, 147, 7, 40, 2, 2, 147, 148, 7, 6, 2, 2, 148,
	21, 3, 2, 2, 2, 149, 150, 8, 12, 1, 2, 150, 151, 5, 20, 11, 2, 151, 156,
	3, 2, 2, 2, 152, 153, 12, 3, 2, 2, 153, 155, 5, 20, 11, 2, 154, 152, 3,
	2, 2, 2, 155, 158, 3, 2, 2, 2, 156, 154, 3, 2, 2, 2, 156, 157, 3, 2, 2,
	2, 157, 23, 3, 2, 2, 2, 158, 156, 3, 2, 2, 2, 159, 160, 7, 39, 2, 2, 160,
	161, 5, 22, 12, 2, 161, 25, 3, 2, 2, 2, 162, 163, 8, 14, 1, 2, 163, 164,
	5, 18, 10, 2, 164, 170, 3, 2, 2, 2, 165, 166, 12, 3, 2, 2, 166, 167, 7,
	3, 2, 2, 167, 169, 5, 18, 10, 2, 168, 165, 3, 2, 2, 2, 169, 172, 3, 2,
	2, 2, 170, 168, 3, 2, 2, 2, 170, 171, 3, 2, 2, 2, 171, 27, 3, 2, 2, 2,
	172, 170, 3, 2, 2, 2, 173, 174, 5, 4, 3, 2, 174, 175, 5, 6, 4, 2, 175,
	196, 3, 2, 2, 2, 176, 177, 5, 4, 3, 2, 177, 178, 7, 22, 2, 2, 178, 179,
	5, 14, 8, 2, 179, 196, 3, 2, 2, 2, 180, 181, 5, 2, 2, 2, 181, 182, 7, 22,
	2, 2, 182, 183, 7, 37, 2, 2, 183, 184, 7, 11, 2, 2, 184, 185, 5, 10, 6,
	2, 185, 186, 7, 3, 2, 2, 186, 187, 7, 11, 2, 2, 187, 188, 5, 26, 14, 2,
	188, 189, 7, 12, 2, 2, 189, 190, 7, 12, 2, 2, 190, 196, 3, 2, 2, 2, 191,
	192, 5, 2, 2, 2, 192, 193, 7, 22, 2, 2, 193, 194, 5, 60, 31, 2, 194, 196,
	3, 2, 2, 2, 195, 173, 3, 2, 2, 2, 195, 176, 3, 2, 2, 2, 195, 180, 3, 2,
	2, 2, 195, 191, 3, 2, 2, 2, 196, 29, 3, 2, 2, 2, 197, 198, 7, 23, 2, 2,
	198, 199, 5, 16, 9, 2, 199, 201, 7, 24, 2, 2, 200, 202, 5, 56, 29, 2, 201,
	200, 3, 2, 2, 2, 201, 202, 3, 2, 2, 2, 202, 203, 3, 2, 2, 2, 203, 214,
	7, 25, 2, 2, 204, 205, 7, 26, 2, 2, 205, 206, 5, 16, 9, 2, 206, 208, 7,
	24, 2, 2, 207, 209, 5, 56, 29, 2, 208, 207, 3, 2, 2, 2, 208, 209, 3, 2,
	2, 2, 209, 210, 3, 2, 2, 2, 210, 211, 7, 25, 2, 2, 211, 213, 3, 2, 2, 2,
	212, 204, 3, 2, 2, 2, 213, 216, 3, 2, 2, 2, 214, 212, 3, 2, 2, 2, 214,
	215, 3, 2, 2, 2, 215, 223, 3, 2, 2, 2, 216, 214, 3, 2, 2, 2, 217, 218,
	7, 27, 2, 2, 218, 220, 7, 24, 2, 2, 219, 221, 5, 56, 29, 2, 220, 219, 3,
	2, 2, 2, 220, 221, 3, 2, 2, 2, 221, 222, 3, 2, 2, 2, 222, 224, 7, 25, 2,
	2, 223, 217, 3, 2, 2, 2, 223, 224, 3, 2, 2, 2, 224, 232, 3, 2, 2, 2, 225,
	226, 5, 16, 9, 2, 226, 227, 7, 28, 2, 2, 227, 228, 5, 54, 28, 2, 228, 229,
	7, 29, 2, 2, 229, 230, 5, 54, 28, 2, 230, 232, 3, 2, 2, 2, 231, 197, 3,
	2, 2, 2, 231, 225, 3, 2, 2, 2, 232, 31, 3, 2, 2, 2, 233, 235, 7, 30, 2,
	2, 234, 236, 5, 16, 9, 2, 235, 234, 3, 2, 2, 2, 235, 236, 3, 2, 2, 2, 236,
	237, 3, 2, 2, 2, 237, 238, 7, 24, 2, 2, 238, 239, 5, 56, 29, 2, 239, 240,
	7, 25, 2, 2, 240, 254, 3, 2, 2, 2, 241, 242, 7, 30, 2, 2, 242, 243, 5,
	28, 15, 2, 243, 245, 7, 31, 2, 2, 244, 246, 5, 16, 9, 2, 245, 244, 3, 2,
	2, 2, 245, 246, 3, 2, 2, 2, 246, 247, 3, 2, 2, 2, 247, 248, 7, 31, 2, 2,
	248, 249, 5, 28, 15, 2, 249, 250, 7, 24, 2, 2, 250, 251, 5, 56, 29, 2,
	251, 252, 7, 25, 2, 2, 252, 254, 3, 2, 2, 2, 253, 233, 3, 2, 2, 2, 253,
	241, 3, 2, 2, 2, 254, 33, 3, 2, 2, 2, 255, 258, 5, 28, 15, 2, 256, 258,
	5, 62, 32, 2, 257, 255, 3, 2, 2, 2, 257, 256, 3, 2, 2, 2, 258, 35, 3, 2,
	2, 2, 259, 260, 8, 19, 1, 2, 260, 261, 5, 34, 18, 2, 261, 266, 3, 2, 2,
	2, 262, 263, 12, 3, 2, 2, 263, 265, 5, 34, 18, 2, 264, 262, 3, 2, 2, 2,
	265, 268, 3, 2, 2, 2, 266, 264, 3, 2, 2, 2, 266, 267, 3, 2, 2, 2, 267,
	37, 3, 2, 2, 2, 268, 266, 3, 2, 2, 2, 269, 273, 5, 6, 4, 2, 270, 273, 5,
	10, 6, 2, 271, 273, 5, 46, 24, 2, 272, 269, 3, 2, 2, 2, 272, 270, 3, 2,
	2, 2, 272, 271, 3, 2, 2, 2, 273, 39, 3, 2, 2, 2, 274, 275, 8, 21, 1, 2,
	275, 276, 5, 38, 20, 2, 276, 282, 3, 2, 2, 2, 277, 278, 12, 3, 2, 2, 278,
	279, 7, 3, 2, 2, 279, 281, 5, 38, 20, 2, 280, 277, 3, 2, 2, 2, 281, 284,
	3, 2, 2, 2, 282, 280, 3, 2, 2, 2, 282, 283, 3, 2, 2, 2, 283, 41, 3, 2,
	2, 2, 284, 282, 3, 2, 2, 2, 285, 286, 5, 4, 3, 2, 286, 287, 5, 38, 20,
	2, 287, 43, 3, 2, 2, 2, 288, 289, 8, 23, 1, 2, 289, 290, 5, 42, 22, 2,
	290, 296, 3, 2, 2, 2, 291, 292, 12, 3, 2, 2, 292, 293, 7, 3, 2, 2, 293,
	295, 5, 42, 22, 2, 294, 291, 3, 2, 2, 2, 295, 298, 3, 2, 2, 2, 296, 294,
	3, 2, 2, 2, 296, 297, 3, 2, 2, 2, 297, 45, 3, 2, 2, 2, 298, 296, 3, 2,
	2, 2, 299, 300, 7, 36, 2, 2, 300, 302, 7, 11, 2, 2, 301, 303, 5, 44, 23,
	2, 302, 301, 3, 2, 2, 2, 302, 303, 3, 2, 2, 2, 303, 304, 3, 2, 2, 2, 304,
	310, 7, 12, 2, 2, 305, 307, 7, 11, 2, 2, 306, 308, 5, 40, 21, 2, 307, 306,
	3, 2, 2, 2, 307, 308, 3, 2, 2, 2, 308, 309, 3, 2, 2, 2, 309, 311, 7, 12,
	2, 2, 310, 305, 3, 2, 2, 2, 310, 311, 3, 2, 2, 2, 311, 47, 3, 2, 2, 2,
	312, 313, 7, 36, 2, 2, 313, 314, 7, 39, 2, 2, 314, 316, 7, 11, 2, 2, 315,
	317, 5, 44, 23, 2, 316, 315, 3, 2, 2, 2, 316, 317, 3, 2, 2, 2, 317, 318,
	3, 2, 2, 2, 318, 324, 7, 12, 2, 2, 319, 321, 7, 11, 2, 2, 320, 322, 5,
	40, 21, 2, 321, 320, 3, 2, 2, 2, 321, 322, 3, 2, 2, 2, 322, 323, 3, 2,
	2, 2, 323, 325, 7, 12, 2, 2, 324, 319, 3, 2, 2, 2, 324, 325, 3, 2, 2, 2,
	325, 49, 3, 2, 2, 2, 326, 331, 5, 12, 7, 2, 327, 331, 5, 16, 9, 2, 328,
	331, 7, 39, 2, 2, 329, 331, 7, 38, 2, 2, 330, 326, 3, 2, 2, 2, 330, 327,
	3, 2, 2, 2, 330, 328, 3, 2, 2, 2, 330, 329, 3, 2, 2, 2, 331, 51, 3, 2,
	2, 2, 332, 333, 8, 27, 1, 2, 333, 334, 5, 50, 26, 2, 334, 340, 3, 2, 2,
	2, 335, 336, 12, 3, 2, 2, 336, 337, 7, 3, 2, 2, 337, 339, 5, 50, 26, 2,
	338, 335, 3, 2, 2, 2, 339, 342, 3, 2, 2, 2, 340, 338, 3, 2, 2, 2, 340,
	341, 3, 2, 2, 2, 341, 53, 3, 2, 2, 2, 342, 340, 3, 2, 2, 2, 343, 351, 5,
	28, 15, 2, 344, 351, 5, 30, 16, 2, 345, 351, 5, 32, 17, 2, 346, 348, 7,
	32, 2, 2, 347, 349, 5, 52, 27, 2, 348, 347, 3, 2, 2, 2, 348, 349, 3, 2,
	2, 2, 349, 351, 3, 2, 2, 2, 350, 343, 3, 2, 2, 2, 350, 344, 3, 2, 2, 2,
	350, 345, 3, 2, 2, 2, 350, 346, 3, 2, 2, 2, 351, 55, 3, 2, 2, 2, 352, 353,
	8, 29, 1, 2, 353, 354, 5, 54, 28, 2, 354, 359, 3, 2, 2, 2, 355, 356, 12,
	3, 2, 2, 356, 358, 5, 54, 28, 2, 357, 355, 3, 2, 2, 2, 358, 361, 3, 2,
	2, 2, 359, 357, 3, 2, 2, 2, 359, 360, 3, 2, 2, 2, 360, 57, 3, 2, 2, 2,
	361, 359, 3, 2, 2, 2, 362, 364, 5, 56, 29, 2, 363, 362, 3, 2, 2, 2, 363,
	364, 3, 2, 2, 2, 364, 59, 3, 2, 2, 2, 365, 366, 5, 46, 24, 2, 366, 368,
	7, 24, 2, 2, 367, 369, 5, 58, 30, 2, 368, 367, 3, 2, 2, 2, 368, 369, 3,
	2, 2, 2, 369, 370, 3, 2, 2, 2, 370, 371, 7, 25, 2, 2, 371, 61, 3, 2, 2,
	2, 372, 373, 5, 48, 25, 2, 373, 375, 7, 24, 2, 2, 374, 376, 5, 58, 30,
	2, 375, 374, 3, 2, 2, 2, 375, 376, 3, 2, 2, 2, 376, 377, 3, 2, 2, 2, 377,
	378, 7, 25, 2, 2, 378, 63, 3, 2, 2, 2, 42, 66, 76, 87, 98, 106, 108, 119,
	133, 140, 156, 170, 195, 201, 208, 214, 220, 223, 231, 235, 245, 253, 257,
	266, 272, 282, 296, 302, 307, 310, 316, 321, 324, 330, 340, 348, 350, 359,
	363, 368, 375,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "','", "'string'", "'['", "']'", "'*'", "'/'", "'+'", "'-'", "'('",
	"')'", "'=='", "'<'", "'>'", "'>='", "'<='", "'!='", "'&&'", "'||'", "'!'",
	"'='", "'if'", "'{'", "'}'", "'elsif'", "'else'", "'?'", "':'", "'for'",
	"';'", "'return'", "'var'", "'int'", "'float'", "'func'", "'list'", "'nil'",
	"", "", "'0.0'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "Var", "Int", "Float",
	"Func", "List", "Nil", "Identifier", "IntegerLiteral", "FloatLiteral",
	"WS",
}

var ruleNames = []string{
	"declarator", "declaratorList", "simpleTypeSpecifier", "listElementTypeSpecifier",
	"listTypeSpecifier", "aExp", "aExpList", "bExp", "integerExpression", "listElement",
	"listElements", "listElementExpression", "tupleSizes", "assignStatement",
	"selectionStatement", "iterationStatement", "entry", "entryList", "typeSpecifier",
	"typeSpecifierList", "typeSpecifierWithIdentity", "typeSpecifierWithIdentityList",
	"funcSpecifier", "funcSpecifierWithName", "funcReturnPara", "funcReturnParaList",
	"funcStatement", "funcStatementList", "funcBody", "funcExpression", "funcExpressionWithName",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type ZZParser struct {
	*antlr.BaseParser
}

func NewZZParser(input antlr.TokenStream) *ZZParser {
	this := new(ZZParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "ZZ.g4"

	return this
}

// ZZParser tokens.
const (
	ZZParserEOF            = antlr.TokenEOF
	ZZParserT__0           = 1
	ZZParserT__1           = 2
	ZZParserT__2           = 3
	ZZParserT__3           = 4
	ZZParserT__4           = 5
	ZZParserT__5           = 6
	ZZParserT__6           = 7
	ZZParserT__7           = 8
	ZZParserT__8           = 9
	ZZParserT__9           = 10
	ZZParserT__10          = 11
	ZZParserT__11          = 12
	ZZParserT__12          = 13
	ZZParserT__13          = 14
	ZZParserT__14          = 15
	ZZParserT__15          = 16
	ZZParserT__16          = 17
	ZZParserT__17          = 18
	ZZParserT__18          = 19
	ZZParserT__19          = 20
	ZZParserT__20          = 21
	ZZParserT__21          = 22
	ZZParserT__22          = 23
	ZZParserT__23          = 24
	ZZParserT__24          = 25
	ZZParserT__25          = 26
	ZZParserT__26          = 27
	ZZParserT__27          = 28
	ZZParserT__28          = 29
	ZZParserT__29          = 30
	ZZParserVar            = 31
	ZZParserInt            = 32
	ZZParserFloat          = 33
	ZZParserFunc           = 34
	ZZParserList           = 35
	ZZParserNil            = 36
	ZZParserIdentifier     = 37
	ZZParserIntegerLiteral = 38
	ZZParserFloatLiteral   = 39
	ZZParserWS             = 40
)

// ZZParser rules.
const (
	ZZParserRULE_declarator                    = 0
	ZZParserRULE_declaratorList                = 1
	ZZParserRULE_simpleTypeSpecifier           = 2
	ZZParserRULE_listElementTypeSpecifier      = 3
	ZZParserRULE_listTypeSpecifier             = 4
	ZZParserRULE_aExp                          = 5
	ZZParserRULE_aExpList                      = 6
	ZZParserRULE_bExp                          = 7
	ZZParserRULE_integerExpression             = 8
	ZZParserRULE_listElement                   = 9
	ZZParserRULE_listElements                  = 10
	ZZParserRULE_listElementExpression         = 11
	ZZParserRULE_tupleSizes                    = 12
	ZZParserRULE_assignStatement               = 13
	ZZParserRULE_selectionStatement            = 14
	ZZParserRULE_iterationStatement            = 15
	ZZParserRULE_entry                         = 16
	ZZParserRULE_entryList                     = 17
	ZZParserRULE_typeSpecifier                 = 18
	ZZParserRULE_typeSpecifierList             = 19
	ZZParserRULE_typeSpecifierWithIdentity     = 20
	ZZParserRULE_typeSpecifierWithIdentityList = 21
	ZZParserRULE_funcSpecifier                 = 22
	ZZParserRULE_funcSpecifierWithName         = 23
	ZZParserRULE_funcReturnPara                = 24
	ZZParserRULE_funcReturnParaList            = 25
	ZZParserRULE_funcStatement                 = 26
	ZZParserRULE_funcStatementList             = 27
	ZZParserRULE_funcBody                      = 28
	ZZParserRULE_funcExpression                = 29
	ZZParserRULE_funcExpressionWithName        = 30
)

// IDeclaratorContext is an interface to support dynamic dispatch.
type IDeclaratorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclaratorContext differentiates from other interfaces.
	IsDeclaratorContext()
}

type DeclaratorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclaratorContext() *DeclaratorContext {
	var p = new(DeclaratorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_declarator
	return p
}

func (*DeclaratorContext) IsDeclaratorContext() {}

func NewDeclaratorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclaratorContext {
	var p = new(DeclaratorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_declarator

	return p
}

func (s *DeclaratorContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclaratorContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ZZParserIdentifier, 0)
}

func (s *DeclaratorContext) ListElementExpression() IListElementExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListElementExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListElementExpressionContext)
}

func (s *DeclaratorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclaratorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclaratorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterDeclarator(s)
	}
}

func (s *DeclaratorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitDeclarator(s)
	}
}

func (p *ZZParser) Declarator() (localctx IDeclaratorContext) {
	localctx = NewDeclaratorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ZZParserRULE_declarator)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(62)
			p.Match(ZZParserIdentifier)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(63)
			p.ListElementExpression()
		}

	}

	return localctx
}

// IDeclaratorListContext is an interface to support dynamic dispatch.
type IDeclaratorListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclaratorListContext differentiates from other interfaces.
	IsDeclaratorListContext()
}

type DeclaratorListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclaratorListContext() *DeclaratorListContext {
	var p = new(DeclaratorListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_declaratorList
	return p
}

func (*DeclaratorListContext) IsDeclaratorListContext() {}

func NewDeclaratorListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclaratorListContext {
	var p = new(DeclaratorListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_declaratorList

	return p
}

func (s *DeclaratorListContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclaratorListContext) Declarator() IDeclaratorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclaratorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclaratorContext)
}

func (s *DeclaratorListContext) DeclaratorList() IDeclaratorListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclaratorListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclaratorListContext)
}

func (s *DeclaratorListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclaratorListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclaratorListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterDeclaratorList(s)
	}
}

func (s *DeclaratorListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitDeclaratorList(s)
	}
}

func (p *ZZParser) DeclaratorList() (localctx IDeclaratorListContext) {
	return p.declaratorList(0)
}

func (p *ZZParser) declaratorList(_p int) (localctx IDeclaratorListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewDeclaratorListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IDeclaratorListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, ZZParserRULE_declaratorList, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)
		p.Declarator()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewDeclaratorListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ZZParserRULE_declaratorList)
			p.SetState(69)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(70)
				p.Match(ZZParserT__0)
			}
			{
				p.SetState(71)
				p.Declarator()
			}

		}
		p.SetState(76)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}

	return localctx
}

// ISimpleTypeSpecifierContext is an interface to support dynamic dispatch.
type ISimpleTypeSpecifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSimpleTypeSpecifierContext differentiates from other interfaces.
	IsSimpleTypeSpecifierContext()
}

type SimpleTypeSpecifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpleTypeSpecifierContext() *SimpleTypeSpecifierContext {
	var p = new(SimpleTypeSpecifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_simpleTypeSpecifier
	return p
}

func (*SimpleTypeSpecifierContext) IsSimpleTypeSpecifierContext() {}

func NewSimpleTypeSpecifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleTypeSpecifierContext {
	var p = new(SimpleTypeSpecifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_simpleTypeSpecifier

	return p
}

func (s *SimpleTypeSpecifierContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleTypeSpecifierContext) Int() antlr.TerminalNode {
	return s.GetToken(ZZParserInt, 0)
}

func (s *SimpleTypeSpecifierContext) Float() antlr.TerminalNode {
	return s.GetToken(ZZParserFloat, 0)
}

func (s *SimpleTypeSpecifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleTypeSpecifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimpleTypeSpecifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterSimpleTypeSpecifier(s)
	}
}

func (s *SimpleTypeSpecifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitSimpleTypeSpecifier(s)
	}
}

func (p *ZZParser) SimpleTypeSpecifier() (localctx ISimpleTypeSpecifierContext) {
	localctx = NewSimpleTypeSpecifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ZZParserRULE_simpleTypeSpecifier)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-2)&-(0x1f+1)) == 0 && ((1<<uint((_la-2)))&((1<<(ZZParserT__1-2))|(1<<(ZZParserInt-2))|(1<<(ZZParserFloat-2)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IListElementTypeSpecifierContext is an interface to support dynamic dispatch.
type IListElementTypeSpecifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListElementTypeSpecifierContext differentiates from other interfaces.
	IsListElementTypeSpecifierContext()
}

type ListElementTypeSpecifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListElementTypeSpecifierContext() *ListElementTypeSpecifierContext {
	var p = new(ListElementTypeSpecifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_listElementTypeSpecifier
	return p
}

func (*ListElementTypeSpecifierContext) IsListElementTypeSpecifierContext() {}

func NewListElementTypeSpecifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListElementTypeSpecifierContext {
	var p = new(ListElementTypeSpecifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_listElementTypeSpecifier

	return p
}

func (s *ListElementTypeSpecifierContext) GetParser() antlr.Parser { return s.parser }

func (s *ListElementTypeSpecifierContext) SimpleTypeSpecifier() ISimpleTypeSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimpleTypeSpecifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISimpleTypeSpecifierContext)
}

func (s *ListElementTypeSpecifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListElementTypeSpecifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListElementTypeSpecifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterListElementTypeSpecifier(s)
	}
}

func (s *ListElementTypeSpecifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitListElementTypeSpecifier(s)
	}
}

func (p *ZZParser) ListElementTypeSpecifier() (localctx IListElementTypeSpecifierContext) {
	localctx = NewListElementTypeSpecifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ZZParserRULE_listElementTypeSpecifier)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		p.SimpleTypeSpecifier()
	}

	return localctx
}

// IListTypeSpecifierContext is an interface to support dynamic dispatch.
type IListTypeSpecifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListTypeSpecifierContext differentiates from other interfaces.
	IsListTypeSpecifierContext()
}

type ListTypeSpecifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListTypeSpecifierContext() *ListTypeSpecifierContext {
	var p = new(ListTypeSpecifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_listTypeSpecifier
	return p
}

func (*ListTypeSpecifierContext) IsListTypeSpecifierContext() {}

func NewListTypeSpecifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListTypeSpecifierContext {
	var p = new(ListTypeSpecifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_listTypeSpecifier

	return p
}

func (s *ListTypeSpecifierContext) GetParser() antlr.Parser { return s.parser }

func (s *ListTypeSpecifierContext) ListElementTypeSpecifier() IListElementTypeSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListElementTypeSpecifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListElementTypeSpecifierContext)
}

func (s *ListTypeSpecifierContext) ListTypeSpecifier() IListTypeSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListTypeSpecifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListTypeSpecifierContext)
}

func (s *ListTypeSpecifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListTypeSpecifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListTypeSpecifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterListTypeSpecifier(s)
	}
}

func (s *ListTypeSpecifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitListTypeSpecifier(s)
	}
}

func (p *ZZParser) ListTypeSpecifier() (localctx IListTypeSpecifierContext) {
	localctx = NewListTypeSpecifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ZZParserRULE_listTypeSpecifier)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(85)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ZZParserT__1, ZZParserInt, ZZParserFloat:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(81)
			p.ListElementTypeSpecifier()
		}

	case ZZParserT__2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(82)
			p.Match(ZZParserT__2)
		}
		{
			p.SetState(83)
			p.Match(ZZParserT__3)
		}
		{
			p.SetState(84)
			p.ListTypeSpecifier()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAExpContext is an interface to support dynamic dispatch.
type IAExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAExpContext differentiates from other interfaces.
	IsAExpContext()
}

type AExpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAExpContext() *AExpContext {
	var p = new(AExpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_aExp
	return p
}

func (*AExpContext) IsAExpContext() {}

func NewAExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AExpContext {
	var p = new(AExpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_aExp

	return p
}

func (s *AExpContext) GetParser() antlr.Parser { return s.parser }

func (s *AExpContext) IntegerLiteral() antlr.TerminalNode {
	return s.GetToken(ZZParserIntegerLiteral, 0)
}

func (s *AExpContext) FloatLiteral() antlr.TerminalNode {
	return s.GetToken(ZZParserFloatLiteral, 0)
}

func (s *AExpContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ZZParserIdentifier, 0)
}

func (s *AExpContext) ListElementExpression() IListElementExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListElementExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListElementExpressionContext)
}

func (s *AExpContext) AllAExp() []IAExpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAExpContext)(nil)).Elem())
	var tst = make([]IAExpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAExpContext)
		}
	}

	return tst
}

func (s *AExpContext) AExp(i int) IAExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAExpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAExpContext)
}

func (s *AExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterAExp(s)
	}
}

func (s *AExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitAExp(s)
	}
}

func (p *ZZParser) AExp() (localctx IAExpContext) {
	return p.aExp(0)
}

func (p *ZZParser) aExp(_p int) (localctx IAExpContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewAExpContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IAExpContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, ZZParserRULE_aExp, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(88)
			p.Match(ZZParserIntegerLiteral)
		}

	case 2:
		{
			p.SetState(89)
			p.Match(ZZParserFloatLiteral)
		}

	case 3:
		{
			p.SetState(90)
			p.Match(ZZParserIdentifier)
		}

	case 4:
		{
			p.SetState(91)
			p.ListElementExpression()
		}

	case 5:
		{
			p.SetState(92)
			p.Match(ZZParserT__8)
		}
		{
			p.SetState(93)
			p.aExp(0)
		}
		{
			p.SetState(94)
			p.Match(ZZParserT__9)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(104)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ZZParserRULE_aExp)
				p.SetState(98)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(99)
					_la = p.GetTokenStream().LA(1)

					if !(_la == ZZParserT__4 || _la == ZZParserT__5) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(100)
					p.aExp(4)
				}

			case 2:
				localctx = NewAExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ZZParserRULE_aExp)
				p.SetState(101)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(102)
					_la = p.GetTokenStream().LA(1)

					if !(_la == ZZParserT__6 || _la == ZZParserT__7) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(103)
					p.aExp(3)
				}

			}

		}
		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// IAExpListContext is an interface to support dynamic dispatch.
type IAExpListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAExpListContext differentiates from other interfaces.
	IsAExpListContext()
}

type AExpListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAExpListContext() *AExpListContext {
	var p = new(AExpListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_aExpList
	return p
}

func (*AExpListContext) IsAExpListContext() {}

func NewAExpListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AExpListContext {
	var p = new(AExpListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_aExpList

	return p
}

func (s *AExpListContext) GetParser() antlr.Parser { return s.parser }

func (s *AExpListContext) AExp() IAExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAExpContext)
}

func (s *AExpListContext) AExpList() IAExpListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAExpListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAExpListContext)
}

func (s *AExpListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AExpListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AExpListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterAExpList(s)
	}
}

func (s *AExpListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitAExpList(s)
	}
}

func (p *ZZParser) AExpList() (localctx IAExpListContext) {
	return p.aExpList(0)
}

func (p *ZZParser) aExpList(_p int) (localctx IAExpListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewAExpListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IAExpListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 12
	p.EnterRecursionRule(localctx, 12, ZZParserRULE_aExpList, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(110)
		p.aExp(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewAExpListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ZZParserRULE_aExpList)
			p.SetState(112)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(113)
				p.Match(ZZParserT__0)
			}
			{
				p.SetState(114)
				p.aExp(0)
			}

		}
		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}

	return localctx
}

// IBExpContext is an interface to support dynamic dispatch.
type IBExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBExpContext differentiates from other interfaces.
	IsBExpContext()
}

type BExpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBExpContext() *BExpContext {
	var p = new(BExpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_bExp
	return p
}

func (*BExpContext) IsBExpContext() {}

func NewBExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BExpContext {
	var p = new(BExpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_bExp

	return p
}

func (s *BExpContext) GetParser() antlr.Parser { return s.parser }

func (s *BExpContext) AllAExp() []IAExpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAExpContext)(nil)).Elem())
	var tst = make([]IAExpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAExpContext)
		}
	}

	return tst
}

func (s *BExpContext) AExp(i int) IAExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAExpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAExpContext)
}

func (s *BExpContext) AllBExp() []IBExpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBExpContext)(nil)).Elem())
	var tst = make([]IBExpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBExpContext)
		}
	}

	return tst
}

func (s *BExpContext) BExp(i int) IBExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBExpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBExpContext)
}

func (s *BExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterBExp(s)
	}
}

func (s *BExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitBExp(s)
	}
}

func (p *ZZParser) BExp() (localctx IBExpContext) {
	return p.bExp(0)
}

func (p *ZZParser) bExp(_p int) (localctx IBExpContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewBExpContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IBExpContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 14
	p.EnterRecursionRule(localctx, 14, ZZParserRULE_bExp, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(121)
			p.aExp(0)
		}
		{
			p.SetState(122)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ZZParserT__10)|(1<<ZZParserT__11)|(1<<ZZParserT__12)|(1<<ZZParserT__13)|(1<<ZZParserT__14)|(1<<ZZParserT__15))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(123)
			p.aExp(0)
		}

	case 2:
		{
			p.SetState(125)
			p.Match(ZZParserT__18)
		}
		{
			p.SetState(126)
			p.bExp(2)
		}

	case 3:
		{
			p.SetState(127)
			p.Match(ZZParserT__8)
		}
		{
			p.SetState(128)
			p.bExp(0)
		}
		{
			p.SetState(129)
			p.Match(ZZParserT__9)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewBExpContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ZZParserRULE_bExp)
			p.SetState(133)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
			}
			{
				p.SetState(134)
				_la = p.GetTokenStream().LA(1)

				if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ZZParserT__10)|(1<<ZZParserT__15)|(1<<ZZParserT__16)|(1<<ZZParserT__17))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(135)
				p.bExp(4)
			}

		}
		p.SetState(140)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}

	return localctx
}

// IIntegerExpressionContext is an interface to support dynamic dispatch.
type IIntegerExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntegerExpressionContext differentiates from other interfaces.
	IsIntegerExpressionContext()
}

type IntegerExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerExpressionContext() *IntegerExpressionContext {
	var p = new(IntegerExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_integerExpression
	return p
}

func (*IntegerExpressionContext) IsIntegerExpressionContext() {}

func NewIntegerExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerExpressionContext {
	var p = new(IntegerExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_integerExpression

	return p
}

func (s *IntegerExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerExpressionContext) IntegerLiteral() antlr.TerminalNode {
	return s.GetToken(ZZParserIntegerLiteral, 0)
}

func (s *IntegerExpressionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ZZParserIdentifier, 0)
}

func (s *IntegerExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterIntegerExpression(s)
	}
}

func (s *IntegerExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitIntegerExpression(s)
	}
}

func (p *ZZParser) IntegerExpression() (localctx IIntegerExpressionContext) {
	localctx = NewIntegerExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ZZParserRULE_integerExpression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(141)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ZZParserIdentifier || _la == ZZParserIntegerLiteral) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IListElementContext is an interface to support dynamic dispatch.
type IListElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListElementContext differentiates from other interfaces.
	IsListElementContext()
}

type ListElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListElementContext() *ListElementContext {
	var p = new(ListElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_listElement
	return p
}

func (*ListElementContext) IsListElementContext() {}

func NewListElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListElementContext {
	var p = new(ListElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_listElement

	return p
}

func (s *ListElementContext) GetParser() antlr.Parser { return s.parser }

func (s *ListElementContext) IntegerLiteral() antlr.TerminalNode {
	return s.GetToken(ZZParserIntegerLiteral, 0)
}

func (s *ListElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterListElement(s)
	}
}

func (s *ListElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitListElement(s)
	}
}

func (p *ZZParser) ListElement() (localctx IListElementContext) {
	localctx = NewListElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ZZParserRULE_listElement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		p.Match(ZZParserT__2)
	}
	{
		p.SetState(144)
		p.Match(ZZParserIntegerLiteral)
	}
	{
		p.SetState(145)
		p.Match(ZZParserT__3)
	}

	return localctx
}

// IListElementsContext is an interface to support dynamic dispatch.
type IListElementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListElementsContext differentiates from other interfaces.
	IsListElementsContext()
}

type ListElementsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListElementsContext() *ListElementsContext {
	var p = new(ListElementsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_listElements
	return p
}

func (*ListElementsContext) IsListElementsContext() {}

func NewListElementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListElementsContext {
	var p = new(ListElementsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_listElements

	return p
}

func (s *ListElementsContext) GetParser() antlr.Parser { return s.parser }

func (s *ListElementsContext) ListElement() IListElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListElementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListElementContext)
}

func (s *ListElementsContext) ListElements() IListElementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListElementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListElementsContext)
}

func (s *ListElementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListElementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListElementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterListElements(s)
	}
}

func (s *ListElementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitListElements(s)
	}
}

func (p *ZZParser) ListElements() (localctx IListElementsContext) {
	return p.listElements(0)
}

func (p *ZZParser) listElements(_p int) (localctx IListElementsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewListElementsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListElementsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 20
	p.EnterRecursionRule(localctx, 20, ZZParserRULE_listElements, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(148)
		p.ListElement()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListElementsContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ZZParserRULE_listElements)
			p.SetState(150)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(151)
				p.ListElement()
			}

		}
		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
	}

	return localctx
}

// IListElementExpressionContext is an interface to support dynamic dispatch.
type IListElementExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListElementExpressionContext differentiates from other interfaces.
	IsListElementExpressionContext()
}

type ListElementExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListElementExpressionContext() *ListElementExpressionContext {
	var p = new(ListElementExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_listElementExpression
	return p
}

func (*ListElementExpressionContext) IsListElementExpressionContext() {}

func NewListElementExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListElementExpressionContext {
	var p = new(ListElementExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_listElementExpression

	return p
}

func (s *ListElementExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ListElementExpressionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ZZParserIdentifier, 0)
}

func (s *ListElementExpressionContext) ListElements() IListElementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListElementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListElementsContext)
}

func (s *ListElementExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListElementExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListElementExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterListElementExpression(s)
	}
}

func (s *ListElementExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitListElementExpression(s)
	}
}

func (p *ZZParser) ListElementExpression() (localctx IListElementExpressionContext) {
	localctx = NewListElementExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ZZParserRULE_listElementExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)
		p.Match(ZZParserIdentifier)
	}
	{
		p.SetState(158)
		p.listElements(0)
	}

	return localctx
}

// ITupleSizesContext is an interface to support dynamic dispatch.
type ITupleSizesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTupleSizesContext differentiates from other interfaces.
	IsTupleSizesContext()
}

type TupleSizesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTupleSizesContext() *TupleSizesContext {
	var p = new(TupleSizesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_tupleSizes
	return p
}

func (*TupleSizesContext) IsTupleSizesContext() {}

func NewTupleSizesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TupleSizesContext {
	var p = new(TupleSizesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_tupleSizes

	return p
}

func (s *TupleSizesContext) GetParser() antlr.Parser { return s.parser }

func (s *TupleSizesContext) IntegerExpression() IIntegerExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerExpressionContext)
}

func (s *TupleSizesContext) TupleSizes() ITupleSizesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITupleSizesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITupleSizesContext)
}

func (s *TupleSizesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TupleSizesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TupleSizesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterTupleSizes(s)
	}
}

func (s *TupleSizesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitTupleSizes(s)
	}
}

func (p *ZZParser) TupleSizes() (localctx ITupleSizesContext) {
	return p.tupleSizes(0)
}

func (p *ZZParser) tupleSizes(_p int) (localctx ITupleSizesContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewTupleSizesContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ITupleSizesContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 24
	p.EnterRecursionRule(localctx, 24, ZZParserRULE_tupleSizes, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(161)
		p.IntegerExpression()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewTupleSizesContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ZZParserRULE_tupleSizes)
			p.SetState(163)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(164)
				p.Match(ZZParserT__0)
			}
			{
				p.SetState(165)
				p.IntegerExpression()
			}

		}
		p.SetState(170)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
	}

	return localctx
}

// IAssignStatementContext is an interface to support dynamic dispatch.
type IAssignStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignStatementContext differentiates from other interfaces.
	IsAssignStatementContext()
}

type AssignStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignStatementContext() *AssignStatementContext {
	var p = new(AssignStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_assignStatement
	return p
}

func (*AssignStatementContext) IsAssignStatementContext() {}

func NewAssignStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignStatementContext {
	var p = new(AssignStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_assignStatement

	return p
}

func (s *AssignStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignStatementContext) DeclaratorList() IDeclaratorListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclaratorListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclaratorListContext)
}

func (s *AssignStatementContext) SimpleTypeSpecifier() ISimpleTypeSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimpleTypeSpecifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISimpleTypeSpecifierContext)
}

func (s *AssignStatementContext) AExpList() IAExpListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAExpListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAExpListContext)
}

func (s *AssignStatementContext) Declarator() IDeclaratorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclaratorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclaratorContext)
}

func (s *AssignStatementContext) List() antlr.TerminalNode {
	return s.GetToken(ZZParserList, 0)
}

func (s *AssignStatementContext) ListTypeSpecifier() IListTypeSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListTypeSpecifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListTypeSpecifierContext)
}

func (s *AssignStatementContext) TupleSizes() ITupleSizesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITupleSizesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITupleSizesContext)
}

func (s *AssignStatementContext) FuncExpression() IFuncExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncExpressionContext)
}

func (s *AssignStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterAssignStatement(s)
	}
}

func (s *AssignStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitAssignStatement(s)
	}
}

func (p *ZZParser) AssignStatement() (localctx IAssignStatementContext) {
	localctx = NewAssignStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ZZParserRULE_assignStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(171)
			p.declaratorList(0)
		}
		{
			p.SetState(172)
			p.SimpleTypeSpecifier()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(174)
			p.declaratorList(0)
		}
		{
			p.SetState(175)
			p.Match(ZZParserT__19)
		}
		{
			p.SetState(176)
			p.aExpList(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(178)
			p.Declarator()
		}
		{
			p.SetState(179)
			p.Match(ZZParserT__19)
		}
		{
			p.SetState(180)
			p.Match(ZZParserList)
		}
		{
			p.SetState(181)
			p.Match(ZZParserT__8)
		}
		{
			p.SetState(182)
			p.ListTypeSpecifier()
		}
		{
			p.SetState(183)
			p.Match(ZZParserT__0)
		}
		{
			p.SetState(184)
			p.Match(ZZParserT__8)
		}
		{
			p.SetState(185)
			p.tupleSizes(0)
		}
		{
			p.SetState(186)
			p.Match(ZZParserT__9)
		}
		{
			p.SetState(187)
			p.Match(ZZParserT__9)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(189)
			p.Declarator()
		}
		{
			p.SetState(190)
			p.Match(ZZParserT__19)
		}
		{
			p.SetState(191)
			p.FuncExpression()
		}

	}

	return localctx
}

// ISelectionStatementContext is an interface to support dynamic dispatch.
type ISelectionStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectionStatementContext differentiates from other interfaces.
	IsSelectionStatementContext()
}

type SelectionStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectionStatementContext() *SelectionStatementContext {
	var p = new(SelectionStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_selectionStatement
	return p
}

func (*SelectionStatementContext) IsSelectionStatementContext() {}

func NewSelectionStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectionStatementContext {
	var p = new(SelectionStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_selectionStatement

	return p
}

func (s *SelectionStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectionStatementContext) AllBExp() []IBExpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBExpContext)(nil)).Elem())
	var tst = make([]IBExpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBExpContext)
		}
	}

	return tst
}

func (s *SelectionStatementContext) BExp(i int) IBExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBExpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBExpContext)
}

func (s *SelectionStatementContext) AllFuncStatementList() []IFuncStatementListContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFuncStatementListContext)(nil)).Elem())
	var tst = make([]IFuncStatementListContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFuncStatementListContext)
		}
	}

	return tst
}

func (s *SelectionStatementContext) FuncStatementList(i int) IFuncStatementListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncStatementListContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFuncStatementListContext)
}

func (s *SelectionStatementContext) AllFuncStatement() []IFuncStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFuncStatementContext)(nil)).Elem())
	var tst = make([]IFuncStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFuncStatementContext)
		}
	}

	return tst
}

func (s *SelectionStatementContext) FuncStatement(i int) IFuncStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFuncStatementContext)
}

func (s *SelectionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectionStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectionStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterSelectionStatement(s)
	}
}

func (s *SelectionStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitSelectionStatement(s)
	}
}

func (p *ZZParser) SelectionStatement() (localctx ISelectionStatementContext) {
	localctx = NewSelectionStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ZZParserRULE_selectionStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.SetState(229)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ZZParserT__20:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(195)
			p.Match(ZZParserT__20)
		}
		{
			p.SetState(196)
			p.bExp(0)
		}
		{
			p.SetState(197)
			p.Match(ZZParserT__21)
		}
		p.SetState(199)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la-9)&-(0x1f+1)) == 0 && ((1<<uint((_la-9)))&((1<<(ZZParserT__8-9))|(1<<(ZZParserT__18-9))|(1<<(ZZParserT__20-9))|(1<<(ZZParserT__27-9))|(1<<(ZZParserT__29-9))|(1<<(ZZParserIdentifier-9))|(1<<(ZZParserIntegerLiteral-9))|(1<<(ZZParserFloatLiteral-9)))) != 0 {
			{
				p.SetState(198)
				p.funcStatementList(0)
			}

		}
		{
			p.SetState(201)
			p.Match(ZZParserT__22)
		}
		p.SetState(212)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(202)
					p.Match(ZZParserT__23)
				}
				{
					p.SetState(203)
					p.bExp(0)
				}
				{
					p.SetState(204)
					p.Match(ZZParserT__21)
				}
				p.SetState(206)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if ((_la-9)&-(0x1f+1)) == 0 && ((1<<uint((_la-9)))&((1<<(ZZParserT__8-9))|(1<<(ZZParserT__18-9))|(1<<(ZZParserT__20-9))|(1<<(ZZParserT__27-9))|(1<<(ZZParserT__29-9))|(1<<(ZZParserIdentifier-9))|(1<<(ZZParserIntegerLiteral-9))|(1<<(ZZParserFloatLiteral-9)))) != 0 {
					{
						p.SetState(205)
						p.funcStatementList(0)
					}

				}
				{
					p.SetState(208)
					p.Match(ZZParserT__22)
				}

			}
			p.SetState(214)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())
		}
		p.SetState(221)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(215)
				p.Match(ZZParserT__24)
			}
			{
				p.SetState(216)
				p.Match(ZZParserT__21)
			}
			p.SetState(218)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if ((_la-9)&-(0x1f+1)) == 0 && ((1<<uint((_la-9)))&((1<<(ZZParserT__8-9))|(1<<(ZZParserT__18-9))|(1<<(ZZParserT__20-9))|(1<<(ZZParserT__27-9))|(1<<(ZZParserT__29-9))|(1<<(ZZParserIdentifier-9))|(1<<(ZZParserIntegerLiteral-9))|(1<<(ZZParserFloatLiteral-9)))) != 0 {
				{
					p.SetState(217)
					p.funcStatementList(0)
				}

			}
			{
				p.SetState(220)
				p.Match(ZZParserT__22)
			}

		}

	case ZZParserT__8, ZZParserT__18, ZZParserIdentifier, ZZParserIntegerLiteral, ZZParserFloatLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(223)
			p.bExp(0)
		}
		{
			p.SetState(224)
			p.Match(ZZParserT__25)
		}
		{
			p.SetState(225)
			p.FuncStatement()
		}
		{
			p.SetState(226)
			p.Match(ZZParserT__26)
		}
		{
			p.SetState(227)
			p.FuncStatement()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIterationStatementContext is an interface to support dynamic dispatch.
type IIterationStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIterationStatementContext differentiates from other interfaces.
	IsIterationStatementContext()
}

type IterationStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIterationStatementContext() *IterationStatementContext {
	var p = new(IterationStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_iterationStatement
	return p
}

func (*IterationStatementContext) IsIterationStatementContext() {}

func NewIterationStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IterationStatementContext {
	var p = new(IterationStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_iterationStatement

	return p
}

func (s *IterationStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IterationStatementContext) FuncStatementList() IFuncStatementListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncStatementListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncStatementListContext)
}

func (s *IterationStatementContext) BExp() IBExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBExpContext)
}

func (s *IterationStatementContext) AllAssignStatement() []IAssignStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAssignStatementContext)(nil)).Elem())
	var tst = make([]IAssignStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAssignStatementContext)
		}
	}

	return tst
}

func (s *IterationStatementContext) AssignStatement(i int) IAssignStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAssignStatementContext)
}

func (s *IterationStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IterationStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IterationStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterIterationStatement(s)
	}
}

func (s *IterationStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitIterationStatement(s)
	}
}

func (p *ZZParser) IterationStatement() (localctx IIterationStatementContext) {
	localctx = NewIterationStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ZZParserRULE_iterationStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(251)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(231)
			p.Match(ZZParserT__27)
		}
		p.SetState(233)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la-9)&-(0x1f+1)) == 0 && ((1<<uint((_la-9)))&((1<<(ZZParserT__8-9))|(1<<(ZZParserT__18-9))|(1<<(ZZParserIdentifier-9))|(1<<(ZZParserIntegerLiteral-9))|(1<<(ZZParserFloatLiteral-9)))) != 0 {
			{
				p.SetState(232)
				p.bExp(0)
			}

		}
		{
			p.SetState(235)
			p.Match(ZZParserT__21)
		}
		{
			p.SetState(236)
			p.funcStatementList(0)
		}
		{
			p.SetState(237)
			p.Match(ZZParserT__22)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(239)
			p.Match(ZZParserT__27)
		}
		{
			p.SetState(240)
			p.AssignStatement()
		}
		{
			p.SetState(241)
			p.Match(ZZParserT__28)
		}
		p.SetState(243)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la-9)&-(0x1f+1)) == 0 && ((1<<uint((_la-9)))&((1<<(ZZParserT__8-9))|(1<<(ZZParserT__18-9))|(1<<(ZZParserIdentifier-9))|(1<<(ZZParserIntegerLiteral-9))|(1<<(ZZParserFloatLiteral-9)))) != 0 {
			{
				p.SetState(242)
				p.bExp(0)
			}

		}
		{
			p.SetState(245)
			p.Match(ZZParserT__28)
		}
		{
			p.SetState(246)
			p.AssignStatement()
		}
		{
			p.SetState(247)
			p.Match(ZZParserT__21)
		}
		{
			p.SetState(248)
			p.funcStatementList(0)
		}
		{
			p.SetState(249)
			p.Match(ZZParserT__22)
		}

	}

	return localctx
}

// IEntryContext is an interface to support dynamic dispatch.
type IEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEntryContext differentiates from other interfaces.
	IsEntryContext()
}

type EntryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEntryContext() *EntryContext {
	var p = new(EntryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_entry
	return p
}

func (*EntryContext) IsEntryContext() {}

func NewEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EntryContext {
	var p = new(EntryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_entry

	return p
}

func (s *EntryContext) GetParser() antlr.Parser { return s.parser }

func (s *EntryContext) AssignStatement() IAssignStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignStatementContext)
}

func (s *EntryContext) FuncExpressionWithName() IFuncExpressionWithNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncExpressionWithNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncExpressionWithNameContext)
}

func (s *EntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EntryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterEntry(s)
	}
}

func (s *EntryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitEntry(s)
	}
}

func (p *ZZParser) Entry() (localctx IEntryContext) {
	localctx = NewEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ZZParserRULE_entry)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(255)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ZZParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(253)
			p.AssignStatement()
		}

	case ZZParserFunc:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(254)
			p.FuncExpressionWithName()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IEntryListContext is an interface to support dynamic dispatch.
type IEntryListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEntryListContext differentiates from other interfaces.
	IsEntryListContext()
}

type EntryListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEntryListContext() *EntryListContext {
	var p = new(EntryListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_entryList
	return p
}

func (*EntryListContext) IsEntryListContext() {}

func NewEntryListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EntryListContext {
	var p = new(EntryListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_entryList

	return p
}

func (s *EntryListContext) GetParser() antlr.Parser { return s.parser }

func (s *EntryListContext) Entry() IEntryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEntryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEntryContext)
}

func (s *EntryListContext) EntryList() IEntryListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEntryListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEntryListContext)
}

func (s *EntryListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EntryListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EntryListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterEntryList(s)
	}
}

func (s *EntryListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitEntryList(s)
	}
}

func (p *ZZParser) EntryList() (localctx IEntryListContext) {
	return p.entryList(0)
}

func (p *ZZParser) entryList(_p int) (localctx IEntryListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewEntryListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IEntryListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 34
	p.EnterRecursionRule(localctx, 34, ZZParserRULE_entryList, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(258)
		p.Entry()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewEntryListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ZZParserRULE_entryList)
			p.SetState(260)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(261)
				p.Entry()
			}

		}
		p.SetState(266)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())
	}

	return localctx
}

// ITypeSpecifierContext is an interface to support dynamic dispatch.
type ITypeSpecifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeSpecifierContext differentiates from other interfaces.
	IsTypeSpecifierContext()
}

type TypeSpecifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeSpecifierContext() *TypeSpecifierContext {
	var p = new(TypeSpecifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_typeSpecifier
	return p
}

func (*TypeSpecifierContext) IsTypeSpecifierContext() {}

func NewTypeSpecifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSpecifierContext {
	var p = new(TypeSpecifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_typeSpecifier

	return p
}

func (s *TypeSpecifierContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeSpecifierContext) SimpleTypeSpecifier() ISimpleTypeSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimpleTypeSpecifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISimpleTypeSpecifierContext)
}

func (s *TypeSpecifierContext) ListTypeSpecifier() IListTypeSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListTypeSpecifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListTypeSpecifierContext)
}

func (s *TypeSpecifierContext) FuncSpecifier() IFuncSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncSpecifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncSpecifierContext)
}

func (s *TypeSpecifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSpecifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeSpecifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterTypeSpecifier(s)
	}
}

func (s *TypeSpecifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitTypeSpecifier(s)
	}
}

func (p *ZZParser) TypeSpecifier() (localctx ITypeSpecifierContext) {
	localctx = NewTypeSpecifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ZZParserRULE_typeSpecifier)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(267)
			p.SimpleTypeSpecifier()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(268)
			p.ListTypeSpecifier()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(269)
			p.FuncSpecifier()
		}

	}

	return localctx
}

// ITypeSpecifierListContext is an interface to support dynamic dispatch.
type ITypeSpecifierListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeSpecifierListContext differentiates from other interfaces.
	IsTypeSpecifierListContext()
}

type TypeSpecifierListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeSpecifierListContext() *TypeSpecifierListContext {
	var p = new(TypeSpecifierListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_typeSpecifierList
	return p
}

func (*TypeSpecifierListContext) IsTypeSpecifierListContext() {}

func NewTypeSpecifierListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSpecifierListContext {
	var p = new(TypeSpecifierListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_typeSpecifierList

	return p
}

func (s *TypeSpecifierListContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeSpecifierListContext) TypeSpecifier() ITypeSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierContext)
}

func (s *TypeSpecifierListContext) TypeSpecifierList() ITypeSpecifierListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecifierListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierListContext)
}

func (s *TypeSpecifierListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSpecifierListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeSpecifierListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterTypeSpecifierList(s)
	}
}

func (s *TypeSpecifierListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitTypeSpecifierList(s)
	}
}

func (p *ZZParser) TypeSpecifierList() (localctx ITypeSpecifierListContext) {
	return p.typeSpecifierList(0)
}

func (p *ZZParser) typeSpecifierList(_p int) (localctx ITypeSpecifierListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewTypeSpecifierListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ITypeSpecifierListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 38
	p.EnterRecursionRule(localctx, 38, ZZParserRULE_typeSpecifierList, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(273)
		p.TypeSpecifier()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(280)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewTypeSpecifierListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ZZParserRULE_typeSpecifierList)
			p.SetState(275)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(276)
				p.Match(ZZParserT__0)
			}
			{
				p.SetState(277)
				p.TypeSpecifier()
			}

		}
		p.SetState(282)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())
	}

	return localctx
}

// ITypeSpecifierWithIdentityContext is an interface to support dynamic dispatch.
type ITypeSpecifierWithIdentityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeSpecifierWithIdentityContext differentiates from other interfaces.
	IsTypeSpecifierWithIdentityContext()
}

type TypeSpecifierWithIdentityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeSpecifierWithIdentityContext() *TypeSpecifierWithIdentityContext {
	var p = new(TypeSpecifierWithIdentityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_typeSpecifierWithIdentity
	return p
}

func (*TypeSpecifierWithIdentityContext) IsTypeSpecifierWithIdentityContext() {}

func NewTypeSpecifierWithIdentityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSpecifierWithIdentityContext {
	var p = new(TypeSpecifierWithIdentityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_typeSpecifierWithIdentity

	return p
}

func (s *TypeSpecifierWithIdentityContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeSpecifierWithIdentityContext) DeclaratorList() IDeclaratorListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclaratorListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclaratorListContext)
}

func (s *TypeSpecifierWithIdentityContext) TypeSpecifier() ITypeSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierContext)
}

func (s *TypeSpecifierWithIdentityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSpecifierWithIdentityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeSpecifierWithIdentityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterTypeSpecifierWithIdentity(s)
	}
}

func (s *TypeSpecifierWithIdentityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitTypeSpecifierWithIdentity(s)
	}
}

func (p *ZZParser) TypeSpecifierWithIdentity() (localctx ITypeSpecifierWithIdentityContext) {
	localctx = NewTypeSpecifierWithIdentityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ZZParserRULE_typeSpecifierWithIdentity)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(283)
		p.declaratorList(0)
	}
	{
		p.SetState(284)
		p.TypeSpecifier()
	}

	return localctx
}

// ITypeSpecifierWithIdentityListContext is an interface to support dynamic dispatch.
type ITypeSpecifierWithIdentityListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeSpecifierWithIdentityListContext differentiates from other interfaces.
	IsTypeSpecifierWithIdentityListContext()
}

type TypeSpecifierWithIdentityListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeSpecifierWithIdentityListContext() *TypeSpecifierWithIdentityListContext {
	var p = new(TypeSpecifierWithIdentityListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_typeSpecifierWithIdentityList
	return p
}

func (*TypeSpecifierWithIdentityListContext) IsTypeSpecifierWithIdentityListContext() {}

func NewTypeSpecifierWithIdentityListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSpecifierWithIdentityListContext {
	var p = new(TypeSpecifierWithIdentityListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_typeSpecifierWithIdentityList

	return p
}

func (s *TypeSpecifierWithIdentityListContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeSpecifierWithIdentityListContext) TypeSpecifierWithIdentity() ITypeSpecifierWithIdentityContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecifierWithIdentityContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierWithIdentityContext)
}

func (s *TypeSpecifierWithIdentityListContext) TypeSpecifierWithIdentityList() ITypeSpecifierWithIdentityListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecifierWithIdentityListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierWithIdentityListContext)
}

func (s *TypeSpecifierWithIdentityListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSpecifierWithIdentityListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeSpecifierWithIdentityListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterTypeSpecifierWithIdentityList(s)
	}
}

func (s *TypeSpecifierWithIdentityListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitTypeSpecifierWithIdentityList(s)
	}
}

func (p *ZZParser) TypeSpecifierWithIdentityList() (localctx ITypeSpecifierWithIdentityListContext) {
	return p.typeSpecifierWithIdentityList(0)
}

func (p *ZZParser) typeSpecifierWithIdentityList(_p int) (localctx ITypeSpecifierWithIdentityListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewTypeSpecifierWithIdentityListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ITypeSpecifierWithIdentityListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 42
	p.EnterRecursionRule(localctx, 42, ZZParserRULE_typeSpecifierWithIdentityList, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(287)
		p.TypeSpecifierWithIdentity()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(294)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewTypeSpecifierWithIdentityListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ZZParserRULE_typeSpecifierWithIdentityList)
			p.SetState(289)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(290)
				p.Match(ZZParserT__0)
			}
			{
				p.SetState(291)
				p.TypeSpecifierWithIdentity()
			}

		}
		p.SetState(296)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())
	}

	return localctx
}

// IFuncSpecifierContext is an interface to support dynamic dispatch.
type IFuncSpecifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncSpecifierContext differentiates from other interfaces.
	IsFuncSpecifierContext()
}

type FuncSpecifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncSpecifierContext() *FuncSpecifierContext {
	var p = new(FuncSpecifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_funcSpecifier
	return p
}

func (*FuncSpecifierContext) IsFuncSpecifierContext() {}

func NewFuncSpecifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncSpecifierContext {
	var p = new(FuncSpecifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_funcSpecifier

	return p
}

func (s *FuncSpecifierContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncSpecifierContext) Func() antlr.TerminalNode {
	return s.GetToken(ZZParserFunc, 0)
}

func (s *FuncSpecifierContext) TypeSpecifierWithIdentityList() ITypeSpecifierWithIdentityListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecifierWithIdentityListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierWithIdentityListContext)
}

func (s *FuncSpecifierContext) TypeSpecifierList() ITypeSpecifierListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecifierListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierListContext)
}

func (s *FuncSpecifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncSpecifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncSpecifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterFuncSpecifier(s)
	}
}

func (s *FuncSpecifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitFuncSpecifier(s)
	}
}

func (p *ZZParser) FuncSpecifier() (localctx IFuncSpecifierContext) {
	localctx = NewFuncSpecifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, ZZParserRULE_funcSpecifier)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(297)
		p.Match(ZZParserFunc)
	}
	{
		p.SetState(298)
		p.Match(ZZParserT__8)
	}
	p.SetState(300)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ZZParserIdentifier {
		{
			p.SetState(299)
			p.typeSpecifierWithIdentityList(0)
		}

	}
	{
		p.SetState(302)
		p.Match(ZZParserT__9)
	}
	p.SetState(308)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(303)
			p.Match(ZZParserT__8)
		}
		p.SetState(305)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ZZParserT__1 || _la == ZZParserT__2 || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(ZZParserInt-32))|(1<<(ZZParserFloat-32))|(1<<(ZZParserFunc-32)))) != 0) {
			{
				p.SetState(304)
				p.typeSpecifierList(0)
			}

		}
		{
			p.SetState(307)
			p.Match(ZZParserT__9)
		}

	}

	return localctx
}

// IFuncSpecifierWithNameContext is an interface to support dynamic dispatch.
type IFuncSpecifierWithNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncSpecifierWithNameContext differentiates from other interfaces.
	IsFuncSpecifierWithNameContext()
}

type FuncSpecifierWithNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncSpecifierWithNameContext() *FuncSpecifierWithNameContext {
	var p = new(FuncSpecifierWithNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_funcSpecifierWithName
	return p
}

func (*FuncSpecifierWithNameContext) IsFuncSpecifierWithNameContext() {}

func NewFuncSpecifierWithNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncSpecifierWithNameContext {
	var p = new(FuncSpecifierWithNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_funcSpecifierWithName

	return p
}

func (s *FuncSpecifierWithNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncSpecifierWithNameContext) Func() antlr.TerminalNode {
	return s.GetToken(ZZParserFunc, 0)
}

func (s *FuncSpecifierWithNameContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ZZParserIdentifier, 0)
}

func (s *FuncSpecifierWithNameContext) TypeSpecifierWithIdentityList() ITypeSpecifierWithIdentityListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecifierWithIdentityListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierWithIdentityListContext)
}

func (s *FuncSpecifierWithNameContext) TypeSpecifierList() ITypeSpecifierListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecifierListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierListContext)
}

func (s *FuncSpecifierWithNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncSpecifierWithNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncSpecifierWithNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterFuncSpecifierWithName(s)
	}
}

func (s *FuncSpecifierWithNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitFuncSpecifierWithName(s)
	}
}

func (p *ZZParser) FuncSpecifierWithName() (localctx IFuncSpecifierWithNameContext) {
	localctx = NewFuncSpecifierWithNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, ZZParserRULE_funcSpecifierWithName)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(310)
		p.Match(ZZParserFunc)
	}
	{
		p.SetState(311)
		p.Match(ZZParserIdentifier)
	}
	{
		p.SetState(312)
		p.Match(ZZParserT__8)
	}
	p.SetState(314)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ZZParserIdentifier {
		{
			p.SetState(313)
			p.typeSpecifierWithIdentityList(0)
		}

	}
	{
		p.SetState(316)
		p.Match(ZZParserT__9)
	}
	p.SetState(322)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ZZParserT__8 {
		{
			p.SetState(317)
			p.Match(ZZParserT__8)
		}
		p.SetState(319)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ZZParserT__1 || _la == ZZParserT__2 || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(ZZParserInt-32))|(1<<(ZZParserFloat-32))|(1<<(ZZParserFunc-32)))) != 0) {
			{
				p.SetState(318)
				p.typeSpecifierList(0)
			}

		}
		{
			p.SetState(321)
			p.Match(ZZParserT__9)
		}

	}

	return localctx
}

// IFuncReturnParaContext is an interface to support dynamic dispatch.
type IFuncReturnParaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncReturnParaContext differentiates from other interfaces.
	IsFuncReturnParaContext()
}

type FuncReturnParaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncReturnParaContext() *FuncReturnParaContext {
	var p = new(FuncReturnParaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_funcReturnPara
	return p
}

func (*FuncReturnParaContext) IsFuncReturnParaContext() {}

func NewFuncReturnParaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncReturnParaContext {
	var p = new(FuncReturnParaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_funcReturnPara

	return p
}

func (s *FuncReturnParaContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncReturnParaContext) AExp() IAExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAExpContext)
}

func (s *FuncReturnParaContext) BExp() IBExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBExpContext)
}

func (s *FuncReturnParaContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ZZParserIdentifier, 0)
}

func (s *FuncReturnParaContext) Nil() antlr.TerminalNode {
	return s.GetToken(ZZParserNil, 0)
}

func (s *FuncReturnParaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncReturnParaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncReturnParaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterFuncReturnPara(s)
	}
}

func (s *FuncReturnParaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitFuncReturnPara(s)
	}
}

func (p *ZZParser) FuncReturnPara() (localctx IFuncReturnParaContext) {
	localctx = NewFuncReturnParaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, ZZParserRULE_funcReturnPara)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(328)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(324)
			p.aExp(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(325)
			p.bExp(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(326)
			p.Match(ZZParserIdentifier)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(327)
			p.Match(ZZParserNil)
		}

	}

	return localctx
}

// IFuncReturnParaListContext is an interface to support dynamic dispatch.
type IFuncReturnParaListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncReturnParaListContext differentiates from other interfaces.
	IsFuncReturnParaListContext()
}

type FuncReturnParaListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncReturnParaListContext() *FuncReturnParaListContext {
	var p = new(FuncReturnParaListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_funcReturnParaList
	return p
}

func (*FuncReturnParaListContext) IsFuncReturnParaListContext() {}

func NewFuncReturnParaListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncReturnParaListContext {
	var p = new(FuncReturnParaListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_funcReturnParaList

	return p
}

func (s *FuncReturnParaListContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncReturnParaListContext) FuncReturnPara() IFuncReturnParaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncReturnParaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncReturnParaContext)
}

func (s *FuncReturnParaListContext) FuncReturnParaList() IFuncReturnParaListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncReturnParaListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncReturnParaListContext)
}

func (s *FuncReturnParaListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncReturnParaListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncReturnParaListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterFuncReturnParaList(s)
	}
}

func (s *FuncReturnParaListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitFuncReturnParaList(s)
	}
}

func (p *ZZParser) FuncReturnParaList() (localctx IFuncReturnParaListContext) {
	return p.funcReturnParaList(0)
}

func (p *ZZParser) funcReturnParaList(_p int) (localctx IFuncReturnParaListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewFuncReturnParaListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IFuncReturnParaListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 50
	p.EnterRecursionRule(localctx, 50, ZZParserRULE_funcReturnParaList, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(331)
		p.FuncReturnPara()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(338)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewFuncReturnParaListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ZZParserRULE_funcReturnParaList)
			p.SetState(333)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(334)
				p.Match(ZZParserT__0)
			}
			{
				p.SetState(335)
				p.FuncReturnPara()
			}

		}
		p.SetState(340)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext())
	}

	return localctx
}

// IFuncStatementContext is an interface to support dynamic dispatch.
type IFuncStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncStatementContext differentiates from other interfaces.
	IsFuncStatementContext()
}

type FuncStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncStatementContext() *FuncStatementContext {
	var p = new(FuncStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_funcStatement
	return p
}

func (*FuncStatementContext) IsFuncStatementContext() {}

func NewFuncStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncStatementContext {
	var p = new(FuncStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_funcStatement

	return p
}

func (s *FuncStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncStatementContext) AssignStatement() IAssignStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignStatementContext)
}

func (s *FuncStatementContext) SelectionStatement() ISelectionStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectionStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectionStatementContext)
}

func (s *FuncStatementContext) IterationStatement() IIterationStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIterationStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIterationStatementContext)
}

func (s *FuncStatementContext) FuncReturnParaList() IFuncReturnParaListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncReturnParaListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncReturnParaListContext)
}

func (s *FuncStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterFuncStatement(s)
	}
}

func (s *FuncStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitFuncStatement(s)
	}
}

func (p *ZZParser) FuncStatement() (localctx IFuncStatementContext) {
	localctx = NewFuncStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, ZZParserRULE_funcStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(348)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(341)
			p.AssignStatement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(342)
			p.SelectionStatement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(343)
			p.IterationStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(344)
			p.Match(ZZParserT__29)
		}
		p.SetState(346)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(345)
				p.funcReturnParaList(0)
			}

		}

	}

	return localctx
}

// IFuncStatementListContext is an interface to support dynamic dispatch.
type IFuncStatementListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncStatementListContext differentiates from other interfaces.
	IsFuncStatementListContext()
}

type FuncStatementListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncStatementListContext() *FuncStatementListContext {
	var p = new(FuncStatementListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_funcStatementList
	return p
}

func (*FuncStatementListContext) IsFuncStatementListContext() {}

func NewFuncStatementListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncStatementListContext {
	var p = new(FuncStatementListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_funcStatementList

	return p
}

func (s *FuncStatementListContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncStatementListContext) FuncStatement() IFuncStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncStatementContext)
}

func (s *FuncStatementListContext) FuncStatementList() IFuncStatementListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncStatementListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncStatementListContext)
}

func (s *FuncStatementListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncStatementListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncStatementListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterFuncStatementList(s)
	}
}

func (s *FuncStatementListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitFuncStatementList(s)
	}
}

func (p *ZZParser) FuncStatementList() (localctx IFuncStatementListContext) {
	return p.funcStatementList(0)
}

func (p *ZZParser) funcStatementList(_p int) (localctx IFuncStatementListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewFuncStatementListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IFuncStatementListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 54
	p.EnterRecursionRule(localctx, 54, ZZParserRULE_funcStatementList, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(351)
		p.FuncStatement()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(357)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewFuncStatementListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ZZParserRULE_funcStatementList)
			p.SetState(353)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(354)
				p.FuncStatement()
			}

		}
		p.SetState(359)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext())
	}

	return localctx
}

// IFuncBodyContext is an interface to support dynamic dispatch.
type IFuncBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncBodyContext differentiates from other interfaces.
	IsFuncBodyContext()
}

type FuncBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncBodyContext() *FuncBodyContext {
	var p = new(FuncBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_funcBody
	return p
}

func (*FuncBodyContext) IsFuncBodyContext() {}

func NewFuncBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncBodyContext {
	var p = new(FuncBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_funcBody

	return p
}

func (s *FuncBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncBodyContext) FuncStatementList() IFuncStatementListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncStatementListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncStatementListContext)
}

func (s *FuncBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterFuncBody(s)
	}
}

func (s *FuncBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitFuncBody(s)
	}
}

func (p *ZZParser) FuncBody() (localctx IFuncBodyContext) {
	localctx = NewFuncBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, ZZParserRULE_funcBody)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(361)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-9)&-(0x1f+1)) == 0 && ((1<<uint((_la-9)))&((1<<(ZZParserT__8-9))|(1<<(ZZParserT__18-9))|(1<<(ZZParserT__20-9))|(1<<(ZZParserT__27-9))|(1<<(ZZParserT__29-9))|(1<<(ZZParserIdentifier-9))|(1<<(ZZParserIntegerLiteral-9))|(1<<(ZZParserFloatLiteral-9)))) != 0 {
		{
			p.SetState(360)
			p.funcStatementList(0)
		}

	}

	return localctx
}

// IFuncExpressionContext is an interface to support dynamic dispatch.
type IFuncExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncExpressionContext differentiates from other interfaces.
	IsFuncExpressionContext()
}

type FuncExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncExpressionContext() *FuncExpressionContext {
	var p = new(FuncExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_funcExpression
	return p
}

func (*FuncExpressionContext) IsFuncExpressionContext() {}

func NewFuncExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncExpressionContext {
	var p = new(FuncExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_funcExpression

	return p
}

func (s *FuncExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncExpressionContext) FuncSpecifier() IFuncSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncSpecifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncSpecifierContext)
}

func (s *FuncExpressionContext) FuncBody() IFuncBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncBodyContext)
}

func (s *FuncExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterFuncExpression(s)
	}
}

func (s *FuncExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitFuncExpression(s)
	}
}

func (p *ZZParser) FuncExpression() (localctx IFuncExpressionContext) {
	localctx = NewFuncExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, ZZParserRULE_funcExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(363)
		p.FuncSpecifier()
	}
	{
		p.SetState(364)
		p.Match(ZZParserT__21)
	}
	p.SetState(366)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(365)
			p.FuncBody()
		}

	}
	{
		p.SetState(368)
		p.Match(ZZParserT__22)
	}

	return localctx
}

// IFuncExpressionWithNameContext is an interface to support dynamic dispatch.
type IFuncExpressionWithNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncExpressionWithNameContext differentiates from other interfaces.
	IsFuncExpressionWithNameContext()
}

type FuncExpressionWithNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncExpressionWithNameContext() *FuncExpressionWithNameContext {
	var p = new(FuncExpressionWithNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_funcExpressionWithName
	return p
}

func (*FuncExpressionWithNameContext) IsFuncExpressionWithNameContext() {}

func NewFuncExpressionWithNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncExpressionWithNameContext {
	var p = new(FuncExpressionWithNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_funcExpressionWithName

	return p
}

func (s *FuncExpressionWithNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncExpressionWithNameContext) FuncSpecifierWithName() IFuncSpecifierWithNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncSpecifierWithNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncSpecifierWithNameContext)
}

func (s *FuncExpressionWithNameContext) FuncBody() IFuncBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncBodyContext)
}

func (s *FuncExpressionWithNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncExpressionWithNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncExpressionWithNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterFuncExpressionWithName(s)
	}
}

func (s *FuncExpressionWithNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitFuncExpressionWithName(s)
	}
}

func (p *ZZParser) FuncExpressionWithName() (localctx IFuncExpressionWithNameContext) {
	localctx = NewFuncExpressionWithNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, ZZParserRULE_funcExpressionWithName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(370)
		p.FuncSpecifierWithName()
	}
	{
		p.SetState(371)
		p.Match(ZZParserT__21)
	}
	p.SetState(373)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(372)
			p.FuncBody()
		}

	}
	{
		p.SetState(375)
		p.Match(ZZParserT__22)
	}

	return localctx
}

func (p *ZZParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *DeclaratorListContext = nil
		if localctx != nil {
			t = localctx.(*DeclaratorListContext)
		}
		return p.DeclaratorList_Sempred(t, predIndex)

	case 5:
		var t *AExpContext = nil
		if localctx != nil {
			t = localctx.(*AExpContext)
		}
		return p.AExp_Sempred(t, predIndex)

	case 6:
		var t *AExpListContext = nil
		if localctx != nil {
			t = localctx.(*AExpListContext)
		}
		return p.AExpList_Sempred(t, predIndex)

	case 7:
		var t *BExpContext = nil
		if localctx != nil {
			t = localctx.(*BExpContext)
		}
		return p.BExp_Sempred(t, predIndex)

	case 10:
		var t *ListElementsContext = nil
		if localctx != nil {
			t = localctx.(*ListElementsContext)
		}
		return p.ListElements_Sempred(t, predIndex)

	case 12:
		var t *TupleSizesContext = nil
		if localctx != nil {
			t = localctx.(*TupleSizesContext)
		}
		return p.TupleSizes_Sempred(t, predIndex)

	case 17:
		var t *EntryListContext = nil
		if localctx != nil {
			t = localctx.(*EntryListContext)
		}
		return p.EntryList_Sempred(t, predIndex)

	case 19:
		var t *TypeSpecifierListContext = nil
		if localctx != nil {
			t = localctx.(*TypeSpecifierListContext)
		}
		return p.TypeSpecifierList_Sempred(t, predIndex)

	case 21:
		var t *TypeSpecifierWithIdentityListContext = nil
		if localctx != nil {
			t = localctx.(*TypeSpecifierWithIdentityListContext)
		}
		return p.TypeSpecifierWithIdentityList_Sempred(t, predIndex)

	case 25:
		var t *FuncReturnParaListContext = nil
		if localctx != nil {
			t = localctx.(*FuncReturnParaListContext)
		}
		return p.FuncReturnParaList_Sempred(t, predIndex)

	case 27:
		var t *FuncStatementListContext = nil
		if localctx != nil {
			t = localctx.(*FuncStatementListContext)
		}
		return p.FuncStatementList_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ZZParser) DeclaratorList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ZZParser) AExp_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ZZParser) AExpList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ZZParser) BExp_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ZZParser) ListElements_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 5:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ZZParser) TupleSizes_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 6:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ZZParser) EntryList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 7:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ZZParser) TypeSpecifierList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 8:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ZZParser) TypeSpecifierWithIdentityList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 9:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ZZParser) FuncReturnParaList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 10:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ZZParser) FuncStatementList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 11:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
