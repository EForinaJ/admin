<template>
    <div class="w-[700px] p-[20px]">
        <ElForm ref="formRef" :model="formData" :rules="rules" label-width="120px">
            <ElFormItem label="默认头像" prop="avatar">
                <ArtUpload  class="w-full"
                    :class="formData.avatar != '' ? 'mb-[-11px]' : ''"
                    @success="handleAvatar"
                    >
                    <ElAvatar v-if="formData.avatar != ''"
                    shape="square"
                        :size="80" :src="formData.avatar" />
                    <ElButton v-if="formData.avatar == ''" type="primary">选择文件</ElButton>
                </ArtUpload>
                <ElText size="small">用户注册默认头像,尺寸建议60*180</ElText>
            </ElFormItem>
            <ElFormItem label="默认等级" prop="levelId">
                <ElInput v-model="formData.levelId" placeholder="请输入默认等级编号ID" />
                <ElText size="small">用户注册后默认等级编号ID</ElText>
            </ElFormItem>
            <ElFormItem label="签到经验" prop="checkInExperience">
                <ElInput v-model="formData.checkInExperience" placeholder="请输入经验值" />
                <ElText size="small">用户签到后所得的经验值(在值内随机例如输入99，随机获取的经验值在99以内)</ElText>
            </ElFormItem>
            <ElFormItem>
                <ElButton @click="handleSubmit" type="primary">更新保存</ElButton>
            </ElFormItem>
        </ElForm>
    </div>
</template>

<script setup lang="ts">
import { fetchGetSystemUser, fetchPostSystemUser } from '@/api/system';
import { FormInstance, FormRules } from 'element-plus';


// 表单实例
const formRef = ref<FormInstance>()
// 表单数据
const formData = reactive<System.Params.User>({
    avatar: "",
    levelId: null,
    checkInExperience: null,
})
// 表单验证规则
const rules: FormRules = {
    avatar: [
        { required: true, message: '请设置用户默认头像', trigger: 'blur' }
    ],
    levelId: [
        { required: true, message: '请设置用户默认等级', trigger: 'blur' }
    ],
}


const handleAvatar = (e:string)=>{
    formData.avatar = e
}



const handleSubmit = async () => {
    if (!formRef.value) return

    try {
        await formRef.value.validate()
        await fetchPostSystemUser(formData)
        // getData()
        ElMessage.success('更新成功')
    } catch (error) {
        console.log('表单验证失败:', error)
    }
  }

const getData = async () =>{
  const res = await fetchGetSystemUser()
  Object.assign(formData,res)
}

onMounted(()=>{
    getData()
})

</script>