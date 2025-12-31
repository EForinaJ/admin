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
            <ElTabPane label="用户详情" name="detail" />
            <ElTabPane label="余额变更" name="balance"/>
        </ElTabs>
        <component :is="userMap[activeKey]" :id="id"/>
    </ElDrawer>
  </template>
  
<script setup lang="ts">
import type { Component } from 'vue';
import UserDetail from './user-detail.vue';
import { TabPaneName } from 'element-plus';
import UserBalanceList from './user-balance-list.vue';
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
const userMap: Record<string, Component> = {
    "detail": UserDetail,
    "balance":UserBalanceList,
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
  