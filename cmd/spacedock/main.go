package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"spacetrash/input"
	"spacetrash/model"
	"spacetrash/service"
	"spacetrash/storage"
)

func main() {
	dbPath := flag.String("db", "spacedock.db", "database path")
	flag.Parse()
	store, err := storage.Open(*dbPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer store.Close()
	player := model.NewPlayer("pilot-1", "Ari")
	vessel := model.NewVessel("vessel-1", player.ID, "Wayfarer")
	weapons := []model.Weapon{model.NewWeapon("weapon-1", vessel.ID, "Pulse", 0), model.NewWeapon("weapon-2", vessel.ID, "Arc", 1)}
	session, err := service.NewSession(store, player, vessel, weapons)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	asteroids := []model.Asteroid{model.NewAsteroid("rock-1", 8, 0, 1), model.NewAsteroid("rock-2", 15, 2, 2)}
	if err := session.BeginMission("mission-1", 2, asteroids); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	executor := input.Executor{Session: session}
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("Space Dock ready. Type help for commands.")
	for reader.Scan() {
		command, parseErr := input.Parse(reader.Text())
		if parseErr != nil {
			fmt.Println(parseErr)
			continue
		}
		output, done, executeErr := executor.Execute(command)
		if executeErr != nil {
			fmt.Println(executeErr)
		} else if output != "" {
			fmt.Println(output)
		}
		if done {
			break
		}
	}
}
