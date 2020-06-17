package schemas

type GetPublicKeyResponse struct {
	PublicKey string `json:"public_key"`
	Id        int64  `json:"id"`
}
