import { http, HttpResponse, passthrough } from "msw";
import {
  type Project,
  type UpdateProjectRequest,
} from "../src/api/project_api";
import { setupWorker } from "msw/browser";
import { type Sprint, type UpdateSprintRequest } from "../src/api/sprint_api";

// NOTE: Do not import this module from spec files. Praywright somewhat stops working. (due to cyclic import?)

// NOTE: VITE_API_HOST env var is not loaded in the test environment
const API_HOST = "http://localhost:1323/api/v1";

const TEST_CREATE_PROJECT: Project = {
  id: 10,
  user_id: 1,
  title: "Test Project",
  sprint_count: 5,
  description: "This is a test project",
  sprint_duration: 1,
  start_date: "2024-01-01",
  total_sp: 100,
  created_at: "2024-01-01T00:00:00Z",
  updated_at: "2024-01-01T00:00:00Z",
};

const TEST_DEMO_PROJECT: Project = {
  id: 1,
  user_id: 1,
  title: "Demo Project",
  sprint_count: 5,
  description: "This is a demo project",
  sprint_duration: 1,
  start_date: "2024-01-01",
  total_sp: 100,
  created_at: "2024-01-01T00:00:00Z",
  updated_at: "2024-01-01T00:00:00Z",
};

const TEST_DEMO_PROJECT_SPRINTS: Sprint[] = [
  {
    id: 1,
    project_id: 1,
    user_id: 1,
    start_date: "2024-01-01",
    end_date: "2024-01-07",
    actual_sp: 20,
    ideal_sp: 20,
    created_at: "2024-01-01T00:00:00Z",
    updated_at: "2024-01-01T00:00:00Z",
  },
  {
    id: 2,
    project_id: 1,
    user_id: 1,
    start_date: "2024-01-08",
    end_date: "2024-01-14",
    actual_sp: 20,
    ideal_sp: 20,
    created_at: "2024-01-01T00:00:00Z",
    updated_at: "2024-01-01T00:00:00Z",
  },
  {
    id: 3,
    project_id: 1,
    user_id: 1,
    start_date: "2024-01-15",
    end_date: "2024-01-21",
    actual_sp: 20,
    ideal_sp: 20,
    created_at: "2024-01-01T00:00:00Z",
    updated_at: "2024-01-01T00:00:00Z",
  },
  {
    id: 4,
    project_id: 1,
    user_id: 1,
    start_date: "2024-01-22",
    end_date: "2024-01-28",
    actual_sp: 20,
    ideal_sp: 20,
    created_at: "2024-01-01T00:00:00Z",
    updated_at: "2024-01-01T00:00:00Z",
  },
  {
    id: 5,
    project_id: 1,
    user_id: 1,
    start_date: "2024-01-29",
    end_date: "2024-02-04",
    actual_sp: 20,
    ideal_sp: 20,
    created_at: "2024-01-01T00:00:00Z",
    updated_at: "2024-01-01T00:00:00Z",
  },
];

const handlers = [
  // auth api
  http.post(`${API_HOST}/sign_up`, () => {
    return HttpResponse.json(
      {
        message: "Registration successful. Please sign in.",
      },
      { status: 201 }
    );
  }),
  http.post(`${API_HOST}/sign_in`, () => {
    return HttpResponse.json(
      {
        token: generateMockJWT(),
        user: {
          id: "1",
          email: "test@example.com",
        },
      },
      { status: 200 }
    );
  }),
  // projects api
  http.get(`${API_HOST}/projects`, () => {
    return HttpResponse.json([TEST_DEMO_PROJECT], { status: 200 });
  }),
  http.post(`${API_HOST}/projects`, async () => {
    return HttpResponse.json(TEST_CREATE_PROJECT, { status: 201 });
  }),
  http.put(`${API_HOST}/projects/:id`, async (request) => {
    const json = await request.request.json();
    const updateReq = json?.valueOf();
    const req = updateReq as UpdateProjectRequest;

    // TEST_CREATE_PROJECTをコピーして更新
    const updatedProject = { ...TEST_CREATE_PROJECT };
    updatedProject.title = req.title;
    updatedProject.description = req.description;
    updatedProject.total_sp = req.total_sp;
    updatedProject.sprint_count = req.sprint_count;

    return HttpResponse.json(updatedProject, { status: 200 });
  }),
  http.delete(`${API_HOST}/projects/:id`, () => {
    return new HttpResponse(null, { status: 204 });
  }),
  http.get(`${API_HOST}/projects/:projectId/sprints`, ({ params }) => {
    const { projectId } = params;
    if (projectId === "1") {
      return HttpResponse.json(TEST_DEMO_PROJECT_SPRINTS, { status: 200 });
    }
    return HttpResponse.json([], { status: 200 });
  }),
  // sprints api
  http.patch(
    `${API_HOST}/projects/:projectId/sprints/:sprintId`,
    async (request) => {
      const json = await request.request.json();
      const updateReq = json?.valueOf();
      const req = updateReq as UpdateSprintRequest;
      const { sprintId } = request.params;

      const sprintIndex = TEST_DEMO_PROJECT_SPRINTS.findIndex(
        (sprint) => sprint.id.toString() === sprintId
      );
      if (sprintIndex === -1) {
        return HttpResponse.json(
          { message: "Sprint not found" },
          { status: 404 }
        );
      }

      const updatedSprint = { ...TEST_DEMO_PROJECT_SPRINTS[sprintIndex] };
      updatedSprint.actual_sp = req.actual_sp;
      updatedSprint.updated_at = new Date().toISOString();

      // Update the TEST_SPRINTS array to maintain state between requests
      TEST_DEMO_PROJECT_SPRINTS[sprintIndex] = updatedSprint;

      return HttpResponse.json(updatedSprint, { status: 200 });
    }
  ),
  // other settings
  // CSS, Vue, TypeScriptファイルを取得するリクエストをそのまま通す（unhandle requestの警告を消すため）
  http.get(new RegExp("\\.(css|vue|ts)$"), () => {
    passthrough();
  }),
];

export const worker = setupWorker(...handlers);

// デコード可能なモックJWTトークン
// ペイロード: { "sub": "1", "email": "test@example.com", "exp": 最新の時刻から1時間後 }
const generateMockJWT = () => {
  const header = btoa(JSON.stringify({ alg: "HS256", typ: "JWT" }));
  const now = Math.floor(Date.now() / 1000);
  const payload = btoa(
    JSON.stringify({
      sub: "1",
      email: "test@example.com",
      exp: now + 3600, // 1時間後
    })
  );
  const signature = btoa("mock-signature"); // 実際の署名は不要
  return `${header}.${payload}.${signature}`;
};
