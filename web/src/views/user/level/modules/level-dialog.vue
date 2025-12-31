<template>
  <ElDialog
    v-model="visible"
    :title="drawerType === 'create' ? '添加等级' : '编辑等级'"
    width="25%"
    align-center
    @close="handleClose"
  >
    <ElForm  ref="formRef" :model="formData" :rules="rules" label-width="80px">
      <ElFormItem label="名称" prop="name">
        <ElInput v-model="formData.name" placeholder="请输入名称" />
      </ElFormItem>
      <ElFormItem label="等级" prop="level">
        <ElInput v-model="formData.level" placeholder="请输入等级" />
      </ElFormItem>
      <ElFormItem label="经验值" prop="experience">
        <ElInputNumber 
            :precision="0"
            style="width: 100%;"
            v-model="formData.experience" placeholder="请输入经验值"
            controls-position="right"> </ElInputNumber>
      </ElFormItem>
      <ElFormItem label="图标" prop="icon">
        <ArtUpload
          @success="handleIcon"
        >
          <ElImage
          class="w-[80px]"
          shape="square"
          v-if="formData.icon != ''" :src="formData.icon!" />
          <ElButton v-if="formData.icon == ''" type="primary">选择文件</ElButton>
        </ArtUpload>
      </ElFormItem>
    </ElForm>
    <template #footer>
      <div class="drawer-footer">
        <ElButton @click="handleClose">取消</ElButton>
        <ElButton type="primary" @click="handleSubmit">提交</ElButton>
      </div>
    </template>
  </ElDialog>
</template>

<script setup lang="ts">

import { fetchGetLevelEdit, fetchPostLevelCreate, fetchPostLevelEdit,  } from '@/api/level';
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

const drawerType = computed(() => props.type)

// 表单实例
const formRef = ref<FormInstance>()
interface model {
  id?: number;
  name: string | null;
  level: number | null;
  icon: string | null;
  experience: number | null;
}
// 表单数据
const formData = reactive<model>({
  id: 0,
  name: null,
  level: null,
  icon: "",
  experience: null,
})

// 表单验证规则
const rules: FormRules = {
  name: [
    { required: true, message: '请输入等级名称', trigger: 'blur' },
  ],
  level: [
    { required: true, message: '请输入等级', trigger: 'blur' },
  ],
  icon: [
    { required: true, message: '请上传等级图标', trigger: 'blur' },
  ]
}

const handleIcon = (e:string) => {
  formData.icon = e
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
      level: null,
      icon: "",
      experience: null,
    })
  }
  if (props.type === 'edit') {
    const res = await fetchGetLevelEdit({id:props.id!})
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
      if (drawerType.value === 'create') {
        await fetchPostLevelCreate(formData)
      }

      if (drawerType.value === 'edit') {
        await fetchPostLevelEdit(formData)
      }

      ElMessage.success(drawerType.value === 'create' ? '添加成功' : '更新成功')
      emit('submit')
      handleClose()
    }
  })
}
</script>
