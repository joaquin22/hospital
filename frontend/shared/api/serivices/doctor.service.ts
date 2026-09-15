import client from "../client";
import type {
  CreateDoctorPayload,
  CreateDoctorResponse,
  ListDoctorsResponse,
} from "../types/doctor";

export const doctorService = {
  create: (payload: CreateDoctorPayload): Promise<CreateDoctorResponse> => {
    return client.post("/doctors", payload);
  },
  list: async (): Promise<ListDoctorsResponse> => {
    const response = await client.get("/doctors");
    return response.data;
  },
};
