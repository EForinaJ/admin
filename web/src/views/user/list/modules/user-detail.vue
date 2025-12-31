<template>
    <ElCard v-loading="loading" shadow="never">
        <template #header>
            <div class="flex items-center gap-3">
                <ElAvatar :src="detail?.avatar" shape="square" :size="60" />
                <div class="flex-col flex gap-2">
                    <h2 class="font-bold">
                        {{detail?.name}}
                    </h2>
                    <img :src="detail?.level.icon" class="w-[48px] h-[20px]">
                </div>
            </div>
        </template>
        <ElDescriptions>
            <ElDescriptionsItem :span="3"  label="手机号码">
                {{ detail?.phone }}
            </ElDescriptionsItem>
            <ElDescriptionsItem :span="3"  label="地址">
                {{ detail != null ? address : '无' }}
            </ElDescriptionsItem>
            <ElDescriptionsItem :span="3"  label="性别">
                <ElTag type="primary">{{sex}}</ElTag>
            </ElDescriptionsItem>
            <ElDescriptionsItem :span="3"  label="出生日期">
                {{ detail?.birthday }}
            </ElDescriptionsItem>
            <ElDescriptionsItem :span="3"  label="注册时间">
                {{ detail?.createTime }}
            </ElDescriptionsItem>
        </ElDescriptions>
    </ElCard>

    <ElCard v-loading="loading" shadow="never" class="mt-[20px]">
        <ElDescriptions  direction="vertical">
            <ElDescriptionsItem :span="1"  label="余额">
                <span class="text-xl font-bold">{{ detail?.balance }} {{ site.symbol }}</span>
            </ElDescriptionsItem>
        </ElDescriptions>
    </ElCard>
</template>

<script setup lang="ts">
import { fetchGetUserDetail } from '@/api/user';
import { SexType } from '@/enums/typeEnum';
import { useSiteStore } from '@/store/modules/site';
import { codeToText } from 'element-china-area-data';

interface Props {
  id?: number;
}

const props = defineProps<Props>();

const {getInfo:site} = useSiteStore()
const detail = ref<User.Response.Detail>()
const loading = ref<boolean>(false)
const getData = async () =>{
    loading.value = true
    const res = await fetchGetUserDetail({id:props.id!})
    detail.value = res
    loading.value = false
}

const sex = computed(()=>{
    const sexMap: Record<number, string> = {
        [SexType.Other]: "其他",
        [SexType.Female]: "女",
        [SexType.Male]: "男",
    };
    return sexMap[detail.value?.sex!]
})

const address = computed(()=>{
    if (detail.value!.address.length  == 2) {
        return codeToText[detail.value!.address[0]]  + '/' + codeToText[detail.value!.address[1]]
    }
    return codeToText[detail.value!.address[0]]
})



onMounted(()=>{
    getData()
})
</script>