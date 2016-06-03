package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {
	e := &Entry{
		Key: proto.String("k1"),
		Val: proto.String("v1"),
	}
	data, err := proto.Marshal(e)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	fmt.Println(data)
}
