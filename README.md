# pezdoc

A utility to generate a CF pushable web application for hosting API documentation.

## Overview
Pezdoc is a CLI that automates many of the steps required when working with the `pivotal-pez/swagger` library.  Pezdoc depends on, and will initialize your go environment with, `pivotal-pez/swagger` for the generation of swagger data.  Further, it will create a CF-pushable web application for hosting your API documentation.

## Installation
With Go and git installed:

``` shell
go get github.com/pivotal-pez/pezdoc
```
will download, compile, and install the package into your `$GOPATH`.

## Usage
To simply install the swagger library dependency:

``` shell
pezdoc init
```

however, the `create` command (below) will check for the swagger dependency and install it should it not be resident in your `$GOPATH` environment.

Use the `create` subcommand to build the web application.  `--target` specifies where to create the application directory.  `--name` is used when generating the `manifest.yml`.  

``` shell
pezdoc create --target <target-dir> --name <app-name>
```

The `create` command will build an empty static web server with boilerplate swagger UI.  *(Note: It will be necessary to generate the swagger documentation -- see below)*

```
.
├── manifest.yml
├── swagger-ui/
│   ├── css/
│   ├── images/
│   ├── index.html
│   ├── lib/
│   ├── o2c.html
│   ├── swagger-ui.js
│   └── swagger-ui.min.js
└── web.go
```

## Annotate your APIs and generate swagger docs
It is necessary to annotate your APIs before you can generate data to drive the Swagger UI.  Refer to the full `swagger` documentation [here](http://github.com/pivotal-pez/swagger).

### Annotations

#### main.go example
The annotations in the main are at the top of the file above the package declaration.
``` go
// @APIVersion 1.0.0
// @APITitle Pez Search
// @APIDescription Search endpoint for Pez
// @SubApi Type Search [/v1/search]
package main

...
```
#### controller.go example
Annotations for the controllers are placed just above the controller method.
``` go
...

// @Title search
// @Description Search API
// @Accept  json
// @Param   X-API-KEY       header  string  true        "APIKEY"
// @Param   q               query   string  true        "Retrieve items matching search terms"
// @Success 200 {object}  ResponseMessage
// @Failure 400 {object}  ResponseMessage
// @Resource /v1/search
// @Router /v1/search [get]
func (c *searchController) find(req *http.Request, render render.Render) {
     // application logic
}

...

```

### Swagger Generation
Once you have annotated your APIs, it is necessary to generate the swagger docs.  You will need to tailor the command based on the structure of your API project. Refer to the full `swagger` documentation [here](http://github.com/pivotal-pez/swagger) for more details.

*Example:*
``` shell
swagger -apiPackage="github.com/pivotal-pez/pezsearch"
        -mainApiFile="github.com/pivotal-pez/pezsearch/main.go"
        -basePath="https://your.apis.url:443"
```

Successful generation of swagger docs should result in the creation of file named `docs.go`.  This document should be placed in the root of your generated web application.

```
.
├── docs.go  <<------------------- generated swagger data
├── manifest.yml
├── swagger-ui/
│   ├── css/
│   ├── images/
│   ├── index.html
│   ├── lib/
│   ├── o2c.html
│   ├── swagger-ui.js
│   └── swagger-ui.min.js
└── web.go
```
### Running the server locally
``` shell
go run web.go docs.go
```

## Wercker integration
`forthcoming`

## TODOs
1. Add library dependencies to Godeps
1. Enable CORS in API project