// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package files implements file utility functions.
package files

import (
	"gopkg.in/goyy/goyy.v0/util/strings"
	"io/ioutil"
	"os"
)

// Reports whether the specified file exists.
// Returns true if the file exists, false if it does not exist.
func IsExist(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

// Read reads the file named by filename and returns the contents.
// A successful call returns err == nil, not err == EOF. Because Read
// reads the whole file, it does not treat an EOF from Read as an error
// to be reported.
func Read(filename string) (string, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// Write writes data to a file named by filename.
// If the file does not exist, Write creates it with permissions perm;
// otherwise Write truncates it before writing.
func Write(filename string, data string, perm os.FileMode) error {
	return ioutil.WriteFile(filename, []byte(data), perm)
}

// Create creates the named file mode 0666 (before umask), truncating
// it if it already exists.  If successful, methods on the returned
// File can be used for I/O; the associated file descriptor has mode
// O_RDWR.
// If there is an error, it will be of type *PathError.
func Create(name string) (*os.File, error) {
	return os.Create(name)
}

// Copy copies from src to dst until either EOF is reached on src or an error occurs.
func Copy(dstfile string, srcfile string, perm os.FileMode) error {
	data, rerr := ioutil.ReadFile(srcfile)
	if rerr != nil {
		return rerr
	}
	werr := ioutil.WriteFile(dstfile, data, perm)
	if werr != nil {
		return werr
	}
	return nil
}

// Mkdir creates a new directory with the specified name and permission bits.
// If there is an error, it will be of type *PathError.
func Mkdir(name string, perm os.FileMode) error {
	return os.Mkdir(name, perm)
}

// MkdirAll creates a directory named path, along with any necessary parents,
// and returns nil, or else returns an error.
// The permission bits perm are used for all directories that MkdirAll creates.
// If path is already a directory, MkdirAll does nothing and returns nil.
func MkdirAll(name string, perm os.FileMode) error {
	return os.MkdirAll(name, perm)
}

// Rename renames a file.
func Rename(oldname, newname string) error {
	return os.Rename(oldname, newname)
}

// Remove removes the named file or directory.
// If there is an error, it will be of type *PathError.
func Remove(name string) error {
	return os.Remove(name)
}

// Returns the <a href="http://en.wikipedia.org/wiki/Filename_extension">file
// extension</a> for the given file name, or the empty string if the file has
// no extension.  The result does not include the '{@code .}'.
func GetExtension(fileName string) string {
	fname := strings.Replace(fileName, "\\", "/", -1)
	ext := strings.AfterLast(fname, ".")
	first := strings.Left(ext, 1)
	if first == "/" {
		return ""
	} else {
		return ext
	}
}

// ModTime returns the file modification time
func ModTime(file string) (int64, error) {
	f, e := os.Stat(file)
	if e != nil {
		return 0, e
	}
	return f.ModTime().Unix(), nil
}

// Size returns the length in bytes for regular files
func Size(file string) (int64, error) {
	f, e := os.Stat(file)
	if e != nil {
		return 0, e
	}
	return f.Size(), nil
}