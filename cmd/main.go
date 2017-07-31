package main

import (
"os"
//"context"
"github.com/urfave/cli"
)

func InitApp() *cli.App {
	app := cli.NewApp()
	app.Name = "PttNotifier"
	app.Usage = "Ptt Notifier"

	app.Flags = cmdFlags
	app.Commands = []cli.Command{
		{
			Name:   "detect",
			Usage:  "Start detect.",
			Action: StartDetectTarget,
			Flags:  app.Flags,
		},
	}

	app.Action = func(c *cli.Context) error {
		return StartDetectTarget()
	}
	return app
}

func StartDetectTarget() error {
	print("Start detecting")
	//ctx := context.Background()

	return nil
}


func main() {
	app := InitApp()
	if err := app.Run(os.Args); err != nil {
		println(err)
	}
}
