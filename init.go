package main

import (
	"os"
	"os/exec"

	"github.com/codegangsta/cli"
)

var initCli = cli.Command{
	Name: "init",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "force, f",
			Usage: "Force upgrade of existing swagger library",
		},
	},
	Aliases: []string{"i"},
	Usage:   "initialize go environment",
	Action: func(c *cli.Context) {
		if !swaggerExists() || c.Bool("force") {
			installSwagger()
			return
		}
		println("Swagger library is already installed.  Use -f to force upgrade.")
	},
}

func swaggerExists() bool {
	if _, err := os.Stat(SWAGGER_ROOT); err != nil {
		if os.IsNotExist(err) {
			return false
		} else {
			panic(err)
		}
	}
	return true
}

func installSwagger() {
	println("Installing swagger library...")
	cmd := exec.Command("go", "get", "-u", "github.com/pivotalservices/swagger")
	out, err := cmd.CombinedOutput()
	if err != nil {
		println("Error: ", out)
		return
	}
	println("Success!")
}
