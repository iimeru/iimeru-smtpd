package main

import (
	"github.com/iimeru/go-guerrilla"
)

func main() {
	goguerrilla.Run(gConfig, saveMail)
}

func saveMail(saveMailChan chan *goguerrilla.Client) {
}
