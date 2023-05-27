<template>
  <v-card class="elevation-1 rounded-xl pa-3">
    <v-card-title>{{ title }}</v-card-title>
    <v-table height="150px" density="compact" fixed-header hover>
      <thead>
        <tr>
          <th class="table-head" v-for="value in headers" :key="value">
            {{ value }}
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in items" :key="item">
          <td v-for="(value, key) in item" :key="value">
            {{ value }}
          </td>
        </tr>
      </tbody>
    </v-table>
  </v-card>
</template>

<script>
import { reactive } from "vue";
export default {
  name: "TableCard",
  props: ["title", "headers", "items"],
  setup() {},
};
</script>

<style scoped>
.table-head {
  font-size: 12px;
}
</style>
