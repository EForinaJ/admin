<!-- 威客管理页面 -->
<!-- art-full-height 自动计算出页面剩余高度 -->
<!-- art-table-card 一个符合系统样式的 class，同时自动撑满剩余高度 -->
<!-- 更多 useTable 使用示例请移步至 功能示例 下面的高级表格示例或者查看官方文档 -->
<!-- useTable 文档：https://www.artd.pro/docs/zh/guide/hooks/use-table.html -->
<template>
  <div class="witkey-page art-full-height">
    <!-- 搜索栏 -->
    <WitkeySearch v-model="searchForm" @search="handleSearch" @reset="resetSearchParams"></WitkeySearch>

    <ElCard class="art-table-card" shadow="never">
      <!-- 表格头部 -->
      <ArtTableHeader v-model:columns="columnChecks" :loading="loading" @refresh="refreshData">
        <template #left>
          <ElSpace wrap>
            <ElButton v-auth="'create'" @click="handleCreate" v-ripple>新增威客</ElButton>
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

      <!-- 弹窗 -->
      <WitkeyDialog
        v-model="dialogVisible"
        :type="dialogType"
        :id="id"
        @submit="refreshData"
      />
      <!-- 威客弹窗 -->
      <WitkeyViewDrawer
        v-model="viewDrawerVisible"
        :id="id"
        @submit="refreshData"
      />

      <!-- 变更弹窗 -->
      <WitkeyChangeCommissionModal
        v-model:visible="changeCommissionModalVisible"
        :id="id"
        @submit="refreshData"
      />
      <!-- 充值弹窗 -->
      <!-- <WitkeyRechargeModal
        v-model:visible="rechargeModalVisible"
        :id="id"
        @submit="refreshData"
      /> -->
    </ElCard>
  </div>
</template>

<script setup lang="ts">
import { useTable } from '@/hooks/core/useTable'
import WitkeySearch from './modules/witkey-search.vue'
import { ElTag, ElMessageBox, ElImage, ElRate } from 'element-plus'
import { useAuth } from '@/hooks'
import WitkeyDialog from './modules/witkey-dialog.vue'
import { ButtonMoreItem } from '@/components/core/forms/art-button-more/index.vue'
import { fetchGetWitkeyList, fetchPostWitkeyDelete } from '@/api/witkey'
import ArtButtonTable from '@/components/core/forms/art-button-table/index.vue'
import ArtButtonMore from '@/components/core/forms/art-button-more/index.vue'
import WitkeyViewDrawer from './modules/witkey-view-drawer.vue'
import WitkeyChangeCommissionModal from './modules/witkey-change-commission-modal.vue'


const { hasAuth } = useAuth();
defineOptions({ name: 'Witkey' })


// 弹窗相关
const dialogType = ref<"create" | "edit">('create')
const dialogVisible = ref(false)
const viewDrawerVisible = ref(false)
const changeCommissionModalVisible = ref(false)
const rechargeModalVisible = ref(false)
const id = ref<number>(0)

// 选中行
const selectedRows = ref<number[]>([])

// 搜索表单
const searchForm = ref({
  name: undefined,
  phone: undefined,
})


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
    apiFn: fetchGetWitkeyList,
    apiParams:{
      name: "",
      phone: "",
    },
    paginationKey:{
      current: 'page', 
      size: 'limit'
    },
    columnsFactory: () => [
      { type: 'selection' }, // 勾选列
      { prop: 'id', width: 60, label: 'ID' }, // 序号
      {
        prop: 'userInfo',
        label: '所属用户',
        width: 280,
        formatter: (row) => {
          return h('div', { class: 'user flex-c' }, [
            h(ElImage, {
              class: 'size-9.5 rounded-md',
              src: row.user.avatar,
              previewSrcList: [row.user.avatar],
              // 图片预览是否插入至 body 元素上，用于解决表格内部图片预览样式异常
              previewTeleported: true
            }),
            h('div', { class: 'ml-2' }, [
              h('p', { class: 'user-name' }, row.user.name),
              h('p', { class: 'email' }, row.user.phone)
            ])
          ])
        }
      },
      { 
        prop: 'game', 
        label: '所属游戏',
        formatter: (row) => {
          return h(ElTag, { type: "primary" }, () => `${row.game}`)
        }
      },
      { 
        prop: 'title', 
        label: '所属头衔',
        formatter: (row) => {
          return h(ElTag, { type: "primary" }, () => `${row.title}`)
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
      // { 
      //   prop: 'commission', 
      //   label: '佣金',
      //   formatter: (row) => {
      //     return h(ElTag, { type: "primary" }, () => `${row.commission}${siteStore.getInfo.symbol}`)
      //   }
      // },
      {
        prop: 'createTime',
        label: '入驻日期',
        sortable: true
      },
      {
        prop: 'operation',
        label: '操作',
        width: 120,
        fixed: 'right', // 固定列
        formatter: (row) =>{
          return h('div', { class: 'witkey flex-c' }, [
            (hasAuth("view") && h(ArtButtonTable, {
              type: 'view',
              onClick: () => handleView(row)
            })),
            h(ArtButtonMore, {
              list: [
                {
                  key: 'changeCommission',
                  label: '变更佣金',
                  icon: 'ep:element-plus',
                  auth:'changeCommission'
                },
                {
                  key: 'edit',
                  label: '编辑角色',
                  icon: 'ri:edit-2-line',
                  auth:'edit'
                },
                {
                  key: 'delete',
                  label: '删除角色',
                  icon: 'ri:delete-bin-4-line',
                  color: '#f56c6c',
                  auth:'delete'
                }
              ],
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

const buttonMoreClick = (item: ButtonMoreItem, row: Witkey.Response.Info) => {
  switch (item.key) {
    case 'changeCommission':
      handleChangeCommission(row)
      break
    case 'recharge':
      handleRecharge(row)
      break
    case 'edit':
      handleEdit(row)
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
  // 搜索参数赋值
  Object.assign(searchParams, params)
  getData()
}

/**
 * 显示威客弹窗
 */
const handleCreate = (): void => {
  dialogType.value = "create"
  id.value = 0
  nextTick(() => {
    dialogVisible.value = true
  })
}
const handleEdit = (row:Witkey.Response.Info) => {
   dialogType.value = "edit"
   id.value = row.id
   nextTick(() => {
    dialogVisible.value = true
  })
}
const handleView = (row:Witkey.Response.Info) => {
    id.value = row.id
    nextTick(() => {
      viewDrawerVisible.value = true
    })
}
const handleChangeCommission = (row:Witkey.Response.Info): void => {
  id.value = row.id
  nextTick(() => {
    changeCommissionModalVisible.value = true
  })
}
const handleRecharge = (row:Witkey.Response.Info): void => {
  id.value = row.id
  nextTick(() => {
    rechargeModalVisible.value = true
  })
}

const handleBatchDelete = () =>{
  if (selectedRows.value.length != 0) {
    ElMessageBox.confirm(`确定要删除该吗？`, '删除威客', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'error'
    }).then(async() => {
      // TODO: 调用删除接口
      ElMessage.success('删除成功')
      await fetchPostWitkeyDelete({ids:selectedRows.value})
      refreshData()
    })
    .catch(() => {
      ElMessage.info('已取消删除')
    })
  }
}


const handleDelete = async (row: Witkey.Response.Info): Promise<void> => {
  ElMessageBox.confirm(`确定要删除该吗？`, '删除威客', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'error'
  }).then(async() => {
    // TODO: 调用删除接口
    ElMessage.success('删除成功')
    await fetchPostWitkeyDelete({ids:[row.id]})
    refreshData()
  })
  .catch(() => {
    ElMessage.info('已取消删除')
  })
}


/**
 * 处理表格行选择变化
 */
const handleSelectionChange = (selection: Witkey.Response.Info[]): void => {
  selectedRows.value = selection.map((item)=>item.id)
}
</script>
