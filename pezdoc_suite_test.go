package main_test

import (
	"bytes"
	"io/ioutil"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPezdoc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pezdoc Suite")
}

func fileExists(filepath string) bool {
	if _, err := os.Stat(filepath); err == nil {
		return true
	}
	return false
}

func matchStringInFile(filename, testname string) bool {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("Error opening file.")
	}

	if i := bytes.Index(f, []byte(testname)); i > -1 {
		return true
	}
	return false
}
