import { type Page } from "@playwright/test";

const API_BASE_URL = "/api";

export interface Project {
  id: number;
  user_id: number;
  title: string;
  sprint_count: number;
  description: string;
  sprint_duration: number;
  start_date: string;
  total_sp: number;
  created_at: string;
  updated_at: string;
}

export interface Sprint {
  id: number;
  project_id: number;
  user_id: number;
  start_date: string;
  end_date: string;
  actual_sp: number;
  ideal_sp: number;
  created_at: string;
  updated_at: string;
}

export interface UpdateProjectRequest {
  title: string;
  description: string;
  total_sp: number;
  sprint_count: number;
}

export interface UpdateSprintRequest {
  actual_sp: number;
}

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

export async function mockSignInApi(page: Page) {
  await page.route("**/api/v1/sign_in", (route) => {
    route.fulfill({
      status: 200,
      body: JSON.stringify({
        token: generateMockJWT(),
      }),
    });
  });
}

export async function mockSignUpApi(page: Page) {
  await page.route("**/api/v1/sign_up", (route) => {
    route.fulfill({
      status: 201,
      body: JSON.stringify({
        id: "1",
        email: "test@example.com",
        email_verified: false,
        created_at: new Date().toISOString(),
        updated_at: new Date().toISOString(),
      }),
    });
  });
}

export async function mockProjectsApi(page: Page) {
  await page.route("**/api/v1/projects", (route) => {
    const method = route.request().method();
    switch (method) {
      case "GET":
        route.fulfill({
          contentType: "application/json",
          status: 200,
          body: JSON.stringify([TEST_DEMO_PROJECT]),
        });
        break;
      case "POST":
        route.fulfill({
          contentType: "application/json",
          status: 201,
          body: JSON.stringify(TEST_CREATE_PROJECT),
        });
        break;
    }
  });
}

export async function mockProjectApi(page: Page) {
  await page.route("**/api/v1/projects/**", async (route) => {
    const method = route.request().method();
    switch (method) {
      case "PUT":
        const body = JSON.parse((await route.request().postData()) || "{}");
        const req = body as UpdateProjectRequest;

        // TEST_CREATE_PROJECTをコピーして更新
        const updatedProject = { ...TEST_CREATE_PROJECT };
        updatedProject.title = req.title;
        updatedProject.description = req.description;
        updatedProject.total_sp = req.total_sp;
        updatedProject.sprint_count = req.sprint_count;

        route.fulfill({
          contentType: "application/json",
          status: 200,
          body: JSON.stringify(updatedProject),
        });
        break;
      case "DELETE":
        route.fulfill({
          contentType: "application/json",
          status: 204,
        });
        break;
    }
  });
}

export async function mockSprintsApi(page: Page) {
  await page.route("**/api/v1/projects/*/sprints", (route) => {
    const method = route.request().method();
    switch (method) {
      case "GET":
        const url = route.request().url();
        const projectIdMatch = url.match(/\/projects\/(\d+)\/sprints/);
        const projectId = projectIdMatch ? projectIdMatch[1] : null;

        if (projectId === "1") {
          route.fulfill({
            contentType: "application/json",
            status: 200,
            body: JSON.stringify(TEST_DEMO_PROJECT_SPRINTS),
          });
        } else {
          route.fulfill({
            contentType: "application/json",
            status: 200,
            body: JSON.stringify([]),
          });
        }
        break;
    }
  });
}

export async function mockSprintApi(page: Page) {
  await page.route("**/api/v1/projects/*/sprints/*", async (route) => {
    const method = route.request().method();
    switch (method) {
      case "PATCH":
        const url = route.request().url();
        const sprintIdMatch = url.match(/\/sprints\/(\d+)$/);
        const sprintId = sprintIdMatch ? sprintIdMatch[1] : null;

        const body = JSON.parse((await route.request().postData()) || "{}");
        const req = body as UpdateSprintRequest;

        // スプリントを検索
        const sprintIndex = TEST_DEMO_PROJECT_SPRINTS.findIndex(
          (sprint) => sprint.id.toString() === sprintId
        );

        if (sprintIndex === -1) {
          route.fulfill({
            contentType: "application/json",
            status: 404,
            body: JSON.stringify({ message: "Sprint not found" }),
          });
          return;
        }

        // スプリントを更新
        const updatedSprint = { ...TEST_DEMO_PROJECT_SPRINTS[sprintIndex] };
        updatedSprint.actual_sp = req.actual_sp;
        updatedSprint.updated_at = new Date().toISOString();

        // 配列を更新して状態を維持
        TEST_DEMO_PROJECT_SPRINTS[sprintIndex] = updatedSprint;

        route.fulfill({
          contentType: "application/json",
          status: 200,
          body: JSON.stringify(updatedSprint),
        });
        break;
    }
  });
}

export async function mockApiKeyStatusApi(page: Page, exists: boolean = false) {
  await page.route("**/api/v1/apikeys/status", (route) => {
    route.fulfill({
      contentType: "application/json",
      status: 200,
      body: JSON.stringify({ exists }),
    });
  });
}

export async function mockApiKeyApi(page: Page) {
  await page.route("**/api/v1/apikeys", (route) => {
    const method = route.request().method();
    switch (method) {
      case "POST":
        route.fulfill({
          contentType: "application/json",
          status: 201,
          body: JSON.stringify({
            raw_key: "testapikey123456789abcdefghijklmn",
          }),
        });
        break;
      case "DELETE":
        route.fulfill({
          contentType: "application/json",
          status: 204,
        });
        break;
    }
  });
}

export async function mockAllApis(page: Page) {
  await mockSignInApi(page);
  await mockSignUpApi(page);
  await mockProjectsApi(page);
  await mockProjectApi(page);
  await mockSprintsApi(page);
  await mockSprintApi(page);
  await mockApiKeyStatusApi(page);
  await mockApiKeyApi(page);
}

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
