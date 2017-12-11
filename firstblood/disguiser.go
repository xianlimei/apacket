package firstblood

type Disguiser interface {
	Fingerprint(data []byte) (identify bool, err error)
	DisguiserData() (data []byte)
}

var DisguiserMap []Disguiser

func init() {
	http := NewHttp()
	DisguiserMap = append(DisguiserMap, http)
}
