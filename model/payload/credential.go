package payload

type Credential struct {
	Username string `json:"username" validate:"nonzero,min=2,max=20" extensions:"x-order=0"`
	Password string `json:"password" validate:"nonzero,min=2,max=50" extensions:"x-order=1"`
}
