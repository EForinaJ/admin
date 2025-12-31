<template>
  <ArtSearchBar
    ref="searchBarRef"
    v-model="formData"
    :items="formItems"
    :rules="rules"
    @reset="handleReset"
    @search="handleSearch"
  >
  </ArtSearchBar>
</template>

<script setup lang="ts">
import { fetchGetGameOptionsList } from '@/api/site';
import { ProductStatus } from '@/enums/statusEnum';

interface Props {
  modelValue: Record<string, any>
}
interface Emits {
  (e: 'update:modelValue', value: Record<string, any>): void
  (e: 'search', params: Record<string, any>): void
  (e: 'reset'): void
}
const props = defineProps<Props>()
const emit = defineEmits<Emits>()

// 表单数据双向绑定
const searchBarRef = ref()
const formData = computed({
  get: () => props.modelValue,
  set: (val) => emit('update:modelValue', val)
})

// 校验规则
const rules = {
  // userName: [{ required: true, message: '请输入商品名', trigger: 'blur' }]
}

// 动态 options
const statusOptions = ref<{ label: string; value: number; disabled?: boolean }[]>([
{
    label:"在售",
    value: ProductStatus.Enable
  },
  {
    label:"下架",
    value: ProductStatus.Disable
  },
])


interface ListItem {
  value: number
  label: string,
}
const gameOptions = ref<ListItem[]>([])
const getGameOptions =  async () => {
  const res = await fetchGetGameOptionsList()
  gameOptions.value = res.list.map((item) => {
    return {
      value: item.id,
      label: item.name
    }
  })
}
onMounted(async () => {
  await getGameOptions()
})

// 表单配置
const formItems = computed(() => [
  {
    label: '商品名',
    key: 'name',
    type: 'input',
    placeholder: '请输入商品名',
    clearable: true
  },
  {
    label: '游戏领域',
    key: 'gameId',
    type: 'select',
    props: {
      placeholder: '请选择游戏领域',
      options: gameOptions.value
    }
  },
  {
    label: '状态',
    key: 'status',
    type: 'select',
    props: {
      placeholder: '请选择状态',
      options: statusOptions.value
    }
  },
])

// 事件
function handleReset() {
  console.log('重置表单')
  emit('reset')
}

async function handleSearch() {
    await searchBarRef.value.validate()
    emit('search', formData.value)
    console.log('表单数据', formData.value)
}
</script>
