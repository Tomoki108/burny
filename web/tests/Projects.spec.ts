import { test, expect } from "@playwright/test";
import { WEB_LOCAL_HOST, pageClick, login, sleep } from "./test_helper";
import { TEST_CREATE_PROJECT } from "./mock_server";

test.describe("Projects page", () => {
  test("User can create, update and delete project", async ({ page }) => {
    await page.goto(WEB_LOCAL_HOST);
    await login(page);

    // create project
    await pageClick(page, "new-project-button");
    await page.getByLabel("Title").fill(TEST_CREATE_PROJECT.title);
    await page.getByLabel("Description").fill(TEST_CREATE_PROJECT.description);
    await page
      .getByLabel("Total SP")
      .fill(String(TEST_CREATE_PROJECT.total_sp));
    await page
      .getByLabel("Sprint Count")
      .fill(String(TEST_CREATE_PROJECT.sprint_count));
    await page
      .getByLabel("Start Date")
      .fill(String(TEST_CREATE_PROJECT.start_date));
    await pageClick(page, "project-save");

    // assert created project
    const projectCard = page.getByTestId("project-card-10");
    await expect(projectCard.locator("h2")).toHaveText("Test Project");
    await expect(projectCard).toBeVisible();

    // update project
    await pageClick(page, "edit-project-button-10");
    await page.getByLabel("Title").fill("Updated Project");
    await page.getByLabel("Description").fill("This project has been updated");
    await page.getByLabel("Total SP").fill("150");
    await page.getByLabel("Sprint Count").fill("6");
    await pageClick(page, "project-save");

    // assert updated project
    await expect(projectCard.locator("h2")).toHaveText("Updated Project");

    // delete project
    await pageClick(page, "delete-project-button-10");
    await pageClick(page, "dialog-proceed");

    // assert deletion
    await expect(projectCard).not.toBeVisible();
  });
});
