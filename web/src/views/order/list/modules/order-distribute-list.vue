<template>
    <div class="art-full-height">
    <!-- 表格 -->
        <ArtTable
            :loading="loading"
            :data="data"
            :columns="columns"
            :pagination="pagination"
            @pagination:size-change="handleSizeChange"
            @pagination:current-change="handleCurrentChange"
        >
        </ArtTable>
        <OrderDistributeCancelModal 
            v-model:visible="cancelModalVisible"
            :id="distributeId"
            @submit="refreshData"
        />
    </div>
</template>

<script setup lang="ts">
import { fetchGetOrderDistributeList } from '@/api/order';
import { useAuth, useTable } from '@/hooks';
import ArtButtonTable from '@/components/core/forms/art-button-table/index.vue'
import { ElTag } from 'element-plus';
import OrderDistributeCancelModal from './order-distribute-cancel-modal.vue';
interface Props {
  id?: number;
}
const props = defineProps<Props>();
const { hasAuth } = useAuth();
const IS_CANCEL = {
  1: { type: 'primary' as const, text: '未取消' },
  2: { type: 'danger' as const, text: '取消' },
} as const

const getIsCancel = (isCancel: number) => {
  return (
    IS_CANCEL[isCancel as keyof typeof IS_CANCEL] || {
      type: 'info' as const,
      text: '未知'
    }
  )
}
const {
    columns,
    data,
    loading,
    pagination,
    refreshData,
    handleSizeChange,
    handleCurrentChange,
} = useTable({
    // 核心配置
    core: {
        apiFn: fetchGetOrderDistributeList,
        apiParams:{
            id: props.id,
            name: "",
        },
        paginationKey:{
            current: 'page', 
            size: 'limit'
        },
        columnsFactory: () => [
            {
                prop: 'witkey',
                label: '接单者',
                width: 160,
                formatter: (row) => {
                return h('p', { }, row.witkey)
                }
            },
            {
                prop: 'game',
                label: '游戏领域',
                formatter: (row) => {
                    return h(ElTag, { type:"primary" }, () => row.game )
                }
            },
            {
                prop: 'title',
                label: '所属头衔',
                formatter: (row) => {
                    return h(ElTag, { type:"primary" }, () => row.title )
                }
            },
            {
                prop: 'isCancel',
                label: '是否取消',
                formatter: (row) => {
                const statusConfig = getIsCancel(row.isCancel)
                return h(ElTag, { type: statusConfig.type }, () => statusConfig.text)
                }
            },
            {
                prop: 'createTime',
                label: '派单时间',
                sortable: true
            },
            {
                prop: 'operation',
                label: '操作',
                width: 120,
                fixed: 'right', // 固定列
                formatter: (row) =>{
             
                return h('div', { class: 'order flex-c' }, [
                        (hasAuth("distributeCancel") && h(ArtButtonTable, {
                            type: 'delete',
                            icon:'solar:close-circle-bold',
                            onClick: () => handleCancel(row.id)
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

const distributeId = ref(0)
const cancelModalVisible = ref(false)
const handleCancel = (id:number) => {
    distributeId.value = id
    nextTick(() => {
        cancelModalVisible.value = true
    })
}

</script>