import { test, expect } from "@playwright/test";
import { WEB_LOCAL_HOST, pageClick, pageFill, login } from "./test_helper";

test.describe("Projects page", () => {
  test.beforeEach(async ({ page }) => {
    await page.goto(WEB_LOCAL_HOST);
    await login(page);
  });

  test("User can create, update and delete project", async ({ page }) => {
    // create project
    await pageClick(page, "new-project-button");
    await page.getByLabel("Title").fill("Test Project");
    await page.getByLabel("Description").fill("This is a test project");
    await page.getByLabel("Total SP").fill("100");
    await page.getByLabel("Sprint Count").fill("5");
    await page.getByLabel("Start Date").fill("2024-01-01");
    await pageClick(page, "project-save");

    // assert created project
    const projectCard = page.getByTestId("project-card-1");
    await expect(projectCard.locator("h2")).toHaveText("Test Project");
    await expect(projectCard).toBeVisible();

    // update project
    await pageClick(page, "edit-project-button-1");
    await page.getByLabel("Title").fill("Updated Project");
    await page.getByLabel("Description").fill("This project has been updated");
    await page.getByLabel("Total SP").fill("150");
    await page.getByLabel("Sprint Count").fill("6");
    await pageClick(page, "project-save");

    // assert updated project
    await expect(projectCard.locator("h2")).toHaveText("Updated Project");

    // delete project
    await pageClick(page, "delete-project-button-1");
    await pageClick(page, "dialog-proceed");

    // assert deletion
    await expect(projectCard).not.toBeVisible();
  });
});
