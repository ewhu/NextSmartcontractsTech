// cmd/nextsmartcontractstech/main.go
package main

import (
"flag"
"log"
"os"

"nextsmartcontractstech/internal/nextsmartcontractstech"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := nextsmartcontractstech.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
