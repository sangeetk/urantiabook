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
	Index, err = bleve.NewMemOnly(mapping)
	if err != nil {
		log.Fatal(err)
	}

	// Index Papers
	for pn, paper := range UBPapers {
		// Index Sections
		for sn, section := range paper.Sections {
			if sn == 0 {
				section.Title = paper.Title
			}
			// Index section
			id := fmt.Sprintf("UB[%d:%d]", pn, sn)
			err = Index.Index(id, section)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

}
