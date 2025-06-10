package cli

import (
  "flag"
  "fmt"
)

func ParseOpts() {
  version := flag.Bool("version", false, "Display version and exit.")
  if version {
    fmt.Printf("Blackroot version: %s", VERSION) 
  }
  flag.Parse();
}
