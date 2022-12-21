package main

import (
	router "github.com/nstoker/stoker.dev/pkg/router"
	"github.com/nstoker/stoker.dev/pkg/version"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Infof("Starting up %s", version.Version())
	_, err := router.GetRouter()
	if err != nil {
		logrus.Fatalf("router init failed %v", err)
	}

	log.Infof(router.Run("127.0.0.1:3000").Error())
}
