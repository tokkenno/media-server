package user

import (
	"net"
	"time"
)

type SessionAccess struct {
	Address   net.IP
	Time      time.Time
	UserAgent string
}

type Session struct {
	User   *User
	Token  string
	Expire time.Time
	Access []SessionAccess
}
