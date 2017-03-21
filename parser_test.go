package xrobotstag

import (
	"testing"
	"net/http"
	"golang.org/x/net/html"
)

func TestHellowWorld(t *testing.T) {
	if true != true{
		t.Fail()
	}
}

func TestSimpleCase(t *testing.T) {
	header := http.Header{}
	header["X-Robots-Tag"] = []string{"googlebot: noindex, nofollow"}
	botname := "googlebot"
	parsedRobotsTag := RobotsTagFromHeaders(&header,botname)
	if parsedRobotsTag.Noindex == false || parsedRobotsTag.Nofollow == false {
		t.Error("Expected noindex and nofollow to be true")
	}
}

func TestParserHeaderTagSimple(t *testing.T) {
	parsedRobotsTag := RobotsTag{}
	parseHeaderTag("googlebot: noindex, nofollow","googlebot",&parsedRobotsTag)
	if parsedRobotsTag.Noindex == false || parsedRobotsTag.Nofollow == false {
		t.Error("Expected noindex and nofollow to be true")
	}
}

func TestParserHeaderTagDifferentBotname(t *testing.T) {
	parsedRobotsTag := RobotsTag{}
	parseHeaderTag("herokubot: noindex, nofollow","googlebot",&parsedRobotsTag)
	if parsedRobotsTag.Noindex == true || parsedRobotsTag.Nofollow == true {
		t.Error("Expected noindex and nofollow to be true")
	}
}

func TestParserHeaderDate(t *testing.T) {
	parsedRobotsTag := RobotsTag{}
	err := parseHeaderTag("unavailable_after: 25 Jun 2010 15:00:00 PST","googlebot",&parsedRobotsTag)
	if err != nil {
		t.Fatal(err)
	}
	if parsedRobotsTag.Noindex == true || parsedRobotsTag.UnavailableAfter == nil {
		t.Error("Expected UnavailableAfter to be set.")
	}
}

func TestParserHtmlTagSimple(t *testing.T) {
	parsedRobotsTag := RobotsTag{}
	err := RobotsTagFromHtmlTag(&html.Node{Data: "meta",Attr:[]html.Attribute{{Key:"name", Val:"robots"},{Key:"content", Val:"noindex, nofollow"}}}, "googlebot", &parsedRobotsTag)
	if err != nil {
		t.Fatal(err)
	}
	if parsedRobotsTag.Noindex == false || parsedRobotsTag.Nofollow == false {
		t.Error("Html tag was not correctly parsed. Expected noindex and nofollow to be set.")
	}
}