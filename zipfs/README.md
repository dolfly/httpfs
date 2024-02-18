# zipfs

Go lang zip http.FileSystem

## Installation

```sh
$ go get github.com/dolfly/httpfs/zipfs
```

## Usage

```go
import "github.com/dolfly/httpfs/zipfs"
```

## Examples

### New

```go
package main

import (
	"github.com/dolfly/httpfs/zipfs"
	"log"
	"net/http"
	"os"
)

func main() {
	b, err := os.ReadFile("public.zip")
	if err != nil {
		log.Fatal(err)
	}

	fs, err := zipfs.New(b, &zipfs.Options{Prefix: "public"})
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", http.FileServer(fs))
	http.ListenAndServe(":8080", nil)
}

```

Simple zip webserver(use New)

### Open

```go
package main

import (
	"github.com/dolfly/httpfs/zipfs"
	"log"
	"net/http"
)

func main() {
	fs, err := zipfs.Open("public.zip", &zipfs.Options{Prefix: "public"})
	if err != nil {
		log.Fatal(err)
	}
	defer fs.Close()

	http.Handle("/", http.FileServer(fs))
	http.ListenAndServe(":8080", nil)
}

```

Simple zip webserver(use Open)

## Author

[gnue](https://github.com/dolfly)

## License

[MIT](LICENSE.txt)

