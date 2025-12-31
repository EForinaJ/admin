<template>
    <ElCard v-loading="loading" shadow="never">
        <template #header>
            <div class="flex items-center gap-3">
                <ElAvatar :src="detail?.user.avatar" shape="square" :size="60" />
                <div class="flex-col flex gap-2">
                    <h2 class="font-bold">
                        {{detail?.user.name}}
                    </h2>
                    <p>{{ detail?.user.phone }}</p>
                </div>
            </div>
        </template>
        <ElDescriptions>
            <ElDescriptionsItem :span="3"  label="所属游戏">
                <ElTag type="primary">{{detail?.game}}</ElTag>
            </ElDescriptionsItem>
            <ElDescriptionsItem :span="3"  label="所得头衔">
                <ElTag type="primary">{{detail?.title}}</ElTag>
            </ElDescriptionsItem>
            <ElDescriptionsItem :span="3"  label="评分">
                <ElRate showScore
                :scoreTemplate="`${detail?.rate}分`"
                disabled
                :model-value="detail?.rate"
                />
            </ElDescriptionsItem>
            <ElDescriptionsItem :span="3"  label="入驻时间">
                {{ detail?.createTime }}
            </ElDescriptionsItem>
        </ElDescriptions>
        <template #footer>
            <ElSpace wrap>
                <div v-for="(item) in detail?.album" 
                    class="w-[100px] h-[100px]">
                    <ElImage
                        class="rounded-lg"
                        :src="item"
                        :preview-src-list="[item]"
                        >
                    </ElImage>
                </div>
            </ElSpace>
        </template>
    </ElCard>

    <ElCard v-loading="loading" shadow="never" class="mt-[20px]">
        <ElDescriptions  direction="vertical">
            <ElDescriptionsItem :span="1"  label="佣金">
                <span class="text-xl font-bold">{{ detail?.commission }} {{ site.symbol }}</span>
            </ElDescriptionsItem>
        </ElDescriptions>
    </ElCard>
</template>

<script setup lang="ts">
import { fetchGetWitkeyDetail } from '@/api/witkey';
import { SexType } from '@/enums/typeEnum';
import { useSiteStore } from '@/store/modules/site';
import { codeToText } from 'element-china-area-data';
import { ElRate } from 'element-plus';

interface Props {
  id?: number;
}

const props = defineProps<Props>();

const {getInfo:site} = useSiteStore()
const detail = ref<Witkey.Response.Detail>()
const loading = ref<boolean>(false)
const getData = async () =>{
    loading.value = true
    const res = await fetchGetWitkeyDetail({id:props.id!})
    detail.value = res
    loading.value = false
}

const sex = computed(()=>{
    // const sexMap: Record<number, string> = {
    //     [SexType.Other]: "其他",
    //     [SexType.Female]: "女",
    //     [SexType.Male]: "男",
    // };
    // return sexMap[detail.value?.sex!]
})

// const address = computed(()=>{
//     if (detail.value!.address.length  == 2) {
//         return codeToText[detail.value!.address[0]]  + '/' + codeToText[detail.value!.address[1]]
//     }
//     return codeToText[detail.value!.address[0]]
// })



onMounted(()=>{
    getData()
})
</script>