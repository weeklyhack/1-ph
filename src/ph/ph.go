package main

import (
  "fmt"
  "os"
  "strings"
  "io/ioutil"
  "path"
  "bufio"
)

func ParseArgs(args []string) (string, RemoteBranchGroup, string) {
  flagsLookup := map[string] string {
    "f": "--force",
    "n": "--dry-run",
    "t": "--set-upstream",
    "q": "--quiet",
    "v": "--verbose",
  }

  // get the data that should be parsed to form the push command
  slug := strings.Join(args[1:], " ")

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
    // help or inject or report?
    if len(os.Args) > 1 && os.Args[1] == "help" {
      fmt.Println("PH: Add some chemistry to your terminal.")
      fmt.Println("Easily shorten git pushes and git pulls by abbreviating remotes and branches.")
      fmt.Println("(The branch and remote names are magically pulled from the current repository)")
      fmt.Println("For example:")
      fmt.Println("  ph om -> git push origin master")
      fmt.Println("  ph od -> git push origin dev")
      fmt.Println("  ph hm -> git push heroku master")
      fmt.Println("Or don't - that works too.")
      fmt.Println("  ph origin m -> git push origin master")
      fmt.Println("  ph origin master -> git push origin master")
      fmt.Println("Also, easily add commonly used flags when performing operations.")
      fmt.Println("  ph oml -> git pull origin master")
      fmt.Println("  ph omf -> git push origin master --force")
      fmt.Println("  ph omu -> git push origin master --set-upstream")
      fmt.Println("  ph omn -> git push origin master --dry-run")
      fmt.Println("  ph omq -> git push origin master --quiet")
      fmt.Println("  ph omv -> git push origin master --verbose")
      fmt.Println("  ph om --flag -> git push origin master --flag")
      fmt.Println("Are you still skeptical? Prove it to yourself by running ph inject. We'll track your git pushes and pulls so you can see how many keystrokes you'd save by using ph.")
      fmt.Println()
      fmt.Println("An app by Ryan Gaus (http:/rgaus.net) built for my weekly hacks challenge. Learn more at http://weeklyhacks.github.io/ph")
    } else if len(os.Args) > 1 && os.Args[1] == "inject" {
      inject()
    } else if len(os.Args) > 1 && os.Args[1] == "report" {
      report()
    } else {
      // a normal command
      action, output, flags := ParseArgs(os.Args)
      if len(output.Remote) > 0 && len(output.Branch) > 0 {
        RunGit(action, output, flags)
      } else {
        fmt.Println("No action specified. Run with ph help for help.")
      }
    }
  } else {
    fmt.Println("Git isn't in your PATH. (You'll need to install git first to use ph)")
  }
}
