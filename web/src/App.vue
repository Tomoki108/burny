<template>
  <v-app>
    <router-view />
  </v-app>
</template>

<style>
#app {
  display: flex;
  flex-direction: column;
  height: 100vh;
  /* min-heightをheightに変更 */  overflow-y: auto; /* コンテンツが高さを超える場合のみスクロール可能に */
}

main {
  flex: 1;
}
</style>
