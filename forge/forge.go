package main

import (
	"flag"

	log "github.com/Sirupsen/logrus"
	"github.com/twosevenska/aesir/forge/server"
	"github.com/twosevenska/aesir/splashes"
)

func main() {
	debug := flag.Bool("debug", false, "Display verbose debug output")

	flag.Parse()

	if *debug {
		log.SetLevel(log.DebugLevel)
		log.Debug("Debug Enabled")
	} else {
		log.SetLevel(log.InfoLevel)
	}

	splashes.Splash("forge")

	log.Info("The flames are now roaring in the forge")

	server.RunServer()

	log.Info("The flames have been extinguished.")
	log.Exit(0)
}
