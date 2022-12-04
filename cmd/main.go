package main

import (
	"github.com/nstoker/stoker.dev/pkg/version"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Infof("Starting up %s", version.Version())
}
