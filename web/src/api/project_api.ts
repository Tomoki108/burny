import { API_HOST } from "../config";
import {
  getAuthHeader,
  replaceDateWithISOString,
  ErrorResponse,
} from "./helper";

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
  return await response.json();
}

export async function createProject(
  project: Project
): Promise<Project | ErrorResponse> {
  const response = await fetch(`${API_HOST}/projects`, {
    method: "POST",
    headers: getAuthHeader(),
    body: JSON.stringify(project, replaceDateWithISOString),
  });
  if (!response.ok) {
    return Object.assign(new ErrorResponse(), response.json());
  }
  return await response.json();
}

export async function fetchProject(
  id: number
): Promise<Project | ErrorResponse> {
  const response = await fetch(`${API_HOST}/projects/${id}`, {
    headers: getAuthHeader(),
  });
  if (!response.ok) {
    return Object.assign(new ErrorResponse(), response.json());
  }
  return await response.json();
}

type updateProjectRequest = {
  title: string;
  description: string;
  total_sp: number;
  sprint_count: number;
};

export async function updateProject(
  project: Project
): Promise<Project | ErrorResponse> {
  const req: updateProjectRequest = {
    title: project.title,
    description: project.description,
    total_sp: project.total_sp,
    sprint_count: project.sprint_count,
  };
  const response = await fetch(`${API_HOST}/projects/${project.id}`, {
    method: "PUT",
    headers: getAuthHeader(),
    body: JSON.stringify(req),
  });
  if (!response.ok) {
    return Object.assign(new ErrorResponse(), response.json());
  }

  return await response.json();
}

export async function deleteProject(id: number): Promise<void | ErrorResponse> {
  const response = await fetch(`${API_HOST}/projects/${id}`, {
    method: "DELETE",
    headers: getAuthHeader(),
  });
  if (!response.ok) {
    return Object.assign(new ErrorResponse(), response.json());
  }
}
