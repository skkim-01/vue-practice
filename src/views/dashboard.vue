<template>
  <v-row class="title"><h2>DASHBOARD</h2></v-row>
  <v-row class="title"><h5>Management console dashboard page.</h5></v-row>
  <v-container class="pa-1" />
  <v-divider />
  <v-container class="ma-1" />
  <v-row class="dashbord_body" justify="space-between">
    <v-col cols="12" class="">
      <v-container class="pa-0">
        <v-row>
          <v-col>
            <v-card class="bg-1dp rounded-xl issue-trend">
              <v-card-title>Trends</v-card-title>
              <v-container>
                <v-row justify="center">
                  <v-col cols="6" class="d-flex justify-end">
                    <SmallCard number="24" topic="Case Open" />
                  </v-col>
                  <v-col cols="6" class="d-flex justify-start">
                    <SmallCard number="13" topic="Case Resolve" />
                  </v-col>
                </v-row>
              </v-container>
            </v-card>
          </v-col>
        </v-row>

        <v-row>
          <v-col>
            <ChartCard title="Trend Chart" />
          </v-col>
        </v-row>

        <v-row>
          <v-col>
            <TableCard
              title="My Cases"
              :headers="state.headers"
              :items="state.desserts"
              :tableHeight="6"
            />
          </v-col>
        </v-row>

        <v-row>
          <v-col>
            <TableCard
              title="Recent Cases"
              :headers="state.recentCaseHeader"
              :items="state.recentCaseItems"
              :tableHeight="6"
            />
          </v-col>
        </v-row>
      </v-container>
    </v-col>
  </v-row>
</template>

<script>
import { reactive } from "vue";
import SmallCard from "../components/SmallCard";

export default {
  // data,
  // methods,
  name: "DashBoard",
  components: { SmallCard },
  setup() {
    const state = reactive({
      chart_options: {
        title: {
          text: "World Population",
        },
        tooltip: {
          trigger: "axis",
          axisPointer: {
            type: "shadow",
          },
        },
        legend: {},
        grid: {
          left: "3%",
          right: "4%",
          bottom: "3%",
          containLabel: true,
        },
        yAxis: {
          type: "value",
          boundaryGap: [0, 0.01],
        },
        xAxis: {
          type: "category",
          data: ["Brazil", "Indonesia", "USA", "India", "China", "World"],
        },
        series: [
          {
            name: "2011",
            type: "bar",
            data: [18203, 23489, 29034, 104970, 131744, 630230],
          },
          {
            name: "2012",
            type: "bar",
            data: [19325, 23438, 31000, 121594, 134141, 681807],
          },
        ],
      },
      headers: ["name", "calories", "fat", "carbs", "protein", "iron"],
      desserts: [
        {
          name: "Frozen Yogurt",
          calories: 159,
          fat: 6.0,
          carbs: 24,
          protein: 4.0,
          iron: "1",
        },
        {
          name: "Ice cream sandwich",
          calories: 237,
          fat: 9.0,
          carbs: 37,
          protein: 4.3,
          iron: "1",
        },
        {
          name: "Eclair",
          calories: 262,
          fat: 16.0,
          carbs: 23,
          protein: 6.0,
          iron: "7",
        },
        {
          name: "Cupcake",
          calories: 305,
          fat: 3.7,
          carbs: 67,
          protein: 4.3,
          iron: "8",
        },
        {
          name: "Gingerbread",
          calories: 356,
          fat: 16.0,
          carbs: 49,
          protein: 3.9,
          iron: "16",
        },
        {
          name: "Jelly bean",
          calories: 375,
          fat: 0.0,
          carbs: 94,
          protein: 0.0,
          iron: "0",
        },
        {
          name: "Lollipop",
          calories: 392,
          fat: 0.2,
          carbs: 98,
          protein: 0,
          iron: "2",
        },
      ],
      recentCaseHeader: [
        "Case ID",
        "Account ID",
        "Payer ID",
        "고객명",
        "Service",
        "Category",
        "Subject",
        "Severity",
        "Status",
        "Updated",
      ],
      recentCaseItems: [],
      recentCmntHeader: [
        "Case ID",
        "Account ID",
        "Service",
        "Category",
        "Subject",
        "comment",
        "Updated",
      ],
      recentCmntItems: [],
    });
    return { state };
  },
};
</script>

<style lang="scss" scoped>
.title {
  margin: 0;
}
.issue-trend {
  // position: static;
  overflow: hidden;
}
</style>
