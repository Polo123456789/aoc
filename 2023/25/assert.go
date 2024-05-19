package main

import "log"

func assert(condition bool, msg interface{}) {
	if !condition {
		log.Fatal(msg)
	}
}
