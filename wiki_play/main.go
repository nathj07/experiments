package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/nathj07/go-wikiparse"
)

func main() {
	p, err := wikiparse.NewParser(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error setting up parser", err)
		os.Exit(1)
	}

	for err == nil {
		var page *wikiparse.Page
		page, err = p.Next()
		if err == nil {
			fmt.Println(page.Title)
			for i, rev := range page.Revisions {
				body := []byte(rev.Text)
				err := ioutil.WriteFile(fmt.Sprintf("/tmp/wiki_out/%d.ml", rev.ID), body, 0644)
				if err != nil {
					fmt.Printf("Error on page %d: %v\n", rev.ID, err)
				}
			}

		}
	}
}
