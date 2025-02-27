import { defineStore } from "pinia";
import { fetchProjects, type Project } from "../api/project_api";

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
    getProjects(): Project[] {
      return this.projects;
    },
  },
});
