import { Page } from "playwright";

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
  
  // Clear localStorage to ensure clean state
  await page.evaluate(() => localStorage.clear());
  
  // Open sign-in modal
  await pageClick(page, "signin-modal-button");
  
  // Fill login credentials
  await pageFill(page, "email", "test@example.com");
  await pageFill(page, "password", "burnyburny");
  
  // Create a navigation promise before clicking
  const navigationPromise = page.waitForNavigation({ 
    url: /\/projects$/,
    timeout: 10000,
    waitUntil: 'networkidle' 
  });
  
  // Click login button
  await pageClick(page, "auth-submit-button");
  
  // Wait for navigation to complete
  await navigationPromise;
  
  // Wait for localStorage to be populated with token
  await page.waitForFunction(() => {
    return localStorage.getItem('token') !== null;
  }, { timeout: 5000 });
  
  // Verify we're on the projects page
  await page.waitForURL(/\/projects$/);
  
  // Wait for the page to be fully loaded
  await page.waitForLoadState('networkidle');
  
  // Wait a bit more to ensure everything is stabilized
  await sleep(500);
}
