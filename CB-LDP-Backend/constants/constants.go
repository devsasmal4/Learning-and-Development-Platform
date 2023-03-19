package constants

type UserRole string

const (
	DefaultTimeOut = 10
)

const (
	Admin        UserRole = "Admin"
	Module_Owner          = "Module Owner"
	User                  = "User"
)

func (r UserRole) String() string {
	switch r {
	case Admin:
		return "Admin"
	case Module_Owner:
		return "Module Owner"
	case User:
		return "User"
	}
	return "User"
}

// Cors Constants
const Origin = "Origin"
const ContentTypeHeader = "Content-Type"
const Accept = "Accept"
const Authorization = "Authorization"
const DateUsed = "DateUsed"
const XRequestedWith = "X-Requested-With"
const Cookie = "Cookie"
const Email = "Email"
const Token = "Token"

var AllowedOrigin = []string{"https://beandigest.coffeebeans.io", "http://localhost:3000", "https://beandigest-dev.coffeebeans.io"}
