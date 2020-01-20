package shortener

import (
	"errors"
	"log"
	"time"

	errs "github.com/pkg/errors"
	"github.com/teris-io/shortid"
	"gopkg.in/validator.v2"
)

var (
	//ErrRedirectNotFound ...
	ErrRedirectNotFound = errors.New("Redirect Not Found")
	//ErrRedirectInvalid ...
	ErrRedirectInvalid = errors.New("Redirect Invalid")
)

type redirectService struct {
	redirectRepo RedirectRepository
}

func (r redirectService) Find(code string) (*Redirect, error) {
	return r.redirectRepo.Find(code)
}

func (r redirectService) Store(redirect *Redirect) error {
	if err := validator.Validate(redirect); err != nil {
		log.Println(err)
		return errs.Wrap(ErrRedirectInvalid, "service.Redirect.Store")
	}
	redirect.Code = shortid.MustGenerate()
	redirect.CreatedAt = time.Now().UTC().Unix()
	log.Println(redirect)
	return r.redirectRepo.Store(redirect)
}

//NewRedirectService ...
func NewRedirectService(redirectRepo RedirectRepository) RedirectService {
	return &redirectService{
		redirectRepo,
	}
}
