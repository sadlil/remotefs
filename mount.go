package main

import (
	"log"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

func mount(mountpoint string) error {
	c, err := fuse.Mount(mountpoint)
	if err != nil {
		log.Println("Err mount "+ mountpoint, err)
		return err
	}
	defer c.Close()

	filesys := &FS{
		files: make(map[string]MemoryFile),
	}
	log.Println("Starting server")
	if err := fs.Serve(c, filesys); err != nil {
		return err
	}

	// check if the mount process has an error to report
	<-c.Ready
	if err := c.MountError; err != nil {
		return err
	}

	return nil
}
