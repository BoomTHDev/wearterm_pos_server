package model

type (
	RegisterRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	RegisterResponse struct {
		ID       uint64 `json:"id"`
		Username string `json:"username"`
	}
)
