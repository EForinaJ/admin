<template>
    <ElDrawer
      v-model="visible"
      title="订单详情"
      size="50%"
      @close="handleClose"
    >
        <ElTabs 
          v-model="activeKey"
          @tab-change="changeMneu">
            <ElTabPane label="订单详情" name="detail" />
            <ElTabPane label="派单记录" name="distribute"/>
            <ElTabPane label="订单日记" name="logs"/>
        </ElTabs>
        <component :is="orderMap[activeKey]" :id="id"/>
    </ElDrawer>
  </template>
  
<script setup lang="ts">
import type { Component } from 'vue';
import { TabPaneName } from 'element-plus';
import OrderDetail from './order-detail.vue';
import OrderDistributeList from './order-distribute-list.vue';
import OrderLogList from './order-log-list.vue';
// import OrderDistributeList from './order-distribute-list.vue';
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

const activeKey = ref<string>("")
const changeMneu = (name: TabPaneName)=>{
  activeKey.value = name.toString()
}
const orderMap: Record<string, Component> = {
    "detail": OrderDetail,
    "distribute":OrderDistributeList,
    "logs":OrderLogList,
};

  
  /**
   * 初始化表单数据
   * 根据对话框类型（新增/编辑）填充表单
   */
  const init = async () => {
    activeKey.value = "detail"
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
    activeKey.value = ""
    emit('submit')
  }
</script>
  