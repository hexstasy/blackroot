package main

import (
  "github.com/hexstasy/blackroot/internal/cli"
  "github.com/hexstasy/blackroot/internal/config"
  "github.com/hexstasy/blackroot/internal/errs"
  "github.com/hexstasy/blackroot/internal/parser"
)

func main() {
  cfg, err := config.Load()
  if err != nil {
    errs.Must(err.Error())
  }

  if err := cli.ParseOpts(cfg); err != nil {
    errs.Must(err.Error())
  }

  if err := parseblackroot.ParseAndExecute(cfg.command); err != nil {
    errs.Must(err.Error())
  }
}
