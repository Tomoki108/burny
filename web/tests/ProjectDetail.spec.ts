import test from "@playwright/test";
import { WEB_LOCAL_HOST, pageClick, login } from "./test_helper";

test.describe("ProjectDetail page", () => {
  test("User can view project details and update sprints actual_sp", async ({
    page,
  }) => {
    await page.goto(WEB_LOCAL_HOST);
    await login(page);
  });
});
