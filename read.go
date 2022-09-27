// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package updenv

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// Read is to read the file from the local filesystem
// We store the bytes & scanned lines with the environ key
// After that we closed the file
func (c *Config) Read() error {

	file, err := os.Open(c.Path)
	if err != nil {
		return err
	}

	c.reader = io.TeeReader(file, &c.buffer)
	c.Bytes, err = io.ReadAll(c.reader)
	if err != nil {
		return err
	}

	c.Environs = make(map[string][]byte)

	read := bufio.NewScanner(&c.buffer)
	for read.Scan() {
		before, _, found := strings.Cut(read.Text(), "=")
		if found {
			c.Environs[before] = read.Bytes()
		}
	}
	c.buffer.Reset()

	err = file.Close()
	if err != nil {
		return err
	}

	return nil

}
