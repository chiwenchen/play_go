package common

import "time"

type Magazine struct {
	Name        string
	Price       int
	PublishedAt time.Time
}

func (m Magazine) GetName() string {
	return m.Name
}
