package main

type magazine struct {
	publication // embad之後，會繼承publication type的value跟func
	frontPerson string
}

func createMagazine(title string, pages int, publisher string) iPublication {
	return &magazine{
		publication: publication{
			title:     title,
			pages:     pages,
			publisher: publisher,
		},
	}
}

// this func only works on magazine
// FrontPerson 封面人物 XD
func (m *magazine) setFrontPerson(name string) {
	m.frontPerson = name
}

// this func only works on magazine
func (m *magazine) getFrontPerson() string {
	return m.frontPerson
}

// setPages have different implement across all underlying concrete types
func (p *magazine) setPages(pages int) {
	p.pages = pages
}
