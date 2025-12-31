<template>
  <ElDialog
    v-model="visible"
    :title="dialogType === 'create' ? '添加游戏' : '编辑游戏'"
    width="25%"
    align-center
    @close="handleClose"
  >
    <ElForm  ref="formRef" :model="formData" :rules="rules" label-width="80px">
      <ElFormItem label="名称" prop="name">
        <ElInput v-model="formData.name" placeholder="请输入名称" />
      </ElFormItem>
      <ElFormItem label="封面" prop="icon">
        <ArtUpload
          @success="handlePic"
        >
          <ElImage
          class="w-[80px]"
          shape="square"
          v-if="formData.pic != ''" :src="formData.pic!" />
          <ElButton v-if="formData.pic == ''" type="primary">选择文件</ElButton>
        </ArtUpload>
      </ElFormItem>
      <ElFormItem label="描述" prop="description">
        <ElInput 
        :rows="2"
        type="textarea"
        v-model="formData.description" placeholder="请输入描述" />
      </ElFormItem>
    </ElForm>
    <template #footer>
      <div class="dialog-footer">
        <ElButton @click="handleClose">取消</ElButton>
        <ElButton type="primary" @click="handleSubmit">提交</ElButton>
      </div>
    </template>
  </ElDialog>
</template>

<script setup lang="ts">

import { fetchGetGameEdit, fetchPostGameCreate, fetchPostGameEdit,  } from '@/api/game';
import type { FormInstance, FormRules  } from 'element-plus'


interface Props {
  modelValue: boolean
  type: "create" | "edit"
  id?: number | null
}
const props = withDefaults(defineProps<Props>(), {
  modelValue: false,
  type: 'create',
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

const dialogType = computed(() => props.type)

// 表单实例
const formRef = ref<FormInstance>()
interface model {
  id?: number;
  name: string | null;
  pic: string | null;
  description: string | null;
}
// 表单数据
const formData = reactive<model>({
  id: 0,
  name: null,
  pic: "",
  description: null,
})

// 表单验证规则
const rules: FormRules = {
  name: [
    { required: true, message: '请输入游戏名称', trigger: 'blur' },
  ],
  pic: [
    { required: true, message: '请上传游戏封面', trigger: 'blur' },
  ]
}

const handlePic = (e:string) => {
  formData.pic = e
}



/**
 * 初始化表单数据
 * 根据对话框类型（新增/编辑）填充表单
 */
const initFormData = async () => {

  if (props.type === 'create') {
    Object.assign(formData, {
      id: 0,
      name: null,
      pic: "",
      description: null,
    })
  }
  if (props.type === 'edit') {
    const res = await fetchGetGameEdit({id:props.id!})
    Object.assign(formData, res)
  }
}
/**
   * 监听弹窗打开，初始化表单数据
   */
watch(
  () => props.modelValue,
  (newVal) => {
    if (newVal) initFormData()
  }
)

/**
 * 关闭弹窗并重置表单
 */
const handleClose = () => {
  visible.value = false
  formRef.value?.resetFields()
}
/**
 * 提交表单
 * 验证通过后触发提交事件
 */
const handleSubmit = async () => {
  
  if (!formRef.value) return
  await formRef.value.validate(async (valid) => {
    if (valid) {
      if (dialogType.value === 'create') {
        await fetchPostGameCreate(formData)
      }

      if (dialogType.value === 'edit') {
        await fetchPostGameEdit(formData)
      }

      ElMessage.success(dialogType.value === 'create' ? '添加成功' : '更新成功')
      emit('submit')
      handleClose()
    }
  })
}
</script>
