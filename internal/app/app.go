package app

import (
	"github.com/vgekko/anistream-template/config"
	"github.com/vgekko/anistream-template/pkg/logger"
)

func Run(cfg *config.Config) {
	_ = logger.New(cfg.Env)
}
