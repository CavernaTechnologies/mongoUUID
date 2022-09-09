package mongoUUID

import (
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
)

type UUID struct {
	uuid.UUID
}

func (u UUID) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bsontype.Binary, bsoncore.AppendBinary(nil, 4, u.UUID[:]), nil
}

func (u *UUID) UnmarshalBSONValue(t bsontype.Type, raw []byte) error {
	if t != bsontype.Binary {
		return fmt.Errorf("invalid format on unmarshal bson value")
	}

	subtype, data, rem, ok := bsoncore.ReadBinary(raw)
	if subtype != 0x03 && subtype != 0x04 {
		return fmt.Errorf("invalid subtype")
	}
	if len(rem) > 0 {
		return fmt.Errorf("too many bytes returned")
	}
	if !ok {
		return fmt.Errorf("not enough bytes to unmarshal bson value")
	}

	copy(u.UUID[:], data)

	return nil
}
