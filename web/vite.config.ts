import vuetify from "vite-plugin-vuetify";
import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";

// ref: https://vite.dev/config/
export default defineConfig({
  plugins: [vue(), vuetify({ autoImport: true })],
  // GCS静的ホスティングのためベースパスを空文字列に設定
  base: "",
  server: {
    port: 5179,
    allowedHosts: ["3459-2001-240-2989-0-b4fc-1491-e7ec-4ca9.ngrok-free.app"],
  },
});
