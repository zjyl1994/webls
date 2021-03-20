package config

import "flag"

var (
	DataDir  = "."
	SiteName = "Webls"
	Author   = "Webls"
	Listen   = ":9496"
	Since    = 0
)

const (
	DebugMode      = false
	ReadmeFilename = "README.md"
)

func LoadArgs() {
	flag.StringVar(&Listen, "listen", ":9496", "listen address")
	flag.StringVar(&DataDir, "path", ".", "path to list")
	flag.StringVar(&SiteName, "sitename", "Webls", "name display in web panel")
	flag.StringVar(&Author, "author", "Webls", "copyright author display in web panel")
	flag.IntVar(&Since, "since", 0, "since year display in web panel")
	flag.Parse()
}
