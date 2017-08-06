package main

import "github.com/urfave/cli"

var (
	flagCrawlingFrequency = ""
	flagTarget            = ""
	flagMaxCrawlerNumber  = 0
	flagNotifierConfig    = ""
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
		Usage:       "Detect target in json format {<boardName1>:[<keyword1>,<keyword2>], <boardName2>...}",
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
	cli.StringFlag{
		Name:        "notifier",
		Usage:       "Notifier config",
		EnvVar:      "PTT_NOTIFIER_CONFIG",
		Destination: &flagNotifierConfig,
	},
}
