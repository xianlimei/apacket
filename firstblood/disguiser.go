package firstblood

type Disguiser interface {
	Fingerprint(request []byte) (identify bool, err error)
	DisguiserResponse(request []byte) (response []byte)
}

var DisguiserMap []Disguiser

func init() {
	http := NewHttp()
	DisguiserMap = append(DisguiserMap, http)
}
