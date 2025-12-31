<!-- 工作台页面 -->
<template>
  <div v-loading="loading">
   
    <CardList 
    :profit-amount="detail?.profitAmount!"
    :order-count="detail?.orderCount!"
    :sales-amount="detail?.salesAmount!"
    :user-count="detail?.userCount!"
    />
   
    <ElRow :gutter="20">
      <ElCol :sm="24" :md="12" :lg="12">
        <TodaySales
          :apply-comment="detail?.applyComment!"
          :apply-settlement="detail?.applySettlement!"
          :pending-service-order="detail?.pendingServiceOrder!"
          :apply-withdraw="detail?.applyWithdraw!"
        />
      </ElCol>
      <ElCol :sm="24" :md="12" :lg="12">
        <OrderStats
          :chart-data="detail?.orderStatistic.count!"
          :x-axis-labels="detail?.orderStatistic.weekdays!"
        />
      </ElCol>
    </ElRow>

    <SalesOverview 
      :data="detail?.userStatistic.count!"
      :x-axis-data="detail?.userStatistic.days!"
    />

    <!-- <ElRow :gutter="20">
      <ElCol :sm="24" :md="24" :lg="8">
        <Dynamic />
      </ElCol>
      <ElCol :sm="24" :md="12" :lg="8">
        <Dynamic />
      </ElCol>
      <ElCol :sm="24" :md="12" :lg="8">
        <TodoList />
      </ElCol>
    </ElRow> -->

    <AboutProject />
  </div>
</template>

<script setup lang="ts">
import CardList from './modules/card-list.vue'
import OrderStats from './modules/order-stats.vue'
import SalesOverview from './modules/sales-overview.vue'
import AboutProject from './modules/about-project.vue'
import { fetchGetDashboardDetail } from '@/api/dashboard'
import TodaySales from './modules/today-sales.vue'

defineOptions({ name: 'Console' })
const detail = ref<Dashboard.Response.Detail>({
  userCount:{
      totalCount:0,
      todayCount:0,
  },
  orderCount:{
      totalCount:0,
      todayCount:0,
  },
  salesAmount:{
      amountTotal:0,
      amountToday:0,
  },
  profitAmount:{
      amountTotal:0,
      amountToday:0,
  },
  orderStatistic:{
      weekdays:[],
      count:[],
  },
  userStatistic:{
      days:[],
      count:[],
  },
  pendingServiceOrder:{
      totalCount:0,
      todayCount:0,
  },
  applySettlement:{
      totalCount:0,
      todayCount:0,
  },
  applyWithdraw:{
      totalCount:0,
      todayCount:0,
  },
  applyComment:{
      totalCount:0,
      todayCount:0,
  }
})
const loading = ref<boolean>(false)
const getData = async () => {
    loading.value = true
    const res = await fetchGetDashboardDetail()
    detail.value = res
    detail.value.userStatistic.count = res.userStatistic.count.reverse()
    detail.value.userStatistic.days = res.userStatistic.days.reverse()
    detail.value.orderStatistic.count = res.orderStatistic.count.reverse()
    detail.value.orderStatistic.weekdays = res.orderStatistic.weekdays.reverse()
    loading.value = false
}


onMounted(()=>{
    getData()
})
</script>
