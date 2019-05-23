package pom

import (
	"encoding/xml"
	"fmt"
)

type Project struct {
	XMLName    xml.Name  `xml:"project"`
	GroupId    string    `xml:"groupId"`
	ArtifactId string    `xml:"artifactId"`
	Version    string    `xml:"version"`
	Name       string    `xml:"name"`
	Url        string    `xml:"url"`
	Licenses   []License `xml:"licenses>license"`
}
type License struct {
	XMLName xml.Name `xml:"license"`
	Name    string   `xml:"name"`
	Url     string   `xml:"url"`
}

func (p Project) String() string {
	return fmt.Sprintf("{%s %s %s}", p.GroupId, p.ArtifactId, p.Version)
}
