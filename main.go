package main

import (
	"embed"
	"flag"
	"fmt"
	"io"
	"os"
)

//go:embed all:content
var content embed.FS

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)

	// No flags for now
	if err := flags.Parse(args[1:]); err != nil {
		return err
	}

	rest := flags.Args()

	// Require a directory name
	if len(rest) < 1 {
		return fmt.Errorf("no name given as the first argument")
	}

	dir := rest[0]

	// Make sure that dir does not already exist
	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		return fmt.Errorf("%q already exists", dir)
	}

	// Create the dir and dir/src
	if err := os.MkdirAll(dir+"/src", os.ModePerm); err != nil {
		return err
	}

	// Enter the new directory
	if err := os.Chdir(dir); err != nil {
		return err
	}

	entries, err := content.ReadDir("content")
	if err != nil {
		return err
	}

	for _, e := range entries {
		if !e.IsDir() {
			if err := writeFile(e.Name()); err != nil {
				return err
			}

		} else {
			if e.Name() == "src" {
				srcEntries, err := content.ReadDir("content/src")
				if err != nil {
					return err
				}

				for _, e := range srcEntries {
					if !e.IsDir() {
						if err := writeFile("src/" + e.Name()); err != nil {
							return err
						}
					}
				}
			}
		}
	}

	return nil
}

func writeFile(name string) error {
	data, err := content.ReadFile("content/" + name)
	if err != nil {
		return fmt.Errorf("writeFile: %w", err)
	}

	return os.WriteFile(name, data, 0644)
}
