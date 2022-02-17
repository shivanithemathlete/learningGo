package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("sloth.txt")
	if err != nil {
		fmt.Printf("error opening sloth.txt: %v", err)
	}
	defer file.Close()
	bytesRead := make([]byte, 33)
	n, err := file.Read(bytesRead)
	if err != nil {
		fmt.Printf("error opening sloth.txt: %v", err)
	}
	fmt.Printf("We read \"%s\" into bytesRead (%d bytes)\n\n", string(bytesRead), n)

	n, err = file.Read(bytesRead)
	if err != nil {
		fmt.Printf("error reading from sloth.txt: %v", err)
	}

	fmt.Printf("We read \"%s\" into bytesRead (%d bytes)",
		string(bytesRead), n)
	n, err = file.Read(bytesRead)
	if err != nil {
		fmt.Printf("error reading from sloth.txt: %v", err)
	}

	fmt.Printf("We read \"%s\" into bytesRead (%d bytes)",
		string(bytesRead), n)

}
