package user

import (
  "github.com/thoj/go-ircevent"
)
type User struct {
  Nick string
  Type int
}

func (u *User) SendType(ircobj *irc.Connection) {
  ircobj.Notice(u.Nick, "Vous etes un "+TypeString(u.Type))
}
