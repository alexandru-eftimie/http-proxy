package main

import "log"

func debug(level int, p ...interface{}) {
	if level <= Co.DebugLevel {
		log.Println(p...)
	}
}
func debugf(level int, format string, p ...interface{}) {
	if level <= Co.DebugLevel {
		log.Printf(format, p...)
	}
}
