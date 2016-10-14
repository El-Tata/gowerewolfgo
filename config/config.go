package config

import (
  "os"
  "strings"

  "github.com/thoj/go-ircevent"
)

var E = map[string]string{
  "SERVER": "irc.iiens.net:6667",
  "CHANNELS": "#test-ircbot",
  "BOT_NAME": "loupgarou",
  "PREFIX": "!loupgarou",
}

var Channels []string
var IRC *irc.Connection

func init() {
  for key := range E {
    value := os.Getenv(key)
    if value != "" {
      E[key] = value
    }
  }
  initEnv()
}

func initEnv() {
  Channels = strings.Split(E["CHANNELS"], ",")
}
