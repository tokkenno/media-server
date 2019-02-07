package media

import (
	"github.com/tokkenno/media-server/server/models/library"
	"github.com/tokkenno/media-server/server/models/metadata"
)

type Media struct {
	Library  *library.Library
	Type     Type
	Public   bool
	Metadata []metadata.Metadata
	Persons  []metadata.Person
	Images   []Image

	Index  int    // Index on parent object (Ex. track number on music album)
	Parent *Media // Parent media object, to nest elements in collections

	Views       uint64 // Number of views
	Original    File   // Original file, if this media has file
	Conversions []File // Pre-converted alternative versions, if this media has file
}
