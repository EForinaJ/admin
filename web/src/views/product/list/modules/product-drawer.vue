<template>
  <ElDrawer
    v-model="visible"
    :title="drawerType === 'create' ? '添加商品' : '编辑商品'"
    size="50%"
    @close="handleClose"
  >
    <ElForm v-loading="loading" ref="formRef" :model="formData" :rules="rules" label-width="80px">
      <ElFormItem label="商品名称" prop="name">
        <ElInput v-model="formData.name" placeholder="请输入商品名称" />
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
      <ElFormItem label="商品分类" prop="categoryId">
        <ElSelect
          clearable
          style="width: 100%;"
          v-model="formData.categoryId"
          placeholder="请选择商品分类"
        >
          <el-option
            v-for="item in categoryOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </ElSelect>
      </ElFormItem>
      <ElFormItem label="商品价格" prop="price">
        <ElInputNumber 
          :precision="2"
         style="width: 100%;"
          v-model="formData.price" placeholder="请输入商品价格"
            controls-position="right"> </ElInputNumber>
      </ElFormItem>
      <ElFormItem label="评分" prop="rate">
        <El-rate v-model="formData.rate" />
      </ElFormItem>
      <ElFormItem label="商品单位" prop="unit">
        <ElInput style="width: 100%;"
        v-model="formData.unit" placeholder="请输入商品单位" />
      </ElFormItem>
      <ElFormItem label="服务人数" prop="witkeyCount">
        <ElInputNumber 
          :precision="0"
          style="width: 100%;"
          v-model="formData.witkeyCount" placeholder="请输入服务人数"
          controls-position="right"/>
      </ElFormItem>
      <ElFormItem label="限购次数" prop="purchaseLimit">
        <ElInputNumber 
          :precision="0"
          style="width: 100%;"
          v-model="formData.purchaseLimit" placeholder="请输入限购次数"
          controls-position="right"/>
      </ElFormItem>
      <ElFormItem label="销售数量" prop="salesCount">
        <ElInputNumber 
          :precision="0"
          style="width: 100%;"
          v-model="formData.salesCount" placeholder="请输入销售数量"
          controls-position="right"> </ElInputNumber>
      </ElFormItem>
      <ElFormItem label="商品描述" prop="description">
        <ElInput  style="width: 100%;" :rows="6"
        type="textarea"
        v-model="formData.description" placeholder="请输入商品描述" />
      </ElFormItem>
      <ElFormItem label="封面" prop="pic">
        <ArtUpload
          @success="handlePic"
        >
          <ElImage
          class="w-[80px]"
          shape="square"
          v-if="formData.pic != ''" :src="formData.pic!" />
          <ElButton v-if="formData.pic == ''" type="primary">选择文件</ElButton>
        </ArtUpload>
      </ElFormItem>
      <ElFormItem label="状态" prop="status">
        <ElRadioGroup v-model="formData.status">
          <ElRadio :value="ProductStatus.Disable">下架</ElRadio>
          <ElRadio :value="ProductStatus.Enable">上架</ElRadio>
        </ElRadioGroup>
      </ElFormItem>
      <ElDivider content-position="left">商品详情</ElDivider>
      <ElFormItem  prop="content">
          <ArtWangEditor 
              v-model="formData.content"
              :uploadConfig="{
                server:'/api/v1/admin/upload/file',
              }"
          />
      </ElFormItem>



      <!-- <ElFormItem label="商品规格" prop="specification">
        <ElCard style="width: 100%;" shadow="never">
          <div v-for="(item,index) in formData.specification" class="mb-[20px]">
            <div class="mb-[10px] flex items-center bg-[#F0F2F5] p-[10px] rounded-lg">
              <div class="flex-1">
                  <ElSpace align="center">
                    <span>规格名称:</span>
                    <ElInput  
                     @input="updateSpecificationName(index,$event)"
                    style="width: 100%;" 
                    :model-value="item.name" placeholder="请输入规格名称" />
                    <ElButton type="danger" @click="deleteSpecification(index)">删除</ElButton>
                  </ElSpace>
              </div>
            </div>
            <div class="bg-white rounded-lg">
              <ElRow :gutter="20">
                <ElCol class="mb-[10px]" v-for="(vitem,vindex) in item.values" :span="12">
                  <ElSpace>
                    <ElInput 
                      :model-value="vitem"
                      @input="updateSpecificationValue(index,vindex,$event)"
                      placeholder="请输入规格名称" />
                      <ElButton type="danger" @click="deleteSpecificationValue(index,vindex)">
                        删除
                      </ElButton>
                  </ElSpace>
                </ElCol>
                <ElCol class="mb-[10px]" :span="8">
                  <ElButton type="success" @click="addSpecificationValue(index)">添加规格值</ElButton>
                </ElCol>
              </ElRow>
            </div>
          </div>
          <template #footer>
            <ElSpace>
              <ElButton type="success" @click="addSpecification">添加规格</ElButton>
              <ElButton v-if="formData.specification.length > 0" type="primary" @click="generateSpecCombinations">生成规格</ElButton>
            </ElSpace>
          </template>
        </ElCard>
      </ElFormItem>
      <ElFormItem v-if="formData.sku.length > 0" label="规格明细" prop="sku">
        <ElTable :data="skuTableData" style="width: 100%">
          <ElTableColumn 
          v-for="item in formData.specification"
          :prop="item.name" :label="item.name"  />
          <ElTableColumn 
          prop="price" fixed="right"
          label="售价">
            <template #default="scope">
              <ElInputNumber 
                :precision="2"
                style="width: 100%;"
                v-model="scope.row.price" placeholder="请输入商品价格"
                controls-position="right"> </ElInputNumber>
              </template>
          </ElTableColumn>
          <ElTableColumn fixed="right"
          prop="originalPrice" 
          label="原价">
            <template #default="scope">
              <ElInputNumber 
                :precision="2"
                style="width: 100%;"
                v-model="scope.row.originalPrice" placeholder="请输入商品原价"
                controls-position="right"> </ElInputNumber>
            </template>
          </ElTableColumn>
          <ElTableColumn fixed="right"
            prop="stock" 
            label="库存">
            <template #default="scope">
              <ElInputNumber 
                :precision="2"
                style="width: 100%;"
                v-model="scope.row.stock" placeholder="请输入商品库存"
                controls-position="right"> </ElInputNumber>
            </template>
          </ElTableColumn>
        </ElTable>
      </ElFormItem> -->

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
import { fetchGetProductEdit, fetchPostProductCreate, fetchPostProductEdit } from '@/api/product';
import { fetchGetGameOptionsList, fetchGetCategoryOptionsList } from '@/api/site';
import { ProductStatus } from '@/enums/statusEnum';
import { SpecCartesianProduct } from '@/utils/form/sku';
import type { FormInstance, FormRules  } from 'element-plus'

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
const formData = reactive<Product.Params.Model>({
  id: 0,
  name: "",
  gameId: null,
  price: null,
  rate:  0,
  witkeyCount:  null,
  purchaseLimit:  null,
  salesCount:  null,
  categoryId: null,
  unit: null,
  pic: "",
  description: "",
  content: "",
  status: ProductStatus.Enable,
  // skuType: 0,
  // specification:[],
  // sku:[],
})

// 表单验证规则
const rules: FormRules = {
  name: [
    { required: true, message: '请输入商品名', trigger: 'blur' },
  ],
  gameId: [
    { required: true, message: '请选择游戏领域', trigger: 'blur' },
  ],
  categoryId: [
    { required: true, message: '请选择商品分类', trigger: 'blur' },
  ],
  price: [
    { required: true, message: '请输入价格', trigger: 'blur' },
  ],
  number: [
    { required: true, message: '请输入服务人数', trigger: 'blur' },
  ],
  unit: [
    { required: true, message: '请输入商品单位', trigger: 'blur' },
  ],
  pic: [
    { required: true, message: '请上传商品封面', trigger: 'blur' },
  ],
}

const handlePic = (e:string) => {
  formData.pic = e
}
// const addSpecification = () => {
//   formData.specification.push({
//     name: "",
//     values: [""]
//   })
// }

// const deleteSpecification = (i:number) => {
//   formData.specification.splice(i,1)
//   if (formData.specification.length == 0) {
//     formData.sku = []
//   }
// }

// const updateSpecificationName = (i:number,value:string) =>{
//   formData.specification[i].name = value
// }

// const addSpecificationValue = (i:number) => {
//   formData.specification[i].values.push("")
// }

// const deleteSpecificationValue = (i:number,v:number) => {
//   formData.specification[i].values.splice(v,1)
// }

// const updateSpecificationValue = (i:number,v:number,value:string) =>{
//   formData.specification[i].values[v] = value
// }

// const generateSpecCombinations =  ()=>{
//   const valueCombinations = SpecCartesianProduct.generateSpecCombinations(formData.specification)
//   formData.sku = valueCombinations
// }

// const skuTableData = computed(()=>{
//   return formData.sku.map((item) => {
//     const specification: Record<string, string> = {};
//     item.attrs.forEach((item) => {
//       specification[item.attr] = item.value;
//     });
//     const tmp = {
//       ...specification,
//       price:item.price,
//       originalPrice:item.originalPrice,
//       stock:item.stock,
//     }
//     return tmp
//   })
// })


interface ListItem {
  value: number
  label: string,
}

const loading = ref(false)
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
  await getCategoryOptions(v)
}
const categoryOptions = ref<ListItem[]>([])
const getCategoryOptions =  async (id:number) => {
  const res = await fetchGetCategoryOptionsList({gameId:id})
  categoryOptions.value = res.list.map((item) => {
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
  if (props.type === 'create') {
    Object.assign(formData, {
      id: 0,
      name: "",
      gameId: null,
      price: 0,
      rate:  0,
      number:  0,
      salesCount:  0,
      categoryId: null,
      unit: null,
      pic: "",
      description: "",
      content: "",
      status: ProductStatus.Enable,
    })
  }
  if (props.type === 'edit') {
    const res = await fetchGetProductEdit({id:props.id!})
    await getCategoryOptions(res.gameId!)
    Object.assign(formData, res)
  }
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

      if (drawerType.value === 'create') {
        await fetchPostProductCreate(formData)
      }

      if (drawerType.value === 'edit') {
        await fetchPostProductEdit(formData)
      }

      ElMessage.success(drawerType.value === 'create' ? '添加成功' : '更新成功')
      emit('submit')
      handleClose()
    }
  })
}
</script>
