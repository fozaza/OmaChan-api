package jp2a

import (
	"fmt"
	"log"
	"os/exec"
)

func Print(image_path string) {
	path := image_path
	// path, err := module.Get_file_path(image_path)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	cmd := exec.Command("jp2a", "--color", "--width=80", path)

	output, err := cmd.Output()

	if err != nil {
		// Log error output for debugging
		if exitError, ok := err.(*exec.ExitError); ok {
			log.Fatalf("jp2a failed with output: %s\nError: %v", exitError.Stderr, err)
		}
		log.Fatalf("Failed to run jp2a: %v", err)
	}
	fmt.Println(string(output))
}
