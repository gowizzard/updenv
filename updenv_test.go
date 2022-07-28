package updenv

import (
	"log"
	"testing"
)

// Test is to test the functionality of the library
func Test(t *testing.T) {

	c := Config{
		Path: "/path/to/your/file/.env",
	}

	err := c.Read()
	if err != nil {
		log.Fatal(err)
	}

	c.Updates = make(map[string]string)
	c.Updates["EXAMPLE"] = "true"

	err = c.Update()
	if err != nil {
		log.Fatal(err)
	}

}
