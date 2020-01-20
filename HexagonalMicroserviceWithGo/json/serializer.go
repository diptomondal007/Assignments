package json

import(
	"encoding/json"
	errs "github.com/pkg/errors"
	"hexagonalmicroservicewithgo/shortener"
	"log"
)

type Redirect struct{}

func (r *Redirect) Encode(redirect *shortener.Redirect) ([]byte, error) {
	raw, err := json.Marshal(redirect)
	if err != nil{
		return nil, errs.Wrap(err, "serializer.Redirect.Encode")
	}
	return raw, nil
}

func (r *Redirect) Decode(b []byte) (*shortener.Redirect, error) {
	redirect := &shortener.Redirect{}
	if err:= json.Unmarshal(b,redirect); err != nil{
		return nil, errs.Wrap(err, "serializer.Redirect.Decode")
	}
	log.Println(redirect)
	return redirect, nil
}
