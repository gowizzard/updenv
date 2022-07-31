package updenv

// Clear is to remove the updates & the bytes
// from the configuration struct
func (c *Config) Clear() {
	c.ClearUpdates()
	c.ClearBytes()
}

// ClearUpdates is to remove the content from
// the updates type in configuration struct
func (c *Config) ClearUpdates() {
	c.Updates = nil
}

// ClearBytes is to clear the read bytes
// from the read function
func (c *Config) ClearBytes() {
	c.Bytes = nil
}
