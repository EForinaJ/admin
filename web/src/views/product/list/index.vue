<!-- 商品管理页面 -->
<!-- art-full-height 自动计算出页面剩余高度 -->
<!-- art-table-card 一个符合系统样式的 class，同时自动撑满剩余高度 -->
<!-- 更多 useTable 使用示例请移步至 功能示例 下面的高级表格示例或者查看官方文档 -->
<!-- useTable 文档：https://www.artd.pro/docs/zh/guide/hooks/use-table.html -->
<template>
  <div class="product-page art-full-height">
    <!-- 搜索栏 -->
    <ProductSearch v-model="searchForm" @search="handleSearch" @reset="resetSearchParams"></ProductSearch>

    <ElCard class="art-table-card" shadow="never">
      <!-- 表格头部 -->
      <ArtTableHeader v-model:columns="columnChecks" :loading="loading" @refresh="refreshData">
        <template #left>
          <ElSpace wrap>
            <ElButton v-auth="'create'" @click="handleCreate" v-ripple>新增商品</ElButton>
            <ElButton 
            :disabled="selectedRows.length == 0"
            v-auth="'delete'" 
            @click="handleBatchDelete" type="danger" v-ripple>批量删除</ElButton>
          </ElSpace>
        </template>
      </ArtTableHeader>

      <!-- 表格 -->
      <ArtTable
        :loading="loading"
        :data="data"
        :columns="columns"
        :pagination="pagination"
        @selection-change="handleSelectionChange"
        @pagination:size-change="handleSizeChange"
        @pagination:current-change="handleCurrentChange"
      >
      </ArtTable>

      <!-- 商品弹窗 -->
      <ProductDrawer
        v-model="drawerVisible"
        :type="drawerType"
        :id="id"
        @submit="refreshData"
      />
    </ElCard>
  </div>
</template>

<script setup lang="ts">
import { useTable } from '@/hooks/core/useTable'
// import ProductSearch from './modules/product-search.vue'
import { ElTag, ElMessageBox, ElImage, ElRate } from 'element-plus'
import { fetchGetProductList, fetchPostProductDelete } from '@/api/product'
import { useSiteStore } from '@/store/modules/site'
import { useAuth } from '@/hooks'
import ArtButtonTable from '@/components/core/forms/art-button-table/index.vue'
import ProductDrawer from './modules/product-drawer.vue'
import ProductSearch from './modules/product-search.vue'


const { hasAuth } = useAuth();
defineOptions({ name: 'Product' })

const siteStore = useSiteStore()



// 弹窗相关
const drawerType = ref<"create" | "edit">('create')
const drawerVisible = ref(false)
const viewDrawerVisible = ref(false)
const id = ref<number>(0)

// 选中行
const selectedRows = ref<number[]>([])

// 搜索表单
const searchForm = ref({
  name: undefined,
  gameId: undefined,
  categoryId: undefined,
  status: undefined,
})

// 商品状态配置
const PRODUCT_STATUS_CONFIG = {
  1: { type: 'info' as const, text: '下架' },
  2: { type: 'success' as const, text: '在售' },
} as const

/**
 * 获取商品状态配置
 */
const getProductStatusConfig = (status: number) => {
  return (
    PRODUCT_STATUS_CONFIG[status as keyof typeof PRODUCT_STATUS_CONFIG] || {
      type: 'info' as const,
      text: '未知'
    }
  )
}

const {
  columns,
  columnChecks,
  data,
  loading,
  pagination,
  getData,
  searchParams,
  resetSearchParams,
  handleSizeChange,
  handleCurrentChange,
  refreshData
} = useTable({
  // 核心配置
  core: {
    apiFn: fetchGetProductList,
    apiParams:{
      name: "",
      status: 0,
    },
    paginationKey:{
      current: 'page', 
      size: 'limit'
    },
    columnsFactory: () => [
      { type: 'selection' }, // 勾选列
      { prop: 'id', width: 60, label: 'ID' }, // 序号
      { 
        prop: 'code', 
        label: '商品编码',
      },
      {
        prop: 'productInfo',
        label: '商品信息',
        width: 280,
        formatter: (row) => {
          return h('div', { class: 'product flex-c' }, [
            h(ElImage, {
              class: 'size-9.5 rounded-md',
              src: row.pic,
              previewSrcList: [row.pic],
              // 图片预览是否插入至 body 元素上，用于解决表格内部图片预览样式异常
              previewTeleported: true
            }),
            h('div', { class: 'ml-2' }, [
              h('p', { class: 'product-name' }, row.name),
              h('p', { class: 'email' }, row.category),
            ])
          ])
        }
      },
      { 
        prop: 'game', 
        label: '游戏领域',
        width: 130,
        formatter: (row) => {
          return h(ElTag, { type: "primary" }, () => `${row.game}`)
        }
      },
      { 
        prop: 'rate', 
        label: '评分',
        formatter: (row) => {
          return h(ElRate , {
            modelValue: row.rate,
            showScore: true, 
            disabled: true,
            scoreTemplate:`${row.rate}分` 
          })
        }
      },
      { 
        prop: 'price', 
        label: '商品售价',
        width: 140,
        formatter: (row) => {
          return h(ElTag, { type: "primary" }, () => `${row.price}${siteStore.getInfo.symbol}`)
        } 
      },
      {
        prop: 'status',
        label: '状态',
        width: 100,
        formatter: (row) => {
          const statusConfig = getProductStatusConfig(row.status)
          return h(ElTag, { type: statusConfig.type }, () => statusConfig.text)
        }
      },
      {
        prop: 'createTime',
        label: '创建日期',
        sortable: true
      },
      {
        prop: 'operation',
        label: '操作',
        width: 120,
        fixed: 'right', // 固定列
        formatter: (row) =>{
          return h('div', { class: 'product flex-c' }, [
            (hasAuth("edit") && h(ArtButtonTable, {
              type: 'edit',
              onClick: () => handleEdit(row)
            })),
            (hasAuth("delete") && h(ArtButtonTable, {
              type: 'delete',
              onClick: () => handleDelete(row)
            })),
          ])
        }
      }
    ],
  },
  // 数据处理
  transform: {
    responseAdapter: (response) => {
      return {
        records: response.list,
        total: response.total,
      };
    },
  },
})


/**
 * 搜索处理
 * @param params 参数
 */
const handleSearch = (params: Record<string, any>) => {
  console.log(params)
  // 搜索参数赋值
  Object.assign(searchParams, params)
  getData()
}

/**
 * 显示商品弹窗
 */
const handleCreate = (): void => {
  drawerType.value = "create"
  id.value = 0
  nextTick(() => {
    drawerVisible.value = true
  })
}
const handleEdit = (row:Product.Response.Info) => {
   drawerType.value = "edit"
   id.value = row.id
   nextTick(() => {
    drawerVisible.value = true
  })
}
const handleView = (row:Product.Response.Info) => {
    id.value = row.id
    nextTick(() => {
      viewDrawerVisible.value = true
    })
}

const handleBatchDelete = () =>{
  if (selectedRows.value.length != 0) {
    ElMessageBox.confirm(`确定要删除该吗？`, '删除商品', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'error'
    }).then(async() => {
      // TODO: 调用删除接口
      ElMessage.success('删除成功')
      await fetchPostProductDelete({ids:selectedRows.value})
      refreshData()
    })
    .catch(() => {
      ElMessage.info('已取消删除')
    })
  }
}


const handleDelete = async (row: Product.Response.Info): Promise<void> => {
  ElMessageBox.confirm(`确定要删除该吗？`, '删除商品', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'error'
  }).then(async() => {
    // TODO: 调用删除接口
    ElMessage.success('删除成功')
    await fetchPostProductDelete({ids:[row.id]})
    refreshData()
  })
  .catch(() => {
    ElMessage.info('已取消删除')
  })
}


/**
 * 处理表格行选择变化
 */
const handleSelectionChange = (selection: Product.Response.Info[]): void => {
  selectedRows.value = selection.map((item)=>item.id)
}
</script>
