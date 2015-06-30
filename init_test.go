package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Swagger not installed", func() {
	It("swaggerExists should return false", func() {
		SWAGGER_ROOT = "garbagetypeofstringthatshouldnotexist"
		_ = SWAGGER_ROOT
		Ω(swaggerExists()).To(BeFalse())
	})
})

var _ = Describe("Swagger is installed", func() {
	It("should return true", func() {
		Ω(swaggerExists()).To(BeTrue())
	})
})

var _ = Describe("Installs swagger lib", func() {
	It("should return true", func() {
		Ω(swaggerExists()).To(BeTrue())
	})
})
