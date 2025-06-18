package parser

import (
  "github.com/ilyakaznacheev/cleanenv"
  "os"
  "io"
  "gopkg.in/yaml.v3"
  "fmt"
)

const BUF_SIZE = 64
const FILE_DIR = "blackroot.yml"
var pairedCommands map[interface{}]interface{}

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
  file, err := os.ReadFile(  
}
