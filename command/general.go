package command

import (
  "os"

  "github.com/johnsudaar/gowerewolfgo/config"
  "github.com/johnsudaar/gowerewolfgo/game"

  "github.com/thoj/go-ircevent"
)

func Start(event *irc.Event, ircobj *irc.Connection){
  err := game.StartGame(event.Arguments[0], event.Nick)
  if err != nil {
    ircobj.Privmsg(event.Arguments[0], err.Error())
  } else {
    ircobj.Privmsg(event.Arguments[0], "Nouvelle partie lancée par : "+event.Nick)
    ircobj.Privmsg(event.Arguments[0], "Tapez "+config.E["PREFIX"]+" join pour rejoindre")
    if err := game.RegisterUser(event.Nick); err != nil {
      ircobj.Privmsg(event.Arguments[0], "J'ai été codé par un débile. Allez j'me casse")
      os.Exit(0)
    }
  }
}

func Register(event *irc.Event, ircobj *irc.Connection){
  if game.HasStarted() && event.Arguments[0] != game.Channel() {
    ircobj.Privmsg(event.Arguments[0], "La partie n'est pas lancée dans ce channel. RDV sur "+game.Channel())
  } else {
    if err := game.RegisterUser(event.Nick); err != nil {
      ircobj.Privmsg(event.Arguments[0],err.Error())
    } else {
      ircobj.Privmsg(game.Channel(), event.Nick+" a rejoint la partie.")
    }
  }
}

func Launch(event *irc.Event, ircobj * irc.Connection){
  if err := game.Launch() ; err != nil {
    ircobj.Privmsg(event.Arguments[0], err.Error())
  }
}

func List(event *irc.Event, ircobj * irc.Connection){
  if game.Phase() == game.NOT_LAUNCHED {
    ircobj.Privmsg(event.Arguments[0], "La partie n'a pas commencée")
  } else {
    ircobj.Privmsg(event.Arguments[0], "Liste des joueurs : ")
    for _, user := range game.Users(){
      ircobj.Privmsg(event.Arguments[0], " - "+user.Nick)
    }
  }
}


func Vote(event *irc.Event, ircobj *irc.Connection) {
  game.Vote(event)
}
