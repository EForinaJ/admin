<!-- 权限管理页面 -->
<template>
  <div class="menu-page art-full-height">

    <ElCard class="art-table-card" shadow="never">
      <!-- 表格头部 -->
      <ArtTableHeader
        :showZebra="false"
        :loading="loading"
        v-model:columns="columnChecks"
        @refresh="handleRefresh"
      >
        <template #left>
          <ElButton v-auth="'create'" @click="handleCreate" v-ripple> 添加权限 </ElButton>
        </template>
      </ArtTableHeader>

      <ArtTable
        ref="tableRef"
        rowKey="id"
        :loading="loading"
        :columns="columns"
        :data="tableData"
        :stripe="false"
        :tree-props="{ children: 'children', hasChildren: 'hasChildren' }"
        :default-expand-all="false"
      />

      <!-- 权限弹窗 -->
      <PermissionDialog
        v-model:visible="dialogVisible"
        v-model:is-child="isChild"
        :dialog-type="dialogType"
        :id="id"
        @submit="handleSubmit"
      />
    </ElCard>
   
  </div>
</template>

<script setup lang="ts">
import ArtButtonTable from '@/components/core/forms/art-button-table/index.vue'
import { useTableColumns } from '@/hooks/core/useTableColumns'
import {ElMessageBox } from 'element-plus'
import { useAuth } from '@/hooks';
import { fetchPermissionAll, fetchPermissionDelete } from '@/api/permission';
import PermissionDialog from './modules/permission-dialog.vue';

const { hasAuth } = useAuth();
defineOptions({ name: 'Menus' })

// 状态管理
const loading = ref(false)
const isExpanded = ref(false)
const tableRef = ref()

// 弹窗相关
const dialogVisible = ref(false)
const isChild = ref(false)
const dialogType = ref<'create' | 'edit'>('create')
const id = ref<number | null>(null)


// 搜索相关
const initialSearchState = {
  name: '',
  route: ''
}

const formFilters = reactive({ ...initialSearchState })
const appliedFilters = reactive({ ...initialSearchState })

onMounted(() => {
  getList()
})

/**
 * 获取权限列表数据
 */
const getList = async (): Promise<void> => {
  loading.value = true
  try {
    const res = await fetchPermissionAll()
    tableData.value = convertMenuToRouteTree(res.list)
  } catch (error) {
    throw error instanceof Error ? error : new Error('获取权限失败')
  } finally {
    loading.value = false
  }
}

/**
 * 应用路由记录接口
 * 扩展 Vue Router 的路由记录类型
 */
interface PermissionTree extends Omit<Permission.Response.Info,  'children'> {
  children?: PermissionTree[]
}

const convertMenuToRouteTree = (list: Permission.Response.Info[]): PermissionTree[] => {
  const routeMap = new Map<number, PermissionTree>();
  const tree: PermissionTree[] = [];
  
  // 第一步：先创建所有非按钮类型的路由映射
  list.forEach(item => {
    routeMap.set(item.id, item);
  });
  
  
  // 第三步：构建树形结构
  list.forEach(item => {
    const route = routeMap.get(item.id)!;
    
    if (item.pId === 0) {
      tree.push(route);
    } else {
      const parent = routeMap.get(item.pId);
      if (parent) {
        if (!parent.children) {
          parent.children = [];
        }
        parent.children.push(route);
      } else {
        tree.push(route);
      }
    }
  });
  
  return tree;
};


// 表格列配置
const { columnChecks, columns } = useTableColumns(() => [
{
    prop: 'id',
    label: '权限ID',
    formatter: (row: PermissionTree) => {
      return row.id
    }
  },
  {
    prop: 'name',
    label: '权限名称',
    minWidth: 120,
    formatter: (row: PermissionTree) => {
      return row.name
    }
  },
  {
    prop: 'permission',
    label: '权限标识',
    formatter: (row: PermissionTree) => {
      return row.permission
    }
  },
  {
    prop: 'createTime',
    label: '创建时间',
    formatter: (row: PermissionTree) => {
      return row.createTime
    }
  },
  {
    prop: 'operation',
    label: '操作',
    width: 180,
    align: 'right',
    formatter: (row: PermissionTree) => {
      const buttonStyle = { style: 'text-align: right' }
      
      return h('div', buttonStyle, [
        (hasAuth("create") && h(ArtButtonTable, {
          type: 'add',
          onClick: () => handleChildren(row),
        })),
        (hasAuth("edit") &&  h(ArtButtonTable, {
          type: 'edit',
          onClick: () => handleEdit(row)
        })),
        (hasAuth("delete") &&  h(ArtButtonTable, {
          type: 'delete',
          onClick: () => handleDelete(row)
        })),
      ])
    }
  }
])

// 数据相关
const tableData = ref<PermissionTree[]>([])

/**
 * 重置搜索条件
 */
const handleReset = (): void => {
  Object.assign(formFilters, { ...initialSearchState })
  Object.assign(appliedFilters, { ...initialSearchState })
  getList()
}

/**
 * 执行搜索
 */
const handleSearch = (): void => {
  getList()
}

/**
 * 刷新权限列表
 */
const handleRefresh = (): void => {
  getList()
}

/**
 * 添加权限
 */
const handleCreate = (): void => {
  dialogType.value = 'create'
  id.value = 0
  dialogVisible.value = true
}


const handleChildren = (row: PermissionTree):  void =>{
  dialogType.value = 'create'
  dialogVisible.value = true
  isChild.value = true
  id.value = row.id!
}


/**
 * 编辑权限
 * @param row 权限行数据
 */
const handleEdit = (row: PermissionTree): void => {
  dialogType.value = 'edit'
  id.value = row.id!
  dialogVisible.value = true
}


/**
 * 提交表单数据
 * @param formData 表单数据
 */
const handleSubmit = (): void => {
  // TODO: 调用API保存数据
  getList()
}

/**
 * 删除权限
 */
const handleDelete = async (row: PermissionTree): Promise<void> => {
  try {
    await ElMessageBox.confirm('确定要删除该权限吗？删除后无法恢复', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    await fetchPermissionDelete({ids:[row.id!]})
    ElMessage.success('删除成功')
    getList()
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error('删除失败')
    }
  }
}

</script>
