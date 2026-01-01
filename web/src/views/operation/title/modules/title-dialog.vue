<template>
  <ElDialog
    v-model="visible"
    :title="dialogType === 'create' ? '添加分类' : '编辑分类'"
    width="25%"
    align-center
    @close="handleClose"
  >
    <ElForm  ref="formRef" :model="formData" :rules="rules" label-width="80px">
      <ElFormItem label="名称" prop="name">
        <ElInput v-model="formData.name" placeholder="请输入名称" />
      </ElFormItem>
      <ElFormItem label="所属游戏" prop="gameId">
        <ElSelect
          clearable
          style="width: 100%;"
          v-model="formData.gameId"
          placeholder="请选择游戏"
        >
          <el-option
            v-for="item in gameOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </ElSelect>
      </ElFormItem>
      <ElFormItem label="服务费率" prop="servicePercent">
        <ElInputNumber 
          :precision="0"
          :max="100"
          :step="10"
          style="width: 100%;"
          v-model="formData.servicePercent" placeholder="请输入服务费率"
          controls-position="right"> </ElInputNumber>
      </ElFormItem>
      <ElFormItem label="封面" prop="pic">
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

import { fetchGetTitleEdit, fetchPostTitleCreate, fetchPostTitleEdit,  } from '@/api/title';
import { fetchGetGameOptionsList } from '@/api/site';
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
  servicePercent: number | null;
  gameId: number | null;
  description: string | null;
}
// 表单数据
const formData = reactive<model>({
  id: 0,
  name: null,
  pic: "",
  gameId: null,
  servicePercent: null,
  description: null,
})

// 表单验证规则
const rules: FormRules = {
  name: [
    { required: true, message: '请输入分类名称', trigger: 'blur' },
  ],
  pic: [
    { required: true, message: '请上传分类封面', trigger: 'blur' },
  ],
  servicePercent: [
    { required: true, message: '请输入服务费率', trigger: 'blur' },
  ],
  gameId: [
    { required: true, message: '请设置所属游戏', trigger: 'blur' },
  ]
}

const handlePic = (e:string) => {
  formData.pic = e
}

interface ListItem {
  value: number
  label: string,
}
const loading = ref(false)
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

/**
 * 初始化表单数据
 * 根据对话框类型（新增/编辑）填充表单
 */
const initFormData = async () => {
  loading.value = true 
  await getGameOptions()
  if (props.type === 'create') {
    Object.assign(formData, {
      id: 0,
      name: null,
      pic: "",
      gameId: null,
      description: null,
      servicePercent: null,
    })
  }
  if (props.type === 'edit') {
    const res = await fetchGetTitleEdit({id:props.id!})
    Object.assign(formData, res)
  }
  loading.value = false 
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
        await fetchPostTitleCreate(formData)
      }

      if (dialogType.value === 'edit') {
        await fetchPostTitleEdit(formData)
      }

      ElMessage.success(dialogType.value === 'create' ? '添加成功' : '更新成功')
      emit('submit')
      handleClose()
    }
  })
}
</script>
