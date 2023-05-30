<template>
  <v-row class="title"><h2>SUPPORT PLAN</h2></v-row>
  <v-row class="support-plan-body" justify="space-between">
    <V-col cols="12">
      <!-- Space -->
      <v-container class="ma-1" />

      <!-- Search Field -->
      <v-container class="pa-1">
        <v-row>
          <v-text>검색</v-text>
          <v-text-field>검색창</v-text-field>
        </v-row>
      </v-container>
    
      <!-- divider -->
      <v-divider />

      <!-- Space -->
      <v-container class="ma-2" />
    
      <!-- Table -->
      <TableCard />

    </V-col>
  </v-row>
</template>
  
  
<script>
export default {
    name: 'supportPlan',
};
</script>


<style lang="scss" scoped>
.title {
  margin: 0;
}
</style>
