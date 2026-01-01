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
import { fetchGetWitkeyCommissionList } from '@/api/witkey';
import { FundMode } from '@/enums/modeEnum';
import { ChangeCommissionType } from '@/enums/typeEnum';
import { useTable } from '@/hooks';
import { useSiteStore } from '@/store/modules/site';
import { ElTag } from 'element-plus';
interface Props {
  id?: number;
}
const props = defineProps<Props>();

const {getInfo:site} = useSiteStore()

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
        apiFn: fetchGetWitkeyCommissionList,
        apiParams:{
            id: props.id,
        },
        paginationKey:{
            current: 'page', 
            size: 'limit'
        },
        columnsFactory: () => [
            
            {
                prop: 'type',
                label: '变更类型',
                formatter: (row) => {
                    const label: Record<number, string> = {
                        [ChangeCommissionType.WitkeyChangeBalanceTypeSystem]: '系统变更佣金',
                        [ChangeCommissionType.WitkeyChangeBalanceTypeSettlement]: '报单结算',
                        [ChangeCommissionType.WitkeyChangeBalanceTypeWithdraw]: '佣金提现',
                    };
                    return h(ElTag, { 
                        type: "primary"
                    }, {
                        default: () => label[row.type]
                    })
                }
            },
            {
                prop: 'commissionInfo',
                label: '变动资金',
                formatter: (row) => {
                    if (row.mode == FundMode.Add) {
                        return  h('div', { class: 'witkey flex-c gap-1' }, [
                            h('span', { class: 'text-primary text-sm' },"+" + row.amount + site.symbol!),
                        ])
                    }
                    if (row.mode == FundMode.Sub) {
                        return  h('div', { class: 'witkey flex-c gap-2' }, [
                            h('span', { class: 'text-error text-sm' }, "-" +row.amount + site.symbol!),
                        ])
                    }
                }
            },
            {
                prop: 'commission',
                label: '当前余额',
                formatter: (row) => {
                    return  h('div', { class: 'witkey flex-c gap-2' }, [
                        h('span', { class: 'text-sm' }, row.after + site.symbol!),
                    ])
                }
            },
            {
                prop: 'remark',
                label: '变更备注',
            },
            {
                prop: 'createTime',
                label: '变动时间',
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