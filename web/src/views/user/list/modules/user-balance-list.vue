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
import { fetchGetUserBalanceList } from '@/api/user';
import { FundMode } from '@/enums/modeEnum';
import { ChangeBalanceType } from '@/enums/typeEnum';
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
        apiFn: fetchGetUserBalanceList,
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
                        [ChangeBalanceType.UserChangeBalanceTypeSystem]: '系统变更余额',
                        [ChangeBalanceType.UserChangeBalanceTypeRecharge]: '用户充值',
                        [ChangeBalanceType.UserChangeBalanceTypePaidOrder]: '订单消费',
                    };
                    return h(ElTag, { 
                        type: "primary"
                    }, {
                        default: () => label[row.type]
                    })
                }
            },
            {
                prop: 'balanceInfo',
                label: '变动资金',
                formatter: (row) => {
                    if (row.mode == FundMode.Add) {
                        return  h('div', { class: 'user flex-c gap-1' }, [
                            h('span', { class: 'text-primary text-sm' },"+" + row.amount + site.symbol!),
                        ])
                    }
                    if (row.mode == FundMode.Sub) {
                        return  h('div', { class: 'user flex-c gap-2' }, [
                            h('span', { class: 'text-error text-sm' }, "-" +row.amount + site.symbol!),
                        ])
                    }
                }
            },
            {
                prop: 'balance',
                label: '当前余额',
                formatter: (row) => {
                    return  h('div', { class: 'user flex-c gap-2' }, [
                        h('span', { class: 'text-sm' }, row.before + site.symbol!),
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