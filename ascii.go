/*
* @Author: scottxiong
* @Date:   2019-07-31 16:40:43
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-07-31 17:59:53
 */

package main

import (
	"fmt"
	"os"
)

func main() {
	printAscii()
	printTable()
}

func printAscii() {
	fmt.Fprintf(os.Stdin, `
ASCII Table
Dec  = Decimal Value
Char = Character

'5' has the int value 53
if we write '5'-'0' it evaluates to 53-48, or the int 5
if we write char c = 'B'+32; then c stores 'b'
`)
}

type ascii struct {
	dec  int
	char string
	note string
}

type asciis []ascii

var asciitable1, asciitable2 asciis

func init() {
	asciitable1 = asciis{
		ascii{0, "NUL", "(null)"},
		ascii{1, "SOH", "(start of heading)"},
		ascii{2, "STX", "(start of text)"},
		ascii{3, "ETX", "(end of text)"},
		ascii{4, "EOT", "(end of transmission)"},
		ascii{5, "ENQ", "(enquiry)"},
		ascii{6, "ACK", "(acknowledge)"},
		ascii{7, "BEL", "(bell)"},
		ascii{8, "BS", "(backspace)"},
		ascii{9, "TAB", "(horizontal tab)"},
		ascii{10, "LF", "(NL line feed, new line)"},
		ascii{11, "VT", "(vertical tab)"},
		ascii{12, "FF", "(NP form feed, new page)"},
		ascii{13, "CR", "(carriage return)"},
		ascii{14, "SO", "(shift out)"},
		ascii{15, "SI", "(shift in)"},
		ascii{16, "DLE", "(data link escape)"},
		ascii{17, "DC1", "(device control 1)"},
		ascii{18, "DC2", "(device control 2)"},
		ascii{19, "DC3", "(device control 3)"},
		ascii{20, "DC4", "(device control 4)"},
		ascii{21, "NAK", "(negative acknowledge)"},
		ascii{22, "SYN", "(synchronous idle)"},
		ascii{23, "ETB", "(end of trans. block)"},
		ascii{24, "CAN", "(cancel)"},
		ascii{25, "EM", "(end of medium)"},
		ascii{26, "SUB", "(substitute)"},
		ascii{27, "ESC", "(escape)"},
		ascii{28, "FS", "(file separator)"},
		ascii{29, "GS", "(group separator)"},
		ascii{30, "RS", "(record separator)"},
		ascii{31, "US", "(unit separator)"},
	}
	asciitable2 = asciis{
		ascii{32, "SPACE", ""},
		ascii{33, "!", ""},
		ascii{34, `"`, ""},
		ascii{35, `#`, ""},
		ascii{36, `$`, ""},
		ascii{37, `%`, ""},
		ascii{38, "&", ""},
		ascii{39, `'`, ""},
		ascii{40, `(`, ""},
		ascii{41, `)`, ""},
		ascii{42, `*`, ""},
		ascii{43, "+", ""},
		ascii{44, `,`, ""},
		ascii{45, `-`, ""},
		ascii{46, `.`, ""},
		ascii{47, `/`, ""},
		ascii{48, "0", ""},
		ascii{49, "1", ""},
		ascii{50, "2", ""},
		ascii{51, "3", ""},
		ascii{52, "4", ""},
		ascii{53, "5", ""},
		ascii{54, "6", ""},
		ascii{55, "7", ""},
		ascii{56, "8", ""},
		ascii{57, "9", ""},
		ascii{58, `:`, ""},
		ascii{59, ";", ""},
		ascii{60, "<", ""},
		ascii{61, "=", ""},
		ascii{62, ">", ""},
		ascii{63, "?", ""},
	}

	asciitable3 = asciis{
		ascii{64, "@", ""},
		ascii{65, "A", ""},
		ascii{66, "B", ""},
		ascii{67, "C", ""},
		ascii{68, "D", ""},
		ascii{69, "E", ""},
		ascii{70, "F", ""},
		ascii{71, "G", ""},
		ascii{72, "H", ""},
		ascii{73, "I", ""},
		ascii{74, "J", ""},
		ascii{75, "K", ""},
		ascii{76, "L", ""},
		ascii{77, "M", ""},
		ascii{78, "N", ""},
		ascii{79, "O", ""},
		ascii{80, "P", ""},
		ascii{81, "Q", ""},
		ascii{82, "R", ""},
		ascii{83, "S", ""},
		ascii{84, "T", ""},
		ascii{85, "U", ""},
		ascii{86, "V", ""},
		ascii{87, "W", ""},
		ascii{88, "X", ""},
		ascii{89, "Y", ""},
		ascii{90, "Z", ""},
		ascii{91, "[", ""},
		ascii{92, `\`, ""},
		ascii{93, "]", ""},
		ascii{94, "^", ""},
		ascii{95, "_", ""},
	}

	asciitable4 = asciis{
		ascii{96, "`", ""},
		ascii{97, "a", ""},
		ascii{98, "b", ""},
		ascii{99, "c", ""},
		ascii{100, "d", ""},
		ascii{101, "e", ""},
		ascii{102, "f", ""},
		ascii{103, "g", ""},
		ascii{104, "h", ""},
		ascii{105, "i", ""},
		ascii{106, "j", ""},
		ascii{107, "k", ""},
		ascii{108, "l", ""},
		ascii{109, "m", ""},
		ascii{110, "n", ""},
		ascii{111, "o", ""},
		ascii{112, "p", ""},
		ascii{113, "q", ""},
		ascii{114, "r", ""},
		ascii{115, "s", ""},
		ascii{116, "t", ""},
		ascii{117, "u", ""},
		ascii{118, "v", ""},
		ascii{119, "w", ""},
		ascii{120, "x", ""},
		ascii{121, "y", ""},
		ascii{122, "z", ""},
		ascii{123, "{", ""},
		ascii{124, "|", ""},
		ascii{125, "}", ""},
		ascii{126, "~", ""},
		ascii{127, "DEL", ""},
	}
}

func printTable() {
	fmt.Println("Dec  Char          Dec    Char")
	fmt.Println("---------")
	fmt.Println(asciitable2[0])

}
