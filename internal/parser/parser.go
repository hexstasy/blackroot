package parser

import (
  "os"
  "gopkg.in/yaml.v3"
  "fmt"
)

const FILE_DIR string = "blackroot"
var pairedCommands map[string]interface{}

func Parser(executableCommand string) (string, error) {
  readyCommand, err := parse(executableCommand)
  if err != nil {
    return "", err 
  }

  return readyCommand, nil
}

func parse(executableCommand string) (string, error) {
  file, err := os.ReadFile(FILE_DIR)
  if err != nil {
    return "", fmt.Errorf("failed to open file: %w", err)
  }

  if err := yaml.Unmarshal(file, &pairedCommands); err != nil {
    return "", fmt.Errorf("failed to parse file: %w", err)
  }

  for k, v := range pairedCommands {
    if k == executableCommand {
      return v.(string), nil
    } 

    return "", fmt.Errorf("could not find command")
  }

  return "", nil
}
