<!-- 提现管理页面 -->
<!-- art-full-height 自动计算出页面剩余高度 -->
<!-- art-table-card 一个符合系统样式的 class，同时自动撑满剩余高度 -->
<!-- 更多 useTable 使用示例请移步至 功能示例 下面的高级表格示例或者查看官方文档 -->
<!-- useTable 文档：https://www.artd.pro/docs/zh/guide/hooks/use-table.html -->
<template>
  <div class="withdraw-page art-full-height">
    <!-- 搜索栏 -->
    <WithdrawSearch v-model="searchForm" @search="handleSearch" @reset="resetSearchParams"></WithdrawSearch>

    <ElCard class="art-table-card" shadow="never">
      <!-- 表格头部 -->
      <ArtTableHeader v-model:columns="columnChecks" :loading="loading" @refresh="refreshData">
        <template #left>
          <ElSpace wrap>
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


      <WithdrawApplyModal
        v-model:visible="applyModalVisible"
        :id="id"
        @submit="refreshData"
      />
      <WithdrawViewDrawer
        v-model="viewVisible"
        :id="id"
        @submit="refreshData"
      />
    </ElCard>
  </div>
</template>

<script setup lang="ts">
import { useTable } from '@/hooks/core/useTable'
import { ElTag, ElMessageBox } from 'element-plus'
import { useSiteStore } from '@/store/modules/site'
import { useAuth } from '@/hooks'
import ArtButtonMore, { ButtonMoreItem } from '@/components/core/forms/art-button-more/index.vue'
import ArtButtonTable from '@/components/core/forms/art-button-table/index.vue'

import WithdrawApplyModal from './modules/withdraw-apply-modal.vue'
import WithdrawSearch from './modules/withdraw-search.vue'

import { ApplyStatus } from '@/enums/statusEnum'
import WithdrawViewDrawer from './modules/withdraw-view-drawer.vue'
import { fetchGetWithdrawList } from '@/api/withdraw'


const { hasAuth } = useAuth();
defineOptions({ name: 'Withdraw' })

const {getInfo:site} = useSiteStore()
// 弹窗相关

const applyModalVisible = ref(false)
const viewVisible = ref(false)
const id = ref<number>(0)

// 选中行
const selectedRows = ref<number[]>([])

// 搜索表单
const searchForm = ref({
  code: undefined,
  name: undefined,
  status: undefined
})



const WITHDRAW_STATUS = {
  [ApplyStatus.Pending]: { type: 'primary' as const, text: '待审核' },
  [ApplyStatus.Success]: { type: 'success' as const, text: '已通过' },
  [ApplyStatus.Fail]: { type: 'danger' as const, text: '拒绝' },
} as const

const getStatus = (status: number) => {
  return (
    WITHDRAW_STATUS[status as keyof typeof WITHDRAW_STATUS] || {
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
    apiFn: fetchGetWithdrawList,
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
        label: '提现编号',
        width: 220,
      },
      {
        prop: 'witkey',
        label: '申请人',
        formatter: (row) => {
          return h('p', { }, row.witkey)
        }
      },
      {
        prop: 'amount',
        label: '提现金额',
        formatter: (row) => {
          return h(ElTag, { type:"primary" }, () => `${row.amount}${site.symbol}` )
        }
      },
      {
        prop: 'serviceFee',
        label: '平台手续费',
        formatter: (row) => {
          return h(ElTag, { type:"primary" }, () => `${row.serviceFee}${site.symbol}` )
        }
      },
      {
        prop: 'settledAmount',
        label: '到账金额',
        formatter: (row) => {
          return h(ElTag, { type:"primary" }, () => `${row.settledAmount}${site.symbol}` )
        }
      },
      {
        prop: 'status',
        label: '提现状态',
        formatter: (row) => {
          const statusConfig = getStatus(row.status)
          return h(ElTag, { type: statusConfig.type }, () => statusConfig.text)
        }
      },
      {
        prop: 'createTime',
        label: '申请时间',
        sortable: true
      },
      {
        prop: 'operation',
        label: '操作',
        width: 120,
        fixed: 'right', // 固定列
        formatter: (row) =>{
          return h('div', { class: 'withdraw flex-c' }, [
            (hasAuth("view") && h(ArtButtonTable, {
              type: 'view',
              onClick: () => handleView(row)
            })),
            ((row.status == ApplyStatus.Pending) && h(ArtButtonMore, {
              list: [{
                  key: 'apply',
                  label: '审核报单',
                  icon: 'ep:element-plus',
                  auth:'apply'
                },{
                  key: 'delete',
                  label: '删除订单',
                  icon: 'ri:delete-bin-4-line',
                  auth:'delete'
                },
              ],
              onClick: (item: ButtonMoreItem) => buttonMoreClick(item, row)
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

const buttonMoreClick = (item: ButtonMoreItem, row: Withdraw.Response.Info) => {
  switch (item.key) {
    case 'apply':
      handleApply(row)
      break
    case 'delete':
      handleDelete(row)
      break
  }
}

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

const handleView = (row:Withdraw.Response.Info) => {
  id.value = row.id
  nextTick(() => {
    viewVisible.value = true
  })
}

const handleApply = (row:Withdraw.Response.Info) => {
  id.value = row.id
  nextTick(() => {
    applyModalVisible.value = true
  })
}


const handleBatchDelete = () =>{
  if (selectedRows.value.length != 0) {
    ElMessageBox.confirm(`确定要删除该吗？`, '删除提现', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'error'
    }).then(async() => {
      // TODO: 调用删除接口
      ElMessage.success('删除成功')
      // await fetchPostWithdrawDelete({ids:selectedRows.value})
      refreshData()
    })
    .catch(() => {
      ElMessage.info('已取消删除')
    })
  }
}


const handleDelete = async (row: Withdraw.Response.Info): Promise<void> => {
  ElMessageBox.confirm(`确定要删除该吗？`, '删除提现', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'error'
  }).then(async() => {
    // TODO: 调用删除接口
    ElMessage.success('删除成功')
    // await fetchPostWithdrawDelete({ids:[row.id]})
    refreshData()
  })
  .catch(() => {
    ElMessage.info('已取消删除')
  })
}


/**
 * 处理表格行选择变化
 */
const handleSelectionChange = (selection: Withdraw.Response.Info[]): void => {
  selectedRows.value = selection.map((item)=>item.id)
}
</script>
