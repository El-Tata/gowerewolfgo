package game

import (
  "log"

  "github.com/johnsudaar/gowerewolfgo/config"
  "github.com/johnsudaar/gowerewolfgo/user"

  "github.com/thoj/go-ircevent"
)

func HandleMP(event *irc.Event){

  log.Println("New mp from "+event.Nick+" : "+event.Message())
  switch phase {
  case SEER:
    SeerMP(event)
  case WEREWOLF:
    WerewolfMP(event)
  default:
    config.IRC.Privmsg(event.Nick, "Désolé, je ne peux pas te parler maintenant.")
  }
}

func SeerMP(event *irc.Event){
  seer := user.GetMembersOf(users, user.Seer)

  if event.Nick == seer[0].Nick {
    message := event.Message()
    u := user.GetUser(users, message)

    if u == nil {
      config.IRC.Privmsg(event.Nick, "Je ne connais pas cette personne")
    } else {
      config.IRC.Privmsg(event.Nick, "Cette personne est un "+ user.TypeString(u.Type))
      phase = WEREWOLF
      config.IRC.Privmsg(channel, "La voyante se rendort.")
      NextPhase()
    }
  } else {
    config.IRC.Privmsg(event.Nick, "Désolé, je ne peux pas te parler maintenant.")
  }
}

func WerewolfMP(event *irc.Event) {
  current_user := user.GetUser(users, event.Nick)

  if current_user == nil || current_user.Type != user.Werewolf {
    config.IRC.Privmsg(event.Nick, "Désolé, je ne peux pas te parler maintenant.")
  } else {
    werewolfs := user.GetMembersOf(users, user.Werewolf)
    for _, u := range werewolfs {
      if u.Nick != event.Nick {
        config.IRC.Privmsg(u.Nick, event.Nick+" : "+event.Message())
      }
    }
  }
}
