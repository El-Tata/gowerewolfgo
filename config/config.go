package config

import (
  "os"
  "strings"
)

var E = map[string]string{
  "SERVER": "irc.iiens.net:6667",
  "CHANNELS": "#test-ircbot",
  "BOT_NAME": "johnbot",
  "PREFIX": "!johnbot",
}

var Channels []string

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
