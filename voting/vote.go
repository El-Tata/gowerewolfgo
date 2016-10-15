package voting

import(
  "errors"
  "sort"
)

type Vote struct {
  possibleAnswers []string
  possibleVoters  []string
  votes           map[string]string
}

func NewVote(answers []string, voters []string) *Vote{
  sort.Strings(answers)
  sort.Strings(voters)
  return &Vote {
    possibleAnswers : answers,
    possibleVoters  : voters,
    votes : make(map[string]string),
  }
}

func (v *Vote) Vote(voter string, votee string) error {
  if ! Contains(v.possibleVoters, voter) {
    return errors.New("Vous ne pouvez pas voter.")
  }

  if ! Contains(v.possibleAnswers, votee) {
    return errors.New("Vous ne pouvez pas voter pour cette personne.")
  }

  v.votes[voter] = votee
  return nil
}

func (v *Vote) Complete() bool {
  for _, voter := range v.possibleVoters {
    if v.votes[voter] == "" {
      return false
    }
  }
  return true
}

func (v *Vote) IsFinished() bool{
  if ! v.Complete() {
    return false
  }

  votes := v.CountVote()
  max := -1
  single := true

  for _, value := range votes {
    if value > max {
      single = true
      max = value
    } else if value == max {
      single = false
    }
  }

  return single
}

func (v *Vote) Winner() (bool, string) {
  if v.IsFinished(){
    votes := v.CountVote()
    winner := ""
    max := -1
    for voter, value := range votes{
      if value > max {
        max = value
        winner = voter
      }
    }
    return true, winner
  } else {
    return false, ""
  }
}

func (v *Vote) CountVote() map[string]int{
  counter := make(map[string]int)
  for _, v := range v.votes {
    if counter[v] == 0 {
      counter[v] = 1
    } else {
      counter[v] = counter[v] + 1
    }
  }
  return counter
}

func Contains(set []string, v string) bool{
  for _, c := range set {
    if c == v {
      return true
    }
  }

  return false
}
