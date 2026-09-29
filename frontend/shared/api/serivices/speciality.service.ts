import client from "../client";
import type {
  CreateSpecialityPayload,
  CreateSpecialityResponse,
  GetSpecialityResponse,
  ListSpecialitiesResponse,
  PatchSpecialityPayload,
  UpdateSpecialityResponse,
} from "../types/specialty";

// El grupo de rutas del backend es `/specialty` (en singular), a diferencia de
// `/doctors` y `/users`.
export const specialityService = {
  create: async (
    payload: CreateSpecialityPayload,
  ): Promise<CreateSpecialityResponse> => {
    const response = await client.post<CreateSpecialityResponse>(
      "/specialty",
      payload,
    );
    return response.data;
  },
  list: async (): Promise<ListSpecialitiesResponse> => {
    const response = await client.get<ListSpecialitiesResponse>("/specialty");
    return response.data;
  },
  get: async (id: number): Promise<GetSpecialityResponse> => {
    const response = await client.get<GetSpecialityResponse>(
      `/specialty/${id}`,
    );
    return response.data;
  },
  update: async (
    id: number,
    payload: PatchSpecialityPayload,
  ): Promise<UpdateSpecialityResponse> => {
    const response = await client.patch<UpdateSpecialityResponse>(
      `/specialty/${id}`,
      payload,
    );
    return response.data;
  },
};
