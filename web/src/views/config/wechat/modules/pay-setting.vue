<template>
    <div class="w-[700px] p-[20px]">
        <ElForm ref="formRef" :model="formData" :rules="rules" label-width="120px">
            <ElFormItem label="微信商户号" prop="mchId">
                <ElInput v-model="formData.mchId" placeholder="请输入微信商户号" />
                <el-text size="small">微信商户号</el-text>
            </ElFormItem>
            <ElFormItem label="商户API密钥" prop="mchKey">
                <ElInput v-model="formData.mchKey" placeholder="请输入微信商户API密钥" />
                <ElText size="small">请使用APIv3密钥，请确保填写正确的商户API密钥</ElText>
            </ElFormItem>
            <ElFormItem label="商户证书序列号" prop="serialNo">
                <ElInput v-model="formData.serialNo" placeholder="请输入微信商户API密钥" />
                <ElText size="small">请使用API证书序列号</ElText>
            </ElFormItem>
            <ElFormItem label="商户API证书" prop="apiclientCert">
                <ArtUpload  class="w-full"
                    @success="handleApiclientCert"
                    >
                    <div 
                    class="flex items-center gap-[10px]"
                    v-if="formData.apiclientCert != '' || formData.apiclientCert != null">
                        <ElButton  type="default">重新上传</ElButton>
                        <ElText size="small">{{ formData.apiclientCert }}</ElText>
                    </div>
                    <ElButton v-if="formData.apiclientCert == '' || formData.apiclientCert == null" type="primary">选择文件</ElButton>
                </ArtUpload>
                <ElText size="small">调用微信商户接口时需要用到的API证书，文件名称一般为apiclient_cert</ElText>
            </ElFormItem>
            <ElFormItem label="商户API证书密钥" prop="apiclientKey">
                <ArtUpload  class="w-full" @success="handleApiclientKey"
                >
                    <div 
                    class="flex items-center gap-[10px]"
                    v-if="formData.apiclientKey != '' || formData.apiclientKey != null">
                        <ElButton  type="default">重新上传</ElButton>
                        <ElText size="small">{{ formData.apiclientKey }}</ElText>
                    </div>
                    <ElButton v-if="formData.apiclientKey == '' || formData.apiclientKey == null" type="primary">选择文件</ElButton>
                </ArtUpload>
                <ElText size="small">调用微信商户接口时需要用到的API证书密钥，文件名称一般为apiclient_key</ElText>
            </ElFormItem>
            <ElFormItem label="微信支付公钥ID" prop="pubId">
                <ElInput v-model="formData.pubId" placeholder="请输入微信支付公钥ID" />
                <ElText size="small">请使用微信支付公钥ID</ElText>
            </ElFormItem>
            <ElFormItem label="微信支付公钥文件" prop="pubKey">
                <ArtUpload  class="w-full"
                  
                    @success="handlePubKey"
                >
                    <div 
                    class="flex items-center gap-[10px]"
                    v-if="formData.pubKey != '' || formData.pubKey != null">
                        <ElButton  type="default">重新上传</ElButton>
                        <ElText size="small">{{ formData.pubKey }}</ElText>
                    </div>
                    <ElButton v-if="formData.pubKey == '' || formData.pubKey == null" type="primary">选择文件</ElButton>
                </ArtUpload>
                <ElText size="small">调用微信商户接口时需要用到的公钥证书，文件名称一般为pub_key</ElText>
            </ElFormItem>

            <ElFormItem>
                <ElButton @click="handleSubmit" type="primary">更新保存</ElButton>
            </ElFormItem>
        </ElForm>
    </div>
</template>

<script setup lang="ts">
import { fetchGetSystemWechatPay, fetchPostSystemWechatPay } from '@/api/system';
import { FormInstance, FormRules } from 'element-plus';


// 表单实例
const formRef = ref<FormInstance>()
// 表单数据
const formData = reactive<System.Params.WechatPay>({
    mchId: null,
    mchKey: null,
    serialNo: null,
    apiclientCert: null,
    apiclientKey: null,
    pubKey: null,
    pubId: null,
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
        await fetchPostSystemWechatPay(formData)
        // getData()
        ElMessage.success('更新成功')
    } catch (error) {
        console.log('表单验证失败:', error)
    }
  }

const getData = async () =>{
  const res = await fetchGetSystemWechatPay()
  Object.assign(formData,res)
}

const handleApiclientCert = (e:string) =>{
    formData.apiclientCert = e
}
const handleApiclientKey = (e:string) =>{
    formData.apiclientKey = e
}
const handlePubKey = (e:string) =>{
    formData.pubKey = e
}
onMounted(()=>{
    getData()
})
</script>