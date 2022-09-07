package updenv_test

import (
	"github.com/gowizzard/updenv"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

// TestRead is to test the read method with table driven tests
func TestRead(t *testing.T) {

	tests := []struct {
		path     string
		perm     os.FileMode
		data     []byte
		expected map[string][]byte
	}{
		{
			path: filepath.Join(os.TempDir(), ".env"),
			perm: 0666,
			data: []byte("# Dummy access data\nUSERNAME=admin\nPASSWORD=admin123"),
			expected: map[string][]byte{
				"USERNAME": []byte("USERNAME=admin"),
				"PASSWORD": []byte("PASSWORD=admin123"),
			},
		},
		{
			path: filepath.Join(os.TempDir(), ".env"),
			perm: 0666,
			data: []byte("URL=https://www.test.com/api/v1\nTOKEN=mySecret123456789"),
			expected: map[string][]byte{
				"URL":   []byte("URL=https://www.test.com/api/v1"),
				"TOKEN": []byte("TOKEN=mySecret123456789"),
			},
		},
	}

	for _, value := range tests {

		err := os.WriteFile(value.path, value.data, value.perm)
		if err != nil {
			log.Fatal(err)
		}

		c := updenv.Config{
			Path: value.path,
		}

		err = c.Read()
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(value.expected, c.Environs) {
			t.Fatalf("expected: \"%s\", got: \"%s\"", value.expected, c.Environs)
		}

	}

}
