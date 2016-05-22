package main

import (
  "fmt"
  "regexp"
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
// contained within. Also, mark the used charaters within haystackCharsUsed
func itemsWithin(haystack string, collection []string, haystackCharsUsed []bool) []int {
  var totals []int

  for index, item := range collection {
    match := regexp.MustCompile(item).FindStringIndex(haystack)
    if (len(match) != 0) {
      // if the specified area hasn't been used, then we're good
      if indexAreaFree(haystackCharsUsed, match[0], match[1]) {
        totals = append(totals, index)

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
func pluckElementsByIndex(haystack []string, needles []int) []string {
  var total []string
  for index, elem := range haystack {
    for _, needle := range needles {
      if index == needle {
        total = append(total, elem)
      }
    }
  }

  return total
}

// the main function call.
func Parse(config RemoteBranchGroup, slug string) RemoteBranchGroup {
  // create an array, specifying whether a charater has been used
  var slugCharsUsed []bool
  for k := 0; k < len(slug); k++ {
    slugCharsUsed = append(slugCharsUsed, false)
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

  // get all matching full-length branches
  branchIndexes := itemsWithin(slug, config.Branch, slugCharsUsed)
  branches := pluckElementsByIndex(config.Branch, branchIndexes)

  // get all matching 1 character branch aliases (ie, m for master, d for dev, etc)
  oneCharbranchIndexes := itemsWithin(slug, firstCharOf(config.Branch), slugCharsUsed)
  branches = append(
    branches,
    pluckElementsByIndex(config.Branch, oneCharbranchIndexes)...
  )

  fmt.Println("Unassigned Chars", slugCharsUsed)

  return RemoteBranchGroup{Remote: remotes, Branch: branches}
}
