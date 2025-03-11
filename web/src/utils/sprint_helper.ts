import { type Sprint } from "../api/sprint_api";

export const isSprintStarted = (sprint: Sprint): boolean => {
  const today = new Date();
  today.setHours(0, 0, 0, 0);
  const sprintStart = new Date(sprint.start_date);
  return sprintStart < today;
};
