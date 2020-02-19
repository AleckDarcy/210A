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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 46, 486,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4,
	50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 3, 2, 3, 2, 3, 3, 3, 3, 5, 3, 109,
	10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 117, 10, 4, 12, 4, 14,
	4, 120, 11, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 128, 10, 6, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 160, 10, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 7, 10, 168, 10, 10, 12, 10, 14, 10,
	171, 11, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 7, 11, 179, 10,
	11, 12, 11, 14, 11, 182, 11, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 195, 10, 12, 3, 12, 3, 12,
	3, 12, 7, 12, 200, 10, 12, 12, 12, 14, 12, 203, 11, 12, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 7, 14, 214, 10, 14, 12, 14,
	14, 14, 217, 11, 14, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 5,
	16, 226, 10, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 7, 17, 234,
	10, 17, 12, 17, 14, 17, 237, 11, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19,
	3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20, 251, 10, 20, 3,
	20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 259, 10, 21, 3, 21, 3, 21,
	3, 22, 3, 22, 3, 22, 5, 22, 266, 10, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3,
	23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 7, 25, 278, 10, 25, 12, 25, 14,
	25, 281, 11, 25, 3, 25, 5, 25, 284, 10, 25, 3, 25, 3, 25, 3, 25, 3, 25,
	5, 25, 290, 10, 25, 3, 26, 3, 26, 5, 26, 294, 10, 26, 3, 27, 3, 27, 3,
	28, 3, 28, 5, 28, 300, 10, 28, 3, 28, 3, 28, 5, 28, 304, 10, 28, 3, 28,
	3, 28, 5, 28, 308, 10, 28, 3, 28, 3, 28, 5, 28, 312, 10, 28, 3, 28, 3,
	28, 3, 29, 3, 29, 5, 29, 318, 10, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30,
	7, 30, 325, 10, 30, 12, 30, 14, 30, 328, 11, 30, 3, 31, 3, 31, 3, 32, 3,
	32, 5, 32, 334, 10, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 7, 33,
	342, 10, 33, 12, 33, 14, 33, 345, 11, 33, 3, 34, 3, 34, 3, 35, 3, 35, 3,
	35, 3, 35, 3, 35, 3, 35, 7, 35, 355, 10, 35, 12, 35, 14, 35, 358, 11, 35,
	3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 7, 37, 369,
	10, 37, 12, 37, 14, 37, 372, 11, 37, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39,
	3, 39, 5, 39, 380, 10, 39, 3, 39, 3, 39, 3, 39, 5, 39, 385, 10, 39, 3,
	39, 5, 39, 388, 10, 39, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 5, 40, 395,
	10, 40, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 7, 41, 403, 10, 41, 12,
	41, 14, 41, 406, 11, 41, 3, 42, 3, 42, 5, 42, 410, 10, 42, 3, 43, 3, 43,
	3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 5, 43, 419, 10, 43, 3, 44, 3, 44, 3,
	44, 3, 44, 3, 44, 7, 44, 426, 10, 44, 12, 44, 14, 44, 429, 11, 44, 3, 45,
	5, 45, 432, 10, 45, 3, 46, 3, 46, 3, 46, 5, 46, 437, 10, 46, 3, 46, 3,
	46, 3, 47, 3, 47, 5, 47, 443, 10, 47, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48,
	3, 48, 7, 48, 451, 10, 48, 12, 48, 14, 48, 454, 11, 48, 3, 49, 3, 49, 3,
	49, 5, 49, 459, 10, 49, 3, 49, 3, 49, 3, 50, 3, 50, 3, 51, 3, 51, 3, 51,
	5, 51, 468, 10, 51, 3, 51, 3, 51, 3, 51, 3, 51, 5, 51, 474, 10, 51, 7,
	51, 476, 10, 51, 12, 51, 14, 51, 479, 11, 51, 3, 52, 3, 52, 3, 52, 3, 52,
	3, 52, 3, 52, 2, 16, 6, 18, 20, 22, 26, 32, 58, 64, 68, 72, 80, 86, 94,
	100, 53, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
	36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70,
	72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 2, 7,
	4, 2, 4, 4, 37, 39, 3, 2, 9, 10, 3, 2, 11, 12, 3, 2, 14, 19, 4, 2, 14,
	14, 19, 21, 2, 494, 2, 104, 3, 2, 2, 2, 4, 108, 3, 2, 2, 2, 6, 110, 3,
	2, 2, 2, 8, 121, 3, 2, 2, 2, 10, 127, 3, 2, 2, 2, 12, 129, 3, 2, 2, 2,
	14, 133, 3, 2, 2, 2, 16, 140, 3, 2, 2, 2, 18, 159, 3, 2, 2, 2, 20, 172,
	3, 2, 2, 2, 22, 194, 3, 2, 2, 2, 24, 204, 3, 2, 2, 2, 26, 208, 3, 2, 2,
	2, 28, 218, 3, 2, 2, 2, 30, 225, 3, 2, 2, 2, 32, 227, 3, 2, 2, 2, 34, 238,
	3, 2, 2, 2, 36, 242, 3, 2, 2, 2, 38, 246, 3, 2, 2, 2, 40, 254, 3, 2, 2,
	2, 42, 262, 3, 2, 2, 2, 44, 269, 3, 2, 2, 2, 46, 273, 3, 2, 2, 2, 48, 289,
	3, 2, 2, 2, 50, 293, 3, 2, 2, 2, 52, 295, 3, 2, 2, 2, 54, 297, 3, 2, 2,
	2, 56, 317, 3, 2, 2, 2, 58, 319, 3, 2, 2, 2, 60, 329, 3, 2, 2, 2, 62, 333,
	3, 2, 2, 2, 64, 335, 3, 2, 2, 2, 66, 346, 3, 2, 2, 2, 68, 348, 3, 2, 2,
	2, 70, 359, 3, 2, 2, 2, 72, 362, 3, 2, 2, 2, 74, 373, 3, 2, 2, 2, 76, 375,
	3, 2, 2, 2, 78, 394, 3, 2, 2, 2, 80, 396, 3, 2, 2, 2, 82, 407, 3, 2, 2,
	2, 84, 418, 3, 2, 2, 2, 86, 420, 3, 2, 2, 2, 88, 431, 3, 2, 2, 2, 90, 433,
	3, 2, 2, 2, 92, 442, 3, 2, 2, 2, 94, 444, 3, 2, 2, 2, 96, 455, 3, 2, 2,
	2, 98, 462, 3, 2, 2, 2, 100, 464, 3, 2, 2, 2, 102, 480, 3, 2, 2, 2, 104,
	105, 7, 43, 2, 2, 105, 3, 3, 2, 2, 2, 106, 109, 5, 2, 2, 2, 107, 109, 5,
	28, 15, 2, 108, 106, 3, 2, 2, 2, 108, 107, 3, 2, 2, 2, 109, 5, 3, 2, 2,
	2, 110, 111, 8, 4, 1, 2, 111, 112, 5, 4, 3, 2, 112, 118, 3, 2, 2, 2, 113,
	114, 12, 3, 2, 2, 114, 115, 7, 3, 2, 2, 115, 117, 5, 4, 3, 2, 116, 113,
	3, 2, 2, 2, 117, 120, 3, 2, 2, 2, 118, 116, 3, 2, 2, 2, 118, 119, 3, 2,
	2, 2, 119, 7, 3, 2, 2, 2, 120, 118, 3, 2, 2, 2, 121, 122, 9, 2, 2, 2, 122,
	9, 3, 2, 2, 2, 123, 128, 5, 8, 5, 2, 124, 125, 7, 5, 2, 2, 125, 126, 7,
	6, 2, 2, 126, 128, 5, 10, 6, 2, 127, 123, 3, 2, 2, 2, 127, 124, 3, 2, 2,
	2, 128, 11, 3, 2, 2, 2, 129, 130, 7, 5, 2, 2, 130, 131, 7, 6, 2, 2, 131,
	132, 5, 10, 6, 2, 132, 13, 3, 2, 2, 2, 133, 134, 7, 41, 2, 2, 134, 135,
	7, 7, 2, 2, 135, 136, 5, 12, 7, 2, 136, 137, 7, 3, 2, 2, 137, 138, 5, 18,
	10, 2, 138, 139, 7, 8, 2, 2, 139, 15, 3, 2, 2, 2, 140, 141, 7, 39, 2, 2,
	141, 142, 7, 7, 2, 2, 142, 143, 5, 20, 11, 2, 143, 144, 7, 8, 2, 2, 144,
	17, 3, 2, 2, 2, 145, 146, 8, 10, 1, 2, 146, 160, 7, 44, 2, 2, 147, 160,
	7, 45, 2, 2, 148, 160, 5, 2, 2, 2, 149, 160, 5, 28, 15, 2, 150, 151, 7,
	7, 2, 2, 151, 152, 5, 18, 10, 2, 152, 153, 7, 8, 2, 2, 153, 160, 3, 2,
	2, 2, 154, 155, 7, 13, 2, 2, 155, 156, 7, 7, 2, 2, 156, 157, 5, 18, 10,
	2, 157, 158, 7, 8, 2, 2, 158, 160, 3, 2, 2, 2, 159, 145, 3, 2, 2, 2, 159,
	147, 3, 2, 2, 2, 159, 148, 3, 2, 2, 2, 159, 149, 3, 2, 2, 2, 159, 150,
	3, 2, 2, 2, 159, 154, 3, 2, 2, 2, 160, 169, 3, 2, 2, 2, 161, 162, 12, 6,
	2, 2, 162, 163, 9, 3, 2, 2, 163, 168, 5, 18, 10, 7, 164, 165, 12, 5, 2,
	2, 165, 166, 9, 4, 2, 2, 166, 168, 5, 18, 10, 6, 167, 161, 3, 2, 2, 2,
	167, 164, 3, 2, 2, 2, 168, 171, 3, 2, 2, 2, 169, 167, 3, 2, 2, 2, 169,
	170, 3, 2, 2, 2, 170, 19, 3, 2, 2, 2, 171, 169, 3, 2, 2, 2, 172, 173, 8,
	11, 1, 2, 173, 174, 5, 18, 10, 2, 174, 180, 3, 2, 2, 2, 175, 176, 12, 3,
	2, 2, 176, 177, 7, 3, 2, 2, 177, 179, 5, 18, 10, 2, 178, 175, 3, 2, 2,
	2, 179, 182, 3, 2, 2, 2, 180, 178, 3, 2, 2, 2, 180, 181, 3, 2, 2, 2, 181,
	21, 3, 2, 2, 2, 182, 180, 3, 2, 2, 2, 183, 184, 8, 12, 1, 2, 184, 185,
	5, 18, 10, 2, 185, 186, 9, 5, 2, 2, 186, 187, 5, 18, 10, 2, 187, 195, 3,
	2, 2, 2, 188, 189, 7, 22, 2, 2, 189, 195, 5, 22, 12, 4, 190, 191, 7, 7,
	2, 2, 191, 192, 5, 22, 12, 2, 192, 193, 7, 8, 2, 2, 193, 195, 3, 2, 2,
	2, 194, 183, 3, 2, 2, 2, 194, 188, 3, 2, 2, 2, 194, 190, 3, 2, 2, 2, 195,
	201, 3, 2, 2, 2, 196, 197, 12, 5, 2, 2, 197, 198, 9, 6, 2, 2, 198, 200,
	5, 22, 12, 6, 199, 196, 3, 2, 2, 2, 200, 203, 3, 2, 2, 2, 201, 199, 3,
	2, 2, 2, 201, 202, 3, 2, 2, 2, 202, 23, 3, 2, 2, 2, 203, 201, 3, 2, 2,
	2, 204, 205, 7, 5, 2, 2, 205, 206, 5, 18, 10, 2, 206, 207, 7, 6, 2, 2,
	207, 25, 3, 2, 2, 2, 208, 209, 8, 14, 1, 2, 209, 210, 5, 24, 13, 2, 210,
	215, 3, 2, 2, 2, 211, 212, 12, 3, 2, 2, 212, 214, 5, 24, 13, 2, 213, 211,
	3, 2, 2, 2, 214, 217, 3, 2, 2, 2, 215, 213, 3, 2, 2, 2, 215, 216, 3, 2,
	2, 2, 216, 27, 3, 2, 2, 2, 217, 215, 3, 2, 2, 2, 218, 219, 5, 2, 2, 2,
	219, 220, 5, 26, 14, 2, 220, 29, 3, 2, 2, 2, 221, 226, 5, 18, 10, 2, 222,
	226, 5, 14, 8, 2, 223, 226, 5, 16, 9, 2, 224, 226, 5, 96, 49, 2, 225, 221,
	3, 2, 2, 2, 225, 222, 3, 2, 2, 2, 225, 223, 3, 2, 2, 2, 225, 224, 3, 2,
	2, 2, 226, 31, 3, 2, 2, 2, 227, 228, 8, 17, 1, 2, 228, 229, 5, 30, 16,
	2, 229, 235, 3, 2, 2, 2, 230, 231, 12, 3, 2, 2, 231, 232, 7, 3, 2, 2, 232,
	234, 5, 30, 16, 2, 233, 230, 3, 2, 2, 2, 234, 237, 3, 2, 2, 2, 235, 233,
	3, 2, 2, 2, 235, 236, 3, 2, 2, 2, 236, 33, 3, 2, 2, 2, 237, 235, 3, 2,
	2, 2, 238, 239, 5, 6, 4, 2, 239, 240, 7, 23, 2, 2, 240, 241, 5, 32, 17,
	2, 241, 35, 3, 2, 2, 2, 242, 243, 5, 6, 4, 2, 243, 244, 7, 24, 2, 2, 244,
	245, 5, 32, 17, 2, 245, 37, 3, 2, 2, 2, 246, 247, 7, 25, 2, 2, 247, 248,
	5, 22, 12, 2, 248, 250, 7, 26, 2, 2, 249, 251, 5, 86, 44, 2, 250, 249,
	3, 2, 2, 2, 250, 251, 3, 2, 2, 2, 251, 252, 3, 2, 2, 2, 252, 253, 7, 27,
	2, 2, 253, 39, 3, 2, 2, 2, 254, 255, 7, 28, 2, 2, 255, 256, 5, 22, 12,
	2, 256, 258, 7, 26, 2, 2, 257, 259, 5, 86, 44, 2, 258, 257, 3, 2, 2, 2,
	258, 259, 3, 2, 2, 2, 259, 260, 3, 2, 2, 2, 260, 261, 7, 27, 2, 2, 261,
	41, 3, 2, 2, 2, 262, 263, 7, 29, 2, 2, 263, 265, 7, 26, 2, 2, 264, 266,
	5, 86, 44, 2, 265, 264, 3, 2, 2, 2, 265, 266, 3, 2, 2, 2, 266, 267, 3,
	2, 2, 2, 267, 268, 7, 27, 2, 2, 268, 43, 3, 2, 2, 2, 269, 270, 5, 22, 12,
	2, 270, 271, 7, 30, 2, 2, 271, 272, 5, 84, 43, 2, 272, 45, 3, 2, 2, 2,
	273, 274, 5, 84, 43, 2, 274, 47, 3, 2, 2, 2, 275, 279, 5, 38, 20, 2, 276,
	278, 5, 40, 21, 2, 277, 276, 3, 2, 2, 2, 278, 281, 3, 2, 2, 2, 279, 277,
	3, 2, 2, 2, 279, 280, 3, 2, 2, 2, 280, 283, 3, 2, 2, 2, 281, 279, 3, 2,
	2, 2, 282, 284, 5, 42, 22, 2, 283, 282, 3, 2, 2, 2, 283, 284, 3, 2, 2,
	2, 284, 290, 3, 2, 2, 2, 285, 286, 5, 44, 23, 2, 286, 287, 7, 31, 2, 2,
	287, 288, 5, 46, 24, 2, 288, 290, 3, 2, 2, 2, 289, 275, 3, 2, 2, 2, 289,
	285, 3, 2, 2, 2, 290, 49, 3, 2, 2, 2, 291, 294, 5, 36, 19, 2, 292, 294,
	5, 34, 18, 2, 293, 291, 3, 2, 2, 2, 293, 292, 3, 2, 2, 2, 294, 51, 3, 2,
	2, 2, 295, 296, 5, 34, 18, 2, 296, 53, 3, 2, 2, 2, 297, 299, 7, 32, 2,
	2, 298, 300, 5, 50, 26, 2, 299, 298, 3, 2, 2, 2, 299, 300, 3, 2, 2, 2,
	300, 301, 3, 2, 2, 2, 301, 303, 7, 33, 2, 2, 302, 304, 5, 22, 12, 2, 303,
	302, 3, 2, 2, 2, 303, 304, 3, 2, 2, 2, 304, 305, 3, 2, 2, 2, 305, 307,
	7, 33, 2, 2, 306, 308, 5, 52, 27, 2, 307, 306, 3, 2, 2, 2, 307, 308, 3,
	2, 2, 2, 308, 309, 3, 2, 2, 2, 309, 311, 7, 26, 2, 2, 310, 312, 5, 86,
	44, 2, 311, 310, 3, 2, 2, 2, 311, 312, 3, 2, 2, 2, 312, 313, 3, 2, 2, 2,
	313, 314, 7, 27, 2, 2, 314, 55, 3, 2, 2, 2, 315, 318, 5, 36, 19, 2, 316,
	318, 5, 90, 46, 2, 317, 315, 3, 2, 2, 2, 317, 316, 3, 2, 2, 2, 318, 57,
	3, 2, 2, 2, 319, 320, 8, 30, 1, 2, 320, 321, 5, 56, 29, 2, 321, 326, 3,
	2, 2, 2, 322, 323, 12, 3, 2, 2, 323, 325, 5, 56, 29, 2, 324, 322, 3, 2,
	2, 2, 325, 328, 3, 2, 2, 2, 326, 324, 3, 2, 2, 2, 326, 327, 3, 2, 2, 2,
	327, 59, 3, 2, 2, 2, 328, 326, 3, 2, 2, 2, 329, 330, 5, 58, 30, 2, 330,
	61, 3, 2, 2, 2, 331, 334, 5, 8, 5, 2, 332, 334, 5, 12, 7, 2, 333, 331,
	3, 2, 2, 2, 333, 332, 3, 2, 2, 2, 334, 63, 3, 2, 2, 2, 335, 336, 8, 33,
	1, 2, 336, 337, 5, 62, 32, 2, 337, 343, 3, 2, 2, 2, 338, 339, 12, 3, 2,
	2, 339, 340, 7, 3, 2, 2, 340, 342, 5, 62, 32, 2, 341, 338, 3, 2, 2, 2,
	342, 345, 3, 2, 2, 2, 343, 341, 3, 2, 2, 2, 343, 344, 3, 2, 2, 2, 344,
	65, 3, 2, 2, 2, 345, 343, 3, 2, 2, 2, 346, 347, 5, 2, 2, 2, 347, 67, 3,
	2, 2, 2, 348, 349, 8, 35, 1, 2, 349, 350, 5, 66, 34, 2, 350, 356, 3, 2,
	2, 2, 351, 352, 12, 3, 2, 2, 352, 353, 7, 3, 2, 2, 353, 355, 5, 66, 34,
	2, 354, 351, 3, 2, 2, 2, 355, 358, 3, 2, 2, 2, 356, 354, 3, 2, 2, 2, 356,
	357, 3, 2, 2, 2, 357, 69, 3, 2, 2, 2, 358, 356, 3, 2, 2, 2, 359, 360, 5,
	68, 35, 2, 360, 361, 5, 62, 32, 2, 361, 71, 3, 2, 2, 2, 362, 363, 8, 37,
	1, 2, 363, 364, 5, 70, 36, 2, 364, 370, 3, 2, 2, 2, 365, 366, 12, 3, 2,
	2, 366, 367, 7, 3, 2, 2, 367, 369, 5, 70, 36, 2, 368, 365, 3, 2, 2, 2,
	369, 372, 3, 2, 2, 2, 370, 368, 3, 2, 2, 2, 370, 371, 3, 2, 2, 2, 371,
	73, 3, 2, 2, 2, 372, 370, 3, 2, 2, 2, 373, 374, 5, 2, 2, 2, 374, 75, 3,
	2, 2, 2, 375, 376, 7, 40, 2, 2, 376, 377, 5, 74, 38, 2, 377, 379, 7, 7,
	2, 2, 378, 380, 5, 72, 37, 2, 379, 378, 3, 2, 2, 2, 379, 380, 3, 2, 2,
	2, 380, 381, 3, 2, 2, 2, 381, 387, 7, 8, 2, 2, 382, 384, 7, 7, 2, 2, 383,
	385, 5, 64, 33, 2, 384, 383, 3, 2, 2, 2, 384, 385, 3, 2, 2, 2, 385, 386,
	3, 2, 2, 2, 386, 388, 7, 8, 2, 2, 387, 382, 3, 2, 2, 2, 387, 388, 3, 2,
	2, 2, 388, 77, 3, 2, 2, 2, 389, 395, 5, 18, 10, 2, 390, 395, 5, 22, 12,
	2, 391, 395, 5, 2, 2, 2, 392, 395, 7, 42, 2, 2, 393, 395, 5, 96, 49, 2,
	394, 389, 3, 2, 2, 2, 394, 390, 3, 2, 2, 2, 394, 391, 3, 2, 2, 2, 394,
	392, 3, 2, 2, 2, 394, 393, 3, 2, 2, 2, 395, 79, 3, 2, 2, 2, 396, 397, 8,
	41, 1, 2, 397, 398, 5, 78, 40, 2, 398, 404, 3, 2, 2, 2, 399, 400, 12, 3,
	2, 2, 400, 401, 7, 3, 2, 2, 401, 403, 5, 78, 40, 2, 402, 399, 3, 2, 2,
	2, 403, 406, 3, 2, 2, 2, 404, 402, 3, 2, 2, 2, 404, 405, 3, 2, 2, 2, 405,
	81, 3, 2, 2, 2, 406, 404, 3, 2, 2, 2, 407, 409, 7, 34, 2, 2, 408, 410,
	5, 80, 41, 2, 409, 408, 3, 2, 2, 2, 409, 410, 3, 2, 2, 2, 410, 83, 3, 2,
	2, 2, 411, 419, 5, 34, 18, 2, 412, 419, 5, 36, 19, 2, 413, 419, 5, 48,
	25, 2, 414, 419, 5, 54, 28, 2, 415, 419, 5, 82, 42, 2, 416, 419, 5, 98,
	50, 2, 417, 419, 5, 102, 52, 2, 418, 411, 3, 2, 2, 2, 418, 412, 3, 2, 2,
	2, 418, 413, 3, 2, 2, 2, 418, 414, 3, 2, 2, 2, 418, 415, 3, 2, 2, 2, 418,
	416, 3, 2, 2, 2, 418, 417, 3, 2, 2, 2, 419, 85, 3, 2, 2, 2, 420, 421, 8,
	44, 1, 2, 421, 422, 5, 84, 43, 2, 422, 427, 3, 2, 2, 2, 423, 424, 12, 3,
	2, 2, 424, 426, 5, 84, 43, 2, 425, 423, 3, 2, 2, 2, 426, 429, 3, 2, 2,
	2, 427, 425, 3, 2, 2, 2, 427, 428, 3, 2, 2, 2, 428, 87, 3, 2, 2, 2, 429,
	427, 3, 2, 2, 2, 430, 432, 5, 86, 44, 2, 431, 430, 3, 2, 2, 2, 431, 432,
	3, 2, 2, 2, 432, 89, 3, 2, 2, 2, 433, 434, 5, 76, 39, 2, 434, 436, 7, 26,
	2, 2, 435, 437, 5, 88, 45, 2, 436, 435, 3, 2, 2, 2, 436, 437, 3, 2, 2,
	2, 437, 438, 3, 2, 2, 2, 438, 439, 7, 27, 2, 2, 439, 91, 3, 2, 2, 2, 440,
	443, 5, 18, 10, 2, 441, 443, 5, 22, 12, 2, 442, 440, 3, 2, 2, 2, 442, 441,
	3, 2, 2, 2, 443, 93, 3, 2, 2, 2, 444, 445, 8, 48, 1, 2, 445, 446, 5, 92,
	47, 2, 446, 452, 3, 2, 2, 2, 447, 448, 12, 3, 2, 2, 448, 449, 7, 3, 2,
	2, 449, 451, 5, 92, 47, 2, 450, 447, 3, 2, 2, 2, 451, 454, 3, 2, 2, 2,
	452, 450, 3, 2, 2, 2, 452, 453, 3, 2, 2, 2, 453, 95, 3, 2, 2, 2, 454, 452,
	3, 2, 2, 2, 455, 456, 5, 74, 38, 2, 456, 458, 7, 7, 2, 2, 457, 459, 5,
	94, 48, 2, 458, 457, 3, 2, 2, 2, 458, 459, 3, 2, 2, 2, 459, 460, 3, 2,
	2, 2, 460, 461, 7, 8, 2, 2, 461, 97, 3, 2, 2, 2, 462, 463, 5, 96, 49, 2,
	463, 99, 3, 2, 2, 2, 464, 467, 8, 51, 1, 2, 465, 468, 5, 18, 10, 2, 466,
	468, 5, 22, 12, 2, 467, 465, 3, 2, 2, 2, 467, 466, 3, 2, 2, 2, 468, 477,
	3, 2, 2, 2, 469, 470, 12, 3, 2, 2, 470, 473, 7, 3, 2, 2, 471, 474, 5, 18,
	10, 2, 472, 474, 5, 22, 12, 2, 473, 471, 3, 2, 2, 2, 473, 472, 3, 2, 2,
	2, 474, 476, 3, 2, 2, 2, 475, 469, 3, 2, 2, 2, 476, 479, 3, 2, 2, 2, 477,
	475, 3, 2, 2, 2, 477, 478, 3, 2, 2, 2, 478, 101, 3, 2, 2, 2, 479, 477,
	3, 2, 2, 2, 480, 481, 7, 35, 2, 2, 481, 482, 7, 7, 2, 2, 482, 483, 5, 100,
	51, 2, 483, 484, 7, 8, 2, 2, 484, 103, 3, 2, 2, 2, 47, 108, 118, 127, 159,
	167, 169, 180, 194, 201, 215, 225, 235, 250, 258, 265, 279, 283, 289, 293,
	299, 303, 307, 311, 317, 326, 333, 343, 356, 370, 379, 384, 387, 394, 404,
	409, 418, 427, 431, 436, 442, 452, 458, 467, 473, 477,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "','", "'string'", "'['", "']'", "'('", "')'", "'*'", "'/'", "'+'",
	"'-'", "'transpose'", "'=='", "'<'", "'>'", "'<='", "'>='", "'!='", "'&&'",
	"'||'", "'!'", "'='", "':='", "'if'", "'{'", "'}'", "'elsif'", "'else'",
	"'?'", "':'", "'for'", "';'", "'return'", "'print'", "'var'", "'int'",
	"'float'", "'matrix'", "'func'", "'list'", "'nil'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "Var",
	"Int", "Float", "Matrix", "Func", "List", "Nil", "Identifier", "IntegerLiteral",
	"FloatLiteral", "WS",
}

var ruleNames = []string{
	"identifier", "declarator", "declaratorList", "simpleTypeSpecifier", "listElementTypeSpecifier",
	"listTypeSpecifier", "listInitExpression", "matrixInitExpression", "aExpr",
	"aExprList", "bExpr", "listElementIndex", "listElementIndexList", "listElementExpression",
	"assignInit", "assignInitList", "assignStatement", "assignInitStatement",
	"ifExpr", "elsifExpr", "elseExpr", "ternaryIfExpr", "ternaryElseExpr",
	"selectionStatement", "iterationAssignInitStatement", "iterationAssignStatement",
	"iterationStatement", "definition", "definitionList", "file", "typeSpecifier",
	"typeSpecifierList", "paraDeclarator", "paraDeclaratorList", "paraDeclaratorWithIdentity",
	"paraDeclaratorWithIdentityList", "funcIdentifier", "funcTypeSpecifierWithName",
	"funcReturnPara", "funcReturnParaList", "funcReturnStatement", "funcStatement",
	"funcStatementList", "funcBody", "funcDefinition", "funcExecutePara", "funcExecuteParaList",
	"funcExecuteExpression", "funcExecuteStatement", "printList", "printStatement",
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
	ZZParserT__30          = 31
	ZZParserT__31          = 32
	ZZParserT__32          = 33
	ZZParserVar            = 34
	ZZParserInt            = 35
	ZZParserFloat          = 36
	ZZParserMatrix         = 37
	ZZParserFunc           = 38
	ZZParserList           = 39
	ZZParserNil            = 40
	ZZParserIdentifier     = 41
	ZZParserIntegerLiteral = 42
	ZZParserFloatLiteral   = 43
	ZZParserWS             = 44
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
	ZZParserRULE_matrixInitExpression           = 7
	ZZParserRULE_aExpr                          = 8
	ZZParserRULE_aExprList                      = 9
	ZZParserRULE_bExpr                          = 10
	ZZParserRULE_listElementIndex               = 11
	ZZParserRULE_listElementIndexList           = 12
	ZZParserRULE_listElementExpression          = 13
	ZZParserRULE_assignInit                     = 14
	ZZParserRULE_assignInitList                 = 15
	ZZParserRULE_assignStatement                = 16
	ZZParserRULE_assignInitStatement            = 17
	ZZParserRULE_ifExpr                         = 18
	ZZParserRULE_elsifExpr                      = 19
	ZZParserRULE_elseExpr                       = 20
	ZZParserRULE_ternaryIfExpr                  = 21
	ZZParserRULE_ternaryElseExpr                = 22
	ZZParserRULE_selectionStatement             = 23
	ZZParserRULE_iterationAssignInitStatement   = 24
	ZZParserRULE_iterationAssignStatement       = 25
	ZZParserRULE_iterationStatement             = 26
	ZZParserRULE_definition                     = 27
	ZZParserRULE_definitionList                 = 28
	ZZParserRULE_file                           = 29
	ZZParserRULE_typeSpecifier                  = 30
	ZZParserRULE_typeSpecifierList              = 31
	ZZParserRULE_paraDeclarator                 = 32
	ZZParserRULE_paraDeclaratorList             = 33
	ZZParserRULE_paraDeclaratorWithIdentity     = 34
	ZZParserRULE_paraDeclaratorWithIdentityList = 35
	ZZParserRULE_funcIdentifier                 = 36
	ZZParserRULE_funcTypeSpecifierWithName      = 37
	ZZParserRULE_funcReturnPara                 = 38
	ZZParserRULE_funcReturnParaList             = 39
	ZZParserRULE_funcReturnStatement            = 40
	ZZParserRULE_funcStatement                  = 41
	ZZParserRULE_funcStatementList              = 42
	ZZParserRULE_funcBody                       = 43
	ZZParserRULE_funcDefinition                 = 44
	ZZParserRULE_funcExecutePara                = 45
	ZZParserRULE_funcExecuteParaList            = 46
	ZZParserRULE_funcExecuteExpression          = 47
	ZZParserRULE_funcExecuteStatement           = 48
	ZZParserRULE_printList                      = 49
	ZZParserRULE_printStatement                 = 50
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
		p.SetState(102)
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

	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(104)
			p.Identifier()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(105)
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
		p.SetState(109)
		p.Declarator()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(116)
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
			p.SetState(111)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(112)
				p.Match(ZZParserT__0)
			}
			{
				p.SetState(113)
				p.Declarator()
			}

		}
		p.SetState(118)
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

func (s *SimpleTypeSpecifierContext) Matrix() antlr.TerminalNode {
	return s.GetToken(ZZParserMatrix, 0)
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
		p.SetState(119)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ZZParserT__1 || (((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(ZZParserInt-35))|(1<<(ZZParserFloat-35))|(1<<(ZZParserMatrix-35)))) != 0)) {
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

	p.SetState(125)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ZZParserT__1, ZZParserInt, ZZParserFloat, ZZParserMatrix:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(121)
			p.SimpleTypeSpecifier()
		}

	case ZZParserT__2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(122)
			p.Match(ZZParserT__2)
		}
		{
			p.SetState(123)
			p.Match(ZZParserT__3)
		}
		{
			p.SetState(124)
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
		p.SetState(127)
		p.Match(ZZParserT__2)
	}
	{
		p.SetState(128)
		p.Match(ZZParserT__3)
	}
	{
		p.SetState(129)
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
		p.SetState(131)
		p.Match(ZZParserList)
	}
	{
		p.SetState(132)
		p.Match(ZZParserT__4)
	}
	{
		p.SetState(133)
		p.ListTypeSpecifier()
	}
	{
		p.SetState(134)
		p.Match(ZZParserT__0)
	}
	{
		p.SetState(135)
		p.aExpr(0)
	}
	{
		p.SetState(136)
		p.Match(ZZParserT__5)
	}

	return localctx
}

// IMatrixInitExpressionContext is an interface to support dynamic dispatch.
type IMatrixInitExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMatrixInitExpressionContext differentiates from other interfaces.
	IsMatrixInitExpressionContext()
}

type MatrixInitExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatrixInitExpressionContext() *MatrixInitExpressionContext {
	var p = new(MatrixInitExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_matrixInitExpression
	return p
}

func (*MatrixInitExpressionContext) IsMatrixInitExpressionContext() {}

func NewMatrixInitExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatrixInitExpressionContext {
	var p = new(MatrixInitExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_matrixInitExpression

	return p
}

func (s *MatrixInitExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *MatrixInitExpressionContext) Matrix() antlr.TerminalNode {
	return s.GetToken(ZZParserMatrix, 0)
}

func (s *MatrixInitExpressionContext) AExprList() IAExprListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAExprListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAExprListContext)
}

func (s *MatrixInitExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatrixInitExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MatrixInitExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterMatrixInitExpression(s)
	}
}

func (s *MatrixInitExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitMatrixInitExpression(s)
	}
}

func (p *ZZParser) MatrixInitExpression() (localctx IMatrixInitExpressionContext) {
	localctx = NewMatrixInitExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ZZParserRULE_matrixInitExpression)

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
		p.SetState(138)
		p.Match(ZZParserMatrix)
	}
	{
		p.SetState(139)
		p.Match(ZZParserT__4)
	}
	{
		p.SetState(140)
		p.aExprList(0)
	}
	{
		p.SetState(141)
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

type AExp_transposeContext struct {
	*AExprContext
}

func NewAExp_transposeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AExp_transposeContext {
	var p = new(AExp_transposeContext)

	p.AExprContext = NewEmptyAExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AExprContext))

	return p
}

func (s *AExp_transposeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AExp_transposeContext) AExpr() IAExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAExprContext)
}

func (s *AExp_transposeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterAExp_transpose(s)
	}
}

func (s *AExp_transposeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitAExp_transpose(s)
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
	_startState := 16
	p.EnterRecursionRule(localctx, 16, ZZParserRULE_aExpr, _p)
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
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAExp_IntergerLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(144)
			p.Match(ZZParserIntegerLiteral)
		}

	case 2:
		localctx = NewAExp_FloatLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(145)
			p.Match(ZZParserFloatLiteral)
		}

	case 3:
		localctx = NewAExp_identifierContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(146)
			p.Identifier()
		}

	case 4:
		localctx = NewAExp_listElementExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(147)
			p.ListElementExpression()
		}

	case 5:
		localctx = NewAExp_bracketExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(148)
			p.Match(ZZParserT__4)
		}
		{
			p.SetState(149)
			p.aExpr(0)
		}
		{
			p.SetState(150)
			p.Match(ZZParserT__5)
		}

	case 6:
		localctx = NewAExp_transposeContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(152)
			p.Match(ZZParserT__10)
		}
		{
			p.SetState(153)
			p.Match(ZZParserT__4)
		}
		{
			p.SetState(154)
			p.aExpr(0)
		}
		{
			p.SetState(155)
			p.Match(ZZParserT__5)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(165)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAExp_multiplicativeExpressionContext(p, NewAExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ZZParserRULE_aExpr)
				p.SetState(159)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(160)
					_la = p.GetTokenStream().LA(1)

					if !(_la == ZZParserT__6 || _la == ZZParserT__7) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(161)
					p.aExpr(5)
				}

			case 2:
				localctx = NewAExp_additiveExpressionContext(p, NewAExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ZZParserRULE_aExpr)
				p.SetState(162)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(163)
					_la = p.GetTokenStream().LA(1)

					if !(_la == ZZParserT__8 || _la == ZZParserT__9) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(164)
					p.aExpr(4)
				}

			}

		}
		p.SetState(169)
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
	_startState := 18
	p.EnterRecursionRule(localctx, 18, ZZParserRULE_aExprList, _p)

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
		p.SetState(171)
		p.aExpr(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(178)
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
			p.SetState(173)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(174)
				p.Match(ZZParserT__0)
			}
			{
				p.SetState(175)
				p.aExpr(0)
			}

		}
		p.SetState(180)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
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
	_startState := 20
	p.EnterRecursionRule(localctx, 20, ZZParserRULE_bExpr, _p)
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
	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		localctx = NewBExpr_aExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(182)
			p.aExpr(0)
		}
		{
			p.SetState(183)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ZZParserT__11)|(1<<ZZParserT__12)|(1<<ZZParserT__13)|(1<<ZZParserT__14)|(1<<ZZParserT__15)|(1<<ZZParserT__16))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(184)
			p.aExpr(0)
		}

	case 2:
		localctx = NewBExpr_bangContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(186)
			p.Match(ZZParserT__19)
		}
		{
			p.SetState(187)
			p.bExpr(2)
		}

	case 3:
		localctx = NewBExpr_bracketExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(188)
			p.Match(ZZParserT__4)
		}
		{
			p.SetState(189)
			p.bExpr(0)
		}
		{
			p.SetState(190)
			p.Match(ZZParserT__5)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewBExpr_bExprContext(p, NewBExprContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, ZZParserRULE_bExpr)
			p.SetState(194)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
			}
			{
				p.SetState(195)
				_la = p.GetTokenStream().LA(1)

				if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ZZParserT__11)|(1<<ZZParserT__16)|(1<<ZZParserT__17)|(1<<ZZParserT__18))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(196)
				p.bExpr(4)
			}

		}
		p.SetState(201)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 22, ZZParserRULE_listElementIndex)

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
		p.SetState(202)
		p.Match(ZZParserT__2)
	}
	{
		p.SetState(203)
		p.aExpr(0)
	}
	{
		p.SetState(204)
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
	_startState := 24
	p.EnterRecursionRule(localctx, 24, ZZParserRULE_listElementIndexList, _p)

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
		p.SetState(207)
		p.ListElementIndex()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(213)
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
			p.SetState(209)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(210)
				p.ListElementIndex()
			}

		}
		p.SetState(215)
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
	p.EnterRule(localctx, 26, ZZParserRULE_listElementExpression)

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
		p.SetState(216)
		p.Identifier()
	}
	{
		p.SetState(217)
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

func (s *AssignInitContext) MatrixInitExpression() IMatrixInitExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMatrixInitExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMatrixInitExpressionContext)
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
	p.EnterRule(localctx, 28, ZZParserRULE_assignInit)

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

	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(219)
			p.aExpr(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(220)
			p.ListInitExpression()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(221)
			p.MatrixInitExpression()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(222)
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
	_startState := 30
	p.EnterRecursionRule(localctx, 30, ZZParserRULE_assignInitList, _p)

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
		p.SetState(226)
		p.AssignInit()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(233)
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
			p.SetState(228)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(229)
				p.Match(ZZParserT__0)
			}
			{
				p.SetState(230)
				p.AssignInit()
			}

		}
		p.SetState(235)
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
	p.EnterRule(localctx, 32, ZZParserRULE_assignStatement)

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
		p.SetState(236)
		p.declaratorList(0)
	}
	{
		p.SetState(237)
		p.Match(ZZParserT__20)
	}
	{
		p.SetState(238)
		p.assignInitList(0)
	}

	return localctx
}

// IAssignInitStatementContext is an interface to support dynamic dispatch.
type IAssignInitStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignInitStatementContext differentiates from other interfaces.
	IsAssignInitStatementContext()
}

type AssignInitStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignInitStatementContext() *AssignInitStatementContext {
	var p = new(AssignInitStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_assignInitStatement
	return p
}

func (*AssignInitStatementContext) IsAssignInitStatementContext() {}

func NewAssignInitStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignInitStatementContext {
	var p = new(AssignInitStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_assignInitStatement

	return p
}

func (s *AssignInitStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignInitStatementContext) DeclaratorList() IDeclaratorListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclaratorListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclaratorListContext)
}

func (s *AssignInitStatementContext) AssignInitList() IAssignInitListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignInitListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignInitListContext)
}

func (s *AssignInitStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignInitStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignInitStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterAssignInitStatement(s)
	}
}

func (s *AssignInitStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitAssignInitStatement(s)
	}
}

func (p *ZZParser) AssignInitStatement() (localctx IAssignInitStatementContext) {
	localctx = NewAssignInitStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ZZParserRULE_assignInitStatement)

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
		p.SetState(240)
		p.declaratorList(0)
	}
	{
		p.SetState(241)
		p.Match(ZZParserT__21)
	}
	{
		p.SetState(242)
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
	p.EnterRule(localctx, 36, ZZParserRULE_ifExpr)
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
		p.SetState(244)
		p.Match(ZZParserT__22)
	}
	{
		p.SetState(245)
		p.bExpr(0)
	}
	{
		p.SetState(246)
		p.Match(ZZParserT__23)
	}
	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ZZParserT__4)|(1<<ZZParserT__10)|(1<<ZZParserT__19)|(1<<ZZParserT__22)|(1<<ZZParserT__29))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(ZZParserT__31-32))|(1<<(ZZParserT__32-32))|(1<<(ZZParserIdentifier-32))|(1<<(ZZParserIntegerLiteral-32))|(1<<(ZZParserFloatLiteral-32)))) != 0) {
		{
			p.SetState(247)
			p.funcStatementList(0)
		}

	}
	{
		p.SetState(250)
		p.Match(ZZParserT__24)
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
	p.EnterRule(localctx, 38, ZZParserRULE_elsifExpr)
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
		p.SetState(252)
		p.Match(ZZParserT__25)
	}
	{
		p.SetState(253)
		p.bExpr(0)
	}
	{
		p.SetState(254)
		p.Match(ZZParserT__23)
	}
	p.SetState(256)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ZZParserT__4)|(1<<ZZParserT__10)|(1<<ZZParserT__19)|(1<<ZZParserT__22)|(1<<ZZParserT__29))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(ZZParserT__31-32))|(1<<(ZZParserT__32-32))|(1<<(ZZParserIdentifier-32))|(1<<(ZZParserIntegerLiteral-32))|(1<<(ZZParserFloatLiteral-32)))) != 0) {
		{
			p.SetState(255)
			p.funcStatementList(0)
		}

	}
	{
		p.SetState(258)
		p.Match(ZZParserT__24)
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
	p.EnterRule(localctx, 40, ZZParserRULE_elseExpr)
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
		p.Match(ZZParserT__26)
	}
	{
		p.SetState(261)
		p.Match(ZZParserT__23)
	}
	p.SetState(263)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ZZParserT__4)|(1<<ZZParserT__10)|(1<<ZZParserT__19)|(1<<ZZParserT__22)|(1<<ZZParserT__29))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(ZZParserT__31-32))|(1<<(ZZParserT__32-32))|(1<<(ZZParserIdentifier-32))|(1<<(ZZParserIntegerLiteral-32))|(1<<(ZZParserFloatLiteral-32)))) != 0) {
		{
			p.SetState(262)
			p.funcStatementList(0)
		}

	}
	{
		p.SetState(265)
		p.Match(ZZParserT__24)
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
	p.EnterRule(localctx, 42, ZZParserRULE_ternaryIfExpr)

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
		p.SetState(267)
		p.bExpr(0)
	}
	{
		p.SetState(268)
		p.Match(ZZParserT__27)
	}
	{
		p.SetState(269)
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
	p.EnterRule(localctx, 44, ZZParserRULE_ternaryElseExpr)

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
		p.SetState(271)
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
	p.EnterRule(localctx, 46, ZZParserRULE_selectionStatement)

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

	p.SetState(287)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ZZParserT__22:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(273)
			p.IfExpr()
		}
		p.SetState(277)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(274)
					p.ElsifExpr()
				}

			}
			p.SetState(279)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
		}
		p.SetState(281)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(280)
				p.ElseExpr()
			}

		}

	case ZZParserT__4, ZZParserT__10, ZZParserT__19, ZZParserIdentifier, ZZParserIntegerLiteral, ZZParserFloatLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(283)
			p.TernaryIfExpr()
		}
		{
			p.SetState(284)
			p.Match(ZZParserT__28)
		}
		{
			p.SetState(285)
			p.TernaryElseExpr()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIterationAssignInitStatementContext is an interface to support dynamic dispatch.
type IIterationAssignInitStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIterationAssignInitStatementContext differentiates from other interfaces.
	IsIterationAssignInitStatementContext()
}

type IterationAssignInitStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIterationAssignInitStatementContext() *IterationAssignInitStatementContext {
	var p = new(IterationAssignInitStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_iterationAssignInitStatement
	return p
}

func (*IterationAssignInitStatementContext) IsIterationAssignInitStatementContext() {}

func NewIterationAssignInitStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IterationAssignInitStatementContext {
	var p = new(IterationAssignInitStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_iterationAssignInitStatement

	return p
}

func (s *IterationAssignInitStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IterationAssignInitStatementContext) AssignInitStatement() IAssignInitStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignInitStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignInitStatementContext)
}

func (s *IterationAssignInitStatementContext) AssignStatement() IAssignStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignStatementContext)
}

func (s *IterationAssignInitStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IterationAssignInitStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IterationAssignInitStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterIterationAssignInitStatement(s)
	}
}

func (s *IterationAssignInitStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitIterationAssignInitStatement(s)
	}
}

func (p *ZZParser) IterationAssignInitStatement() (localctx IIterationAssignInitStatementContext) {
	localctx = NewIterationAssignInitStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, ZZParserRULE_iterationAssignInitStatement)

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

	p.SetState(291)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(289)
			p.AssignInitStatement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(290)
			p.AssignStatement()
		}

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
	p.EnterRule(localctx, 50, ZZParserRULE_iterationAssignStatement)

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
		p.SetState(293)
		p.AssignStatement()
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

func (s *IterationStatementContext) IterationAssignInitStatement() IIterationAssignInitStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIterationAssignInitStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIterationAssignInitStatementContext)
}

func (s *IterationStatementContext) BExpr() IBExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBExprContext)
}

func (s *IterationStatementContext) IterationAssignStatement() IIterationAssignStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIterationAssignStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIterationAssignStatementContext)
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
	p.EnterRule(localctx, 52, ZZParserRULE_iterationStatement)
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
		p.SetState(295)
		p.Match(ZZParserT__29)
	}
	p.SetState(297)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ZZParserIdentifier {
		{
			p.SetState(296)
			p.IterationAssignInitStatement()
		}

	}
	{
		p.SetState(299)
		p.Match(ZZParserT__30)
	}
	p.SetState(301)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ZZParserT__4)|(1<<ZZParserT__10)|(1<<ZZParserT__19))) != 0) || (((_la-41)&-(0x1f+1)) == 0 && ((1<<uint((_la-41)))&((1<<(ZZParserIdentifier-41))|(1<<(ZZParserIntegerLiteral-41))|(1<<(ZZParserFloatLiteral-41)))) != 0) {
		{
			p.SetState(300)
			p.bExpr(0)
		}

	}
	{
		p.SetState(303)
		p.Match(ZZParserT__30)
	}
	p.SetState(305)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ZZParserIdentifier {
		{
			p.SetState(304)
			p.IterationAssignStatement()
		}

	}
	{
		p.SetState(307)
		p.Match(ZZParserT__23)
	}
	p.SetState(309)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ZZParserT__4)|(1<<ZZParserT__10)|(1<<ZZParserT__19)|(1<<ZZParserT__22)|(1<<ZZParserT__29))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(ZZParserT__31-32))|(1<<(ZZParserT__32-32))|(1<<(ZZParserIdentifier-32))|(1<<(ZZParserIntegerLiteral-32))|(1<<(ZZParserFloatLiteral-32)))) != 0) {
		{
			p.SetState(308)
			p.funcStatementList(0)
		}

	}
	{
		p.SetState(311)
		p.Match(ZZParserT__24)
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

func (s *DefinitionContext) AssignInitStatement() IAssignInitStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignInitStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignInitStatementContext)
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
	p.EnterRule(localctx, 54, ZZParserRULE_definition)

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

	p.SetState(315)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ZZParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(313)
			p.AssignInitStatement()
		}

	case ZZParserFunc:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(314)
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
	_startState := 56
	p.EnterRecursionRule(localctx, 56, ZZParserRULE_definitionList, _p)

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
		p.SetState(318)
		p.Definition()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(324)
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
			p.SetState(320)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(321)
				p.Definition()
			}

		}
		p.SetState(326)
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
	p.EnterRule(localctx, 58, ZZParserRULE_file)

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
		p.SetState(327)
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
	p.EnterRule(localctx, 60, ZZParserRULE_typeSpecifier)

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

	p.SetState(331)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ZZParserT__1, ZZParserInt, ZZParserFloat, ZZParserMatrix:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(329)
			p.SimpleTypeSpecifier()
		}

	case ZZParserT__2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(330)
			p.ListTypeSpecifier()
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
	_startState := 62
	p.EnterRecursionRule(localctx, 62, ZZParserRULE_typeSpecifierList, _p)

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
		p.SetState(334)
		p.TypeSpecifier()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(341)
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
			p.SetState(336)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(337)
				p.Match(ZZParserT__0)
			}
			{
				p.SetState(338)
				p.TypeSpecifier()
			}

		}
		p.SetState(343)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 64, ZZParserRULE_paraDeclarator)

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
		p.SetState(344)
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
	_startState := 66
	p.EnterRecursionRule(localctx, 66, ZZParserRULE_paraDeclaratorList, _p)

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
		p.SetState(347)
		p.ParaDeclarator()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(354)
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
			p.SetState(349)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(350)
				p.Match(ZZParserT__0)
			}
			{
				p.SetState(351)
				p.ParaDeclarator()
			}

		}
		p.SetState(356)
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
	p.EnterRule(localctx, 68, ZZParserRULE_paraDeclaratorWithIdentity)

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
		p.SetState(357)
		p.paraDeclaratorList(0)
	}
	{
		p.SetState(358)
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
	_startState := 70
	p.EnterRecursionRule(localctx, 70, ZZParserRULE_paraDeclaratorWithIdentityList, _p)

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
		p.SetState(361)
		p.ParaDeclaratorWithIdentity()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(368)
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
			p.SetState(363)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(364)
				p.Match(ZZParserT__0)
			}
			{
				p.SetState(365)
				p.ParaDeclaratorWithIdentity()
			}

		}
		p.SetState(370)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 72, ZZParserRULE_funcIdentifier)

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
		p.SetState(371)
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
	p.EnterRule(localctx, 74, ZZParserRULE_funcTypeSpecifierWithName)
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
		p.SetState(373)
		p.Match(ZZParserFunc)
	}
	{
		p.SetState(374)
		p.FuncIdentifier()
	}
	{
		p.SetState(375)
		p.Match(ZZParserT__4)
	}
	p.SetState(377)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ZZParserIdentifier {
		{
			p.SetState(376)
			p.paraDeclaratorWithIdentityList(0)
		}

	}
	{
		p.SetState(379)
		p.Match(ZZParserT__5)
	}
	p.SetState(385)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ZZParserT__4 {
		{
			p.SetState(380)
			p.Match(ZZParserT__4)
		}
		p.SetState(382)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ZZParserT__1 || _la == ZZParserT__2 || (((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(ZZParserInt-35))|(1<<(ZZParserFloat-35))|(1<<(ZZParserMatrix-35)))) != 0) {
			{
				p.SetState(381)
				p.typeSpecifierList(0)
			}

		}
		{
			p.SetState(384)
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

func (s *FuncReturnParaContext) FuncExecuteExpression() IFuncExecuteExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncExecuteExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncExecuteExpressionContext)
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
	p.EnterRule(localctx, 76, ZZParserRULE_funcReturnPara)

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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(387)
			p.aExpr(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(388)
			p.bExpr(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(389)
			p.Identifier()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(390)
			p.Match(ZZParserNil)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(391)
			p.FuncExecuteExpression()
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
	_startState := 78
	p.EnterRecursionRule(localctx, 78, ZZParserRULE_funcReturnParaList, _p)

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
		p.FuncReturnPara()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(402)
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
			p.SetState(397)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(398)
				p.Match(ZZParserT__0)
			}
			{
				p.SetState(399)
				p.FuncReturnPara()
			}

		}
		p.SetState(404)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 80, ZZParserRULE_funcReturnStatement)

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
		p.SetState(405)
		p.Match(ZZParserT__31)
	}
	p.SetState(407)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(406)
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

func (s *FuncStatementContext) AssignInitStatement() IAssignInitStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignInitStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignInitStatementContext)
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

func (s *FuncStatementContext) PrintStatement() IPrintStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrintStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrintStatementContext)
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
	p.EnterRule(localctx, 82, ZZParserRULE_funcStatement)

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

	p.SetState(416)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(409)
			p.AssignStatement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(410)
			p.AssignInitStatement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(411)
			p.SelectionStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(412)
			p.IterationStatement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(413)
			p.FuncReturnStatement()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(414)
			p.FuncExecuteStatement()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(415)
			p.PrintStatement()
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
	_startState := 84
	p.EnterRecursionRule(localctx, 84, ZZParserRULE_funcStatementList, _p)

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
		p.SetState(419)
		p.FuncStatement()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(425)
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
			p.SetState(421)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(422)
				p.FuncStatement()
			}

		}
		p.SetState(427)
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
	p.EnterRule(localctx, 86, ZZParserRULE_funcBody)
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
	p.SetState(429)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ZZParserT__4)|(1<<ZZParserT__10)|(1<<ZZParserT__19)|(1<<ZZParserT__22)|(1<<ZZParserT__29))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(ZZParserT__31-32))|(1<<(ZZParserT__32-32))|(1<<(ZZParserIdentifier-32))|(1<<(ZZParserIntegerLiteral-32))|(1<<(ZZParserFloatLiteral-32)))) != 0) {
		{
			p.SetState(428)
			p.funcStatementList(0)
		}

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
	p.EnterRule(localctx, 88, ZZParserRULE_funcDefinition)

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
		p.SetState(431)
		p.FuncTypeSpecifierWithName()
	}
	{
		p.SetState(432)
		p.Match(ZZParserT__23)
	}
	p.SetState(434)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(433)
			p.FuncBody()
		}

	}
	{
		p.SetState(436)
		p.Match(ZZParserT__24)
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
	p.EnterRule(localctx, 90, ZZParserRULE_funcExecutePara)

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

	p.SetState(440)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(438)
			p.aExpr(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(439)
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
	_startState := 92
	p.EnterRecursionRule(localctx, 92, ZZParserRULE_funcExecuteParaList, _p)

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
		p.SetState(443)
		p.FuncExecutePara()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(450)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewFuncExecuteParaListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ZZParserRULE_funcExecuteParaList)
			p.SetState(445)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(446)
				p.Match(ZZParserT__0)
			}
			{
				p.SetState(447)
				p.FuncExecutePara()
			}

		}
		p.SetState(452)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 94, ZZParserRULE_funcExecuteExpression)
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
		p.SetState(453)
		p.FuncIdentifier()
	}
	{
		p.SetState(454)
		p.Match(ZZParserT__4)
	}
	p.SetState(456)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ZZParserT__4)|(1<<ZZParserT__10)|(1<<ZZParserT__19))) != 0) || (((_la-41)&-(0x1f+1)) == 0 && ((1<<uint((_la-41)))&((1<<(ZZParserIdentifier-41))|(1<<(ZZParserIntegerLiteral-41))|(1<<(ZZParserFloatLiteral-41)))) != 0) {
		{
			p.SetState(455)
			p.funcExecuteParaList(0)
		}

	}
	{
		p.SetState(458)
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
	p.EnterRule(localctx, 96, ZZParserRULE_funcExecuteStatement)

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
		p.SetState(460)
		p.FuncExecuteExpression()
	}

	return localctx
}

// IPrintListContext is an interface to support dynamic dispatch.
type IPrintListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrintListContext differentiates from other interfaces.
	IsPrintListContext()
}

type PrintListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintListContext() *PrintListContext {
	var p = new(PrintListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_printList
	return p
}

func (*PrintListContext) IsPrintListContext() {}

func NewPrintListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintListContext {
	var p = new(PrintListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_printList

	return p
}

func (s *PrintListContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintListContext) AExpr() IAExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAExprContext)
}

func (s *PrintListContext) BExpr() IBExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBExprContext)
}

func (s *PrintListContext) PrintList() IPrintListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrintListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrintListContext)
}

func (s *PrintListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterPrintList(s)
	}
}

func (s *PrintListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitPrintList(s)
	}
}

func (p *ZZParser) PrintList() (localctx IPrintListContext) {
	return p.printList(0)
}

func (p *ZZParser) printList(_p int) (localctx IPrintListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewPrintListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPrintListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 98
	p.EnterRecursionRule(localctx, 98, ZZParserRULE_printList, _p)

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
	p.SetState(465)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(463)
			p.aExpr(0)
		}

	case 2:
		{
			p.SetState(464)
			p.bExpr(0)
		}

	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(475)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewPrintListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ZZParserRULE_printList)
			p.SetState(467)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(468)
				p.Match(ZZParserT__0)
			}
			p.SetState(471)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(469)
					p.aExpr(0)
				}

			case 2:
				{
					p.SetState(470)
					p.bExpr(0)
				}

			}

		}
		p.SetState(477)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext())
	}

	return localctx
}

// IPrintStatementContext is an interface to support dynamic dispatch.
type IPrintStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrintStatementContext differentiates from other interfaces.
	IsPrintStatementContext()
}

type PrintStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintStatementContext() *PrintStatementContext {
	var p = new(PrintStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZZParserRULE_printStatement
	return p
}

func (*PrintStatementContext) IsPrintStatementContext() {}

func NewPrintStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintStatementContext {
	var p = new(PrintStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZZParserRULE_printStatement

	return p
}

func (s *PrintStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintStatementContext) PrintList() IPrintListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrintListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrintListContext)
}

func (s *PrintStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.EnterPrintStatement(s)
	}
}

func (s *PrintStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZZListener); ok {
		listenerT.ExitPrintStatement(s)
	}
}

func (p *ZZParser) PrintStatement() (localctx IPrintStatementContext) {
	localctx = NewPrintStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, ZZParserRULE_printStatement)

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
		p.SetState(478)
		p.Match(ZZParserT__32)
	}
	{
		p.SetState(479)
		p.Match(ZZParserT__4)
	}
	{
		p.SetState(480)
		p.printList(0)
	}
	{
		p.SetState(481)
		p.Match(ZZParserT__5)
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

	case 8:
		var t *AExprContext = nil
		if localctx != nil {
			t = localctx.(*AExprContext)
		}
		return p.AExpr_Sempred(t, predIndex)

	case 9:
		var t *AExprListContext = nil
		if localctx != nil {
			t = localctx.(*AExprListContext)
		}
		return p.AExprList_Sempred(t, predIndex)

	case 10:
		var t *BExprContext = nil
		if localctx != nil {
			t = localctx.(*BExprContext)
		}
		return p.BExpr_Sempred(t, predIndex)

	case 12:
		var t *ListElementIndexListContext = nil
		if localctx != nil {
			t = localctx.(*ListElementIndexListContext)
		}
		return p.ListElementIndexList_Sempred(t, predIndex)

	case 15:
		var t *AssignInitListContext = nil
		if localctx != nil {
			t = localctx.(*AssignInitListContext)
		}
		return p.AssignInitList_Sempred(t, predIndex)

	case 28:
		var t *DefinitionListContext = nil
		if localctx != nil {
			t = localctx.(*DefinitionListContext)
		}
		return p.DefinitionList_Sempred(t, predIndex)

	case 31:
		var t *TypeSpecifierListContext = nil
		if localctx != nil {
			t = localctx.(*TypeSpecifierListContext)
		}
		return p.TypeSpecifierList_Sempred(t, predIndex)

	case 33:
		var t *ParaDeclaratorListContext = nil
		if localctx != nil {
			t = localctx.(*ParaDeclaratorListContext)
		}
		return p.ParaDeclaratorList_Sempred(t, predIndex)

	case 35:
		var t *ParaDeclaratorWithIdentityListContext = nil
		if localctx != nil {
			t = localctx.(*ParaDeclaratorWithIdentityListContext)
		}
		return p.ParaDeclaratorWithIdentityList_Sempred(t, predIndex)

	case 39:
		var t *FuncReturnParaListContext = nil
		if localctx != nil {
			t = localctx.(*FuncReturnParaListContext)
		}
		return p.FuncReturnParaList_Sempred(t, predIndex)

	case 42:
		var t *FuncStatementListContext = nil
		if localctx != nil {
			t = localctx.(*FuncStatementListContext)
		}
		return p.FuncStatementList_Sempred(t, predIndex)

	case 46:
		var t *FuncExecuteParaListContext = nil
		if localctx != nil {
			t = localctx.(*FuncExecuteParaListContext)
		}
		return p.FuncExecuteParaList_Sempred(t, predIndex)

	case 49:
		var t *PrintListContext = nil
		if localctx != nil {
			t = localctx.(*PrintListContext)
		}
		return p.PrintList_Sempred(t, predIndex)

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
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 3)

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

func (p *ZZParser) BExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
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

func (p *ZZParser) FuncExecuteParaList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 13:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ZZParser) PrintList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 14:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
