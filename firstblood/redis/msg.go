package redis

type RedisMsg struct {
	Payload []byte `json:"payload,omitempty"`
}
