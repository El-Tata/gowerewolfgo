package user

import(
  "math"
  "math/rand"
)

const (
  None = iota
  Villager
  Werewolf
  Seer
  Hunter
  Witch
)


func TypeString(a int) string {
  switch a{
  case None:
    return "Aucun"
  case Villager:
    return "Villageois"
  case Werewolf:
    return "Loup Garou"
  case Seer:
    return "Voyante"
  case Hunter:
    return "Chasseur"
  case Witch:
    return "Sorciere"
  }
  return "BUUUUGGGGG"
}

func DistributeTypes(users []*User){
  count := len(users)

  werewolfCount := math.Ceil(float64(count) / 3)
  for i := 0.0; i < werewolfCount; i++ {
    SetType(users, None, Werewolf)
  }

  for _, user := range users {
    if user.Type == None {
      user.Type = Villager
    }
  }

  if count >= 5 {
    SetType(users, Villager, Seer)
  }

  if count >= 7 {
    SetType(users, Villager, Witch)
  }

  if count >= 9 {
    SetType(users, Villager, Hunter)
  }
}

func SetType(users []*User, from int, to int){
  a := rand.Int() % len(users)
  for ; users[a].Type != from; {
    a = rand.Int() % len(users)
  }
  users[a].Type = to
}
