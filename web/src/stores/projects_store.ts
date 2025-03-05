import { defineStore } from "pinia";
import {
  fetchProjects,
  createProject,
  deleteProject,
  type Project,
  updateProject,
} from "../api/project_api";
import { ErrorResponse, isErrorResponse } from "../api/helper";

export const useProjectsStore = defineStore("projects", {
  state: () => ({
    projects: [] as Project[],
  }),
  actions: {
    async fetchProjects() {
      try {
        this.projects = await fetchProjects();
      } catch (error) {
        console.error("Error fetching projects:", error);
      }
    },
    async createProject(project: Project) {
      try {
        this.projects.push(await createProject(project));
      } catch (error) {
        console.error("Error creating project:", error);
      }
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
      try {
        this.projects = this.projects.filter((p) => p.id !== id);
        await deleteProject(id);
      } catch (error) {
        console.error("Error deleting project:", error);
      }
    },

    getProjects(): Project[] {
      return this.projects;
    },
  },
});
