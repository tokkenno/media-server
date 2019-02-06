package media

type ImageType string

const (
	Poster ImageType = "poster"
)

type Image struct {
	Type ImageType
	File *File
}
