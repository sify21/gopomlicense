package pom

import (
	"encoding/xml"
	"fmt"
)

type Project struct {
	XMLName          xml.Name  `xml:"project"`
	ParentGroupId    string    `xml:"parent>groupId"`
	ParentArtifactId string    `xml:"parent>artifactId"`
	ParentVersion    string    `xml:"parent>version"`
	GroupId          string    `xml:"groupId"`
	ArtifactId       string    `xml:"artifactId"`
	Version          string    `xml:"version"`
	Name             string    `xml:"name"`
	Url              string    `xml:"url"`
	Licenses         []License `xml:"licenses>license"`
}
type License struct {
	XMLName xml.Name `xml:"license"`
	Name    string   `xml:"name"`
	Url     string   `xml:"url"`
}

func (p Project) String() string {
	g := p.GroupId
	if g == "" {
		g = p.ParentGroupId
	}
	v := p.Version
	if v == "" {
		v = p.ParentVersion
	}
	return fmt.Sprintf("{%s %s %s}", g, p.ArtifactId, v)
}
