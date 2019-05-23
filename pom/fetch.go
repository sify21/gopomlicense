package pom

import (
	"encoding/xml"
	"fmt"
	"golang.org/x/net/html/charset"
	"log"
	"net/http"
	"strings"
)

func FetchPom(mvnUrl string, artifact Artifact, converter ConverterFunc) (interface{}, error) {
	orgPath := strings.ReplaceAll(artifact.GroupId, ".", "/")
	pomFile := artifact.ArtifactId + "-" + artifact.Version + ".pom"
	if !strings.HasSuffix(mvnUrl, "/") {
		mvnUrl = mvnUrl + "/"
	}
	pomUrl := mvnUrl + orgPath + "/" + artifact.ArtifactId + "/" + artifact.Version + "/" + pomFile
	log.Printf("fetching pom %s", pomUrl)
	resp, err := http.Get(pomUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return converter(resp)
}

type ConverterFunc func(response *http.Response) (interface{}, error)

func ResolveToProject(mvnUrl string, artifact Artifact) (*Project, error) {
	result, err := FetchPom(mvnUrl, artifact, func(response *http.Response) (interface{}, error) {
		var project Project
		decoder := xml.NewDecoder(response.Body)
		decoder.CharsetReader = charset.NewReaderLabel
		if err := decoder.Decode(&project); err != nil {
			return nil, err
		}
		return project, nil
	})
	if err != nil {
		return nil, err
	}
	if p, ok := result.(Project); ok {
		return &p, nil
	} else {
		return nil, fmt.Errorf("%s result can't convert to project", artifact.String())
	}
}

func FetchMavenLicense(mvnUrl string, project *Project) (*Project, error) {
	if len(project.Licenses) > 0 {
		return project, nil
	} else if project.ParentGroupId != "" {
		parent := Artifact{
			GroupId:    project.ParentGroupId,
			ArtifactId: project.ParentArtifactId,
			Version:    project.ParentVersion,
		}
		p, err := ResolveToProject(mvnUrl, parent)
		if err != nil {
			return nil, err
		}
		return FetchMavenLicense(mvnUrl, p)
	} else {
		return nil, nil
	}
}
