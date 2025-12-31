<template>
    <div class="p-[20px]">
        <ElForm ref="formRef" :model="formData" :rules="rules">
            <ElFormItem  prop="symbol">
               <ArtWangEditor 
                    v-model="formData.content"
               />
            </ElFormItem>

            <ElFormItem>
                <ElButton @click="handleSubmit" type="primary">更新保存</ElButton>
            </ElFormItem>
        </ElForm>
    </div>
</template>

<script setup lang="ts">
import { fetchGetSystemAboutUs, fetchPostSystemAboutUs,  } from '@/api/system';
import { FormInstance, FormRules } from 'element-plus';


// 表单实例
const formRef = ref<FormInstance>()
// 表单数据
const formData = reactive({
    content: "",
})
// 表单验证规则
const rules: FormRules = {
    content: [
        { required: true, message: '请输入内容', trigger: 'blur' }
    ],
}


const handleSubmit = async () => {
    if (!formRef.value) return

    try {
        await formRef.value.validate()
        await fetchPostSystemAboutUs(formData)
        getData()
        ElMessage.success('更新成功')
    } catch (error) {
        console.log('表单验证失败:', error)
    }
  }

const getData = async () =>{
  const res = await fetchGetSystemAboutUs()
  Object.assign(formData,res)
}

onMounted(()=>{
    getData()
})
</script>