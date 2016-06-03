package main

import (
    "bytes"
    "encoding/gob"
    "fmt"
)
type Entry struct {
    Key string
    Val string
}

func main() {
    var buf bytes.Buffer
    enc := gob.NewEncoder(&buf)
    e := Entry { "k1", "v1" }
    enc.Encode(e)
    fmt.Println(buf.Bytes())
}
