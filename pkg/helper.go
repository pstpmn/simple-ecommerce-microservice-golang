package pkg

import (
	"github.com/google/uuid"
)

type IHelper interface {
	GenUuid() string
}
type helper struct {
}

// GenUuid implements IHelper.
func (*helper) GenUuid() string {
	return uuid.NewString()
}

func NewHelper() IHelper {
	return &helper{}
}
