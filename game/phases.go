package game

import(
  "log"
  "strconv"

  "github.com/johnsudaar/gowerewolfgo/config"
  "github.com/johnsudaar/gowerewolfgo/user"
  "github.com/johnsudaar/gowerewolfgo/voting"
)

const (
  NOT_LAUNCHED = iota
  REGISTER
  START_OF_NIGHT
  SEER
  WEREWOLF
  WITCH
  DAY
)

const (
  KILL_POTION_PHASE_1 = iota
  KILL_POTION_PHASE_2
  SAVE_POTION
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
  var werewolfsString []string
  var usersString []string

  for _, u := range users {
    if u.Type == user.Werewolf{
      werewolfsString = append(werewolfsString, u.Nick)
    } else {
      usersString = append(usersString, u.Nick)
    }
  }

  currentVote = voting.NewVote(usersString, werewolfsString)
  for _, user := range werewolfs {
    config.IRC.Privmsg(user.Nick, "Tapez vote <NOM DE LA PERSONNE> pour voter pour cette personne.")
    config.IRC.Privmsg(user.Nick, "Tout ce que vous tapez ici sera partagé avec les autres loup garou.")
  }
}

func Witch(){
  witch := user.GetMembersOf(users, user.Witch)
  if len(witch) == 0 {
    phase = DAY
    NextPhase()
  } else {
    if werewolfKill == witch[0].Nick {
      phase = DAY
      NextPhase()
    }else if killPotionUsed && savePotionUsed {
      phase = DAY
      NextPhase()
    } else {
      config.IRC.Privmsg(channel, "La sorciere se reveille.")
      witchPhase = SAVE_POTION
      NextWitchPhase()
    }
  }
}

func NextWitchPhase(){
  witch := user.GetMembersOf(users, user.Witch)[0]
  if witchPhase == SAVE_POTION {
    if savePotionUsed {
      witchPhase = KILL_POTION_PHASE_1
      NextWitchPhase()
    } else {
      config.IRC.Privmsg(witch.Nick, "Les loup garou ont tué "+werewolfKill)
      config.IRC.Privmsg(witch.Nick, "Voulez vous utiliser votre potion pour le sauver ?")
    }
  } else if witchPhase == KILL_POTION_PHASE_1 {
    if killPotionUsed {
      config.IRC.Privmsg(channel, "La sorcière se rendort.")
      witchPhase = -1
      phase = DAY
      NextPhase()
    } else {
      config.IRC.Privmsg(witch.Nick, "Voulez vous tuer une personne ?")
    }
  } else if witchPhase == KILL_POTION_PHASE_2 {
    config.IRC.Privmsg(witch.Nick, "Qui voulez vous tuer ?")
  }
}

func Day(){
  config.IRC.Privmsg(channel, "Le village se reveil.")
  if werewolfKill == "" && witchKill == "" {
    config.IRC.Privmsg(channel, "Cette nuit, personne n'est mort.")
  }
  if werewolfKill != "" {
    dead := user.GetUser(users, werewolfKill)
    config.IRC.Privmsg(channel, werewolfKill+" est mort. C'était un " + user.TypeString(dead.Type))
    DeleteUser(werewolfKill)
  }
  if witchKill != "" {
    dead := user.GetUser(users, witchKill)
    config.IRC.Privmsg(channel, witchKill+" est mort. C'était un " + user.TypeString(dead.Type))
    DeleteUser(witchKill)
  }

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

  if cont{
    var usersString []string
    for _, u := range users {
      usersString = append(usersString, u.Nick)
    }
    currentVote = voting.NewVote(usersString, usersString)

    config.IRC.Privmsg(channel, "Le village procede aux vote.")
    config.IRC.Privmsg(channel, "Pour voter tapez "+config.E["PREFIX"]+" vote <NOM>")
  } else {
    config.IRC.Privmsg(channel, "Fin de la pertie !")
    phase = NOT_LAUNCHED
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
  case WITCH:
    Witch()
  case DAY:
    Day()
  }
}
