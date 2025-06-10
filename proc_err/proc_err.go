package procerr

import (
  "fmt"
  "os"
  "github.com/fatih/color"
)

func Err(msg string) {
  err_msg := fmt.Sprintf("blackroot %s: %s\n", color.RedString("error"), msg) 
  os.Stderr.WriteString(err_msg)
  os.Exit(1)
}
