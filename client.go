package recognition

type Client interface {
	Recognize(string, []byte) (string, error)
	Stop()
}
