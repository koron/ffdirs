package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var srcenv string
var verbose bool

func ffdirs() error {
	flag.StringVar(&srcenv, "srcenv", "PATH", "name of environment variables of directory list")
	flag.BoolVar(&verbose, "verbose", false, "verbose message")
	flag.Parse()
	if flag.NArg() == 0 {
		return errors.New("require one or more file names to find")
	}
	dirs := filepath.SplitList(os.Getenv(srcenv))
	names := flag.Args()
	if len(dirs) == 0 {
		return fmt.Errorf("environment %q has no dirs", srcenv)
	}
	checked := map[string]struct{}{}
	for _, dir := range dirs {
		if _, ok := checked[dir]; ok {
			log.Printf("skip a checked entry: %s", dir)
			continue
		}
		checked[dir] = struct{}{}
		di, err := os.Stat(dir)
		if err != nil {
			log.Printf("ignore an entry: %s: %s", dir, err)
			continue
		}
		if !di.IsDir() {
			log.Printf("not a directory: %s", dir)
			continue
		}
		if verbose {
			log.Printf("scanning a directory: %s", dir)
		}
		for _, name := range names {
			p := filepath.Join(dir, name)
			_, err := os.Stat(p)
			if err != nil {
				if !os.IsNotExist(err) {
					log.Printf("scan error: %s: %s", p, err)
				}
				continue
			}
			fmt.Println(p)
		}
	}
	return nil
}

func main() {
	if err := ffdirs(); err != nil {
		log.Fatal(err)
	}
}
