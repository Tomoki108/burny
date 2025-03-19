import { test, expect } from "@playwright/test";
import {
  WEB_LOCAL_HOST,
  login,
  sleep,
  pageClick,
  pageFill,
} from "./test_helper";
import path from "path";
import { mkdir } from "fs/promises";
import { generateMockJWT } from "./mock_server";

// // Mock the Vite environment variables
// // @ts-ignore - Mocking import.meta.env
// import.meta = {
//   env: {
//     VITE_API_HOST: "http://localhost:1323/api/v1",
//     VITE_MOCK_API: "true"
//   }
// };

test.describe("Account page", () => {
  test("User can view their account information", async ({ page }) => {
    // Create screenshots directory if it doesn't exist
    const screenshotsDir = path.join(process.cwd(), "test-results/screenshots");
    await mkdir(screenshotsDir, { recursive: true });

    await page.goto(WEB_LOCAL_HOST);

    await page.route("**/api/v1/sign_in", (route) => {
      route.fulfill({
        status: 200,
        body: JSON.stringify({
          token: generateMockJWT(),
          user: {
            id: "1",
            email: "test@example.com",
          },
        }),
      });
    });

    await pageClick(page, "signin-modal-button");
    await pageFill(page, "email", "test@example.com");
    await pageFill(page, "password", "burnyburny");

    await page.screenshot({
      path: path.join(screenshotsDir, "login_input.png"),
      fullPage: true,
    });

    await pageClick(page, "auth-submit-button");

    await page.screenshot({
      path: path.join(screenshotsDir, "login_submit.png"),
      fullPage: true,
    });

    // プロジェクトページに正しく遷移したことを確認
    await expect(page).toHaveURL(/\/projects$/);

    await pageClick(page, "nav-account");

    // アカウントページに遷移したことを確認
    await page.waitForURL(/\/account$/, { timeout: 10000 });

    // スクリーンショットを撮る
    await page.screenshot({
      path: path.join(screenshotsDir, "5-account-page.png"),
      fullPage: true,
    });

    // アカウント情報が表示されていることを確認
    await expect(page.getByText("test@example.com")).toBeVisible();
    await expect(page.getByText("********")).toBeVisible();
  });
});
