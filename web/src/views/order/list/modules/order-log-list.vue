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
    </div>
</template>

<script setup lang="ts">
import {  fetchGetOrderLogList } from '@/api/order';
import { useAuth, useTable } from '@/hooks';
import ArtButtonTable from '@/components/core/forms/art-button-table/index.vue'
import { ElTag } from 'element-plus';
import { OrderLogType } from '../../../../enums/typeEnum';
interface Props {
  id?: number;
}
const props = defineProps<Props>();

const TYPE = {
  [OrderLogType.Create]: { type: 'primary' as const, text: '创建订单' },
  [OrderLogType.AddDiscount]: { type: 'primary' as const, text: '设置优惠金额' },
  [OrderLogType.Cancel]: { type: 'primary' as const, text: '关闭订单' },
  [OrderLogType.Complete]: { type: 'primary' as const, text: '完成订单' },
  [OrderLogType.Paid]: { type: 'primary' as const, text: '确认收款订单' },
  [OrderLogType.AfterSales]: { type: 'primary' as const, text: '添加售后工单' },
  [OrderLogType.Distribute]: { type: 'primary' as const, text: '派发威客' },
  [OrderLogType.DistributeCancel]: { type: 'primary' as const, text: '取消派单服务' },
  [OrderLogType.Start]: { type: 'primary' as const, text: '开始服务' },
  [OrderLogType.Refund]: { type: 'primary' as const, text: '订单手动退款' },
} as const

const getType = (type: number) => {
  return (
    TYPE[type as keyof typeof TYPE] || {
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
    handleSizeChange,
    handleCurrentChange,
} = useTable({
    // 核心配置
    core: {
        apiFn: fetchGetOrderLogList,
        apiParams:{
            id: props.id,
        },
        paginationKey:{
            current: 'page', 
            size: 'limit'
        },
        columnsFactory: () => [
            {
                prop: 'manage',
                label: '操作人',
                width: 160,
                formatter: (row) => {
                return h('p', { }, row.manage)
                }
            },
            {
                prop: 'type',
                label: '操作类型',
                formatter: (row) => {
                    const type = getType(row.type)
                    return h(ElTag, { type:"primary" }, () => type.text )
                }
            },
            {
                prop: 'createTime',
                label: '派单时间',
                sortable: true
            },
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


</script>