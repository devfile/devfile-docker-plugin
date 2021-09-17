package cmd

import (
	"errors"
	"github.com/devfile/devrunner/detection/miniBenchmarker"
	"github.com/go-git/go-git/v5"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

const (
	DirTypeFlag = "dir"
	GitTypeFlag = "git"
)

func GetInputStrategy(pathType string) (InputStrategy, error) {
	switch pathType {
	case DirTypeFlag:
		return DirInputStrategy{}, nil
	case GitTypeFlag:
		return GitInputStrategy{}, nil
	default:
		return nil, errors.New("unknown processor")
	}
}

type InputStrategy interface {
	GetPath(path string) (string, error)
}

type GitInputStrategy struct{}
type DirInputStrategy struct{}

func (this GitInputStrategy) GetPath(path string) (string, error) {
	miniBenchmarker.GetInstance().StartStage("GitInputStrategy")
	rand.Seed(time.Now().UnixNano())

	dirName := filepath.Join(os.TempDir(), randSeq(10))

	_, err := git.PlainClone(dirName, false, &git.CloneOptions{
		URL:      path,
		Progress: os.Stdout,
	})
	if err != nil {
		return "", err
	}
	miniBenchmarker.GetInstance().EndStage("GitInputStrategy")

	return dirName, nil
}

func (this DirInputStrategy) GetPath(path string) (string, error) {
	miniBenchmarker.GetInstance().StartStage("DirInputStrategy")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", errors.New("dir does not exist")
	}
	miniBenchmarker.GetInstance().EndStage("DirInputStrategy")

	return filepath.Clean(path), nil
}

func randSeq(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
