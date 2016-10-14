package game

import(
  "log"
  "strconv"

  "github.com/johnsudaar/gowerewolfgo/config"
  "github.com/johnsudaar/gowerewolfgo/user"
)

const (
  NOT_LAUNCHED = iota
  REGISTER
  START_OF_NIGHT
  SEER
  WEREWOLF
)

func StartOfNight(){
  config.IRC.Privmsg(channel, "La nuit tombe. Tout le monde s'endort.")
  phase = SEER
  NextPhase();
}

func Seer(){
  seer := user.GetMembersOf(users, user.Seer)
  if len(seer) == 0 {
    phase = WEREWOLF
    NextPhase();
  } else {
    config.IRC.Privmsg(channel, "La voyante se reveille et design la carte qu'elle veut voir.")
    config.IRC.Privmsg(seer[0].Nick, "Vous êtes la voyante. Entrez le nom du joueur qui vous interesse.")
  }
}

func Werewolf(){
  config.IRC.Privmsg(channel, "Les loup garous se reveillent et vont choisir une personne a tuer.")
  werewolfs := user.GetMembersOf(users, user.Werewolf)
  for _, user := range werewolfs {
    config.IRC.Privmsg(user.Nick, "Tapez vote <NOM DE LA PERSONNE> pour voter pour cette personne.")
    config.IRC.Privmsg(user.Nick, "Tout ce que vous tapez ici sera partagé avec les autres loup garou.")
  }
}

func NextPhase() {
  log.Println("Nouvelle phase: "+strconv.Itoa(phase))
  switch phase {
  case START_OF_NIGHT:
    StartOfNight()
  case SEER:
    Seer()
  case WEREWOLF:
    Werewolf()
  }
}
