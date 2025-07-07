package constants

const (
	RoleUser      string = "user"
	RoleAdmin     string = "admin"
	RoleModerator string = "moderator"
)

func IsValidRole(role string) bool {
	switch role {
	case RoleUser, RoleAdmin:
		return true
	default:
		return false
	}
}
