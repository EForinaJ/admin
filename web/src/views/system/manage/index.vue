<!-- 员工管理页面 -->
<!-- art-full-height 自动计算出页面剩余高度 -->
<!-- art-table-card 一个符合系统样式的 class，同时自动撑满剩余高度 -->
<!-- 更多 useTable 使用示例请移步至 功能示例 下面的高级表格示例或者查看官方文档 -->
<!-- useTable 文档：https://www.artd.pro/docs/zh/guide/hooks/use-table.html -->
<template>
  <div class="user-page art-full-height">
    <!-- 搜索栏 -->
    <ManageSearch v-model="searchForm" 
    @search="handleSearch"
     @reset="resetSearchParams" />

    <ElCard class="art-table-card" shadow="never">
      <!-- 表格头部 -->
      <ArtTableHeader v-model:columns="columnChecks" :loading="loading" @refresh="refreshData">
        <template #left>
          <ElSpace wrap>
            <ElButton v-auth="'create'" @click="handleCreate" v-ripple>新增员工</ElButton>
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

      <!-- 员工弹窗 -->
      <ManageDrawer
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
import ManageDrawer from './modules/manage-drawer.vue'
import { ElTag, ElMessageBox, ElImage, ElSpace } from 'element-plus'
import { fetchManageDelete, fetchManageList } from '@/api/manage'
import { useAuth } from '@/hooks'
import ArtButtonTable from '@/components/core/forms/art-button-table/index.vue'
import ManageSearch from './modules/manage-search.vue'

const { hasAuth } = useAuth();
defineOptions({ name: 'Manage' })


// 弹窗相关
const drawerType = ref<'create' | 'edit'>('create')
const drawerVisible = ref(false)
const id = ref<number | null>(null)

// 选中行
const selectedRows = ref<number[]>([])

// 搜索表单
const searchForm = ref({
  name: undefined,
})

// 员工状态配置
const USER_STATUS_CONFIG = {
  1: { type: 'info' as const, text: '禁用' },
  2: { type: 'success' as const, text: '启用' },
  3: { type: 'warning' as const, text: '异常' },
  4: { type: 'danger' as const, text: '注销' }
} as const

/**
 * 获取员工状态配置
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
    apiFn: fetchManageList,
    apiParams: {
      ...searchForm.value
    },
    // 自定义分页字段映射，未设置时将使用全局配置 tableConfig.ts 中的 paginationKey
    paginationKey:{
      current: 'page', 
      size: 'limit'
    },
    columnsFactory: () => [
      { type: 'selection' }, // 勾选列
      {  
        prop: 'id', width: 100, label: '员工ID' }, // 序号
      {
        prop: 'userInfo',
        label: '员工名',
        width: 280,
        // visible: false, // 默认是否显示列
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
        prop: 'role',
        label: '角色',
        formatter: (row) => {
          const roles = row.roles.map((item)=>{
            return h(ElTag, { type: 'success' }, () => item)
          })
          return h(ElSpace,{},() => roles )
        }
      },
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
        formatter: (row) =>
          h('div', [
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
    ]
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
  // 搜索参数赋值
  Object.assign(searchParams, params)
  getData()
}

/**
 * 显示员工弹窗
 */
const handleCreate = (): void => {
  drawerType.value = "create"
  // currentUserData.value = row || {}
  nextTick(() => {
    drawerVisible.value = true
  })
}

const handleEdit = (row:Manage.Response.Info) => {
   drawerType.value = "edit"
   id.value = row.id
   nextTick(() => {
    drawerVisible.value = true
  })
}

const handleBatchDelete = () =>{
  if (selectedRows.value.length != 0) {
    ElMessageBox.confirm(`确定要删除该员工吗？`, '删除员工', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'error'
    }).then(async() => {
      // TODO: 调用删除接口
      ElMessage.success('删除成功')
      await fetchManageDelete({ids:selectedRows.value})
      refreshData()
    })
    .catch(() => {
      ElMessage.info('已取消删除')
    })
  }
}


const handleDelete = async (row: Manage.Response.Info): Promise<void> => {
  ElMessageBox.confirm(`确定要删除该员工吗？`, '删除员工', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'error'
  }).then(async() => {
    // TODO: 调用删除接口
    ElMessage.success('删除成功')
    await fetchManageDelete({ids:[row.id]})
    refreshData()
  })
  .catch(() => {
    ElMessage.info('已取消删除')
  })
}


/**
 * 处理表格行选择变化
 */
const handleSelectionChange = (selection:Manage.Response.Info[]): void => {
  selectedRows.value = selection.map((item)=>item.id)
}
</script>
