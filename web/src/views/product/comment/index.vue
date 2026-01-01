<!-- 提现管理页面 -->
<!-- art-full-height 自动计算出页面剩余高度 -->
<!-- art-table-card 一个符合系统样式的 class，同时自动撑满剩余高度 -->
<!-- 更多 useTable 使用示例请移步至 功能示例 下面的高级表格示例或者查看官方文档 -->
<!-- useTable 文档：https://www.artd.pro/docs/zh/guide/hooks/use-table.html -->
<template>
  <div class="comment-page art-full-height">
    <!-- 搜索栏 -->
    <CommentSearch v-model="searchForm" @search="handleSearch" @reset="resetSearchParams"></CommentSearch>

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
import { ElTag, ElMessageBox, ElImage, ElRate, ElSpace, ElPopconfirm } from 'element-plus'
import { useSiteStore } from '@/store/modules/site'
import { useAuth } from '@/hooks'
import ArtButtonTable from '@/components/core/forms/art-button-table/index.vue'


import { ApplyStatus } from '@/enums/statusEnum'
import { fetchGetCommentList, fetchPostCommentApply, fetchPostCommentDelete } from '@/api/comment'
import CommentSearch from './modules/comment-search.vue'


const { hasAuth } = useAuth();
defineOptions({ name: 'Comment' })

const {getInfo:site} = useSiteStore()
// 弹窗相关
const id = ref<number>(0)

// 选中行
const selectedRows = ref<number[]>([])

// 搜索表单
const searchForm = ref({
  code: undefined,
  status: undefined
})



const COMMENT_STATUS = {
  [ApplyStatus.Pending]: { type: 'primary' as const, text: '待审核' },
  [ApplyStatus.Success]: { type: 'success' as const, text: '已通过' },
  [ApplyStatus.Fail]: { type: 'danger' as const, text: '拒绝' },
} as const

const getStatus = (status: number) => {
  return (
    COMMENT_STATUS[status as keyof typeof COMMENT_STATUS] || {
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
    apiFn: fetchGetCommentList,
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
        type: 'expand',
        prop: 'content',
        formatter: (row) => {
          return h('p', { }, `评价内容: ${row.content}`)
        }
      }, 
      {
        prop: 'user',
        label: '评价人',
        formatter: (row) => {
          return h('p', { }, row.user)
        }
      },
      {
        prop: 'productInfo',
        label: '商品信息',
        width: 280,
        formatter: (row) => {
          return h('div', { class: 'product flex-c' }, [
            h(ElImage, {
              class: 'size-9.5 rounded-md',
              src: row.product.pic,
              previewSrcList: [row.product.pic],
              // 图片预览是否插入至 body 元素上，用于解决表格内部图片预览样式异常
              previewTeleported: true
            }),
            h('div', { class: 'ml-2' }, [
              h('p', { class: 'product-name' }, row.product.name),
              h('p', { class: 'email' }, row.product.category),
            ])
          ])
        }
      },
      { 
        prop: 'rate', 
        label: '评分',
        formatter: (row) => {
          return h(ElRate , {
            modelValue: row.rate,
            disabled: true,
          })
        }
      },
      { 
        prop: 'images', 
        label: '晒单',
        formatter: (row) => {
          const images = row.images.map((item) => {
            return  h(ElImage, {
              class: 'size-9.5 rounded-md',
              src: item,
              previewSrcList: [item],
              // 图片预览是否插入至 body 元素上，用于解决表格内部图片预览样式异常
              previewTeleported: true
            })
          })
          return h(ElSpace, {wrap:true},images)
        }
      },
      {
        prop: 'status',
        label: '审核状态',
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
          return h('div', { class: 'comment flex-c' }, [
            // (hasAuth("apply") && h(ArtButtonTable, {
            //   type: 'view',
            //   icon: 'solar:checklist-minimalistic-bold',
            //   onClick: () => handleApply(row)
            // })),
            ((hasAuth("apply") && row.status == ApplyStatus.Pending) && h(ElPopconfirm, {
              title: "评论是否通过审核",
              confirmButtonText: '是',
              cancelButtonText: '否',
              onConfirm: () => handleApply({id:row.id,status:ApplyStatus.Success}),
              onCancel: () => handleApply({id:row.id,status:ApplyStatus.Fail}),
            },{reference: () =>h(ArtButtonTable, {
              type: 'view',
              icon: 'solar:checklist-minimalistic-bold'
            })})),

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



const handleApply = async (data:{id:number,status:number}) => {
    // TODO: 调用删除接口
  await fetchPostCommentApply(data)
  refreshData()
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
      await fetchPostCommentDelete({ids:selectedRows.value})
      refreshData()
    })
    .catch(() => {
      ElMessage.info('已取消删除')
    })
  }
}


const handleDelete = async (row: Comment.Response.Info): Promise<void> => {
  ElMessageBox.confirm(`确定要删除该吗？`, '删除提现', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'error'
  }).then(async() => {
    // TODO: 调用删除接口
    ElMessage.success('删除成功')
   await fetchPostCommentDelete({ids:[row.id]})
    refreshData()
  })
  .catch(() => {
    ElMessage.info('已取消删除')
  })
}


/**
 * 处理表格行选择变化
 */
const handleSelectionChange = (selection: Comment.Response.Info[]): void => {
  selectedRows.value = selection.map((item)=>item.id)
}
</script>
