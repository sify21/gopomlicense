package main

import (
	"fmt"
	"github.com/sify21/gopomlicense/config"
	"github.com/sify21/gopomlicense/pom"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	pflag.String(config.MAVEN_URL, "http://central.maven.org/maven2/", "maven repository url for retrieving pom file")
	pflag.String(config.POM_FILE, "", "pom file (absolute path)")
	pflag.Bool("help", false, "show help message")
	pflag.String(config.MVN_CMD, "mvn", "maven command location")
	pflag.String(config.FORMAT, "%i. %nArtifact Name: %a%nWebsite: %d%n(License: %b%nLicense Url: %c%n)----%n", "output format.\n\t%n: new line\n\t%i: artifact index(begin from 1)\n\t%a: artifact name\n\t(): license related format should be put in()\n\t%b: license name\n\t%c: license url\n\t%d: artifact website")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	if viper.GetBool("help") || viper.GetString(config.POM_FILE) == "" {
		pflag.PrintDefaults()
	} else {
		artifacts := pom.ListDep(viper.GetString(config.MVN_CMD), viper.GetString(config.POM_FILE))
		var errArtifacts []pom.Artifact
		var projectsWithLicense []pom.Project
		var projectsWithoutLicense []pom.Project
		ch := make(chan interface{}, len(artifacts))
		for _, v := range artifacts {
			a := v
			go func() {
				project, err := pom.ResolveToProject(viper.GetString(config.MAVEN_URL), a)
				if err != nil {
					log.Printf("Can't resolve to project. %s %v", a.String(), err)
					ch <- a
					return
				}
				projectWithLicense, err := pom.FetchMavenLicense(viper.GetString(config.MAVEN_URL), project)
				if err != nil {
					log.Printf("Fetch license error. %s %v", a.String(), err)
					ch <- a
					return
				}
				if projectWithLicense == nil {
					ch <- *project
					return
				}
				if project.Url == "" {
					project.Url = projectWithLicense.Url
				}
				project.Licenses = projectWithLicense.Licenses
				ch <- *project
			}()
		}
		for i := 0; i < len(artifacts); i++ {
			ret := <-ch
			switch res := ret.(type) {
			case pom.Artifact:
				errArtifacts = append(errArtifacts, res)
			case pom.Project:
				if len(res.Licenses) > 0 {
					projectsWithLicense = append(projectsWithLicense, res)
				} else {
					projectsWithoutLicense = append(projectsWithoutLicense, res)
				}
			}
		}
		log.Printf("Finished. Total: %d, Success: %d", len(artifacts), len(projectsWithLicense))
		log.Printf("There are %d artifacts that failed to parse: \n", len(errArtifacts))
		for _, v := range errArtifacts {
			fmt.Println("\t" + v.String())
		}
		log.Printf("There are %d artifacts that don't have license info: \n", len(projectsWithoutLicense))
		for _, v := range projectsWithoutLicense {
			fmt.Println("\t" + v.String())
		}
		log.Println("Formatted output: ")
		format := viper.GetString(config.FORMAT)
		reg, _ := regexp.Compile(`(.*)\((.*)\)(.*)`)
		ff := reg.FindStringSubmatch(format)
		artifactFormatBefore := strings.ReplaceAll(ff[1], "%n", "\n")
		licenseFormat := strings.ReplaceAll(ff[2], "%n", "\n")
		artifactFormatAfter := strings.ReplaceAll(ff[3], "%n", "\n")
		for k, v := range projectsWithLicense {
			before := strings.ReplaceAll(artifactFormatBefore, "%i", strconv.FormatInt(int64(k+1), 10))
			after := strings.ReplaceAll(artifactFormatAfter, "%i", strconv.FormatInt(int64(k+1), 10))
			before = strings.ReplaceAll(before, "%a", v.Name)
			after = strings.ReplaceAll(after, "%a", v.Name)
			before = strings.ReplaceAll(before, "%d", v.Url)
			after = strings.ReplaceAll(after, "%d", v.Url)
			licStr := ""
			for _, l := range v.Licenses {
				lic := strings.ReplaceAll(licenseFormat, "%b", l.Name)
				lic = strings.ReplaceAll(lic, "%c", l.Url)
				licStr += lic
			}
			fmt.Print(before + licStr + after)
		}
	}
}
