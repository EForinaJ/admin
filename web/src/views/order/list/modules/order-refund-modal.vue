<template>
    <ElDialog
      title="退款申请"
      width="25%"
      :model-value="visible"
      align-center
      @update:model-value="handleCancel"
      @close="handleClose"
    >
      <ElForm ref="formRef" :model="form" :rules="rules" label-width="auto">
        <ElFormItem label="退款类型" prop="type">
            <ElRadioGroup  v-model="form.type">
                <ElRadio :value="AfterSalesType.Cheater">开挂作弊</ElRadio>
                <ElRadio :value="AfterSalesType.ServiceNotCompleted">炸单</ElRadio>
                <ElRadio :value="AfterSalesType.Other">其他</ElRadio>
            </ElRadioGroup>
        </ElFormItem>
        <ElFormItem label="退款金额" prop="amount">
            <ElInputNumber 
            :precision="2"
            style="width: 100%;"
            v-model="form.amount"
            placeholder="请输入退款金额" 
            controls-position="right"/>
        </ElFormItem>
        <ElFormItem label="退款原因" prop="reason">
            <ElInput
            v-model="form.reason"
            type="textarea"
            :rows="3"
            placeholder="请输入退款原因"
          />
        </ElFormItem>
      </ElForm>
      <template #footer>
        <ElButton @click="handleCancel">取消</ElButton>
        <ElButton type="primary" @click="handleSubmit">提交</ElButton>
      </template>
    </ElDialog>
  </template>
  
<script setup lang="ts">
import { fetchPostOrderRefund } from '@/api/order';
import { AfterSalesType } from '@/enums/typeEnum';
import type { FormInstance, FormRules } from 'element-plus'


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
const form = reactive<Order.Params.Refund>({
    id: 0, // 权限ID
    type: AfterSalesType.ServiceNotCompleted,
    amount: null,
    reason: null,
})

/**
 * 表单验证规则
 */
const rules = reactive<FormRules>({
    type: [
        { required: true, message: '请选择退款类型', trigger: 'blur' },
    ],
    amount: [
        { required: true, message: '请输入退款金额', trigger: 'blur' },
    ],
    reason: [
        { required: true, message: '请输入退款原因', trigger: 'blur' },
    ],
})


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
        amount: 0,
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
        await fetchPostOrderRefund(form)

        ElMessage.success('更新成功')
        emit('submit')
        handleClose()
        handleCancel()
    } catch (error) {
        console.log('表单验证失败:', error)
    }
}
</script>