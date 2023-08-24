package main

import (
	"github.com/vgekko/anistream-template/config"
	"github.com/vgekko/anistream-template/internal/app"
)

func main() {
	cfg := config.Load()

	app.Run(cfg)
}
