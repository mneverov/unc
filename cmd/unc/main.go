package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/mneverov/unc/pkg/unc"
)

func main() {
	singlechecker.Main(unc.Analyzer)
}
