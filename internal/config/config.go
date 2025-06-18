package config

import (
  "github.com/ilyakaznacheev/cleanenv"
  "fmt"
)

type settings struct {
  Settings struct {
    Version string `yml:"version"`
  } `yml:"settings"`
}

type Config struct {
  Version string
  Command string
}

const CFG_FILE_NAME = "config.yml"

func MustLoad() (*Config, error) {
  var cfg_settings settings;
  if err := cleanenv.ReadConfig(CFG_FILE_NAME, &cfg_settings); err != nil {
    return &Config{}, fmt.Errorf("failed read %s file", CFG_FILE_NAME)
  }
  
  return &Config{
    Version: cfg_settings.Settings.Version,
    Command: "",
  }, nil
}
