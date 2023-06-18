package helper

import "gopi/internal/config"

func (h *helper) GetAppScheme() string {
	scheme := "http"
	if h.config.AppEnv == config.PRODUCTION {
		scheme += "s"
	}
	return scheme
}
