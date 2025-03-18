import { test, expect } from "@playwright/test";
import {
  WEB_LOCAL_HOST,
  pageClick,
  login,
  pageFill,
  sleep,
} from "./test_helper";

test.describe("ProjectDetail page", () => {
  test("User can view project details and update sprints actual_sp", async ({
    page,
  }) => {
    await page.goto(WEB_LOCAL_HOST);
    await login(page);

    // Create a project first (since we need a project to view details)
    await pageClick(page, "new-project-button");
    await page.getByLabel("Title").fill("Test Project");
    await page.getByLabel("Description").fill("This is a test project");
    await page.getByLabel("Total SP").fill("100");
    await page.getByLabel("Sprint Count").fill("5");
    await page.getByLabel("Start Date").fill("2024-01-01");
    await pageClick(page, "project-save");

    // Navigate to project detail page by clicking on the project card
    await page.getByText("Test Project").first().click();
    await expect(page).toHaveURL(/\/projects\/10$/);

    // Check that project details are displayed correctly
    await expect(page.getByText("Projects > Test Project")).toBeVisible();
    await expect(
      page.getByText("2024-01-01 to 02-05, 5 sprints, 100 sp")
    ).toBeVisible();
    await expect(page.getByText("This is a test project")).toBeVisible();

    // Verify that the first sprint shows actual_sp value from our mock
    const firstSprintRow = page
      .getByRole("row")
      .filter({ hasText: "Sprint 1" });
    await expect(firstSprintRow.getByTestId("actual_sp")).toHaveText("20");

    // Find an Update button for a sprint and click it
    await page.getByTestId("update-sprint-button").first().click();
    await sleep(1000);

    // Check if the modal is displayed
    await expect(page.getByTestId("sprint-modal-title")).toHaveText(
      "Update Sprint 1"
    );

    // Update the actual_sp in the modal
    await page.getByLabel("actual_sp").fill("25");
    await pageClick(page, "sprint-update-button");

    // Verify the update was successful
    await expect(page.getByText("Sprint updated successfully")).toBeVisible();
    await expect(firstSprintRow.getByTestId("actual_sp")).toHaveText("25");
  });
});
