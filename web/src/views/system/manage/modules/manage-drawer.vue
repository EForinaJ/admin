<template>
  <ElDrawer
    v-model="visible"
    :title="drawerType === 'create' ? '添加员工' : '编辑员工'"
    width="30%"
    @close="handleClose"
  >
    <ElForm ref="formRef" :model="formData" :rules="rules" label-width="80px">
      <ElFormItem label="员工名" prop="name">
        <ElInput v-model="formData.name" placeholder="请输入员工名" />
      </ElFormItem>
      <ElFormItem label="手机号" prop="phone">
        <ElInput v-model="formData.phone" placeholder="请输入手机号" />
      </ElFormItem>
      <ElFormItem label="密码" prop="password">
        <ElInput v-model="formData.password" placeholder="请输入密码" />
      </ElFormItem>
      <ElFormItem label="性别" prop="sex">
        <ElSelect v-model="formData.sex">
          <ElOption label="男" :value="SexType.Male" />
          <ElOption label="女" :value="SexType.Female" />
          <ElOption label="其他" :value="SexType.Other" />
        </ElSelect>
      </ElFormItem>
      <ElFormItem label="角色" prop="roles">
        <ElSelect v-model="formData.roles" multiple>
          <ElOption
            v-for="role in roleList"
            :key="role.code"
            :value="role.id"
            :label="role.name"
          />
        </ElSelect>
      </ElFormItem>
      <ElFormItem label="头像" prop="avatar">
        <ArtUpload
        @success="handleAvatar"
        >
          <ElAvatar 
          shape="square"
          v-if="formData.avatar != ''" :size="80" :src="formData.avatar" />
          <ElButton v-if="formData.avatar == ''" type="primary">选择文件</ElButton>
        </ArtUpload>
      </ElFormItem>
      <ElFormItem label="状态" prop="status">
        <ElRadioGroup v-model="formData.status">
          <ElRadio :value="ManageStatus.Disable">禁用</ElRadio>
          <ElRadio :value="ManageStatus.Enable">启用</ElRadio>
        </ElRadioGroup>
      </ElFormItem>
    </ElForm>
    <template #footer>
      <div class="drawer-footer">
        <ElButton @click="handleClose">取消</ElButton>
        <ElButton type="primary" @click="handleSubmit">提交</ElButton>
      </div>
    </template>
  </ElDrawer>
</template>

<script setup lang="ts">
import { fetchManageAllRole, fetchManageCreate, fetchManageEdit, fetchManageGetEdit } from '@/api/manage';
import { ManageStatus } from '@/enums/statusEnum';
import { SexType } from '@/enums/typeEnum';
import type { FormInstance, FormRules } from 'element-plus'



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

// 表单数据
const formData = reactive({
  id: 0,
  name:  '',
  wechat:  '',
  password: '',
  phone: '',
  sex:  0,
  avatar:  '',
  roles: [],
  status: ManageStatus.Enable,
})

// 表单验证规则
const rules: FormRules = {
  name: [
    { required: true, message: '请输入员工名', trigger: 'blur' },
    { min: 2, max: 20, message: '长度在 2 到 20 个字符', trigger: 'blur' }
  ],
  phone: [
    { required: true, message: '请输入手机号', trigger: 'blur' },
    { pattern: /^1[3-9]\d{9}$/, message: '请输入正确的手机号格式', trigger: 'blur' }
  ],
  sex: [{ required: true, message: '请选择性别', trigger: 'blur' }],
  roles: [{ required: true, message: '请选择角色', trigger: 'blur' }]
}

const handleAvatar = (e:string) => {
  formData.avatar = e
}

// 角色列表数据
const roleList = ref<Role.Response.Info[]>([])
/**
 * 初始化表单数据
 * 根据对话框类型（新增/编辑）填充表单
 */
const initFormData = async () => {
  const roleRes = await fetchManageAllRole()
  roleList.value = roleRes != null ? roleRes.list : []
  if (props.type === 'create') {
    Object.assign(formData, {
      id: 0,
      name:  '',
      wechat:  '',
      description:  '',
      password: '',
      phone: '',
      sex:  SexType.Other,
      avatar:  '',
      roles: [],
      status: ManageStatus.Enable,
    })
  }
  if (props.type === 'edit') {
    const res = await fetchManageGetEdit({id:props.id!})
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
        await fetchManageCreate(formData)
      }

      if (drawerType.value === 'edit') {
        await fetchManageEdit(formData)
      }

      ElMessage.success(drawerType.value === 'create' ? '添加成功' : '更新成功')
      emit('submit')
      handleClose()
    }
  })
}
</script>
