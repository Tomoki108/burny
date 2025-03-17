import { test, expect } from "@playwright/test";
import { WEB_LOCAL_HOST, pageClick, pageFill, login } from "./test_helper";

test.describe("Projects page", () => {
  test.beforeEach(async ({ page }) => {
    await page.goto(WEB_LOCAL_HOST);
    await login(page);
  });

  test("User can create a new project", async ({ page }) => {
    await pageClick(page, "new-project-button");

    // Fill project form
    await page.getByLabel("Title").fill("Test Project");
    await page.getByLabel("Description").fill("This is a test project");
    await page.getByLabel("Total SP").fill("100");
    await page.getByLabel("Sprint Count").fill("5");
    // vuetify combo box is somewhat not testable with playwright, so to use default value
    // await page.getByLabel("Sprint Weeks").selectOption("1");
    await page.getByLabel("Start Date").fill("2024-01-01");

    // Save project
    await pageClick(page, "project-save");

    // Verify project is created and visible in the list
    const projectCard = page.locator('[data-testid="project-card-1"]');
    await expect(projectCard).toBeVisible();
    await expect(projectCard.locator("h2")).toHaveText("Test Project");
  });

  //   test("User can edit an existing project", async ({ page }) => {
  //     // Click edit button on the first project
  //     await pageClick(page, "edit-project-button");

  //     // Update project details
  //     await pageFill(page, "project-title", "Updated Project");
  //     await pageFill(
  //       page,
  //       "project-description",
  //       "This project has been updated"
  //     );
  //     await pageFill(page, "project-total-sp", "150");
  //     await pageFill(page, "project-sprint-count", "6");

  //     // Save changes
  //     await pageClick(page, "project-save");

  //     // Verify project is updated
  //     const projectCard = page.locator('[data-testid="project-card-1"]');
  //     await expect(projectCard).toBeVisible();
  //     await expect(projectCard.locator("h2")).toHaveText("Updated Project");
  //   });

  //   test("Projects are listed correctly", async ({ page }) => {
  //     // Verify project cards are visible
  //     const projectCards = page.locator(".project-card");
  //     await expect(projectCards).toHaveCount(1); // Assuming one project exists from previous test

  //     // Verify new project button is visible
  //     const newProjectButton = page.locator('[data-testid="new-project-button"]');
  //     await expect(newProjectButton).toBeVisible();
  //   });
});
