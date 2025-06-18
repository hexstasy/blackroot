package cli

import (
  "flag"
  "fmt"
  "os"
  "github.com/hexstasy/blackroot/config"
  "github.com/fatih/color"
)

func Parse(cfg *config.Config) error {
  version := flag.Bool("v", false, "Display version and exit.")
  flag.Parse();
  
  args := os.Args
  if *version {
    fmt.Printf("Blackroot version: %s\n", color.GreenString(cfg.Version)) 
    return nil
  }
  
  if len(args) > 1 {
    cfg.Command = args[len(args)-1]
    return nil
  }

  return fmt.Errorf("no command line arguments specified!")
}
