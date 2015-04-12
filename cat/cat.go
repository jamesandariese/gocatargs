package main

import "github.com/jamesandariese/gocatargs"
import "flag"
import "os"
import "io"

func main() {
	flag.Parse()

	if reader, err := gocatargs.NewOneReader(); err != nil {
		panic(err)
	} else {
		defer reader.Close()
		io.Copy(os.Stdout, reader)
	}
}
