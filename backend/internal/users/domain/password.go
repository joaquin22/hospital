package domain

// MinPasswordLength es la longitud mínima exigida para una contraseña.
const MinPasswordLength = 8

// ValidatePassword comprueba la contraseña en texto plano y devuelve
// ErrWeakPassword si no cumple el mínimo.
//
// Se invoca desde los casos de uso ANTES de hashear. No se puede hacer dentro
// de NewUser porque a este solo llega el hash de bcrypt, cuya longitud no dice
// nada sobre la de la contraseña original.
func ValidatePassword(plain string) error {
	if len(plain) < MinPasswordLength {
		return ErrWeakPassword
	}
	return nil
}
