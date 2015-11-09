package model

type Response struct {
	Hits struct {
		Hit []struct {
			Id     string                 `json:"_id"`
			Source map[string]interface{} `json:"_source"`
		} `json:"hits"`
		Total int `json:"total"`
	} `json:"hits"`
}
