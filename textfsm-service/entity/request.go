package entity

type CloudEventData struct {
	Meta struct {
		Timestamp     int64  `json:"timestamp"`
		Username      string `json:"username"`
		TxID          int    `json:"txId"`
		TxEventID     int    `json:"txEventId"`
		TxEventsCount int    `json:"txEventsCount"`
		Operation     string `json:"operation"`
		Source        struct {
			Hostname string `json:"hostname"`
		} `json:"source"`
	} `json:"meta"`
	Payload struct {
		ID    string `json:"id"`
		Start struct {
			ID     string                 `json:"id"`
			Labels []string               `json:"labels"`
			IDs    map[string]interface{} `json:"ids"`
		} `json:"start"`
		End struct {
			ID     string                 `json:"id"`
			Labels []string               `json:"labels"`
			IDs    map[string]interface{} `json:"ids"`
		} `json:"end"`
		Before interface{} `json:"before"`
		After  struct {
			Properties map[string]interface{} `json:"properties"`
		} `json:"after"`
		Label string `json:"label"`
		Type  string `json:"type"`
	} `json:"payload"`
	Schema struct {
		Properties  map[string]interface{} `json:"properties"`
		Constraints []string               `json:"constraints"`
	} `json:"schema"`
}
