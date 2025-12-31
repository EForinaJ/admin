<template>
    <div>
      <ElUpload
        ref="uploadRef"
        action="#" 
        :auto-upload="true" 
        :multiple="false"
        :on-exceed="handleExceed"
        :before-upload="beforeUpload"
        :http-request="customRequest"
        :show-file-list="false"
      >
        <slot></slot>
      </ElUpload>
    </div>
  </template>
  
<script setup lang="ts">
import { ref } from 'vue';
import { ElMessage } from 'element-plus';
import { fetchUpload } from '@/api/upload';
import { useSiteStore } from '@/store/modules/site';

interface Emits {
  (e: 'success', link: string): void
}
const emit = defineEmits<Emits>()
const {getInfo:site} = useSiteStore()
const uploadRef = ref();

// 1. 文件数量超限提示
const handleExceed = () => {
  ElMessage.warning('最多只能上传3个文件');
};

// 2. 上传前校验（大小、类型）
const beforeUpload = (file:any) => {
  // const isLt2M = file.size / 1024 / 1024 < site.minFileSize!;
  // if (!isLt2M) {
  //   ElMessage.error(`单个文件大小不能超过 ${site.minFileSize}MB!`);
  //   return false; // 阻止上传
  // }
  return true;
};



// 4. 核心：自定义上传请求
const customRequest = async (options: { file: any; onProgress: any; onSuccess: any; onError: any; }) => {
  const { file, onProgress, onSuccess, onError } = options;
  console.log(file)
  const formData = new FormData();
  formData.append('file', file); // 字段名需与后端约定
  // 可以追加其他参数

  try {
    const res = await fetchUpload(formData)
    // const response = await axios.post('https://your-api.com/upload', formData, {
    //   headers: {
    //     'Content-Type': 'multipart/form-data',
    //     // 其他自定义Header，如token
    //     // 'Authorization': `Bearer ${token}`
    //   },
    //   // 上传进度支持
    //   onUploadProgress: (progressEvent) => {
    //     const percent = Math.round((progressEvent.loaded * 100) / progressEvent.total);
    //     onProgress({ percent }); // 调用钩子更新组件内部进度显示
    //   }
    // });

    // 请求成功，调用onSuccess通知组件，参数会传递给on-success回调
    emit("success",res.links[0])
    onSuccess(res.links);
    ElMessage.success(`${file.name} 上传成功！`);
  } catch (err) {
    // 请求失败，调用onError通知组件，参数会传递给on-error回调
    onError(err);
    ElMessage.error(`${file.name} 上传失败！`);
  }
};
  
</script>