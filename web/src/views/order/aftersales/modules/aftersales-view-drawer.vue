<template>
    <ElDrawer
        v-model="visible"
        title="订单详情"
        size="50%"
        @close="handleClose"
        >
        <ElCard v-loading="loading" shadow="never">
            <template #header>
                <div class="flex items-center gap-3">
                    <div class="font-bold">
                        订单编号: {{ detail?.order.code }}
                    </div>
                    <el-divider direction="vertical" />
                    <div class="font-bold">
                        下单时间: {{ detail?.order.createTime }}
                    </div>
                </div>
            </template>
            <template #footer>
                <div class="flex items-center gap-3">
                    <ElAvatar :src="detail?.order.user.avatar" shape="square" :size="28" />
                    <div class="flex-col flex gap-2">
                        <h2 class="font-bold">
                            {{detail?.order.user.name}}
                        </h2>
                    </div>
                </div>
            </template>
            <ElDescriptions>
                <ElDescriptionsItem :span="1"  label="订单状态">
                    <ElTag :type="getOrderStatusConfig(detail?.order.status!).type">{{ getOrderStatusConfig(detail?.order.status!).text }}</ElTag>
                </ElDescriptionsItem>
                <ElDescriptionsItem :span="1"  label="支付状态">
                    <ElTag :type="getOrderPayStatus(detail?.order.payStatus!).type">{{ getOrderPayStatus(detail?.order.payStatus!).text }}</ElTag>
                </ElDescriptionsItem>
                <ElDescriptionsItem :span="1"  label="支付方式">
                    <ElTag :type="getOrderPayMode(detail?.order.payMode!).type">{{ getOrderPayMode(detail?.order.payMode!).text }}</ElTag>
                </ElDescriptionsItem>
                <ElDescriptionsItem :span="1"  label="支付时间">
                    {{ detail?.order.payTime  ? detail?.order.payTime :  '未支付' }}
                </ElDescriptionsItem>
                <ElDescriptionsItem :span="1"  label="开始时间">
                    {{ detail?.order.startTime  ? detail?.order.startTime : "未开始" }}
                </ElDescriptionsItem>
                <ElDescriptionsItem :span="1"  label="结束时间">
                    {{ detail?.order.finishTime  ? detail?.order.finishTime :  "未结束" }}
                </ElDescriptionsItem>
                <ElDescriptionsItem :span="3"  label="用户要求">
                    {{ detail?.order.requirements == "" ? '无备注' : detail?.order.requirements }}
                </ElDescriptionsItem>
            </ElDescriptions>
        </ElCard>

        <ElCard v-loading="loading" shadow="never" class="mt-[20px]">
            <template #header>
                <div class="font-bold">
                    商品详情
                </div>
            </template>
            <ElTable :data="productData" style="width: 100%">
                <ElTableColumn prop="productInfo" label="商品信息" width="280">
                    <template #default="scope">
                        <div class="flex gap-3">
                            <ElImage class="size-12 rounded-md" :src="scope.row.pic"/>
                            <h2 class="text-base line-clamp-1">{{ scope.row.name }}</h2>
                        </div>
                    </template>
                </ElTableColumn>
                <ElTableColumn prop="game" label="所属游戏" width="180" >
                    <template #default="scope">
                        <ElTag type="primary">{{ scope.row.game }}</ElTag>
                    </template>
                </ElTableColumn>
                <ElTableColumn prop="category" label="所属分类" width="180" >
                    <template #default="scope">
                        <ElTag type="primary">{{ scope.row.category }}</ElTag>
                    </template>
                </ElTableColumn>
                <ElTableColumn prop="unitPrice" label="单价">
                    <template #default="scope">
                        {{ `${scope.row.unitPrice}${site.symbol}` }}
                    </template>
                </ElTableColumn>
                <ElTableColumn prop="quantity" label="数量">
                    <template #default="scope">
                        {{ `${scope.row.quantity}${scope.row.unit}` }}
                    </template>
                </ElTableColumn>
            </ElTable>
            <template #footer>
                <div class="flex justify-end">
                    <div class="flex items-center gap-2">
                        <div class="flex items-center">
                        订单总额：<span class="text-[red] font-bold">{{ detail?.order.totalAmount }} {{ site.symbol }}</span>
                        </div>
                        <div class="font-bold">
                        -
                        </div>
                        <div class="flex items-center">
                        优惠金额：<span class="text-[red] font-bold">{{ detail?.order.discountAmount }}  {{ site.symbol }}</span>
                        </div>
                    </div>
                </div>
                <div class="flex justify-end">
                    <div class="flex items-center">
                        实付金额：<span class="text-[red] font-bold">{{ detail?.order.actualAmount }} {{ site.symbol }}</span>
                    </div>
                </div>
                <div class="flex justify-end">
                    <div class="flex items-center">
                        已支付：<span class="text-[red] font-bold">{{ detail?.order.payAmount ? detail?.order.payAmount :  0 }} {{ site.symbol }}</span>
                    </div>
                </div>
            </template>
        </ElCard>
        <ElCard v-loading="loading" shadow="never" class="mt-[20px]">
            <template #header>
                <div class="font-bold">
                    售后号: {{ detail?.code }}
                </div>
            </template>
            <ElDescriptions>
                <ElDescriptionsItem :span="3"  label="审核人">
                    {{ detail?.manage }}
                </ElDescriptionsItem>
                <ElDescriptionsItem :span="3"  label="退款金额">
                    {{ detail?.amount }}{{ site.symbol }}
                </ElDescriptionsItem>
                <ElDescriptionsItem :span="3"  label="售后状态">
                    <ElTag :type="getAftersalesStatusConfig(detail?.status!).type">{{ getAftersalesStatusConfig(detail?.status!).text }}</ElTag>
                </ElDescriptionsItem>
                <ElDescriptionsItem :span="3"  label="售后类型">
                    <ElTag :type="getAftersalesType(detail?.type!).type">{{ getAftersalesType(detail?.type!).text }}</ElTag>
                </ElDescriptionsItem>
            </ElDescriptions>
            <div v-if="detail?.status == ApplyStatus.Fail" class="bg-primary/10 text-sm p-2 rounded">
                驳回原因:{{ detail?.reject }}
            </div>
            <template #footer>
                <div class="font-bold">
                    {{detail?.reason}}
                </div>
            </template>
        </ElCard>
    </ElDrawer>
</template>
  
<script setup lang="ts">
import { fetchGetAftersalesDetail } from '@/api/aftersales';

import { PayMode } from '@/enums/modeEnum';
import {  ApplyStatus, OrderStatus, PayStatus } from '@/enums/statusEnum';
import { AfterSalesType } from '@/enums/typeEnum';
import { useSiteStore } from '@/store/modules/site';


interface Props {
    modelValue: boolean
    id?: number | null
}
const props = withDefaults(defineProps<Props>(), {
    modelValue: false,
    id: 0
})
interface Emits {
    (e: 'update:modelValue', value: boolean): void
    (e: 'submit'): void
}
const emit = defineEmits<Emits>()
// 对话框显示控制
const visible = computed({
    get: () => props.modelValue,
    set: (value) => emit('update:modelValue', value)
})
// 售后类型配置
const AFTERSALES_TYPE_CONFIG = {
  [AfterSalesType.Cheater]: { type: 'primary' as const, text: '开挂作弊' },
  [AfterSalesType.Other]: { type: 'primary' as const, text: '其他' },
  [AfterSalesType.ServiceNotCompleted]: { type: 'primary' as const, text: '炸单' },
} as const

/**
 * 获取类型配置
 */
const getAftersalesType = (status: number) => {
  return (
    AFTERSALES_TYPE_CONFIG[status as keyof typeof AFTERSALES_TYPE_CONFIG] || {
      type: 'info' as const,
      text: '未知'
    }
  )
}

// 售后状态配置
const AFTERSALES_STATUS_CONFIG = {
  [ApplyStatus.Pending]: { type: 'info' as const, text: '待审核' },
  [ApplyStatus.Success]: { type: 'success' as const, text: '已通过' },
  [ApplyStatus.Fail]: { type: 'danger' as const, text: '拒绝通过' },
} as const

/**
 * 获取售后状态配置
 */
const getAftersalesStatusConfig = (status: number) => {
  return (
    AFTERSALES_STATUS_CONFIG[status as keyof typeof AFTERSALES_STATUS_CONFIG] || {
      type: 'info' as const,
      text: '未知'
    }
  )
}


// 订单状态配置
const ORDER_STATUS_CONFIG = {
  [OrderStatus.PendingPayment]: { type: 'info' as const, text: '待支付' },
  [OrderStatus.PendingService]: { type: 'primary' as const, text: '待服务' },
  [OrderStatus.InProgress]: { type: 'primary' as const, text: '进行中' },
  [OrderStatus.Completed]: { type: 'success' as const, text: '已完成' },
  [OrderStatus.Cancel]: { type: 'danger' as const, text: '已取消' },
  [OrderStatus.Refund]: { type: 'warning' as const, text: '已退款' },
} as const

/**
 * 获取订单状态配置
 */
const getOrderStatusConfig = (status: number) => {
  return (
    ORDER_STATUS_CONFIG[status as keyof typeof ORDER_STATUS_CONFIG] || {
      type: 'info' as const,
      text: '未知'
    }
  )
}

// 支付方式配置
const ORDER_PAY_MODE = {
  [PayMode.AlyPay]: { type: 'primary' as const, text: '支付宝' },
  [PayMode.Wechat]: { type: 'success' as const, text: '微信' },
  [PayMode.Balance]: { type: 'warning' as const, text: '余额' },
  [PayMode.PersonalTransfer]: { type: 'primary' as const, text: '个人转账' },
} as const

/**
 * 获取支付方式配置
 */
const getOrderPayMode = (mode: number) => {
  return (
    ORDER_PAY_MODE[mode as keyof typeof ORDER_PAY_MODE] || {
      type: 'info' as const,
      text: '未支付'
    }
  )
}

// 支付状态配置
const ORDER_PAY_STATUS = {
  [PayStatus.Paid]: { type: 'success' as const, text: '已支付' },
  [PayStatus.PendingPayment]: { type: 'primary' as const, text: '待支付' },
  [PayStatus.Refunded]: { type: 'warning' as const, text: '已退款' },
} as const

/**
 * 获取支付状态配置
 */
const getOrderPayStatus = (status: number) => {
  return (
    ORDER_PAY_STATUS[status as keyof typeof ORDER_PAY_STATUS] || {
      type: 'info' as const,
      text: '未支付'
    }
  )
}

const {getInfo:site} = useSiteStore()
const detail = ref<Aftersales.Response.Detail>()
const loading = ref<boolean>(false)
const getData = async () =>{
    loading.value = true
    const res = await fetchGetAftersalesDetail({id:props.id!})
    detail.value = res
    loading.value = false
}

const productData = computed(()=>{
    return detail.value ? [detail.value?.order.product] :[]
})


/**
 * 初始化表单数据
 * 根据对话框类型（新增/编辑）填充表单
 */
const init = async () => {
  await getData()
}
/**
   * 监听弹窗打开，初始化表单数据
   */
watch(
  () => props.modelValue,
  (newVal) => {
    if (newVal) init()
  }
)

/**
 * 关闭弹窗并重置表单
 */
const handleClose = () => {
  visible.value = false
}
</script>
  