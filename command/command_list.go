package command

import(
  "regexp"
)

var commands = map[string]*Command{
  "start" : &Command{
    Pattern: regexp.MustCompile(`^ +start$`),
    Description: "Démarre une nouvelle partie de Loup Garou",
    UsePrefix: true,
    ShowInHelp: true,
    Function: Start,
  },

  "join" : &Command{
    Pattern: regexp.MustCompile(`^ *join$`),
    Description: "Rejoindre une partie en cours",
    UsePrefix: true,
    ShowInHelp: true,
    Function: Register,
  },
  "launch" : &Command{
    Pattern: regexp.MustCompile(`^ +launch$`),
    Description: "Clos la phase d'inscription",
    UsePrefix: true,
    ShowInHelp: true,
    Function: Launch,
  },
  "list" : &Command{
    Pattern: regexp.MustCompile(`^ +list$`),
    Description: "Liste les joueurs dans la partie",
    UsePrefix: true,
    ShowInHelp: true,
    Function: List,
  },
  "vote" : &Command{
    Pattern: regexp.MustCompile(`^ +vote`),
    Description: "Permet de voter pour un joueur",
    UsePrefix: true,
    ShowInHelp: true,
    Function: Vote,
  },
}

func init(){
  commands["help"] = &Command{
    Pattern: regexp.MustCompile(`^ +help`),
    Description: "Show help",
    UsePrefix: true,
    ShowInHelp: true,
    Function: Help,
  }
}
