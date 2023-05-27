<template>
  <v-card class="bg-16dp rounded-lg smallcard">
    <v-container
      class="smallcard-container d-flex flex-column justify-space-around align-center"
    >
      <v-row class="smallcard-amount" align="center">
        <v-col> {{ number }}</v-col>
      </v-row>
      <v-row class="ma-0" align="center">
        <v-col>{{ topic }}</v-col>
      </v-row>
    </v-container>
  </v-card>
</template>

<script>
export default {
  props: ["number", "topic"],
};
</script>

<style lang="scss" scope>
.v-card.smallcard {
  width: 100%;
  position: relative;
  top: 0;
  left: 0;
  overflow: hidden;
  &::after {
    content: "";
    display: block;
    padding-bottom: 100%;
  }
  .v-container.smallcard-container {
    position: absolute;
    width: 100%;
    height: 100%;
    .smallcard-amount {
      font-size: 60px;
    }
    .v-col {
      padding: 0;
    }
  }
}
</style>
