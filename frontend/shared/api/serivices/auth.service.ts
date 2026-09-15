import client from "../client";
import type { AuthCredentials, AuthResponse } from "../types/auth";

export const authService = {
  login: async (credentials: AuthCredentials): Promise<AuthResponse> => {
    const response = await client.post<AuthResponse>(
      "/auth/login",
      credentials,
    );
    const data = response.data;
    return data;
  },
};
