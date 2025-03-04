import { API_HOST } from "../config";
import { getAuthHeader } from "./helper";

export interface Project {
  id: number;
  user_id: number;
  title: string;
  sprint_count: number;
  description: string;
  sprint_duration: number;
  start_date: string;
  total_sp: number;
  created_at: string;
  updated_at: string;
}

export const defaultProject: Project = {
  id: 0,
  user_id: 0,
  title: "",
  sprint_count: 0,
  description: "",
  sprint_duration: 1,
  start_date: "",
  total_sp: 0,
  created_at: "",
  updated_at: "",
};

export async function fetchProjects(): Promise<Project[]> {
  const response = await fetch(`${API_HOST}/projects`, {
    headers: getAuthHeader(),
  });
  if (!response.ok) {
    throw new Error("Failed to fetch projects");
  }
  return await response.json();
}

export async function createProject(project: Project): Promise<Project> {
  const response = await fetch(`${API_HOST}/projects`, {
    method: "POST",
    headers: getAuthHeader(),
    body: JSON.stringify(project),
  });
  if (!response.ok) {
    throw new Error("Failed to create project");
  }
  return await response.json();
}

export async function fetchProject(id: number): Promise<Project> {
  const response = await fetch(`${API_HOST}/projects/${id}`, {
    headers: getAuthHeader(),
  });
  if (!response.ok) {
    throw new Error("Failed to fetch project");
  }
  return await response.json();
}

export async function updateProjects(project: Project): Promise<Project> {
  const response = await fetch(`${API_HOST}/projects/${project.id}`, {
    method: "PUT",
    headers: getAuthHeader(),
    body: JSON.stringify(project),
  });
  if (!response.ok) {
    throw new Error("Failed to update project");
  }
  return await response.json();
}

export async function deleteProjects(project: Project): Promise<void> {
  const response = await fetch(`${API_HOST}/projects/${project.id}`, {
    method: "DELETE",
    headers: getAuthHeader(),
  });
  if (!response.ok) {
    throw new Error("Failed to delete project");
  }
}
