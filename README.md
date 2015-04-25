gocatargs
===

Arguments sort of like cat with - representing stdin.

This uses the remaining arguments from a flag.Parse and returns a list of readers and a list of errs.

Most will only care about the contents of all of the
files together.  This is easily accomplished with
the OneReader which is used in cat/cat.go.

For slightly more control, you can also have access to each
file with the same sort of thing as the OneReader by using
io.MultiReader (that's how OneReader currently works).
