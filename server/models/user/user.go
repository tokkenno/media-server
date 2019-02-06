package user

import (
	"github.com/tokkenno/media-server/server/models/file"
	"github.com/tokkenno/media-server/server/models/library"
)

type Permission string

type User struct {
	Username       string
	password       string
	Email          string
	Image          *file.File         // Profile image
	Permissions    []Permission       // User permissions
	Notifications  []Notification     // Notification configuration
	Language       string             // System language
	MediaLanguages []MediaLanguage    // Rules for media track selection based in language
	Libraries      []*library.Library // List of subscribed libraries
}
