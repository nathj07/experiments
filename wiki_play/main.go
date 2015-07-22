package main

import (
	"fmt"
	"os"

	"github.com/dustin/go-wikiparse"
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
			for _, rev := range page.Revisions {
				fmt.Println("Revision:", rev.ID)
				fmt.Println("By      :", rev.Contributor)
				fmt.Println("Body:\n\t", rev.Text)
			}
			fmt.Println("=========")
		}
	}
}
