<template>
  <div id="app">
    <Sidebar v-if="showSidebar" />
    <div class="content" :class="{ 'with-sidebar': showSidebar }">
      <router-view />
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import Sidebar from './components/Sidebar.vue'

const route = useRoute()
const showSidebar = computed(() => route.path !== '/login')
</script>

<style scoped>
#app {
  display: flex;
}

.content {
  flex: 1;
  padding: 20px;
}

.with-sidebar {
  margin-left: 200px;
  /* サイドバーの幅に合わせて左マージンを追加 */
}
</style>