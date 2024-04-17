package http



type UserIDPathParm struct {
	UserID string `uri:"user-id" binding:"required,uuid_rfc4122"`
}