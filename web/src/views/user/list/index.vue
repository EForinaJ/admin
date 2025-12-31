<!-- 用户管理页面 -->
<!-- art-full-height 自动计算出页面剩余高度 -->
<!-- art-table-card 一个符合系统样式的 class，同时自动撑满剩余高度 -->
<!-- 更多 useTable 使用示例请移步至 功能示例 下面的高级表格示例或者查看官方文档 -->
<!-- useTable 文档：https://www.artd.pro/docs/zh/guide/hooks/use-table.html -->
<template>
  <div class="user-page art-full-height">
    <!-- 搜索栏 -->
    <UserSearch v-model="searchForm" @search="handleSearch" @reset="resetSearchParams"></UserSearch>

    <ElCard class="art-table-card" shadow="never">
      <!-- 表格头部 -->
      <ArtTableHeader v-model:columns="columnChecks" :loading="loading" @refresh="refreshData">
        <template #left>
          <ElSpace wrap>
            <ElButton v-auth="'create'" @click="handleCreate" v-ripple>新增用户</ElButton>
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

      <!-- 用户弹窗 -->
      <UserDrawer
        v-model="drawerVisible"
        :type="drawerType"
        :id="id"
        @submit="refreshData"
      />
      <!-- 用户弹窗 -->
      <UserViewDrawer
        v-model="viewDrawerVisible"
        :id="id"
        @submit="refreshData"
      />

      <!-- 变更弹窗 -->
      <UserChangeBalanceModal
        v-model:visible="changeBalanceModalVisible"
        :id="id"
        @submit="refreshData"
      />
      <!-- 充值弹窗 -->
      <UserRechargeModal
        v-model:visible="rechargeModalVisible"
        :id="id"
        @submit="refreshData"
      />
    </ElCard>
  </div>
</template>

<script setup lang="ts">
import { useTable } from '@/hooks/core/useTable'
import UserSearch from './modules/user-search.vue'
import { ElTag, ElMessageBox, ElImage } from 'element-plus'
import { fetchGetUserList } from '@/api/user'
import { useSiteStore } from '@/store/modules/site'
import { fetchPostUserDelete } from '@/api/user'
import { useAuth } from '@/hooks'
import UserDrawer from './modules/user-drawer.vue'

import UserViewDrawer from './modules/user-view-drawer.vue'
import { ButtonMoreItem } from '@/components/core/forms/art-button-more/index.vue'
import ArtButtonTable from '@/components/core/forms/art-button-table/index.vue'
import ArtButtonMore from '@/components/core/forms/art-button-more/index.vue'
import UserChangeBalanceModal from './modules/user-change-balance-modal.vue'
import UserRechargeModal from './modules/user-recharge-modal.vue'

const { hasAuth } = useAuth();
defineOptions({ name: 'User' })

const siteStore = useSiteStore()



// 弹窗相关
const drawerType = ref<"create" | "edit">('create')
const drawerVisible = ref(false)
const viewDrawerVisible = ref(false)
const changeBalanceModalVisible = ref(false)
const rechargeModalVisible = ref(false)
const id = ref<number>(0)

// 选中行
const selectedRows = ref<number[]>([])

// 搜索表单
const searchForm = ref({
  name: undefined,
  phone: undefined,
  status: undefined
})

// 用户状态配置
const USER_STATUS_CONFIG = {
  1: { type: 'info' as const, text: '禁用' },
  2: { type: 'success' as const, text: '启用' },
  3: { type: 'warning' as const, text: '异常' },
  4: { type: 'danger' as const, text: '注销' }
} as const

/**
 * 获取用户状态配置
 */
const getUserStatusConfig = (status: number) => {
  return (
    USER_STATUS_CONFIG[status as keyof typeof USER_STATUS_CONFIG] || {
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
    apiFn: fetchGetUserList,
    apiParams:{
      name: "",
      phone: "",
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
        prop: 'userInfo',
        label: '用户名',
        width: 280,
        formatter: (row) => {
          return h('div', { class: 'user flex-c' }, [
            h(ElImage, {
              class: 'size-9.5 rounded-md',
              src: row.avatar,
              previewSrcList: [row.avatar],
              // 图片预览是否插入至 body 元素上，用于解决表格内部图片预览样式异常
              previewTeleported: true
            }),
            h('div', { class: 'ml-2' }, [
              h('p', { class: 'user-name' }, row.name),
              h('p', { class: 'email' }, row.phone)
            ])
          ])
        }
      },
      {
        prop: 'userSex',
        label: '性别',
        sortable: true,
        formatter: (row) => {
          const label: Record<number, string> = {
            1: '男',
            2: '女',
            3: '其他'
          };

          return label[row.sex]
        }
      },
      { 
        prop: 'balance', 
        label: '余额',
        formatter: (row) => {
          return h(ElTag, { type: "primary" }, () => `${row.balance}${siteStore.getInfo.symbol}`)
        } },
      {
        prop: 'status',
        label: '状态',
        formatter: (row) => {
          const statusConfig = getUserStatusConfig(row.status)
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
          return h('div', { class: 'user flex-c' }, [
            (hasAuth("view") && h(ArtButtonTable, {
              type: 'view',
              onClick: () => handleView(row)
            })),
            h(ArtButtonMore, {
              list: [
                {
                  key: 'changeBalance',
                  label: '变更余额',
                  icon: 'ep:element-plus',
                  auth:'changeBalance'
                },
                {
                  key: 'recharge',
                  label: '预存余额',
                  icon: 'ep:pear',
                  auth:'recharge'
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

const buttonMoreClick = (item: ButtonMoreItem, row: User.Response.Info) => {
  switch (item.key) {
    case 'changeBalance':
      handleChangeBalance(row)
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
  console.log(params)
  // 搜索参数赋值
  Object.assign(searchParams, params)
  getData()
}

/**
 * 显示用户弹窗
 */
const handleCreate = (): void => {
  drawerType.value = "create"
  id.value = 0
  nextTick(() => {
    drawerVisible.value = true
  })
}
const handleEdit = (row:User.Response.Info) => {
   drawerType.value = "edit"
   id.value = row.id
   nextTick(() => {
    drawerVisible.value = true
  })
}
const handleView = (row:User.Response.Info) => {
    id.value = row.id
    nextTick(() => {
      viewDrawerVisible.value = true
    })
}
const handleChangeBalance = (row:User.Response.Info): void => {
  id.value = row.id
  nextTick(() => {
    changeBalanceModalVisible.value = true
  })
}
const handleRecharge = (row:User.Response.Info): void => {
  id.value = row.id
  nextTick(() => {
    rechargeModalVisible.value = true
  })
}

const handleBatchDelete = () =>{
  if (selectedRows.value.length != 0) {
    ElMessageBox.confirm(`确定要删除该吗？`, '删除用户', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'error'
    }).then(async() => {
      // TODO: 调用删除接口
      ElMessage.success('删除成功')
      await fetchPostUserDelete({ids:selectedRows.value})
      refreshData()
    })
    .catch(() => {
      ElMessage.info('已取消删除')
    })
  }
}


const handleDelete = async (row: User.Response.Info): Promise<void> => {
  ElMessageBox.confirm(`确定要删除该吗？`, '删除用户', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'error'
  }).then(async() => {
    // TODO: 调用删除接口
    ElMessage.success('删除成功')
    await fetchPostUserDelete({ids:[row.id]})
    refreshData()
  })
  .catch(() => {
    ElMessage.info('已取消删除')
  })
}


/**
 * 处理表格行选择变化
 */
const handleSelectionChange = (selection: User.Response.Info[]): void => {
  selectedRows.value = selection.map((item)=>item.id)
}
</script>
