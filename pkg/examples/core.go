package examples

import (
	"io/ioutil"
)

// split a string
func Ex() (string, error) {
	bts, e := ioutil.ReadFile("/path/to/file.txt")
	if e != nil {
		return "", e
	}
	return string(bts), nil
}

// read a file

// create a client

// check file existence

// check a map holds a key

// read an environment variable

// sort a slice

// parse a time

// format a time

// base64 encode and decode

// get sha256 hash

// format a string

// read a stream to a string
