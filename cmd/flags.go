package main

import "github.com/urfave/cli"

var (
	flagCrawlingFrequency = ""
	flagTarget            = ""
	flagMaxCrawlerNumber  = 0
)

var cmdFlags = []cli.Flag{
	cli.StringFlag{
		Name:        "frequency",
		Usage:       "Frequency to crawling ptt archives",
		EnvVar:      "PTT_NOTIFIER_FREQUENCY",
		Value:       "5",
		Destination: &flagCrawlingFrequency,
	},
	cli.StringFlag{
		Name:        "target",
		Usage:       "Detect target in <board name>-<keyword> format with ';' for multiple pairs",
		EnvVar:      "PTT_NOTIFIER_TARGET",
		Destination: &flagTarget,
	},
	cli.IntFlag{
		Name:        "maxCrawlerNumber",
		Usage:       "Max Crawler number",
		EnvVar:      "PTT_NOTIFIER_MAX_CRAWLER_NUM",
		Value:       5,
		Destination: &flagMaxCrawlerNumber,
	},
}
