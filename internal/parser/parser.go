package parser

import (
  "github.com/ilyakaznacheev/cleanenv"
  "os"
  "io"
  "gopkg.in/yaml.v3"
  "fmt"
)

const FILE_DIR = "blackroot.yml"
var pairedCommands map[string]interface{}

func Parser(executableCommand string) (string, error) {
  if os.IsNotExist(FILE_DIR) { 
    return "", fmt.Errorf("failed parse file, file is exist: %s", file_dir)
  }
  
  readyCommand, err := parse(executableCommand)
  if err != nil {
    return "", fmt.Errorf("failed parse blackroot file: %w", err)
  }

  return readyCommand, nil
}

func parse(executableCommand string) (string, error) {
  file, err := os.Open(FILE_DIR)
  if err != nil {
    return "", fmt.Errorf("failed to open file, dir: %s: %w", FILE_DIR, err)
  }

  if err := yaml.Unmarshal(file, &pairedCommands); err != nil {
    return "", fmt.Errorf("failed to parse file: %w", err)
  }

  for k, v := range pairedCommands {} 
}
