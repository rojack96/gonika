package codec

type Encoder interface {
	Encode() []byte
}
