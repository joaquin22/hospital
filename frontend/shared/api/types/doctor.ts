export interface CreateDoctorPayload {
	first_name: string;
	last_name: string;
	email: string;
	password: string;
	role: string;
	dni: string;
	specialty_id: number;
	license_number: string;
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

export interface DoctorOutput {
	id: number;
	speciality_id: number;
	license_number: string;
}

export interface DoctorUserOutput {
	id: number;
	user_id?: number;
	first_name: string;
	last_name: string;
	email: string;
	role: string;
	dni: string;
	speciality_id: number;
	license_number: string;
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