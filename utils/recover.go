package utils

import (
	log "github.com/ReshiAdavan/CryptGO/utils/glog"
)

func CryptGORecover() {
	if r := recover(); r != nil {
		log.Debug().Interface("recover", r).Msg("")
	}
}
