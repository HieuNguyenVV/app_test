package common

type Role int

const (
	Admin Role = 1
	Other Role = 99
)

func (s Role) IsValid() bool {
	switch s {
	case Admin, Other:
		return true
	}
	return false
}
