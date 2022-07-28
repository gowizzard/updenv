package updenv

import (
	"bytes"
	"log"
	"os"
	"strconv"
	"strings"
)

// Update is to save the updated data to the dot env file
// First we read the bytes of the file, after that we check the updates & if the bytes exists in the file
// After that, we add quotes to the value & join the string, to replace the bytes in the file
func (c *Config) Update() error {

	read, err := os.ReadFile(c.Path)
	if err != nil {
		return err
	}

	for index, value := range c.Updates {

		if bytes.Contains(read, []byte(index)) {

			value = strconv.Quote(value)
			environ := strings.Join([]string{index, value}, "=")

			read = bytes.ReplaceAll(read, c.Bytes[index], []byte(environ))

		}

	}

	err = os.WriteFile(c.Path, read, 0666)
	if err != nil {
		log.Fatal(err)
	}

	return nil

}
