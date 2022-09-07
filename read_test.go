package updenv_test

import (
	"github.com/gowizzard/updenv"
	"reflect"
	"testing"
	"testing/fstest"
)

// TestRead is to test the read method with table driven tests
func TestRead(t *testing.T) {

	tests := []struct {
		data     []byte
		expected map[string][]byte
	}{
		{
			data: []byte("# Dummy access data\nUSERNAME=admin\nPASSWORD=admin123"),
			expected: map[string][]byte{
				"USERNAME": []byte("USERNAME=admin"),
				"PASSWORD": []byte("PASSWORD=admin123"),
			},
		},
		{
			data: []byte("URL=https://www.test.com/api/v1\nTOKEN=mySecret123456789"),
			expected: map[string][]byte{
				"URL":   []byte("URL=https://www.test.com/api/v1"),
				"TOKEN": []byte("TOKEN=mySecret123456789"),
			},
		},
	}

	for _, value := range tests {

		c := updenv.Config{
			Filesystem: fstest.MapFS{
				".env": {
					Data: value.data,
				},
			},
			Path: ".env",
		}

		err := c.Read()
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(value.expected, c.Environs) {
			t.Fatalf("expected: \"%s\", got: \"%s\"", value.expected, c.Environs)
		}

	}

}
