package main

import "fblviewer/command"

func main() {
	err := command.Execute()
	if  err != nil {
		panic("command Execute error" + err.Error())

	}
}
