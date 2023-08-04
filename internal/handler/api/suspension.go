package api

type CreateSuspensionRequest struct {
	UserID int64  `json:"user_id" binding:"required" example:"21"`
	Reason string `json:"reason" binding:"required" example:"Bad behaviour"`

	// Time in minutes
	ExpiresIn int64 `json:"expires_in" binding:"required,min=1" swaggertype:"primitive,integer"`
}

type UpdateSuspensionRequest struct {
	ID
	Reason string `json:"reason" binding:"omitempty" example:"Bad behaviour"`

	// Time in minutes
	ExpiresIn int64 `json:"expires_in" binding:"omitempty,min=1" swaggertype:"primitive,integer"`
}
