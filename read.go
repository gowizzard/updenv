package updenv

import (
	"bufio"
	"os"
	"strings"
)

// Read is to read the file from the local filesystem
// We store the scanned lines with the environ key & the bytes from the file
// After that we closed the file
func (c *Config) Read() error {

	file, err := os.Open(c.Path)
	if err != nil {
		return err
	}

	c.Bytes = make(map[string][]byte)

	read := bufio.NewScanner(file)
	for read.Scan() {

		before, _, found := strings.Cut(read.Text(), "=")
		if found {
			c.Bytes[before] = read.Bytes()
		}

	}

	err = file.Close()
	if err != nil {
		return err
	}

	return nil

}
