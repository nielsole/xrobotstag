package xrobotstag

import (
	"time"
	"net/http"
	"strings"
	"errors"
	"golang.org/x/net/html"
)

type RobotsTag struct {
	All bool
	Noindex bool
	Nofollow bool
	// None is equivalent to noindex and nofollow. Omitted
	Noarchive bool
	Nosnippet bool
	Noodp bool
	Notranslate bool
	Noimageindex bool
	UnavailableAfter *time.Time
}

func RobotsTagFromHeaders(header *http.Header, botname string) *RobotsTag {
	var tags []string = (*header)["X-Robots-Tag"]
	robotsTag := &RobotsTag{}
	for _, tag := range tags {
		parseHeaderTag(tag, botname, robotsTag)
	}
	return robotsTag
}

func RobotsTagFromHtmlTag(token *html.Node, botname string, parsedTag *RobotsTag) error {
	if (token.Data != "meta") {
		return errors.New("Node is not a meta-node")
	}
	var attributes string
	var affectedBot string
	for _, attr := range token.Attr {
		if attr.Key == "content" {
			attributes = attr.Val
		}
		if attr.Key == "name" {
			affectedBot = attr.Val
		}
	}
	if (attributes == ""){
		return errors.New("No content flag there")
	}
	// Check whether we are meant by this flag
	if (affectedBot != botname && affectedBot != "robots") {
		return nil
	}
	parseHtmlTag(attributes, parsedTag)
	return nil
}

func parseHtmlTag(rawString string, parsedTag *RobotsTag) error {
	var botnameAndBody []string = strings.SplitN(rawString, ":", 2)
	if len(botnameAndBody) > 1 && botnameAndBody[0] == "unavailable_after" {
		unavAfter, err := parseUnavailableAfter(botnameAndBody[1])
		if err != nil{
			return errors.New("Could not parse datetime")
		}
		parsedTag.UnavailableAfter = unavAfter
		return nil
	}else {
		parseTags(rawString,parsedTag)
	}
	return nil
}
func parseHeaderTag(rawTag string, botname string, parsedTag *RobotsTag) error {
	var botnameAndBody []string = strings.SplitN(rawTag, ":", 2)
	// Position of the tags if no botname is present.
	var i int = 0
	if len(botnameAndBody) > 1 {
		if botnameAndBody[0] == "unavailable_after" {
			unavAfter, err := parseUnavailableAfter(botnameAndBody[1])
			if err != nil{
				return errors.New("Could not parse datetime")
			}
			parsedTag.UnavailableAfter = unavAfter
		}
		//offset the tags by one
		i = 1
		// check whether this bot is affected
		if botnameAndBody[0] != botname {
			// We are not meant, which is not an error
			return nil
		}
	}
	parseTags(botnameAndBody[i], parsedTag)
	return nil
}


func parseUnavailableAfter(rawTime string) (*time.Time, error) {
	trimmedTime := strings.Trim(rawTime, " ")
	println(trimmedTime)
	/*
	Google says they use RFC850 but they show an example where they have a different format.
	Since we will eventually try a lot of different date time formats depending on what is widespread on the internet,
	we will try one after the other and log if we could not make sense of it.
	 */
	var timeFormats []string = []string{
		time.RFC850,
		"02 Jan 2006 15:04:05 MST", // Example: 25 Jun 2010 15:00:00 PST
		"_2 Jan 2006 15:04:05 MST", // Example: 25 Jun 2010 15:00:00 PST
		}
	for _, timeFormat := range timeFormats {
		unavAfter, err := time.Parse(timeFormat, trimmedTime)
		if err == nil {
			return &unavAfter, nil
		}
	}
	return nil, errors.New("Could not parse time. Unknown format: " + trimmedTime)
}

func parseTags (rawTags string, parsedTag *RobotsTag) []string {
	var tags []string = strings.Split(rawTags, ",")
	for _, tag := range tags {
		tag := strings.Trim(tag, " ")
		switch tag {
		case "noindex":
			parsedTag.Noindex = true
		case "nofollow":
			parsedTag.Nofollow = true
		case "nosnippet":
			parsedTag.Nosnippet = true
		case "noarchive":
			parsedTag.Noarchive = true
		case "noodp":
			parsedTag.Noodp = true
		}
	}
	return tags
}