package user

import (
  "github.com/thoj/go-ircevent"
)
type User struct {
  Nick string
  Type int
}

func (u *User) SendType(ircobj *irc.Connection) {
  ircobj.Privmsg(u.Nick, "Vous etes un "+TypeString(u.Type))
}

func GetUser(users []*User, nick string)*User{
  for _, user := range users {
    if user.Nick == nick {
      return user
    }
  }
  return nil
}
