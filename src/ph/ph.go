package main

import (
  "fmt"
  "os"
  "strings"

  "github.com/codegangsta/cli"
)


func main() {
  // get the data that should be parsed to form the push command
  slug := strings.Join(os.Args[1:], " ")

  // start with the remotes and origins that already exist
  config := RemoteBranchGroup{
    Remote: []string { "origin", "heroku", "gin"},
    Branch: []string { "master", "dev" },
  }

  // do the parsing
  fmt.Println(Parse(config, slug))

  app := cli.NewApp()
  app.Name = "ph"
  app.Usage = "fight the loneliness!"
  app.Action = func(c *cli.Context) error {
    // fmt.Println("Hello friend!")
    return nil
  }

  app.Run(os.Args)
}
