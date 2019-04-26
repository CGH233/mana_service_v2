package config

import (
	"github.com/asynccnu/mana_service_v2/model"
)

type UpdateRequest struct {
	Config model.IOSConfig `json:"config"`
}

type UpdateResponse struct {
}
