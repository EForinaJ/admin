<template>
  <ElDialog
    :title="dialogType === 'create' ? '新增权限' : '编辑权限'"
    width="30%"
    :model-value="visible"
    align-center
    @update:model-value="handleCancel"
    @close="handleClose"
  >
    <ElForm ref="formRef" :model="form" :rules="rules" label-width="auto">
      <ElFormItem label="权限名称" prop="name">
        <ElInput v-model="form.name" placeholder="请输入名称" />
      </ElFormItem>
      <ElFormItem label="权限标识" prop="permission">
        <ElInput v-model="form.permission" placeholder="请输入权限标识" />
      </ElFormItem>
      <ElFormItem label="权限描述" prop="description">
        <ElInput
          v-model="form.description"
          type="textarea"
          :rows="3"
          placeholder="请输入角色描述"
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
  import { fetchPermissionCreate, fetchPermissionEdit, fetchPermissionGetEdit } from '@/api/permission';
import type { FormInstance, FormRules } from 'element-plus'


  interface Props {
    visible: boolean
    dialogType: 'create' | 'edit'
    id?: number | null
    isChild: boolean
  }

  interface Emits {
    (e: 'update:visible', value: boolean): void
    (e: 'update:isChild', value: boolean): void
    (e: 'submit'): void
  }

  const props = withDefaults(defineProps<Props>(), {
    visible: false,
    dialogType: 'create',
    id: 0,
    isChild: false,
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
   * 表单验证规则
   */
  const rules = reactive<FormRules>({
    name: [
      { required: true, message: '请输入权限名称', trigger: 'blur' },
    ],
    permission: [
      { required: true, message: '请输入权限标识', trigger: 'blur' },
    ]
  })

  /**
   * 表单数据
   */
  const form = reactive<Permission.Params.Create>({
    id: 0, // 权限ID
    pId: 0, // 父Id
    permission: "", // 权限标识
    name: "", // 权限名称
    description: "", // 权限描述
  })

  /**
   * 监听弹窗打开，初始化表单数据
   */
  watch(
    () => props.visible,
    (newVal) => {
      if (props.isChild) {
        form.pId = props.id!
      }
      if (newVal) initForm()
    }
  )

  /**
   * 监听角色数据变化，更新表单
   */
  // watch(
  //   () => props.roleData,
  //   (newData) => {
  //     if (newData && props.visible) initForm()
  //   },
  //   { deep: true }
  // )

  /**
   * 初始化表单数据
   * 根据弹窗类型填充表单或重置表单
   */
  const initForm = async () => {
    if (props.dialogType === 'edit') {
      const res = await fetchPermissionGetEdit({id:props.id!})
      form.id = res.id
      form.pId = res.pId
      form.permission = res.permission
      form.name = res.name
      form.description = res.description
    } else {
      Object.assign(form, {
        id: 0, // 权限ID
        permission: "", // 权限标识
        name: "", // 权限名称
        description: "", // 权限描述
      })
    }
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
  emit('update:visible', false)
  emit('update:isChild', false)
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
      if (props.dialogType === "create") {
        await fetchPermissionCreate(form)
      }
      if (props.dialogType === 'edit') {
        await fetchPermissionEdit(form)
      }

      const message = props.dialogType === 'create' ? '新增成功' : '修改成功'
      ElMessage.success(message)
      emit('submit')
      handleClose()
      handleCancel()
    } catch (error) {
      console.log('表单验证失败:', error)
    }
  }
</script>
