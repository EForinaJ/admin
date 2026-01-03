<template>
    <ElDrawer
        v-model="visible"
        title="预存详情"
        size="35%"
        @close="handleClose"
        >
        <ElCard v-loading="loading" shadow="never">
            <template #header>
                <div class="flex items-center gap-3">
                    <div class="font-bold">
                        申请时间: {{ detail?.createTime }}
                    </div>
                </div>
            </template>
            <ElDescriptions>
                <ElDescriptionsItem :span="3"  label="预存用户">
                    {{ detail?.user }}
                </ElDescriptionsItem>
                <ElDescriptionsItem  :span="3"  label="操作人">
                    {{ detail?.manage }}
                </ElDescriptionsItem>
                <ElDescriptionsItem :span="3"  label="预存金额">
                    <ElTag type="primary">{{ detail?.amount }} {{ site.symbol }}</ElTag>
                </ElDescriptionsItem>
                <ElDescriptionsItem :span="3"  label="赠送金额">
                    <ElTag type="primary">{{ detail?.bonusAmount }} {{ site.symbol }}</ElTag>
                </ElDescriptionsItem>
                <ElDescriptionsItem :span="3"  label="预存状态">
                    <ElTag :type="getPrestoreStatusConfig(detail?.status!).type">{{ getPrestoreStatusConfig(detail?.status!).text }}</ElTag>
                </ElDescriptionsItem>
            </ElDescriptions>
            <template v-if="detail?.status == ApplyStatus.Fail" #footer>
                {{ detail?.reason }}
            </template>
        </ElCard>
    </ElDrawer>
</template>
  
<script setup lang="ts">
import { fetchGetPrestoreDetail } from '@/api/prestore';
import {  ApplyStatus } from '@/enums/statusEnum';
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
const PRESTORE_STATUS_CONFIG = {
  [ApplyStatus.Pending]: { type: 'info' as const, text: '待审核' },
  [ApplyStatus.Success]: { type: 'success' as const, text: '已通过' },
  [ApplyStatus.Fail]: { type: 'danger' as const, text: '拒绝通过' },
} as const

/**
 * 获取售后状态配置
 */
const getPrestoreStatusConfig = (status: number) => {
  return (
    PRESTORE_STATUS_CONFIG[status as keyof typeof PRESTORE_STATUS_CONFIG] || {
      type: 'info' as const,
      text: '未知'
    }
  )
}

const {getInfo:site} = useSiteStore()
const detail = ref<Prestore.Response.Detail>()
const loading = ref<boolean>(false)
const getData = async () =>{
    loading.value = true
    const res = await fetchGetPrestoreDetail({id:props.id!})
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
  