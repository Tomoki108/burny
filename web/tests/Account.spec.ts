import { test, expect } from "@playwright/test";
import { WEB_LOCAL_HOST, login, pageClick, sleep } from "./test_helper";
import path from "path";
import { mkdir } from "fs/promises";

test.describe("Account page", () => {
  test("User can view their account information", async ({ page }) => {
    // Create screenshots directory if it doesn't exist
    const screenshotsDir = path.join(process.cwd(), "test-results/screenshots");
    await mkdir(screenshotsDir, { recursive: true });

    await page.goto(WEB_LOCAL_HOST);

    // Take screenshot after initial page load
    await page.screenshot({
      path: path.join(screenshotsDir, "1-home-page.png"),
      fullPage: true,
    });

    await login(page);

    // Take screenshot after login
    await page.screenshot({
      path: path.join(screenshotsDir, "2-after-login.png"),
      fullPage: true,
    });

    // Log DOM structure to help debug
    console.log("Page content after login:", await page.content());

    // Take screenshot specifically of the sidebar with nav elements
    const sidebarElement = page.locator(".sidebar");
    if (await sidebarElement.isVisible())
      await sidebarElement.screenshot({
        path: path.join(screenshotsDir, "3-sidebar.png"),
      });
    else console.log("Sidebar is not visible");

    // Make sure the navigation is fully loaded and visible before clicking
    await page.waitForSelector('[data-testid="nav-account"]', {
      state: "visible",
      timeout: 10000,
    });

    // Log info about the nav-account element
    const navAccountElement = page.locator('[data-testid="nav-account"]');
    const isVisible = await navAccountElement.isVisible();
    console.log("nav-account element visible:", isVisible);

    if (isVisible) {
      const boundingBox = await navAccountElement.boundingBox();
      console.log("nav-account position:", boundingBox);

      // Take screenshot with the element highlighted
      await page.evaluate((selector) => {
        const element = document.querySelector(selector);
        if (element && element instanceof HTMLElement) {
          element.style.border = "3px solid red";
        }
      }, '[data-testid="nav-account"]');

      await page.screenshot({
        path: path.join(screenshotsDir, "4-nav-account-highlighted.png"),
        fullPage: true,
      });
    }

    // Add a short delay to ensure any animations or state changes are complete
    await sleep(500);

    // Use force: true to ensure the click works even if something might be obscuring it
    await page.click('[data-testid="nav-account"]', { force: true });

    // Take screenshot after clicking
    await page.screenshot({
      path: path.join(screenshotsDir, "5-after-click.png"),
      fullPage: true,
    });

    await expect(page).toHaveURL(/\/account$/);

    // Check if the account information is displayed
    await expect(page.getByText("test@example.com")).toBeVisible();
    await expect(page.getByText("********")).toBeVisible();

    // Final screenshot showing the account page
    await page.screenshot({
      path: path.join(screenshotsDir, "6-account-page.png"),
      fullPage: true,
    });
  });
});
