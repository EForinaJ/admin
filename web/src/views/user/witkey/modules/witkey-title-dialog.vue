<template>
  <ElDialog
    title="变更头衔"
    width="25%"
    :model-value="visible"
    align-center
    @update:model-value="handleCancel"
    @close="handleClose"
  >
    <ElForm ref="formRef" :model="form" :rules="rules" label-width="auto">
      <ElFormItem label="所属游戏" prop="gameId">
        <ElSelect
          clearable
          style="width: 100%;"
          v-model="form.gameId"
          placeholder="请选择游戏"
          @change="handleChangeGame"
        >
          <el-option
            v-for="item in gameOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </ElSelect>
      </ElFormItem>
      <ElFormItem v-if="form.gameId" label="所属头衔" prop="titleId">
        <ElSelect
          clearable
          style="width: 100%;"
          v-model="form.titleId"
          placeholder="请选择头衔"
        >
          <el-option
            v-for="item in titleOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </ElSelect>
      </ElFormItem>
    </ElForm>
    <template #footer>
      <ElButton @click="handleCancel">取消</ElButton>
      <ElButton type="primary" @click="handleSubmit">提交</ElButton>
    </template>
  </ElDialog>
</template>

<script setup lang="ts">
import { fetchGetGameOptionsList, fetchGetTitleOptionsList } from '@/api/site';
import type { FormInstance, FormRules } from 'element-plus'
import { fetchPostWitkeyChangeTitle } from '@/api/witkey';
interface Props {
  visible: boolean
  id?: number | null
}

interface Emits {
  (e: 'update:visible', value: boolean): void
  (e: 'submit'): void
}

const props = withDefaults(defineProps<Props>(), {
  visible: false,
  id: 0,
})
const emit = defineEmits<Emits>()


const formRef = ref<FormInstance>()

/**
* 弹窗显示状态双向绑定
*/
const visible = computed({
  get: () => props.visible,
  set: (value) => emit('update:visible', value)
})
/**
* 表单数据
*/
const form = reactive<Witkey.Params.ChangeTitle>({
  id: null,
  gameId: null,
  titleId: null,
})


const rules = reactive<FormRules>({
  gameId: [
      { required: true, message: '请选择游戏领域', trigger: 'blur' },
  ],
  titleId: [
      { required: true, message: '请选择头衔', trigger: 'blur' },
  ],
})


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
const handleChangeGame = async (v:any) =>{
  await getTitleOptions(v)
}
const titleOptions = ref<ListItem[]>([])
const getTitleOptions = async (id:number) => {
  const res = await fetchGetTitleOptionsList({gameId:id})
  titleOptions.value = res.list.map((item) => {
    return {
      value: item.id,
      label: item.name
    }
  })
}

/**
* 监听弹窗打开，初始化表单数据
*/
watch(
  () => props.visible,
  (newVal) => {
  if (newVal) initForm()
  }
)

/**
* 初始化表单数据
* 根据弹窗类型填充表单或重置表单
*/
const initForm = async () => {
  Object.assign(form, {
      id: props.id!,
  })
  getGameOptions()
}

/**
* 关闭弹窗并重置表单
*/
const handleClose = () => {
  formRef.value?.resetFields()
}

/**
* 取消操作
*/
const handleCancel = (): void => {
  handleClose()
  emit('update:visible', false)
}

/**
* 提交表单
* 验证通过后调用接口保存数据
*/
const handleSubmit = async () => {
  if (!formRef.value) return
  try {
      await formRef.value.validate()
      // TODO: 调用新增/编辑接口
      const res = await fetchPostWitkeyChangeTitle(form)
      ElMessage.success('更新成功')
      emit('submit')
      handleClose()
      handleCancel()
  } catch (error) {
      console.log('表单验证失败:', error)
  }
}
</script>