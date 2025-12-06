package token

import "time"

type Maker interface {

	// create a token
	CreateToken(username string, duration time.Duration) (string, error)
	//verify a token
	VerifyToken(token string) (*Payload, error)
}
