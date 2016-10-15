package game

import(
  "errors"
  "strconv"
  "strings"

  "github.com/johnsudaar/gowerewolfgo/config"
  "github.com/johnsudaar/gowerewolfgo/user"
  "github.com/johnsudaar/gowerewolfgo/voting"

  "github.com/thoj/go-ircevent"
)

var phase int = NOT_LAUNCHED
var channel string
var users []*user.User
var launcher string
var turn int
var currentVote *voting.Vote
var werewolfKill string
var witchKill string
var killPotionUsed bool
var savePotionUsed bool
var witchPhase int

func HasStarted() bool {
  return phase == NOT_LAUNCHED
}

func Channel() string {
  return channel
}

func Users() []*user.User{
  return users
}

func Phase() int{
  return phase
}

func StartGame(c string, u string) error {
  if phase == NOT_LAUNCHED {
    phase = REGISTER
    launcher = u
    channel = c
    users = []*user.User {}
    turn = 1
    werewolfKill = ""
    currentVote = nil
    killPotionUsed = false
    savePotionUsed = false
    witchPhase = -1
    return nil
  } else {
    return errors.New("La partie à déjà été lancée par "+launcher)
  }
}

func RegisterUser(u string) error{
  if phase == REGISTER {
    for _, user := range users {
      if user.Nick == u {
        return errors.New(user.Nick + " : Tu es déjà enregistré")
      }
    }
    users = append(users,&user.User{
      Nick: u,
      Type: user.None,
    })
    return nil
  } else {
    return errors.New("Ce n'est pas le moment de s'enregistrer")
  }
}

func Launch() error{
  if phase == REGISTER {
    if(UserCount() < 5) {
      return errors.New("Impossible de lancer la partie avec moins de 5 utilisateurs")
    } else {
      user.DistributeTypes(users)
      phase = START_OF_NIGHT
      for _, user := range users {
        user.SendType(config.IRC)
      }
      NextPhase()
      return nil
    }
  } else {
    return errors.New("Nous ne sommes pas en phase d'inscription")
  }
}

func HasWerewolfWin() bool{
  werewolfs := user.GetMembersOf(users, user.Werewolf)
  return len(users) == len(werewolfs)
}

func HasVillagersWin() bool{
  werewolfs := user.GetMembersOf(users, user.Werewolf)
  return len(werewolfs) == 0
}

func UserCount() int{
  return len(users)
}

func DeleteUser(u string){
  for i := 0; i < len(users); i++ {
    if u == users[i].Nick {
      users = append(users[:i], users[i+1:]...)
      return
    }
  }
}

func Vote(event *irc.Event) {
  if phase == DAY {
    message := strings.TrimPrefix(event.Message(), config.E["PREFIX"] + " vote ")
    current_user := user.GetUser(users, message)
    if current_user == nil {
      config.IRC.Privmsg(channel, "Je ne connais pas cette personne")
    } else {
      if err := currentVote.Vote(event.Nick, current_user.Nick); err != nil {
        config.IRC.Privmsg(channel, err.Error())
      } else {
        currentVotes := currentVote.CountVote()
        config.IRC.Privmsg(channel, "Etat des votes : ")
        for perso, count := range currentVotes {
          config.IRC.Privmsg(channel, " - "+perso+ " : "+strconv.Itoa(count))
        }

        winned, winner := currentVote.Winner()

        if winned {
          dead := user.GetUser(users, winner)
          config.IRC.Privmsg(channel, winner+" est mort. C'était un " + user.TypeString(dead.Type))
          DeleteUser(winner)
          cont := true

          if HasWerewolfWin() && !HasVillagersWin() {
            config.IRC.Privmsg(channel, "Les loup garou ont gagné !")
            cont = false
          } else if HasVillagersWin() && !HasWerewolfWin() {
            config.IRC.Privmsg(channel, "Les villageois ont gagné !")
            cont = false
          } else if len(users) == 0 {
            config.IRC.Privmsg(channel, "Tout le monde est mort !")
            cont = false
          }
          if cont {
            turn ++
            phase = START_OF_NIGHT
            NextPhase()
          } else {
            config.IRC.Privmsg(channel, "Fin de la partie !")
            phase = NOT_LAUNCHED
          }
        }
      }
    }
  } else {
    config.IRC.Privmsg(event.Arguments[0], "Vous ne pouvez pas voter maintenant.")
  }
}
