package updenv

import (
	"bytes"
	"io"
)

// Config is to save essential data, like the file path
// The updates & the bytes map
type Config struct {
	Path     string
	Quotes   bool
	Updates  map[string]string
	Bytes    []byte
	Environs map[string][]byte
	buffer   bytes.Buffer
	reader   io.Reader
}
