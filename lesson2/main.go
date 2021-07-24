package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {

	getHandle()

	//Query for multiple rows
	name := "John Coltrane"
	albums, err := albumsByArtist(name)
	if err != nil {

		log.Fatal(err)
		err = fmt.Errorf("albumsByArtist %q: %v", name, errors.Unwrap(err))
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)

}
