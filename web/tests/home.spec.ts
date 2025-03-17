import { test, expect, Page } from "@playwright/test";
import {
  WEB_LOCAL_HOST,
  pageClick,
  pageFill,
  pageTextContent,
} from "./test_helper";

test("User can sign up and sign in", async ({ page }) => {
  await page.goto(WEB_LOCAL_HOST);

  await pageClick(page, "signin-modal-button");
  await pageClick(page, "signup-tab");
  await pageFill(page, "email", "test@example.com");
  await pageFill(page, "password", "burnyburny");
  await pageClick(page, "auth-submit-button");

  const message = await pageTextContent(page, "auth-success");
  expect(message).toBe("Registration successful. Please sign in.");
});
