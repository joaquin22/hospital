import client from "../client";
import type {
  CreateDoctorPayload,
  CreateDoctorResponse,
  GetDoctorResponse,
  ListDoctorsResponse,
  PatchDoctorPayload,
  UpdateDoctorResponse,
} from "../types/doctor";

export const doctorService = {
  create: async (payload: CreateDoctorPayload): Promise<CreateDoctorResponse> => {
    const response = await client.post<CreateDoctorResponse>(
      "/doctors",
      payload,
    );
    return response.data;
  },
  list: async (): Promise<ListDoctorsResponse> => {
    const response = await client.get<ListDoctorsResponse>("/doctors");
    return response.data;
  },
  get: async (id: number): Promise<GetDoctorResponse> => {
    const response = await client.get<GetDoctorResponse>(`/doctors/${id}`);
    return response.data;
  },
  update: async (
    id: number,
    payload: PatchDoctorPayload,
  ): Promise<UpdateDoctorResponse> => {
    const response = await client.patch<UpdateDoctorResponse>(
      `/doctors/${id}`,
      payload,
    );
    return response.data;
  },
};
