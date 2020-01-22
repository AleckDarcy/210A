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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 42, 401,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 3, 2, 3, 2, 5, 2, 75, 10, 2, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 83, 10, 3, 12, 3, 14, 3, 86, 11, 3, 3,
	4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 94, 10, 5, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 118, 10, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 7, 8, 126, 10, 8, 12, 8, 14, 8, 129, 11, 8, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 137, 10, 9, 12, 9, 14, 9, 140, 11,
	9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 5, 10, 153, 10, 10, 3, 10, 3, 10, 3, 10, 7, 10, 158, 10, 10, 12,
	10, 14, 10, 161, 11, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 7, 13, 174, 10, 13, 12, 13, 14, 13, 177, 11,
	13, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 5, 15, 185, 10, 15, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 7, 16, 193, 10, 16, 12, 16, 14, 16,
	196, 11, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 5,
	18, 206, 10, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 213, 10, 18,
	3, 18, 3, 18, 7, 18, 217, 10, 18, 12, 18, 14, 18, 220, 11, 18, 3, 18, 3,
	18, 3, 18, 5, 18, 225, 10, 18, 3, 18, 5, 18, 228, 10, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 236, 10, 18, 3, 19, 3, 19, 5, 19, 240,
	10, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 248, 10, 19, 3,
	19, 3, 19, 5, 19, 252, 10, 19, 3, 19, 3, 19, 5, 19, 256, 10, 19, 3, 19,
	3, 19, 3, 19, 3, 19, 5, 19, 262, 10, 19, 3, 20, 3, 20, 5, 20, 266, 10,
	20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 7, 21, 273, 10, 21, 12, 21, 14,
	21, 276, 11, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 5, 23, 283, 10, 23,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 7, 24, 291, 10, 24, 12, 24, 14,
	24, 294, 11, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 7, 25, 302,
	10, 25, 12, 25, 14, 25, 305, 11, 25, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27,
	3, 27, 3, 27, 3, 27, 3, 27, 7, 27, 316, 10, 27, 12, 27, 14, 27, 319, 11,
	27, 3, 28, 3, 28, 3, 28, 5, 28, 324, 10, 28, 3, 28, 3, 28, 3, 28, 5, 28,
	329, 10, 28, 3, 28, 5, 28, 332, 10, 28, 3, 29, 3, 29, 3, 29, 3, 29, 5,
	29, 338, 10, 29, 3, 29, 3, 29, 3, 29, 5, 29, 343, 10, 29, 3, 29, 5, 29,
	346, 10, 29, 3, 30, 3, 30, 3, 30, 3, 30, 5, 30, 352, 10, 30, 3, 31, 3,
	31, 3, 31, 3, 31, 3, 31, 3, 31, 7, 31, 360, 10, 31, 12, 31, 14, 31, 363,
	11, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 5, 32, 370, 10, 32, 5, 32, 372,
	10, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 7, 33, 379, 10, 33, 12, 33,
	14, 33, 382, 11, 33, 3, 34, 5, 34, 385, 10, 34, 3, 35, 3, 35, 3, 35, 5,
	35, 390, 10, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 5, 36, 397, 10, 36,
	3, 36, 3, 36, 3, 36, 2, 14, 4, 14, 16, 18, 24, 30, 40, 46, 48, 52, 60,
	64, 37, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
	36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70,
	2, 8, 4, 2, 4, 4, 34, 35, 3, 2, 9, 10, 3, 2, 11, 12, 3, 2, 13, 18, 4, 2,
	13, 13, 18, 20, 3, 2, 39, 40, 2, 418, 2, 74, 3, 2, 2, 2, 4, 76, 3, 2, 2,
	2, 6, 87, 3, 2, 2, 2, 8, 93, 3, 2, 2, 2, 10, 95, 3, 2, 2, 2, 12, 99, 3,
	2, 2, 2, 14, 117, 3, 2, 2, 2, 16, 130, 3, 2, 2, 2, 18, 152, 3, 2, 2, 2,
	20, 162, 3, 2, 2, 2, 22, 164, 3, 2, 2, 2, 24, 168, 3, 2, 2, 2, 26, 178,
	3, 2, 2, 2, 28, 184, 3, 2, 2, 2, 30, 186, 3, 2, 2, 2, 32, 197, 3, 2, 2,
	2, 34, 235, 3, 2, 2, 2, 36, 261, 3, 2, 2, 2, 38, 265, 3, 2, 2, 2, 40, 267,
	3, 2, 2, 2, 42, 277, 3, 2, 2, 2, 44, 282, 3, 2, 2, 2, 46, 284, 3, 2, 2,
	2, 48, 295, 3, 2, 2, 2, 50, 306, 3, 2, 2, 2, 52, 309, 3, 2, 2, 2, 54, 320,
	3, 2, 2, 2, 56, 333, 3, 2, 2, 2, 58, 351, 3, 2, 2, 2, 60, 353, 3, 2, 2,
	2, 62, 371, 3, 2, 2, 2, 64, 373, 3, 2, 2, 2, 66, 384, 3, 2, 2, 2, 68, 386,
	3, 2, 2, 2, 70, 393, 3, 2, 2, 2, 72, 75, 7, 39, 2, 2, 73, 75, 5, 26, 14,
	2, 74, 72, 3, 2, 2, 2, 74, 73, 3, 2, 2, 2, 75, 3, 3, 2, 2, 2, 76, 77, 8,
	3, 1, 2, 77, 78, 5, 2, 2, 2, 78, 84, 3, 2, 2, 2, 79, 80, 12, 3, 2, 2, 80,
	81, 7, 3, 2, 2, 81, 83, 5, 2, 2, 2, 82, 79, 3, 2, 2, 2, 83, 86, 3, 2, 2,
	2, 84, 82, 3, 2, 2, 2, 84, 85, 3, 2, 2, 2, 85, 5, 3, 2, 2, 2, 86, 84, 3,
	2, 2, 2, 87, 88, 9, 2, 2, 2, 88, 7, 3, 2, 2, 2, 89, 94, 5, 6, 4, 2, 90,
	91, 7, 5, 2, 2, 91, 92, 7, 6, 2, 2, 92, 94, 5, 8, 5, 2, 93, 89, 3, 2, 2,
	2, 93, 90, 3, 2, 2, 2, 94, 9, 3, 2, 2, 2, 95, 96, 7, 5, 2, 2, 96, 97, 7,
	6, 2, 2, 97, 98, 5, 8, 5, 2, 98, 11, 3, 2, 2, 2, 99, 100, 7, 37, 2, 2,
	100, 101, 7, 7, 2, 2, 101, 102, 5, 10, 6, 2, 102, 103, 7, 3, 2, 2, 103,
	104, 7, 7, 2, 2, 104, 105, 5, 14, 8, 2, 105, 106, 7, 8, 2, 2, 106, 107,
	7, 8, 2, 2, 107, 13, 3, 2, 2, 2, 108, 109, 8, 8, 1, 2, 109, 118, 7, 40,
	2, 2, 110, 118, 7, 41, 2, 2, 111, 118, 7, 39, 2, 2, 112, 118, 5, 26, 14,
	2, 113, 114, 7, 7, 2, 2, 114, 115, 5, 14, 8, 2, 115, 116, 7, 8, 2, 2, 116,
	118, 3, 2, 2, 2, 117, 108, 3, 2, 2, 2, 117, 110, 3, 2, 2, 2, 117, 111,
	3, 2, 2, 2, 117, 112, 3, 2, 2, 2, 117, 113, 3, 2, 2, 2, 118, 127, 3, 2,
	2, 2, 119, 120, 12, 5, 2, 2, 120, 121, 9, 3, 2, 2, 121, 126, 5, 14, 8,
	6, 122, 123, 12, 4, 2, 2, 123, 124, 9, 4, 2, 2, 124, 126, 5, 14, 8, 5,
	125, 119, 3, 2, 2, 2, 125, 122, 3, 2, 2, 2, 126, 129, 3, 2, 2, 2, 127,
	125, 3, 2, 2, 2, 127, 128, 3, 2, 2, 2, 128, 15, 3, 2, 2, 2, 129, 127, 3,
	2, 2, 2, 130, 131, 8, 9, 1, 2, 131, 132, 5, 14, 8, 2, 132, 138, 3, 2, 2,
	2, 133, 134, 12, 3, 2, 2, 134, 135, 7, 3, 2, 2, 135, 137, 5, 14, 8, 2,
	136, 133, 3, 2, 2, 2, 137, 140, 3, 2, 2, 2, 138, 136, 3, 2, 2, 2, 138,
	139, 3, 2, 2, 2, 139, 17, 3, 2, 2, 2, 140, 138, 3, 2, 2, 2, 141, 142, 8,
	10, 1, 2, 142, 143, 5, 14, 8, 2, 143, 144, 9, 5, 2, 2, 144, 145, 5, 14,
	8, 2, 145, 153, 3, 2, 2, 2, 146, 147, 7, 21, 2, 2, 147, 153, 5, 18, 10,
	4, 148, 149, 7, 7, 2, 2, 149, 150, 5, 18, 10, 2, 150, 151, 7, 8, 2, 2,
	151, 153, 3, 2, 2, 2, 152, 141, 3, 2, 2, 2, 152, 146, 3, 2, 2, 2, 152,
	148, 3, 2, 2, 2, 153, 159, 3, 2, 2, 2, 154, 155, 12, 5, 2, 2, 155, 156,
	9, 6, 2, 2, 156, 158, 5, 18, 10, 6, 157, 154, 3, 2, 2, 2, 158, 161, 3,
	2, 2, 2, 159, 157, 3, 2, 2, 2, 159, 160, 3, 2, 2, 2, 160, 19, 3, 2, 2,
	2, 161, 159, 3, 2, 2, 2, 162, 163, 9, 7, 2, 2, 163, 21, 3, 2, 2, 2, 164,
	165, 7, 5, 2, 2, 165, 166, 5, 14, 8, 2, 166, 167, 7, 6, 2, 2, 167, 23,
	3, 2, 2, 2, 168, 169, 8, 13, 1, 2, 169, 170, 5, 22, 12, 2, 170, 175, 3,
	2, 2, 2, 171, 172, 12, 3, 2, 2, 172, 174, 5, 22, 12, 2, 173, 171, 3, 2,
	2, 2, 174, 177, 3, 2, 2, 2, 175, 173, 3, 2, 2, 2, 175, 176, 3, 2, 2, 2,
	176, 25, 3, 2, 2, 2, 177, 175, 3, 2, 2, 2, 178, 179, 7, 39, 2, 2, 179,
	180, 5, 24, 13, 2, 180, 27, 3, 2, 2, 2, 181, 185, 5, 14, 8, 2, 182, 185,
	5, 12, 7, 2, 183, 185, 5, 68, 35, 2, 184, 181, 3, 2, 2, 2, 184, 182, 3,
	2, 2, 2, 184, 183, 3, 2, 2, 2, 185, 29, 3, 2, 2, 2, 186, 187, 8, 16, 1,
	2, 187, 188, 5, 28, 15, 2, 188, 194, 3, 2, 2, 2, 189, 190, 12, 3, 2, 2,
	190, 191, 7, 3, 2, 2, 191, 193, 5, 28, 15, 2, 192, 189, 3, 2, 2, 2, 193,
	196, 3, 2, 2, 2, 194, 192, 3, 2, 2, 2, 194, 195, 3, 2, 2, 2, 195, 31, 3,
	2, 2, 2, 196, 194, 3, 2, 2, 2, 197, 198, 5, 4, 3, 2, 198, 199, 7, 22, 2,
	2, 199, 200, 5, 30, 16, 2, 200, 33, 3, 2, 2, 2, 201, 202, 7, 23, 2, 2,
	202, 203, 5, 18, 10, 2, 203, 205, 7, 24, 2, 2, 204, 206, 5, 64, 33, 2,
	205, 204, 3, 2, 2, 2, 205, 206, 3, 2, 2, 2, 206, 207, 3, 2, 2, 2, 207,
	218, 7, 25, 2, 2, 208, 209, 7, 26, 2, 2, 209, 210, 5, 18, 10, 2, 210, 212,
	7, 24, 2, 2, 211, 213, 5, 64, 33, 2, 212, 211, 3, 2, 2, 2, 212, 213, 3,
	2, 2, 2, 213, 214, 3, 2, 2, 2, 214, 215, 7, 25, 2, 2, 215, 217, 3, 2, 2,
	2, 216, 208, 3, 2, 2, 2, 217, 220, 3, 2, 2, 2, 218, 216, 3, 2, 2, 2, 218,
	219, 3, 2, 2, 2, 219, 227, 3, 2, 2, 2, 220, 218, 3, 2, 2, 2, 221, 222,
	7, 27, 2, 2, 222, 224, 7, 24, 2, 2, 223, 225, 5, 64, 33, 2, 224, 223, 3,
	2, 2, 2, 224, 225, 3, 2, 2, 2, 225, 226, 3, 2, 2, 2, 226, 228, 7, 25, 2,
	2, 227, 221, 3, 2, 2, 2, 227, 228, 3, 2, 2, 2, 228, 236, 3, 2, 2, 2, 229,
	230, 5, 18, 10, 2, 230, 231, 7, 28, 2, 2, 231, 232, 5, 62, 32, 2, 232,
	233, 7, 29, 2, 2, 233, 234, 5, 62, 32, 2, 234, 236, 3, 2, 2, 2, 235, 201,
	3, 2, 2, 2, 235, 229, 3, 2, 2, 2, 236, 35, 3, 2, 2, 2, 237, 239, 7, 30,
	2, 2, 238, 240, 5, 18, 10, 2, 239, 238, 3, 2, 2, 2, 239, 240, 3, 2, 2,
	2, 240, 241, 3, 2, 2, 2, 241, 242, 7, 24, 2, 2, 242, 243, 5, 64, 33, 2,
	243, 244, 7, 25, 2, 2, 244, 262, 3, 2, 2, 2, 245, 247, 7, 30, 2, 2, 246,
	248, 5, 32, 17, 2, 247, 246, 3, 2, 2, 2, 247, 248, 3, 2, 2, 2, 248, 249,
	3, 2, 2, 2, 249, 251, 7, 31, 2, 2, 250, 252, 5, 18, 10, 2, 251, 250, 3,
	2, 2, 2, 251, 252, 3, 2, 2, 2, 252, 253, 3, 2, 2, 2, 253, 255, 7, 31, 2,
	2, 254, 256, 5, 32, 17, 2, 255, 254, 3, 2, 2, 2, 255, 256, 3, 2, 2, 2,
	256, 257, 3, 2, 2, 2, 257, 258, 7, 24, 2, 2, 258, 259, 5, 64, 33, 2, 259,
	260, 7, 25, 2, 2, 260, 262, 3, 2, 2, 2, 261, 237, 3, 2, 2, 2, 261, 245,
	3, 2, 2, 2, 262, 37, 3, 2, 2, 2, 263, 266, 5, 32, 17, 2, 264, 266, 5, 70,
	36, 2, 265, 263, 3, 2, 2, 2, 265, 264, 3, 2, 2, 2, 266, 39, 3, 2, 2, 2,
	267, 268, 8, 21, 1, 2, 268, 269, 5, 38, 20, 2, 269, 274, 3, 2, 2, 2, 270,
	271, 12, 3, 2, 2, 271, 273, 5, 38, 20, 2, 272, 270, 3, 2, 2, 2, 273, 276,
	3, 2, 2, 2, 274, 272, 3, 2, 2, 2, 274, 275, 3, 2, 2, 2, 275, 41, 3, 2,
	2, 2, 276, 274, 3, 2, 2, 2, 277, 278, 5, 4, 3, 2, 278, 43, 3, 2, 2, 2,
	279, 283, 5, 6, 4, 2, 280, 283, 5, 10, 6, 2, 281, 283, 5, 54, 28, 2, 282,
	279, 3, 2, 2, 2, 282, 280, 3, 2, 2, 2, 282, 281, 3, 2, 2, 2, 283, 45, 3,
	2, 2, 2, 284, 285, 8, 24, 1, 2, 285, 286, 5, 44, 23, 2, 286, 292, 3, 2,
	2, 2, 287, 288, 12, 3, 2, 2, 288, 289, 7, 3, 2, 2, 289, 291, 5, 44, 23,
	2, 290, 287, 3, 2, 2, 2, 291, 294, 3, 2, 2, 2, 292, 290, 3, 2, 2, 2, 292,
	293, 3, 2, 2, 2, 293, 47, 3, 2, 2, 2, 294, 292, 3, 2, 2, 2, 295, 296, 8,
	25, 1, 2, 296, 297, 7, 39, 2, 2, 297, 303, 3, 2, 2, 2, 298, 299, 12, 3,
	2, 2, 299, 300, 7, 3, 2, 2, 300, 302, 7, 39, 2, 2, 301, 298, 3, 2, 2, 2,
	302, 305, 3, 2, 2, 2, 303, 301, 3, 2, 2, 2, 303, 304, 3, 2, 2, 2, 304,
	49, 3, 2, 2, 2, 305, 303, 3, 2, 2, 2, 306, 307, 5, 4, 3, 2, 307, 308, 5,
	44, 23, 2, 308, 51, 3, 2, 2, 2, 309, 310, 8, 27, 1, 2, 310, 311, 5, 50,
	26, 2, 311, 317, 3, 2, 2, 2, 312, 313, 12, 3, 2, 2, 313, 314, 7, 3, 2,
	2, 314, 316, 5, 50, 26, 2, 315, 312, 3, 2, 2, 2, 316, 319, 3, 2, 2, 2,
	317, 315, 3, 2, 2, 2, 317, 318, 3, 2, 2, 2, 318, 53, 3, 2, 2, 2, 319, 317,
	3, 2, 2, 2, 320, 321, 7, 36, 2, 2, 321, 323, 7, 7, 2, 2, 322, 324, 5, 52,
	27, 2, 323, 322, 3, 2, 2, 2, 323, 324, 3, 2, 2, 2, 324, 325, 3, 2, 2, 2,
	325, 331, 7, 8, 2, 2, 326, 328, 7, 7, 2, 2, 327, 329, 5, 46, 24, 2, 328,
	327, 3, 2, 2, 2, 328, 329, 3, 2, 2, 2, 329, 330, 3, 2, 2, 2, 330, 332,
	7, 8, 2, 2, 331, 326, 3, 2, 2, 2, 331, 332, 3, 2, 2, 2, 332, 55, 3, 2,
	2, 2, 333, 334, 7, 36, 2, 2, 334, 335, 7, 39, 2, 2, 335, 337, 7, 7, 2,
	2, 336, 338, 5, 52, 27, 2, 337, 336, 3, 2, 2, 2, 337, 338, 3, 2, 2, 2,
	338, 339, 3, 2, 2, 2, 339, 345, 7, 8, 2, 2, 340, 342, 7, 7, 2, 2, 341,
	343, 5, 46, 24, 2, 342, 341, 3, 2, 2, 2, 342, 343, 3, 2, 2, 2, 343, 344,
	3, 2, 2, 2, 344, 346, 7, 8, 2, 2, 345, 340, 3, 2, 2, 2, 345, 346, 3, 2,
	2, 2, 346, 57, 3, 2, 2, 2, 347, 352, 5, 14, 8, 2, 348, 352, 5, 18, 10,
	2, 349, 352, 7, 39, 2, 2, 350, 352, 7, 38, 2, 2, 351, 347, 3, 2, 2, 2,
	351, 348, 3, 2, 2, 2, 351, 349, 3, 2, 2, 2, 351, 350, 3, 2, 2, 2, 352,
	59, 3, 2, 2, 2, 353, 354, 8, 31, 1, 2, 354, 355, 5, 58, 30, 2, 355, 361,
	3, 2, 2, 2, 356, 357, 12, 3, 2, 2, 357, 358, 7, 3, 2, 2, 358, 360, 5, 58,
	30, 2, 359, 356, 3, 2, 2, 2, 360, 363, 3, 2, 2, 2, 361, 359, 3, 2, 2, 2,
	361, 362, 3, 2, 2, 2, 362, 61, 3, 2, 2, 2, 363, 361, 3, 2, 2, 2, 364, 372,
	5, 32, 17, 2, 365, 372, 5, 34, 18, 2, 366, 372, 5, 36, 19, 2, 367, 369,
	7, 32, 2, 2, 368, 370, 5, 60, 31, 2, 369, 368, 3, 2, 2, 2, 369, 370, 3,
	2, 2, 2, 370, 372, 3, 2, 2, 2, 371, 364, 3, 2, 2, 2, 371, 365, 3, 2, 2,
	2, 371, 366, 3, 2, 2, 2, 371, 367, 3, 2, 2, 2, 372, 63, 3, 2, 2, 2, 373,
	374, 8, 33, 1, 2, 374, 375, 5, 62, 32, 2, 375, 380, 3, 2, 2, 2, 376, 377,
	12, 3, 2, 2, 377, 379, 5, 62, 32, 2, 378, 376, 3, 2, 2, 2, 379, 382, 3,
	2, 2, 2, 380, 378, 3, 2, 2, 2, 380, 381, 3, 2, 2, 2, 381, 65, 3, 2, 2,
	2, 382, 380, 3, 2, 2, 2, 383, 385, 5, 64, 33, 2, 384, 383, 3, 2, 2, 2,
	384, 385, 3, 2, 2, 2, 385, 67, 3, 2, 2, 2, 386, 387, 5, 54, 28, 2, 387,
	389, 7, 24, 2, 2, 388, 390, 5, 66, 34, 2, 389, 388, 3, 2, 2, 2, 389, 390,
	3, 2, 2, 2, 390, 391, 3, 2, 2, 2, 391, 392, 7, 25, 2, 2, 392, 69, 3, 2,
	2, 2, 393, 394, 5, 56, 29, 2, 394, 396, 7, 24, 2, 2, 395, 397, 5, 66, 34,
	2, 396, 395, 3, 2, 2, 2, 396, 397, 3, 2, 2, 2, 397, 398, 3, 2, 2, 2, 398,
	399, 7, 25, 2, 2, 399, 71, 3, 2, 2, 2, 45, 74, 84, 93, 117, 125, 127, 138,
	152, 159, 175, 184, 194, 205, 212, 218, 224, 227, 235, 239, 247, 251, 255,
	261, 265, 274, 282, 292, 303, 317, 323, 328, 331, 337, 342, 345, 351, 361,
	369, 371, 380, 384, 389, 396,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "','", "'string'", "'['", "']'", "'('", "')'", "'*'", "'/'", "'+'",
	"'-'", "'=='", "'<'", "'>'", "'>='", "'<='", "'!='", "'&&'", "'||'", "'!'",
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
	"listTypeSpecifier", "listInitExpression", "aExpr", "aExprList", "bExp",
	"integerExpression", "listElementIndex", "listElementIndexList", "listElementExpression",
	"assignInit", "assignInitList", "assignStatement", "selectionStatement",
	"iterationStatement", "definition", "definitionList", "file", "typeSpecifier",
	"typeSpecifierList", "paraDeclaratorList", "paraDeclaratorWithIdentity",
	"paraDeclaratorWithIdentityList", "funcTypeSpecifier", "funcTypeSpecifierWithName",
	"funcReturnPara", "funcReturnParaList", "funcStatement", "funcStatementList",
	"funcBody", "funcInitExpression", "funcDefinition",
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
	ZZParserRULE_declarator                     = 0
	ZZParserRULE_declaratorList                 = 1
	ZZParserRULE_simpleTypeSpecifier            = 2
	ZZParserRULE_listElementTypeSpecifier       = 3
	ZZParserRULE_listTypeSpecifier              = 4
	ZZParserRULE_listInitExpression             = 5
	ZZParserRULE_aExpr                          = 6
	ZZParserRULE_aExprList                      = 7
	ZZParserRULE_bExp                           = 8
	ZZParserRULE_integerExpression              = 9
	ZZParserRULE_listElementIndex               = 10
	ZZParserRULE_listElementIndexList           = 11
	ZZParserRULE_listElementExpression          = 12
	ZZParserRULE_assignInit                     = 13
	ZZParserRULE_assignInitList                 = 14
	ZZParserRULE_assignStatement                = 15
	ZZParserRULE_selectionStatement             = 16
	ZZParserRULE_iterationStatement             = 17
	ZZParserRULE_definition                     = 18
	ZZParserRULE_definitionList                 = 19
	ZZParserRULE_file                           = 20
	ZZParserRULE_typeSpecifier                  = 21
	ZZParserRULE_typeSpecifierList              = 22
	ZZParserRULE_paraDeclaratorList             = 23
	ZZParserRULE_paraDeclaratorWithIdentity     = 24
	ZZParserRULE_paraDeclaratorWithIdentityList = 25
	ZZParserRULE_funcTypeSpecifier              = 26
	ZZParserRULE_funcTypeSpecifierWithName      = 27
	ZZParserRULE_funcReturnPara                 = 28
	ZZParserRULE_funcReturnParaList             = 29
	ZZParserRULE_funcStatement                  = 30
	ZZParserRULE_funcStatementList              = 31
	ZZParserRULE_funcBody                       = 32
	ZZParserRULE_funcInitExpression             = 33
	ZZParserRULE_funcDefinition                 = 34
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

	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(70)
			p.Match(ZZParserIdentifier)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(71)
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
		p.SetState(75)
		p.Declarator()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(82)
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
			p.SetState(77)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(78)
				p.Match(ZZParserT__0)
			}
			{
				p.SetState(79)
				p.Declarator()
			}

		}
		p.SetState(84)
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
		p.SetState(85)
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

func (s *ListElementTypeSpecifierContext) ListElementTypeSpecifier() IListElementTypeSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListElementTypeSpecifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListElementTypeSpecifierContext)
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

	p.SetState(91)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ZZParserT__1, ZZParserInt, ZZParserFloat:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(87)
			p.SimpleTypeSpecifier()
		}

	case ZZParserT__2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(88)
			p.Match(ZZParserT__2)
		}
		{
			p.SetState(89)
			p.Match(ZZParserT__3)
		}
		{
			p.SetState(90)
			p.ListElementTypeSpecifier()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(93)
		p.Match(ZZParserT__2)
	}
	{
		p.SetState(94)
		p.Match(ZZParserT__3)
	}
	{
		p.SetState(95)
		p.ListElementTypeSpecifier()
	}

	return localctx
}

// IListInitExpressionContext is an interface to support dynamic dispatch.
type IListInitExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListInitExpressionContext differentiates from other interfaces.
	IsListInitExpressionContext()
}

type ListInitExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListInitExpressionContext() *ListInitExpressionContext {
	var p = new(ListInitExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_listInitExpression
	return p
}

func (*ListInitExpressionContext) IsListInitExpressionContext() {}

func NewListInitExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListInitExpressionContext {
	var p = new(ListInitExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_listInitExpression

	return p
}

func (s *ListInitExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ListInitExpressionContext) List() antlr.TerminalNode {
	return s.GetToken(ZZParserList, 0)
}

func (s *ListInitExpressionContext) ListTypeSpecifier() IListTypeSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListTypeSpecifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListTypeSpecifierContext)
}

func (s *ListInitExpressionContext) AExpr() IAExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAExprContext)
}

func (s *ListInitExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListInitExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListInitExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterListInitExpression(s)
	}
}

func (s *ListInitExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitListInitExpression(s)
	}
}

func (p *ZZParser) ListInitExpression() (localctx IListInitExpressionContext) {
	localctx = NewListInitExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ZZParserRULE_listInitExpression)

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
		p.SetState(97)
		p.Match(ZZParserList)
	}
	{
		p.SetState(98)
		p.Match(ZZParserT__4)
	}
	{
		p.SetState(99)
		p.ListTypeSpecifier()
	}
	{
		p.SetState(100)
		p.Match(ZZParserT__0)
	}
	{
		p.SetState(101)
		p.Match(ZZParserT__4)
	}
	{
		p.SetState(102)
		p.aExpr(0)
	}
	{
		p.SetState(103)
		p.Match(ZZParserT__5)
	}
	{
		p.SetState(104)
		p.Match(ZZParserT__5)
	}

	return localctx
}

// IAExprContext is an interface to support dynamic dispatch.
type IAExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAExprContext differentiates from other interfaces.
	IsAExprContext()
}

type AExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAExprContext() *AExprContext {
	var p = new(AExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_aExpr
	return p
}

func (*AExprContext) IsAExprContext() {}

func NewAExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AExprContext {
	var p = new(AExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_aExpr

	return p
}

func (s *AExprContext) GetParser() antlr.Parser { return s.parser }

func (s *AExprContext) CopyFrom(ctx *AExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *AExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AExp_bracketExpressionContext struct {
	*AExprContext
}

func NewAExp_bracketExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AExp_bracketExpressionContext {
	var p = new(AExp_bracketExpressionContext)

	p.AExprContext = NewEmptyAExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AExprContext))

	return p
}

func (s *AExp_bracketExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AExp_bracketExpressionContext) AExpr() IAExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAExprContext)
}

func (s *AExp_bracketExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterAExp_bracketExpression(s)
	}
}

func (s *AExp_bracketExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitAExp_bracketExpression(s)
	}
}

type AExp_FloatLiteralContext struct {
	*AExprContext
}

func NewAExp_FloatLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AExp_FloatLiteralContext {
	var p = new(AExp_FloatLiteralContext)

	p.AExprContext = NewEmptyAExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AExprContext))

	return p
}

func (s *AExp_FloatLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AExp_FloatLiteralContext) FloatLiteral() antlr.TerminalNode {
	return s.GetToken(ZZParserFloatLiteral, 0)
}

func (s *AExp_FloatLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterAExp_FloatLiteral(s)
	}
}

func (s *AExp_FloatLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitAExp_FloatLiteral(s)
	}
}

type AExp_multiplicativeExpressionContext struct {
	*AExprContext
}

func NewAExp_multiplicativeExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AExp_multiplicativeExpressionContext {
	var p = new(AExp_multiplicativeExpressionContext)

	p.AExprContext = NewEmptyAExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AExprContext))

	return p
}

func (s *AExp_multiplicativeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AExp_multiplicativeExpressionContext) AllAExpr() []IAExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAExprContext)(nil)).Elem())
	var tst = make([]IAExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAExprContext)
		}
	}

	return tst
}

func (s *AExp_multiplicativeExpressionContext) AExpr(i int) IAExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAExprContext)
}

func (s *AExp_multiplicativeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterAExp_multiplicativeExpression(s)
	}
}

func (s *AExp_multiplicativeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitAExp_multiplicativeExpression(s)
	}
}

type AExp_additiveExpressionContext struct {
	*AExprContext
}

func NewAExp_additiveExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AExp_additiveExpressionContext {
	var p = new(AExp_additiveExpressionContext)

	p.AExprContext = NewEmptyAExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AExprContext))

	return p
}

func (s *AExp_additiveExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AExp_additiveExpressionContext) AllAExpr() []IAExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAExprContext)(nil)).Elem())
	var tst = make([]IAExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAExprContext)
		}
	}

	return tst
}

func (s *AExp_additiveExpressionContext) AExpr(i int) IAExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAExprContext)
}

func (s *AExp_additiveExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterAExp_additiveExpression(s)
	}
}

func (s *AExp_additiveExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitAExp_additiveExpression(s)
	}
}

type AExp_IdentifierContext struct {
	*AExprContext
}

func NewAExp_IdentifierContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AExp_IdentifierContext {
	var p = new(AExp_IdentifierContext)

	p.AExprContext = NewEmptyAExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AExprContext))

	return p
}

func (s *AExp_IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AExp_IdentifierContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ZZParserIdentifier, 0)
}

func (s *AExp_IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterAExp_Identifier(s)
	}
}

func (s *AExp_IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitAExp_Identifier(s)
	}
}

type AExp_IntergerLiteralContext struct {
	*AExprContext
}

func NewAExp_IntergerLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AExp_IntergerLiteralContext {
	var p = new(AExp_IntergerLiteralContext)

	p.AExprContext = NewEmptyAExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AExprContext))

	return p
}

func (s *AExp_IntergerLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AExp_IntergerLiteralContext) IntegerLiteral() antlr.TerminalNode {
	return s.GetToken(ZZParserIntegerLiteral, 0)
}

func (s *AExp_IntergerLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterAExp_IntergerLiteral(s)
	}
}

func (s *AExp_IntergerLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitAExp_IntergerLiteral(s)
	}
}

type AExp_listElementExpressionContext struct {
	*AExprContext
}

func NewAExp_listElementExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AExp_listElementExpressionContext {
	var p = new(AExp_listElementExpressionContext)

	p.AExprContext = NewEmptyAExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AExprContext))

	return p
}

func (s *AExp_listElementExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AExp_listElementExpressionContext) ListElementExpression() IListElementExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListElementExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListElementExpressionContext)
}

func (s *AExp_listElementExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterAExp_listElementExpression(s)
	}
}

func (s *AExp_listElementExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitAExp_listElementExpression(s)
	}
}

func (p *ZZParser) AExpr() (localctx IAExprContext) {
	return p.aExpr(0)
}

func (p *ZZParser) aExpr(_p int) (localctx IAExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewAExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IAExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 12
	p.EnterRecursionRule(localctx, 12, ZZParserRULE_aExpr, _p)
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
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAExp_IntergerLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(107)
			p.Match(ZZParserIntegerLiteral)
		}

	case 2:
		localctx = NewAExp_FloatLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(108)
			p.Match(ZZParserFloatLiteral)
		}

	case 3:
		localctx = NewAExp_IdentifierContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(109)
			p.Match(ZZParserIdentifier)
		}

	case 4:
		localctx = NewAExp_listElementExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(110)
			p.ListElementExpression()
		}

	case 5:
		localctx = NewAExp_bracketExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(111)
			p.Match(ZZParserT__4)
		}
		{
			p.SetState(112)
			p.aExpr(0)
		}
		{
			p.SetState(113)
			p.Match(ZZParserT__5)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(123)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAExp_multiplicativeExpressionContext(p, NewAExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ZZParserRULE_aExpr)
				p.SetState(117)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(118)
					_la = p.GetTokenStream().LA(1)

					if !(_la == ZZParserT__6 || _la == ZZParserT__7) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(119)
					p.aExpr(4)
				}

			case 2:
				localctx = NewAExp_additiveExpressionContext(p, NewAExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ZZParserRULE_aExpr)
				p.SetState(120)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(121)
					_la = p.GetTokenStream().LA(1)

					if !(_la == ZZParserT__8 || _la == ZZParserT__9) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(122)
					p.aExpr(3)
				}

			}

		}
		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// IAExprListContext is an interface to support dynamic dispatch.
type IAExprListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAExprListContext differentiates from other interfaces.
	IsAExprListContext()
}

type AExprListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAExprListContext() *AExprListContext {
	var p = new(AExprListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_aExprList
	return p
}

func (*AExprListContext) IsAExprListContext() {}

func NewAExprListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AExprListContext {
	var p = new(AExprListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_aExprList

	return p
}

func (s *AExprListContext) GetParser() antlr.Parser { return s.parser }

func (s *AExprListContext) AExpr() IAExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAExprContext)
}

func (s *AExprListContext) AExprList() IAExprListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAExprListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAExprListContext)
}

func (s *AExprListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AExprListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AExprListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterAExprList(s)
	}
}

func (s *AExprListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitAExprList(s)
	}
}

func (p *ZZParser) AExprList() (localctx IAExprListContext) {
	return p.aExprList(0)
}

func (p *ZZParser) aExprList(_p int) (localctx IAExprListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewAExprListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IAExprListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 14
	p.EnterRecursionRule(localctx, 14, ZZParserRULE_aExprList, _p)

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
		p.SetState(129)
		p.aExpr(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewAExprListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ZZParserRULE_aExprList)
			p.SetState(131)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(132)
				p.Match(ZZParserT__0)
			}
			{
				p.SetState(133)
				p.aExpr(0)
			}

		}
		p.SetState(138)
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

func (s *BExpContext) AllAExpr() []IAExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAExprContext)(nil)).Elem())
	var tst = make([]IAExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAExprContext)
		}
	}

	return tst
}

func (s *BExpContext) AExpr(i int) IAExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAExprContext)
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
	_startState := 16
	p.EnterRecursionRule(localctx, 16, ZZParserRULE_bExp, _p)
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
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(140)
			p.aExpr(0)
		}
		{
			p.SetState(141)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ZZParserT__10)|(1<<ZZParserT__11)|(1<<ZZParserT__12)|(1<<ZZParserT__13)|(1<<ZZParserT__14)|(1<<ZZParserT__15))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(142)
			p.aExpr(0)
		}

	case 2:
		{
			p.SetState(144)
			p.Match(ZZParserT__18)
		}
		{
			p.SetState(145)
			p.bExp(2)
		}

	case 3:
		{
			p.SetState(146)
			p.Match(ZZParserT__4)
		}
		{
			p.SetState(147)
			p.bExp(0)
		}
		{
			p.SetState(148)
			p.Match(ZZParserT__5)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(157)
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
			p.SetState(152)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
			}
			{
				p.SetState(153)
				_la = p.GetTokenStream().LA(1)

				if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ZZParserT__10)|(1<<ZZParserT__15)|(1<<ZZParserT__16)|(1<<ZZParserT__17))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(154)
				p.bExp(4)
			}

		}
		p.SetState(159)
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
	p.EnterRule(localctx, 18, ZZParserRULE_integerExpression)
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
		p.SetState(160)
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

// IListElementIndexContext is an interface to support dynamic dispatch.
type IListElementIndexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListElementIndexContext differentiates from other interfaces.
	IsListElementIndexContext()
}

type ListElementIndexContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListElementIndexContext() *ListElementIndexContext {
	var p = new(ListElementIndexContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_listElementIndex
	return p
}

func (*ListElementIndexContext) IsListElementIndexContext() {}

func NewListElementIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListElementIndexContext {
	var p = new(ListElementIndexContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_listElementIndex

	return p
}

func (s *ListElementIndexContext) GetParser() antlr.Parser { return s.parser }

func (s *ListElementIndexContext) AExpr() IAExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAExprContext)
}

func (s *ListElementIndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListElementIndexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListElementIndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterListElementIndex(s)
	}
}

func (s *ListElementIndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitListElementIndex(s)
	}
}

func (p *ZZParser) ListElementIndex() (localctx IListElementIndexContext) {
	localctx = NewListElementIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ZZParserRULE_listElementIndex)

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
		p.SetState(162)
		p.Match(ZZParserT__2)
	}
	{
		p.SetState(163)
		p.aExpr(0)
	}
	{
		p.SetState(164)
		p.Match(ZZParserT__3)
	}

	return localctx
}

// IListElementIndexListContext is an interface to support dynamic dispatch.
type IListElementIndexListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListElementIndexListContext differentiates from other interfaces.
	IsListElementIndexListContext()
}

type ListElementIndexListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListElementIndexListContext() *ListElementIndexListContext {
	var p = new(ListElementIndexListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_listElementIndexList
	return p
}

func (*ListElementIndexListContext) IsListElementIndexListContext() {}

func NewListElementIndexListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListElementIndexListContext {
	var p = new(ListElementIndexListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_listElementIndexList

	return p
}

func (s *ListElementIndexListContext) GetParser() antlr.Parser { return s.parser }

func (s *ListElementIndexListContext) ListElementIndex() IListElementIndexContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListElementIndexContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListElementIndexContext)
}

func (s *ListElementIndexListContext) ListElementIndexList() IListElementIndexListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListElementIndexListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListElementIndexListContext)
}

func (s *ListElementIndexListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListElementIndexListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListElementIndexListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterListElementIndexList(s)
	}
}

func (s *ListElementIndexListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitListElementIndexList(s)
	}
}

func (p *ZZParser) ListElementIndexList() (localctx IListElementIndexListContext) {
	return p.listElementIndexList(0)
}

func (p *ZZParser) listElementIndexList(_p int) (localctx IListElementIndexListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewListElementIndexListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListElementIndexListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 22
	p.EnterRecursionRule(localctx, 22, ZZParserRULE_listElementIndexList, _p)

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
		p.SetState(167)
		p.ListElementIndex()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListElementIndexListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ZZParserRULE_listElementIndexList)
			p.SetState(169)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(170)
				p.ListElementIndex()
			}

		}
		p.SetState(175)
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

func (s *ListElementExpressionContext) ListElementIndexList() IListElementIndexListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListElementIndexListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListElementIndexListContext)
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
	p.EnterRule(localctx, 24, ZZParserRULE_listElementExpression)

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
		p.SetState(176)
		p.Match(ZZParserIdentifier)
	}
	{
		p.SetState(177)
		p.listElementIndexList(0)
	}

	return localctx
}

// IAssignInitContext is an interface to support dynamic dispatch.
type IAssignInitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignInitContext differentiates from other interfaces.
	IsAssignInitContext()
}

type AssignInitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignInitContext() *AssignInitContext {
	var p = new(AssignInitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_assignInit
	return p
}

func (*AssignInitContext) IsAssignInitContext() {}

func NewAssignInitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignInitContext {
	var p = new(AssignInitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_assignInit

	return p
}

func (s *AssignInitContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignInitContext) AExpr() IAExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAExprContext)
}

func (s *AssignInitContext) ListInitExpression() IListInitExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListInitExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListInitExpressionContext)
}

func (s *AssignInitContext) FuncInitExpression() IFuncInitExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncInitExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncInitExpressionContext)
}

func (s *AssignInitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignInitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignInitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterAssignInit(s)
	}
}

func (s *AssignInitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitAssignInit(s)
	}
}

func (p *ZZParser) AssignInit() (localctx IAssignInitContext) {
	localctx = NewAssignInitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ZZParserRULE_assignInit)

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

	p.SetState(182)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ZZParserT__4, ZZParserIdentifier, ZZParserIntegerLiteral, ZZParserFloatLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(179)
			p.aExpr(0)
		}

	case ZZParserList:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(180)
			p.ListInitExpression()
		}

	case ZZParserFunc:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(181)
			p.FuncInitExpression()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAssignInitListContext is an interface to support dynamic dispatch.
type IAssignInitListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignInitListContext differentiates from other interfaces.
	IsAssignInitListContext()
}

type AssignInitListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignInitListContext() *AssignInitListContext {
	var p = new(AssignInitListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_assignInitList
	return p
}

func (*AssignInitListContext) IsAssignInitListContext() {}

func NewAssignInitListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignInitListContext {
	var p = new(AssignInitListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_assignInitList

	return p
}

func (s *AssignInitListContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignInitListContext) AssignInit() IAssignInitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignInitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignInitContext)
}

func (s *AssignInitListContext) AssignInitList() IAssignInitListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignInitListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignInitListContext)
}

func (s *AssignInitListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignInitListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignInitListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterAssignInitList(s)
	}
}

func (s *AssignInitListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitAssignInitList(s)
	}
}

func (p *ZZParser) AssignInitList() (localctx IAssignInitListContext) {
	return p.assignInitList(0)
}

func (p *ZZParser) assignInitList(_p int) (localctx IAssignInitListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewAssignInitListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IAssignInitListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 28
	p.EnterRecursionRule(localctx, 28, ZZParserRULE_assignInitList, _p)

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
		p.SetState(185)
		p.AssignInit()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewAssignInitListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ZZParserRULE_assignInitList)
			p.SetState(187)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(188)
				p.Match(ZZParserT__0)
			}
			{
				p.SetState(189)
				p.AssignInit()
			}

		}
		p.SetState(194)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
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

func (s *AssignStatementContext) AssignInitList() IAssignInitListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignInitListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignInitListContext)
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
	p.EnterRule(localctx, 30, ZZParserRULE_assignStatement)

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
		p.SetState(195)
		p.declaratorList(0)
	}
	{
		p.SetState(196)
		p.Match(ZZParserT__19)
	}
	{
		p.SetState(197)
		p.assignInitList(0)
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
	p.EnterRule(localctx, 32, ZZParserRULE_selectionStatement)
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

	p.SetState(233)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ZZParserT__20:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(199)
			p.Match(ZZParserT__20)
		}
		{
			p.SetState(200)
			p.bExp(0)
		}
		{
			p.SetState(201)
			p.Match(ZZParserT__21)
		}
		p.SetState(203)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ZZParserT__4)|(1<<ZZParserT__18)|(1<<ZZParserT__20)|(1<<ZZParserT__27)|(1<<ZZParserT__29))) != 0) || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(ZZParserIdentifier-37))|(1<<(ZZParserIntegerLiteral-37))|(1<<(ZZParserFloatLiteral-37)))) != 0) {
			{
				p.SetState(202)
				p.funcStatementList(0)
			}

		}
		{
			p.SetState(205)
			p.Match(ZZParserT__22)
		}
		p.SetState(216)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(206)
					p.Match(ZZParserT__23)
				}
				{
					p.SetState(207)
					p.bExp(0)
				}
				{
					p.SetState(208)
					p.Match(ZZParserT__21)
				}
				p.SetState(210)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ZZParserT__4)|(1<<ZZParserT__18)|(1<<ZZParserT__20)|(1<<ZZParserT__27)|(1<<ZZParserT__29))) != 0) || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(ZZParserIdentifier-37))|(1<<(ZZParserIntegerLiteral-37))|(1<<(ZZParserFloatLiteral-37)))) != 0) {
					{
						p.SetState(209)
						p.funcStatementList(0)
					}

				}
				{
					p.SetState(212)
					p.Match(ZZParserT__22)
				}

			}
			p.SetState(218)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())
		}
		p.SetState(225)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(219)
				p.Match(ZZParserT__24)
			}
			{
				p.SetState(220)
				p.Match(ZZParserT__21)
			}
			p.SetState(222)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ZZParserT__4)|(1<<ZZParserT__18)|(1<<ZZParserT__20)|(1<<ZZParserT__27)|(1<<ZZParserT__29))) != 0) || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(ZZParserIdentifier-37))|(1<<(ZZParserIntegerLiteral-37))|(1<<(ZZParserFloatLiteral-37)))) != 0) {
				{
					p.SetState(221)
					p.funcStatementList(0)
				}

			}
			{
				p.SetState(224)
				p.Match(ZZParserT__22)
			}

		}

	case ZZParserT__4, ZZParserT__18, ZZParserIdentifier, ZZParserIntegerLiteral, ZZParserFloatLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(227)
			p.bExp(0)
		}
		{
			p.SetState(228)
			p.Match(ZZParserT__25)
		}
		{
			p.SetState(229)
			p.FuncStatement()
		}
		{
			p.SetState(230)
			p.Match(ZZParserT__26)
		}
		{
			p.SetState(231)
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
	p.EnterRule(localctx, 34, ZZParserRULE_iterationStatement)
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

	p.SetState(259)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(235)
			p.Match(ZZParserT__27)
		}
		p.SetState(237)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ZZParserT__4 || _la == ZZParserT__18 || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(ZZParserIdentifier-37))|(1<<(ZZParserIntegerLiteral-37))|(1<<(ZZParserFloatLiteral-37)))) != 0) {
			{
				p.SetState(236)
				p.bExp(0)
			}

		}
		{
			p.SetState(239)
			p.Match(ZZParserT__21)
		}
		{
			p.SetState(240)
			p.funcStatementList(0)
		}
		{
			p.SetState(241)
			p.Match(ZZParserT__22)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(243)
			p.Match(ZZParserT__27)
		}
		p.SetState(245)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ZZParserIdentifier {
			{
				p.SetState(244)
				p.AssignStatement()
			}

		}
		{
			p.SetState(247)
			p.Match(ZZParserT__28)
		}
		p.SetState(249)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ZZParserT__4 || _la == ZZParserT__18 || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(ZZParserIdentifier-37))|(1<<(ZZParserIntegerLiteral-37))|(1<<(ZZParserFloatLiteral-37)))) != 0) {
			{
				p.SetState(248)
				p.bExp(0)
			}

		}
		{
			p.SetState(251)
			p.Match(ZZParserT__28)
		}
		p.SetState(253)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ZZParserIdentifier {
			{
				p.SetState(252)
				p.AssignStatement()
			}

		}
		{
			p.SetState(255)
			p.Match(ZZParserT__21)
		}
		{
			p.SetState(256)
			p.funcStatementList(0)
		}
		{
			p.SetState(257)
			p.Match(ZZParserT__22)
		}

	}

	return localctx
}

// IDefinitionContext is an interface to support dynamic dispatch.
type IDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefinitionContext differentiates from other interfaces.
	IsDefinitionContext()
}

type DefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefinitionContext() *DefinitionContext {
	var p = new(DefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_definition
	return p
}

func (*DefinitionContext) IsDefinitionContext() {}

func NewDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefinitionContext {
	var p = new(DefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_definition

	return p
}

func (s *DefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *DefinitionContext) AssignStatement() IAssignStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignStatementContext)
}

func (s *DefinitionContext) FuncDefinition() IFuncDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncDefinitionContext)
}

func (s *DefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterDefinition(s)
	}
}

func (s *DefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitDefinition(s)
	}
}

func (p *ZZParser) Definition() (localctx IDefinitionContext) {
	localctx = NewDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ZZParserRULE_definition)

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

	p.SetState(263)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ZZParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(261)
			p.AssignStatement()
		}

	case ZZParserFunc:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(262)
			p.FuncDefinition()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDefinitionListContext is an interface to support dynamic dispatch.
type IDefinitionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefinitionListContext differentiates from other interfaces.
	IsDefinitionListContext()
}

type DefinitionListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefinitionListContext() *DefinitionListContext {
	var p = new(DefinitionListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_definitionList
	return p
}

func (*DefinitionListContext) IsDefinitionListContext() {}

func NewDefinitionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefinitionListContext {
	var p = new(DefinitionListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_definitionList

	return p
}

func (s *DefinitionListContext) GetParser() antlr.Parser { return s.parser }

func (s *DefinitionListContext) Definition() IDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefinitionContext)
}

func (s *DefinitionListContext) DefinitionList() IDefinitionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefinitionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefinitionListContext)
}

func (s *DefinitionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefinitionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefinitionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterDefinitionList(s)
	}
}

func (s *DefinitionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitDefinitionList(s)
	}
}

func (p *ZZParser) DefinitionList() (localctx IDefinitionListContext) {
	return p.definitionList(0)
}

func (p *ZZParser) definitionList(_p int) (localctx IDefinitionListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewDefinitionListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IDefinitionListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 38
	p.EnterRecursionRule(localctx, 38, ZZParserRULE_definitionList, _p)

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
		p.SetState(266)
		p.Definition()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(272)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewDefinitionListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ZZParserRULE_definitionList)
			p.SetState(268)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(269)
				p.Definition()
			}

		}
		p.SetState(274)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())
	}

	return localctx
}

// IFileContext is an interface to support dynamic dispatch.
type IFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFileContext differentiates from other interfaces.
	IsFileContext()
}

type FileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFileContext() *FileContext {
	var p = new(FileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_file
	return p
}

func (*FileContext) IsFileContext() {}

func NewFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileContext {
	var p = new(FileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_file

	return p
}

func (s *FileContext) GetParser() antlr.Parser { return s.parser }

func (s *FileContext) DeclaratorList() IDeclaratorListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclaratorListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclaratorListContext)
}

func (s *FileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterFile(s)
	}
}

func (s *FileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitFile(s)
	}
}

func (p *ZZParser) File() (localctx IFileContext) {
	localctx = NewFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ZZParserRULE_file)

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
		p.SetState(275)
		p.declaratorList(0)
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

func (s *TypeSpecifierContext) FuncTypeSpecifier() IFuncTypeSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncTypeSpecifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncTypeSpecifierContext)
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
	p.EnterRule(localctx, 42, ZZParserRULE_typeSpecifier)

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

	p.SetState(280)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ZZParserT__1, ZZParserInt, ZZParserFloat:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(277)
			p.SimpleTypeSpecifier()
		}

	case ZZParserT__2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(278)
			p.ListTypeSpecifier()
		}

	case ZZParserFunc:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(279)
			p.FuncTypeSpecifier()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	_startState := 44
	p.EnterRecursionRule(localctx, 44, ZZParserRULE_typeSpecifierList, _p)

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
		p.SetState(283)
		p.TypeSpecifier()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(290)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewTypeSpecifierListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ZZParserRULE_typeSpecifierList)
			p.SetState(285)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(286)
				p.Match(ZZParserT__0)
			}
			{
				p.SetState(287)
				p.TypeSpecifier()
			}

		}
		p.SetState(292)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())
	}

	return localctx
}

// IParaDeclaratorListContext is an interface to support dynamic dispatch.
type IParaDeclaratorListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParaDeclaratorListContext differentiates from other interfaces.
	IsParaDeclaratorListContext()
}

type ParaDeclaratorListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParaDeclaratorListContext() *ParaDeclaratorListContext {
	var p = new(ParaDeclaratorListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_paraDeclaratorList
	return p
}

func (*ParaDeclaratorListContext) IsParaDeclaratorListContext() {}

func NewParaDeclaratorListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParaDeclaratorListContext {
	var p = new(ParaDeclaratorListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_paraDeclaratorList

	return p
}

func (s *ParaDeclaratorListContext) GetParser() antlr.Parser { return s.parser }

func (s *ParaDeclaratorListContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ZZParserIdentifier, 0)
}

func (s *ParaDeclaratorListContext) ParaDeclaratorList() IParaDeclaratorListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParaDeclaratorListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParaDeclaratorListContext)
}

func (s *ParaDeclaratorListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParaDeclaratorListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParaDeclaratorListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterParaDeclaratorList(s)
	}
}

func (s *ParaDeclaratorListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitParaDeclaratorList(s)
	}
}

func (p *ZZParser) ParaDeclaratorList() (localctx IParaDeclaratorListContext) {
	return p.paraDeclaratorList(0)
}

func (p *ZZParser) paraDeclaratorList(_p int) (localctx IParaDeclaratorListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewParaDeclaratorListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IParaDeclaratorListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 46
	p.EnterRecursionRule(localctx, 46, ZZParserRULE_paraDeclaratorList, _p)

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
		p.SetState(294)
		p.Match(ZZParserIdentifier)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(301)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewParaDeclaratorListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ZZParserRULE_paraDeclaratorList)
			p.SetState(296)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(297)
				p.Match(ZZParserT__0)
			}
			{
				p.SetState(298)
				p.Match(ZZParserIdentifier)
			}

		}
		p.SetState(303)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())
	}

	return localctx
}

// IParaDeclaratorWithIdentityContext is an interface to support dynamic dispatch.
type IParaDeclaratorWithIdentityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParaDeclaratorWithIdentityContext differentiates from other interfaces.
	IsParaDeclaratorWithIdentityContext()
}

type ParaDeclaratorWithIdentityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParaDeclaratorWithIdentityContext() *ParaDeclaratorWithIdentityContext {
	var p = new(ParaDeclaratorWithIdentityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_paraDeclaratorWithIdentity
	return p
}

func (*ParaDeclaratorWithIdentityContext) IsParaDeclaratorWithIdentityContext() {}

func NewParaDeclaratorWithIdentityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParaDeclaratorWithIdentityContext {
	var p = new(ParaDeclaratorWithIdentityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_paraDeclaratorWithIdentity

	return p
}

func (s *ParaDeclaratorWithIdentityContext) GetParser() antlr.Parser { return s.parser }

func (s *ParaDeclaratorWithIdentityContext) DeclaratorList() IDeclaratorListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclaratorListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclaratorListContext)
}

func (s *ParaDeclaratorWithIdentityContext) TypeSpecifier() ITypeSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierContext)
}

func (s *ParaDeclaratorWithIdentityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParaDeclaratorWithIdentityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParaDeclaratorWithIdentityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterParaDeclaratorWithIdentity(s)
	}
}

func (s *ParaDeclaratorWithIdentityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitParaDeclaratorWithIdentity(s)
	}
}

func (p *ZZParser) ParaDeclaratorWithIdentity() (localctx IParaDeclaratorWithIdentityContext) {
	localctx = NewParaDeclaratorWithIdentityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, ZZParserRULE_paraDeclaratorWithIdentity)

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
		p.SetState(304)
		p.declaratorList(0)
	}
	{
		p.SetState(305)
		p.TypeSpecifier()
	}

	return localctx
}

// IParaDeclaratorWithIdentityListContext is an interface to support dynamic dispatch.
type IParaDeclaratorWithIdentityListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParaDeclaratorWithIdentityListContext differentiates from other interfaces.
	IsParaDeclaratorWithIdentityListContext()
}

type ParaDeclaratorWithIdentityListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParaDeclaratorWithIdentityListContext() *ParaDeclaratorWithIdentityListContext {
	var p = new(ParaDeclaratorWithIdentityListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_paraDeclaratorWithIdentityList
	return p
}

func (*ParaDeclaratorWithIdentityListContext) IsParaDeclaratorWithIdentityListContext() {}

func NewParaDeclaratorWithIdentityListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParaDeclaratorWithIdentityListContext {
	var p = new(ParaDeclaratorWithIdentityListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_paraDeclaratorWithIdentityList

	return p
}

func (s *ParaDeclaratorWithIdentityListContext) GetParser() antlr.Parser { return s.parser }

func (s *ParaDeclaratorWithIdentityListContext) ParaDeclaratorWithIdentity() IParaDeclaratorWithIdentityContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParaDeclaratorWithIdentityContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParaDeclaratorWithIdentityContext)
}

func (s *ParaDeclaratorWithIdentityListContext) ParaDeclaratorWithIdentityList() IParaDeclaratorWithIdentityListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParaDeclaratorWithIdentityListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParaDeclaratorWithIdentityListContext)
}

func (s *ParaDeclaratorWithIdentityListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParaDeclaratorWithIdentityListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParaDeclaratorWithIdentityListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterParaDeclaratorWithIdentityList(s)
	}
}

func (s *ParaDeclaratorWithIdentityListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitParaDeclaratorWithIdentityList(s)
	}
}

func (p *ZZParser) ParaDeclaratorWithIdentityList() (localctx IParaDeclaratorWithIdentityListContext) {
	return p.paraDeclaratorWithIdentityList(0)
}

func (p *ZZParser) paraDeclaratorWithIdentityList(_p int) (localctx IParaDeclaratorWithIdentityListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewParaDeclaratorWithIdentityListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IParaDeclaratorWithIdentityListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 50
	p.EnterRecursionRule(localctx, 50, ZZParserRULE_paraDeclaratorWithIdentityList, _p)

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
		p.SetState(308)
		p.ParaDeclaratorWithIdentity()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(315)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewParaDeclaratorWithIdentityListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ZZParserRULE_paraDeclaratorWithIdentityList)
			p.SetState(310)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(311)
				p.Match(ZZParserT__0)
			}
			{
				p.SetState(312)
				p.ParaDeclaratorWithIdentity()
			}

		}
		p.SetState(317)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())
	}

	return localctx
}

// IFuncTypeSpecifierContext is an interface to support dynamic dispatch.
type IFuncTypeSpecifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncTypeSpecifierContext differentiates from other interfaces.
	IsFuncTypeSpecifierContext()
}

type FuncTypeSpecifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncTypeSpecifierContext() *FuncTypeSpecifierContext {
	var p = new(FuncTypeSpecifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_funcTypeSpecifier
	return p
}

func (*FuncTypeSpecifierContext) IsFuncTypeSpecifierContext() {}

func NewFuncTypeSpecifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncTypeSpecifierContext {
	var p = new(FuncTypeSpecifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_funcTypeSpecifier

	return p
}

func (s *FuncTypeSpecifierContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncTypeSpecifierContext) Func() antlr.TerminalNode {
	return s.GetToken(ZZParserFunc, 0)
}

func (s *FuncTypeSpecifierContext) ParaDeclaratorWithIdentityList() IParaDeclaratorWithIdentityListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParaDeclaratorWithIdentityListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParaDeclaratorWithIdentityListContext)
}

func (s *FuncTypeSpecifierContext) TypeSpecifierList() ITypeSpecifierListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecifierListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierListContext)
}

func (s *FuncTypeSpecifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncTypeSpecifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncTypeSpecifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterFuncTypeSpecifier(s)
	}
}

func (s *FuncTypeSpecifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitFuncTypeSpecifier(s)
	}
}

func (p *ZZParser) FuncTypeSpecifier() (localctx IFuncTypeSpecifierContext) {
	localctx = NewFuncTypeSpecifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, ZZParserRULE_funcTypeSpecifier)
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
		p.SetState(318)
		p.Match(ZZParserFunc)
	}
	{
		p.SetState(319)
		p.Match(ZZParserT__4)
	}
	p.SetState(321)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ZZParserIdentifier {
		{
			p.SetState(320)
			p.paraDeclaratorWithIdentityList(0)
		}

	}
	{
		p.SetState(323)
		p.Match(ZZParserT__5)
	}
	p.SetState(329)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(324)
			p.Match(ZZParserT__4)
		}
		p.SetState(326)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ZZParserT__1 || _la == ZZParserT__2 || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(ZZParserInt-32))|(1<<(ZZParserFloat-32))|(1<<(ZZParserFunc-32)))) != 0) {
			{
				p.SetState(325)
				p.typeSpecifierList(0)
			}

		}
		{
			p.SetState(328)
			p.Match(ZZParserT__5)
		}

	}

	return localctx
}

// IFuncTypeSpecifierWithNameContext is an interface to support dynamic dispatch.
type IFuncTypeSpecifierWithNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncTypeSpecifierWithNameContext differentiates from other interfaces.
	IsFuncTypeSpecifierWithNameContext()
}

type FuncTypeSpecifierWithNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncTypeSpecifierWithNameContext() *FuncTypeSpecifierWithNameContext {
	var p = new(FuncTypeSpecifierWithNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_funcTypeSpecifierWithName
	return p
}

func (*FuncTypeSpecifierWithNameContext) IsFuncTypeSpecifierWithNameContext() {}

func NewFuncTypeSpecifierWithNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncTypeSpecifierWithNameContext {
	var p = new(FuncTypeSpecifierWithNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_funcTypeSpecifierWithName

	return p
}

func (s *FuncTypeSpecifierWithNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncTypeSpecifierWithNameContext) Func() antlr.TerminalNode {
	return s.GetToken(ZZParserFunc, 0)
}

func (s *FuncTypeSpecifierWithNameContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ZZParserIdentifier, 0)
}

func (s *FuncTypeSpecifierWithNameContext) ParaDeclaratorWithIdentityList() IParaDeclaratorWithIdentityListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParaDeclaratorWithIdentityListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParaDeclaratorWithIdentityListContext)
}

func (s *FuncTypeSpecifierWithNameContext) TypeSpecifierList() ITypeSpecifierListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecifierListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierListContext)
}

func (s *FuncTypeSpecifierWithNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncTypeSpecifierWithNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncTypeSpecifierWithNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterFuncTypeSpecifierWithName(s)
	}
}

func (s *FuncTypeSpecifierWithNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitFuncTypeSpecifierWithName(s)
	}
}

func (p *ZZParser) FuncTypeSpecifierWithName() (localctx IFuncTypeSpecifierWithNameContext) {
	localctx = NewFuncTypeSpecifierWithNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, ZZParserRULE_funcTypeSpecifierWithName)
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
		p.SetState(331)
		p.Match(ZZParserFunc)
	}
	{
		p.SetState(332)
		p.Match(ZZParserIdentifier)
	}
	{
		p.SetState(333)
		p.Match(ZZParserT__4)
	}
	p.SetState(335)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ZZParserIdentifier {
		{
			p.SetState(334)
			p.paraDeclaratorWithIdentityList(0)
		}

	}
	{
		p.SetState(337)
		p.Match(ZZParserT__5)
	}
	p.SetState(343)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ZZParserT__4 {
		{
			p.SetState(338)
			p.Match(ZZParserT__4)
		}
		p.SetState(340)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ZZParserT__1 || _la == ZZParserT__2 || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(ZZParserInt-32))|(1<<(ZZParserFloat-32))|(1<<(ZZParserFunc-32)))) != 0) {
			{
				p.SetState(339)
				p.typeSpecifierList(0)
			}

		}
		{
			p.SetState(342)
			p.Match(ZZParserT__5)
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

func (s *FuncReturnParaContext) AExpr() IAExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAExprContext)
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
	p.EnterRule(localctx, 56, ZZParserRULE_funcReturnPara)

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

	p.SetState(349)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(345)
			p.aExpr(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(346)
			p.bExp(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(347)
			p.Match(ZZParserIdentifier)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(348)
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
	_startState := 58
	p.EnterRecursionRule(localctx, 58, ZZParserRULE_funcReturnParaList, _p)

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
		p.SetState(352)
		p.FuncReturnPara()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(359)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewFuncReturnParaListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ZZParserRULE_funcReturnParaList)
			p.SetState(354)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(355)
				p.Match(ZZParserT__0)
			}
			{
				p.SetState(356)
				p.FuncReturnPara()
			}

		}
		p.SetState(361)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 60, ZZParserRULE_funcStatement)

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

	p.SetState(369)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(362)
			p.AssignStatement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(363)
			p.SelectionStatement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(364)
			p.IterationStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(365)
			p.Match(ZZParserT__29)
		}
		p.SetState(367)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(366)
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
	_startState := 62
	p.EnterRecursionRule(localctx, 62, ZZParserRULE_funcStatementList, _p)

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
		p.SetState(372)
		p.FuncStatement()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(378)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewFuncStatementListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ZZParserRULE_funcStatementList)
			p.SetState(374)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(375)
				p.FuncStatement()
			}

		}
		p.SetState(380)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 64, ZZParserRULE_funcBody)
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
	p.SetState(382)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ZZParserT__4)|(1<<ZZParserT__18)|(1<<ZZParserT__20)|(1<<ZZParserT__27)|(1<<ZZParserT__29))) != 0) || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(ZZParserIdentifier-37))|(1<<(ZZParserIntegerLiteral-37))|(1<<(ZZParserFloatLiteral-37)))) != 0) {
		{
			p.SetState(381)
			p.funcStatementList(0)
		}

	}

	return localctx
}

// IFuncInitExpressionContext is an interface to support dynamic dispatch.
type IFuncInitExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncInitExpressionContext differentiates from other interfaces.
	IsFuncInitExpressionContext()
}

type FuncInitExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncInitExpressionContext() *FuncInitExpressionContext {
	var p = new(FuncInitExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_funcInitExpression
	return p
}

func (*FuncInitExpressionContext) IsFuncInitExpressionContext() {}

func NewFuncInitExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncInitExpressionContext {
	var p = new(FuncInitExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_funcInitExpression

	return p
}

func (s *FuncInitExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncInitExpressionContext) FuncTypeSpecifier() IFuncTypeSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncTypeSpecifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncTypeSpecifierContext)
}

func (s *FuncInitExpressionContext) FuncBody() IFuncBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncBodyContext)
}

func (s *FuncInitExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncInitExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncInitExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterFuncInitExpression(s)
	}
}

func (s *FuncInitExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitFuncInitExpression(s)
	}
}

func (p *ZZParser) FuncInitExpression() (localctx IFuncInitExpressionContext) {
	localctx = NewFuncInitExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, ZZParserRULE_funcInitExpression)

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
		p.SetState(384)
		p.FuncTypeSpecifier()
	}
	{
		p.SetState(385)
		p.Match(ZZParserT__21)
	}
	p.SetState(387)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(386)
			p.FuncBody()
		}

	}
	{
		p.SetState(389)
		p.Match(ZZParserT__22)
	}

	return localctx
}

// IFuncDefinitionContext is an interface to support dynamic dispatch.
type IFuncDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncDefinitionContext differentiates from other interfaces.
	IsFuncDefinitionContext()
}

type FuncDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncDefinitionContext() *FuncDefinitionContext {
	var p = new(FuncDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_funcDefinition
	return p
}

func (*FuncDefinitionContext) IsFuncDefinitionContext() {}

func NewFuncDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncDefinitionContext {
	var p = new(FuncDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_funcDefinition

	return p
}

func (s *FuncDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncDefinitionContext) FuncTypeSpecifierWithName() IFuncTypeSpecifierWithNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncTypeSpecifierWithNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncTypeSpecifierWithNameContext)
}

func (s *FuncDefinitionContext) FuncBody() IFuncBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncBodyContext)
}

func (s *FuncDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterFuncDefinition(s)
	}
}

func (s *FuncDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitFuncDefinition(s)
	}
}

func (p *ZZParser) FuncDefinition() (localctx IFuncDefinitionContext) {
	localctx = NewFuncDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, ZZParserRULE_funcDefinition)

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
		p.SetState(391)
		p.FuncTypeSpecifierWithName()
	}
	{
		p.SetState(392)
		p.Match(ZZParserT__21)
	}
	p.SetState(394)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(393)
			p.FuncBody()
		}

	}
	{
		p.SetState(396)
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

	case 6:
		var t *AExprContext = nil
		if localctx != nil {
			t = localctx.(*AExprContext)
		}
		return p.AExpr_Sempred(t, predIndex)

	case 7:
		var t *AExprListContext = nil
		if localctx != nil {
			t = localctx.(*AExprListContext)
		}
		return p.AExprList_Sempred(t, predIndex)

	case 8:
		var t *BExpContext = nil
		if localctx != nil {
			t = localctx.(*BExpContext)
		}
		return p.BExp_Sempred(t, predIndex)

	case 11:
		var t *ListElementIndexListContext = nil
		if localctx != nil {
			t = localctx.(*ListElementIndexListContext)
		}
		return p.ListElementIndexList_Sempred(t, predIndex)

	case 14:
		var t *AssignInitListContext = nil
		if localctx != nil {
			t = localctx.(*AssignInitListContext)
		}
		return p.AssignInitList_Sempred(t, predIndex)

	case 19:
		var t *DefinitionListContext = nil
		if localctx != nil {
			t = localctx.(*DefinitionListContext)
		}
		return p.DefinitionList_Sempred(t, predIndex)

	case 22:
		var t *TypeSpecifierListContext = nil
		if localctx != nil {
			t = localctx.(*TypeSpecifierListContext)
		}
		return p.TypeSpecifierList_Sempred(t, predIndex)

	case 23:
		var t *ParaDeclaratorListContext = nil
		if localctx != nil {
			t = localctx.(*ParaDeclaratorListContext)
		}
		return p.ParaDeclaratorList_Sempred(t, predIndex)

	case 25:
		var t *ParaDeclaratorWithIdentityListContext = nil
		if localctx != nil {
			t = localctx.(*ParaDeclaratorWithIdentityListContext)
		}
		return p.ParaDeclaratorWithIdentityList_Sempred(t, predIndex)

	case 29:
		var t *FuncReturnParaListContext = nil
		if localctx != nil {
			t = localctx.(*FuncReturnParaListContext)
		}
		return p.FuncReturnParaList_Sempred(t, predIndex)

	case 31:
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

func (p *ZZParser) AExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ZZParser) AExprList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
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

func (p *ZZParser) ListElementIndexList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 5:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ZZParser) AssignInitList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 6:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ZZParser) DefinitionList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
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

func (p *ZZParser) ParaDeclaratorList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 9:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ZZParser) ParaDeclaratorWithIdentityList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 10:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ZZParser) FuncReturnParaList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 11:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ZZParser) FuncStatementList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 12:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
