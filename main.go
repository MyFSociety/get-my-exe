package main

import (
	"os"
)

func main() {
	// help code
	for _, val := range os.Args {
		if val == "help" {
			getHelp()
			return
		}
	}

	// getmyexe code
	//getFileFromURL("https://mirrors.abhy.me/archlinux/iso/2023.02.01/archlinux-2023.02.01-x86_64.iso")

}
