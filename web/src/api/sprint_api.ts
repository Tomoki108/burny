import { API_HOST } from "../config";
import { getAuthHeader, ErrorResponse } from "./api_helper";

export interface Sprint {
  id: number;
  project_id: number;
  user_id: number;
  start_date: string;
  end_date: string;
  actual_sp: number;
  ideal_sp: number;
  created_at: string;
  updated_at: string;
}

export async function fetchSprints(projectID: number): Promise<Sprint[]> {
  const response = await fetch(`${API_HOST}/projects/${projectID}/sprints`, {
    headers: getAuthHeader(),
  });
  return await response.json();
}

export interface UpdateSprintRequest {
  actual_sp: number;
}

export async function updateSprint(
  projectID: number,
  sprintID: number,
  req: UpdateSprintRequest
): Promise<Sprint | ErrorResponse> {
  const response = await fetch(
    `${API_HOST}/projects/${projectID}/sprints/${sprintID}`,
    {
      method: "PATCH",
      headers: getAuthHeader(),
      body: JSON.stringify(req),
    }
  );
  if (!response.ok) {
    const errorData = await response.json();
    return Object.assign(new ErrorResponse(), errorData);
  }
  return await response.json();
}
