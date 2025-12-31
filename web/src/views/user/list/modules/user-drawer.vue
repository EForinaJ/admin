<template>
  <ElDrawer
    v-model="visible"
    :title="drawerType === 'create' ? '添加用户' : '编辑用户'"
    width="30%"
    @close="handleClose"
  >
    <ElForm ref="formRef" :model="formData" :rules="rules" label-width="80px">
      <ElFormItem label="用户名" prop="name">
        <ElInput v-model="formData.name" placeholder="请输入用户名" />
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
      <ElFormItem label="所在地" prop="address">
        <ElCascader
          style="width: 100%;"
          v-model="formData.address"
          :options="region"
          placeholder="请选择省市区"
        />
      </ElFormItem>
      <ElFormItem label="出生日期" prop="birthday">
        <ElDatePicker
          style="width: 100%;"
          v-model="formData.birthday"
          type="date"
          placeholder="请选择出生日期"
          format="YYYY/MM/DD"
          value-format="x"
        />
      </ElFormItem>
      <ElFormItem label="个人介绍" prop="description">
        <ElInput :rows="5" type="textarea" v-model="formData.description" placeholder="请输入个人介绍" />
      </ElFormItem>
      <ElFormItem label="状态" prop="status">
        <ElRadioGroup v-model="formData.status">
          <ElRadio :value="UserStatus.Disable">禁用</ElRadio>
          <ElRadio :value="UserStatus.Enable">启用</ElRadio>
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
import { fetchGetUserEdit, fetchPostUserCreate, fetchPostUserEdit } from '@/api/user';
import { UserStatus } from '@/enums/statusEnum';
import { SexType } from '@/enums/typeEnum';
import type { FormInstance, FormRules  } from 'element-plus'
import type { CascaderOption } from 'element-plus';
import { provinceAndCityData } from 'element-china-area-data';

const region = provinceAndCityData as CascaderOption[]
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
const formData = reactive<User.Params.Model>({
  id: 0,
  name: null,
  phone: null,
  password: null,
  address: [],
  birthday: 0,
  description: null,
  sex: SexType.Other,
  avatar: "",
  status: UserStatus.Enable
})

// 表单验证规则
const rules: FormRules = {
  name: [
    { required: true, message: '请输入用户名', trigger: 'blur' },
    { min: 2, max: 20, message: '长度在 2 到 20 个字符', trigger: 'blur' }
  ],
  phone: [
    { required: true, message: '请输入手机号', trigger: 'blur' },
    { pattern: /^1[3-9]\d{9}$/, message: '请输入正确的手机号格式', trigger: 'blur' }
  ],
}

const handleAvatar = (e:string) => {
  formData.avatar = e
}






/**
 * 初始化表单数据
 * 根据对话框类型（新增/编辑）填充表单
 */
const initFormData = async () => {
  // const roleRes = await fetchManageAllRole()
  // roleList.value = roleRes != null ? roleRes.list : []
  if (props.type === 'create') {
    Object.assign(formData, {
      id: 0,
      name: null,
      phone: null,
      password: null,
      address: [],
      birthday: 0,
      description: null,
      sex: SexType.Other,
      avatar: "",
      status: UserStatus.Enable
    })
  }
  if (props.type === 'edit') {
    const res = await fetchGetUserEdit({id:props.id!})
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
        await fetchPostUserCreate(formData)
      }

      if (drawerType.value === 'edit') {
        await fetchPostUserEdit(formData)
      }

      ElMessage.success(drawerType.value === 'create' ? '添加成功' : '更新成功')
      emit('submit')
      handleClose()
    }
  })
}
</script>
