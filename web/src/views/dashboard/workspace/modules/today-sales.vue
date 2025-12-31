<template>
  <div class="art-card h-82 p-5 mb-5 overflow-hidden max-lg:h-auto max-sm:mb-4">
    <div class="art-card-header pr-0">
      <div class="title">
        <h4>待办事项</h4>
      </div>
      <!-- <div class="flex-cc h-7.5 min-w-17 border border-g-300 rounded-lg text-g-500 c-p">
        <ArtSvgIcon icon="ri:arrow-up-line" class="text-base mr-1.5" />
        <span class="text-xs">导出</span>
      </div> -->
    </div>

    <div class="mt-2">
      <ElRow :gutter="20">
        <ElCol :span="6" :xs="24">
          <div
            class="flex px-5 flex-col justify-center h-55 border border-g-300/85 rounded-xl max-lg:mb-4 max-sm:flex-row max-sm:justify-between max-sm:items-center max-sm:h-40"
          >
            <div class="size-12 rounded-lg flex-cc bg-theme/10">
              <ArtSvgIcon icon="solar:filters-bold" class="text-xl text-theme" />
            </div>

            <div class="max-sm:ml-4 mt-3.5 max-sm:mt-0 max-sm:text-end">
              <ArtCountTo class="text-2xl font-medium" :target="pendingServiceOrder.totalCount" :duration="1500" />
              <p class="mt-2 text-base text-g-600 max-sm:mt-1">订单待派单</p>
              <small class="text-g-500 mt-1 max-sm:mt-0.5">
                今日
                <span class="font-medium text-success"> {{ pendingServiceOrder.todayCount }} </span>
              </small>
            </div>
          </div>
        </ElCol>
        <ElCol :span="6" :xs="24">
          <div
            class="flex px-5 flex-col justify-center h-55 border border-g-300/85 rounded-xl max-lg:mb-4 max-sm:flex-row max-sm:justify-between max-sm:items-center max-sm:h-40"
          >
            <div class="size-12 rounded-lg flex-cc bg-theme/10">
              <ArtSvgIcon icon="solar:square-transfer-vertical-bold" class="text-xl text-theme" />
            </div>

            <div class="max-sm:ml-4 mt-3.5 max-sm:mt-0 max-sm:text-end">
              <ArtCountTo class="text-2xl font-medium" :target="applySettlement.totalCount" :duration="1500" />
              <p class="mt-2 text-base text-g-600 max-sm:mt-1">报单待结算</p>
              <small class="text-g-500 mt-1 max-sm:mt-0.5">
                今日
                <span class="font-medium text-success"> {{ applySettlement.todayCount }} </span>
              </small>
            </div>
          </div>
        </ElCol>
        <ElCol :span="6" :xs="24">
          <div
            class="flex px-5 flex-col justify-center h-55 border border-g-300/85 rounded-xl max-lg:mb-4 max-sm:flex-row max-sm:justify-between max-sm:items-center max-sm:h-40"
          >
            <div class="size-12 rounded-lg flex-cc bg-theme/10">
              <ArtSvgIcon icon="solar:wad-of-money-bold" class="text-xl text-theme" />
            </div>

            <div class="max-sm:ml-4 mt-3.5 max-sm:mt-0 max-sm:text-end">
              <ArtCountTo class="text-2xl font-medium" :target="applyWithdraw.totalCount" :duration="1500" />
              <p class="mt-2 text-base text-g-600 max-sm:mt-1">待审核提现</p>
              <small class="text-g-500 mt-1 max-sm:mt-0.5">
                今日
                <span class="font-medium text-success"> {{ applyWithdraw.todayCount }} </span>
              </small>
            </div>
          </div>
        </ElCol>
        <ElCol :span="6" :xs="24">
          <div
            class="flex px-5 flex-col justify-center h-55 border border-g-300/85 rounded-xl max-lg:mb-4 max-sm:flex-row max-sm:justify-between max-sm:items-center max-sm:h-40"
          >
            <div class="size-12 rounded-lg flex-cc bg-theme/10">
              <ArtSvgIcon icon="solar:chat-round-dots-bold" class="text-xl text-theme" />
            </div>

            <div class="max-sm:ml-4 mt-3.5 max-sm:mt-0 max-sm:text-end">
              <ArtCountTo class="text-2xl font-medium" :target="applyComment.totalCount" :duration="1500" />
              <p class="mt-2 text-base text-g-600 max-sm:mt-1">评论待审核</p>
              <small class="text-g-500 mt-1 max-sm:mt-0.5">
                今日
                <span class="font-medium text-success"> {{ applyComment.todayCount }} </span>
              </small>
            </div>
          </div>
        </ElCol>
      </ElRow>
    </div>
  </div>
</template>

<script setup lang="ts">
interface Props {
  pendingServiceOrder: {
    totalCount:number;
    todayCount:number;
  },
  applySettlement: {
    totalCount:number;
    todayCount:number;
  },
  applyWithdraw: {
    totalCount:number;
    todayCount:number;
  },
  applyComment: {
    totalCount:number;
    todayCount:number;
  },
}
defineProps<Props>()
</script>