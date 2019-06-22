package service

import (
	"log"
	"os"

	"github.com/blevesearch/bleve"
)

var Index bleve.Index

type SearchItem struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

func init() {
	var err error
	// mapping.DefaultMapping = documentMapping
	mapping := bleve.NewIndexMapping()

	if os.Getenv("STORAGE") == "MEMORY" {
		Index, err = bleve.NewMemOnly(mapping)
		createIndex()
	} else {
		Index, err = bleve.Open("db/urantiabook.idx")
		if err != nil {
			if Index, err = bleve.New("db/urantiabook.idx", mapping); err == nil {
				createIndex()
			}
		} else {
			if os.Getenv("REINDEX") == "TRUE" || os.Getenv("REINDEX") == "YES" {
				createIndex()
			}
		}
	}
	if err != nil {
		log.Fatal(err)
	}

}

func createIndex() {
	// Index Papers
	for _, paper := range UBPapers {
		log.Printf("Indexing: [%s] %s - [by %s]\n", paper.ID, paper.Title, paper.Author)
		Index.Index(paper.ID, &SearchItem{paper.ID, paper.Title})
		// Sections
		for _, section := range paper.Sections {
			// fmt.Printf("\t[%s] %s\n", section.ID, section.Title)
			if section.Title != "" {
				Index.Index(section.ID, &SearchItem{section.ID, section.Title})
			}
			// Paragraphs
			for _, para := range section.Paragraphs {
				// fmt.Printf("\t\t[%s]\n", para.ID)
				Index.Index(para.ID, &SearchItem{para.ID, para.Text})
			}
		}
	}
}
