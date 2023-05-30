<template>
  <v-row class="title"><h2>SUPPORT CASE</h2></v-row>
  <v-row class="title"><h5>List of support case.</h5></v-row>
  <v-container class="pa-1" />
  <v-divider />
  <v-container class="pa-1" />
  <v-row class="support-case-body" justify="space-between">
    <V-col cols="12">
      <!-- Space -->
      <v-container class="ma-1" />

      <!-- Search Field -->
      <v-container class="pa-1">
        <v-row align="center">
          <v-spacer />
          <v-text-field></v-text-field>
          <v-spacer />
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
