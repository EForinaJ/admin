<!-- 游戏管理页面 -->
<!-- art-full-height 自动计算出页面剩余高度 -->
<!-- art-table-card 一个符合系统样式的 class，同时自动撑满剩余高度 -->
<!-- 更多 useTable 使用示例请移步至 功能示例 下面的高级表格示例或者查看官方文档 -->
<!-- useTable 文档：https://www.artd.pro/docs/zh/guide/hooks/use-table.html -->
<template>
  <div class="game-page art-full-height">
    <!-- 搜索栏 -->
    <GameSearch v-model="searchForm" @search="handleSearch" @reset="resetSearchParams"></GameSearch>

    <ElCard class="art-table-card" shadow="never">
      <!-- 表格头部 -->
      <ArtTableHeader v-model:columns="columnChecks" :loading="loading" @refresh="refreshData">
        <template #left>
          <ElSpace wrap>
            <ElButton v-auth="'create'" @click="handleCreate" v-ripple>新增游戏</ElButton>
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

      <!-- 游戏弹窗 -->
      <GameDialog
        v-model="dialogVisible"
        :type="dialogType"
        :id="id"
        @submit="refreshData"
      />
    </ElCard>
  </div>
</template>

<script setup lang="ts">
import { useTable } from '@/hooks/core/useTable'
import { ElTag, ElMessageBox, ElImage } from 'element-plus'
import { useAuth } from '@/hooks'
import ArtButtonTable from '@/components/core/forms/art-button-table/index.vue'

import { fetchGetGameList } from '@/api/game'
import GameDialog from './modules/game-dialog.vue'
import GameSearch from './modules/game-search.vue'

const { hasAuth } = useAuth();
defineOptions({ name: 'Game' })



// 弹窗相关
const dialogType = ref<"create" | "edit">('create')
const dialogVisible = ref(false)
const id = ref<number>(0)

// 选中行
const selectedRows = ref<number[]>([])

// 搜索表单
const searchForm = ref({
  name: undefined,
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
    apiFn: fetchGetGameList,
    apiParams:{
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
        prop: 'pic',
        label: '游戏封面',
        width: 120,
        formatter: (row) => {
          return  h(ElImage, {
              class: 'size-9.5 rounded-md',
              src: row.pic,
              previewSrcList: [row.pic],
              // 图片预览是否插入至 body 元素上，用于解决表格内部图片预览样式异常
              previewTeleported: true
            })
        }
      },
      { 
        prop: 'name', 
        label: '游戏名称'
      },
      {
        prop: 'description',
        label: '描述',
        showOverflowTooltip: true,
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
          return h('div', { class: 'game flex-c' }, [
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
 * 显示游戏弹窗
 */
const handleCreate = (): void => {
  dialogType.value = "create"
  id.value = 0
  nextTick(() => {
    dialogVisible.value = true
  })
}
const handleEdit = (row:Game.Response.Info) => {
   dialogType.value = "edit"
   id.value = row.id
   nextTick(() => {
    dialogVisible.value = true
  })
}

const handleBatchDelete = () =>{
  if (selectedRows.value.length != 0) {
    ElMessageBox.confirm(`确定要删除该吗？`, '删除游戏', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'error'
    }).then(async() => {
      // TODO: 调用删除接口
      ElMessage.success('删除成功')
      // await fetchPostGameDelete({ids:selectedRows.value})
      refreshData()
    })
    .catch(() => {
      ElMessage.info('已取消删除')
    })
  }
}


const handleDelete = async (row: Game.Response.Info): Promise<void> => {
  ElMessageBox.confirm(`确定要删除该吗？`, '删除游戏', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'error'
  }).then(async() => {
    // TODO: 调用删除接口
    ElMessage.success('删除成功')
    // await fetchPostGameDelete({ids:[row.id]})
    refreshData()
  })
  .catch(() => {
    ElMessage.info('已取消删除')
  })
}


/**
 * 处理表格行选择变化
 */
const handleSelectionChange = (selection: Game.Response.Info[]): void => {
  selectedRows.value = selection.map((item)=>item.id)
}
</script>
