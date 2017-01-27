package main

import (
	"flag"

	log "github.com/Sirupsen/logrus"

	"aesir/formatter"
	"aesir/server"
	"aesir/utils"
)

func main() {
	debug := flag.Bool("debug", false, "Display verbose debug output")
	flag.Parse()
	if *version {
		printVersion()
	}

	if *debug {
		log.SetLevel(log.DebugLevel)
		log.Debug("Debug Enabled")
	} else {
		log.SetLevel(log.InfoLevel)
	}

	log.SetFormatter(&formatter.SimpleFormatter{})

	utils.Splash()

	log.Info("Aesir is starting")

	/*
	  mongo := mongo.Connect("aesir", "aesir", conf.MongoPassword)
	*/

	server.RunServer(conf, mongo, producer, nflex, babel)

	log.Info("The council has now finished.")
	log.Exit(0)
}
