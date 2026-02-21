package server

type Message struct {
	Type string      `json:"type"` //ход, чат, состояние и система
	Data interface{} `json:"data"`
}

type TurnMessage struct {
	Hit   string `json:"hit"`
	Block string `json:"block"`
}
