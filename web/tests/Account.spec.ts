import { test, expect } from "@playwright/test";
import { WEB_LOCAL_HOST, login, pageClick, sleep } from "./test_helper";

test.describe("Account page", () => {
  test("User can view their account information", async ({ page }) => {
    await page.goto(WEB_LOCAL_HOST);
    await login(page);

    // Make sure the navigation is fully loaded and visible before clicking
    await page.waitForSelector('[data-testid="nav-account"]', {
      state: "visible",
      timeout: 10000,
    });
    // Add a short delay to ensure any animations or state changes are complete
    await sleep(500);

    // Click on the account page
    await pageClick(page, "nav-account");
    await expect(page).toHaveURL(/\/account$/);

    // Check if the account information is displayed
    await expect(page.getByText("test@example.com")).toBeVisible();
    await expect(page.getByText("********")).toBeVisible();
  });
});
