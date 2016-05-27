package main

import (
  // "fmt"
  "regexp"
  "strings"
)

type RemoteBranchGroup struct {
  Remote []string
  Branch []string
}

// given a slug used array, start and end, this function returns a bool
// singifying if the region specified is clear (ie, all false)
func indexAreaFree(slugCharsUsed []bool, start int, end int) bool {
  for ct, i := range slugCharsUsed {
    if ct >= start && ct < end && i == true { // not free
      return false
    }
  }
  return true
}

// given a haystack and a collection of needles, return all item indexes that are
// contained within. Also, mark the used charaters within haystackCharsUsed.
// A few edge cases that this handles:
// - Because we are searching through branches in sequence, the range
// branch2:branch1 may noe be parsed correctly if branch1 comes up before
// branch2. Therefore, the below code handles both orders independantly.
func itemsWithin(haystack string, collection []string, haystackCharsUsed []bool) [][]int {
  var totals [][]int

  // mark the locations of the "starts" to colon-joined fields
  endsWithColon := make(map[int] int) // for foo: in `foo:bar`
  startsWithColon := make(map[int] int) // for :bar in `foo:bar`

  for index, item := range collection {
    match := regexp.MustCompile(regexp.QuoteMeta(item)).FindStringIndex(haystack)
    if (len(match) != 0) {

      // if the specified area hasn't been used, then we're good
      if indexAreaFree(haystackCharsUsed, match[0], match[1]) {
        // if this branch has a colon before or after, make sure it's treated specially
        isColonBranch := false

        // search for colons after and before the current branch
        if match[0] > 0 && haystack[match[0] - 1] == ':' {
          // would match def in `abc:def`
          startsWithColon[match[0] - 1] = index;
          isColonBranch = true
        } else if len(haystack) > match[1] && haystack[match[1]] == ':' {
          // would match abc in `abc:def`
          endsWithColon[match[1] + 1] = index;
          isColonBranch = true
        }

        // is this the first part of a colon-seperated branch group, or
        // is this the second part? (or neither?)
        if afterColon, ok := startsWithColon[match[1]]; ok {
          totals = append(totals, []int{index, afterColon})
        } else if beforeColon, ok := endsWithColon[match[0]]; ok {
          totals = append(totals, []int{beforeColon, index})
        } else if !isColonBranch {
          // no colon, so just add the new branch name
          totals = append(totals, []int{index})
        }

        // mark chars as used
        for k := match[0]; k < match[1]; k++ {
          haystackCharsUsed[k] = true;
        }
      }
    }
  }

  return totals
}

// return the first character of each of the heystack elements passed
func firstCharOf(haystack []string) []string {
  var total []string
  for _, v := range haystack {
    total = append(total, string(v[0]))
  }
  return total
}

// given a slice of strings and another slice of indecies from the original
// slice, extract the original data at the specified indexes.
func pluckElementsByIndex(haystack []string, needles [][]int) []string {
  // given: [[1], [2, 3], [4]]
  var total []string
  // fmt.Println(needles)


  for _, needle := range needles {
    if len(needle) == 1 {
      // just a normal element, so only needle[0] really matters
      // just search though and look for a match
      for index, elem := range haystack {
        if index == needle[0] {
          total = append(total, elem)
        }
      }
    } else if len(needle) == 2 {
      // colon-separated element
      // look for both the part after the colon and the part before
      // then, merge them together and push to the array
      firstPart := ""
      secondPart := ""

      for index, elem := range haystack {
        if index == needle[0] {
          firstPart = elem
        } else if index == needle[1] {
          secondPart = elem
        }
      }

      if len(firstPart) != 0 && len(secondPart) != 0 {
        total = append(total, firstPart + ":" + secondPart)
      }
    }
  }


  // for index, elem := range haystack {
  //   for _, needle := range needles {
  //     if index == needle {
  //       total = append(total, elem)
  //     }
  //   }
  // }

  return total
}

// the main function call. What happens here is reletively simple: We start by
// marking an array of whether each passed character has been used. Then, we
// loop though the passed slug to find remote/branch names, or their aliases. As
// we go, we "check them off" in the array created at the top. At the end, we
// know which characters are unused and the array with their positions is
// reterned to be used later.
func Parse(config RemoteBranchGroup, slug string) (RemoteBranchGroup, []bool, []string) {
  // create an array, specifying whether a charater has been used
  var slugCharsUsed []bool
  for k := 0; k < len(slug); k++ {
    slugCharsUsed = append(slugCharsUsed, false)
  }

  ////////////
  // FLAGS
  ////////////

  // return all flags that the user has added outside of the app
  r := regexp.MustCompile("(-[^ -] [^ ]+|--[^ ]+ [^ ]+|-[^ -]|--[^ ]+|--)")
  flags := r.FindAllString(slug, -1)

  // mark the flag characters as used
  for _, i := range flags {
    start := strings.Index(slug, i)
    end := start + len(i)

    // mark chars as used that are flags
    for k := start; k < end; k++ {
      slugCharsUsed[k] = true;
    }
  }

  ////////////
  // REMOTES
  ////////////

  // get all matching full-length remotes
  remoteIndexes := itemsWithin(slug, config.Remote, slugCharsUsed)
  remotes := pluckElementsByIndex(config.Remote, remoteIndexes)

  // get all matching 1 character remote aliases (ie, o for origin, h for heroku, etc)
  oneCharRemoteIndexes := itemsWithin(slug, firstCharOf(config.Remote), slugCharsUsed)
  remotes = append(
    remotes,
    pluckElementsByIndex(config.Remote, oneCharRemoteIndexes)...
  )

  ////////////
  // BRANCHES
  ////////////
  config.Branch = append(config.Branch, "current") // add current branch as an alias

  // get all matching full-length branches
  branchIndexes := itemsWithin(slug, config.Branch, slugCharsUsed)
  branches := pluckElementsByIndex(config.Branch, branchIndexes)

  // get all matching 1 character branch aliases (ie, m for master, d for dev, etc)
  oneCharbranchIndexes := itemsWithin(slug, firstCharOf(config.Branch), slugCharsUsed)
  branches = append(
    branches,
    pluckElementsByIndex(config.Branch, oneCharbranchIndexes)...
  )

  return RemoteBranchGroup{Remote: remotes, Branch: branches}, slugCharsUsed, flags;
}
