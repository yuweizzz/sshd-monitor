package main

import (
	"log"
	"time"

	"github.com/BurntSushi/toml"
)

type Configs struct {
	PamLib	string
	NetIface       string
	MaxFailedCount int
	TimeUnit       time.Duration
	ReportPeriod   time.Duration
	BuildDictOnly  bool
	Users          []string
}

func ReadConfig(path string) (cfg *Configs) {
	if _, err := toml.DecodeFile(path, &cfg); err != nil {
		panic(err)
	}
	cfg.TimeUnit = cfg.TimeUnit * time.Second
	cfg.ReportPeriod = cfg.ReportPeriod * time.Second
	log.Println("Using Config:")
	log.Printf("PamLib: %s\n", cfg.PamLib)
	log.Printf("NetIface: %s\n", cfg.NetIface)
	log.Printf("MaxFailedCount: %d\n", cfg.MaxFailedCount)
	log.Printf("TimeUnit: %s\n", cfg.TimeUnit)
	log.Printf("ReportPeriod: %s\n", cfg.ReportPeriod)
	log.Printf("BuildDictOnly: %v\n", cfg.BuildDictOnly)
	log.Printf("Users: %s\n", cfg.Users)
	return
}
