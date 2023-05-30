<template>
  <!-- 가운데 정렬이 싫은 경우 justify를 삭제하시거나 원하는 위치로 수정하시면 됩니다. -->
  <v-row align="center" justify="center">
    <v-col cols="1">
      <div class="text-caption text-end">{{ label }}</div>
    </v-col>

    <slot />
  </v-row>
</template>

<script>
export default {
  name: "FormItem",
  props: ["label"],
};
</script>

<style></style>
