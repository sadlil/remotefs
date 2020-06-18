package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

type Dir struct {
	fs *FS
	// path from  root to this dir; empty for root bucket
	path string
}

var _ = fs.Node(&Dir{})

func (d *Dir) Attr(ctx context.Context, a *fuse.Attr) error {
	a.Mode = os.ModeDir | 0755
	return nil
}

var _ = fs.HandleReadDirAller(&Dir{})

func (d *Dir) ReadDirAll(ctx context.Context) ([]fuse.Dirent, error) {
	fmt.Println("ReadDirAll Invoked")
	fmt.Println(d.path)

	var res []fuse.Dirent

	for path, file := range d.fs.files {
		de := fuse.Dirent{
			Name: path,
			Type: fuse.DT_File,
		}

		if file.IsDir {
			de.Type = fuse.DT_Dir
		}

		res = append(res, de)
	}
	return res, nil
}

var _ = fs.NodeStringLookuper(&Dir{})

func (d *Dir) Lookup(ctx context.Context, name string) (fs.Node, error) {
	file, ok := d.fs.files[name]
	if !ok {
		return nil, fuse.ENOENT
	}

	switch file.IsDir {
	case true:
		return &Dir{
			fs:   d.fs,
			path: name,
		}, nil
	case false:
		return &File{
			dir:  d,
			name: []byte(name),
		}, nil
	}
	return nil, fuse.ENOENT
}

var _ = fs.NodeMkdirer(&Dir{})

func (d *Dir) Mkdir(ctx context.Context, req *fuse.MkdirRequest) (fs.Node, error) {
	d.fs.files[req.Name + "/"] = MemoryFile{
		IsDir:   true,
		Created: time.Now(),
	}

	return &Dir{
		fs:   d.fs,
		path: req.Name,
	}, nil
}

var _ = fs.NodeCreater(&Dir{})

func (d *Dir) Create(ctx context.Context, req *fuse.CreateRequest, resp *fuse.CreateResponse) (fs.Node, fs.Handle, error) {
	d.fs.files[req.Name] = MemoryFile{
		Created: time.Now(),
	}

	f := &File{
		dir:     d,
		name:    []byte(req.Name),
		writers: 1,
		// file is empty at Create time, no need to set data
	}
	return f, f, nil
}

var _ = fs.NodeRemover(&Dir{})

func (d *Dir) Remove(ctx context.Context, req *fuse.RemoveRequest) error {
	delete(d.fs.files, req.Name+"/")
	return nil
}
