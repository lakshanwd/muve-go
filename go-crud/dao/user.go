package dao

//User - holds user
type User struct {
	UserID int64  `json:"userID,omitempty"`
	Name   string `json:"name,omitempty"`
	Role   string `json:"role,omitempty"`
}
