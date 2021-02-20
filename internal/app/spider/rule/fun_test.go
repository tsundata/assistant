package rule

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestParseFun(t *testing.T) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(`<tr class="athing" id="25782861">
      <td align="right" valign="top" class="title"><span class="rank">3.</span></td>      
		<td valign="top" class="votelinks"><center><a id="up_25782861" href="vote?id=25782861&amp;how=up&amp;goto=news">
		<div class="votearrow" title="upvote"></div></a></center></td><td class="title">
		<a href="http://demo.com" class="storylink">demo</a>
		<span class="sitebit comhead"> (<a href="from?site=demo.com">
		<span class="sitestr">demo.com</span></a>)
		</span></td></tr>`))
	if err != nil {
		t.Fatal(err)
	}

	sel := doc.First()
	t.Parallel()
	tests := []struct {
		name   string
		fun    string
		expect string
	}{
		{
			"text",
			`$("a.storylink").text`,
			"demo",
		},
		{
			"expand",
			`$(".rank").text.expand("(\d+)", "#$1")`,
			"#3",
		},
		{
			"match",
			`$(".rank").text.match("(\d+)")`,
			"3",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := ParseFun(sel, tt.fun)
			r, err := f.Invoke()
			if err != nil {
				t.Fatal(err)
			}
			require.Equal(t, r, tt.expect)
		})
	}
}
