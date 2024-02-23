package pkg

import (
	"encoding/hex"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IHelper interface {
	ConvertStrToPrimitiveObjectId(value string, target *primitive.ObjectID) error
	GenUuid() string
}
type helper struct {
}

// ConvertStrToPrimitiveObjectId implements IHelper.
func (*helper) ConvertStrToPrimitiveObjectId(value string, target *primitive.ObjectID) error {
	_, err := hex.DecodeString(value)
	if err != nil {
		return fmt.Errorf("Invalid hexadecimal string: %v", err)
	}
	*target, err = primitive.ObjectIDFromHex(value)
	return err
}

// GenUuid implements IHelper.
func (*helper) GenUuid() string {
	return uuid.NewString()
}

func NewHelper() IHelper {
	return &helper{}
}
