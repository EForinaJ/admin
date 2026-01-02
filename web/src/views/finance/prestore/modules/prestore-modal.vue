<template>
    <ElDialog
      title="预存充值"
      width="25%"
      :model-value="visible"
      align-center
      @update:model-value="handleCancel"
      @close="handleClose"
    >
      <ElForm ref="formRef" :model="form" :rules="rules" label-width="auto">
        <ElFormItem label="所属用户" prop="userId">
            <ElSelect
                v-model="form.userId"
                filterable
                remote
                reserve-keyword
                placeholder="请输入用户手机号"
                :remote-method="remoteMethod"
                :loading="loading"
        
            >
                <ElOption
                v-for="item in userOptions"
                :key="item.value"
                :label="item.label"
                :value="item.value"
                />
            </ElSelect>
        </ElFormItem>
        <ElFormItem label="预存充值" prop="amount">
            <ElInputNumber 
            :precision="2"
            style="width: 100%;"
            v-model="form.amount" placeholder="请输入预存充值" controls-position="right"> </ElInputNumber>
        </ElFormItem>
        <ElFormItem label="赠送金额" prop="bonusAmount">
            <ElInputNumber 
            :precision="2"
            style="width: 100%;"
            v-model="form.bonusAmount" placeholder="请输入赠送金额" controls-position="right"> </ElInputNumber>
        </ElFormItem>
      </ElForm>
      <template #footer>
        <ElButton @click="handleCancel">取消</ElButton>
        <ElButton type="primary" @click="handleSubmit">提交</ElButton>
      </template>
    </ElDialog>
  </template>
  
<script setup lang="ts">
import type { FormInstance, FormRules } from 'element-plus'
import { useDebounce } from '@/hooks/index';
import { fetchGetUserOptionsList } from '@/api/site';
import {  fetchPostPrestoreCreate } from '@/api/prestore';

interface Props {
    visible: boolean
}

interface Emits {
    (e: 'update:visible', value: boolean): void
    (e: 'submit'): void
}

interface ListItem {
  value: number
  label: string
}

const props = withDefaults(defineProps<Props>(), {
    visible: false,
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
 const form = reactive<Prestore.Params.Moadl>({
    userId: null,
    amount: null,
    bonusAmount: null,
})

/**
 * 表单验证规则
 */
const rules = reactive<FormRules>({
    userId: [
        { required: true, message: '请选择所属用户', trigger: 'blur' },
    ],
    amount: [
        { required: true, message: '请输入预存充值', trigger: 'blur' },
    ],
})

const { debounce } = useDebounce()
const loading = ref(false)
const userOptions = ref<ListItem[]>([])
const getUserOptions =  async (phone:string) => {
  try {
      const res = await fetchGetUserOptionsList({phone})
      userOptions.value = res.list.map((item) => {
          return {
              value: item.id,
              label: item.name
          }
      })
  } finally {
      loading.value = false
  }
}
// 创建防抖版本
const getUserOptionsDebounced = debounce((query: string) => {
  loading.value = true
  getUserOptions(query)
}, 500)
const remoteMethod = (query: string) => {
  if (query) {
    getUserOptionsDebounced(query)
  } else {
    userOptions.value = []
  }
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
        money: 0,
    })
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
        await fetchPostPrestoreCreate(form)
        ElMessage.success('添加成功')
        emit('submit')
        handleClose()
        handleCancel()
    } catch (error) {
        console.log('表单验证失败:', error)
    }
}
</script>