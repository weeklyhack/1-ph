package main

import (
  "fmt"
  "os"
  "strings"
  // "exec"

  "github.com/codegangsta/cli"
)


func main() {
  flagsLookup := map[string] string {
    "f": "--force",
    "n": "--dry-run",
    "u": "--set-upstream",
    "q": "--quiet",
    "v": "--verbose",
  }

  // get the data that should be parsed to form the push command
  slug := strings.Join(os.Args[1:], " ")

  // start with the remotes and origins that already exist
  config := RemoteBranchGroup{
    Remote: []string { "origin", "heroku", "gin"},
    Branch: []string { "master", "dev" },
  }

  // do the parsing
  output, availableChars := Parse(config, slug)

  // given the rest of the flags, see if they have been included
  flags := ""
  action := "push"
  for index, unused := range availableChars {
    if !unused {
      // are we performing a push or a pull?
      if slug[index] == 'l' {
        action = "pull"
      } else if longFlag, ok := flagsLookup[string(slug[index])]; ok {
        flags = flags + longFlag + " "
        availableChars[index] = true // mark this character as used
      }
    }
  }
  flags = strings.Trim(flags, " ")

  // run the command
  for _, remote := range output.Remote {
    for _, branch := range output.Branch {
      fmt.Println("$ git", action, remote, branch, flags)
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
