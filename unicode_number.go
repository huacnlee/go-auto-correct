package autocorrect

import "strings"

var (
	unicodeNumbers = map[string]string{
		// White Circled Number
		"①": "1", "②": "2", "③": "3", "④": "4", "⑤": "5", "⑥": "6", "⑦": "7", "⑧": "8", "⑨": "9", "⑩": "10", "⑪": "11", "⑫": "12", "⑬": "13", "⑭": "14", "⑮": "15", "⑯": "16", "⑰": "17", "⑱": "18", "⑲": "19", "⑳": "20",
		// Double-Circled Number
		"⓵": "1", "⓶": "2", "⓷": "3", "⓸": "4", "⓹": "5", "⓺": "6", "⓻": "7", "⓼": "8", "⓽": "9", "⓾": "10",
		// Black Circled Number
		"⓿": "0", "❶": "1", "❷": "2", "❸": "3", "❹": "4", "❺": "5", "❻": "6", "❼": "7", "❽": "8", "❾": "9", "❿": "10", "⓫": "11", "⓬": "12", "⓭": "13", "⓮": "14", "⓯": "15", "⓰": "16", "⓱": "17", "⓲": "18", "⓳": "19", "⓴": "20",
		// Black Circled Number with sans-serif
		"🄌": "0", "➊": "1", "➋": "2", "➌": "3", "➍": "4", "➎": "5", "➏": "6", "➐": "7", "➑": "8", "➒": "9", "➓": "10",
		// White Circled Letters
		"Ⓐ": "A", "Ⓑ": "B", "Ⓒ": "C", "Ⓓ": "D", "Ⓔ": "E", "Ⓕ": "F", "Ⓖ": "G", "Ⓗ": "H", "Ⓘ": "I", "Ⓙ": "J", "Ⓚ": "K", "Ⓛ": "L", "Ⓜ": "M", "Ⓝ": "N", "Ⓞ": "O", "Ⓟ": "P", "Ⓠ": "Q", "Ⓡ": "R", "Ⓢ": "S", "Ⓣ": "T", "Ⓤ": "U", "Ⓥ": "V", "Ⓦ": "W", "Ⓧ": "X", "Ⓨ": "Y", "Ⓩ": "Z",
		"ⓐ": "a", "ⓑ": "b", "ⓒ": "c", "ⓓ": "d", "ⓔ": "e", "ⓕ": "f", "ⓖ": "g", "ⓗ": "h", "ⓘ": "i", "ⓙ": "j", "ⓚ": "k", "ⓛ": "l", "ⓜ": "m", "ⓝ": "n", "ⓞ": "o", "ⓟ": "p", "ⓠ": "q", "ⓡ": "r", "ⓢ": "s", "ⓣ": "t", "ⓤ": "u", "ⓥ": "v", "ⓦ": "w", "ⓧ": "x", "ⓨ": "y", "ⓩ": "z",
		// Black Circled Letters
		"🅐": "A", "🅑": "B", "🅒": "C", "🅓": "D", "🅔": "E", "🅕": "F", "🅖": "G", "🅗": "H", "🅘": "I", "🅙": "J", "🅚": "K", "🅛": "L", "🅜": "M", "🅝": "N", "🅞": "O", "🅟": "P", "🅠": "Q", "🅡": "R", "🅢": "S", "🅣": "T", "🅤": "U", "🅥": "V", "🅦": "W", "🅧": "X", "🅨": "Y", "🅩": "Z",
		// White Squared Letters
		"🄰": "A", "🄱": "B", "🄲": "C", "🄳": "D", "🄴": "E", "🄵": "F", "🄶": "G", "🄷": "H", "🄸": "I", "🄹": "J", "🄺": "K", "🄻": "L", "🄼": "M", "🄽": "N", "🄾": "O", "🄿": "P", "🅀": "Q", "🅁": "R", "🅂": "S", "🅃": "T", "🅄": "U", "🅅": "V", "🅆": "W", "🅇": "X", "🅈": "Y", "🅉": "Z",
		// Black Squered Letters
		"🅰": "A", "🅱": "B", "🅲": "C", "🅳": "D", "🅴": "E", "🅵": "F", "🅶": "G", "🅷": "H", "🅸": "I", "🅹": "J", "🅺": "K", "🅻": "L", "🅼": "M", "🅽": "N", "🅾": "O", "🅿": "P", "🆀": "Q", "🆁": "R", "🆂": "S", "🆃": "T", "🆄": "U", "🆅": "V", "🆆": "W", "🆇": "X", "🆈": "Y", "🆉": "Z",
	}
)

// WithUnicodeNumber normalize unicode numbers
//
// ① ② ... ⑳, ⓵ ⓶, ❶ ❷ -> 1 2 ... 20
//
// Ⓐ Ⓑ ... Ⓩ, 🅐 🅑, 🄰 🄱, 🅰 🅱 -> A B ... Z
//
// ⓐ ... ⓩ -> a ... z
//
// Reference:
//
// http://xahlee.info/comp/unicode_circled_numbers.html
//
// https://www.unicode.org/charts/nameslist/n_2460.html
//
func WithUnicodeNumber() Option {
	return unicodeNumber{}
}

type unicodeNumber struct{}

func (un unicodeNumber) Format(text string) string {
	strs := strings.Split(text, "")
	for i, str := range strs {
		newChar := unicodeNumbers[str]
		if newChar != "" {
			strs[i] = newChar
			continue
		}
	}

	return strings.Join(strs, "")
}
