package main

import (
	"github.com/piotr-gladysz/estate-compare/pkg/cli"
	_ "github.com/piotr-gladysz/estate-compare/pkg/hack/pprof"
)

func main() {
	cmd := cli.CreateCLICommand()

	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
