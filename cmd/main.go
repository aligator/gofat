package main

import (
	"fmt"
	"github.com/spf13/afero"
	"os"

	"github.com/aligator/gofat"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) <= 0 {
		fmt.Println("Please provide a filename.")
		os.Exit(1)
	}

	fsFile, err := os.Open(argsWithoutProg[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer fsFile.Close()

	fat := gofat.New(fsFile)
	fmt.Printf("Opened volume '%v'\n\n", fat.Label())

	afero.Walk(fat, "/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println(path, info.IsDir())
		return nil
	})

	file, err := fat.Open("/grub/themes/starfield/README")
	if err != nil {
		fmt.Println("could not open the root file", err)
		os.Exit(1)
	}

	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		fmt.Println("could not stat the file", err)
		os.Exit(1)
	}
	buffer := make([]byte, stat.Size())
	n, err := file.Read(buffer)
	if err != nil {
		fmt.Println("could not read the file", err)
		os.Exit(1)
	}
	fmt.Println(stat.Size(), n)
	fmt.Println(string(buffer))
}
