package main

import (
	"errors"

	"github.com/nobe4/go-cli-comparison/internal/library"
	"github.com/nobe4/go-cli-comparison/internal/result"
	"github.com/nobe4/go-cli-comparison/internal/spec"
)

var errNoMarker = errors.New("could not find marker comments")

const (
	expectedParts = 3
)

type templateData struct {
	Libs          []*library.Library
	Tests         []spec.Test
	ResultByLibs  result.Result
	ResultByTests result.Result
}

type repo struct {
	PushedAt       string `json:"pushed_at"`
	StargazerCount int    `json:"stargazers_count"`
}

func main() {
	if err := updateReadme(); err != nil {
		panic(err)
	}

	if err := updateSpecs(); err != nil {
		panic(err)
	}
}
