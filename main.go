package main

import (
	"encoding/xml"
	"github.com/sify21/gopomlicense/config"
	"github.com/sify21/gopomlicense/pom"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func main() {
	pflag.String(config.MAVEN_URL, "http://central.maven.org/maven2/", "maven repository url for retrieving pom file")
	pflag.String(config.POM_FILE, "", "pom file absolute path")
	pflag.Bool("help", false, "show help message")
	pflag.String(config.MVN_CMD, "mvn", "maven command location")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	if viper.GetBool("help") || viper.GetString(config.POM_FILE) == "" {
		pflag.PrintDefaults()
	} else {
		//pomFile, err := os.Open(viper.GetString("pomFile"))
		//if err != nil {
		//	log.Fatalf("error open pom file. %v", err)
		//}
		//defer pomFile.Close()
		//var project pom.Project
		//if err = xml.NewDecoder(pomFile).Decode(&project); err != nil {
		//	log.Fatalf("error decode pom file. %v", err)
		//}
		//log.Printf("%+v", project)
		artifacts := pom.ListDep(viper.GetString(config.MVN_CMD), viper.GetString(config.POM_FILE))
		var errArtifacts []pom.Artifact
		ch := make(chan interface{}, len(artifacts))
		for _, v := range artifacts {
			a := v
			go func() {
				result, err := pom.Fetch(viper.GetString(config.MAVEN_URL), a, func(response *http.Response) (interface{}, error) {
					var project pom.Project
					if err := xml.NewDecoder(response.Body).Decode(&project); err != nil {
						return nil, err
					}
					return project, nil
				})
				if err != nil {
					log.Printf("fetch pom error. %+v %v", a, err)
					ch <- a
				} else if project, ok := result.(pom.Project); ok {
					ch <- project
				} else {
					log.Printf("fetch pom error. %+v %s", a, "can't convert to project")
					ch <- a
				}
			}()
		}
		for i := 0; i < len(artifacts); i++ {
			ret := <-ch
			switch res := ret.(type) {
			case pom.Artifact:
				errArtifacts = append(errArtifacts, res)
			case pom.Project:
				log.Printf("license info %+v", res)
			}
		}
		log.Printf("Finished. Total: %d, Success: %d, Fail: %d", len(artifacts), len(artifacts)-len(errArtifacts), len(errArtifacts))
		log.Println("Failed artifact: ")
		for _, v := range errArtifacts {
			log.Printf("%v", v)
		}
	}
}
