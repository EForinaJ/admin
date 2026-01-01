<template>
  <ElDialog
    v-model="visible"
    :title="dialogType === 'create' ? '新增角色' : '编辑角色'"
    width="30%"
    align-center
    @close="handleClose"
  >
    <ElForm ref="formRef" :model="form" :rules="rules" label-width="120px">
      <ElFormItem label="角色名称" prop="name">
        <ElInput v-model="form.name" placeholder="请输入角色名称" />
      </ElFormItem>
      <ElFormItem label="角色编码" prop="code">
        <ElInput v-model="form.code" placeholder="请输入角色编码" />
      </ElFormItem>
      <ElFormItem label="角色类型" prop="type">
        <ElSelect v-model="form.type">
          <ElOption label="超管" :value="RoleType.Super" />
          <ElOption label="客服" :value="RoleType.Kefu" />
        </ElSelect>
      </ElFormItem>
      <ElFormItem label="描述" prop="description">
        <ElInput
          v-model="form.description"
          type="textarea"
          :rows="3"
          placeholder="请输入角色描述"
        />
      </ElFormItem>
    </ElForm>
    <template #footer>
      <ElButton @click="handleClose">取消</ElButton>
      <ElButton type="primary" @click="handleSubmit">提交</ElButton>
    </template>
  </ElDialog>
</template>

<script setup lang="ts">
  import { fetchRoleCreate, fetchRoleEdit, fetchRoleGetEdit } from '@/api/role';
import { RoleType } from '@/enums/typeEnum';
import type { FormInstance, FormRules } from 'element-plus'


  interface Props {
    modelValue: boolean
    dialogType: 'create' | 'edit'
    id?: number | null
  }

  interface Emits {
    (e: 'update:modelValue', value: boolean): void
    (e: 'submit'): void
  }

  const props = withDefaults(defineProps<Props>(), {
    modelValue: false,
    dialogType: 'create',
    id: 0
  })

  const emit = defineEmits<Emits>()

  const formRef = ref<FormInstance>()

  /**
   * 弹窗显示状态双向绑定
   */
  const visible = computed({
    get: () => props.modelValue,
    set: (value) => emit('update:modelValue', value)
  })

  /**
   * 表单验证规则
   */
  const rules = reactive<FormRules>({
    name: [
      { required: true, message: '请输入角色名称', trigger: 'blur' },
      { min: 2, max: 20, message: '长度在 2 到 20 个字符', trigger: 'blur' }
    ],
    code: [
      { required: true, message: '请输入角色编码', trigger: 'blur' },
      { min: 2, max: 50, message: '长度在 2 到 50 个字符', trigger: 'blur' }
    ],
    type: [
      { required: true, message: '请选择角色类型', trigger: 'blur' },
    ],
  })

  /**
   * 表单数据
   */
  const form = reactive<Role.Params.Model>({
    id: 0,
    name: '',
    code: '',
    type: RoleType.Kefu,
    description: '',
  })

  /**
   * 监听弹窗打开，初始化表单数据
   */
  watch(
    () => props.modelValue,
    (newVal) => {
      if (newVal) initForm()
    }
  )

  /**
   * 初始化表单数据
   * 根据弹窗类型填充表单或重置表单
   */
  const initForm = async () => {
    if (props.dialogType === 'edit') {
      const res = await fetchRoleGetEdit({id:props.id!})
      Object.assign(form, res)
    } else {
      Object.assign(form, {
        id: 0,
        name: '',
        code: '',
        type: RoleType.Kefu,
        description: '',
        createTime: '',

      })
    }
  }

  /**
   * 关闭弹窗并重置表单
   */
  const handleClose = () => {
    visible.value = false
    formRef.value?.resetFields()
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
      if (props.dialogType === 'create') {
        await fetchRoleCreate(form)
      }
      // TODO: 调用新增/编辑接口
      if (props.dialogType === 'edit') {
            await fetchRoleEdit(form)
      }
      const message = props.dialogType === 'create' ? '新增成功' : '修改成功'
      ElMessage.success(message)
      emit('submit')
      handleClose()
    } catch (error) {
      console.log('表单验证失败:', error)
    }
  }
</script>
