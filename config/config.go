package config

import (
  "github.com/ilyakaznacheev/cleanenv"
  "fmt"
)

type Config struct {
  Settings struct {
    Version string `yml:"version"`
  } `yml:"settings"`
}

const CFG_FILE_NAME = "config.yml"

func Load() (Config, error) {
  var cfg Config;
  if err := cleanenv.ReadConfig(CFG_FILE_NAME, &cfg); err != nil {
    return Config{}, fmt.Errorf("failed read %s file", CFG_FILE_NAME)
  }
  
  return cfg, nil
}
