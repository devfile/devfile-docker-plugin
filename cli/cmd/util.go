package cmd

import (
	"github.com/compose-spec/compose-go/types"
	convert "github.com/devfile/devrunner/pkg/devfile"
	"github.com/devfile/library/pkg/devfile"
	"github.com/devfile/library/pkg/devfile/parser"
)

func convertToProject(devFilepath string) (prj *types.Project, err error) {
	d, _, err := devfile.ParseDevfileAndValidate(parser.ParserArgs{
		Path: devFilepath,
	})
	if err != nil {
		return nil, err
	}

	result, err := convert.ToComposeProject(d)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
