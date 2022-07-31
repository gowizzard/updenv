# updenv

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gowizzard/updenv.svg)](https://golang.org/) [![Go](https://github.com/gowizzard/updenv/actions/workflows/go.yml/badge.svg)](https://github.com/gowizzard/updenv/actions/workflows/go.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/gowizzard/updenv)](https://goreportcard.com/report/github.com/gowizzard/updenv) [![Go Reference](https://pkg.go.dev/badge/github.com/gowizzard/updenv.svg)](https://pkg.go.dev/github.com/gowizzard/updenv) [![GitHub issues](https://img.shields.io/github/issues/gowizzard/updenv)](https://github.com/gowizzard/updenv/issues) [![GitHub forks](https://img.shields.io/github/forks/gowizzard/updenv)](https://github.com/gowizzard/updenv/network) [![GitHub stars](https://img.shields.io/github/stars/gowizzard/updenv)](https://github.com/gowizzard/updenv/stargazers) [![GitHub license](https://img.shields.io/github/license/gowizzard/updenv)](https://github.com/gowizzard/updenv/blob/master/LICENSE)

With this small library it should be possible to update environment variables in a dot env file. With this library it is not possible to read out the variables, because there are already very good other libraries for this.

# Install



```console
go get github.com/gowizzard/updenv
```

# How to use?

Here is an example of how to update a dot env file. Of course, only variables that exactly match the key can be updated. New variables cannot be added.

If you want to edit multiple files, just create another config struct and fill it with the path to the file. The functions can then be applied to the variable that is linked to the struct. If you want to clear the update struct, you can use the `updenv.ClearUpdates()` function. The function also exists for the read bytes here `updenv.ClearBytes()`. If you want to empty both, you can use the following function `updenv.Clear()`. Execute this function if you want to read a file multiple times with adjustments.

If you prefer not to use quotes for your variables, then you can set this in the config struct **by setting the quotes value to false**, or omitting it directly.

```go
c := updenv.Config{
	Path:   "/path/to/your/file/.env",
	Quotes: true,
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
```

## Special thanks

Thanks to [JetBrains](https://github.com/JetBrains) for supporting me with this and other [open source projects](https://www.jetbrains.com/community/opensource/#support).