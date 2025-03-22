import { getAuthHeader, ErrorResponse, API_BASE_URL } from "./api_helper";

interface RawAPIKey {
  raw_key: string;
}

interface APIKeyStatus {
  exists: boolean;
}

export async function checkAPIKeyStatus(): Promise<APIKeyStatus> {
  const response = await fetch(`${API_BASE_URL}/apikeys/status`, {
    headers: getAuthHeader(),
  });
  return await response.json();
}

export async function createAPIKey(): Promise<RawAPIKey | ErrorResponse> {
  const response = await fetch(`${API_BASE_URL}/apikeys`, {
    method: "POST",
    headers: getAuthHeader(),
  });
  if (!response.ok) {
    const errorData = await response.json();
    return Object.assign(new ErrorResponse(), errorData);
  }
  return await response.json();
}

export async function deleteAPIKey(): Promise<void | ErrorResponse> {
  const response = await fetch(`${API_BASE_URL}/apikeys`, {
    method: "DELETE",
    headers: getAuthHeader(),
  });
  if (!response.ok) {
    const errorData = await response.json();
    return Object.assign(new ErrorResponse(), errorData);
  }
  return await response.json();
}
