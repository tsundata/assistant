package rule

import (
	"github.com/PuerkitoBio/goquery"
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
	// text
	f := ParseFun(sel, `$("a.storylink").text`)
	r, err := f.Invoke()
	if err != nil {
		t.Fatal(err)
	}
	if r != "demo" {
		t.Fatal("error ParseFun")
	}
	// expand
	f = ParseFun(sel, `$(".rank").text.expand("(\d+)", "#$1")`)
	r, err = f.Invoke()
	if err != nil {
		t.Fatal(err)
	}
	if r != "#3" {
		t.Fatal("error ParseFun")
	}
	// match
	f = ParseFun(sel, `$(".rank").text.match("(\d+)")`)
	r, err = f.Invoke()
	if err != nil {
		t.Fatal(err)
	}
	if r != "3" {
		t.Fatal("error ParseFun")
	}
}
