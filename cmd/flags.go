package main

import "github.com/urfave/cli"

var (
	flagCrawlingFrequency = 0
	flagTarget            = ""
)

var cmdFlags = []cli.Flag{
	cli.IntFlag{
		Name:        "frequency",
		Usage:       "Frequency to crawling ptt archives",
		EnvVar:      "PTT_NOTIFIER_FREQUENCY",
		Value:       60,
		Destination: &flagCrawlingFrequency,
	},
	cli.StringFlag{
		Name:        "target",
		Usage:       "Detect target in <board name>-<keyword> format with space for multiple pairs",
		EnvVar:      "PTT_NOTIFIER_TARGET",
		Destination: &flagTarget,
	},
}
