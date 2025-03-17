import { test, expect, Page } from "@playwright/test";

test("User can sign up and sign in", async ({ page }) => {
  await page.goto("localhost:5179");

  await pageClick(page, "signin-modal-button");
  await pageClick(page, "signup-tab");
  await pageFill(page, "email", "test@example.com");
  await pageFill(page, "password", "burnyburny");
  await pageClick(page, "auth-submit-button");

  const message = await pageTextContent(page, "auth-success");
  expect(message).toBe("Registration successful. Please sign in.");
});

async function pageFill(page: Page, dataTestId: string, value: string) {
  return page.fill(`[data-testid="${dataTestId}"]`, value);
}

async function pageClick(page: Page, dataTestId: string) {
  return page.click(`[data-testid="${dataTestId}"]`);
}

async function pageTextContent(page: Page, dataTestId: string) {
  return page.textContent(`[data-testid="${dataTestId}"]`);
}
