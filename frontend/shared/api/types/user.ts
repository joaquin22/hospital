export interface CreateUserPayload {
	first_name: string;
	last_name: string;
	email: string;
	password: string;
	dni: string;
	role: string;
}

export type PatchUserPayload = Partial<
  Omit<CreateUserPayload, "password">
>;

export interface ChangePasswordPayload {
	password: string;
}

export interface CreateUserResponse {
	code: number;
	message: string;
	data: UserOutput;
}

export interface GetUserResponse {
	code: number;
	message: string;
	data: UserOutput;
}

export interface UpdateUserResponse {
	code: number;
	message: string;
	data: UserOutput;
}

export interface UserOutput {
	id: number;
	first_name: string;
	last_name: string;
	email: string;
	dni: string;
	role: string;
	active: boolean;
	created_at: string;
	updated_at: string;
}

export interface ListUsersResponse {
	code: number;
	message: string;
	data: UserOutput[];
}