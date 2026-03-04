package main

import (
	"log"
)



func main(){
	
	config := LoadConfig()
	app := &application{
		config: *config,
	}
	
	log.Fatal(app.runApp(app.mount()))

}