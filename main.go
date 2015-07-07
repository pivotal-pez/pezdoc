package main

import (
	"os"
	"path"

	"github.com/codegangsta/cli"
)

const (
	WEB_SRC       = "/static/web.go_"
	WEB_DEST      = "./web.go"
	UI_SRC        = "/ui/pez"
	UI_DEST       = "./swagger-ui"
	MANIFEST_FILE = "./manifest.yml"
	MANIFEST_TXT  = `
---
applications:
- name: {{.Name}}
  memory: 128M
  disk_quota: 128M
  host: {{.Name}}
  buildpack: http://github.com/ryandotsmith/null-buildpack.git
  command: ./build/pezdoc_linux_amd64
`
)

var (
	GOPATH       string = os.Getenv("GOPATH")
	ORGPATH      string = path.Join(GOPATH, "/src/github.com/pivotal-pez")
	SWAGGER_ROOT string = path.Join(ORGPATH, "/swagger")
	PROJECT_ROOT string = path.Join(ORGPATH, "/pezdoc")
)

func main() {

	app := NewApp()
	app.Run(os.Args)
}

// NewApp creates a new cli app
func NewApp() *cli.App {
	app := cli.NewApp()
	app.Name = "PezDoc"
	app.Usage = "Bootstrap Swagger Documentation for Pez"

	app.Commands = append(app.Commands, []cli.Command{
		initCli,
		createCli,
	}...)

	return app
}
