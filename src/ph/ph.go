package main

import (
  "fmt"
  "os"
  "strings"
  "io/ioutil"
  "path"
  "bufio"
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

func inject() {
  fmt.Println("In order for ph to give you feedback, we're going to install a small function in your ~/.profile file to track all the git pushes and pulls you make. None of the commands you type will leave this system.")

  fmt.Println("Press any key to continue...")
  bio := bufio.NewReader(os.Stdin)
  _, _, err := bio.ReadLine()

  fmt.Println("Injecting shim...")

  // inject this command as an alias
  profile := path.Join(os.Getenv("HOME"), ".profile")
  f, err := os.OpenFile(profile, os.O_APPEND|os.O_WRONLY, 0600)
  if err != nil { panic(err) }
  defer f.Close()

  // write the command to the file
  _, err2 := f.WriteString(`
# pass all git commands through ph so it can log them to give statistics
function git {
  if [ "$1" = "push" ] || [ "$1" = "pull" ]; then
    echo "git $@">> /tmp/dat2
  fi
  $(which git) $@
}`)
  if err2 != nil { panic(err) }

  fmt.Println("Done!")
  fmt.Println("Before ph can start helping you save keystrokes, restart your terminal.")
}

func report() {
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
}

func main() {
  if GitExists() {
    // inject or report?
    if len(os.Args) > 1 && os.Args[1] == "inject" {
      inject()
    } else if len(os.Args) > 1 && os.Args[1] == "report" {
      report()
    } else {
      // a normal command
      action, output, flags := parseArgs()
      if len(output.Remote) > 0 && len(output.Branch) > 0 {
        RunGit(action, output, flags)
      } else {
        fmt.Println("No action specified. Run with --help for help.")
      }
    }
  } else {
    fmt.Println("Git isn't in your PATH. (You'll need to install git first to use ph)")
  }
}
