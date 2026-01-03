<template>
    <ElDialog
      title="派单威客"
      width="50%"
      :model-value="visible"
      align-center
      :close-on-click-modal="false"
      @update:model-value="handleCancel"
      @close="handleClose"
    >
        <OrderDistributeSearch class="mb-4"
         v-model="searchForm" @search="handleSearch" @reset="resetSearchParams"/>
        <!-- 表格头部 -->
        <ArtTableHeader v-model:columns="columnChecks" :loading="loading" @refresh="refreshData"></ArtTableHeader>
        <!-- 表格 -->
        <ArtTable
            :height="300"
            :loading="loading"
            :data="data"
            :columns="columns"
            :pagination="pagination"
            @selection-change="handleSelectionChange"
            @pagination:size-change="handleSizeChange"
            @pagination:current-change="handleCurrentChange"
        >
        </ArtTable>
    </ElDialog>
  </template>
  
<script setup lang="ts">
import {  fetchPostOrderDistribute } from '@/api/order';
import { fetchGetWitkeyList } from '@/api/witkey';
import { useTable } from '@/hooks';
import { ElButton, ElImage, ElMessageBox, ElRate, ElTag } from 'element-plus'
import OrderDistributeSearch from './order-distribute-search.vue';
import { DistributeType } from '@/enums/typeEnum';

interface Props {
    visible: boolean
    id?: number | null
}

interface Emits {
    (e: 'update:visible', value: boolean): void
    (e: 'submit'): void
}

const props = withDefaults(defineProps<Props>(), {
    visible: false,
    id: 0,
})

const emit = defineEmits<Emits>()

/**
 * 弹窗显示状态双向绑定
 */
const visible = computed({
    get: () => props.visible,
    set: (value) => emit('update:visible', value)
})


// 选中行
const selectedRows = ref<number[]>([])
// 搜索表单
const searchForm = ref({
    name: undefined,
    phone: undefined
})

const {
    columns,
    columnChecks,
    data,
    loading,
    pagination,
    fetchData,
    getData,
    searchParams,
    resetSearchParams,
    handleSizeChange,
    handleCurrentChange,
    refreshData
} = useTable({
    // 核心配置
    core: {
        immediate:false,
        apiFn: fetchGetWitkeyList,
        apiParams:{
            phone: "",
            status: 0,
        },
        paginationKey:{
            current: 'page', 
            size: 'limit'
        },
        columnsFactory: () => [
            {
                prop: 'userInfo',
                label: '所属用户',
                width: 240,
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
                    disabled: true,
                })
                }
            },
            {
                prop: 'operation',
                label: '操作',
                width: 200,
                fixed: 'right', // 固定列
                formatter: (row) =>{
                    return h('div', { class: 'order flex-c' }, [
                        h(ElButton, {
                            type: "primary",
                            plain: true,
                            size: 'small',
                            onClick: () => handleSubmit(row.id,DistributeType.Team)
                        },()=>"自带队伍",),
                        h(ElButton, {
                            type: "primary",
                            plain: true,
                            size: 'small',
                            onClick: () => handleSubmit(row.id,DistributeType.Self)
                        },()=>"个人接单",),
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
 * 处理表格行选择变化
 */
const handleSelectionChange = (selection: Order.Response.Info[]): void => {
    selectedRows.value = selection.map((item)=>item.id)
}
/**
 * 监听弹窗打开，初始化表单数据
 */
watch(
    () => props.visible,
    (newVal) => {
    if (newVal) init()
    }
)

/**
 * 初始化表单数据
 * 根据弹窗类型填充表单或重置表单
 */
const init = async () => {
    fetchData()
}

/**
 * 关闭弹窗并重置表单
 */
const handleClose = () => {
    emit('update:visible', false)
}

/**
 * 取消操作
 */
const handleCancel = (): void => {
    handleClose()
}

/**
 * 提交表单
 * 验证通过后调用接口保存数据
 */
const handleSubmit = async (id:number,type:number) => {
    if (type == DistributeType.Self) {
        ElMessageBox.confirm(`确认派发给他个人吗`, '提示', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'primary'
        }).then(async() => {
            // TODO: 调用删除接口
            await fetchPostOrderDistribute({
                id:props.id!,
                witkeyId: id,
                type:DistributeType.Self
            })
            ElMessage.success('派单成功')
            emit('submit')
            handleCancel()
        })
        .catch(() => {
            ElMessage.info('已取消')
        })
   }
   if (type == DistributeType.Team) {
        ElMessageBox.confirm(`确认派发给他队伍吗`, '提示', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'primary'
        }).then(async() => {
            // TODO: 调用删除接口
            await fetchPostOrderDistribute({
                id:props.id!,
                witkeyId: id,
                type:DistributeType.Team
            })
            ElMessage.success('派单成功')
            emit('submit')
            handleCancel()
        })
        .catch(() => {
            ElMessage.info('已取消')
        })
   }

}
</script>