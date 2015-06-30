package main

import (
	"fmt"
	"io"
	"os"
)

func CopyFile(source string, dest string) (err error) {
	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourcefile.Close()

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destfile.Close()

	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		sourceinfo, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, sourceinfo.Mode())
		}
	}
	return
}

func CopyDir(source string, dest string) (err error) {
	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		return err
	}

	directory, _ := os.Open(source)
	objects, err := directory.Readdir(-1)

	for _, obj := range objects {
		srcFilePointer := source + "/" + obj.Name()
		dstFilePointer := dest + "/" + obj.Name()

		if obj.IsDir() {
			err = CopyDir(srcFilePointer, dstFilePointer)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			err = CopyFile(srcFilePointer, dstFilePointer)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	return
}

func DirectoryExists(dirpath string) bool {
	_, err := os.Open(dirpath)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func RemoveDir(dirpath string) bool {
	_, err := os.Open(dirpath)
	if os.IsNotExist(err) {
		fmt.Println("Directory not found:", dirpath)
		return true
	} else {
		fmt.Println("Removing:", dirpath)
		if err := os.RemoveAll(dirpath); err != nil {
			return false
		}
		fmt.Println("Done.")
	}
	return true
}
