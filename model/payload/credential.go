package payload

type Credential struct {
	Username string `json:"username" validate:"nonzero,min=2,max=20"`
	Password string `json:"password" validate:"nonzero,min=2,max=50"`
}
