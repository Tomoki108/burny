import { ErrorResponse, API_BASE_URL } from "./api_helper";

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
): Promise<SignInResponse | ErrorResponse> {
  const response = await fetch(`${API_BASE_URL}/sign_in`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ email, password }),
  });

  if (!response.ok) {
    const errorData = await response.json();
    return Object.assign(new ErrorResponse(), errorData);
  }

  return await response.json();
}

export async function signUp(
  email: string,
  password: string
): Promise<SignUpResponse | ErrorResponse> {
  const response = await fetch(`${API_BASE_URL}/sign_up`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ email, password }),
  });

  if (!response.ok) {
    const errorData = await response.json();
    return Object.assign(new ErrorResponse(), errorData);
  }

  return await response.json();
}
