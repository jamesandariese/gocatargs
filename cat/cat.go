package main

import "github.com/jamesandariese/gocatargs"
import "flag"
import "os"
import "io"

func main() {
	flag.Parse()

	filereaders, errs := gocatargs.NewReaders()

	if len(errs) > 0 {
		panic(errs)
	}
	readers := []io.Reader{}
	for _, elt := range filereaders {
		readers = append(readers, elt)
		defer elt.Close()
	}
	io.Copy(os.Stdout, io.MultiReader(readers...))

}
