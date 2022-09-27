// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package updenv

import (
	"bytes"
	"log"
	"os"
	"strconv"
	"strings"
)

// Update is to save the updated data to the dot env file
// First we check the updates & if the bytes exists in the file
// After that, we join the string, to replace the bytes in the file
func (c *Config) Update() error {

	for index, value := range c.Updates {
		if bytes.Contains(c.Bytes, []byte(index)) {
			if c.Quotes {
				value = strconv.Quote(value)
			}
			environ := strings.Join([]string{index, value}, "=")
			c.Bytes = bytes.ReplaceAll(c.Bytes, c.Environs[index], []byte(environ))
		}
	}

	err := os.WriteFile(c.Path, c.Bytes, 0666)
	if err != nil {
		log.Fatal(err)
	}

	return nil

}
