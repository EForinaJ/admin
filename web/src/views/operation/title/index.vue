<!-- 头衔管理页面 -->
<!-- art-full-height 自动计算出页面剩余高度 -->
<!-- art-table-card 一个符合系统样式的 class，同时自动撑满剩余高度 -->
<!-- 更多 useTable 使用示例请移步至 功能示例 下面的高级表格示例或者查看官方文档 -->
<!-- useTable 文档：https://www.artd.pro/docs/zh/guide/hooks/use-table.html -->
<template>
  <div class="title-page art-full-height">
    <!-- 搜索栏 -->
    <TitleSearch v-model="searchForm" @search="handleSearch" @reset="resetSearchParams"></TitleSearch>

    <ElCard class="art-table-card" shadow="never">
      <!-- 表格头部 -->
      <ArtTableHeader v-model:columns="columnChecks" :loading="loading" @refresh="refreshData">
        <template #left>
          <ElSpace wrap>
            <ElButton v-auth="'create'" @click="handleCreate" v-ripple>新增头衔</ElButton>
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

      <!-- 头衔弹窗 -->
      <TitleDialog
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
import TitleSearch from './modules/title-search.vue'
import { ElTag, ElMessageBox, ElImage } from 'element-plus'
import { useAuth } from '@/hooks'
import ArtButtonTable from '@/components/core/forms/art-button-table/index.vue'
import { fetchGetTitleList } from '@/api/title';
import TitleDialog from './modules/title-dialog.vue'



const { hasAuth } = useAuth();
defineOptions({ name: 'Title' })



// 弹窗相关
const dialogType = ref<"create" | "edit">('create')
const dialogVisible = ref(false)
const id = ref<number>(0)

// 选中行
const selectedRows = ref<number[]>([])

// 搜索表单
const searchForm = ref({
  name: undefined,
  gameId: undefined,
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
    apiFn: fetchGetTitleList,
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
        label: '头衔封面',
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
        label: '头衔名称'
      },
      { 
        prop: 'game', 
        label: '所属游戏',
        formatter: (row) => {
          return h(ElTag, { type: "primary" }, () => `${row.game}`)
        }
      },
      { 
        prop: 'servicePercent', 
        label: '服务费率',
        formatter: (row) => {
          return h(ElTag, { type: "primary" }, () => `${row.servicePercent}%`)
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
          return h('div', { class: 'title flex-c' }, [
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
 * 显示头衔弹窗
 */
const handleCreate = (): void => {
  dialogType.value = "create"
  id.value = 0
  nextTick(() => {
    dialogVisible.value = true
  })
}
const handleEdit = (row:Title.Response.Info) => {
   dialogType.value = "edit"
   id.value = row.id
   nextTick(() => {
    dialogVisible.value = true
  })
}

const handleBatchDelete = () =>{
  if (selectedRows.value.length != 0) {
    ElMessageBox.confirm(`确定要删除该吗？`, '删除头衔', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'error'
    }).then(async() => {
      // TODO: 调用删除接口
      ElMessage.success('删除成功')
      // await fetchPostTitleDelete({ids:selectedRows.value})
      refreshData()
    })
    .catch(() => {
      ElMessage.info('已取消删除')
    })
  }
}


const handleDelete = async (row: Title.Response.Info): Promise<void> => {
  ElMessageBox.confirm(`确定要删除该吗？`, '删除头衔', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'error'
  }).then(async() => {
    // TODO: 调用删除接口
    ElMessage.success('删除成功')
    // await fetchPostTitleDelete({ids:[row.id]})
    refreshData()
  })
  .catch(() => {
    ElMessage.info('已取消删除')
  })
}


/**
 * 处理表格行选择变化
 */
const handleSelectionChange = (selection: Title.Response.Info[]): void => {
  selectedRows.value = selection.map((item)=>item.id)
}
</script>
