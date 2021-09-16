package devfile

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/compose-spec/compose-go/types"
	"github.com/devfile/api/v2/pkg/apis/workspaces/v1alpha2"
	"github.com/devfile/api/v2/pkg/devfile"
	. "github.com/devfile/devrunner/tests"
	"github.com/devfile/library/pkg/devfile/parser"
	v2 "github.com/devfile/library/pkg/devfile/parser/data/v2"
)

func TestToComposeProject(t *testing.T) {
	type args struct {
		devfile parser.DevfileObj
	}
	type testCase struct {
		name    string
		args    args
		want    types.Project
		wantErr bool
	}
	tests := []testCase{
		func() testCase {
			return testCase{
				name: "converts fields",
				args: args{
					devfile: sampleDevfile("test", 3, 3),
				},
				want: sampleProject("test", 3, 3),
			}
		}(),
		func() testCase {
			return testCase{
				name: "no services",
				args: args{
					devfile: sampleDevfile("test", 0, 1),
				},
				want: sampleProject("test", 0, 1),
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToComposeProject(tt.args.devfile)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToComposeProject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			ExpectEqual(t, got, tt.want, "ToComposeProject()")
		})
	}
}

func sampleDevfile(name string, numComponents, numVol int) parser.DevfileObj {
	components := make([]v1alpha2.Component, 0)
	for i := 0; i < numComponents; i++ {
		components = append(components, sampleContainer(fmt.Sprintf("cont-%d", i), "ubuntu"))
	}
	for i := 0; i < numVol; i++ {
		components = append(components, sampleVolume(fmt.Sprintf("vol-%d", i)))
	}
	devObj := parser.DevfileObj{
		Data: &v2.DevfileV2{
			Devfile: v1alpha2.Devfile{
				DevfileHeader: devfile.DevfileHeader{
					Metadata: devfile.DevfileMetadata{Name: name},
				},
				DevWorkspaceTemplateSpec: v1alpha2.DevWorkspaceTemplateSpec{
					DevWorkspaceTemplateSpecContent: v1alpha2.DevWorkspaceTemplateSpecContent{
						Components: components,
					},
				},
			},
		},
	}

	return devObj
}

func sampleProject(name string, numComponents, numVol int) types.Project {
	prj := types.Project{
		Name: name,
	}

	for i := 0; i < numComponents; i++ {
		prj.Services = append(prj.Services, sampleService(fmt.Sprintf("cont-%d", i), "ubuntu"))
	}

	prj.Volumes = make(types.Volumes)
	for i := 0; i < numVol; i++ {
		vol := fmt.Sprintf("vol-%d", i)
		prj.Volumes[vol] = types.VolumeConfig{Name: vol}
	}

	return prj
}

func sampleContainer(name, image string) v1alpha2.Component {
	return v1alpha2.Component{
		Name: name,
		ComponentUnion: v1alpha2.ComponentUnion{
			Container: &v1alpha2.ContainerComponent{
				Container: v1alpha2.Container{
					Image:        image,
					Command:      []string{"/bin/bash", "-c"},
					Args:         []string{"echo", "hello"},
					Env:          []v1alpha2.EnvVar{{Name: "e1", Value: "v1"}, {Name: "e2", Value: "v2"}},
					VolumeMounts: []v1alpha2.VolumeMount{{Name: "vol1", Path: "/tmp/v1"}, {Name: "vol2", Path: "/tmp/v2"}},
				},
			},
		},
	}
}

func sampleService(name, image string) types.ServiceConfig {
	result := types.ServiceConfig{
		Name:        name,
		Image:       image,
		Entrypoint:  types.ShellCommand{"/bin/bash", "-c"},
		Command:     types.ShellCommand{"echo", "hello"},
		NetworkMode: "host",
		Environment: make(types.MappingWithEquals),
		Volumes:     []types.ServiceVolumeConfig{{Type: "volume", Source: "vol1", Target: "/tmp/v1"}, {Type: "volume", Source: "vol2", Target: "/tmp/v2"}},
	}
	env := "v1"
	result.Environment["e1"] = &env
	env = "v2"
	result.Environment["e2"] = &env
	return result
}
func sampleVolume(name string) v1alpha2.Component {
	return v1alpha2.Component{Name: name, ComponentUnion: v1alpha2.ComponentUnion{ComponentType: v1alpha2.VolumeComponentType, Volume: &v1alpha2.VolumeComponent{}}}
}

func Test_convertToService(t *testing.T) {
	type args struct {
		devContainer v1alpha2.Component
	}
	type testCase struct {
		name string
		args args
		want types.ServiceConfig
	}

	tests := []testCase{
		func() testCase {
			return testCase{
				name: "all fields",
				args: args{
					devContainer: sampleContainer("cnt1", "ubuntu"),
				},
				want: sampleService("cnt1", "ubuntu"),
			}
		}(),
		func() testCase {
			return testCase{
				name: "no source code",
				args: args{
					devContainer: sampleContainer("cnt1", "ubuntu"),
				},
				want: sampleService("cnt1", "ubuntu"),
			}
		}(),
		func() testCase {
			cnt := sampleContainer("cnt1", "ubuntu")
			cnt.Container.MountSources = nil
			svc := sampleService("cnt1", "ubuntu")
			return testCase{
				name: "nil MountSources is equivalent to true",
				args: args{
					devContainer: cnt,
				},
				want: svc,
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToService(tt.args.devContainer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertToService() = \n%+v, want \n%+v", got, tt.want)
			}
		})
	}
}

func Test_convertToVolume(t *testing.T) {
	type args struct {
		vol v1alpha2.Component
	}
	tests := []struct {
		name string
		args args
		want types.VolumeConfig
	}{
		{
			name: "works",
			args: args{vol: sampleVolume("my-vol")},
			want: types.VolumeConfig{Name: "my-vol"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToVolume(tt.args.vol); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertToVolume() = %v, want %v", got, tt.want)
			}
		})
	}
}
