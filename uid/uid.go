package uid

import (
	"encoding/json"

	"github.com/google/uuid"
)

// A UID is used to uniquely identify any resource in the system.
type UID struct {
	uid *uuid.UUID
}

// NewID returns a new ID.
func NewUID() UID {
	uid := uuid.New()
	return UID{
		uid: &uid,
	}
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
