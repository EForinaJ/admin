<template>
    <ElDialog
      title="变更余额"
      width="25%"
      :model-value="visible"
      align-center
      @update:model-value="handleCancel"
      @close="handleClose"
    >
      <ElForm ref="formRef" :model="form" :rules="rules" label-width="auto">
        <ElFormItem label="变更方式" prop="mode">
            <ElSelect v-model="form.mode">
                <ElOption label="减少" :value="FundMode.Sub" />
                <ElOption label="增加" :value="FundMode.Add" />
            </ElSelect>
        </ElFormItem>
        <ElFormItem label="变更金额" prop="amount">
            <ElInputNumber 
            :precision="2"
            style="width: 100%;"
            v-model="form.amount" placeholder="请输入变更金额" controls-position="right"> </ElInputNumber>
        </ElFormItem>
        <ElFormItem label="变更备注" prop="remark">
          <ElInput
            v-model="form.remark"
            type="textarea"
            :rows="3"
            placeholder="请输入变更备注"
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

import { fetchPostUserChangeBalance } from '@/api/user';
import { FundMode } from '@/enums/modeEnum';
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
 const form = reactive<User.Params.ChangeBalance>({
    id: 0, // 权限ID
    amount: null,
    mode: null,
    remark: null,
})

/**
 * 表单验证规则
 */
const rules = reactive<FormRules>({
    amount: [
        { required: true, message: '请输入变更金额', trigger: 'blur' },
    ],
    mode: [
        { required: true, message: '请输入变更方式', trigger: 'blur' },
    ],
    remark: [
        { required: true, message: '请输入变更备注', trigger: 'blur' },
    ]
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
        id: props.id!, // 权限ID
        amount: null,
        mode: null,
        remark: null,
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
        await fetchPostUserChangeBalance(form)


        ElMessage.success('更新成功')
        emit('submit')
        handleClose()
        handleCancel()
    } catch (error) {
        console.log('表单验证失败:', error)
    }
}
</script>