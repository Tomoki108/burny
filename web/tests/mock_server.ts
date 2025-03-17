import { http, HttpResponse } from "msw";
import { setupWorker } from "msw/browser";
import { API_HOST } from "../src/config";

export const handlers = [
  http.post(`${API_HOST}/sign_up`, () => {
    return HttpResponse.json(
      {
        message: "Registration successful. Please sign in.",
      },
      { status: 201 }
    );
  }),
  http.post(`${API_HOST}/api/v1/sign_in`, () => {
    return HttpResponse.json(
      {
        token: "mock-jwt-token",
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
