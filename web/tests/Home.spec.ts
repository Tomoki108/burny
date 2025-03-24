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
    const signUpInfoMsg = await pageTextContent(page, "auth-info");
    expect(signUpInfoMsg).toBe(
      "Verification email sent to your email address. Please check your inbox."
    );

    // verify email
    await page.goto(WEB_LOCAL_HOST + "?email_verified=true");
    const emailVerifiedMsg = await pageTextContent(page, "auth-success");
    expect(emailVerifiedMsg).toBe(
      "Email verified successfully! You can now sign in."
    );

    // sign in
    await pageFill(page, "email", "test@example.com");
    await pageFill(page, "password", "burnyburny");
    await pageClick(page, "auth-submit-button");
    await expect(page).toHaveURL(/\/projects$/);
  });
});
