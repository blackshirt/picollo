package model

import (
	"fmt"
	"io"
	"strconv"
)

// Rolle access
// A custom enum that uses integers to represent the values in memory
// but serialize as string for graphql
type Role string

const (
	RoleAdmin Role = "Admin"
	RoleUser  Role = "User"
	RoleGuest Role = "Guest"
)

var AllRole = []Role{
	RoleAdmin,
	RoleUser,
	RoleGuest,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleAdmin, RoleUser, RoleGuest:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
