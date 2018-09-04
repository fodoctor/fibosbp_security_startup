package main

import (
	"os"
	"github.com/urfave/cli"
	"github.com/fodoctor/fibosbp_security_startup/log"
	"io/ioutil"
	"sort"
)

func main() {

	app := cli.NewApp()
	app.Version = "v0.0.1"
	app.Name = "FibOS BP Security Program"
	app.Usage = "An useful program for fibos bps. Code by BlockChain Doctor's fans."

	//god bless you
	//may be you can delete it ,but nobody can help you but The Buddha
	{
		godBytes, err := ioutil.ReadFile("god_bless_you")
		if err != nil {
			log.Error(err.Error())
			return
		}
		app.Description = string(godBytes)
	}

	// cli flags config
	{
		app.Flags = []cli.Flag{
			cli.StringFlag{
				Name:  "bp_start, bs",
				Usage: "Start To Launch bp.js",
			},
		}

		sort.Sort(cli.FlagsByName(app.Flags))
		//TODO bp start action
		//TODO claim bp rewards action
		//TODO get bp account balance
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Error(err.Error())
		return
	}

}
