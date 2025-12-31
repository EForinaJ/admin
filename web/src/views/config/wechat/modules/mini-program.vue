<template>
    <div class="w-[700px] p-[20px]">
        <ElForm ref="formRef" :model="formData" :rules="rules" label-width="120px">
            <ElFormItem label="小程序AppId" prop="appId">
                <ElInput v-model="formData.appId" placeholder="请输入微信小程序AppId" />
                <el-text size="small">请从微信小程序后台获取，编辑完成后请确保填写正确的小程序AppId</el-text>
            </ElFormItem>
            <ElFormItem label="小程序Secret" prop="secret">
                <ElInput v-model="formData.secret" placeholder="请输入微信小程序Secret" />
                <ElText size="small">请从微信小程序后台获取，编辑完成后请确保填写正确的小程序Secret</ElText>
            </ElFormItem>
            <ElFormItem>
                <ElButton @click="handleSubmit" type="primary">更新保存</ElButton>
            </ElFormItem>
        </ElForm>
    </div>
</template>

<script setup lang="ts">
import { fetchGetSystemWechatMiniProgram, fetchPostSystemWechatMiniProgram } from '@/api/system';
import { FormInstance, FormRules } from 'element-plus';


// 表单实例
const formRef = ref<FormInstance>()
// 表单数据
const formData = reactive<System.Params.WechatMiniProgram>({
    appId: null,
    secret: null,
})
// 表单验证规则
const rules: FormRules = {
    appId: [
        { required: true, message: '请输入微信小程序AppId', trigger: 'blur' }
    ],
    secret: [
        { required: true, message: '请输入微信小程序Secret', trigger: 'blur' }
    ],
}



const handleSubmit = async () => {
    if (!formRef.value) return

    try {
        await formRef.value.validate()
        await fetchPostSystemWechatMiniProgram(formData)
        // getData()
        ElMessage.success('更新成功')
    } catch (error) {
        console.log('表单验证失败:', error)
    }
  }

const getData = async () =>{
  const res = await fetchGetSystemWechatMiniProgram()
  Object.assign(formData,res)
}

onMounted(()=>{
    getData()
})
</script>