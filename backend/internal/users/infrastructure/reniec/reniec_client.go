package reniec

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/joaquin22/hospital-api/internal/users/domain"
)

// defaultTimeout limita la espera por respuesta de RENIEC: el registro de un
// usuario no debe quedar colgado si la API externa se demora.
const defaultTimeout = 10 * time.Second

// dniResponse es el body que devuelve la API cuando el DNI existe.
type dniResponse struct {
	DocumentNumber string `json:"document_number"`
}

// apiError es el body de error de la API: {"message": "..."} en el 404 de
// "no encontrado" y {"error": "..."} en el 401 de API key faltante.
type apiError struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

func (e apiError) String() string {
	if e.Message != "" {
		return e.Message
	}
	return e.Error
}

// ReniecClient implementa domain.DniVerifier contra la API de Decolecta.
type ReniecClient struct {
	apiKey  string
	baseURL string
	http    *http.Client
}

var _ domain.DniVerifier = (*ReniecClient)(nil)

func NewReniecClient(apiKey string, baseURL string) *ReniecClient {
	return &ReniecClient{
		apiKey:  apiKey,
		baseURL: baseURL,
		http:    &http.Client{Timeout: defaultTimeout},
	}
}

func (c *ReniecClient) Exists(dni string) (bool, error) {
	if c.apiKey == "" {
		return false, errors.New("decolecta: falta la API key (DNI_KEY)")
	}

	endpoint := fmt.Sprintf("%s/reniec/dni?numero=%s", c.baseURL, url.QueryEscape(dni))

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return false, err
	}
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Accept", "application/json")

	res, err := c.http.Do(req)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case http.StatusOK:
		var body dniResponse
		if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
			return false, err
		}
		// La API devuelve el DNI encontrado; si no viene o no coincide con el
		// consultado, no lo damos por válido.
		return body.DocumentNumber == dni, nil
	case http.StatusNotFound:
		// 404 = DNI inexistente en RENIEC, no es un fallo de la API.
		return false, nil
	default:
		return false, fmt.Errorf("decolecta: status inesperado %d: %s", res.StatusCode, readError(res.Body))
	}
}

func readError(body io.Reader) string {
	var apiErr apiError
	if err := json.NewDecoder(body).Decode(&apiErr); err != nil {
		return "sin cuerpo de respuesta"
	}
	if msg := apiErr.String(); msg != "" {
		return msg
	}
	return "sin mensaje de error"
}
