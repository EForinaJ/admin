<template>
    <div class="w-[700px] p-[20px]">
        <ElForm ref="formRef" :model="formData" :rules="rules" label-width="120px">
            <ElFormItem label="货币符号" prop="symbol">
                <ElInput v-model="formData.symbol" placeholder="请输入货币符号" />
                <ElText size="small">站点显示金额的货币符合</ElText>
            </ElFormItem>
            <ElFormItem label="最小提现额度" prop="minWithdraw">
                <ElInput v-model="formData.minWithdraw" placeholder="请输入最小提现额度" />
                <ElText size="small">用户提现最低金额限制</ElText>
            </ElFormItem>
            <ElFormItem label="提现服务费率" prop="withDrawPercent">
                <ElInput v-model="formData.withDrawPercent" placeholder="请输入提现服务费率" />
                <ElText size="small">提现金额百分比（0~100）%，例如：提现100，手续费20%，实际到账100 * （1-20/100）</ElText>
            </ElFormItem>
            <ElFormItem>
                <ElButton @click="handleSubmit" type="primary">更新保存</ElButton>
            </ElFormItem>
        </ElForm>
    </div>
</template>

<script setup lang="ts">
import { fetchGetSystemWithdraw, fetchPostSystemWithdraw } from '@/api/system';
import { FormInstance, FormRules } from 'element-plus';


// 表单实例
const formRef = ref<FormInstance>()
// 表单数据
const formData = reactive<System.Params.Withdraw>({
    symbol: null,
    minWithdraw: null,
    withDrawPercent: null,
})
// 表单验证规则
const rules: FormRules = {
    symbol: [
        { required: true, message: '请输入货币符号', trigger: 'blur' }
    ],
    minWithdraw: [
        { required: true, message: '请输入最小提现额度', trigger: 'blur' }
    ],
    withDrawPercent: [
        { required: true, message: '请输入提现服务费率', trigger: 'blur' }
    ],
}


const handleSubmit = async () => {
    if (!formRef.value) return

    try {
        await formRef.value.validate()
        await fetchPostSystemWithdraw(formData)
        // getData()
        ElMessage.success('更新成功')
    } catch (error) {
        console.log('表单验证失败:', error)
    }
  }

const getData = async () =>{
  const res = await fetchGetSystemWithdraw()
  Object.assign(formData,res)
}

onMounted(()=>{
    getData()
})
</script>