<!-- 角色管理页面 -->
<template>
  <div class="art-full-height">
    <RoleSearch
      v-show="showSearchBar"
      v-model="searchForm"
      @search="handleSearch"
      @reset="resetSearchParams"
    ></RoleSearch>

    <ElCard
      class="art-table-card"
      shadow="never"
      :style="{ 'margin-top': showSearchBar ? '12px' : '0' }"
    >
      <ArtTableHeader
        v-model:columns="columnChecks"
        v-model:showSearchBar="showSearchBar"
        :loading="loading"
        @refresh="refreshData"
      >
        <template #left>
          <ElSpace wrap>
            <ElButton @click="handleCreate" v-ripple>新增角色</ElButton>
          </ElSpace>
        </template>
      </ArtTableHeader>

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
    </ElCard>

    <!-- 角色编辑弹窗 -->
    <RoleDialog
      v-model="dialogVisible"
      :dialog-type="dialogType"
      :id="id"
      @submit="refreshData"
    />

    <!-- 菜单权限弹窗 -->
    <RoleMenuDialog
      v-model="menuDialog"
      :id="id"
      @submit="refreshData"
    />
    <RolePermissionDialog
      v-model="permissionDialog"
     :id="id"
      @submit="refreshData"
    />
  </div>
</template>

<script setup lang="ts">
import { ButtonMoreItem } from '@/components/core/forms/art-button-more/index.vue'
import { useTable } from '@/hooks/core/useTable'

import ArtButtonMore from '@/components/core/forms/art-button-more/index.vue'
import RoleSearch from './modules/role-search.vue'
import RoleDialog from './modules/role-dialog.vue'
import RolePermissionDialog from './modules/role-permission-dialog.vue'
import { ElTag, ElMessageBox } from 'element-plus'
import { fetchRoleDelete, fetchRoleList } from '@/api/role'
import RoleMenuDialog from './modules/role-menu-dialog.vue'

  defineOptions({ name: 'Role' })


  // 搜索表单
  const searchForm = ref({
    name: undefined,
  })

  const showSearchBar = ref(true)

  const dialogVisible = ref(false)
  const permissionDialog = ref(false)
  const menuDialog = ref(false)
  const id = ref<number | null>(null)

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
      apiFn: fetchRoleList,
      apiParams:{
        name: "",
      },
      paginationKey:{
        current: 'page', 
        size: 'limit'
      },
      // 排除 apiParams 中的属性
      excludeParams: ['daterange'],
      columnsFactory: () => [
        {
          prop: 'id',
          label: '角色ID',
          width: 100
        },
        {
          prop: 'name',
          label: '角色名称',
          minWidth: 120
        },
        {
          prop: 'code',
          label: '角色编码',
          minWidth: 120
        },
        {
          prop: 'description',
          label: '角色描述',
          minWidth: 150,
          showOverflowTooltip: true
        },
        {
          prop: 'status',
          label: '角色状态',
          width: 100,
          formatter: (row) => {

            const tagMap: Record<number, string> = {
              1: 'warning',
              2: 'success'
            };

            const label: Record<number, string> = {
              1: '禁用',
              2: '启用'
            };
            return h(
              ElTag,
              { type: tagMap[row.status] as 'success' | 'warning' },
              () => label[row.status]
            )
          }
        },
        {
          prop: 'createTime',
          label: '创建日期',
          width: 180,
          sortable: true
        },
        {
          prop: 'operation',
          label: '操作',
          width: 80,
          fixed: 'right',
          formatter: (row) =>
            h('div', [
              h(ArtButtonMore, {
                list: [
                  {
                    key: 'setMenu',
                    label: '设置菜单',
                    icon: 'ep:element-plus',
                    auth:'setMenu'
                  },
                  {
                    key: 'setPermission',
                    label: '设置权限',
                    icon: 'ep:pear',
                    auth:'setPermission'
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
      ]
    },
    transform: {
      responseAdapter: (response) => {
        return {
          records: response.list,
          total: response.total,
        };
      },
    },
  })

  const dialogType = ref<'create' | 'edit'>('create')

  const handleCreate = () => {
    dialogVisible.value = true
    dialogType.value =  'create'
    id.value = 0
    // currentRoleData.value = row
  }

  const handleEdit = (row:Role.Response.Info) => {
    dialogVisible.value = true
    dialogType.value =  'edit'
    id.value = row.id
    // currentRoleData.value = row
  }

  /**
   * 搜索处理
   * @param params 搜索参数
   */
  const handleSearch = (params: Record<string, any>) => {
    Object.assign(searchParams, { ...params })
    getData()
  }

  const buttonMoreClick = (item: ButtonMoreItem, row: Role.Response.Info) => {
    switch (item.key) {
      case 'setMenu':
        showMenuDialog(row)
        break
      case 'setPermission':
        showPermissionDialog(row)
        break
      case 'edit':
        handleEdit(row)
        break
      case 'delete':
        handleDelete(row)
        break
    }
  }

  const showMenuDialog = (row?: Role.Response.Info) => {
    menuDialog.value = true
    id.value = row?.id!
  }

  const showPermissionDialog = (row?: Role.Response.Info) => {
    permissionDialog.value = true
    id.value = row?.id!
  }

  const handleDelete = (row: Role.Response.Info) => {
    ElMessageBox.confirm(`确定删除角色"${row.name}"吗？此操作不可恢复！`, '删除确认', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
      .then(async() => {
        // TODO: 调用删除接口
        ElMessage.success('删除成功')
        await fetchRoleDelete({ids:[row.id]})
        refreshData()
      })
      .catch(() => {
        ElMessage.info('已取消删除')
      })
  }
</script>
