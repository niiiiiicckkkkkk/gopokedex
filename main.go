package main

import (
    
    "fmt"
    "bufio"
    "os"
)

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
        
        if cmd, ok := getCommands()[input]; ok {
            stat := cmd.callback()
            if stat != nil {
                break
            }
        } else {
            fmt.Println("invalid command enter help for more info")
        }
        prompt()
    }
}
