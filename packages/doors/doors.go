package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var (
	// keys represents all valid keys for any/all classes.
	keys = []int{1, 2, 3, 4, 5, 6}
)

// Main is the entrypoint of the request.
func Main(args map[string]interface{}) map[string]interface{} {
  raw := args["class"]
  n := ndoors(raw)
  if n == 0 {
    return map[string]interface{}{
			"error": fmt.Sprintf("unknown class %v", raw),
		}
  }

  return map[string]interface{}{
		"body": map[string]interface{}{
			"keys": randomize(n),
		},
	}
}

// ndoors returns the number of doors necessary to open for the given
// class `cl`. Return 0 if `cl` is unknown/invalid.
func ndoors(cl any) int {
	class, ok := cl.(string)
	if !ok {
		return 0
	}
	switch strings.ToLower(strings.TrimSpace(class)) {
	case "fighter", "barbarian", "rogue", "magician", "guardian":
		return 2
	case "samurai", "paladin":
		return 3
	case "monk", "ninja":
		return 4
	case "warlock", "headhunter":
		return 5
	case "alchemist":
		return 6
	default:
		return 0
	}
}

// randomize takes a number of 'keys' n and returns a slice with
// n elements of randomized 'keys'.
func randomize(n int) []int {
  // Get a slice with n total keys
  ret := keys[:n]
  // Randomize the slice and return it
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(ret), func(i, j int) {
		ret[i], ret[j] = ret[j], ret[i]
	})
	return ret
}