package metadata

import (
	"net/url"
	"time"
)

type Person struct {
	Name           string
	LastName       string
	PatronymicName string
	Description    string
	Gender         Gender
	Birthday       time.Time
	Links          []url.URL
	Tags           []string
	Relationships  []Relationship
}
