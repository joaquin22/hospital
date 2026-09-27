package domain

// DniVerifier consulta una fuente externa (RENIEC) para confirmar que un DNI
// realmente existe antes de dar de alta a una persona. La implementación real
// (cliente HTTP) vive en infrastructure; el dominio solo depende de esta
// interfaz para poder testear el caso de uso sin salir a la red.
type DniVerifier interface {
	// Exists devuelve true si el DNI está registrado en RENIEC.
	// Devuelve error solo si no se pudo determinar la existencia
	// (timeout, API key inválida, API caída, etc).
	Exists(dni string) (bool, error)
}
