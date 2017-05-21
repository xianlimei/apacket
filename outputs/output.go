package outputs

type Outputer interface {
	Output(msg []byte)
}
