<template>
    <div class="page-content">
        <ElTabs v-model="activeKey"
        @tab-change="changeMneu">
            <ElTabPane label="小程序配置" name="miniProgram"></ElTabPane>
            <ElTabPane label="支付配置" name="pay"></ElTabPane>
        </ElTabs>
        <component :is="map[activeKey]"/>
    </div>
</template>
    
<script setup lang="ts">
import MiniProgram from './modules/mini-program.vue';
import PaySetting from './modules/pay-setting.vue';

import { TabPaneName } from 'element-plus';
import type { Component } from 'vue';

const activeKey = ref<string>("")
const changeMneu = (name: TabPaneName)=>{
  activeKey.value = name.toString()
}
const map: Record<string, Component> = {
    "miniProgram": MiniProgram,
    "pay": PaySetting,
};
onMounted(()=>{
    activeKey.value = "miniProgram"
})
</script>