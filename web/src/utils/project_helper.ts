import { type Project } from "../api/project_api";

export function getEndDate(project: Project): string {
  const endDate = new Date(project.start_date);
  endDate.setDate(
    endDate.getDate() + project.sprint_count * 7 * project.sprint_duration
  );

  // yyyy-mm-dd形式に変換
  let str = endDate.toISOString().split("T")[0];

  // もしstart_dateと同じ年なら、月と日だけにする
  if (str.startsWith(project.start_date.split("-")[0])) {
    str = str.split("-").slice(1).join("-");
  }

  return str;
}
