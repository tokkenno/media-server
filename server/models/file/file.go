package file

import "time"

type File struct {
	Path      string
	Size      uint64
	CreatedAt time.Time
}
