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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 42, 448,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 3, 2, 3, 2, 3,
	3, 3, 3, 5, 3, 101, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 109,
	10, 4, 12, 4, 14, 4, 112, 11, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 5,
	6, 120, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 142,
	10, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 150, 10, 9, 12, 9, 14,
	9, 153, 11, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 5, 10, 166, 10, 10, 3, 10, 3, 10, 3, 10, 7, 10, 171,
	10, 10, 12, 10, 14, 10, 174, 11, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 185, 10, 12, 12, 12, 14, 12, 188, 11,
	12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 197, 10, 14,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 7, 15, 205, 10, 15, 12, 15, 14,
	15, 208, 11, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17,
	5, 17, 218, 10, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 226,
	10, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 5, 19, 233, 10, 19, 3, 19, 3,
	19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 7, 22, 245,
	10, 22, 12, 22, 14, 22, 248, 11, 22, 3, 22, 5, 22, 251, 10, 22, 3, 22,
	3, 22, 3, 22, 3, 22, 5, 22, 257, 10, 22, 3, 23, 3, 23, 5, 23, 261, 10,
	23, 3, 24, 3, 24, 5, 24, 265, 10, 24, 3, 24, 3, 24, 5, 24, 269, 10, 24,
	3, 24, 3, 24, 5, 24, 273, 10, 24, 3, 24, 3, 24, 5, 24, 277, 10, 24, 3,
	24, 3, 24, 3, 25, 3, 25, 5, 25, 283, 10, 25, 3, 26, 3, 26, 3, 26, 3, 26,
	3, 26, 7, 26, 290, 10, 26, 12, 26, 14, 26, 293, 11, 26, 3, 27, 3, 27, 3,
	28, 3, 28, 3, 28, 5, 28, 300, 10, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29,
	3, 29, 7, 29, 308, 10, 29, 12, 29, 14, 29, 311, 11, 29, 3, 30, 3, 30, 3,
	31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 7, 31, 321, 10, 31, 12, 31, 14,
	31, 324, 11, 31, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33,
	3, 33, 7, 33, 335, 10, 33, 12, 33, 14, 33, 338, 11, 33, 3, 34, 3, 34, 3,
	34, 5, 34, 343, 10, 34, 3, 34, 3, 34, 3, 34, 5, 34, 348, 10, 34, 3, 34,
	5, 34, 351, 10, 34, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 36, 5, 36, 359,
	10, 36, 3, 36, 3, 36, 3, 36, 5, 36, 364, 10, 36, 3, 36, 5, 36, 367, 10,
	36, 3, 37, 3, 37, 3, 37, 3, 37, 5, 37, 373, 10, 37, 3, 38, 3, 38, 3, 38,
	3, 38, 3, 38, 3, 38, 7, 38, 381, 10, 38, 12, 38, 14, 38, 384, 11, 38, 3,
	39, 3, 39, 5, 39, 388, 10, 39, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 5, 40,
	395, 10, 40, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 7, 41, 402, 10, 41, 12,
	41, 14, 41, 405, 11, 41, 3, 42, 5, 42, 408, 10, 42, 3, 43, 3, 43, 3, 43,
	5, 43, 413, 10, 43, 3, 43, 3, 43, 3, 44, 3, 44, 3, 44, 5, 44, 420, 10,
	44, 3, 44, 3, 44, 3, 45, 3, 45, 5, 45, 426, 10, 45, 3, 46, 3, 46, 3, 46,
	3, 46, 3, 46, 3, 46, 7, 46, 434, 10, 46, 12, 46, 14, 46, 437, 11, 46, 3,
	47, 3, 47, 3, 47, 5, 47, 442, 10, 47, 3, 47, 3, 47, 3, 48, 3, 48, 3, 48,
	2, 14, 6, 16, 18, 22, 28, 50, 56, 60, 64, 74, 80, 90, 49, 2, 4, 6, 8, 10,
	12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46,
	48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82,
	84, 86, 88, 90, 92, 94, 2, 7, 4, 2, 4, 4, 34, 35, 3, 2, 9, 10, 3, 2, 11,
	12, 3, 2, 13, 18, 4, 2, 13, 13, 18, 20, 2, 457, 2, 96, 3, 2, 2, 2, 4, 100,
	3, 2, 2, 2, 6, 102, 3, 2, 2, 2, 8, 113, 3, 2, 2, 2, 10, 119, 3, 2, 2, 2,
	12, 121, 3, 2, 2, 2, 14, 125, 3, 2, 2, 2, 16, 141, 3, 2, 2, 2, 18, 165,
	3, 2, 2, 2, 20, 175, 3, 2, 2, 2, 22, 179, 3, 2, 2, 2, 24, 189, 3, 2, 2,
	2, 26, 196, 3, 2, 2, 2, 28, 198, 3, 2, 2, 2, 30, 209, 3, 2, 2, 2, 32, 213,
	3, 2, 2, 2, 34, 221, 3, 2, 2, 2, 36, 229, 3, 2, 2, 2, 38, 236, 3, 2, 2,
	2, 40, 240, 3, 2, 2, 2, 42, 256, 3, 2, 2, 2, 44, 260, 3, 2, 2, 2, 46, 262,
	3, 2, 2, 2, 48, 282, 3, 2, 2, 2, 50, 284, 3, 2, 2, 2, 52, 294, 3, 2, 2,
	2, 54, 299, 3, 2, 2, 2, 56, 301, 3, 2, 2, 2, 58, 312, 3, 2, 2, 2, 60, 314,
	3, 2, 2, 2, 62, 325, 3, 2, 2, 2, 64, 328, 3, 2, 2, 2, 66, 339, 3, 2, 2,
	2, 68, 352, 3, 2, 2, 2, 70, 354, 3, 2, 2, 2, 72, 372, 3, 2, 2, 2, 74, 374,
	3, 2, 2, 2, 76, 385, 3, 2, 2, 2, 78, 394, 3, 2, 2, 2, 80, 396, 3, 2, 2,
	2, 82, 407, 3, 2, 2, 2, 84, 409, 3, 2, 2, 2, 86, 416, 3, 2, 2, 2, 88, 425,
	3, 2, 2, 2, 90, 427, 3, 2, 2, 2, 92, 438, 3, 2, 2, 2, 94, 445, 3, 2, 2,
	2, 96, 97, 7, 39, 2, 2, 97, 3, 3, 2, 2, 2, 98, 101, 5, 2, 2, 2, 99, 101,
	5, 24, 13, 2, 100, 98, 3, 2, 2, 2, 100, 99, 3, 2, 2, 2, 101, 5, 3, 2, 2,
	2, 102, 103, 8, 4, 1, 2, 103, 104, 5, 4, 3, 2, 104, 110, 3, 2, 2, 2, 105,
	106, 12, 3, 2, 2, 106, 107, 7, 3, 2, 2, 107, 109, 5, 4, 3, 2, 108, 105,
	3, 2, 2, 2, 109, 112, 3, 2, 2, 2, 110, 108, 3, 2, 2, 2, 110, 111, 3, 2,
	2, 2, 111, 7, 3, 2, 2, 2, 112, 110, 3, 2, 2, 2, 113, 114, 9, 2, 2, 2, 114,
	9, 3, 2, 2, 2, 115, 120, 5, 8, 5, 2, 116, 117, 7, 5, 2, 2, 117, 118, 7,
	6, 2, 2, 118, 120, 5, 10, 6, 2, 119, 115, 3, 2, 2, 2, 119, 116, 3, 2, 2,
	2, 120, 11, 3, 2, 2, 2, 121, 122, 7, 5, 2, 2, 122, 123, 7, 6, 2, 2, 123,
	124, 5, 10, 6, 2, 124, 13, 3, 2, 2, 2, 125, 126, 7, 37, 2, 2, 126, 127,
	7, 7, 2, 2, 127, 128, 5, 12, 7, 2, 128, 129, 7, 3, 2, 2, 129, 130, 5, 16,
	9, 2, 130, 131, 7, 8, 2, 2, 131, 15, 3, 2, 2, 2, 132, 133, 8, 9, 1, 2,
	133, 142, 7, 40, 2, 2, 134, 142, 7, 41, 2, 2, 135, 142, 5, 2, 2, 2, 136,
	142, 5, 24, 13, 2, 137, 138, 7, 7, 2, 2, 138, 139, 5, 16, 9, 2, 139, 140,
	7, 8, 2, 2, 140, 142, 3, 2, 2, 2, 141, 132, 3, 2, 2, 2, 141, 134, 3, 2,
	2, 2, 141, 135, 3, 2, 2, 2, 141, 136, 3, 2, 2, 2, 141, 137, 3, 2, 2, 2,
	142, 151, 3, 2, 2, 2, 143, 144, 12, 5, 2, 2, 144, 145, 9, 3, 2, 2, 145,
	150, 5, 16, 9, 6, 146, 147, 12, 4, 2, 2, 147, 148, 9, 4, 2, 2, 148, 150,
	5, 16, 9, 5, 149, 143, 3, 2, 2, 2, 149, 146, 3, 2, 2, 2, 150, 153, 3, 2,
	2, 2, 151, 149, 3, 2, 2, 2, 151, 152, 3, 2, 2, 2, 152, 17, 3, 2, 2, 2,
	153, 151, 3, 2, 2, 2, 154, 155, 8, 10, 1, 2, 155, 156, 5, 16, 9, 2, 156,
	157, 9, 5, 2, 2, 157, 158, 5, 16, 9, 2, 158, 166, 3, 2, 2, 2, 159, 160,
	7, 21, 2, 2, 160, 166, 5, 18, 10, 4, 161, 162, 7, 7, 2, 2, 162, 163, 5,
	18, 10, 2, 163, 164, 7, 8, 2, 2, 164, 166, 3, 2, 2, 2, 165, 154, 3, 2,
	2, 2, 165, 159, 3, 2, 2, 2, 165, 161, 3, 2, 2, 2, 166, 172, 3, 2, 2, 2,
	167, 168, 12, 5, 2, 2, 168, 169, 9, 6, 2, 2, 169, 171, 5, 18, 10, 6, 170,
	167, 3, 2, 2, 2, 171, 174, 3, 2, 2, 2, 172, 170, 3, 2, 2, 2, 172, 173,
	3, 2, 2, 2, 173, 19, 3, 2, 2, 2, 174, 172, 3, 2, 2, 2, 175, 176, 7, 5,
	2, 2, 176, 177, 5, 16, 9, 2, 177, 178, 7, 6, 2, 2, 178, 21, 3, 2, 2, 2,
	179, 180, 8, 12, 1, 2, 180, 181, 5, 20, 11, 2, 181, 186, 3, 2, 2, 2, 182,
	183, 12, 3, 2, 2, 183, 185, 5, 20, 11, 2, 184, 182, 3, 2, 2, 2, 185, 188,
	3, 2, 2, 2, 186, 184, 3, 2, 2, 2, 186, 187, 3, 2, 2, 2, 187, 23, 3, 2,
	2, 2, 188, 186, 3, 2, 2, 2, 189, 190, 5, 2, 2, 2, 190, 191, 5, 22, 12,
	2, 191, 25, 3, 2, 2, 2, 192, 197, 5, 16, 9, 2, 193, 197, 5, 14, 8, 2, 194,
	197, 5, 84, 43, 2, 195, 197, 5, 92, 47, 2, 196, 192, 3, 2, 2, 2, 196, 193,
	3, 2, 2, 2, 196, 194, 3, 2, 2, 2, 196, 195, 3, 2, 2, 2, 197, 27, 3, 2,
	2, 2, 198, 199, 8, 15, 1, 2, 199, 200, 5, 26, 14, 2, 200, 206, 3, 2, 2,
	2, 201, 202, 12, 3, 2, 2, 202, 203, 7, 3, 2, 2, 203, 205, 5, 26, 14, 2,
	204, 201, 3, 2, 2, 2, 205, 208, 3, 2, 2, 2, 206, 204, 3, 2, 2, 2, 206,
	207, 3, 2, 2, 2, 207, 29, 3, 2, 2, 2, 208, 206, 3, 2, 2, 2, 209, 210, 5,
	6, 4, 2, 210, 211, 7, 22, 2, 2, 211, 212, 5, 28, 15, 2, 212, 31, 3, 2,
	2, 2, 213, 214, 7, 23, 2, 2, 214, 215, 5, 18, 10, 2, 215, 217, 7, 24, 2,
	2, 216, 218, 5, 80, 41, 2, 217, 216, 3, 2, 2, 2, 217, 218, 3, 2, 2, 2,
	218, 219, 3, 2, 2, 2, 219, 220, 7, 25, 2, 2, 220, 33, 3, 2, 2, 2, 221,
	222, 7, 26, 2, 2, 222, 223, 5, 18, 10, 2, 223, 225, 7, 24, 2, 2, 224, 226,
	5, 80, 41, 2, 225, 224, 3, 2, 2, 2, 225, 226, 3, 2, 2, 2, 226, 227, 3,
	2, 2, 2, 227, 228, 7, 25, 2, 2, 228, 35, 3, 2, 2, 2, 229, 230, 7, 27, 2,
	2, 230, 232, 7, 24, 2, 2, 231, 233, 5, 80, 41, 2, 232, 231, 3, 2, 2, 2,
	232, 233, 3, 2, 2, 2, 233, 234, 3, 2, 2, 2, 234, 235, 7, 25, 2, 2, 235,
	37, 3, 2, 2, 2, 236, 237, 5, 18, 10, 2, 237, 238, 7, 28, 2, 2, 238, 239,
	5, 78, 40, 2, 239, 39, 3, 2, 2, 2, 240, 241, 5, 78, 40, 2, 241, 41, 3,
	2, 2, 2, 242, 246, 5, 32, 17, 2, 243, 245, 5, 34, 18, 2, 244, 243, 3, 2,
	2, 2, 245, 248, 3, 2, 2, 2, 246, 244, 3, 2, 2, 2, 246, 247, 3, 2, 2, 2,
	247, 250, 3, 2, 2, 2, 248, 246, 3, 2, 2, 2, 249, 251, 5, 36, 19, 2, 250,
	249, 3, 2, 2, 2, 250, 251, 3, 2, 2, 2, 251, 257, 3, 2, 2, 2, 252, 253,
	5, 38, 20, 2, 253, 254, 7, 29, 2, 2, 254, 255, 5, 40, 21, 2, 255, 257,
	3, 2, 2, 2, 256, 242, 3, 2, 2, 2, 256, 252, 3, 2, 2, 2, 257, 43, 3, 2,
	2, 2, 258, 261, 3, 2, 2, 2, 259, 261, 5, 30, 16, 2, 260, 258, 3, 2, 2,
	2, 260, 259, 3, 2, 2, 2, 261, 45, 3, 2, 2, 2, 262, 264, 7, 30, 2, 2, 263,
	265, 5, 44, 23, 2, 264, 263, 3, 2, 2, 2, 264, 265, 3, 2, 2, 2, 265, 266,
	3, 2, 2, 2, 266, 268, 7, 31, 2, 2, 267, 269, 5, 18, 10, 2, 268, 267, 3,
	2, 2, 2, 268, 269, 3, 2, 2, 2, 269, 270, 3, 2, 2, 2, 270, 272, 7, 31, 2,
	2, 271, 273, 5, 44, 23, 2, 272, 271, 3, 2, 2, 2, 272, 273, 3, 2, 2, 2,
	273, 274, 3, 2, 2, 2, 274, 276, 7, 24, 2, 2, 275, 277, 5, 80, 41, 2, 276,
	275, 3, 2, 2, 2, 276, 277, 3, 2, 2, 2, 277, 278, 3, 2, 2, 2, 278, 279,
	7, 25, 2, 2, 279, 47, 3, 2, 2, 2, 280, 283, 5, 30, 16, 2, 281, 283, 5,
	86, 44, 2, 282, 280, 3, 2, 2, 2, 282, 281, 3, 2, 2, 2, 283, 49, 3, 2, 2,
	2, 284, 285, 8, 26, 1, 2, 285, 286, 5, 48, 25, 2, 286, 291, 3, 2, 2, 2,
	287, 288, 12, 3, 2, 2, 288, 290, 5, 48, 25, 2, 289, 287, 3, 2, 2, 2, 290,
	293, 3, 2, 2, 2, 291, 289, 3, 2, 2, 2, 291, 292, 3, 2, 2, 2, 292, 51, 3,
	2, 2, 2, 293, 291, 3, 2, 2, 2, 294, 295, 5, 50, 26, 2, 295, 53, 3, 2, 2,
	2, 296, 300, 5, 8, 5, 2, 297, 300, 5, 12, 7, 2, 298, 300, 5, 66, 34, 2,
	299, 296, 3, 2, 2, 2, 299, 297, 3, 2, 2, 2, 299, 298, 3, 2, 2, 2, 300,
	55, 3, 2, 2, 2, 301, 302, 8, 29, 1, 2, 302, 303, 5, 54, 28, 2, 303, 309,
	3, 2, 2, 2, 304, 305, 12, 3, 2, 2, 305, 306, 7, 3, 2, 2, 306, 308, 5, 54,
	28, 2, 307, 304, 3, 2, 2, 2, 308, 311, 3, 2, 2, 2, 309, 307, 3, 2, 2, 2,
	309, 310, 3, 2, 2, 2, 310, 57, 3, 2, 2, 2, 311, 309, 3, 2, 2, 2, 312, 313,
	5, 2, 2, 2, 313, 59, 3, 2, 2, 2, 314, 315, 8, 31, 1, 2, 315, 316, 5, 58,
	30, 2, 316, 322, 3, 2, 2, 2, 317, 318, 12, 3, 2, 2, 318, 319, 7, 3, 2,
	2, 319, 321, 5, 58, 30, 2, 320, 317, 3, 2, 2, 2, 321, 324, 3, 2, 2, 2,
	322, 320, 3, 2, 2, 2, 322, 323, 3, 2, 2, 2, 323, 61, 3, 2, 2, 2, 324, 322,
	3, 2, 2, 2, 325, 326, 5, 60, 31, 2, 326, 327, 5, 54, 28, 2, 327, 63, 3,
	2, 2, 2, 328, 329, 8, 33, 1, 2, 329, 330, 5, 62, 32, 2, 330, 336, 3, 2,
	2, 2, 331, 332, 12, 3, 2, 2, 332, 333, 7, 3, 2, 2, 333, 335, 5, 62, 32,
	2, 334, 331, 3, 2, 2, 2, 335, 338, 3, 2, 2, 2, 336, 334, 3, 2, 2, 2, 336,
	337, 3, 2, 2, 2, 337, 65, 3, 2, 2, 2, 338, 336, 3, 2, 2, 2, 339, 340, 7,
	36, 2, 2, 340, 342, 7, 7, 2, 2, 341, 343, 5, 64, 33, 2, 342, 341, 3, 2,
	2, 2, 342, 343, 3, 2, 2, 2, 343, 344, 3, 2, 2, 2, 344, 350, 7, 8, 2, 2,
	345, 347, 7, 7, 2, 2, 346, 348, 5, 56, 29, 2, 347, 346, 3, 2, 2, 2, 347,
	348, 3, 2, 2, 2, 348, 349, 3, 2, 2, 2, 349, 351, 7, 8, 2, 2, 350, 345,
	3, 2, 2, 2, 350, 351, 3, 2, 2, 2, 351, 67, 3, 2, 2, 2, 352, 353, 5, 2,
	2, 2, 353, 69, 3, 2, 2, 2, 354, 355, 7, 36, 2, 2, 355, 356, 5, 68, 35,
	2, 356, 358, 7, 7, 2, 2, 357, 359, 5, 64, 33, 2, 358, 357, 3, 2, 2, 2,
	358, 359, 3, 2, 2, 2, 359, 360, 3, 2, 2, 2, 360, 366, 7, 8, 2, 2, 361,
	363, 7, 7, 2, 2, 362, 364, 5, 56, 29, 2, 363, 362, 3, 2, 2, 2, 363, 364,
	3, 2, 2, 2, 364, 365, 3, 2, 2, 2, 365, 367, 7, 8, 2, 2, 366, 361, 3, 2,
	2, 2, 366, 367, 3, 2, 2, 2, 367, 71, 3, 2, 2, 2, 368, 373, 5, 16, 9, 2,
	369, 373, 5, 18, 10, 2, 370, 373, 5, 2, 2, 2, 371, 373, 7, 38, 2, 2, 372,
	368, 3, 2, 2, 2, 372, 369, 3, 2, 2, 2, 372, 370, 3, 2, 2, 2, 372, 371,
	3, 2, 2, 2, 373, 73, 3, 2, 2, 2, 374, 375, 8, 38, 1, 2, 375, 376, 5, 72,
	37, 2, 376, 382, 3, 2, 2, 2, 377, 378, 12, 3, 2, 2, 378, 379, 7, 3, 2,
	2, 379, 381, 5, 72, 37, 2, 380, 377, 3, 2, 2, 2, 381, 384, 3, 2, 2, 2,
	382, 380, 3, 2, 2, 2, 382, 383, 3, 2, 2, 2, 383, 75, 3, 2, 2, 2, 384, 382,
	3, 2, 2, 2, 385, 387, 7, 32, 2, 2, 386, 388, 5, 74, 38, 2, 387, 386, 3,
	2, 2, 2, 387, 388, 3, 2, 2, 2, 388, 77, 3, 2, 2, 2, 389, 395, 5, 30, 16,
	2, 390, 395, 5, 42, 22, 2, 391, 395, 5, 46, 24, 2, 392, 395, 5, 76, 39,
	2, 393, 395, 5, 94, 48, 2, 394, 389, 3, 2, 2, 2, 394, 390, 3, 2, 2, 2,
	394, 391, 3, 2, 2, 2, 394, 392, 3, 2, 2, 2, 394, 393, 3, 2, 2, 2, 395,
	79, 3, 2, 2, 2, 396, 397, 8, 41, 1, 2, 397, 398, 5, 78, 40, 2, 398, 403,
	3, 2, 2, 2, 399, 400, 12, 3, 2, 2, 400, 402, 5, 78, 40, 2, 401, 399, 3,
	2, 2, 2, 402, 405, 3, 2, 2, 2, 403, 401, 3, 2, 2, 2, 403, 404, 3, 2, 2,
	2, 404, 81, 3, 2, 2, 2, 405, 403, 3, 2, 2, 2, 406, 408, 5, 80, 41, 2, 407,
	406, 3, 2, 2, 2, 407, 408, 3, 2, 2, 2, 408, 83, 3, 2, 2, 2, 409, 410, 5,
	66, 34, 2, 410, 412, 7, 24, 2, 2, 411, 413, 5, 82, 42, 2, 412, 411, 3,
	2, 2, 2, 412, 413, 3, 2, 2, 2, 413, 414, 3, 2, 2, 2, 414, 415, 7, 25, 2,
	2, 415, 85, 3, 2, 2, 2, 416, 417, 5, 70, 36, 2, 417, 419, 7, 24, 2, 2,
	418, 420, 5, 82, 42, 2, 419, 418, 3, 2, 2, 2, 419, 420, 3, 2, 2, 2, 420,
	421, 3, 2, 2, 2, 421, 422, 7, 25, 2, 2, 422, 87, 3, 2, 2, 2, 423, 426,
	5, 16, 9, 2, 424, 426, 5, 18, 10, 2, 425, 423, 3, 2, 2, 2, 425, 424, 3,
	2, 2, 2, 426, 89, 3, 2, 2, 2, 427, 428, 8, 46, 1, 2, 428, 429, 5, 88, 45,
	2, 429, 435, 3, 2, 2, 2, 430, 431, 12, 3, 2, 2, 431, 432, 7, 3, 2, 2, 432,
	434, 5, 88, 45, 2, 433, 430, 3, 2, 2, 2, 434, 437, 3, 2, 2, 2, 435, 433,
	3, 2, 2, 2, 435, 436, 3, 2, 2, 2, 436, 91, 3, 2, 2, 2, 437, 435, 3, 2,
	2, 2, 438, 439, 5, 68, 35, 2, 439, 441, 7, 7, 2, 2, 440, 442, 5, 90, 46,
	2, 441, 440, 3, 2, 2, 2, 441, 442, 3, 2, 2, 2, 442, 443, 3, 2, 2, 2, 443,
	444, 7, 8, 2, 2, 444, 93, 3, 2, 2, 2, 445, 446, 5, 92, 47, 2, 446, 95,
	3, 2, 2, 2, 47, 100, 110, 119, 141, 149, 151, 165, 172, 186, 196, 206,
	217, 225, 232, 246, 250, 256, 260, 264, 268, 272, 276, 282, 291, 299, 309,
	322, 336, 342, 347, 350, 358, 363, 366, 372, 382, 387, 394, 403, 407, 412,
	419, 425, 435, 441,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "','", "'string'", "'['", "']'", "'('", "')'", "'*'", "'/'", "'+'",
	"'-'", "'=='", "'<'", "'>'", "'<='", "'>='", "'!='", "'&&'", "'||'", "'!'",
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
	"identifier", "declarator", "declaratorList", "simpleTypeSpecifier", "listElementTypeSpecifier",
	"listTypeSpecifier", "listInitExpression", "aExpr", "bExpr", "listElementIndex",
	"listElementIndexList", "listElementExpression", "assignInit", "assignInitList",
	"assignStatement", "ifExpr", "elsifExpr", "elseExpr", "ternaryIfExpr",
	"ternaryElseExpr", "selectionStatement", "iterationAssignStatement", "iterationStatement",
	"definition", "definitionList", "file", "typeSpecifier", "typeSpecifierList",
	"paraDeclarator", "paraDeclaratorList", "paraDeclaratorWithIdentity", "paraDeclaratorWithIdentityList",
	"funcTypeSpecifier", "funcIdentifier", "funcTypeSpecifierWithName", "funcReturnPara",
	"funcReturnParaList", "funcReturnStatement", "funcStatement", "funcStatementList",
	"funcBody", "funcInitExpression", "funcDefinition", "funcExecutePara",
	"funcExecuteParaList", "funcExecuteExpression", "funcExecuteStatement",
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
	ZZParserRULE_identifier                     = 0
	ZZParserRULE_declarator                     = 1
	ZZParserRULE_declaratorList                 = 2
	ZZParserRULE_simpleTypeSpecifier            = 3
	ZZParserRULE_listElementTypeSpecifier       = 4
	ZZParserRULE_listTypeSpecifier              = 5
	ZZParserRULE_listInitExpression             = 6
	ZZParserRULE_aExpr                          = 7
	ZZParserRULE_bExpr                          = 8
	ZZParserRULE_listElementIndex               = 9
	ZZParserRULE_listElementIndexList           = 10
	ZZParserRULE_listElementExpression          = 11
	ZZParserRULE_assignInit                     = 12
	ZZParserRULE_assignInitList                 = 13
	ZZParserRULE_assignStatement                = 14
	ZZParserRULE_ifExpr                         = 15
	ZZParserRULE_elsifExpr                      = 16
	ZZParserRULE_elseExpr                       = 17
	ZZParserRULE_ternaryIfExpr                  = 18
	ZZParserRULE_ternaryElseExpr                = 19
	ZZParserRULE_selectionStatement             = 20
	ZZParserRULE_iterationAssignStatement       = 21
	ZZParserRULE_iterationStatement             = 22
	ZZParserRULE_definition                     = 23
	ZZParserRULE_definitionList                 = 24
	ZZParserRULE_file                           = 25
	ZZParserRULE_typeSpecifier                  = 26
	ZZParserRULE_typeSpecifierList              = 27
	ZZParserRULE_paraDeclarator                 = 28
	ZZParserRULE_paraDeclaratorList             = 29
	ZZParserRULE_paraDeclaratorWithIdentity     = 30
	ZZParserRULE_paraDeclaratorWithIdentityList = 31
	ZZParserRULE_funcTypeSpecifier              = 32
	ZZParserRULE_funcIdentifier                 = 33
	ZZParserRULE_funcTypeSpecifierWithName      = 34
	ZZParserRULE_funcReturnPara                 = 35
	ZZParserRULE_funcReturnParaList             = 36
	ZZParserRULE_funcReturnStatement            = 37
	ZZParserRULE_funcStatement                  = 38
	ZZParserRULE_funcStatementList              = 39
	ZZParserRULE_funcBody                       = 40
	ZZParserRULE_funcInitExpression             = 41
	ZZParserRULE_funcDefinition                 = 42
	ZZParserRULE_funcExecutePara                = 43
	ZZParserRULE_funcExecuteParaList            = 44
	ZZParserRULE_funcExecuteExpression          = 45
	ZZParserRULE_funcExecuteStatement           = 46
)

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_identifier
	return p
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ZZParserIdentifier, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (p *ZZParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ZZParserRULE_identifier)

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
		p.SetState(94)
		p.Match(ZZParserIdentifier)
	}

	return localctx
}

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

func (s *DeclaratorContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
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
	p.EnterRule(localctx, 2, ZZParserRULE_declarator)

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

	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(96)
			p.Identifier()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(97)
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
	_startState := 4
	p.EnterRecursionRule(localctx, 4, ZZParserRULE_declaratorList, _p)

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
		p.SetState(101)
		p.Declarator()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(108)
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
			p.SetState(103)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(104)
				p.Match(ZZParserT__0)
			}
			{
				p.SetState(105)
				p.Declarator()
			}

		}
		p.SetState(110)
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
	p.EnterRule(localctx, 6, ZZParserRULE_simpleTypeSpecifier)
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
		p.SetState(111)
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
	p.EnterRule(localctx, 8, ZZParserRULE_listElementTypeSpecifier)

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

	p.SetState(117)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ZZParserT__1, ZZParserInt, ZZParserFloat:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(113)
			p.SimpleTypeSpecifier()
		}

	case ZZParserT__2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(114)
			p.Match(ZZParserT__2)
		}
		{
			p.SetState(115)
			p.Match(ZZParserT__3)
		}
		{
			p.SetState(116)
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
	p.EnterRule(localctx, 10, ZZParserRULE_listTypeSpecifier)

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
		p.SetState(119)
		p.Match(ZZParserT__2)
	}
	{
		p.SetState(120)
		p.Match(ZZParserT__3)
	}
	{
		p.SetState(121)
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
	p.EnterRule(localctx, 12, ZZParserRULE_listInitExpression)

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
		p.SetState(123)
		p.Match(ZZParserList)
	}
	{
		p.SetState(124)
		p.Match(ZZParserT__4)
	}
	{
		p.SetState(125)
		p.ListTypeSpecifier()
	}
	{
		p.SetState(126)
		p.Match(ZZParserT__0)
	}
	{
		p.SetState(127)
		p.aExpr(0)
	}
	{
		p.SetState(128)
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

type AExp_identifierContext struct {
	*AExprContext
}

func NewAExp_identifierContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AExp_identifierContext {
	var p = new(AExp_identifierContext)

	p.AExprContext = NewEmptyAExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AExprContext))

	return p
}

func (s *AExp_identifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AExp_identifierContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *AExp_identifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterAExp_identifier(s)
	}
}

func (s *AExp_identifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitAExp_identifier(s)
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
	_startState := 14
	p.EnterRecursionRule(localctx, 14, ZZParserRULE_aExpr, _p)
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
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAExp_IntergerLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(131)
			p.Match(ZZParserIntegerLiteral)
		}

	case 2:
		localctx = NewAExp_FloatLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(132)
			p.Match(ZZParserFloatLiteral)
		}

	case 3:
		localctx = NewAExp_identifierContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(133)
			p.Identifier()
		}

	case 4:
		localctx = NewAExp_listElementExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(134)
			p.ListElementExpression()
		}

	case 5:
		localctx = NewAExp_bracketExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(135)
			p.Match(ZZParserT__4)
		}
		{
			p.SetState(136)
			p.aExpr(0)
		}
		{
			p.SetState(137)
			p.Match(ZZParserT__5)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(147)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAExp_multiplicativeExpressionContext(p, NewAExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ZZParserRULE_aExpr)
				p.SetState(141)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(142)
					_la = p.GetTokenStream().LA(1)

					if !(_la == ZZParserT__6 || _la == ZZParserT__7) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(143)
					p.aExpr(4)
				}

			case 2:
				localctx = NewAExp_additiveExpressionContext(p, NewAExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ZZParserRULE_aExpr)
				p.SetState(144)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(145)
					_la = p.GetTokenStream().LA(1)

					if !(_la == ZZParserT__8 || _la == ZZParserT__9) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(146)
					p.aExpr(3)
				}

			}

		}
		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// IBExprContext is an interface to support dynamic dispatch.
type IBExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBExprContext differentiates from other interfaces.
	IsBExprContext()
}

type BExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBExprContext() *BExprContext {
	var p = new(BExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_bExpr
	return p
}

func (*BExprContext) IsBExprContext() {}

func NewBExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BExprContext {
	var p = new(BExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_bExpr

	return p
}

func (s *BExprContext) GetParser() antlr.Parser { return s.parser }

func (s *BExprContext) CopyFrom(ctx *BExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *BExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BExpr_aExprContext struct {
	*BExprContext
}

func NewBExpr_aExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BExpr_aExprContext {
	var p = new(BExpr_aExprContext)

	p.BExprContext = NewEmptyBExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BExprContext))

	return p
}

func (s *BExpr_aExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BExpr_aExprContext) AllAExpr() []IAExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAExprContext)(nil)).Elem())
	var tst = make([]IAExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAExprContext)
		}
	}

	return tst
}

func (s *BExpr_aExprContext) AExpr(i int) IAExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAExprContext)
}

func (s *BExpr_aExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterBExpr_aExpr(s)
	}
}

func (s *BExpr_aExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitBExpr_aExpr(s)
	}
}

type BExpr_bangContext struct {
	*BExprContext
}

func NewBExpr_bangContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BExpr_bangContext {
	var p = new(BExpr_bangContext)

	p.BExprContext = NewEmptyBExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BExprContext))

	return p
}

func (s *BExpr_bangContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BExpr_bangContext) BExpr() IBExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBExprContext)
}

func (s *BExpr_bangContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterBExpr_bang(s)
	}
}

func (s *BExpr_bangContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitBExpr_bang(s)
	}
}

type BExpr_bExprContext struct {
	*BExprContext
}

func NewBExpr_bExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BExpr_bExprContext {
	var p = new(BExpr_bExprContext)

	p.BExprContext = NewEmptyBExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BExprContext))

	return p
}

func (s *BExpr_bExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BExpr_bExprContext) AllBExpr() []IBExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBExprContext)(nil)).Elem())
	var tst = make([]IBExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBExprContext)
		}
	}

	return tst
}

func (s *BExpr_bExprContext) BExpr(i int) IBExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBExprContext)
}

func (s *BExpr_bExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterBExpr_bExpr(s)
	}
}

func (s *BExpr_bExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitBExpr_bExpr(s)
	}
}

type BExpr_bracketExpressionContext struct {
	*BExprContext
}

func NewBExpr_bracketExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BExpr_bracketExpressionContext {
	var p = new(BExpr_bracketExpressionContext)

	p.BExprContext = NewEmptyBExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BExprContext))

	return p
}

func (s *BExpr_bracketExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BExpr_bracketExpressionContext) BExpr() IBExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBExprContext)
}

func (s *BExpr_bracketExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterBExpr_bracketExpression(s)
	}
}

func (s *BExpr_bracketExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitBExpr_bracketExpression(s)
	}
}

func (p *ZZParser) BExpr() (localctx IBExprContext) {
	return p.bExpr(0)
}

func (p *ZZParser) bExpr(_p int) (localctx IBExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewBExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IBExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 16
	p.EnterRecursionRule(localctx, 16, ZZParserRULE_bExpr, _p)
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
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		localctx = NewBExpr_aExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(153)
			p.aExpr(0)
		}
		{
			p.SetState(154)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ZZParserT__10)|(1<<ZZParserT__11)|(1<<ZZParserT__12)|(1<<ZZParserT__13)|(1<<ZZParserT__14)|(1<<ZZParserT__15))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(155)
			p.aExpr(0)
		}

	case 2:
		localctx = NewBExpr_bangContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(157)
			p.Match(ZZParserT__18)
		}
		{
			p.SetState(158)
			p.bExpr(2)
		}

	case 3:
		localctx = NewBExpr_bracketExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(159)
			p.Match(ZZParserT__4)
		}
		{
			p.SetState(160)
			p.bExpr(0)
		}
		{
			p.SetState(161)
			p.Match(ZZParserT__5)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewBExpr_bExprContext(p, NewBExprContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, ZZParserRULE_bExpr)
			p.SetState(165)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
			}
			{
				p.SetState(166)
				_la = p.GetTokenStream().LA(1)

				if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ZZParserT__10)|(1<<ZZParserT__15)|(1<<ZZParserT__16)|(1<<ZZParserT__17))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(167)
				p.bExpr(4)
			}

		}
		p.SetState(172)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 18, ZZParserRULE_listElementIndex)

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
		p.SetState(173)
		p.Match(ZZParserT__2)
	}
	{
		p.SetState(174)
		p.aExpr(0)
	}
	{
		p.SetState(175)
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
	_startState := 20
	p.EnterRecursionRule(localctx, 20, ZZParserRULE_listElementIndexList, _p)

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
		p.SetState(178)
		p.ListElementIndex()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListElementIndexListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ZZParserRULE_listElementIndexList)
			p.SetState(180)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(181)
				p.ListElementIndex()
			}

		}
		p.SetState(186)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
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

func (s *ListElementExpressionContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
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
		p.SetState(187)
		p.Identifier()
	}
	{
		p.SetState(188)
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

func (s *AssignInitContext) FuncExecuteExpression() IFuncExecuteExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncExecuteExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncExecuteExpressionContext)
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
	p.EnterRule(localctx, 24, ZZParserRULE_assignInit)

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

	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(190)
			p.aExpr(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(191)
			p.ListInitExpression()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(192)
			p.FuncInitExpression()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(193)
			p.FuncExecuteExpression()
		}

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
	_startState := 26
	p.EnterRecursionRule(localctx, 26, ZZParserRULE_assignInitList, _p)

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
		p.SetState(197)
		p.AssignInit()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewAssignInitListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ZZParserRULE_assignInitList)
			p.SetState(199)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(200)
				p.Match(ZZParserT__0)
			}
			{
				p.SetState(201)
				p.AssignInit()
			}

		}
		p.SetState(206)
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
	p.EnterRule(localctx, 28, ZZParserRULE_assignStatement)

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
		p.SetState(207)
		p.declaratorList(0)
	}
	{
		p.SetState(208)
		p.Match(ZZParserT__19)
	}
	{
		p.SetState(209)
		p.assignInitList(0)
	}

	return localctx
}

// IIfExprContext is an interface to support dynamic dispatch.
type IIfExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfExprContext differentiates from other interfaces.
	IsIfExprContext()
}

type IfExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfExprContext() *IfExprContext {
	var p = new(IfExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_ifExpr
	return p
}

func (*IfExprContext) IsIfExprContext() {}

func NewIfExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfExprContext {
	var p = new(IfExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_ifExpr

	return p
}

func (s *IfExprContext) GetParser() antlr.Parser { return s.parser }

func (s *IfExprContext) BExpr() IBExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBExprContext)
}

func (s *IfExprContext) FuncStatementList() IFuncStatementListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncStatementListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncStatementListContext)
}

func (s *IfExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterIfExpr(s)
	}
}

func (s *IfExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitIfExpr(s)
	}
}

func (p *ZZParser) IfExpr() (localctx IIfExprContext) {
	localctx = NewIfExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ZZParserRULE_ifExpr)
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
		p.SetState(211)
		p.Match(ZZParserT__20)
	}
	{
		p.SetState(212)
		p.bExpr(0)
	}
	{
		p.SetState(213)
		p.Match(ZZParserT__21)
	}
	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ZZParserT__4)|(1<<ZZParserT__18)|(1<<ZZParserT__20)|(1<<ZZParserT__27)|(1<<ZZParserT__29))) != 0) || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(ZZParserIdentifier-37))|(1<<(ZZParserIntegerLiteral-37))|(1<<(ZZParserFloatLiteral-37)))) != 0) {
		{
			p.SetState(214)
			p.funcStatementList(0)
		}

	}
	{
		p.SetState(217)
		p.Match(ZZParserT__22)
	}

	return localctx
}

// IElsifExprContext is an interface to support dynamic dispatch.
type IElsifExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElsifExprContext differentiates from other interfaces.
	IsElsifExprContext()
}

type ElsifExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElsifExprContext() *ElsifExprContext {
	var p = new(ElsifExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_elsifExpr
	return p
}

func (*ElsifExprContext) IsElsifExprContext() {}

func NewElsifExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElsifExprContext {
	var p = new(ElsifExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_elsifExpr

	return p
}

func (s *ElsifExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ElsifExprContext) BExpr() IBExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBExprContext)
}

func (s *ElsifExprContext) FuncStatementList() IFuncStatementListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncStatementListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncStatementListContext)
}

func (s *ElsifExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElsifExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElsifExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterElsifExpr(s)
	}
}

func (s *ElsifExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitElsifExpr(s)
	}
}

func (p *ZZParser) ElsifExpr() (localctx IElsifExprContext) {
	localctx = NewElsifExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ZZParserRULE_elsifExpr)
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
		p.SetState(219)
		p.Match(ZZParserT__23)
	}
	{
		p.SetState(220)
		p.bExpr(0)
	}
	{
		p.SetState(221)
		p.Match(ZZParserT__21)
	}
	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ZZParserT__4)|(1<<ZZParserT__18)|(1<<ZZParserT__20)|(1<<ZZParserT__27)|(1<<ZZParserT__29))) != 0) || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(ZZParserIdentifier-37))|(1<<(ZZParserIntegerLiteral-37))|(1<<(ZZParserFloatLiteral-37)))) != 0) {
		{
			p.SetState(222)
			p.funcStatementList(0)
		}

	}
	{
		p.SetState(225)
		p.Match(ZZParserT__22)
	}

	return localctx
}

// IElseExprContext is an interface to support dynamic dispatch.
type IElseExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElseExprContext differentiates from other interfaces.
	IsElseExprContext()
}

type ElseExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseExprContext() *ElseExprContext {
	var p = new(ElseExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_elseExpr
	return p
}

func (*ElseExprContext) IsElseExprContext() {}

func NewElseExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseExprContext {
	var p = new(ElseExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_elseExpr

	return p
}

func (s *ElseExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseExprContext) FuncStatementList() IFuncStatementListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncStatementListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncStatementListContext)
}

func (s *ElseExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterElseExpr(s)
	}
}

func (s *ElseExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitElseExpr(s)
	}
}

func (p *ZZParser) ElseExpr() (localctx IElseExprContext) {
	localctx = NewElseExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ZZParserRULE_elseExpr)
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
		p.SetState(227)
		p.Match(ZZParserT__24)
	}
	{
		p.SetState(228)
		p.Match(ZZParserT__21)
	}
	p.SetState(230)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ZZParserT__4)|(1<<ZZParserT__18)|(1<<ZZParserT__20)|(1<<ZZParserT__27)|(1<<ZZParserT__29))) != 0) || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(ZZParserIdentifier-37))|(1<<(ZZParserIntegerLiteral-37))|(1<<(ZZParserFloatLiteral-37)))) != 0) {
		{
			p.SetState(229)
			p.funcStatementList(0)
		}

	}
	{
		p.SetState(232)
		p.Match(ZZParserT__22)
	}

	return localctx
}

// ITernaryIfExprContext is an interface to support dynamic dispatch.
type ITernaryIfExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTernaryIfExprContext differentiates from other interfaces.
	IsTernaryIfExprContext()
}

type TernaryIfExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTernaryIfExprContext() *TernaryIfExprContext {
	var p = new(TernaryIfExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_ternaryIfExpr
	return p
}

func (*TernaryIfExprContext) IsTernaryIfExprContext() {}

func NewTernaryIfExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TernaryIfExprContext {
	var p = new(TernaryIfExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_ternaryIfExpr

	return p
}

func (s *TernaryIfExprContext) GetParser() antlr.Parser { return s.parser }

func (s *TernaryIfExprContext) BExpr() IBExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBExprContext)
}

func (s *TernaryIfExprContext) FuncStatement() IFuncStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncStatementContext)
}

func (s *TernaryIfExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TernaryIfExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TernaryIfExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterTernaryIfExpr(s)
	}
}

func (s *TernaryIfExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitTernaryIfExpr(s)
	}
}

func (p *ZZParser) TernaryIfExpr() (localctx ITernaryIfExprContext) {
	localctx = NewTernaryIfExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ZZParserRULE_ternaryIfExpr)

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
		p.SetState(234)
		p.bExpr(0)
	}
	{
		p.SetState(235)
		p.Match(ZZParserT__25)
	}
	{
		p.SetState(236)
		p.FuncStatement()
	}

	return localctx
}

// ITernaryElseExprContext is an interface to support dynamic dispatch.
type ITernaryElseExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTernaryElseExprContext differentiates from other interfaces.
	IsTernaryElseExprContext()
}

type TernaryElseExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTernaryElseExprContext() *TernaryElseExprContext {
	var p = new(TernaryElseExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_ternaryElseExpr
	return p
}

func (*TernaryElseExprContext) IsTernaryElseExprContext() {}

func NewTernaryElseExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TernaryElseExprContext {
	var p = new(TernaryElseExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_ternaryElseExpr

	return p
}

func (s *TernaryElseExprContext) GetParser() antlr.Parser { return s.parser }

func (s *TernaryElseExprContext) FuncStatement() IFuncStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncStatementContext)
}

func (s *TernaryElseExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TernaryElseExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TernaryElseExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterTernaryElseExpr(s)
	}
}

func (s *TernaryElseExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitTernaryElseExpr(s)
	}
}

func (p *ZZParser) TernaryElseExpr() (localctx ITernaryElseExprContext) {
	localctx = NewTernaryElseExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ZZParserRULE_ternaryElseExpr)

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
		p.SetState(238)
		p.FuncStatement()
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

func (s *SelectionStatementContext) IfExpr() IIfExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfExprContext)
}

func (s *SelectionStatementContext) AllElsifExpr() []IElsifExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IElsifExprContext)(nil)).Elem())
	var tst = make([]IElsifExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IElsifExprContext)
		}
	}

	return tst
}

func (s *SelectionStatementContext) ElsifExpr(i int) IElsifExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElsifExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IElsifExprContext)
}

func (s *SelectionStatementContext) ElseExpr() IElseExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElseExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElseExprContext)
}

func (s *SelectionStatementContext) TernaryIfExpr() ITernaryIfExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITernaryIfExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITernaryIfExprContext)
}

func (s *SelectionStatementContext) TernaryElseExpr() ITernaryElseExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITernaryElseExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITernaryElseExprContext)
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
	p.EnterRule(localctx, 40, ZZParserRULE_selectionStatement)

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

	p.SetState(254)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ZZParserT__20:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(240)
			p.IfExpr()
		}
		p.SetState(244)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(241)
					p.ElsifExpr()
				}

			}
			p.SetState(246)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())
		}
		p.SetState(248)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(247)
				p.ElseExpr()
			}

		}

	case ZZParserT__4, ZZParserT__18, ZZParserIdentifier, ZZParserIntegerLiteral, ZZParserFloatLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(250)
			p.TernaryIfExpr()
		}
		{
			p.SetState(251)
			p.Match(ZZParserT__26)
		}
		{
			p.SetState(252)
			p.TernaryElseExpr()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIterationAssignStatementContext is an interface to support dynamic dispatch.
type IIterationAssignStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIterationAssignStatementContext differentiates from other interfaces.
	IsIterationAssignStatementContext()
}

type IterationAssignStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIterationAssignStatementContext() *IterationAssignStatementContext {
	var p = new(IterationAssignStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_iterationAssignStatement
	return p
}

func (*IterationAssignStatementContext) IsIterationAssignStatementContext() {}

func NewIterationAssignStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IterationAssignStatementContext {
	var p = new(IterationAssignStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_iterationAssignStatement

	return p
}

func (s *IterationAssignStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IterationAssignStatementContext) AssignStatement() IAssignStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignStatementContext)
}

func (s *IterationAssignStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IterationAssignStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IterationAssignStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterIterationAssignStatement(s)
	}
}

func (s *IterationAssignStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitIterationAssignStatement(s)
	}
}

func (p *ZZParser) IterationAssignStatement() (localctx IIterationAssignStatementContext) {
	localctx = NewIterationAssignStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ZZParserRULE_iterationAssignStatement)

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

	p.SetState(258)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ZZParserT__21, ZZParserT__28:
		p.EnterOuterAlt(localctx, 1)

	case ZZParserIdentifier:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(257)
			p.AssignStatement()
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

func (s *IterationStatementContext) AllIterationAssignStatement() []IIterationAssignStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIterationAssignStatementContext)(nil)).Elem())
	var tst = make([]IIterationAssignStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIterationAssignStatementContext)
		}
	}

	return tst
}

func (s *IterationStatementContext) IterationAssignStatement(i int) IIterationAssignStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIterationAssignStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIterationAssignStatementContext)
}

func (s *IterationStatementContext) BExpr() IBExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBExprContext)
}

func (s *IterationStatementContext) FuncStatementList() IFuncStatementListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncStatementListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncStatementListContext)
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
	p.EnterRule(localctx, 44, ZZParserRULE_iterationStatement)
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
		p.SetState(260)
		p.Match(ZZParserT__27)
	}
	p.SetState(262)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(261)
			p.IterationAssignStatement()
		}

	}
	{
		p.SetState(264)
		p.Match(ZZParserT__28)
	}
	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ZZParserT__4 || _la == ZZParserT__18 || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(ZZParserIdentifier-37))|(1<<(ZZParserIntegerLiteral-37))|(1<<(ZZParserFloatLiteral-37)))) != 0) {
		{
			p.SetState(265)
			p.bExpr(0)
		}

	}
	{
		p.SetState(268)
		p.Match(ZZParserT__28)
	}
	p.SetState(270)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(269)
			p.IterationAssignStatement()
		}

	}
	{
		p.SetState(272)
		p.Match(ZZParserT__21)
	}
	p.SetState(274)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ZZParserT__4)|(1<<ZZParserT__18)|(1<<ZZParserT__20)|(1<<ZZParserT__27)|(1<<ZZParserT__29))) != 0) || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(ZZParserIdentifier-37))|(1<<(ZZParserIntegerLiteral-37))|(1<<(ZZParserFloatLiteral-37)))) != 0) {
		{
			p.SetState(273)
			p.funcStatementList(0)
		}

	}
	{
		p.SetState(276)
		p.Match(ZZParserT__22)
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
	p.EnterRule(localctx, 46, ZZParserRULE_definition)

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
	case ZZParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(278)
			p.AssignStatement()
		}

	case ZZParserFunc:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(279)
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
	_startState := 48
	p.EnterRecursionRule(localctx, 48, ZZParserRULE_definitionList, _p)

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
		p.Definition()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(289)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewDefinitionListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ZZParserRULE_definitionList)
			p.SetState(285)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(286)
				p.Definition()
			}

		}
		p.SetState(291)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())
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

func (s *FileContext) DefinitionList() IDefinitionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefinitionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefinitionListContext)
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
	p.EnterRule(localctx, 50, ZZParserRULE_file)

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
		p.SetState(292)
		p.definitionList(0)
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
	p.EnterRule(localctx, 52, ZZParserRULE_typeSpecifier)

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

	p.SetState(297)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ZZParserT__1, ZZParserInt, ZZParserFloat:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(294)
			p.SimpleTypeSpecifier()
		}

	case ZZParserT__2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(295)
			p.ListTypeSpecifier()
		}

	case ZZParserFunc:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(296)
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
	_startState := 54
	p.EnterRecursionRule(localctx, 54, ZZParserRULE_typeSpecifierList, _p)

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
		p.SetState(300)
		p.TypeSpecifier()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(307)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewTypeSpecifierListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ZZParserRULE_typeSpecifierList)
			p.SetState(302)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(303)
				p.Match(ZZParserT__0)
			}
			{
				p.SetState(304)
				p.TypeSpecifier()
			}

		}
		p.SetState(309)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())
	}

	return localctx
}

// IParaDeclaratorContext is an interface to support dynamic dispatch.
type IParaDeclaratorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParaDeclaratorContext differentiates from other interfaces.
	IsParaDeclaratorContext()
}

type ParaDeclaratorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParaDeclaratorContext() *ParaDeclaratorContext {
	var p = new(ParaDeclaratorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_paraDeclarator
	return p
}

func (*ParaDeclaratorContext) IsParaDeclaratorContext() {}

func NewParaDeclaratorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParaDeclaratorContext {
	var p = new(ParaDeclaratorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_paraDeclarator

	return p
}

func (s *ParaDeclaratorContext) GetParser() antlr.Parser { return s.parser }

func (s *ParaDeclaratorContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ParaDeclaratorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParaDeclaratorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParaDeclaratorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterParaDeclarator(s)
	}
}

func (s *ParaDeclaratorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitParaDeclarator(s)
	}
}

func (p *ZZParser) ParaDeclarator() (localctx IParaDeclaratorContext) {
	localctx = NewParaDeclaratorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, ZZParserRULE_paraDeclarator)

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
		p.Identifier()
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

func (s *ParaDeclaratorListContext) ParaDeclarator() IParaDeclaratorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParaDeclaratorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParaDeclaratorContext)
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
	_startState := 58
	p.EnterRecursionRule(localctx, 58, ZZParserRULE_paraDeclaratorList, _p)

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
		p.SetState(313)
		p.ParaDeclarator()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(320)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewParaDeclaratorListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ZZParserRULE_paraDeclaratorList)
			p.SetState(315)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(316)
				p.Match(ZZParserT__0)
			}
			{
				p.SetState(317)
				p.ParaDeclarator()
			}

		}
		p.SetState(322)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())
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

func (s *ParaDeclaratorWithIdentityContext) ParaDeclaratorList() IParaDeclaratorListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParaDeclaratorListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParaDeclaratorListContext)
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
	p.EnterRule(localctx, 60, ZZParserRULE_paraDeclaratorWithIdentity)

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
		p.SetState(323)
		p.paraDeclaratorList(0)
	}
	{
		p.SetState(324)
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
	_startState := 62
	p.EnterRecursionRule(localctx, 62, ZZParserRULE_paraDeclaratorWithIdentityList, _p)

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
		p.SetState(327)
		p.ParaDeclaratorWithIdentity()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(334)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewParaDeclaratorWithIdentityListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ZZParserRULE_paraDeclaratorWithIdentityList)
			p.SetState(329)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(330)
				p.Match(ZZParserT__0)
			}
			{
				p.SetState(331)
				p.ParaDeclaratorWithIdentity()
			}

		}
		p.SetState(336)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 64, ZZParserRULE_funcTypeSpecifier)
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
		p.SetState(337)
		p.Match(ZZParserFunc)
	}
	{
		p.SetState(338)
		p.Match(ZZParserT__4)
	}
	p.SetState(340)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ZZParserIdentifier {
		{
			p.SetState(339)
			p.paraDeclaratorWithIdentityList(0)
		}

	}
	{
		p.SetState(342)
		p.Match(ZZParserT__5)
	}
	p.SetState(348)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(343)
			p.Match(ZZParserT__4)
		}
		p.SetState(345)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ZZParserT__1 || _la == ZZParserT__2 || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(ZZParserInt-32))|(1<<(ZZParserFloat-32))|(1<<(ZZParserFunc-32)))) != 0) {
			{
				p.SetState(344)
				p.typeSpecifierList(0)
			}

		}
		{
			p.SetState(347)
			p.Match(ZZParserT__5)
		}

	}

	return localctx
}

// IFuncIdentifierContext is an interface to support dynamic dispatch.
type IFuncIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncIdentifierContext differentiates from other interfaces.
	IsFuncIdentifierContext()
}

type FuncIdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncIdentifierContext() *FuncIdentifierContext {
	var p = new(FuncIdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_funcIdentifier
	return p
}

func (*FuncIdentifierContext) IsFuncIdentifierContext() {}

func NewFuncIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncIdentifierContext {
	var p = new(FuncIdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_funcIdentifier

	return p
}

func (s *FuncIdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncIdentifierContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *FuncIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterFuncIdentifier(s)
	}
}

func (s *FuncIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitFuncIdentifier(s)
	}
}

func (p *ZZParser) FuncIdentifier() (localctx IFuncIdentifierContext) {
	localctx = NewFuncIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, ZZParserRULE_funcIdentifier)

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
		p.SetState(350)
		p.Identifier()
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

func (s *FuncTypeSpecifierWithNameContext) FuncIdentifier() IFuncIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncIdentifierContext)
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
	p.EnterRule(localctx, 68, ZZParserRULE_funcTypeSpecifierWithName)
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
		p.SetState(352)
		p.Match(ZZParserFunc)
	}
	{
		p.SetState(353)
		p.FuncIdentifier()
	}
	{
		p.SetState(354)
		p.Match(ZZParserT__4)
	}
	p.SetState(356)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ZZParserIdentifier {
		{
			p.SetState(355)
			p.paraDeclaratorWithIdentityList(0)
		}

	}
	{
		p.SetState(358)
		p.Match(ZZParserT__5)
	}
	p.SetState(364)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ZZParserT__4 {
		{
			p.SetState(359)
			p.Match(ZZParserT__4)
		}
		p.SetState(361)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ZZParserT__1 || _la == ZZParserT__2 || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(ZZParserInt-32))|(1<<(ZZParserFloat-32))|(1<<(ZZParserFunc-32)))) != 0) {
			{
				p.SetState(360)
				p.typeSpecifierList(0)
			}

		}
		{
			p.SetState(363)
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

func (s *FuncReturnParaContext) BExpr() IBExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBExprContext)
}

func (s *FuncReturnParaContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
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
	p.EnterRule(localctx, 70, ZZParserRULE_funcReturnPara)

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

	p.SetState(370)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(366)
			p.aExpr(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(367)
			p.bExpr(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(368)
			p.Identifier()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(369)
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
	_startState := 72
	p.EnterRecursionRule(localctx, 72, ZZParserRULE_funcReturnParaList, _p)

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
		p.SetState(373)
		p.FuncReturnPara()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(380)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewFuncReturnParaListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ZZParserRULE_funcReturnParaList)
			p.SetState(375)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(376)
				p.Match(ZZParserT__0)
			}
			{
				p.SetState(377)
				p.FuncReturnPara()
			}

		}
		p.SetState(382)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext())
	}

	return localctx
}

// IFuncReturnStatementContext is an interface to support dynamic dispatch.
type IFuncReturnStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncReturnStatementContext differentiates from other interfaces.
	IsFuncReturnStatementContext()
}

type FuncReturnStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncReturnStatementContext() *FuncReturnStatementContext {
	var p = new(FuncReturnStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_funcReturnStatement
	return p
}

func (*FuncReturnStatementContext) IsFuncReturnStatementContext() {}

func NewFuncReturnStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncReturnStatementContext {
	var p = new(FuncReturnStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_funcReturnStatement

	return p
}

func (s *FuncReturnStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncReturnStatementContext) FuncReturnParaList() IFuncReturnParaListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncReturnParaListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncReturnParaListContext)
}

func (s *FuncReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncReturnStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncReturnStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterFuncReturnStatement(s)
	}
}

func (s *FuncReturnStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitFuncReturnStatement(s)
	}
}

func (p *ZZParser) FuncReturnStatement() (localctx IFuncReturnStatementContext) {
	localctx = NewFuncReturnStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, ZZParserRULE_funcReturnStatement)

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
		p.SetState(383)
		p.Match(ZZParserT__29)
	}
	p.SetState(385)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(384)
			p.funcReturnParaList(0)
		}

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

func (s *FuncStatementContext) FuncReturnStatement() IFuncReturnStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncReturnStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncReturnStatementContext)
}

func (s *FuncStatementContext) FuncExecuteStatement() IFuncExecuteStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncExecuteStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncExecuteStatementContext)
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
	p.EnterRule(localctx, 76, ZZParserRULE_funcStatement)

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

	p.SetState(392)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(387)
			p.AssignStatement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(388)
			p.SelectionStatement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(389)
			p.IterationStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(390)
			p.FuncReturnStatement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(391)
			p.FuncExecuteStatement()
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
	_startState := 78
	p.EnterRecursionRule(localctx, 78, ZZParserRULE_funcStatementList, _p)

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
		p.SetState(395)
		p.FuncStatement()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(401)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewFuncStatementListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ZZParserRULE_funcStatementList)
			p.SetState(397)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(398)
				p.FuncStatement()
			}

		}
		p.SetState(403)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 80, ZZParserRULE_funcBody)
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
	p.SetState(405)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ZZParserT__4)|(1<<ZZParserT__18)|(1<<ZZParserT__20)|(1<<ZZParserT__27)|(1<<ZZParserT__29))) != 0) || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(ZZParserIdentifier-37))|(1<<(ZZParserIntegerLiteral-37))|(1<<(ZZParserFloatLiteral-37)))) != 0) {
		{
			p.SetState(404)
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
	p.EnterRule(localctx, 82, ZZParserRULE_funcInitExpression)

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
		p.SetState(407)
		p.FuncTypeSpecifier()
	}
	{
		p.SetState(408)
		p.Match(ZZParserT__21)
	}
	p.SetState(410)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(409)
			p.FuncBody()
		}

	}
	{
		p.SetState(412)
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
	p.EnterRule(localctx, 84, ZZParserRULE_funcDefinition)

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
		p.SetState(414)
		p.FuncTypeSpecifierWithName()
	}
	{
		p.SetState(415)
		p.Match(ZZParserT__21)
	}
	p.SetState(417)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(416)
			p.FuncBody()
		}

	}
	{
		p.SetState(419)
		p.Match(ZZParserT__22)
	}

	return localctx
}

// IFuncExecuteParaContext is an interface to support dynamic dispatch.
type IFuncExecuteParaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncExecuteParaContext differentiates from other interfaces.
	IsFuncExecuteParaContext()
}

type FuncExecuteParaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncExecuteParaContext() *FuncExecuteParaContext {
	var p = new(FuncExecuteParaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_funcExecutePara
	return p
}

func (*FuncExecuteParaContext) IsFuncExecuteParaContext() {}

func NewFuncExecuteParaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncExecuteParaContext {
	var p = new(FuncExecuteParaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_funcExecutePara

	return p
}

func (s *FuncExecuteParaContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncExecuteParaContext) AExpr() IAExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAExprContext)
}

func (s *FuncExecuteParaContext) BExpr() IBExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBExprContext)
}

func (s *FuncExecuteParaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncExecuteParaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncExecuteParaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterFuncExecutePara(s)
	}
}

func (s *FuncExecuteParaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitFuncExecutePara(s)
	}
}

func (p *ZZParser) FuncExecutePara() (localctx IFuncExecuteParaContext) {
	localctx = NewFuncExecuteParaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, ZZParserRULE_funcExecutePara)

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

	p.SetState(423)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(421)
			p.aExpr(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(422)
			p.bExpr(0)
		}

	}

	return localctx
}

// IFuncExecuteParaListContext is an interface to support dynamic dispatch.
type IFuncExecuteParaListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncExecuteParaListContext differentiates from other interfaces.
	IsFuncExecuteParaListContext()
}

type FuncExecuteParaListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncExecuteParaListContext() *FuncExecuteParaListContext {
	var p = new(FuncExecuteParaListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_funcExecuteParaList
	return p
}

func (*FuncExecuteParaListContext) IsFuncExecuteParaListContext() {}

func NewFuncExecuteParaListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncExecuteParaListContext {
	var p = new(FuncExecuteParaListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_funcExecuteParaList

	return p
}

func (s *FuncExecuteParaListContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncExecuteParaListContext) FuncExecutePara() IFuncExecuteParaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncExecuteParaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncExecuteParaContext)
}

func (s *FuncExecuteParaListContext) FuncExecuteParaList() IFuncExecuteParaListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncExecuteParaListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncExecuteParaListContext)
}

func (s *FuncExecuteParaListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncExecuteParaListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncExecuteParaListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterFuncExecuteParaList(s)
	}
}

func (s *FuncExecuteParaListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitFuncExecuteParaList(s)
	}
}

func (p *ZZParser) FuncExecuteParaList() (localctx IFuncExecuteParaListContext) {
	return p.funcExecuteParaList(0)
}

func (p *ZZParser) funcExecuteParaList(_p int) (localctx IFuncExecuteParaListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewFuncExecuteParaListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IFuncExecuteParaListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 88
	p.EnterRecursionRule(localctx, 88, ZZParserRULE_funcExecuteParaList, _p)

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
		p.SetState(426)
		p.FuncExecutePara()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(433)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewFuncExecuteParaListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ZZParserRULE_funcExecuteParaList)
			p.SetState(428)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(429)
				p.Match(ZZParserT__0)
			}
			{
				p.SetState(430)
				p.FuncExecutePara()
			}

		}
		p.SetState(435)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext())
	}

	return localctx
}

// IFuncExecuteExpressionContext is an interface to support dynamic dispatch.
type IFuncExecuteExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncExecuteExpressionContext differentiates from other interfaces.
	IsFuncExecuteExpressionContext()
}

type FuncExecuteExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncExecuteExpressionContext() *FuncExecuteExpressionContext {
	var p = new(FuncExecuteExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_funcExecuteExpression
	return p
}

func (*FuncExecuteExpressionContext) IsFuncExecuteExpressionContext() {}

func NewFuncExecuteExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncExecuteExpressionContext {
	var p = new(FuncExecuteExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_funcExecuteExpression

	return p
}

func (s *FuncExecuteExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncExecuteExpressionContext) FuncIdentifier() IFuncIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncIdentifierContext)
}

func (s *FuncExecuteExpressionContext) FuncExecuteParaList() IFuncExecuteParaListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncExecuteParaListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncExecuteParaListContext)
}

func (s *FuncExecuteExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncExecuteExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncExecuteExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterFuncExecuteExpression(s)
	}
}

func (s *FuncExecuteExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitFuncExecuteExpression(s)
	}
}

func (p *ZZParser) FuncExecuteExpression() (localctx IFuncExecuteExpressionContext) {
	localctx = NewFuncExecuteExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, ZZParserRULE_funcExecuteExpression)
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
		p.SetState(436)
		p.FuncIdentifier()
	}
	{
		p.SetState(437)
		p.Match(ZZParserT__4)
	}
	p.SetState(439)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ZZParserT__4 || _la == ZZParserT__18 || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(ZZParserIdentifier-37))|(1<<(ZZParserIntegerLiteral-37))|(1<<(ZZParserFloatLiteral-37)))) != 0) {
		{
			p.SetState(438)
			p.funcExecuteParaList(0)
		}

	}
	{
		p.SetState(441)
		p.Match(ZZParserT__5)
	}

	return localctx
}

// IFuncExecuteStatementContext is an interface to support dynamic dispatch.
type IFuncExecuteStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncExecuteStatementContext differentiates from other interfaces.
	IsFuncExecuteStatementContext()
}

type FuncExecuteStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncExecuteStatementContext() *FuncExecuteStatementContext {
	var p = new(FuncExecuteStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_funcExecuteStatement
	return p
}

func (*FuncExecuteStatementContext) IsFuncExecuteStatementContext() {}

func NewFuncExecuteStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncExecuteStatementContext {
	var p = new(FuncExecuteStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_funcExecuteStatement

	return p
}

func (s *FuncExecuteStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncExecuteStatementContext) FuncExecuteExpression() IFuncExecuteExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncExecuteExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncExecuteExpressionContext)
}

func (s *FuncExecuteStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncExecuteStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncExecuteStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterFuncExecuteStatement(s)
	}
}

func (s *FuncExecuteStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitFuncExecuteStatement(s)
	}
}

func (p *ZZParser) FuncExecuteStatement() (localctx IFuncExecuteStatementContext) {
	localctx = NewFuncExecuteStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, ZZParserRULE_funcExecuteStatement)

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
		p.SetState(443)
		p.FuncExecuteExpression()
	}

	return localctx
}

func (p *ZZParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *DeclaratorListContext = nil
		if localctx != nil {
			t = localctx.(*DeclaratorListContext)
		}
		return p.DeclaratorList_Sempred(t, predIndex)

	case 7:
		var t *AExprContext = nil
		if localctx != nil {
			t = localctx.(*AExprContext)
		}
		return p.AExpr_Sempred(t, predIndex)

	case 8:
		var t *BExprContext = nil
		if localctx != nil {
			t = localctx.(*BExprContext)
		}
		return p.BExpr_Sempred(t, predIndex)

	case 10:
		var t *ListElementIndexListContext = nil
		if localctx != nil {
			t = localctx.(*ListElementIndexListContext)
		}
		return p.ListElementIndexList_Sempred(t, predIndex)

	case 13:
		var t *AssignInitListContext = nil
		if localctx != nil {
			t = localctx.(*AssignInitListContext)
		}
		return p.AssignInitList_Sempred(t, predIndex)

	case 24:
		var t *DefinitionListContext = nil
		if localctx != nil {
			t = localctx.(*DefinitionListContext)
		}
		return p.DefinitionList_Sempred(t, predIndex)

	case 27:
		var t *TypeSpecifierListContext = nil
		if localctx != nil {
			t = localctx.(*TypeSpecifierListContext)
		}
		return p.TypeSpecifierList_Sempred(t, predIndex)

	case 29:
		var t *ParaDeclaratorListContext = nil
		if localctx != nil {
			t = localctx.(*ParaDeclaratorListContext)
		}
		return p.ParaDeclaratorList_Sempred(t, predIndex)

	case 31:
		var t *ParaDeclaratorWithIdentityListContext = nil
		if localctx != nil {
			t = localctx.(*ParaDeclaratorWithIdentityListContext)
		}
		return p.ParaDeclaratorWithIdentityList_Sempred(t, predIndex)

	case 36:
		var t *FuncReturnParaListContext = nil
		if localctx != nil {
			t = localctx.(*FuncReturnParaListContext)
		}
		return p.FuncReturnParaList_Sempred(t, predIndex)

	case 39:
		var t *FuncStatementListContext = nil
		if localctx != nil {
			t = localctx.(*FuncStatementListContext)
		}
		return p.FuncStatementList_Sempred(t, predIndex)

	case 44:
		var t *FuncExecuteParaListContext = nil
		if localctx != nil {
			t = localctx.(*FuncExecuteParaListContext)
		}
		return p.FuncExecuteParaList_Sempred(t, predIndex)

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

func (p *ZZParser) BExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ZZParser) ListElementIndexList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ZZParser) AssignInitList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 5:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ZZParser) DefinitionList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 6:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ZZParser) TypeSpecifierList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 7:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ZZParser) ParaDeclaratorList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 8:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ZZParser) ParaDeclaratorWithIdentityList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
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

func (p *ZZParser) FuncExecuteParaList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 12:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
