// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package updenv_test

import (
	"bytes"
	"github.com/gowizzard/updenv"
	"io"
	"os"
	"path/filepath"
	"testing"
)

// TestUpdate is to test the update method with table driven tests
func TestUpdate(t *testing.T) {

	tests := []struct {
		path     string
		bytes    []byte
		entries  map[string][]byte
		updates  map[string]string
		expected []byte
	}{
		{
			path:  filepath.Join(os.TempDir(), ".env"),
			bytes: []byte("USERNAME=admin\nPASSWORD=admin123"),
			entries: map[string][]byte{
				"USERNAME": []byte("USERNAME=admin"),
				"PASSWORD": []byte("PASSWORD=admin123"),
			},
			updates: map[string]string{
				"USERNAME": "gowizzard",
			},
			expected: []byte("USERNAME=gowizzard"),
		},
		{
			path:  filepath.Join(os.TempDir(), ".env"),
			bytes: []byte("URL=https://www.test.com/api/v1\nTOKEN=mySecret123456789"),
			entries: map[string][]byte{
				"URL":   []byte("URL=https://www.test.com/api/v1"),
				"TOKEN": []byte("TOKEN=mySecret123456789"),
			},
			updates: map[string]string{
				"TOKEN": "myNewSecret123456789",
			},
			expected: []byte("TOKEN=myNewSecret123456789"),
		},
	}

	for _, value := range tests {

		c := updenv.Config{
			Path:     value.path,
			Bytes:    value.bytes,
			Environs: value.entries,
		}

		c.Updates = value.updates
		err := c.Update()
		if err != nil {
			t.Fatal(err)
		}

		file, err := os.Open(value.path)
		if err != nil {
			t.Fatal(err)
		}

		read, err := io.ReadAll(file)
		if !bytes.Contains(read, value.expected) {
			t.Fatalf("expected: \"%s\", got: \"%s\"", value.expected, read)
		}

		err = file.Close()
		if err != nil {
			t.Fatal(err)
		}

	}

}
