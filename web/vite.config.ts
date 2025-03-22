import vuetify from "vite-plugin-vuetify";
import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";

// ref: https://vite.dev/config/
export default defineConfig({
  plugins: [vue(), vuetify({ autoImport: true })],
  base: "/",
  server: {
    port: 5179,
    // ngrokでローカルを公開してスマホ実機での検証をしたいときに設定する
    // allowedHosts: [""],
  },
});
