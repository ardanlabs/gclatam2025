package role

import (
	"fmt"
)

// The set of roles that can be used.
var (
	Admin = newRole("ADMIN")
	User  = newRole("USER")
)

// =============================================================================

// Set of known roles.
var roles = make(map[string]Role)

// Role represents a role in the system.
type Role struct {
	value string
}

func newRole(role string) Role {
	r := Role{role}
	roles[role] = r
	return r
}

// String returns the name of the role.
func (r Role) String() string {
	return r.value
}

// Equal provides support for the go-cmp package and testing.
func (r Role) Equal(r2 Role) bool {
	return r.value == r2.value
}

// MarshalText provides support for logging and any marshal needs.
func (r Role) MarshalText() ([]byte, error) {
	return []byte(r.value), nil
}

// =============================================================================

// Parse parses the string value and returns a role if one exists.
func Parse(value string) (Role, error) {
	role, exists := roles[value]
	if !exists {
		return Role{}, fmt.Errorf("invalid role %q", value)
	}

	return role, nil
}
