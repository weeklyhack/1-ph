package main

import (
  "fmt"
  "os"
  "strings"
  "io/ioutil"

  "github.com/codegangsta/cli"
)

func parseArgs() (string, RemoteBranchGroup, string) {
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
  remotes := GetCwdRemotes()
  branches := GetCwdBranches()
  config := RemoteBranchGroup{Remote: remotes, Branch: branches}
  fmt.Println(config)

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

  return action, output, flags
}

func main() {
  app := cli.NewApp()
  app.Name = "ph"
  app.Usage = "Add some chemistry to your git push."
  app.Commands = []cli.Command {
    {
      Name:      "inject",
      Aliases:     []string{},
      Usage:     "inject ph into git to get stats",
      Action: func(c *cli.Context) error {
        command := strings.Join(os.Args[3:], " ")
        binary := c.Args().First()

        // open the file
        f, err := os.OpenFile("/tmp/dat2", os.O_APPEND|os.O_WRONLY, 0600)
        if err != nil { panic(err) }
        defer f.Close()

        // write the command to the file
        _, err2 := f.WriteString(binary+" "+command+"\n")
        if err2 != nil { panic(err) }

        // run the command
        return RunCmd(binary, os.Args[3:])
      },
    },
    {
      Name:      "report",
      Aliases:     []string{},
      Usage:     "report on how many characters ph could have saved you",
      Action: func(c *cli.Context) error {
        // open the file
        data, err := ioutil.ReadFile("/tmp/dat2")
        if err != nil { panic(err) }

        // read commands
        commands := strings.Split(string(data), "\n")
        totalSavings := 0
        for _, command := range commands {
          phCommand := DecodeCommand(command)
          savings := len(command) - (3 + len(phCommand))
          totalSavings += savings
          if len(phCommand) > 0 {
            fmt.Println(command, " -> ph", phCommand, "saving", savings, "characters")
          }
        }

        fmt.Println("In total, you could save", totalSavings, "characters by using ph instead of git push")
        return nil
      },
    },
  }

  app.Action = func(c *cli.Context) error {
    if GitExists() {
      action, output, flags := parseArgs()
      if len(output.Remote) > 0 && len(output.Branch) > 0 {
        RunGit(action, output, flags)
      } else {
        fmt.Println("No action specified. Run with --help for help.")
      }
    } else {
      fmt.Println("Git isn't in your PATH. (You'll need to install git first to use ph)")
    }

    return nil
  }

  app.Run(os.Args)
}
