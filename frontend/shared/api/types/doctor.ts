export interface CreateDoctorPayload {
	first_name: string;
	last_name: string;
	email: string;
	password: string;
	role: string;
	dni: string;
	// El backend lo expone como `speciality_id` (no `specialty_id`).
	speciality_id: number;
	license_number: string;
}

// El PATCH acepta cualquier subconjunto de campos: los que llegan en null no
// se tocan. La contraseña se deja fuera a propósito, se cambia por separado.
export type PatchDoctorPayload = Partial<
	Omit<CreateDoctorPayload, "password">
> & {
	password?: string;
	// El backend actualiza con este único campo las dos banderas `active`
	// (la de la ficha y la de la cuenta) y las deja sincronizadas.
	active?: boolean;
};

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

export interface DoctorOutput {
	id: number;
	speciality_id: number;
	license_number: string;
}

export interface DoctorUserOutput {
	id: number;
	user_id: number;
	first_name: string;
	last_name: string;
	email: string;
	role: string;
	dni: string;
	speciality_id: number;
	license_number: string;
	active: boolean;
}

export interface CreateDoctorResponse {
	code: number;
	message: string;
	data: {
		user: UserOutput;
		doctor: DoctorOutput;
	};
}

export interface ListDoctorsResponse {
	code: number;
	message: string;
	data: DoctorUserOutput[];
}

export interface GetDoctorResponse {
	code: number;
	message: string;
	data: DoctorUserOutput;
}

export interface UpdateDoctorResponse {
	code: number;
	message: string;
	data: DoctorUserOutput;
}
