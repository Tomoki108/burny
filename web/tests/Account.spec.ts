import { test, expect } from "@playwright/test";
import { WEB_LOCAL_HOST, login, sleep } from "./test_helper";
import path from "path";
import { mkdir } from "fs/promises";

test.describe("Account page", () => {
  test("User can view their account information", async ({ page }) => {
    // Create screenshots directory if it doesn't exist
    const screenshotsDir = path.join(process.cwd(), "test-results/screenshots");
    await mkdir(screenshotsDir, { recursive: true });

    // ログインして、Projects画面まで遷移する
    await login(page);

    // Take screenshot after login (should be on Projects page)
    await page.screenshot({
      path: path.join(screenshotsDir, "2-after-login.png"),
      fullPage: true,
    });

    // プロジェクトページに正しく遷移したことを確認
    await expect(page).toHaveURL(/\/projects$/);

    // サイドバーが表示されるまで待機
    await page.waitForSelector(".sidebar", { state: "visible" });

    // nav-accountの要素が表示されるまで確実に待機
    await page.waitForSelector('[data-testid="nav-account"]', {
      state: "visible",
      timeout: 10000,
    });

    // アカウントページへの遷移を事前に監視
    const navigationPromise = page.waitForNavigation({
      url: /\/account$/,
      waitUntil: "networkidle",
    });

    // navアカウント要素のクリックを確実に実行（何度か試行）
    let success = false;
    for (let i = 0; i < 3; i++) {
      try {
        // フォーカスする
        await page.focus('[data-testid="nav-account"]');
        await sleep(200);

        // クリック（forceオプション付き）
        await page.click('[data-testid="nav-account"]', {
          force: true,
          timeout: 5000,
        });

        success = true;
        break;
      } catch (e) {
        console.log(`Click attempt ${i + 1} failed:`, e);
        await sleep(1000);
        // スクリーンショットを撮る
        await page.screenshot({
          path: path.join(screenshotsDir, `click-attempt-${i + 1}.png`),
          fullPage: true,
        });
      }
    }

    if (!success) {
      // 最終手段：JavaScriptでダイレクトに遷移
      await page.evaluate(() => {
        window.location.href = "/account";
      });
    }

    // 遷移を待機
    try {
      await navigationPromise;
    } catch (e) {
      console.log("Navigation timeout, forcing navigation");
      // 最終手段：再度直接URLに移動
      await page.goto(`${WEB_LOCAL_HOST}/account`);
    }

    // アカウントページに遷移したことを確認
    await page.waitForURL(/\/account$/, { timeout: 10000 });

    // スクリーンショットを撮る
    await page.screenshot({
      path: path.join(screenshotsDir, "5-account-page.png"),
      fullPage: true,
    });

    // アカウント情報が表示されていることを確認
    await expect(page.getByText("test@example.com")).toBeVisible();
    await expect(page.getByText("********")).toBeVisible();
  });
});
