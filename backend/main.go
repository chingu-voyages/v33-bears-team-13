package main

import (
	"log"
	"fmt"
	"os"

	// Shortening the import reference name seems to make it a bit easier
	owm "github.com/briandowns/openweathermap"
)



func main() {
	w, err := owm.NewCurrent("F", "ru", apiKey) // fahrenheit (imperial) with Russian output
	if err != nil {
		log.Fatalln(err)
	}

	w.CurrentByName("Phoenix")
	fmt.Println(w)
}
