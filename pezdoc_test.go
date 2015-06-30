package main_test

import (
	"path"

	. "github.com/pivotal-pez/pezdoc"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PezDoc", func() {

	Describe("pez ui at specified location", func() {
		var (
			T2_TARGET        = path.Join(ORGPATH, "/test")
			T2_NAME          = "peztest"
			T2_UI_DIR        = path.Join(T2_TARGET, "/swagger-ui")
			T2_MANIFEST_PATH = path.Join(T2_TARGET, "/manifest.yml")
			T2_UI_INDEX      = path.Join(T2_UI_DIR, "/index.html")
		)

		It("is created successfully", func() {
			BuildTarget(T2_TARGET, T2_NAME)

			By("creating the project and ui directories")
			Ω(DirectoryExists(T2_TARGET)).To(BeTrue())
			Ω(DirectoryExists(T2_UI_DIR)).To(BeTrue())

			By("correctly generating the manifest")
			Ω(fileExists(T2_MANIFEST_PATH)).To(BeTrue())
			Ω(matchStringInFile(T2_MANIFEST_PATH, T2_NAME)).To(BeTrue())
			Ω(matchStringInFile(T2_UI_INDEX, "pez")).To(BeTrue())
		})
	})
})

var _ = AfterSuite(func() {
	var (
		target = path.Join(ORGPATH, "/test")
	)

	var _ = RemoveDir(target)
})
