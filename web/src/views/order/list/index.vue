<!-- 订单管理页面 -->
<!-- art-full-height 自动计算出页面剩余高度 -->
<!-- art-table-card 一个符合系统样式的 class，同时自动撑满剩余高度 -->
<!-- 更多 useTable 使用示例请移步至 功能示例 下面的高级表格示例或者查看官方文档 -->
<!-- useTable 文档：https://www.artd.pro/docs/zh/guide/hooks/use-table.html -->
<template>
  <div class="order-page art-full-height">
    <!-- 搜索栏 -->
    <OrderSearch v-model="searchForm" @search="handleSearch" @reset="resetSearchParams"></OrderSearch>

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

      <!-- 订单弹窗 -->
      <OrderViewDrawer
        v-model="viewDrawerVisible"
        :id="id"
        @submit="refreshData"
      />

      <OrderRefundModal
        v-model:visible="refundModalVisible"
        :id="id"
        @submit="refreshData"
      />
      <OrderAddDiscountModal
        v-model:visible="addDiscountModalVisible"
        :id="id"
        @submit="refreshData"
      />
    </ElCard>
  </div>
</template>

<script setup lang="ts">
import { useTable } from '@/hooks/core/useTable'
import { ElTag, ElMessageBox, ElImage } from 'element-plus'
import { useSiteStore } from '@/store/modules/site'
import { useAuth } from '@/hooks'
import { ButtonMoreItem } from '@/components/core/forms/art-button-more/index.vue'
import ArtButtonTable from '@/components/core/forms/art-button-table/index.vue'
import ArtButtonMore from '@/components/core/forms/art-button-more/index.vue'
import { fetchGetOrderList, fetchPostOrderCancel, fetchPostOrderComplete, fetchPostOrderDelete, fetchPostOrderPaid, fetchPostOrderStart } from '@/api/order'
import { OrderStatus} from '@/enums/statusEnum'
import OrderAddDiscountModal from './modules/order-add-discount-modal.vue'
import { PayMode } from '@/enums/modeEnum'
import OrderViewDrawer from './modules/order-view-drawer.vue'
import OrderRefundModal from './modules/order-refund-modal.vue'
import OrderSearch from './modules/order-search.vue'


const { hasAuth } = useAuth();
defineOptions({ name: 'Order' })

const siteStore = useSiteStore()



// 弹窗相关
const refundModalVisible = ref(false)
const viewDrawerVisible = ref(false)
const addDiscountModalVisible = ref(false)
const id = ref<number>(0)

// 选中行
const selectedRows = ref<number[]>([])

// 搜索表单
const searchForm = ref({
  name: undefined,
  phone: undefined,
  status: undefined
})

// 订单状态配置
const ORDER_STATUS_CONFIG = {
  [OrderStatus.PendingPayment]: { type: 'info' as const, text: '待支付' },
  [OrderStatus.PendingService]: { type: 'primary' as const, text: '待服务' },
  [OrderStatus.InProgress]: { type: 'primary' as const, text: '进行中' },
  [OrderStatus.Completed]: { type: 'success' as const, text: '已完成' },
  [OrderStatus.Cancel]: { type: 'danger' as const, text: '已取消' },
  [OrderStatus.Refund]: { type: 'warning' as const, text: '已退款' },
} as const

/**
 * 获取订单状态配置
 */
const getOrderStatusConfig = (status: number) => {
  return (
    ORDER_STATUS_CONFIG[status as keyof typeof ORDER_STATUS_CONFIG] || {
      type: 'info' as const,
      text: '未知'
    }
  )
}

// 支付状态配置
const ORDER_PAY_MODE = {
  [PayMode.AlyPay]: { type: 'primary' as const, text: '支付宝' },
  [PayMode.Wechat]: { type: 'success' as const, text: '微信' },
  [PayMode.Balance]: { type: 'warning' as const, text: '余额' },
  [PayMode.PersonalTransfer]: { type: 'primary' as const, text: '个人转账' },
} as const

/**
 * 获取支付状态配置
 */
const getOrderPayMode = (mode: number) => {
  return (
    ORDER_PAY_MODE[mode as keyof typeof ORDER_PAY_MODE] || {
      type: 'info' as const,
      text: '未支付'
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
    apiFn: fetchGetOrderList,
    apiParams:{
      code: "",
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
        label: '订单号',
        width: 200,
      },
      {
        prop: 'productInfo',
        label: '商品信息',
        width: 200,
        formatter: (row) => {
          return h('div', { class: 'flex-c' }, [
            h(ElImage, {
              class: 'size-12 rounded-md',
              src: row.product.pic,
              previewSrcList: [row.product.pic],
              // 图片预览是否插入至 body 元素上，用于解决表格内部图片预览样式异常
              previewTeleported: true
            }),
            h('div', { class: 'ml-2 flex-1' }, [
              h('p', { class: 'line-clamp-1' }, row.product.name),
              h(ElTag, { type:"primary",size:"small" }, () => row.product.game)
            ])
          ])
        }
      },
      {
        prop: 'orderInfo',
        label: '用户信息',
        width: 160,
        formatter: (row) => {
          return h('p', { }, row.user)
        }
      },
      {
        prop: 'totalAmount',
        label: '实付金额',
        formatter: (row) => {
          return h(ElTag, { type:"primary" }, () => `${row.payAmount}${siteStore.getInfo.symbol}`)
        }
      },
      {
        prop: 'payMode',
        label: '支付方式',
        formatter: (row) => {
          const mode = getOrderPayMode(row.payMode)
          return h(ElTag, { type: mode.type }, () => mode.text)
        }
      },
      {
        prop: 'status',
        label: '订单状态',
        formatter: (row) => {
          const statusConfig = getOrderStatusConfig(row.status)
          return h(ElTag, { type: statusConfig.type }, () => statusConfig.text)
        }
      },
      {
        prop: 'createTime',
        label: '下单时间',
        sortable: true
      },
      {
        prop: 'operation',
        label: '操作',
        width: 120,
        fixed: 'right', // 固定列
        formatter: (row) =>{
          const btnList:ButtonMoreItem[] = []
          
          if (row.status == OrderStatus.PendingPayment) {
            btnList.push({
              key: 'addDiscount',
              label: '添加优惠',
              icon: 'ep:element-plus',
              auth:'addDiscount'
            },
            {
              key: 'paid',
              label: '确认付款',
              icon: 'solar:hand-money-bold',
              auth:'paid'
            },
            {
              key: 'cancel',
              label: '关闭订单',
              icon: 'solar:close-circle-bold',
              color: '#f56c6c',
              auth:'cancel'
            },{
              key: 'delete',
              label: '删除订单',
              icon: 'ri:delete-bin-4-line',
              auth:'delete'
            })
          }

          if (row.status == OrderStatus.PendingService) {
            btnList.push({
              key: 'start',
              label: '开始服务',
              icon: 'solar:stars-minimalistic-bold',
              auth:'start'
            },{
              key: 'refund',
              label: '立即退款',
              icon: 'solar:electric-refueling-bold',
              auth:'refund'
            },{
              key: 'delete',
              label: '删除订单',
              icon: 'ri:delete-bin-4-line',
              auth:'delete'
            })
          }

          if (row.status == OrderStatus.InProgress) {
            btnList.push({
              key: 'complete',
              label: '完成订单',
              icon: 'solar:traffic-economy-bold',
              auth:'complete'
            },{
              key: 'refund',
              label: '立即退款',
              icon: 'solar:electric-refueling-bold',
              auth:'refund'
            },{
              key: 'delete',
              label: '删除订单',
              icon: 'ri:delete-bin-4-line',
              auth:'delete'
            })
          }

          if (row.status == OrderStatus.Completed) {
            btnList.push({
              key: 'refund',
              label: '立即退款',
              icon: 'solar:electric-refueling-bold',
              auth:'refund'
            },{
              key: 'delete',
              label: '删除订单',
              icon: 'ri:delete-bin-4-line',
              auth:'delete'
            })
          }

          if (row.status == OrderStatus.Cancel) {
            btnList.push({
              key: 'delete',
              label: '删除订单',
              icon: 'ri:delete-bin-4-line',
              auth:'delete'
            })
          }

          if (row.status == OrderStatus.Refund) {
            btnList.push({
              key: 'delete',
              label: '删除订单',
              icon: 'ri:delete-bin-4-line',
              auth:'delete'
            })
          }

          return h('div', { class: 'order flex-c' }, [
            (hasAuth("view") && h(ArtButtonTable, {
              type: 'view',
              onClick: () => handleView(row)
            })),
            h(ArtButtonMore, {
              list: btnList,
              onClick: (item: ButtonMoreItem) => buttonMoreClick(item, row)
            })
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

const buttonMoreClick = (item: ButtonMoreItem, row: Order.Response.Info) => {
  switch (item.key) {
    case 'addDiscount':
      handleAddDiscount(row)
      break
    case 'paid':
      handlePaid(row)
      break
    case 'start':
      handleStart(row)
      break
    case 'cancel':
      handleCancel(row)
      break
    case 'refund':
      handleRefund(row)
      break
    case 'complete':
      handleComplete(row)
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


const handleRefund = (row:Order.Response.Info) => {
   id.value = row.id
  nextTick(() => {
    refundModalVisible.value = true
  })
}
const handleView = (row:Order.Response.Info) => {
    id.value = row.id
    nextTick(() => {
      viewDrawerVisible.value = true
    })
}
const handleAddDiscount = (row:Order.Response.Info): void => {
  id.value = row.id
  nextTick(() => {
    addDiscountModalVisible.value = true
  })
}
const handlePaid = (row:Order.Response.Info): void => {
  ElMessageBox.confirm(`确定要修改订单为已支付吗？`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'primary'
  }).then(async() => {
    // TODO: 调用删除接口
    await fetchPostOrderPaid({id:row.id})
    refreshData()
  })
  .catch(() => {
    ElMessage.info('已取消')
  })
}
const handleCancel = (row:Order.Response.Info): void => {
  ElMessageBox.confirm(`确定要关闭订单吗？`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'primary'
  }).then(async() => {
    // TODO: 调用删除接口
    await fetchPostOrderCancel({id:row.id})
    refreshData()
  })
  .catch(() => {
    ElMessage.info('已取消')
  })
}

const handleStart = (row:Order.Response.Info): void => {
  ElMessageBox.confirm(`确定要开始服务吗？`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'primary'
  }).then(async() => {
    // TODO: 调用删除接口
    await fetchPostOrderStart({id:row.id})
    refreshData()
  })
  .catch(() => {
    ElMessage.info('已取消')
  })
}

const handleComplete = (row:Order.Response.Info): void => {
  ElMessageBox.confirm(`确定服务完成吗？`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'primary'
  }).then(async() => {
    // TODO: 调用删除接口
    await fetchPostOrderComplete({id:row.id})
    refreshData()
  })
  .catch(() => {
    ElMessage.info('已取消')
  })
}



const handleBatchDelete = () =>{
  if (selectedRows.value.length != 0) {
    ElMessageBox.confirm(`确定要删除该吗？`, '删除订单', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'error'
    }).then(async() => {
      // TODO: 调用删除接口
      ElMessage.success('删除成功')
      await fetchPostOrderDelete({ids:selectedRows.value})
      refreshData()
    })
    .catch(() => {
      ElMessage.info('已取消删除')
    })
  }
}


const handleDelete = async (row: Order.Response.Info): Promise<void> => {
  ElMessageBox.confirm(`确定要删除该吗？`, '删除订单', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'error'
  }).then(async() => {
    // TODO: 调用删除接口
    ElMessage.success('删除成功')
    await fetchPostOrderDelete({ids:[row.id]})
    refreshData()
  })
  .catch(() => {
    ElMessage.info('已取消删除')
  })
}


/**
 * 处理表格行选择变化
 */
const handleSelectionChange = (selection: Order.Response.Info[]): void => {
  selectedRows.value = selection.map((item)=>item.id)
}
</script>
