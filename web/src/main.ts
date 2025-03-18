import { createApp } from "vue";
import App from "./App.vue";
import router from "./router.ts";
import vuetify from "./plugins/vuetify";
import { library } from "@fortawesome/fontawesome-svg-core";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import {
  faProjectDiagram,
  faSignOutAlt,
  faUser,
} from "@fortawesome/free-solid-svg-icons";
import { createPinia } from "pinia";
import "./assets/styles/global.css";
import "@mdi/font/css/materialdesignicons.css";

library.add(faProjectDiagram, faSignOutAlt, faUser);

enableMocking().then(() => {
  createApp(App)
    .component("font-awesome-icon", FontAwesomeIcon)
    .use(createPinia())
    .use(router)
    .use(vuetify)
    .mount("#app");
});

async function enableMocking() {
  if (import.meta.env.VITE_MOCK_API == "true") {
    const { worker } = await import("../tests/mock_server.ts");
    return worker.start();
  }
}
