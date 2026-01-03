<!-- 充值管理页面 -->
<!-- art-full-height 自动计算出页面剩余高度 -->
<!-- art-table-card 一个符合系统样式的 class，同时自动撑满剩余高度 -->
<!-- 更多 useTable 使用示例请移步至 功能示例 下面的高级表格示例或者查看官方文档 -->
<!-- useTable 文档：https://www.artd.pro/docs/zh/guide/hooks/use-table.html -->
<template>
  <div class="recharge-page art-full-height">
    <!-- 搜索栏 -->
    <RechargeSearch v-model="searchForm" @search="handleSearch" @reset="resetSearchParams"></RechargeSearch>

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

    </ElCard>
  </div>
</template>

<script setup lang="ts">
import { useTable } from '@/hooks/core/useTable'
import { ElTag, ElMessageBox } from 'element-plus'
import { useSiteStore } from '@/store/modules/site'
import { useAuth } from '@/hooks'
import ArtButtonTable from '@/components/core/forms/art-button-table/index.vue'
import RechargeSearch from './modules/recharge-search.vue'
import { ApplyStatus } from '@/enums/statusEnum'
import { fetchGetRechargeList, fetchPostRechargeDelete } from '@/api/recharge'
import { PayMode } from '@/enums/modeEnum'


const { hasAuth } = useAuth();
defineOptions({ name: 'Recharge' })

const {getInfo:site} = useSiteStore()
// 弹窗相关

const id = ref<number>(0)

// 选中行
const selectedRows = ref<number[]>([])

// 搜索表单
const searchForm = ref({
  code: undefined,
  name: undefined,
  status: undefined
})

const RECHARGE_PAY_MODE = {
  [PayMode.AlyPay]: { type: 'primary' as const, text: '支付宝' },
  [PayMode.Wechat]: { type: 'primary' as const, text: '微信' },
  [PayMode.Balance]: { type: 'primary' as const, text: '余额充值' },
  [PayMode.PersonalTransfer]: { type: 'primary' as const, text: '人工转账' },
} as const

const getPayMode = (status: number) => {
  return (
    RECHARGE_PAY_MODE[status as keyof typeof RECHARGE_PAY_MODE] || {
      type: 'info' as const,
      text: '未知'
    }
  )
}

const RECHARGE_STATUS = {
  [ApplyStatus.Pending]: { type: 'primary' as const, text: '未充值' },
  [ApplyStatus.Success]: { type: 'success' as const, text: '已充值' },
  [ApplyStatus.Fail]: { type: 'danger' as const, text: '充值失败' },
} as const

const getStatus = (status: number) => {
  return (
    RECHARGE_STATUS[status as keyof typeof RECHARGE_STATUS] || {
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
    apiFn: fetchGetRechargeList,
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
        label: '充值编号',
        width: 270,
      },
      {
        prop: 'user',
        label: '用户名称',
        formatter: (row) => {
          return h('p', { }, row.user)
        }
      },
      {
        prop: 'amount',
        label: '充值金额',
        formatter: (row) => {
          return h(ElTag, { type:"primary" }, () => `${row.amount}${site.symbol}` )
        }
      },
      {
        prop: 'payMode',
        label: '充值方式',
        formatter: (row) => {
          const modeConfig = getPayMode(row.payMode)
          return h(ElTag, { type:modeConfig.type }, () => modeConfig.text )
        }
      },
      {
        prop: 'status',
        label: '充值状态',
        formatter: (row) => {
          const statusConfig = getStatus(row.status)
          return h(ElTag, { type: statusConfig.type }, () => statusConfig.text)
        }
      },
      {
        prop: 'createTime',
        label: '充值时间',
        sortable: true
      },
      {
        prop: 'operation',
        label: '操作',
        width: 120,
        fixed: 'right', // 固定列
        formatter: (row) =>{
          return h('div', { class: 'recharge flex-c' }, [
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


const handleBatchDelete = () =>{
  if (selectedRows.value.length != 0) {
    ElMessageBox.confirm(`确定要删除该吗？`, '删除充值', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'error'
    }).then(async() => {
      // TODO: 调用删除接口
      ElMessage.success('删除成功')
      await fetchPostRechargeDelete({ids:selectedRows.value})
      refreshData()
    })
    .catch(() => {
      ElMessage.info('已取消删除')
    })
  }
}


const handleDelete = async (row: Recharge.Response.Info): Promise<void> => {
  ElMessageBox.confirm(`确定要删除该吗？`, '删除充值', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'error'
  }).then(async() => {
    // TODO: 调用删除接口
    ElMessage.success('删除成功')
    await fetchPostRechargeDelete({ids:[row.id]})
    refreshData()
  })
  .catch(() => {
    ElMessage.info('已取消删除')
  })
}


/**
 * 处理表格行选择变化
 */
const handleSelectionChange = (selection: Recharge.Response.Info[]): void => {
  selectedRows.value = selection.map((item)=>item.id)
}
</script>
