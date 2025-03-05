import { defineStore } from "pinia";
import {
  fetchProjects,
  createProject,
  deleteProject,
  type Project,
  updateProject,
} from "../api/project_api";
import { ErrorResponse } from "../api/helper";

export const useProjectsStore = defineStore("projects", {
  state: () => ({
    projects: [] as Project[],
  }),
  actions: {
    async fetchProjects() {
      const res = await fetchProjects();
      this.projects = res;
    },
    async createProject(project: Project) {
      const res = await createProject(project);
      if (res instanceof ErrorResponse) {
        throw new Error(res.getMessage());
      }
      this.projects.push(res);
    },
    async updateProject(project: Project) {
      const res = await updateProject(project);
      if (res instanceof ErrorResponse) {
        throw new Error(res.getMessage());
      }
      const index = this.projects.findIndex((p) => p.id === res.id);
      this.projects[index] = res;
    },
    async deleteProject(id: number) {
      const res = await deleteProject(id);
      if (res instanceof ErrorResponse) {
        throw new Error(res.getMessage());
      }
      this.projects = this.projects.filter((p) => p.id !== id);
    },
    getProjects(): Project[] {
      return this.projects;
    },
  },
});
