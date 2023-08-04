package entity

import "strings"

type Role int

const (
	USER Role = iota
	MODERATOR
	ADMIN
)

var RoleMap = map[string]Role{
	"USER":      USER,
	"MODERATOR": MODERATOR,
	"ADMIN":     ADMIN,
}

func StringToRole(str string) (Role, bool) {
	c, ok := RoleMap[strings.ToUpper(str)]
	return c, ok
}

func (role Role) String() string {
	for k, v := range RoleMap {
		if v == role {
			return k
		}
	}

	return ""
}
