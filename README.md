## gopomlicense
A util for finding dependency licenses of your maven project.   
It fetches license info from maven repository.
## Usage
      --format string    output format.
                                %n: new line
                                %i: artifact index(begin from 1)
                                %a: artifact name
                                (): license related format should be put in()
                                %b: license name
                                %c: license url
                                 (default "%i. %nArtifact Name: %a%n(License: %b%nLicense Url: %c%n)----%n")
      --help             show help message
      --mvnCmd string    maven command location (default "mvn")
      --mvnUrl string    maven repository url for retrieving pom file (default "http://central.maven.org/maven2/")
      --pomFile string   pom file (absolute path)
