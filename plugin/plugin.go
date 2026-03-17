package main

import (
	"github.com/Bane3482/loglinter/pkg/loglinter"
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("loglinter", New)
}

type plugin struct{}

func New(settings any) (register.LinterPlugin, error) {
	return &plugin{}, nil
}

func (plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{loglinter.Analyzer}, nil
}

func (plugin) GetLoadMode() string {
	return register.LoadModeTypesInfo
}
