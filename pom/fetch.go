package pom

import (
	"log"
	"net/http"
	"strings"
)

func Fetch(mvnUrl string, artifact Artifact, converter ConverterFunc) (interface{}, error) {
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
