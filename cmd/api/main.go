package main

import (
	"log"

	"github.com/Genexis-6/social/internal/store"
)



func main(){
	newStorage := store.NewStorage(nil)
	config := LoadConfig()
	app := &application{
		config: Config{
			addr: config.addr,
			store: *newStorage,
		},
	}
	
	log.Fatal(app.runApp(app.mount()))

}