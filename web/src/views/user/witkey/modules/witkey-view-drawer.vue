<template>
    <ElDrawer
      v-model="visible"
      title="用户详情"
      size="50%"
      @close="handleClose"
    >
        <ElTabs 
          v-model="activeKey"
          @tab-change="changeMneu">
            <ElTabPane label="威客详情" name="detail" />
            <ElTabPane label="佣金变更" name="balance"/>
            <ElTabPane label="接单记录" name="recharge"/>
        </ElTabs>
        <component :is="witkeyMap[activeKey]" :id="id"/>
    </ElDrawer>
  </template>
  
<script setup lang="ts">
import type { Component } from 'vue';
import WitkeyDetail from './witkey-detail.vue';
import { TabPaneName } from 'element-plus';
import WitkeyBalanceList from './witkey-commission-list.vue';
// import WitkeyRechargeList from './witkey-recharge-list.vue';
// import WitkeyLogs from './witkey-logs.vue';
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
const witkeyMap: Record<string, Component> = {
    "detail": WitkeyDetail,
    "balance":WitkeyBalanceList,
    // "recharge": WitkeyRechargeList,
    // "log": WitkeyLogs,
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
  }
  </script>
  