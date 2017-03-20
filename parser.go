package xrobotstag

import (
	"time"
	"net/http"
	"regexp"
	"fmt"
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
func parseHeaderTag(rawTag string, botname string, parsedTag *RobotsTag) {
	re := regexp.MustCompile(`^[a-z]+: (?:([a-z]+)(?:[, ]|$)+)+`)
	// Only return first hundred finds. More than ~10 would already be redundant -> 100 is crazy
	tags := re.FindAllStringSubmatch(rawTag,100)
	fmt.Println(tags)
	var tag string
	var counter int
	for _, tagArray := range tags {
		for counter, tag = range tagArray {
			fmt.Println(counter, tag)
			if counter == 0 {
				continue
			}
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

	}
}