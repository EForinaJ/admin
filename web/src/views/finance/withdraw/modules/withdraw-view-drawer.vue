<template>
    <ElDrawer
        v-model="visible"
        title="提现详情"
        size="35%"
        @close="handleClose"
        >
        <ElCard v-loading="loading" shadow="never">
            <template #header>
                <div class="flex items-center gap-3">
                    <div class="font-bold">
                        提现编号: {{ detail?.code }}
                    </div>
                    <el-divider direction="vertical" />
                    <div class="font-bold">
                        申请时间: {{ detail?.createTime }}
                    </div>
                </div>
            </template>
            <ElDescriptions>
                <ElDescriptionsItem :span="3"  label="申请人">
                    {{ detail?.witkey }}
                </ElDescriptionsItem>
                <ElDescriptionsItem v-if="detail?.status != ApplyStatus.Pending" :span="3"  label="审核人">
                    {{ detail?.manage }}
                </ElDescriptionsItem>
                <ElDescriptionsItem :span="3"  label="提现金额">
                    <ElTag type="primary">{{ detail?.amount }} {{ site.symbol }}</ElTag>
                </ElDescriptionsItem>
                <ElDescriptionsItem :span="3"  label="手续费">
                    <ElTag type="primary">{{ detail?.serviceFee }} {{ site.symbol }}</ElTag>
                </ElDescriptionsItem>
                <ElDescriptionsItem :span="3"  label="到账金额">
                    <ElTag type="primary">{{ detail?.settledAmount }} {{ site.symbol }}</ElTag>
                </ElDescriptionsItem>
                <ElDescriptionsItem :span="3"  label="到账平台">
                    <ElTag :type="getWithdrawType(detail?.type!).type">{{ getWithdrawType(detail?.type!).text }}</ElTag>
                </ElDescriptionsItem>
                <ElDescriptionsItem :span="3"  label="真实姓名">
                    {{ detail?.name }}
                </ElDescriptionsItem>
                <ElDescriptionsItem :span="3"  label="平台号码">
                    {{ detail?.number }}
                </ElDescriptionsItem>
                <ElDescriptionsItem :span="3"  label="提现状态">
                    <ElTag :type="getWithdrawStatusConfig(detail?.status!).type">{{ getWithdrawStatusConfig(detail?.status!).text }}</ElTag>
                </ElDescriptionsItem>
            </ElDescriptions>
            <template v-if="detail?.status == ApplyStatus.Fail" #footer>
                {{ detail?.reason }}
            </template>
        </ElCard>
    </ElDrawer>
</template>
  
<script setup lang="ts">
import { fetchGetWithdrawDetail } from '@/api/withdraw';
import {  ApplyStatus, PayStatus } from '@/enums/statusEnum';
import { WithdrawType } from '@/enums/typeEnum';
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


// 售后状态配置
const WITHDRAW_STATUS_CONFIG = {
  [ApplyStatus.Pending]: { type: 'info' as const, text: '待审核' },
  [ApplyStatus.Success]: { type: 'success' as const, text: '已通过' },
  [ApplyStatus.Fail]: { type: 'danger' as const, text: '拒绝通过' },
} as const

/**
 * 获取售后状态配置
 */
const getWithdrawStatusConfig = (status: number) => {
  return (
    WITHDRAW_STATUS_CONFIG[status as keyof typeof WITHDRAW_STATUS_CONFIG] || {
      type: 'info' as const,
      text: '未知'
    }
  )
}



// 支付方式配置
const WITHDRAW_TYPE = {
  [WithdrawType.AlyPay]: { type: 'primary' as const, text: '支付宝' },
  [WithdrawType.Wechat]: { type: 'success' as const, text: '微信' },
} as const

/**
 * 获取支付方式配置
 */
const getWithdrawType = (mode: number) => {
  return (
    WITHDRAW_TYPE[mode as keyof typeof WITHDRAW_TYPE] || {
      type: 'info' as const,
      text: '未知'
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
const detail = ref<Withdraw.Response.Detail>()
const loading = ref<boolean>(false)
const getData = async () =>{
    loading.value = true
    const res = await fetchGetWithdrawDetail({id:props.id!})
    detail.value = res
    loading.value = false
}

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
  