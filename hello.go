// hello.go

package main

import (
    "fmt"
    "log"
    "Exper_3/greetings"
)

func main() {
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)
    
    // Request a greeting message.
    message, err := greetings.Hello("Sam")
    // If an error was returned, print it to the console and
    // exit the program.
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(message)
}