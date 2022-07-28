package updenv

// Config is to save essential data, like the file path
// The updates & the bytes map
type Config struct {
	Path    string
	Updates map[string]string
	Bytes   map[string][]byte
}
