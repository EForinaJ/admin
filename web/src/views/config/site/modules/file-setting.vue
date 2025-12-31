<template>
    <div class="w-[700px] p-[20px]">
        <ElForm ref="formRef" :model="formData" :rules="rules" label-width="120px">
            <ElFormItem label="默认驱动" prop="symbol">
                <ElSelect v-model="formData.engine" placeholder="请选择存储驱动" style="width: 240px">
                    <ElOption
                        label="本地"
                        value="local"
                    />
                </ElSelect>
            </ElFormItem>
            <ElFormItem label="存储路径" prop="path">
                <ElInput v-model="formData.path" placeholder="请输入存储路径" />
                <ElText size="small">填对外访问的相对路径</ElText>
            </ElFormItem>
            <ElFormItem label="图片大小限制" prop="imageSize">
                <ElInput v-model="formData.imageSize" placeholder="请输入图片最大上传体积" />
                <ElText size="small">图片上传时最大限制，超出不可上传体积</ElText>
            </ElFormItem>
            <ElFormItem label="图片类型限制" prop="imageType">
                <ElInput v-model="formData.imageType" placeholder="请输入图片类型" />
                <ElText size="small">图片格式上传限制（.png,.jpg等用英文逗号隔开）</ElText>
            </ElFormItem>
            <ElFormItem label="文件大小限制" prop="fileSize">
                <ElInput v-model="formData.fileSize" placeholder="请输入文件最大上传体积" />
                <ElText size="small">文件上传时最大限制，超出不可上传体积</ElText>
            </ElFormItem>
            <ElFormItem label="文件类型限制" prop="fileType">
                <ElInput v-model="formData.fileType" placeholder="请输入文件类型" />
                <ElText size="small">文件格式上传限制（.mp3,.mp4等用英文逗号隔开）</ElText>
            </ElFormItem>
            <ElFormItem>
                <ElButton @click="handleSubmit" type="primary">更新保存</ElButton>
            </ElFormItem>
        </ElForm>
    </div>
</template>

<script setup lang="ts">
import { fetchGetSystemFile, fetchPostSystemFile } from '@/api/system';
import { FormInstance, FormRules } from 'element-plus';


// 表单实例
const formRef = ref<FormInstance>()
// 表单数据
const formData = reactive<System.Params.File>({
    engine: null,
    path: null,
    fileSize: null,
    fileType: null,
    imageSize: null,
    imageType: null,
})
// 表单验证规则
const rules: FormRules = {
    engine: [
        { required: true, message: '请选择存储驱动', trigger: 'blur' }
    ],
    path: [
        { required: true, message: '请输入存储路径', trigger: 'blur' }
    ],
    imageSize: [
        { required: true, message: '请输入图片最大上传体积', trigger: 'blur' }
    ],
    imageType: [
        { required: true, message: '请输入图片类型', trigger: 'blur' }
    ],
    fileSize: [
        { required: true, message: '请输入文件最大上传体积', trigger: 'blur' }
    ],
    fileType: [
        { required: true, message: '请输入文件类型', trigger: 'blur' }
    ],
}


const handleSubmit = async () => {
    if (!formRef.value) return

    try {
        await formRef.value.validate()
        formData.fileType =  formData.fileType?.split(',')
        formData.imageType =  formData.imageType?.split(',')
        await fetchPostSystemFile(formData)
        // getData()
        ElMessage.success('更新成功')
    } catch (error) {
        console.log('表单验证失败:', error)
    }
  }

const getData = async () =>{
  const res = await fetchGetSystemFile()
  Object.assign(formData,res)
  formData.fileType = res.fileType.join(',')
  formData.imageType = res.imageType.join(',')
}

onMounted(()=>{
    getData()
})
</script>