import { Page } from "playwright";

export const WEB_LOCAL_HOST = "http://localhost:5179"; // vite.config.tsのserver.portと合わせてある

export async function pageFill(page: Page, dataTestId: string, value: string) {
  return page.fill(`[data-testid="${dataTestId}"]`, value);
}

export async function pageClick(page: Page, dataTestId: string) {
  return page.click(`[data-testid="${dataTestId}"]`);
}

export async function pageTextContent(page: Page, dataTestId: string) {
  return page.textContent(`[data-testid="${dataTestId}"]`);
}

export const sleep = (ms: number) => new Promise((res) => setTimeout(res, ms));

export async function login(page: Page) {
  await page.goto(WEB_LOCAL_HOST);
  await pageClick(page, "signin-modal-button");
  await pageFill(page, "email", "test@example.com");
  await pageFill(page, "password", "burnyburny");
  await pageClick(page, "auth-submit-button");
}
