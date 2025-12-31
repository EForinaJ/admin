<!-- 资金管理页面 -->
<!-- art-full-height 自动计算出页面剩余高度 -->
<!-- art-table-card 一个符合系统样式的 class，同时自动撑满剩余高度 -->
<!-- 更多 useTable 使用示例请移步至 功能示例 下面的高级表格示例或者查看官方文档 -->
<!-- useTable 文档：https://www.artd.pro/docs/zh/guide/hooks/use-table.html -->
<template>
  <div class="capital-page art-full-height">
    <!-- 搜索栏 -->
    <CapitalSearch v-model="searchForm" @search="handleSearch" @reset="resetSearchParams"></CapitalSearch>

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

import { ApplyStatus } from '@/enums/statusEnum'
import { fetchGetCapitalList, fetchPostCapitalDelete } from '@/api/capital'
import { PayMode } from '@/enums/modeEnum'
import CapitalSearch from './modules/capital-search.vue'
import { CpitalType } from '@/enums/typeEnum'



const { hasAuth } = useAuth();
defineOptions({ name: 'Capital' })

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

const CAPITAL_PAY_TYPE = {
  [CpitalType.PaymentOrder]: { type: 'primary' as const, text: '订单支付' },
  [CpitalType.RefundOrder]: { type: 'primary' as const, text: '订单退款' },
  [CpitalType.PaymentRecharge]: { type: 'primary' as const, text: '余额充值' },
  [CpitalType.WithdrawCommission]: { type: 'primary' as const, text: '佣金提现' },
} as const

const getPayType = (type: number) => {
  return (
    CAPITAL_PAY_TYPE[type as keyof typeof CAPITAL_PAY_TYPE] || {
      type: 'info' as const,
      text: '未知'
    }
  )
}

const CAPITAL_PAY_MODE = {
  [PayMode.AlyPay]: { type: 'primary' as const, text: '资金宝' },
  [PayMode.Wechat]: { type: 'primary' as const, text: '微信' },
  [PayMode.Balance]: { type: 'primary' as const, text: '余额资金' },
  [PayMode.PersonalTransfer]: { type: 'primary' as const, text: '人工转账' },
} as const

const getPayMode = (status: number) => {
  return (
    CAPITAL_PAY_MODE[status as keyof typeof CAPITAL_PAY_MODE] || {
      type: 'info' as const,
      text: '未知'
    }
  )
}

const CAPITAL_STATUS = {
  [ApplyStatus.Pending]: { type: 'primary' as const, text: '未资金' },
  [ApplyStatus.Success]: { type: 'success' as const, text: '已资金' },
  [ApplyStatus.Fail]: { type: 'danger' as const, text: '资金失败' },
} as const

const getStatus = (status: number) => {
  return (
    CAPITAL_STATUS[status as keyof typeof CAPITAL_STATUS] || {
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
    apiFn: fetchGetCapitalList,
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
        label: '交易单号',
        width: 220,
        showOverflowTooltip: true,
      },
      {
        prop: 'related',
        label: '关联单号',
        width: 220,
        showOverflowTooltip: true,
      },
      {
        prop: 'user',
        label: '用户名称',
        showOverflowTooltip: true,
        formatter: (row) => {
          return h('p', { }, row.user)
        }
      },
      {
        prop: 'type',
        label: '交易类型',
        formatter: (row) => {
          const modeConfig = getPayType(row.type)
          return h(ElTag, { type:modeConfig.type }, () => modeConfig.text )
        }
      },
      {
        prop: 'amount',
        label: '交易金额',
        formatter: (row) => {
          return h(ElTag, { type:"primary" }, () => `${row.amount}${site.symbol}` )
        }
      },
      {
        prop: 'mode',
        label: '支付方式',
        formatter: (row) => {
          const modeConfig = getPayMode(row.mode)
          return h(ElTag, { type:modeConfig.type }, () => modeConfig.text )
        }
      },
      {
        prop: 'createTime',
        label: '交易时间',
        sortable: true,
        fixed: 'right', // 固定列
        minWidth: 120,
      },
      {
        prop: 'operation',
        label: '操作',
        minWidth: 60,
        fixed: 'right', // 固定列
        formatter: (row) =>{
          return h('div', { class: 'capital flex-c' }, [
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
    ElMessageBox.confirm(`确定要删除该吗？`, '删除资金', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'error'
    }).then(async() => {
      // TODO: 调用删除接口
      ElMessage.success('删除成功')
      await fetchPostCapitalDelete({ids:selectedRows.value})
      refreshData()
    })
    .catch(() => {
      ElMessage.info('已取消删除')
    })
  }
}


const handleDelete = async (row: Capital.Response.Info): Promise<void> => {
  ElMessageBox.confirm(`确定要删除该吗？`, '删除资金', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'error'
  }).then(async() => {
    // TODO: 调用删除接口
    ElMessage.success('删除成功')
    await fetchPostCapitalDelete({ids:[row.id]})
    refreshData()
  })
  .catch(() => {
    ElMessage.info('已取消删除')
  })
}


/**
 * 处理表格行选择变化
 */
const handleSelectionChange = (selection: Capital.Response.Info[]): void => {
  selectedRows.value = selection.map((item)=>item.id)
}
</script>
