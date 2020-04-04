package cinlog

import (
	v1 "github.com/sofyan48/go-cinlog/cinlog/v1"
	"github.com/sofyan48/go-cinlog/config"
	"github.com/sofyan48/go-cinlog/entity"
)

// Session ...
type Session struct {
	URL     string
	Service string
}

// NewSession ...
func NewSession(cfg *config.Configs) *Session {
	return &Session{
		URL:     cfg.URL,
		Service: cfg.Service,
	}
}

// V1 ...
func (sess *Session) V1() *v1.V1Session {
	session := &entity.SessionGlobal{}
	session.URL = sess.URL
	session.Service = sess.Service
	return v1.V1SessionHandler(session)
}
