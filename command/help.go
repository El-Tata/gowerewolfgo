package command

import(
  "log"
  "strings"

  "github.com/thoj/go-ircevent"

)


func Help(event *irc.Event, ircobj *irc.Connection) {
  log.Println("Showing help in "+event.Arguments[0]+" requested by "+event.Nick)
  list :=  ListCommands()
  for _, line := range strings.Split(list, "\n") {
    ircobj.Privmsg(event.Arguments[0], line)
  }
}
