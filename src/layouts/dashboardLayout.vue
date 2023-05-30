<template>
  <v-app theme="dark">
    <toolbar app primary fixed />
    <navigation app dark fixed />
    <v-main>
      <v-container class="dashboard pa-4">
        <slot />
      </v-container>
      <page-footer />
    </v-main>
  </v-app>
</template>

<style lang="scss">
.dashboard {
  width: 100%;
  // height: calc(100% - 64px);
  height: calc(100% - 31px);
  margin: 0;
  max-width: 100%;
  .v-row {
    margin: 0;
    width: 100%;
  }
}
</style>
