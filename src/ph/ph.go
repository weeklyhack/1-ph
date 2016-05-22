package main

import (
  "fmt"
  "os"
  "strings"
  // "exec"

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
  output, _ := Parse(config, slug)

  for _, remote := range output.Remote {
    for _, branch := range output.Branch {
      fmt.Println("$ git push", remote, branch)
    }
  }
  // out, err := exec.Command("git", "push", ).Output()

  app := cli.NewApp()
  app.Name = "ph"
  app.Usage = "fight the loneliness!"
  app.Action = func(c *cli.Context) error {
    fmt.Println("Hello friend!")
    return nil
  }

  app.Run(os.Args)
}
