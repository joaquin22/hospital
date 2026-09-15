export interface AuthCredentials {
	email: string;
	password: string;
}

export interface LoginData {
	token: string;
	expires_in: number;
}

export interface AuthResponse {
	code: number;
	message: string;
	data: LoginData;
}