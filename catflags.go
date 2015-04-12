package catargs

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

func (r OneReader) Close(p []byte) (err error) {
	for _, reader := range(r.readers) {
		terr := reader.Close()
		if terr != nil {
			err = terr
		}
	}
	return
}

func Readers() (readers []*Reader, errs []error) {
	return testable_Readers(os.Stdin, flag.Args())
}

func testable_Readers(stdin io.ReadCloser, args []string) (readers []*Reader, errs []error) {
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
