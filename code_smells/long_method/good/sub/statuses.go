package sub

type status int

const (
	userActive    status = 1
	userNotActive status = 0
)

func (s status) isActive() bool {
	return s == userActive
}
