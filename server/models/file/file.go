package file

import (
	"github.com/tokkenno/media-server/server/models/user"
	"time"
)

type File struct {
	Path      string
	Size      uint64
	Hash      string // SHA256. To detect file changes and errors
	CreatedAt time.Time
	CreatedBy *user.User
}
