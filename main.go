package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

const strNonSquadPlayer = "Non Squad Player"

func secondsBetweenTimestamps(t1, t2 string) (int64, error) {
	layout := "2006-01-02 15:04:05 -07"

	time1, err := time.Parse(layout, t1)
	if err != nil {
		return 0, err
	}

	time2, err := time.Parse(layout, t2)
	if err != nil {
		return 0, err
	}

	difference := time2.Sub(time1)
	return int64(difference.Seconds()), nil
}

func main() {
	// Open the JSON file
	file, err := os.Open("input/input.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the JSON data
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Parse the JSON data
	var schema Gw2eiSchema
	err = json.Unmarshal(data, &schema)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// Print accounts and support values
	fmt.Printf("Recorded by: %s, Recorded Account by: %s\n", schema.RecordedBy, schema.RecordedAccountBy)

	seconds, err := secondsBetweenTimestamps(schema.TimeStart, schema.TimeEnd)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Time Start: %s, Time End: %s, Duration: %s, Duration(seconds): %d\n\n", schema.TimeStart, schema.TimeEnd, schema.Duration, seconds)

	for _, player := range schema.Players {
		isContains := strings.Contains(player.Account, strNonSquadPlayer)
		if !isContains {
			for _, support := range player.Support {
				fmt.Printf("Name: %s, Account: %s, Profession: %s, CondiCleanse: %d, BoonStrips: %d, Healing: %d", player.Name, player.Account, player.Profession, support.CondiCleanse, support.BoonStrips, player.ExtHealingStats.OutGoingHealing[0].Healing)
			}
			for _, statsTargets := range player.StatsTargets[0] {
				fmt.Printf(" Killed: %d, Downed: %d\n", statsTargets.Killed, statsTargets.Downed)
			}
		}

	}
}
