package shortener

//RedirectSerializer ...
type RedirectSerializer interface {
	Encode(redirect *Redirect) ([]byte, error)
	Decode(b []byte) (*Redirect, error)
}
