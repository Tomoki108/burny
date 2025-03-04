import { defineStore } from "pinia";
import { fetchProjects, createProject, type Project } from "../api/project_api";

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
    getProjects(): Project[] {
      return this.projects;
    },
  },
});
