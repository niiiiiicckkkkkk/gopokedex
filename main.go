package main

import (
    
    "fmt"
    "bufio"
    "os"
    "errors"
)

type Command struct {
    name string 
    description string
    callback func() error 
}

var commands = map[string]Command{
    "help" : Command{
        name : "help",
        description : "describe pokedex behavior",
        callback : func() error { 
            fmt.Println("help command")  
            return nil
        },
    },
    "exit" : Command {
        name : "exit",
        description : "exit the pokedex",
        callback : func() error { 
            return errors.New("exit gracefully")
        },
    },
}


func prompt() {
    fmt.Printf("pokedex > ")
}


func main() {
    
    scanner := bufio.NewScanner(os.Stdin)
    for prompt(); scanner.Scan(); {
        err := scanner.Err()
        input := scanner.Text()

        if err != nil {
            break
        }
        
        if cmd, ok := commands[input]; ok {
            fmt.Println("valid command")
            stat := cmd.callback()
            if stat != nil {
                break
            }
        } else {
            fmt.Println("invalid command")
        }
        prompt()
    }
    fmt.Printf("\n")
}
