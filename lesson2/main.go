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

	// Hard-code ID 2 here to test the query.
	alb, err := albumByID(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Album found: %v\n", alb)

	albID, err := addAlbum(Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added album: %v\n", albID)
}
