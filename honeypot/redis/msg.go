package redis

type RedisMsg struct {
	Payload []byte `json:"-"`
	Ctype   string `json:"ctype,omitempty"`
}
