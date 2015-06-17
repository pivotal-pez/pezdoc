package main_test

import (
	"os"
	"path"

	. "github.com/pivotalservices/bootswag"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Bootswag", func() {
	var (
		TEST_TARGET        string
		TEST_NAME          string
		TEST_UI_NAME       string
		TEST_UI_DIR        string
		TEST_MANIFEST_PATH string
		TEST_UI_INDEX      string
	)

	AfterEach(func() {
		_, err := os.Open(TEST_TARGET)
		if os.IsNotExist(err) {
			return
		} else {
			if err := os.RemoveAll(TEST_TARGET); err != nil {
				panic("Could not clean up test artifacts.")
			}
		}
	})

	Describe("default ui at specified location", func() {
		TEST_TARGET = path.Join(ORGPATH, "/testdir")
		TEST_NAME = "testname"
		TEST_UI_NAME = "default"
		TEST_UI_DIR = path.Join(TEST_TARGET, "/swagger-ui")
		TEST_MANIFEST_PATH = path.Join(TEST_TARGET, "/manifest.yml")
		TEST_UI_INDEX = path.Join(TEST_UI_DIR, "/index.html")

		It("is created successfully", func() {
			BuildTarget(TEST_TARGET, TEST_NAME, TEST_UI_NAME)
			Ω(DirectoryExists(TEST_TARGET)).To(BeTrue())
			Ω(DirectoryExists(TEST_UI_DIR)).To(BeTrue())
			Ω(fileExists(TEST_MANIFEST_PATH)).To(BeTrue())
			Ω(matchStringInFile(TEST_MANIFEST_PATH, TEST_NAME)).To(BeTrue())
			Ω(matchStringInFile(TEST_UI_INDEX, "Pez")).To(BeFalse())
		})
	})

	Describe("pez ui at specified location", func() {
		TEST_TARGET = path.Join(ORGPATH, "/testdir")
		TEST_NAME = "peztest"
		TEST_UI_NAME = "pez"
		TEST_UI_DIR = path.Join(TEST_TARGET, "/swagger-ui")
		TEST_MANIFEST_PATH = path.Join(TEST_TARGET, "/manifest.yml")
		TEST_UI_INDEX = path.Join(TEST_UI_DIR, "/index.html")

		It("is created successfully", func() {
			BuildTarget(TEST_TARGET, TEST_NAME, TEST_UI_NAME)
			Ω(DirectoryExists(TEST_TARGET)).To(BeTrue())
			Ω(DirectoryExists(TEST_UI_DIR)).To(BeTrue())
			Ω(fileExists(TEST_MANIFEST_PATH)).To(BeTrue())
			Ω(matchStringInFile(TEST_MANIFEST_PATH, TEST_NAME)).To(BeTrue())
			Ω(matchStringInFile(TEST_UI_INDEX, "pez")).To(BeTrue())
		})
	})
})
