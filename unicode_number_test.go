package autocorrect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertUnicodeNumber(t *testing.T, expected string, text string) {
	assert.Equal(t, expected, Format(text, WithUnicodeNumber()))
}

func Test_WithUnicodeNumber(t *testing.T) {

	// Reference:
	// http://xahlee.info/comp/unicode_circled_numbers.html
	// https://www.unicode.org/charts/nameslist/n_2460.html

	// White Circled Number
	assertUnicodeNumber(t, "1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20", "① ② ③ ④ ⑤ ⑥ ⑦ ⑧ ⑨ ⑩ ⑪ ⑫ ⑬ ⑭ ⑮ ⑯ ⑰ ⑱ ⑲ ⑳")

	// Double-Circled Number
	assertUnicodeNumber(t, "1 2 3 4 5 6 7 8 9 10", "⓵ ⓶ ⓷ ⓸ ⓹ ⓺ ⓻ ⓼ ⓽ ⓾")

	// Black Circled Number
	assertUnicodeNumber(t, "0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20", "⓿ ❶ ❷ ❸ ❹ ❺ ❻ ❼ ❽ ❾ ❿ ⓫ ⓬ ⓭ ⓮ ⓯ ⓰ ⓱ ⓲ ⓳ ⓴")

	// Black Circled Number with Sans-serif
	assertUnicodeNumber(t, "0 1 2 3 4 5 6 7 8 9 10", "🄌 ➊ ➋ ➌ ➍ ➎ ➏ ➐ ➑ ➒ ➓")

	// Double-Circled Number
	assertUnicodeNumber(t, "1 2 3 4 5 6 7 8 9 10", "⓵ ⓶ ⓷ ⓸ ⓹ ⓺ ⓻ ⓼ ⓽ ⓾")

	// White Circled Letters
	assertUnicodeNumber(t, "A B C D E F G H I J K L M N O P Q R S T U V W X Y Z", "Ⓐ Ⓑ Ⓒ Ⓓ Ⓔ Ⓕ Ⓖ Ⓗ Ⓘ Ⓙ Ⓚ Ⓛ Ⓜ Ⓝ Ⓞ Ⓟ Ⓠ Ⓡ Ⓢ Ⓣ Ⓤ Ⓥ Ⓦ Ⓧ Ⓨ Ⓩ")
	assertUnicodeNumber(t, "a b c d e f g h i j k l m n o p q r s t u v w x y z", "ⓐ ⓑ ⓒ ⓓ ⓔ ⓕ ⓖ ⓗ ⓘ ⓙ ⓚ ⓛ ⓜ ⓝ ⓞ ⓟ ⓠ ⓡ ⓢ ⓣ ⓤ ⓥ ⓦ ⓧ ⓨ ⓩ")

	// Black Circled Letters
	assertUnicodeNumber(t, "A B C D E F G H I J K L M N O P Q R S T U V W X Y Z", "🅐 🅑 🅒 🅓 🅔 🅕 🅖 🅗 🅘 🅙 🅚 🅛 🅜 🅝 🅞 🅟 🅠 🅡 🅢 🅣 🅤 🅥 🅦 🅧 🅨 🅩")

	// Squared Letters
	assertUnicodeNumber(t, "A B C D E F G H I J K L M N O P Q R S T U V W X Y Z", "🄰 🄱 🄲 🄳 🄴 🄵 🄶 🄷 🄸 🄹 🄺 🄻 🄼 🄽 🄾 🄿 🅀 🅁 🅂 🅃 🅄 🅅 🅆 🅇 🅈 🅉")
	// Black Squared Letters
	assertUnicodeNumber(t, "A B C D E F G H I J K L M N O P Q R S T U V W X Y Z", "🅰 🅱 🅲 🅳 🅴 🅵 🅶 🅷 🅸 🅹 🅺 🅻 🅼 🅽 🅾 🅿 🆀 🆁 🆂 🆃 🆄 🆅 🆆 🆇 🆈 🆉")
}
