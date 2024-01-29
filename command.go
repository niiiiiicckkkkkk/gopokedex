package main

import (
    "fmt"
    "os"
)

type Command struct {
    name string 
    description string
    documentation string
    callback func() error 
}


func cmdHelp() error {
    for name, cmd := range(getCommands()) {
        fmt.Println(name, ":", cmd.description)
        fmt.Println(cmd.documentation)
    }
    return nil
}

func cmdExit() error {
    os.Exit(0)
    return nil
}

func getCommands() map[string]Command {
    var commands = map[string]Command{
        "help" : Command{
            name : "help",
            description : "list available pokedex commands",
            documentation : "",
            callback : cmdHelp,
        },
        "exit" : Command {
            name : "exit",
            description : "exit the pokedex",
            documentation : "",
            callback : cmdExit,
        },
    }
    return commands
}
