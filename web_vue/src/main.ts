import { createApp } from "vue";
import "./style.css";
import App from "./App.vue";
import router from "./router";
import vuetify from "./plugins/vuetify";
import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { faHome, faProjectDiagram, faCog, faSignOutAlt } from '@fortawesome/free-solid-svg-icons'

library.add(faHome, faProjectDiagram, faCog, faSignOutAlt)

createApp(App)
.component('font-awesome-icon', FontAwesomeIcon)
.use(router)
.use(vuetify)
.mount("#app");
