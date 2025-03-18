import { http, HttpResponse, passthrough } from "msw";
import {
  type Project,
  type UpdateProjectRequest,
} from "../src/api/project_api";
import { setupWorker } from "msw/browser";

// NOTE: VITE_API_HOST env var is not loaded in the test environment
const API_HOST = "http://localhost:1323/api/v1";

export const TEST_CREATE_PROJECT: Project = {
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

const handlers = [
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

  http.get(`${API_HOST}/projects`, () => {
    return HttpResponse.json([], { status: 200 });
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
