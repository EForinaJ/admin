<template>
  <div  class="art-card h-82 p-5 mb-5 max-sm:mb-4">
    <div class="art-card-header">
      <div class="title">
        <h4>订单统计</h4>
        <p>近一周的订单统计结果</p>
      </div>
    </div>
    <ArtBarChart
      class="box-border p-2"
      barWidth="50%"
       height="calc(100% - 56px)"
      :showAxisLine="false"
      :data="chartData"
      :xAxisData="xAxisLabels"
    />
  </div>
</template>

<script setup lang="ts">
  
interface Props {
  xAxisLabels: string[],
  chartData: number[],
}
defineProps<Props>()
</script>