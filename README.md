catargs

Arguments sort of like cat with - representing stdin.

This returns a list of readers and a list of errs.

Most will only care about the contents of all of the
files together.  This is easily accomplished with
an io.MultiReader and is the example given in cat/cat.go.
