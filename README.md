[![Build Status](https://drone.niels-ole.com/api/badges/nielsole/xrobotstag/status.svg)](https://drone.niels-ole.com/nielsole/xrobotstag)
# X-Robots-Tag

This is a parser for X-Robots tags both in the header and within the html.
Currently this is untested and has never run in production!

# Usage
## HTTP header

    robotsTag := RobotsTag{}
    RobotsTagFromHeaders(&resp.Header, "mybotname", &robotsTag)

## Meta Tag

    robotsTag := RobotsTag{}
    RobotsTagFromHtmlTag(htmlNode, "mybotname", robotsTag)
