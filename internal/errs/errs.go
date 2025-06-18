package errs

import (
  "fmt"
  "os"
  "github.com/fatih/color"
)


// an invalid error, after processing which the program must be terminated
func Must(msg string) {
  err_msg := fmt.Sprintf("blackroot %s: %s\n", color.RedString("error"), msg) 
  os.Stderr.WriteString(err_msg)
  os.Exit(1)
}
