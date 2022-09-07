package updenv

import (
	"bufio"
	"io"
	"strings"
)

// Read is to read the file from the local filesystem
// We store the bytes & scanned lines with the environ key
// After that we closed the file
func (c *Config) Read() error {

	file, err := c.Filesystem.Open(c.Path)
	if err != nil {
		return err
	}

	c.reader = io.TeeReader(file, &c.buffer)
	c.Bytes, err = io.ReadAll(c.reader)
	if err != nil {
		return err
	}

	c.Entries = make(map[string][]byte)

	entries := bufio.NewScanner(&c.buffer)
	for entries.Scan() {
		before, _, found := strings.Cut(entries.Text(), "=")
		if found {
			c.Entries[before] = entries.Bytes()
		}
	}
	c.buffer.Reset()

	err = file.Close()
	if err != nil {
		return err
	}

	return nil

}
