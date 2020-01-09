package representations

// wrappers to be used for boolean pointers. (needed for optional boolean values where the default value in Keycloak is not false)
var (
	TRUE  = true
	FALSE = false
)

type Map map[string]interface{}
type Object interface{}
