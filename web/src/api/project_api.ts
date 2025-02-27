import { API_HOST } from "../config";
import { getAuthHeader } from "./helper";

export interface Project {
  id: number;
  title: string;
  sprint_count: number;
  description: string;
  sprint_duration: number;
  start_date: string;
  total_sp: number;
  created_at: string;
  updated_at: string;
  user_id: number;
}

export async function fetchProjects(): Promise<Project[]> {
  const response = await fetch(`${API_HOST}/projects`, {
    headers: getAuthHeader(),
  });
  if (!response.ok) {
    throw new Error("Failed to fetch projects");
  }
  return await response.json();
}
