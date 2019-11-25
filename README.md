# go-run-sass

A simple command line tool wrapper to generate css using `go-libsass` which otherwise recompiles the whole library when simply imported as `github.com/wellington/go-libsass`.

Example:

```go
err := execute("go-run-sass", "-i", "./web/wireframe.scss", "-o", "./web/generated.css")
```