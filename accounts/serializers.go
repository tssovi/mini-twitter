package accounts

// LoginSerializer Serializer for login form
type LoginSerializer struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// UserLoginSerializer Serializer for user details on login
type UserLoginSerializer struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Token    string `json:"token"`
}

// UserDetailsSerializer Serializer for user details
type UserDetailsSerializer struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}
