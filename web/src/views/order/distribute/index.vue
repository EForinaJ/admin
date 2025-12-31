<!-- 派单管理页面 -->
<!-- art-full-height 自动计算出页面剩余高度 -->
<!-- art-table-card 一个符合系统样式的 class，同时自动撑满剩余高度 -->
<!-- 更多 useTable 使用示例请移步至 功能示例 下面的高级表格示例或者查看官方文档 -->
<!-- useTable 文档：https://www.artd.pro/docs/zh/guide/hooks/use-table.html -->
<template>
  <div class="distribute-page art-full-height">
    <!-- 搜索栏 -->
    <DistributeSearch v-model="searchForm" @search="handleSearch" @reset="resetSearchParams"></DistributeSearch>

    <ElCard class="art-table-card" shadow="never">
      <!-- 表格头部 -->
      <ArtTableHeader v-model:columns="columnChecks" :loading="loading" @refresh="refreshData">
        <template #left>
          <ElSpace wrap>
            <ElButton v-auth="'create'" @click="handleCreate" v-ripple>新增派单</ElButton>
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


      <DistributeCancelModal
        v-model:visible="cancelModalVisible"
        :id="id"
        @submit="refreshData"
      />

      <DistributeModal
        v-model:visible="distributeModalVisible"
        :id="id"
        @submit="refreshData"
      />
    </ElCard>
  </div>
</template>

<script setup lang="ts">
import { useTable } from '@/hooks/core/useTable'
import { ElTag, ElMessageBox } from 'element-plus'
import { useAuth } from '@/hooks'
import ArtButtonTable from '@/components/core/forms/art-button-table/index.vue'
import DistributeModal from './modules/distribute-modal.vue'
import { fetchGetDistributeList } from '@/api/distribute'
import DistributeCancelModal from './modules/distribute-cancel-modal.vue'
import DistributeSearch from './modules/distribute-search.vue'

const { hasAuth } = useAuth();
defineOptions({ name: 'Distribute' })


// 弹窗相关
const cancelModalVisible = ref(false)
const distributeModalVisible = ref(false)
const id = ref<number>(0)
// 选中行
const selectedRows = ref<number[]>([])

// 搜索表单
const searchForm = ref({
  name: undefined,
  phone: undefined,
  status: undefined
})


const IS_CANCEL = {
  1: { type: 'primary' as const, text: '未取消' },
  2: { type: 'danger' as const, text: '取消' },
} as const

const getIsCancel = (isCancel: number) => {
  return (
    IS_CANCEL[isCancel as keyof typeof IS_CANCEL] || {
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
    apiFn: fetchGetDistributeList,
    apiParams:{
      code: "",
      name: "",
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
        label: '订单号',
        width: 200,
      },
      {
        prop: 'manage',
        label: '派单者',
        width: 160,
        formatter: (row) => {
          return h('p', { }, row.manage)
        }
      },
      {
        prop: 'witkey',
        label: '接单者',
        width: 160,
        formatter: (row) => {
          return h('p', { }, row.witkey)
        }
      },
      {
        prop: 'game',
        label: '游戏领域',
        formatter: (row) => {
            return h(ElTag, { type:"primary" }, () => row.game )
        }
      },
      {
        prop: 'title',
        label: '所属头衔',
        formatter: (row) => {
            return h(ElTag, { type:"primary" }, () => row.title )
        }
      },
      {
        prop: 'isCancel',
        label: '是否取消',
        formatter: (row) => {
          const statusConfig = getIsCancel(row.isCancel)
          return h(ElTag, { type: statusConfig.type }, () => statusConfig.text)
        }
      },
      {
        prop: 'createTime',
        label: '派单时间',
        sortable: true
      },
      {
        prop: 'operation',
        label: '操作',
        width: 120,
        fixed: 'right', // 固定列
        formatter: (row) =>{
          return h('div', { class: 'distribute flex-c' }, [
            ((hasAuth("view") && row.isCancel == 2) && h(ArtButtonTable, {
              type: 'view',
              onClick: () => handleView(row)
            })),
            ((hasAuth("cancel") && row.isCancel == 1) && h(ArtButtonTable, {
              icon: 'solar:close-circle-bold',
              type: 'delete',
              onClick: () => handleCancel(row)
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

const handleView = (row:Distribute.Response.Info) => {
  if (row.isCancel != 2) {
    ElMessage.error('派单未取消')
    return
  }
  ElMessageBox.confirm(row.reason, '取消原因', {
    // confirmButtonText: '确定',
    // cancelButtonText: '取消',
    showCancelButton:false,
    showConfirmButton:false,
    type: 'info'
  }).then(async() => {
  })
  .catch(() => {
  })
}

const handleCancel = (row:Distribute.Response.Info) => {
   id.value = row.id
  nextTick(() => {
    cancelModalVisible.value = true
  })
}


/**
 * 显示派单弹窗
 */
const handleCreate = (): void => {
  nextTick(() => {
    distributeModalVisible.value = true
  })
}

const handleBatchDelete = () =>{
  if (selectedRows.value.length != 0) {
    ElMessageBox.confirm(`确定要删除该吗？`, '删除派单', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'error'
    }).then(async() => {
      // TODO: 调用删除接口
      ElMessage.success('删除成功')
      // await fetchPostDistributeDelete({ids:selectedRows.value})
      refreshData()
    })
    .catch(() => {
      ElMessage.info('已取消删除')
    })
  }
}


const handleDelete = async (row: Distribute.Response.Info): Promise<void> => {
  ElMessageBox.confirm(`确定要删除该吗？`, '删除派单', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'error'
  }).then(async() => {
    // TODO: 调用删除接口
    ElMessage.success('删除成功')
    // await fetchPostDistributeDelete({ids:[row.id]})
    refreshData()
  })
  .catch(() => {
    ElMessage.info('已取消删除')
  })
}


/**
 * 处理表格行选择变化
 */
const handleSelectionChange = (selection: Distribute.Response.Info[]): void => {
  selectedRows.value = selection.map((item)=>item.id)
}
</script>
