package main

import (
  "github.com/hexstasy/blackroot/cli"
  "github.com/hexstasy/blackroot/config"
  "github.com/hexstasy/blackroot/proc_err"
)

func main() {
  cfg, err := config.Load()
  if err != nil {
    procerr.Err(err.Error())
  }
  cli.ParseOpts(&cfg)
}
