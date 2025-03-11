import {
  type Sprint,
  type UpdateSprintRequest,
  fetchSprints,
  updateSprint,
} from "../api/sprint_api";
import { defineStore } from "pinia";
import { ErrorResponse } from "../api/api_helper";

export const useSprintsStore = defineStore("sprints", {
  state: () => ({
    sprints: [] as Sprint[],
  }),
  actions: {
    async fetchSprints(projectID: number) {
      const res = await fetchSprints(projectID);
      this.sprints = res;
    },
    getSprints(): Sprint[] {
      return this.sprints;
    },
    async updateSprint(sprint: Sprint) {
      const req: UpdateSprintRequest = {
        actual_sp: sprint.actual_sp,
      };
      const res = await updateSprint(sprint.project_id, sprint.id, req);
      if (res instanceof ErrorResponse) {
        throw new Error(res.getMessage());
      }
      const index = this.sprints.findIndex((s) => s.id === res.id);
      this.sprints[index] = res;
    },
  },
});
