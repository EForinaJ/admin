<template>
<div class="page-content">
    <ElTabs 
    v-model="activeKey"
    @tab-change="changeMneu">
        <ElTabPane label="基础信息" name="base"></ElTabPane>
        <ElTabPane label="用户设置" name="user"></ElTabPane>
        <ElTabPane label="提现设置" name="withdraw"></ElTabPane>
        <ElTabPane label="文件上传" name="file"></ElTabPane>
    </ElTabs>
    <component :is="map[activeKey]"/>
</div>
</template>

<script setup lang="ts">
import BaseSetting from './modules/base-setting.vue';
import UserSetting from './modules/user-setting.vue';
import FileSetting from './modules/file-setting.vue';
import WithdrawSetting from './modules/withdraw-setting.vue';
import { TabPaneName } from 'element-plus';
import type { Component } from 'vue';

const activeKey = ref<string>("")
const changeMneu = (name: TabPaneName)=>{
  activeKey.value = name.toString()
}
const map: Record<string, Component> = {
    "base": BaseSetting,
    "user": UserSetting,
    "withdraw": WithdrawSetting,
    "file": FileSetting,
};
onMounted(()=>{
    activeKey.value = "base"
})
</script>