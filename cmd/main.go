package main

import (
	"os"
	"github.com/urfave/cli"
	"github.com/andychuang/pttnotifier"
	"log"
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
		return StartDetectTarget(c)
	}
	return app
}

func StartDetectTarget(c *cli.Context) error {
	log.Println("Start detecting")

	dc, err := pttnotifier.NewDetectorCenter(flagTarget, flagCrawlingFrequency, flagMaxCrawlerNumber)
	if err != nil {
		log.Printf("New detector center failed. %s", err.Error())
		return err
	}
	if err := dc.Run(); err != nil {
		log.Printf("Error when detecting. %s", err.Error())
		return err
	}

	return nil
}

func main() {
	app := InitApp()
	if err := app.Run(os.Args); err != nil {
		println(err)
	}
}
