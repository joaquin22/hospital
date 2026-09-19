import client from "../client";
import type {
  ChangePasswordPayload,
  CreateUserPayload,
  CreateUserResponse,
  GetUserResponse,
  ListUsersResponse,
  PatchUserPayload,
  UpdateUserResponse,
} from "../types/user";

export const userService = {
  create: async (payload: CreateUserPayload): Promise<CreateUserResponse> => {
    const response = await client.post<CreateUserResponse>(
      "/auth/register",
      payload,
    );
    return response.data;
  },
  list: async (): Promise<ListUsersResponse> => {
    const response = await client.get("/users");
    return response.data;
  },
  get: async (id: number): Promise<GetUserResponse> => {
    const response = await client.get<GetUserResponse>(`/users/${id}`);
    return response.data;
  },
  update: async (
    id: number,
    payload: PatchUserPayload,
  ): Promise<UpdateUserResponse> => {
    const response = await client.patch<UpdateUserResponse>(
      `/users/${id}`,
      payload,
    );
    return response.data;
  },
  changePassword: async (
    id: number,
    payload: ChangePasswordPayload,
  ): Promise<UpdateUserResponse> => {
    const response = await client.patch<UpdateUserResponse>(
      `/users/${id}`,
      payload,
    );
    return response.data;
  },
};