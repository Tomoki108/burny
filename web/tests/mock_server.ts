import { http, HttpResponse } from "msw";
import { setupWorker } from "msw/browser";

const API_HOST = "http://localhost:1323/api/v1";

export const handlers = [
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
