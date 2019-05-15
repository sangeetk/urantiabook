package service

import (
	"git.urantiatech.com/urantiabook/urantiabook/api"
)

var UBParts = []api.Part{
	{
		ID:         "1",
		Title:      "The Central and Superuniverses",
		PaperStart: 0,
		PaperEnd:   31,
		Authors:    "Sponsored by a Uversa Corps of Superuniverse Personalities acting by authority of the Orvonton Ancients of Days",
	},
	{
		ID:         "2",
		Title:      "The Local Universe",
		PaperStart: 32,
		PaperEnd:   56,
		Authors:    "Sponsored by a Nebadon Corps of Local Universe Personalities acting by authority of Gabriel of Salvington",
	},
	{
		ID:         "3",
		Title:      "The History of Urantia",
		PaperStart: 57,
		PaperEnd:   119,
		Authors:    "These papers were sponsored by a Corps of Local Universe Personalities acting by authority of Gabriel of Salvington",
	},
	{
		ID:         "4",
		Title:      "The Life and Teachings of Jesus",
		PaperStart: 120,
		PaperEnd:   196,
		Authors:    "This group of papers was sponsored by a commission of twelve of a Melchizedek revelatory director.</p>",
	},
}

var UBPapers = []api.Paper{
	{
		ID:    "0",
		Title: " Foreword",
		Sections: []api.Section{
			{
				ID:    "0:0",
				Title: "",
				Text:  " <p> IN THE MINDS of the mortals of Urantia — that being the name of your world — there exists great confusion.</p>",
			},
			{
				ID:    "0:1",
				Title: " I. Deity and Divinity",
				Text:  " <p> The universe of universes presents phenomena of deity activities on diverse levels of cosmic realities. </p>",
			},
			{
				ID:    "0:2",
				Title: " II. God",
				Text:  " <p> Evolving mortal creatures experience an irresistible urge to symbolize their finite concepts of God. </p>",
			},
			{
				ID:    "0:3",
				Title: " III. The First Source and Center",
				Text:  " <p> Total, infinite reality is existential in seven phases and as seven co-ordinate Absolutes:</p>",
			},
			{
				ID:    "0:4",
				Title: " IV. Universe Reality",
				Text:  " <p> Reality differentially actualizes on diverse universe levels; </p>",
			},
			{
				ID:    "0:5",
				Title: " V. Personality Realities",
				Text:  " <p> Personality is a level of deified reality and ranges from the mortal and midwayer level of the higher mind activation </p>",
			},
		},
		Author: "Divine Counselor",
	},
	{
		ID:    "1",
		Title: "The Universal Father ",
		Sections: []api.Section{
			{
				ID:    "1:0",
				Title: "",
				Text:  " <p> THE Universal Father is the God of all creation, the First Source and Center of all things and beings. </p>",
			},
			{
				ID:    "1:1",
				Title: " 1. The Father’s Name",
				Text:  " <p> Of all the names by which God the Father is known throughout the universes </p>",
			},
			{
				ID:    "1:2",
				Title: " 2. The Reality of God",
				Text:  " <p> God is primal reality in the spirit world; God is the source of truth in the mind spheres. </p>",
			},
			{
				ID:    "1:3",
				Title: " 3. God is a Universal Spirit",
				Text:  " <p> “God is spirit.” He is a universal spiritual presence. The Universal Father is an infinite spiritual reality. </p>",
			},
			{
				ID:    "1:4",
				Title: " 4. The Mystery of God",
				Text:  " <p> The infinity of the perfection of God is such that it eternally constitutes him mystery. </p>",
			},
		},
		Author: "Divine Counselor",
	},
	{
		ID:    "2",
		Title: "The Nature of God",
		Sections: []api.Section{
			{
				ID:    "2:0",
				Title: "",
				Text:  " <p> INASMUCH as man’s highest possible concept of God is embraced within the human idea. </p>",
			},
			{
				ID:    "2:1",
				Title: " 1. The Infinity of God",
				Text:  " <p> “Touching the Infinite, we cannot find him out. The divine footsteps are not known.”</p>",
			},
			{
				ID:    "2:2",
				Title: " 2. The Father’s Eternal Perfection",
				Text:  " <p> Even your olden prophets understood the eternal, never-beginning, never-ending, circular nature of the Universal Father. </p>",
			},
			{
				ID:    "2:3",
				Title: " 3. Justice and Righteousness",
				Text:  " <p> God is righteous; therefore is he just. “The Lord is righteous in all his ways.” </p>",
			},
		},
		Author: "Divine Counselor",
	},
}
