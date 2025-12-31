<template>
    <div class="w-[700px] p-[20px]">
        <ElForm ref="formRef" :model="formData" :rules="rules" label-width="120px">
            <ElFormItem label="站点名称" prop="title">
                <ElInput v-model="formData.title" placeholder="请输入站点名称" />
                <el-text size="small">站点名称很多地方会显示的，建议认真填写</el-text>
            </ElFormItem>
            <ElFormItem label="站点网址" prop="domain">
                <ElInput v-model="formData.domain" placeholder="请输入站点网址" />
                <ElText size="small">站点网址很多地方会使用，请不要随便更换</ElText>
            </ElFormItem>
            <ElFormItem label="站点Logo" prop="logo">
                <ArtUpload  class="w-full"
                    :class="formData.logo != '' ? 'mb-[-11px]' : ''"
                    @success="handleLogo"
                    >
                    <ElAvatar v-if="formData.logo != ''"
                    shape="square"
                        :size="80" :src="formData.logo" />
                    <ElButton v-if="formData.logo == ''" type="primary">选择文件</ElButton>
                </ArtUpload>
                <ElText size="small">站点LOGO,尺寸建议60*180</ElText>
            </ElFormItem>
            <ElFormItem label="站点图标" prop="icon">
                <ArtUpload  class="w-full"
                    :class="formData.icon != '' ? 'mb-[-11px]' : ''"
                    @success="handleIcon"
                    >
                    <ElAvatar v-if="formData.icon != ''"
                    shape="square"
                        :size="80" :src="formData.icon" />
                    <ElButton v-if="formData.icon == ''" type="primary">选择文件</ElButton>
                </ArtUpload>
                <ElText size="small">站点icon,尺寸建议20*20</ElText>
            </ElFormItem>
            <ElFormItem label="备案号" prop="icp">
                <ElInput v-model="formData.icp" placeholder="请输入备案号" />
            </ElFormItem>
            <ElFormItem label="站点描述" prop="description">
                <ElInput 
                    type="textarea"
                    :autosize="{ minRows: 5, maxRows: 5 }"
                    v-model="formData.description" placeholder="站点描述" />
            </ElFormItem>
            <ElFormItem>
                <ElButton @click="handleSubmit" type="primary">更新保存</ElButton>
            </ElFormItem>
        </ElForm>
    </div>
</template>

<script setup lang="ts">
import { fetchGetSystemBase, fetchPostSystemBase } from '@/api/system';
import { FormInstance, FormRules } from 'element-plus';


// 表单实例
const formRef = ref<FormInstance>()
// 表单数据
const formData = reactive<System.Params.Base>({
    title: "",
    domain: "",
    logo: "",
    icon: "",
    description: "",
    icp: "",
})
// 表单验证规则
const rules: FormRules = {
    title: [
        { required: true, message: '请输入站点标题', trigger: 'blur' }
    ],
    domain: [
        { required: true, message: '请输入站点标题', trigger: 'blur' }
    ],
    logo: [
        { required: true, message: '请输入站点标题', trigger: 'blur' }
    ],
    icon: [
        { required: true, message: '请输入站点标题', trigger: 'blur' }
    ],
}

const handleLogo = (e:string)=>{
    formData.logo = e
}

const handleIcon = (e:string)=>{
    formData.icon = e
}

const handleSubmit = async () => {
    if (!formRef.value) return

    try {
        await formRef.value.validate()
        await fetchPostSystemBase(formData)
        // getData()
        ElMessage.success('更新成功')
    } catch (error) {
        console.log('表单验证失败:', error)
    }
  }

const getData = async () =>{
  const res = await fetchGetSystemBase()
  Object.assign(formData,res)
}

onMounted(()=>{
    getData()
})
</script>