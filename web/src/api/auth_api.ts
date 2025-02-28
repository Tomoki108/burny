import { API_HOST } from "../config";

export interface SignInResponse {
  token: string;
}

export interface SignUpResponse {
  id: number;
  email: string;
  created_at: string;
  updated_at: string;
}

export async function signIn(
  email: string,
  password: string
): Promise<SignInResponse> {
  const response = await fetch(`${API_HOST}/sign_in`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ email, password }),
  });

  if (!response.ok) {
    throw new Error("Login failed");
  }

  return await response.json();
}

export async function signUp(
  email: string,
  password: string
): Promise<SignUpResponse> {
  const response = await fetch(`${API_HOST}/sign_up`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ email, password }),
  });

  if (!response.ok) {
    throw new Error("Sign up failed");
  }

  return await response.json();
}
