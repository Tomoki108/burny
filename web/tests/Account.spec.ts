import { test, expect } from "@playwright/test";
import { WEB_LOCAL_HOST, login, pageClick } from "./test_helper";
import { mockAllApis } from "./test_mock";

test.describe("Account page", () => {
  test("User can view their account information", async ({ page }) => {
    await mockAllApis(page);

    await login(page);

    // Click on the account page
    await pageClick(page, "nav-drawer");
    await pageClick(page, "nav-account");
    await expect(page).toHaveURL(/\/account$/);

    // Check if the account information is displayed
    await expect(page.getByText("test@example.com")).toBeVisible();
    await expect(page.getByText("********")).toBeVisible();
  });
});
