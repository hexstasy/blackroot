package cli

import (
  "flag"
  "fmt"
  "os"
  "github.com/hexstasy/blackroot/internal/config"
  "github.com/hexstasy/blackroot/internal/parser"
  "github.com/fatih/color"
  "github.com/hexstasy/blackroot/internal/runner"
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
    command, err := parser.Parser(args[len(args)-1])
    if err != nil {
      return err
    }

    if err := runner.Run(command); err != nil {
      return err
    }

    return nil
  }

  return fmt.Errorf("no command line arguments specified!")
}
