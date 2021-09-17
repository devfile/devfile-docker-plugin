package javaHelpers

type MavenProject struct {
	GroupId    string `xml:"groupId"`
	ArtifactId string `xml:"artifactId"`
	Version    string `xml:"version"`
	Properties struct {
		JavaVersion string `xml:"java.version"`
	} `xml:"properties"`
	Parent struct {
		GroupId    string `xml:"groupId"`
		ArtifactId string `xml:"artifactId"`
		Version    string `xml:"version"`
	} `xml:"parent"`
	Name         string `xml:"name"`
	Dependencies struct {
		Dependency []struct {
			GroupId    string `xml:"groupId"`
			ArtifactId string `xml:"artifactId"`
			Scope      string `xml:"scope"`
			Exclusions struct {
				Exclusion struct {
					Text       string `xml:",chardata"`
					GroupId    string `xml:"groupId"`
					ArtifactId string `xml:"artifactId"`
				} `xml:"exclusion"`
			} `xml:"exclusions"`
			Version  string `xml:"version"`
			Optional string `xml:"optional"`
		} `xml:"dependency"`
	} `xml:"dependencies"`
	Modules struct {
		Text   string   `xml:",chardata"`
		Module []string `xml:"module"`
	} `xml:"modules"`
}
