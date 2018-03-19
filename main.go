package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

func main() {
	if (len(os.Args) > 1) {
		fmt.Println("Cleaning...")
		cleanup("."+os.Args[1])
	} else {
		fmt.Println("Usage: cleanup extension [new-directory]")
	}
}
var c = false;
func cleanup(extension string) {
	current_time := time.Now().Local()
	ct := current_time.Format("2006-01-02")
	pathS, err := os.Getwd()
	name := ""
	if err != nil {
		panic(err)
	}
	if (len(os.Args) > 2) {
		name = os.Args[2]
	} else {
		name = os.Args[1]+" "+ct
	}
	_ = os.Mkdir(name, os.ModePerm)
	filepath.Walk(pathS, func(path string, f os.FileInfo, _ error) error {
		if f.IsDir() {
			if (c == false) {
				c = true;
			} else {
				return filepath.SkipDir
			}
		} else if !f.IsDir() {
			r, err := regexp.MatchString(extension, f.Name())
			if err == nil && r {
				fmt.Println(f.Name())
				os.Rename(f.Name(), name+"/"+f.Name())
			}
		}
		return nil
	})
	fmt.Println("Done.")
}