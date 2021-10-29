package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var (
	allKeys = []int{1, 2, 3, 4, 5, 6}

	ErrInvalidClass = errors.New("invalid class")
)

func Main(args map[string]interface{}) map[string]interface{} {
	class := "feeder"
	if c, ok := args["class"].(string); ok {
		class = c
	}

	msg := make(map[string]interface{})

	keys, err := keysForClass(class)
	if err != nil {
		msg["error"] = err.Error()
	}

	shuffled := randomize(keys...)

	shuffledMap := map[string]interface{}{
		"keys": shuffled,
	}
	shuffledJson, err := json.Marshal(shuffledMap)
	if err != nil {
		msg["error"] = err.Error()
	}

	msg["body"] = fmt.Sprintf("%v", string(shuffledJson))
	return msg
}

func keysForClass(class string) ([]int, error) {
	numDoors, err := numDoorsForClass(class)
	if err != nil {
		return nil, err
	}

	return allKeys[:numDoors], nil
}

func numDoorsForClass(class string) (int, error) {
	switch class {
	case "feeder",
		"barbarian",
		"fighter",
		"guardian",
		"magician",
		"rogue":
		return 2, nil
	case "samurai", "paladin":
		return 3, nil
	case "monk", "ninja":
		return 4, nil
	case "warlock", "headhunter":
		return 5, nil
	case "alchemist":
		return 6, nil
	default:
		return 0, ErrInvalidClass
	}
}

func randomize(keys ...int) []int {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(keys), func(i, j int) {
		keys[i], keys[j] = keys[j], keys[i]
	})
	return keys
}
