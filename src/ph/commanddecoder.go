package main

import (
  // "fmt
  "regexp"
  "strings"
)

func DecodeCommand(command string) string {
  r := regexp.MustCompile("git (push|pull) ([^ ]+) ([^ ]+)")
  match := r.FindStringSubmatch(command)
  final := "" // start with an empty command

  if len(match) != 0 {
    // push or pull
    if match[1] == "pull" {
      final += "l"
    }

    // first letter of remote
    if len(match[2]) > 0 {
      final += string(match[2][0])
    }

    // first letter of branch
    if len(match[3]) > 0 {
      index := strings.Index(match[3], ":")
      if index != -1 && len(match[3]) > index {
        // there's a colon, so treat specially
        final += string(match[3][0]) + ":" + string(match[3][index+1])
      } else {
        // nothing fancy, so just add the branch like normal
        final += string(match[3][0])
      }
    }
    return final
  } else {
    return ""
  }
}
