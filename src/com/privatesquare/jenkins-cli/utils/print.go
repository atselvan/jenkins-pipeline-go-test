package utils

import "fmt"

func PrintFolderList(folderList []string) {
	for _, folder := range folderList {
		fmt.Println(folder)
	}
	fmt.Printf("Number of folders in jenkins : %d\n", len(folderList))
}
