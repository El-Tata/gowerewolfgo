package command

import(
  "bytes"
  "fmt"
  "log"
  "regexp"
  "strings"
  "text/tabwriter"

  "github.com/thoj/go-ircevent"

  "github.com/johnsudaar/gowerewolfgo/config"

)

type Command struct {
  Pattern *regexp.Regexp
  Description string
  UsePrefix bool
  Function func(*irc.Event, *irc.Connection)
}

func RunCommand(event *irc.Event, ircobj *irc.Connection) bool{

  containPrefix := false
  index := strings.Index(event.Message(), config.E["PREFIX"])
  var messageTrimmed string

  if index != -1 {
    containPrefix = true
    messageTrimmed = event.Message()[index+len(config.E["PREFIX"]):len(event.Message())]
  }

  for commandIndex, command := range commands {
    log.Println("Testing with command: "+commandIndex);
    if (! command.UsePrefix && command.Pattern.MatchString(event.Message() )) ||
      (command.UsePrefix && containPrefix && command.Pattern.MatchString(messageTrimmed)){
      command.Function(event, ircobj)
      return true
    }
  }

  return false
}


func ListCommands() string{
  var b bytes.Buffer

  const padding = 3
  w := tabwriter.NewWriter(&b, 0, 0, padding, ' ', tabwriter.Debug)
  fmt.Fprintln(w, "Command\tDescription\t")
  fmt.Fprintln(w, "\t\t")
  for commandIndex, command := range commands {
    fmt.Fprintln(w, commandIndex+"\t"+ command.Description+"\t")
  }

  w.Flush()
  return b.String()
}
