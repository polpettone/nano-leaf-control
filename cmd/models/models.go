package models

import "encoding/json"

type Effects struct {
	EffectsList []string `json:"effectsList"`
	Selected    string   `json:"select"`
}

func ConvertJsonToEffects(jsonData []byte) (*Effects, error) {
	var effects Effects
	err := json.Unmarshal(jsonData, &effects)
	if err != nil {
		return nil, err
	}
	return &effects, nil
}
