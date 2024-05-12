package f1_2023

import "github.com/DaanV2/go-f1-library/game"

var Game = game.Game{
	Drivers:       new(Drivers),
	Infringements: new(Infringements),
	Nationalities: new(Nationalities),
	Penalties:     new(Penalties),
	Teams:         new(Teams),
	Tracks:        new(Tracks),
}
