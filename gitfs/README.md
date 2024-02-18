# gitfs

Go lang git http.FileSystem

## Installation

```sh
$ go get github.com/dolfly/httpfs/gitfs
```

## Usage

```go
import "github.com/dolfly/httpfs/gitfs"
```

## Examples

### GitFS

```go
package main

import (
	"github.com/dolfly/httpfs/gitfs"
	"net/http"
)

func main() {
	fs := gitfs.New("sites.gif", "master")

	http.Handle("/", http.FileServer(fs))
	http.ListenAndServe(":8080", nil)
}

```

Simple git webserver

## Author

[gnue](https://github.com/dolfly)

## License

[MIT](LICENSE.txt)

