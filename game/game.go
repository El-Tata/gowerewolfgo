package game

import(
  "errors"

  "github.com/johnsudaar/gowerewolfgo/config"
  "github.com/johnsudaar/gowerewolfgo/user"

)

var phase int = NOT_LAUNCHED
var channel string
var users []*user.User
var launcher string
var turn int

func HasStarted() bool {
  return phase == NOT_LAUNCHED
}

func Channel() string {
  return channel
}

func StartGame(c string, u string) error {
  if phase == NOT_LAUNCHED {
    phase = REGISTER
    launcher = u
    channel = c
    users = []*user.User {}
    turn = 1
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

func UserCount() int{
  return len(users)
}
