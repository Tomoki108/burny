import { type Page } from "@playwright/test";
import path from "path";
import { mkdir } from "fs/promises";

export const WEB_LOCAL_HOST = "http://localhost:5179"; // vite.config.tsのserver.portと合わせてある

export async function pageFill(page: Page, dataTestId: string, value: string) {
  return page.fill(`[data-testid="${dataTestId}"]`, value);
}

export async function pageClick(page: Page, dataTestId: string) {
  // For dialog buttons and other potentially problematic elements, add extra wait and force option
  if (dataTestId.includes("dialog")) {
    // Wait briefly for any animations to complete
    await sleep(500);
    return page.click(`[data-testid="${dataTestId}"]`, { force: true });
  }
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

// for debug use, can pass to page.screenshot()
export async function createScreenShotDir(testName: string) {
  const screenshotsDir = path.join(
    process.cwd(),
    `test-results/screenshots/${testName}`
  );
  await mkdir(screenshotsDir, { recursive: true });

  return screenshotsDir;
}
