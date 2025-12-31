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
  // userName: [{ required: true, message: '请输入用户名', trigger: 'blur' }]
}

// 动态 options
// const statusOptions = ref<{ label: string; value: number; disabled?: boolean }[]>([])

// // 模拟接口返回状态数据
// function fetchStatusOptions(): Promise<typeof statusOptions.value> {
//   return new Promise((resolve) => {
//     setTimeout(() => {
//       resolve([
//         { label: '启用', value: UserStatus.Enable },
//         { label: '禁用', value: UserStatus.Disable },
//       ])
//     }, 1000)
//   })
// }
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
  // statusOptions.value = await fetchStatusOptions()
})

// 表单配置
const formItems = computed(() => [
  {
    label: '分类名称',
    key: 'name',
    type: 'input',
    placeholder: '请输入分类名称',
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
