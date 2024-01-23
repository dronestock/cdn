package main

import (
	"github.com/dronestock/cdn/internal"
	"github.com/dronestock/drone"
)

func main() {
	drone.New(internal.New).Boot()
}
