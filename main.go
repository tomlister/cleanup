package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

func main() {
	fmt.Println("Cleaning...")
	cleanup("."+os.Args[1])
}
var c = false;
func cleanup(extension string) {
	current_time := time.Now().Local()
	ct := current_time.Format("2006-01-02")
	pathS, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	_ = os.Mkdir(os.Args[1]+" "+ct, os.ModePerm)
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
				os.Rename(f.Name(), os.Args[1]+" "+ct+"/"+f.Name())
			}
		}
		return nil
	})
	fmt.Println("Done.")
}