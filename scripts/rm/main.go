package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	matches, err := filepath.Glob("api/pb/*.pb.go")
	if err != nil {
		panic(err)
	}
	for _, i := range matches {
		err = os.Remove(i)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("remove %s\n", i)
	}
}
