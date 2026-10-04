package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
)




func main () {

	
	homePath, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error getting user home directory: %v\n", err)
	}

	dotTalosPath := filepath.Join(homePath, ".talos")

	sourcePath := flag.String("source", filepath.Join(dotTalosPath, "config.d"), "Source directory")
	outputPath := filepath.Join(dotTalosPath, "config")

	flag.Parse()

	fmt.Printf("Merging files at %v into %v\n", *sourcePath, outputPath)

	err = os.RemoveAll(outputPath)
	if err != nil {
		fmt.Println("Error deleting file:", err)
		return
	}

	err = filepath.WalkDir(*sourcePath, func(s string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}
		

		fmt.Printf("Merging %v \n", s)

		cmd := exec.Command("talosctl", "config", "merge", s)
		_, err = cmd.CombinedOutput()

		return err
			
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Generated %v\n", outputPath)

}