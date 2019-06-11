## gopomlicense
A utility for finding licenses of the dependencies used in your maven project.   
Transitive dependencies are ignored.  
It fetches license info from maven repository.
## Usage
      --format string    output format.
                                %n: new line
                                %i: artifact index(begin from 1)
                                %a: artifact name
                                (): license related format should be put in()
                                %b: license name
                                %c: license url
                                %d: artifact website (default "%i. %nArtifact Name: %a%nWebsite: %d%n(License: %b%nLicense Url: %c%n)----%n")
      --help             show help message
      --mvnCmd string    maven command location (default "mvn")
      --mvnUrl string    maven repository url for retrieving pom file (default "http://central.maven.org/maven2/")
      --pomFile string   pom file (absolute path)
## Sample output
With the default format string: %i. %nArtifact Name: %a%nWebsite: %d%n(License: %b%nLicense Url: %c%n)----%n
```
1. 
Artifact Name: MySQL Connector/J
Website: http://dev.mysql.com/doc/connector-j/en/
License: The GNU General Public License, Version 2
License Url: http://www.gnu.org/licenses/old-licenses/gpl-2.0.html
----
2. 
Artifact Name: spring-security-core
Website: http://spring.io/spring-security
License: The Apache Software License, Version 2.0
License Url: http://www.apache.org/licenses/LICENSE-2.0.txt
----
```
