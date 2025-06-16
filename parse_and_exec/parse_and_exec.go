package parseblackroot

import (
  "github.com/ilyakaznacheev/cleanenv"
  "os"
  "io"
  "gopkg.in/yaml.v3"
)

const buf_size = 64
const file_dir = "blackroot.yml"

func ParseAndExecute(command_name string) error {
  if os.IsNotExist(file_dir) {
    return fmt.Errorf("failed parse file, file is exist: %s", file_dir)
  }
  
  cmd_parsed, err := parse(command_name)
  if err != nil {
    return fmt.Errorf("failed parse blackroot file: %w", err)
  }

  fmt.Println(cmd_parsed)
/*
  if err := execute(cmd_parsed); err != nil {
    return fmt.Errorf("failed execute command: %w", err)
  }
*/
  return nil
}

func execute(command string) error {}

func parse(command_name string) (string, error) {
  file, err := os.ReadFile(file_dir)
  if err != nil {
    return "", fmt.Errorf("failed open %s file: %w", file_dir, err)
  }

  defer file.Close() 
  
  data := make(map[interface{}]interface{])
  if err := yaml.Unmarshal(file, data); err != nil {
    return "", fmt.Errorf("failed parse yaml file: %w", err)
  }

  for key, val := range data {
    if key == command_name {
      return val
    }

    return fmt.Errorf("command not found")
  }
}
