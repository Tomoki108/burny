import { test, expect } from "@playwright/test";

test.describe("Yahoo Finance Japan", () => {
  test("Extract 3 main financial news headlines", async ({ page }) => {
    // Yahoo Finance Japanのトップページにアクセス
    await page.goto("https://finance.yahoo.co.jp/");

    // ページタイトルを検証
    await expect(page).toHaveTitle(/Yahoo!ファイナンス/);

    console.log("Yahoo Finance Japan page loaded successfully");

    // ヘッドラインセクションが表示されるのを待つ
    const headlineSection = page.locator("text=ヘッドライン").first();
    await expect(headlineSection).toBeVisible();

    // ニュース記事のリストを取得
    const newsArticles = page.locator("article a h2").filter({ hasText: /.+/ });

    // 少なくとも3つのニュース記事があることを確認
    const count = await newsArticles.count();
    expect(count).toBeGreaterThanOrEqual(3);
    console.log(`Found ${count} news articles`);

    // 最初の3つのニュース見出しを抽出
    const headlines = [];
    for (let i = 0; i < Math.min(3, count); i++) {
      const headline = await newsArticles.nth(i).textContent();
      const timeElement = page.locator("article a").nth(i).locator("time");
      const time = await timeElement.textContent();

      const sourceElement = page
        .locator("article a")
        .nth(i)
        .locator("li")
        .nth(1);
      const source = await sourceElement.textContent();

      headlines.push({
        headline: headline?.trim(),
        time: time?.trim(),
        source: source?.trim(),
      });

      console.log(`Headline ${i + 1}: ${headline?.trim()}`);
    }

    // 少なくとも3つの見出しが取得できたことを確認
    expect(headlines.length).toBeGreaterThanOrEqual(3);

    // 各見出しが空でないことを確認
    headlines.forEach((item, index) => {
      expect(item.headline).toBeTruthy();
      console.log(
        `Headline ${index + 1}: ${item.headline} (${item.time} - ${
          item.source
        })`
      );
    });

    // スクリーンショットを撮る（デバッグ用）
    await page.screenshot({ path: "yahoo-finance-japan.png" });

    console.log("Test completed successfully");
  });
});
