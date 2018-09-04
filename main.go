package main

import (
	"os"
	"github.com/urfave/cli"
	"github.com/fodoctor/fibosbp_security_startup/log"
	"github.com/fodoctor/fibosbp_security_startup/action"
)

func main() {

	app := cli.NewApp()
	app.Name = "FibOS BP Security Program"
	app.Description = `
	
		

	`
	app.Version = "v0.0.1"

	app.Flags = []cli.Flag {}

	//bp start action
	app.Action = action.BpStartAction

	err := app.Run(os.Args)
	if err != nil {
		log.Error(err.Error())
		return
	}

}
