package main

import (
	"fmt"
	"os"
	"path"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	TEST_GOPATH     = path.Join(PROJECT_ROOT, "/test")
	original_gopath = GOPATH
)

var _ = Describe("Initialize swagger", func() {
	Context("Swagger not installed", func() {
		SWAGGER_ROOT = path.Join(TEST_GOPATH, "/src/github.com/pivotal-pez/swagger")
		_ = SWAGGER_ROOT
		It("swaggerExists should return false", func() {
			Ω(swaggerExists()).To(BeFalse())
		})
	})

	Context("Swagger is installed", func() {
		BeforeEach(func() {
			if err := os.Setenv("GOPATH", TEST_GOPATH); err != nil {
				panic("Cannot set GOPATH for testing.")
			}

			if err := os.Mkdir(TEST_GOPATH, 0777); err != nil {
				panic("Could not create test directory")
			}
			installSwagger()
		})

		It("swaggerExists should return true", func() {
			Ω(swaggerExists()).To(BeTrue())
		})

		AfterEach(func() {
			if err := os.Setenv("GOPATH", original_gopath); err != nil {
				panic("OH SHIT: cannot reset GOPATH when done testing.")
			}
			if !RemoveDir(TEST_GOPATH) {
				fmt.Println("Could not remove test directory")
			}
		})
	})
})
