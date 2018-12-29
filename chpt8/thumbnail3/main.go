package main

import (
	"gopl.io/ch8/thumbnail"
)

func makeThumbnails(filenames []string) error {
	errors := make(chan error)
	for _, f := range filenames {
		go func(f string) {
			_, err := thumbnail.ImageFile(f)
			if err != nil {
				errors <- err
			}
		}(f)
	}
	for range filenames {
		if err := <-errors; err != nil {
			return err // NOTE: incorrect: goroutine leak
		}
	}
	return nil
}

func main() {
	filenames := []string{
		"../Photos/IMG_0045.jpg",
		"../Photos/IMG_0046.jpg",
		"../Photos/IMG_0047.jpg",
		"../Photos/IMG_0048.jpg",
		"../Photos/IMG_0049.jpg",
		"../Photos/IMG_0050.jpg",
		"../Photos/IMG_0051.jpg",
	}
	makeThumbnails(filenames)
}
