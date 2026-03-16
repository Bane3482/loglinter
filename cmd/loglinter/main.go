package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/Bane3482/loglinter/pkg/loglinter"
)

func main() {
	singlechecker.Main(loglinter.Analyzer)
}
