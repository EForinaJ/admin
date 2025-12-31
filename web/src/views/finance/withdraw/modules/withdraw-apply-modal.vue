<template>
    <ElDialog
      title="审核提现"
      width="25%"
      :model-value="visible"
      align-center
      @update:model-value="handleCancel"
      @close="handleClose"
    >
      <ElForm ref="formRef" :model="form" :rules="rules" label-width="auto">
        <ElFormItem label="提现类型" prop="status">
            <ElRadioGroup  v-model="form.status">
                <ElRadio :value="ApplyStatus.Success">同意提现</ElRadio>
                <ElRadio :value="ApplyStatus.Fail">拒绝提现</ElRadio>
            </ElRadioGroup>
        </ElFormItem>
        <ElFormItem v-if="form.status == ApplyStatus.Fail" label="拒绝原因" prop="reason">
            <ElInput
            v-model="form.reason"
            type="textarea"
            :rows="3"
            placeholder="请输入拒绝原因"
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

import { fetchPostWithdrawApply } from '@/api/withdraw';
import { ApplyStatus } from '@/enums/statusEnum';
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
const form = reactive<Withdraw.Params.Apply>({
    id: 0, // 权限ID
    status: ApplyStatus.Success,
    reason: null,
})


const rules = reactive<FormRules>({
    status: [
        { required: true, message: '请选择提现类型', trigger: 'blur' },
    ],
    reason: [
        {
            required: form.status == ApplyStatus.Fail, // 动态判断是否必填
            message: '请输入拒绝原因',
            trigger: 'blur',
            validator: (rule, value, callback) => {
                if (form.status == ApplyStatus.Fail) {
                    if (value === undefined || value === null || value === '') {
                        callback(new Error('请输入拒绝原因'));
                    } else {
                        callback();
                    }
                } else {
                    callback();
                }
            }
        }
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
        await fetchPostWithdrawApply(form)

        ElMessage.success('审核成功')
        emit('submit')
        handleClose()
        handleCancel()
    } catch (error) {
        console.log('表单验证失败:', error)
    }
}
</script>