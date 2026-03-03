package main

import (
	"log"


)



func main(){
	config := LoadConfig()
	app := &application{
		config: Config{
			addr: config.addr,
		},
	}
	
	log.Fatal(app.runApp(app.mount()))

}