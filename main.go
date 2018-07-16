package main

import "os"

func main() {
	document := getDocument(os.Args[1])
	result := getResult(document)
	if result != nil {
		 result.Show()
	}

}




