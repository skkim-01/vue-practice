<template>
  <v-card class="bg-16dp rounded-lg smallcard">
    <v-container
      class="smallcard-container d-flex flex-column justify-center align-center"
    >
      <v-row class="smallcard-amount" align="center">
        <v-col> {{ number }}</v-col>
      </v-row>
      <v-row class="ma-0 topic" align="center">
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
  // 크기 수정하고 싶으실때 아래 width값만 (원하시는 값으로)수정하시면 됩니다.
  width: 150px;
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
    text-align: center;
    .smallcard-amount {
      // 숫자 size 설정
      font-size: 70px;
    }
    .topic {
      // 숫자 아래 topic font size를 설정
      font-size: 14px;
    }
    .v-col {
      padding: 0;
    }
  }
}
</style>
