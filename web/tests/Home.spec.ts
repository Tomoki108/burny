import { test, expect } from "@playwright/test";
import {
  WEB_LOCAL_HOST,
  pageClick,
  pageFill,
  pageTextContent,
} from "./test_helper";
import { mockAllApis } from "./test_mock";

test.describe("Home page", () => {
  test("User can sign up and sign in", async ({ page }) => {
    await mockAllApis(page);

    await page.goto(WEB_LOCAL_HOST);

    // sign up
    await pageClick(page, "signin-modal-button");
    await pageClick(page, "signup-tab");
    await pageFill(page, "email", "test@example.com");
    await pageFill(page, "password", "burnyburny");
    await pageClick(page, "auth-submit-button");
    const signUpSuccessMsg = await pageTextContent(page, "auth-success");
    expect(signUpSuccessMsg).toBe("Registration successful. Please sign in.");

    // sign in
    await pageClick(page, "auth-submit-button");
    await expect(page).toHaveURL(/\/projects$/);
  });
});
