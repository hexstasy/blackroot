package procerr

import (
  "fmt"
  "os"
)

func Err(msg string) {
  err_msg := fmt.Sprintf("blackroot error: %s\n", msg) 
  os.Stderr.WriteString(err_msg)
  os.Exit(1)
}
