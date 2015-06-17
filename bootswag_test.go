package main_test

import (
	"os"
	"path"

	. "github.com/pivotalservices/bootswag"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	T1_TARGET string
	T2_TARGET string
)

var _ = Describe("Bootswag", func() {
	Describe("default ui at specified location", func() {
		var (
			T1_TARGET        = path.Join(ORGPATH, "/test1")
			T1_NAME          = "testname"
			T1_UI_NAME       = "default"
			T1_UI_DIR        = path.Join(T1_TARGET, "/swagger-ui")
			T1_MANIFEST_PATH = path.Join(T1_TARGET, "/manifest.yml")
			T1_UI_INDEX      = path.Join(T1_UI_DIR, "/index.html")
		)

		It("is created successfully", func() {
			BuildTarget(T1_TARGET, T1_NAME, T1_UI_NAME)
			Ω(DirectoryExists(T1_TARGET)).To(BeTrue())
			Ω(DirectoryExists(T1_UI_DIR)).To(BeTrue())
			Ω(fileExists(T1_MANIFEST_PATH)).To(BeTrue())
			Ω(matchStringInFile(T1_MANIFEST_PATH, T1_NAME)).To(BeTrue())
			Ω(matchStringInFile(T1_UI_INDEX, "pez")).To(BeFalse())
		})
	})

	Describe("pez ui at specified location", func() {
		var (
			T2_TARGET        = path.Join(ORGPATH, "/test2")
			T2_NAME          = "peztest"
			T2_UI_NAME       = "pez"
			T2_UI_DIR        = path.Join(T2_TARGET, "/swagger-ui")
			T2_MANIFEST_PATH = path.Join(T2_TARGET, "/manifest.yml")
			T2_UI_INDEX      = path.Join(T2_UI_DIR, "/index.html")
		)

		It("is created successfully", func() {
			BuildTarget(T2_TARGET, T2_NAME, T2_UI_NAME)
			Ω(DirectoryExists(T2_TARGET)).To(BeTrue())
			Ω(DirectoryExists(T2_UI_DIR)).To(BeTrue())
			Ω(fileExists(T2_MANIFEST_PATH)).To(BeTrue())
			Ω(matchStringInFile(T2_MANIFEST_PATH, T2_NAME)).To(BeTrue())
			Ω(matchStringInFile(T2_UI_INDEX, "pez")).To(BeTrue())
		})
	})
})

var _ = AfterSuite(func() {
	cleanUp()
})

func cleanUp() {
	var (
		t1 = path.Join(ORGPATH, "/test1")
		t2 = path.Join(ORGPATH, "/test2")
	)
	_, err := os.Open(t1)
	if !os.IsNotExist(err) {
		if err := os.RemoveAll(t1); err != nil {
			panic("Could not clean up test artifacts.")
		}
	}

	_, err = os.Open(t2)
	if !os.IsNotExist(err) {
		if err := os.RemoveAll(t2); err != nil {
			panic("Could not clean up test artifacts.")
		}
	}
}
