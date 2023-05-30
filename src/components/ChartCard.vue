<template>
  <v-card class="elevation-1 rounded-xl pa-3">
    <vue-echarts
      :option="chart_options"
      style="height: 400px"
      ref="chart"
    />  
  </v-card>
</template>

<script setup>
import { VueEcharts } from "vue3-echarts";

// TODO: Using props
const props = defineProps({
    chartOptions: {
        type: Object,
        required: true,
    }
})

const chart_options= {
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
      };
</script>