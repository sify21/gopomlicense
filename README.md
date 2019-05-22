## gopomlicense
A utility for finding licenses of the dependencies used in your maven project.   
Transitive dependencies are ignored.  
It fetches license info from maven repository.
## Usage
      --help             show help message
      --mvnCmd string    maven command location (default "mvn")
      --mvnUrl string    maven repository url for retrieving pom file (default "http://central.maven.org/maven2/")
      --pomFile string   pom file absolute path
