import { http, HttpResponse, passthrough } from "msw";
import { setupWorker } from "msw/browser";
import type { Project } from "../src/api/project_api";

const API_HOST = "http://localhost:1323/api/v1";

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
  http.post(`${API_HOST}/projects`, async (request) => {
    const json = await request.request.json();
    const pj = json?.valueOf();
    const project = pj as Project;

    return HttpResponse.json(
      {
        id: 1,
        user_id: project.user_id,
        title: project.title,
        sprint_count: project.sprint_count,
        description: project.description,
        sprint_duration: project.sprint_duration,
        start_date: project.start_date,
        total_sp: project.total_sp,
        created_at: "2024-01-01T00:00:00Z",
        updated_at: "2024-01-01T00:00:00Z",
      },
      { status: 201 }
    );
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
