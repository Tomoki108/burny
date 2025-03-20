# Burny Web

SPA for Burny.

## Teck Stack

| Category             | Tool                                                                 |
| -------------------- | -------------------------------------------------------------------- |
| **FW**               | [Vue.js (Compostion API)](https://vuejs.org/guide/introduction.html) |
| **Build**            | [Vite](https://ja.vite.dev/)                                         |
| **UI Component**     | [Vuetify](https://vuetifyjs.com/ja/)                                 |
| **Routing**          | [Vue Router](https://router.vuejs.org/)                              |
| **State Management** | [Pinia](https://pinia.vuejs.org/)                                    |
| **Test**             | [Playwright](https://playwright.dev/)                                |

## How to run

```shell
npm run dev
```

### Prerequisites

```shell
cp .env.sample .env
npm install
```

### Other Commands

```shell
# run scenario tests
#  - web server must be running on port 5179
#  - you can specify test file path as argument
npm run test
npm run testh # headed mode
npm run testhd # headed mode & debug mode
```
