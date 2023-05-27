<template>
  <v-app theme="dark">
    <toolbar app primary fixed />
    <navigation app dark fixed />
    <v-main>
      <v-container class="dashboard">
        <slot />
      </v-container>
      <page-footer />
    </v-main>
  </v-app>
</template>

<style lang="scss" scoped>
.dashboard {
  width: 100%;
  // height: calc(100% - 64px);
  height: calc(100% - 31px);
  margin: 0;
  max-width: 100%;
}
</style>
