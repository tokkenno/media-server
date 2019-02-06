package user

import (
	"github.com/tokkenno/media-server/server/models/notification"
)

type Notification struct {
	Agent      notification.Agent
	Language   string
	Categories []notification.Category
}
