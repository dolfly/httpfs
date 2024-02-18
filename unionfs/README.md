# unionfs

Go lang union http.FileSystem

## Installation

```sh
$ go get github.com/dolfly/httpfs/unionfs
```

## Usage

```go
import "github.com/dolfly/httpfs/unionfs"
```

## Examples

### UnionFS

```go
package main

import (
	"github.com/dolfly/httpfs/unionfs"
	"net/http"
)

func main() {
	fs := unionfs.New(http.Dir("A"), http.Dir("B"))

	http.Handle("/", http.FileServer(fs))
	http.ListenAndServe(":8080", nil)
}

```

Simple union file system webserver

## Author

[gnue](https://github.com/dolfly)

## License

[MIT](LICENSE.txt)

