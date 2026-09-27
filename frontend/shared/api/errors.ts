import type { AxiosError } from "axios";

type Field = "dni" | "email" | "password" | "role";

interface ErrorRule {
	// Fragmento del mensaje que devuelve el backend (errores de dominio, en español).
	backend: string;
	// Campo del formulario al que se asocia el error, si corresponde.
	field?: Field;
	// Clave de i18n con la que se muestra el mensaje.
	key: string;
}

const ERROR_RULES: ErrorRule[] = [
	{ backend: "DNI no existe", field: "dni", key: "formulario.errorDniNoEncontrado" },
	{ backend: "DNI ya está registrado", field: "dni", key: "usuarios.errorDniDuplicado" },
	{ backend: "no se pudo verificar el DNI", field: "dni", key: "usuarios.errorDniVerificacion" },
	{ backend: "DNI inválido", field: "dni", key: "formulario.errorDniInvalido" },
	{ backend: "email ya está registrado", field: "email", key: "usuarios.errorEmailDuplicado" },
	{ backend: "rol inválido", field: "role", key: "usuarios.errorRolInvalido" },
	{ backend: "contraseña", field: "password", key: "usuarios.errorPasswordCorta" },
];

export interface ResolvedApiError {
	field?: Field;
	message: string;
}

// getBackendError saca el mensaje de error del body que devuelve la API
// ({ code, message, error }). Axios guarda la respuesta en error.response,
// no en error.status ni en error.message.
export const getBackendError = (error: unknown): string => {
	const axiosError = error as AxiosError<{ error?: string; message?: string }>;
	const body = axiosError?.response?.data;

	if (body && typeof body.error === "string" && body.error) return body.error;
	if (body && typeof body.message === "string" && body.message) return body.message;

	return "";
};

// resolveApiError traduce el error del backend a un mensaje de i18n y lo
// asocia al campo correspondiente, para mostrarlo debajo del input en vez de
// dejar al usuario con un error genérico.
export const resolveApiError = (
	error: unknown,
	t: (key: string) => string,
	fallbackKey: string,
): ResolvedApiError => {
	const backendError = getBackendError(error);
	const rule = ERROR_RULES.find((item) => backendError.includes(item.backend));

	if (rule) return { field: rule.field, message: t(rule.key) };

	return { message: t(fallbackKey) };
};
