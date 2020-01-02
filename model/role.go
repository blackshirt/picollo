package model

import (
	"fmt"
	"io"
	"strconv"
)

// Rolle access
// A custom enum that uses integers to represent the values in memory
// but serialize as string for graphql
type Role uint

const (
	Guest    Role = iota // 0
	User                 // 1
	Admin                // 2
	Superman             // 3
)

func RoleFrom(str string) (Role, error) {
	switch str {
	case "Guest":
		return Guest, nil
	case "User":
		return User, nil
	case "Admin":
		return Admin, nil
	case "Superman":
		return Superman, nil
	default:
		return 0, fmt.Errorf("%s is not a valid Role", str)
	}
}

func (e Role) IsValid() bool {
	switch e {
	case Guest, User, Admin, Superman:
		return true
	}
	return false
}

func (e Role) String() string {
	switch e {
	case Guest:
		return "Guest"
	case User:
		return "User"
	case Admin:
		return "Admin"
	case Superman:
		return "Superman"

	default:
		panic("invalid enum value")
	}
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	var err error
	*e, err = RoleFrom(str)
	return err
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
