package xrobotstag

import (
	"testing"
	"net/http"
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