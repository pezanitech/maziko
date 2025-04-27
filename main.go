package main

import (
	"fmt"
	"maziko/backend/cmd"
	"os"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "dev" {
		fmt.Println("Starting development mode...")
		cmd.Dev()

		return
	}

	cmd.Execute()
}
