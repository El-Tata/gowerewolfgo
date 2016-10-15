package game

import (
  "log"
  "strconv"
  "strings"

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
  case WITCH:
    WitchMP(event)
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

    SendMPTo(werewolfs, event.Nick+" : "+event.Message())

    if strings.HasPrefix(event.Message(), "vote "){
      nickname := strings.TrimPrefix(event.Message(), "vote ")
      if err := currentVote.Vote(event.Nick, nickname); err != nil {
        config.IRC.Privmsg(event.Nick, err.Error())
      } else {
        currentVotes := currentVote.CountVote()
        SendMPTo(werewolfs, "Etat des votes : ")
        for perso, count := range currentVotes {
          SendMPTo(werewolfs, " - "+perso+" : "+strconv.Itoa(count))
        }

        winned, winner := currentVote.Winner()
        if winned {
          werewolfKill = winner
          config.IRC.Privmsg(channel, "Les loup se rendorment.")
          phase = WITCH
          NextPhase()
        }
      }
    }
  }
}

func WitchMP(event *irc.Event){
  current_user := user.GetUser(users, event.Nick)

  if current_user == nil || current_user.Type != user.Witch {
    config.IRC.Privmsg(event.Nick, "Désolé, je ne peux pas te parler maintenant.")
  } else {
    if witchPhase == SAVE_POTION {
      if event.Message() == "oui" {
        config.IRC.Privmsg(event.Nick, "Vous avez sauvé "+ werewolfKill)
        werewolfKill = ""
        savePotionUsed = true
        witchPhase = KILL_POTION_PHASE_1
        NextWitchPhase()
      } else if event.Message() == "non" {
        config.IRC.Privmsg(event.Nick, "Vous n'avez pas sauvé "+ werewolfKill)
        witchPhase = KILL_POTION_PHASE_1
        NextWitchPhase()
      } else {
        config.IRC.Privmsg(event.Nick, "Je n'ai pas compris merci de répondre par oui ou non")
      }
    } else if witchPhase == KILL_POTION_PHASE_1 {
      if event.Message() == "oui" {
        witchPhase = KILL_POTION_PHASE_2
        NextWitchPhase()
      } else if event.Message() == "non" {
        config.IRC.Privmsg(event.Nick, "Vous n'avez tué personne.")
        witchPhase = -1
        phase = DAY
        config.IRC.Privmsg(channel, "La sorcière se rendort.")
        NextPhase()
      } else {
        config.IRC.Privmsg(event.Nick, "Je n'ai pas compris merci de répondre par oui ou non")
      }
    } else if witchPhase == KILL_POTION_PHASE_2 {
      message := event.Message()
      u := user.GetUser(users, message)

      if u == nil {
        config.IRC.Privmsg(event.Nick, "Je ne connais pas cette personne")
      } else {
        config.IRC.Privmsg(event.Nick, "Vous avez tué "+u.Nick)
        phase = DAY
        witchPhase = -1
        killPotionUsed = true
        witchKill = u.Nick
        config.IRC.Privmsg(channel, "La sorcière se rendort.")
        NextPhase()
      }
    }
  }
}

func SendMPTo(users []*user.User, message string) {
  for _, u := range users {
    config.IRC.Privmsg(u.Nick, message)
  }
}
