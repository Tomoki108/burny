import { test, expect } from "@playwright/test";
import { WEB_LOCAL_HOST, pageClick, login, sleep } from "./test_helper";
import { mockAllApis, mockApiKeyStatusApi, mockApiKeyApi } from "./test_mock";

test.describe("Account page", () => {
  test("User can create and delete API key", async ({
    page,
    playwright,
    context,
  }) => {
    // セットアップ：APIレスポンスをモックする（初期状態ではAPIキーが存在しない）
    await mockAllApis(page);
    await mockApiKeyStatusApi(page, false);

    // ログインする
    await login(page);

    // グローバルメニューからAccountページに移動
    await page.getByTestId("global-menu-account").click();

    // スナップショットを取得して、アクセシビリティツリーを確認
    const snapshot1 = await page.accessibility.snapshot();
    console.log("Page navigated to Account section");

    // 最初はAPI Keyが登録されていないことを確認
    const noApiKeyElement = page.getByTestId("no-api-key-message");
    await expect(noApiKeyElement).toBeVisible();
    expect(await noApiKeyElement.textContent()).toContain("No API Key");
    await expect(page.getByTestId("api-key-value")).not.toBeVisible();

    // API Keyを作成する
    const createButton = page.getByTestId("create-api-key-button");
    await createButton.click();

    // API作成後のモックレスポンスを設定
    await mockApiKeyStatusApi(page, true);

    // 作成したAPIキーが表示されていることを確認
    const apiKeyElement = page.getByTestId("api-key-value");
    await expect(apiKeyElement).toBeVisible();

    // 画面に表示されるAPIキーがモックで指定した値（testapikey123456789abcdefghijklmn）と一致するか確認
    const displayedApiKey = await apiKeyElement.textContent();
    expect(displayedApiKey).toBe("testapikey123456789abcdefghijklmn");

    // スナップショットを取得して、アクセシビリティツリーを確認（APIキーが表示されている状態）
    const snapshot2 = await page.accessibility.snapshot();
    console.log("API Key has been created and displayed");

    // API Keyを削除する
    const deleteButton = page.getByTestId("delete-api-key-button");
    await deleteButton.click();

    // 削除確認ダイアログで「はい」を選択
    await page.getByTestId("dialog-proceed").click();

    // 削除後のモックレスポンスを設定
    await mockApiKeyStatusApi(page, false);

    // 削除後に「No API Key」メッセージが表示されることを確認
    await expect(noApiKeyElement).toBeVisible();
    expect(await noApiKeyElement.textContent()).toContain("No API Key");
    await expect(apiKeyElement).not.toBeVisible();

    // スナップショットを取得して、アクセシビリティツリーを確認（APIキーが削除された状態）
    const snapshot3 = await page.accessibility.snapshot();
    console.log("API Key has been deleted");
  });
});
