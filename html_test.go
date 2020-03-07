package autocorrect

import (
	"regexp"
	"testing"
)

var (
	htmlSpaceRe = regexp.MustCompile(`>[\s]+<`)
)

func assertHTMLEqual(t *testing.T, exptected, actual string) {
	if htmlSpaceRe.ReplaceAllString(exptected, "><") != htmlSpaceRe.ReplaceAllString(actual, "><") {
		t.Errorf("\nexptected:\n%s\nactual   :\n%s", exptected, actual)
	}
}

func BenchmarkFormatHTML(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// about 1.4ms/op
		FormatHTML(readFile("example.txt"))
	}
}
func TestFormatHTMLWithFixtuires(t *testing.T) {
	expected := readFile("example.txt.expected")
	out, err := FormatHTML(readFile("example.txt"))
	if err != nil {
		t.Error(err)
	}
	assertHTMLEqual(t, expected, out)
}

func TestFormatHTMLWithEscapedHTML(t *testing.T) {
	html := `<p>据2019年12月27日，三生制药获JP Morgan Chase &amp; Co.每股均价9.582港元，增持270.3万股</p>`
	expected := `<p>据2019年12月27日，三生制药获 JP Morgan Chase &amp; Co.每股均价 9.582 港元，增持 270.3 万股</p>`
	out, err := FormatHTML(html)
	if err != nil {
		t.Error(err)
	}
	assertHTMLEqual(t, expected, out)

	html = `<p>据2019年12月27日，三生制药获JP Morgan Chase & Co.每股均价9.582港元，增持270.3万股</p>`
	expected = `<p>据2019年12月27日，三生制药获 JP Morgan Chase &amp; Co.每股均价 9.582 港元，增持 270.3 万股</p>`
	out, err = FormatHTML(html)
	if err != nil {
		t.Error(err)
	}
	assertHTMLEqual(t, expected, out)
}
