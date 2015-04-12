package gocatargs

import (
	"testing"
	"os"
	"io"
	"io/ioutil"
	"bytes"
)

func TestReaders(t *testing.T) {
	testfile1, err := ioutil.TempFile("", "catflags_test")
	if err != nil {
		panic(err)
	}
	defer testfile1.Close()
	defer os.Remove(testfile1.Name())
	testfile2, err := ioutil.TempFile("", "catflags_test")
	if err != nil {
		panic(err)
	}
	defer testfile2.Close()
	defer os.Remove(testfile2.Name())

	testfile1.Write([]byte{1,2,3,4,5})
	testfile2.Write([]byte{6,7,8,9})
	
	args := []string{testfile1.Name(), testfile2.Name()}
	filereaders, errs := testable_NewReaders(os.Stdin, args)

	readers := []io.Reader{}

	for _, elt := range(filereaders) {
		readers = append(readers, elt)
		defer elt.Close()
	}

	if len(readers) != 2 {
		t.Error("length of readers is not 2", readers)
	}

	if len(errs) != 0 {
		t.Error("length of errs is not 0", errs)
	}

	mrbytes, err := ioutil.ReadAll(io.MultiReader(readers...))
	if err != nil {
		t.Errorf("error reading, expected 1 2 3 4 5 6 7 8 9: %#v", err)
	}
	if bytes.Compare(mrbytes, []byte{1,2,3,4,5,6,7,8,9}) != 0 {
		t.Errorf("error reading, expected 1 2 3 4 5 6 7 8 9: %#v", mrbytes)
	}
}

func TestReaderImpliedStdin(t *testing.T) {
	fakestdin := ioutil.NopCloser(bytes.NewReader([]byte{1,2,3}))
	filereaders, errs := testable_NewReaders(fakestdin, []string{})

	readers := []io.Reader{}

	for _, elt := range(filereaders) {
		readers = append(readers, elt)
		defer elt.Close()
	}

	if len(readers) != 1 {
		t.Error("length of readers is not 1", readers)
	}

	if len(errs) != 0 {
		t.Error("length of errs is not 0", errs)
	}

	mrbytes, err := ioutil.ReadAll(io.MultiReader(readers...))
	if err != nil {
		t.Errorf("error reading, expected 1 2 3: %#v", err)
	}
	if bytes.Compare(mrbytes, []byte{1,2,3}) != 0 {
		t.Errorf("error reading, expected 1 2 3: %#v", mrbytes)
	}
}

func TestReadersMixedStdin(t *testing.T) {
	testfile1, err := ioutil.TempFile("", "catflags_test")
	if err != nil {
		panic(err)
	}
	defer testfile1.Close()
	defer os.Remove(testfile1.Name())
	testfile2, err := ioutil.TempFile("", "catflags_test")
	if err != nil {
		panic(err)
	}
	defer testfile2.Close()
	defer os.Remove(testfile2.Name())

	testfile1.Write([]byte{1,2,3,4,5})
	testfile2.Write([]byte{6,7,8,9})
	
	args := []string{testfile1.Name(), "-", testfile2.Name()}

	fakestdin := ioutil.NopCloser(bytes.NewReader([]byte{11,12,13}))

	filereaders, errs := testable_NewReaders(fakestdin, args)

	readers := []io.Reader{}

	for _, elt := range(filereaders) {
		readers = append(readers, elt)
		defer elt.Close()
	}

	if len(readers) != 3 {
		t.Error("length of readers is not 3", readers)
	}

	if len(errs) != 0 {
		t.Error("length of errs is not 0", errs)
	}

	mrbytes, err := ioutil.ReadAll(io.MultiReader(readers...))
	if err != nil {
		t.Errorf("error reading, expected 1 2 3 4 5 11 12 13 6 7 8 9: %#v", err)
	}
	if bytes.Compare(mrbytes, []byte{1,2,3,4,5,11,12,13,6,7,8,9}) != 0 {
		t.Errorf("error reading, expected 1 2 3 4 5 11 12 13 6 7 8 9: %#v", mrbytes)
	}
}

func TestOneReaderMixedStdin(t *testing.T) {
	testfile1, err := ioutil.TempFile("", "catflags_test")
	if err != nil {
		panic(err)
	}
	defer testfile1.Close()
	defer os.Remove(testfile1.Name())
	testfile2, err := ioutil.TempFile("", "catflags_test")
	if err != nil {
		panic(err)
	}
	defer testfile2.Close()
	defer os.Remove(testfile2.Name())

	testfile1.Write([]byte{1,2,3,4,5})
	testfile2.Write([]byte{6,7,8,9})
	
	args := []string{testfile1.Name(), "-", testfile2.Name()}

	fakestdin := ioutil.NopCloser(bytes.NewReader([]byte{11,12,13}))

	reader, err := testable_NewOneReader(fakestdin, args)
	if err != nil {
		t.Errorf("Error from testable_OneReader: %#v", err)
	}
	defer reader.Close()
	
	mrbytes, err := ioutil.ReadAll(reader)
	if err != nil {
		t.Errorf("error reading, expected 1 2 3 4 5 11 12 13 6 7 8 9: %#v", err)
	}
	if bytes.Compare(mrbytes, []byte{1,2,3,4,5,11,12,13,6,7,8,9}) != 0 {
		t.Errorf("error reading, expected 1 2 3 4 5 11 12 13 6 7 8 9: %#v", mrbytes)
	}
}

