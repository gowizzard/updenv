// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// Package updenv is used to exchange information or
// configurations in an environment variables file.
// Different configurations can be applied.
package updenv

import (
	"bytes"
	"io"
)

// Config is to save essential configuration data, like
// the Path, Quotes & Updates. In addition, there are types
// in which information for the update process is stored.
type Config struct {
	Path     string
	Quotes   bool
	Updates  map[string]string
	Bytes    []byte
	Environs map[string][]byte
	buffer   bytes.Buffer
	reader   io.Reader
}
