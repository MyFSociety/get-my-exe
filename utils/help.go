package utils

import "fmt"

func getHelp() {
	help_string := "GetMyExe help :) \n\n\n Example: \n \t >> getmyexe.exe <package_name> <hashing_algorithm> <file_hash> \n\n\n Available Algorithm : [arch-linux] \n\n"

	fmt.Println(help_string)
}
