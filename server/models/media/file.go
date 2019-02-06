package media

import (
	"github.com/tokkenno/media-server/server/models/file"
	"github.com/tokkenno/media-server/server/models/track"
)

type File struct {
	File   *file.File
	Tracks []track.Track
}
