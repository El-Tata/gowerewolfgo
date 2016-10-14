package command

import(
  "regexp"

  "github.com/thoj/go-ircevent"
)

var commands = map[string]*Command{
  "hello" : &Command{
    Pattern: regexp.MustCompile(`hello`),
    Description: "Say hello",
    UsePrefix: false,
    Function: func(event *irc.Event, ircobj *irc.Connection){
      ircobj.Privmsg(event.Arguments[0], "Hello "+event.Nick)
    },
  },

  "up" : &Command{
    Pattern: regexp.MustCompile(`^ +up`),
    Description: "Yup",
    UsePrefix: true,
    Function: func(event *irc.Event, ircobj *irc.Connection){
      ircobj.Privmsg(event.Arguments[0], "Yup")
    },
  },
}

func init(){
  commands["help"] = &Command{
    Pattern: regexp.MustCompile(`^ +help`),
    Description: "Show help",
    UsePrefix: true,
    Function: Help,
  }
}
