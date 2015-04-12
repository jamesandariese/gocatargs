package gocatargs

import (
	"flag"
	"os"
	"io"
)

// func CloseAll(readers []io.Closer) {
// 	for _, elt := range readers {
// 		elt.Close()
// 	}
// }

type Reader struct {
	r io.ReadCloser
}

type OneReader struct {
	r io.Reader
	readers []io.ReadCloser
}

func (r *Reader) Read(p []byte) (n int, err error) {
	return r.r.Read(p)
}

func (r *Reader) Close() error {
	return r.r.Close()
}

func (r OneReader) Read(p []byte) (n int, err error) {
	return r.r.Read(p)
}

func (r OneReader) Close() (err error) {
	for _, reader := range(r.readers) {
		terr := reader.Close()
		if terr != nil {
			err = terr
		}
	}
	return
}

func NewReaders() (readers []*Reader, errs []error) {
	return testable_NewReaders(os.Stdin, flag.Args())
}

func testable_NewReaders(stdin io.ReadCloser, args []string) (readers []*Reader, errs []error) {
	if len(args) > 0 {
		for _, arg := range args {
			if arg == "-" {
				readers = append(readers, &Reader{stdin})
			} else {
				if newfile, err := os.Open(arg); err != nil {
					errs = append(errs, err)
				} else {
					readers = append(readers, &Reader{newfile})
				}
			}
		}
	} else {
		return []*Reader{&Reader{stdin}}, nil
	}
	return
}

func NewOneReader() (OneReader, error) {
	return testable_NewOneReader(os.Stdin, flag.Args())
}

func testable_NewOneReader(stdin io.ReadCloser, args []string) (r OneReader, err error) {
	readers, errs := testable_NewReaders(stdin, args)
	if len(errs) > 0 {
		err = errs[0]
	} else {
		justreaders := []io.Reader{}
		for _, elt := range readers {                       
			r.readers = append(r.readers, elt)
			justreaders = append(justreaders, elt)
		}                                                       
		r.r = io.MultiReader(justreaders...)
	}
	return
}
