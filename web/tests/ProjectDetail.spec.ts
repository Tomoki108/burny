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

    // Navigate to project detail page by clicking on the project card
    await page.getByText("Demo Project").first().click();
    await expect(page).toHaveURL(/\/projects\/1$/);

    // Check that project details are displayed correctly
    await expect(page.getByText("Projects > Demo Project")).toBeVisible();
    await expect(
      page.getByText("2024-01-01 to 02-05, 5 sprints, 100 sp")
    ).toBeVisible();
    await expect(page.getByText("This is a demo project")).toBeVisible();

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
