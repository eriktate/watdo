package uid

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/google/uuid"
)

// A UID is used to uniquely identify any resource in the system.
type UID struct {
	uid *uuid.UUID
}

// New returns a new ID.
func New() UID {
	uid := uuid.New()
	return UID{
		uid: &uid,
	}
}

// Nil returns a blank UID.
func Nil() UID {
	return UID{}
}

// ParseString creates a UID from some stringified ID.
func ParseString(id string) (UID, error) {
	var uniqueID UID
	if id == "" || id == "null" {
		return uniqueID, nil
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		return uniqueID, err
	}

	uniqueID.uid = &uid
	return uniqueID, nil
}

// String implements the Stringer interface.
func (u UID) String() string {
	if !u.Empty() {
		return u.uid.String()
	}

	return ""
}

func (u UID) JSONString() string {
	str := u.String()
	if str != "" {
		return "\"" + str + "\""
	}

	return ""
}

// Empty returns whether or not the UID has a value yet.
func (u UID) Empty() bool {
	return u.uid == nil
}

// Equal returns whether or not the given UID matches the receiver.
func (u UID) Equal(uid UID) bool {
	return u.String() == uid.String()
}

func (u *UID) UnmarshalJSON(data []byte) error {
	var dest uuid.UUID
	if len(data) == 0 || string(data) == "null" {
		return nil
	}

	if err := json.Unmarshal(data, &dest); err != nil {
		return err
	}

	u.uid = &dest
	return nil
}

func (u UID) MarshalJSON() ([]byte, error) {
	if u.uid != nil {
		return json.Marshal(u.uid)
	}

	return json.Marshal(nil)
}

func (u *UID) Scan(src interface{}) error {
	var uid uuid.UUID
	if err := uid.Scan(src); err != nil {
		return err
	}

	u.uid = &uid
	return nil
}

func (u UID) Value() (driver.Value, error) {
	if u.Empty() {
		return nil, nil
	}

	return u.String(), nil
}
