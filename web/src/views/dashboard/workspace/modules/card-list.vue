<template>
  <ElRow :gutter="20" class="flex">
    <ElCol :sm="12" :md="6" :lg="6">
      <div class="art-card relative flex flex-col justify-center h-35 px-5 mb-5 max-sm:mb-4">
        <span class="text-g-700 text-sm">新用户</span>
        <ArtCountTo class="text-[26px] font-medium mt-2" :target="userCount.todayCount" :duration="1300" />
        <div class="flex-c mt-1">
          <span class="text-xs text-g-600">用户总量</span>
          <span
            class="ml-1 text-xs font-semibold text-success"
          >
            {{ userCount.totalCount }}
          </span>
        </div>
        <div
          class="absolute top-0 bottom-0 right-5 m-auto size-12.5 rounded-xl flex-cc bg-theme/10"
        >
          <ArtSvgIcon icon="solar:user-bold" class="text-xl text-theme" />
        </div>
      </div>
    </ElCol>
    <ElCol :sm="12" :md="6" :lg="6">
      <div class="art-card relative flex flex-col justify-center h-35 px-5 mb-5 max-sm:mb-4">
        <span class="text-g-700 text-sm">今日下单</span>
        <ArtCountTo class="text-[26px] font-medium mt-2" :target="orderCount.todayCount" :duration="1300" />
        <div class="flex-c mt-1">
          <span class="text-xs text-g-600">总订单数</span>
          <span
            class="ml-1 text-xs font-semibold text-success"
          >
            {{ orderCount.totalCount }}
          </span>
        </div>
        <div
          class="absolute top-0 bottom-0 right-5 m-auto size-12.5 rounded-xl flex-cc bg-theme/10"
        >
          <ArtSvgIcon icon="solar:shop-bold" class="text-xl text-theme" />
        </div>
      </div>
    </ElCol>
    <ElCol :sm="12" :md="6" :lg="6">
      <div class="art-card relative flex flex-col justify-center h-35 px-5 mb-5 max-sm:mb-4">
        <span class="text-g-700 text-sm">销售额</span>
        <ArtCountTo class="text-[26px] font-medium mt-2" 
          :suffix="site.symbol"
          :target="salesAmount.amountToday" :duration="1300" />
        <div class="flex-c mt-1">
          <span class="text-xs text-g-600">总销售额度</span>
          <span
            class="ml-1 text-xs font-semibold text-success"
          >
            {{ salesAmount.amountTotal }} {{ site.symbol }}
          </span>
        </div>
        <div
          class="absolute top-0 bottom-0 right-5 m-auto size-12.5 rounded-xl flex-cc bg-theme/10"
        >
          <ArtSvgIcon icon="solar:wallet-money-bold" class="text-xl text-theme" />
        </div>
      </div>
    </ElCol>
    <ElCol :sm="12" :md="6" :lg="6">
      <div class="art-card relative flex flex-col justify-center h-35 px-5 mb-5 max-sm:mb-4">
        <span class="text-g-700 text-sm">今日收益</span>
        <ArtCountTo class="text-[26px] font-medium mt-2" 
          :suffix="site.symbol"
          :target="profitAmount.amountToday" :duration="1300" />
        <div class="flex-c mt-1">
          <span class="text-xs text-g-600">累计收益金额</span>
          <span
            class="ml-1 text-xs font-semibold text-success"
          >
            {{ profitAmount.amountTotal }} {{ site.symbol }}
          </span>
        </div>
        <div
          class="absolute top-0 bottom-0 right-5 m-auto size-12.5 rounded-xl flex-cc bg-theme/10"
        >
          <ArtSvgIcon icon="solar:chat-round-money-bold" class="text-xl text-theme" />
        </div>
      </div>
    </ElCol>
  </ElRow>
</template>

<script setup lang="ts">
import { useSiteStore } from '@/store/modules/site';

interface Props {
  userCount: {
    totalCount:number;
    todayCount:number;
  },
  orderCount: {
    totalCount:number;
    todayCount:number;
  },
  salesAmount: {
    amountTotal:number;
    amountToday:number;
  },
  profitAmount: {
    amountTotal:number;
    amountToday:number;
  },
}
defineProps<Props>()

const { getInfo:site } = useSiteStore()
</script>
