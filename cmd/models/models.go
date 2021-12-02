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

type OnStateValue struct {
	Value bool `json:"value"`
}

type OnState struct {
	On OnStateValue `json:"on"`
}
