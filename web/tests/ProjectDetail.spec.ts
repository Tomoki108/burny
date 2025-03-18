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
      page.getByText("2024-01-01 to", { exact: false })
    ).toBeVisible();
    await expect(page.getByText("5 sprints, 100 sp")).toBeVisible();
    await expect(page.getByText("This is a test project")).toBeVisible();

    // Check sprint table headers using the exact text from the Vue template
    await expect(page.getByText("sprint_no")).toBeVisible();
    await expect(page.getByText("start_date")).toBeVisible();
    await expect(page.getByText("end_date")).toBeVisible();
    await expect(page.getByText("ideal_sp")).toBeVisible();
    await expect(page.getByText("actual_sp")).toBeVisible();

    // Check if sprint rows are displayed correctly
    await expect(page.getByText("Sprint 1")).toBeVisible();
    await expect(page.getByText("Sprint 5")).toBeVisible();

    // Use getByRole to specifically find the cell containing the date
    await expect(page.getByRole("cell", { name: "2024-01-01" })).toBeVisible(); // First sprint start date

    // Verify that the first sprint shows actual_sp value from our mock
    const firstSprintRow = page
      .getByRole("row")
      .filter({ hasText: "Sprint 1" });
    await expect(firstSprintRow.getByText("20")).toBeVisible(); // actual_sp value for the first sprint

    // Find an Update button for a sprint and click it
    await page.getByText("Update").first().click();

    // Check if the modal is displayed
    await expect(page.getByText("Update Sprint")).toBeVisible();

    // Update the actual_sp in the modal
    await page.getByLabel("Actual SP").fill("25");
    await pageClick(page, "sprint-update-submit");

    // Verify the update was successful
    await expect(page.getByText("Sprint updated successfully")).toBeVisible();

    // Verify that the actual_sp was updated in the table
    await sleep(500); // Small wait to ensure UI updates
    await expect(firstSprintRow.getByText("25")).toBeVisible();

    // Check that the chart is displayed
    await expect(page.locator("canvas")).toBeVisible();

    // Check Basic Info card
    const basicInfoCard = page.getByText("Basic Info").first();
    await expect(basicInfoCard).toBeVisible();

    // Go back to the projects list
    await pageClick(page, "nav-projects");
    await expect(page).toHaveURL(/\/projects$/);
  });
});
