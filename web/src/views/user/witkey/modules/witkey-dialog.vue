<template>
  <ElDialog
    v-model="visible"
    title="添加威客"
    width="30%"
    align-center
    @close="handleClose"
  >
    <ElForm v-loading="loading" ref="formRef" :model="formData" :rules="rules" label-width="80px">
      <ElFormItem label="威客昵称" prop="name">
        <ElInput  v-model="formData.name" placeholder="请输入威客昵称" />
      </ElFormItem>
      <ElFormItem label="所属用户" prop="userId">
          <ElSelect
              v-model="formData.userId"
              filterable
              remote
              reserve-keyword
              placeholder="请输入用户手机号"
              :remote-method="remoteMethod"
              :loading="loading"
        
          >
              <ElOption
                  v-for="item in userOptions"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value"
              />
          </ElSelect>
      </ElFormItem>
      <ElFormItem label="所属游戏" prop="gameId">
        <ElSelect
          clearable
          style="width: 100%;"
          v-model="formData.gameId"
          placeholder="请选择游戏"
          @change="handleChangeGame"
        >
          <el-option
            v-for="item in gameOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </ElSelect>
      </ElFormItem>
      <ElFormItem v-if="formData.gameId" label="所属头衔" prop="titleId">
        <ElSelect
          clearable
          style="width: 100%;"
          v-model="formData.titleId"
          placeholder="请选择头衔"
        >
          <el-option
            v-for="item in titleOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </ElSelect>
      </ElFormItem>
      <ElFormItem label="评分" prop="rate">
        <El-rate v-model="formData.rate" />
      </ElFormItem>
      <ElFormItem label="相册" prop="album">
        <ElSpace wrap>
          <div v-for="(item,index) in formData.album" 
              class="w-[100px] h-[100px] relative">
              <ElImage
                class="rounded-lg"
                :src="item"
                :preview-src-list="[item]"
                >
              </ElImage>
              <div @click="deleteAlbum(index)"  class="absolute top-0 right-0 cursor-pointer">
                <ArtSvgIcon icon="solar:close-circle-bold" class="text-red-500 text-2xl"/>
              </div>
          </div>
          <ArtUpload
          @success="handleAlbum"
          >
            <div class="w-[100px] h-[100px] flex items-center justify-center border-1 border-solid rounded-lg cursor-pointer">
              <ArtSvgIcon icon="solar:upload-square-bold" class="text-2xl"/>
            </div>
          </ArtUpload>
        </ElSpace>
      </ElFormItem>
    </ElForm>
    <template #footer>
      <div class="drawer-footer">
        <ElButton @click="handleClose">取消</ElButton>
        <ElButton type="primary" @click="handleSubmit">提交</ElButton>
      </div>
    </template>
  </ElDialog>
</template>

<script setup lang="ts">
import { fetchGetGameOptionsList, fetchGetTitleOptionsList, fetchGetUserOptionsList } from '@/api/site';
import { fetchPostWitkeyCreate } from '@/api/witkey';
import type { FormInstance, FormRules  } from 'element-plus'
import { useDebounce } from '@/hooks/index';

interface Props {
  modelValue: boolean
}
const props = withDefaults(defineProps<Props>(), {
  modelValue: false,
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

interface ListItem {
  value: number
  label: string
}

// 表单实例
const formRef = ref<FormInstance>()

// 表单数据
const formData = reactive<Witkey.Params.Model>({
  id: 0,
  name: null,
  titleId: null,
  gameId: null,
  userId: null,
  album: [],
  rate: 0,
})

// 表单验证规则
const rules: FormRules = {
  name: [
    { required: true, message: '请输入威客昵称', trigger: 'blur' },
  ],
  userId: [
    { required: true, message: '请选择所属用户', trigger: 'blur' },
  ],
  gameId: [
    { required: true, message: '请选择所属游戏', trigger: 'blur' },
  ],
  titleId: [
    { required: true, message: '请选择所属头衔', trigger: 'blur' },
  ]
}

const handleAlbum = (e:string) => {
  formData.album.push(e)
}

const deleteAlbum = (i:number) => {
  formData.album.splice(i,1)
}
const { debounce } = useDebounce()
const loading = ref(false)
const userOptions = ref<ListItem[]>([])
const getUserOptions =  async (phone:string) => {
  try {
      const res = await fetchGetUserOptionsList({phone})
      userOptions.value = res.list.map((item) => {
          return {
              value: item.id,
              label: item.name
          }
      })
  } finally {
      loading.value = false
  }
}

// 创建防抖版本
const getUserOptionsDebounced = debounce((query: string) => {
loading.value = true
getUserOptions(query)
}, 500)
const remoteMethod = (query: string) => {
  if (query) {
    getUserOptionsDebounced(query)
  } else {
    userOptions.value = []
  }
}


const gameOptions = ref<ListItem[]>([])
const getGameOptions =  async () => {
  const res = await fetchGetGameOptionsList()
  gameOptions.value = res.list.map((item) => {
    return {
      value: item.id,
      label: item.name
    }
  })
}

const handleChangeGame = async (v:any) =>{
  await getTitleOptions(v)
}
const titleOptions = ref<ListItem[]>([])
const getTitleOptions = async (id:number) => {
  const res = await fetchGetTitleOptionsList({gameId:id})
  titleOptions.value = res.list.map((item) => {
    return {
      value: item.id,
      label: item.name
    }
  })
}


/**
 * 初始化表单数据
 * 根据对话框类型（新增/编辑）填充表单
 */
const initFormData = async () => {
  loading.value = true 
  await getGameOptions()
  Object.assign(formData, {
    name: null,
    titleId: null,
    gameId: null,
    userId: null,
    album: [],
    rate: 0,
  })
  loading.value = false 
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
      await fetchPostWitkeyCreate(formData)
      ElMessage.success('添加成功')
      emit('submit')
      handleClose()
    }
  })
}
</script>
