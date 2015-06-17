package main

import (
	"fmt"
	"html/template"
	"os"

	"github.com/codegangsta/cli"
)

type manifest struct {
	Name string
}

var createCli = cli.Command{
	Name: "create",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "target, t",
			Usage: "Specify web root for swagger ui project",
		},
		cli.StringFlag{
			Name:  "name, n",
			Value: "api-docs",
			Usage: "Specify cf app hostname (api-docs)",
		},
		cli.BoolFlag{
			Name:  "pez, p",
			Usage: "Use pez-specific ui",
		},
	},
	Aliases: []string{"c"},
	Usage:   "bootstrap swagger-ui server",
	Action: func(c *cli.Context) {
		n := c.String("name")
		t := c.String("target")
		if len(t) == 0 {
			fmt.Println("USAGE: bootswag create -t <target-dir>")
			return
		}

		u := "default"
		if c.Bool("pez") {
			u = "pez"
		}

		if dst, err := os.Stat(t); (err != nil) || (!dst.IsDir()) {
			BuildTarget(t, n, u)
		} else {
			fmt.Println("Target directory already exists. Exiting.")
			os.Exit(1)
		}
		os.Exit(0)
	},
}

//BuildTarget - creates project at specified location and copies/generates necessary assets.
func BuildTarget(target, name, ui string) {
	fmt.Println("Creating: " + target)

	if err := os.MkdirAll(target, 0755); err != nil {
		fmt.Println("Could not create directory.")
		panic(err)
	}

	if err := os.Chdir(target); err != nil {
		fmt.Println("Could not chdir.")
		panic(err)
	}

	if err := CopyFile(PROJECT_ROOT+WEB_SRC, WEB_DEST); err != nil {
		fmt.Println("Could not copy web.go.")
		panic(err)
	}

	if err := GenerateManifest(name); err != nil {
		fmt.Println("Could not generate manifest.yml.")
		panic(err)
	}

	if err := CopyDir(PROJECT_ROOT+UI_SRC+"/"+ui, UI_DEST); err != nil {
		fmt.Println("Could not copy UI directory.")
		panic(err)
	}
	fmt.Println("Done.")
}

//GenerateManifest - creates a manifest with the specified name.
func GenerateManifest(name string) (err error) {
	m := manifest{name}

	tmpl, err := template.New("fest").Parse(MANIFEST_TXT)
	if err != nil {
		fmt.Println("Could not create manifest template.")
		return
	}

	f, err := os.Create(MANIFEST_FILE)
	if err != nil {
		fmt.Println("Could not create manifest.")
		return
	}

	defer f.Close()

	err = tmpl.Execute(f, m)
	if err != nil {
		fmt.Println("Could not execute template.")
	}
	return
}
