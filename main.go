package main

import (
	"fmt"
	"harry/get-my-exe/core"
	"harry/get-my-exe/utils"
	"log"
	"os"
)

func main() {
	// help code
	for _, val := range os.Args {
		if val == "help" {
			fmt.Println(utils.GetHelp())
			return
		}
	}

	// getmyexe code
	if len(os.Args) < 2 {
		fmt.Println("Not enough arguments")
		return
	}

	package_name := os.Args[1]

	package_struct := core.GetPackage(package_name)

	if package_struct.Name == "" {
		fmt.Println("Package not found")
		return
	}

	fileName := core.GetFileFromURL(package_struct.URL)

	hashing_algorithm := package_struct.SHAType

	// file integrity check

	download_file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	file_good := utils.ValidateFile(download_file, package_struct.SHA, hashing_algorithm)

	if file_good {
		fmt.Println("Hashes match good to go :)")
	} else {
		fmt.Println("Hashes don't match u at your own risk ;(")
	}
}
