export interface CreateSpecialityPayload {
	name: string;
	description: string;
}

// El PATCH solo aplica los campos que llegan; los ausentes se conservan.
export type PatchSpecialityPayload = Partial<CreateSpecialityPayload>;

export interface SpecialityOutput {
	id: number;
	name: string;
	description: string;
	active: boolean;
	created_at: string;
	updated_at: string;
}

export interface CreateSpecialityResponse {
	code: number;
	message: string;
	data: SpecialityOutput;
}

export interface ListSpecialitiesResponse {
	code: number;
	message: string;
	data: SpecialityOutput[];
}

export interface GetSpecialityResponse {
	code: number;
	message: string;
	data: SpecialityOutput;
}

export interface UpdateSpecialityResponse {
	code: number;
	message: string;
	data: SpecialityOutput;
}
