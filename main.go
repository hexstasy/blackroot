package main

import (
  "github.com/hexstasy/blackroot/internal/cli"
  "github.com/hexstasy/blackroot/internal/config"
  "github.com/hexstasy/blackroot/internal/errs"
)

func main() {
  cfg, err := config.MustLoad()
  if err != nil {
    errs.Must(err.Error())
  }

  if err := cli.Parse(cfg); err != nil {
    errs.Must(err.Error())
  }
}
