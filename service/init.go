package service

import (
	"fmt"
	"log"

	"github.com/blevesearch/bleve"
)

var Index bleve.Index

func init() {
	var err error
	// mapping.DefaultMapping = documentMapping
	mapping := bleve.NewIndexMapping()
	Index, err := bleve.NewMemOnly(mapping)
	if err != nil {
		log.Fatal(err)
	}

	// Index Papers
	for n, paper := range UBPapers {
		// Index Sections
		for i, section := range paper.Sections {
			if i == 0 {
				section.Title = paper.Title
			}
			// Index section
			id := fmt.Sprintf("UB[%d:%d]", n, i)
			err = Index.Index(id, section)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

}
