package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"loglinter/internal/loglinter"
)

func main() {
	singlechecker.Main(loglinter.Analyzer)
}
