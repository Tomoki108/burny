import { defineStore } from "pinia";
import {
  fetchProjects,
  createProject,
  deleteProject,
  type Project,
  updateProject,
} from "../api/project_api";

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
      try {
        const updatedProject = await updateProject(project);
        const index = this.projects.findIndex(
          (p) => p.id === updatedProject.id
        );
        this.projects[index] = updatedProject;
      } catch (error) {
        console.error("Error updating project:", error);
      }
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
