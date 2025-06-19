package runner

import (
  "os/exec"
  "fmt"
  "strings"
)

func Run(command string) error {
  args := strings.Fields(command)
  cmd := exec.Command(args[0], args[1:]...)
  
  out, err := cmd.CombinedOutput()
  if err != nil {
    return fmt.Errorf("failed to execute command: %w", err)
  }

  fmt.Println(string(out))
  return nil
}
