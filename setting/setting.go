package setting

import (
	"log"
)

import (
	"github.com/lxn/walk"
)

type Setting struct {
	form walk.Form
}

func (s *Setting) Init(owner walk.Form) {
	log.Println("setting.init")
	s.form = owner
}

func (s *Setting) Create() {

}