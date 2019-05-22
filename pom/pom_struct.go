package pom

import "encoding/xml"

type Project struct {
	XMLName  xml.Name  `xml:"project"`
	Url      string    `xml:"url"`
	Licenses []License `xml:"licenses>license"`
}
type License struct {
	XMLName xml.Name `xml:"license"`
	Name    string   `xml:"name"`
	Url     string   `xml:"url"`
}
