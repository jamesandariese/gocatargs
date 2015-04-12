package main

import "github.com/jamesandariese/catargs"
import "flag"
import "os"
import "io"

func main() {
	flag.Parse()

	filereaders, errs := catargs.Readers()

	if len(errs) > 0 {
		panic(errs)
	}
	readers := []io.Reader{}
	for _, elt := range filereaders {
		readers = append(readers, elt)
	}
	io.Copy(os.Stdout, io.MultiReader(readers...))

}
