package main

type newspaper struct {
	publication
}

func createNewsPaper(title string, pages int, publisher string) iPublication {
	return &newspaper{
		publication: publication{
			title:     title,
			pages:     pages,
			publisher: publisher,
		},
	}
}

// setPages have different implement across all underlying concrete types
func (p *newspaper) setPages(pages int) {
	p.pages = pages
}
