package main

import (
	"bazil.org/fuse/fs"
)

type FS struct {
	files map[string]MemoryFile
}

var _ = fs.FS(&FS{})

func (f *FS) Root() (fs.Node, error) {
	n := &Dir{
		fs: f,
	}
	return n, nil
}
