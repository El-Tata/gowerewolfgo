package main

import (
    "log"

    "github.com/thoj/go-ircevent"

    "github.com/johnsudaar/gowerewolfgo/config"
    "github.com/johnsudaar/gowerewolfgo/command"
)

func main() {
    var a chan bool
    ircobj := irc.IRC(config.E["BOT_NAME"], config.E["BOT_NAME"])
    config.IRC = ircobj
    log.Println("Connecting to "+config.E["SERVER"]+" with username "+config.E["BOT_NAME"]+"...")

    ircobj.AddCallback("001", func(event *irc.Event){
        log.Println("Connected.")
        for _, channel := range config.Channels {
            log.Println("Joining "+channel)
            ircobj.Join(channel)
        }
    });

    ircobj.AddCallback("PRIVMSG", func(event *irc.Event){
        go func(event *irc.Event){
            log.Println("New message from "+event.Nick+" ("+event.Arguments[0]+"): "+event.Message())
            if command.RunCommand(event, ircobj) {
                log.Println("Command launched")
            } else {
                log.Println("Unknown command")
            }
        }(event)
    })

    log.Println(command.ListCommands())

    ircobj.Connect(config.E["SERVER"])
    <- a

}
