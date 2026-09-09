package domain

type TokenGenerator interface {
	GenerateToken(userID uint, role Role) (token string, expiresInSecond int, err error)
}
