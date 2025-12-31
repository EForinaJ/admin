<template>
    <div class="page-content">
        <ElTabs v-model="activeKey"
        @tab-change="changeMneu">
            <ElTabPane label="用户协议"  name="userAgreement"></ElTabPane>
            <ElTabPane label="隐私政策"  name="privacyPolicy"></ElTabPane>
            <ElTabPane label="关于我们"  name="about"></ElTabPane>
        </ElTabs>
        <component :is="map[activeKey]"/>
    </div>
</template>
    
<script setup lang="ts">
import AboutUs from './modules/about-us.vue';
import PrivacyPolicy from './modules/privacy-policy.vue';
import UserAgreement from './modules/user-agreement.vue';


import { TabPaneName } from 'element-plus';
import type { Component } from 'vue';

const activeKey = ref<string>("")
const changeMneu = (name: TabPaneName)=>{
  activeKey.value = name.toString()
}
const map: Record<string, Component> = {
    "about": AboutUs,
    "privacyPolicy": PrivacyPolicy,
    "userAgreement": UserAgreement,
};
onMounted(()=>{
    activeKey.value = "userAgreement"
})
</script>