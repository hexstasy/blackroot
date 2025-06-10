package cli

import (
  "flag"
  "fmt"
  "github.com/hexstasy/blackroot/config"
  "github.com/fatih/color"
)

func ParseOpts(cfg *config.Config) {
  display_version := flag.Bool("version", false, "Display version and exit.")
  flag.Parse();
  if *display_version {
    fmt.Printf("Blackroot version: %s\n", color.GreenString(cfg.Settings.Version)) 
  }
}
