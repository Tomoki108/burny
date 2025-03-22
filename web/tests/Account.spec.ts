import { test, expect } from "@playwright/test";
import { WEB_LOCAL_HOST, login, pageClick, sleep } from "./test_helper";
import { mockAllApis, mockApiKeyStatusApi } from "./test_mock";

test.describe("Account page", () => {
  test("User can view their account information", async ({ page }) => {
    await mockAllApis(page);

    await login(page);

    // Click on the account page
    await pageClick(page, "nav-drawer");
    await pageClick(page, "nav-account");
    await expect(page).toHaveURL(/\/account$/);

    // Check if the account information is displayed
    await expect(page.getByText("test@example.com")).toBeVisible();
    await expect(page.getByText("********")).toBeVisible();
  });

  test("User can create and delete API key", async ({ page, context }) => {
    await context.grantPermissions(["clipboard-read", "clipboard-write"]);
    await mockApiKeyStatusApi(page, false);
    await mockAllApis(page);

    await login(page);

    // Click on the account page
    await pageClick(page, "nav-drawer");
    await pageClick(page, "nav-account");
    await expect(page).toHaveURL(/\/account$/);

    // Check text and button states
    await expect(page.getByText("No API Key")).toBeVisible();
    await expect(page.getByTestId("create-apikey-button")).not.toBeDisabled();
    await expect(page.getByTestId("delete-apikey-button")).toBeDisabled();

    // Create API Key
    await pageClick(page, "create-apikey-button");
    await expect(
      page.getByText("testapikey123456789abcdefghijklmn")
    ).toBeVisible();

    // Copy API key
    await pageClick(page, "copy-apikey-button");
    const handle = await page.evaluateHandle(() =>
      navigator.clipboard.readText()
    );
    const clipboardContent = await handle.jsonValue();
    expect(clipboardContent).toEqual("testapikey123456789abcdefghijklmn");

    // Close the dialog, Check text and button states
    await pageClick(page, "dialog-close");
    await expect(page.getByText("************")).toBeVisible();
    await expect(page.getByTestId("create-apikey-button")).toBeDisabled();
    await expect(page.getByTestId("delete-apikey-button")).not.toBeDisabled();

    // Delete API Key
    await pageClick(page, "delete-apikey-button");

    // Confirmation dialog should be visible
    await expect(page.getByText("Delete API Key")).toBeVisible();
    await expect(
      page.getByText("Are you sure to delete your API Key?")
    ).toBeVisible();

    // Confirm deletion
    await page.getByText("Proceed").click();

    // Success alert should appear
    await expect(page.getByText("API Key deleted successfully")).toBeVisible();

    // Check text and button states
    await expect(page.getByText("No API Key")).toBeVisible();
    await expect(page.getByTestId("create-apikey-button")).not.toBeDisabled();
    await expect(page.getByTestId("delete-apikey-button")).toBeDisabled();
  });
});
