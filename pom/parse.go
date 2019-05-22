package pom

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
)

type Artifact struct {
	GroupId    string
	ArtifactId string
	Version    string
}

func ListDep(mvnCmd, pomFile string) []Artifact {
	cmd := exec.Command(mvnCmd, "dependency:list", "-DexcludeTransitive")
	cmd.Dir = filepath.Dir(pomFile)
	if out, err := cmd.Output(); err != nil {
		log.Fatalf("run mvn command error. %v", err)
		return nil
	} else {
		reader := bufio.NewReader(bytes.NewBuffer(out))
		var line string
		am := make(map[Artifact]struct{})
		for {
			line, err = reader.ReadString('\n')
			if err != nil {
				break
			}
			if strings.HasPrefix(line, "[INFO]    ") && strings.Contains(line, ":jar:") {
				a := strings.Split(line[10:], ":")
				ar := Artifact{
					GroupId:    a[0],
					ArtifactId: a[1],
					Version:    a[len(a)-2],
				}
				am[ar] = struct{}{}
			}
		}
		if err != io.EOF {
			log.Fatalf("error read mvn output. %v", err)
		}
		ret := make([]Artifact, 0, len(am))
		for k := range am {
			ret = append(ret, k)
		}
		return ret
	}
}
