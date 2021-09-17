module github.com/devfile/devrunner

go 1.16

require (
	github.com/compose-spec/compose-go v0.0.0-20210901090333-feb401cda7f7
	github.com/containerd/console v1.0.2
	github.com/devfile/api/v2 v2.0.0-20210420202853-ff3c01bf8292
	github.com/devfile/library v1.0.0-alpha.3
	github.com/docker/cli v20.10.7+incompatible
	github.com/docker/compose/v2 v2.0.0-rc.3
	github.com/fatih/color v1.9.0 // indirect
	github.com/go-git/go-git/v5 v5.4.2
	github.com/magiconair/properties v1.8.5
	github.com/onsi/ginkgo v1.14.2 // indirect
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.2.1
	golang.org/x/mod v0.5.0
	k8s.io/apiextensions-apiserver v0.20.1 // indirect
	sigs.k8s.io/yaml v1.2.0
)

replace (
	// protobuf: corresponds to containerd
	github.com/golang/protobuf => github.com/golang/protobuf v1.5.2
	github.com/hashicorp/go-immutable-radix => github.com/tonistiigi/go-immutable-radix v0.0.0-20170803185627-826af9ccf0fe
	github.com/jaguilar/vt100 => github.com/tonistiigi/vt100 v0.0.0-20190402012908-ad4c4a574305
	// genproto: corresponds to containerd
	google.golang.org/genproto => google.golang.org/genproto v0.0.0-20200224152610-e50cd9704f63
	// grpc: corresponds to protobuf
	google.golang.org/grpc => google.golang.org/grpc v1.30.0
)
