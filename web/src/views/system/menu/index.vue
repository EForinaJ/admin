<!-- 菜单管理页面 -->
<template>
  <div class="menu-page art-full-height">
    <!-- 搜索栏 -->
    <ArtSearchBar
      v-model="formFilters"
      :items="formItems"
      :showExpand="false"
      @reset="handleReset"
      @search="handleSearch"
    />

    <ElCard class="art-table-card" shadow="never">
      <!-- 表格头部 -->
      <ArtTableHeader
        :showZebra="false"
        :loading="loading"
        v-model:columns="columnChecks"
        @refresh="handleRefresh"
      >
        <template #left>
          <ElButton v-auth="'create'" @click="handleCreate" v-ripple> 添加菜单 </ElButton>
          <ElButton @click="toggleExpand" v-ripple>
            {{ isExpanded ? '收起' : '展开' }}
          </ElButton>
        </template>
      </ArtTableHeader>

      <ArtTable
        ref="tableRef"
        rowKey="path"
        :loading="loading"
        :columns="columns"
        :data="filteredTableData"
        :stripe="false"
        :tree-props="{ children: 'children', hasChildren: 'hasChildren' }"
        :default-expand-all="false"
      />

      <!-- 菜单弹窗 -->
      <MenuDialog
        v-model:visible="dialogVisible"
        v-model:is-child="isChild"
        :type="dialogType"
        :id="id"
        :lockType="lockMenuType"
        @submit="handleSubmit"
      />
    </ElCard>
   
  </div>
</template>

<script setup lang="ts">
import { formatMenuTitle } from '@/utils/router'
import ArtButtonTable from '@/components/core/forms/art-button-table/index.vue'
import ArtSvgIcon from "@/components/core/base/art-svg-icon/index.vue";
import { useTableColumns } from '@/hooks/core/useTableColumns'
import type { AppRouteRecord } from '@/types/router'
import MenuDialog from './modules/menu-dialog.vue'
import { ElTag, ElMessageBox } from 'element-plus'
import { fetchMenuDelete, fetchMenuAll } from '@/api/menu'
import { useAuth } from '@/hooks';

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
const lockMenuType = ref(false)

// 搜索相关
const initialSearchState = {
  name: '',
  route: ''
}

const formFilters = reactive({ ...initialSearchState })
const appliedFilters = reactive({ ...initialSearchState })

const formItems = computed(() => [
  {
    label: '菜单名称',
    key: 'name',
    type: 'input',
    props: { clearable: true }
  },
  {
    label: '路由地址',
    key: 'route',
    type: 'input',
    props: { clearable: true }
  }
])

onMounted(() => {
  getMenuList()
})

/**
 * 获取菜单列表数据
 */
const getMenuList = async (): Promise<void> => {
  loading.value = true

  try {
    const res = await fetchMenuAll()
    tableData.value = convertMenuToRouteTree(res.list)
  } catch (error) {
    throw error instanceof Error ? error : new Error('获取菜单失败')
  } finally {
    loading.value = false
  }
}


const convertMenuToRouteTree = (menuList: Menu.Response.List[]): AppRouteRecord[] => {
  const routeMap = new Map<number, AppRouteRecord>();
  const tree: AppRouteRecord[] = [];
  
  // 第一步：先创建所有非按钮类型的路由映射
  menuList.forEach(menu => {
    if (menu.menuType === "button") return;
    
    const route: AppRouteRecord = {
      id: menu.id,
      name: menu.label as string,
      path: menu.path,
      meta: {
        title: menu.name,
        icon: menu.icon,
        showBadge: menu.showBadge,
        showTextBadge: menu.showTextBadge,
        isHide: menu.isHide,
        isHideTab: menu.isHideTab,
        link: menu.link,
        isIframe: menu.isIframe,
        keepAlive: menu.keepAlive,
        fixedTab: menu.fixedTab,
        activePath: menu.activePath,
        isFullPage: menu.isFullPage,
        authMark: menu.authName,
        isAuthButton: false,
        isFirstLevel: menu.pId === 0,
        authList: [] // 初始化为空数组
      },
      component: menu.component ? menu.component : undefined
    };
    
    routeMap.set(menu.id, route);
  });
  
  // 第二步：处理按钮类型的菜单，添加到对应父节点的authList
  menuList.forEach(menu => {
    if (menu.menuType === "button") {
      const parentRoute = routeMap.get(menu.pId);
      if (parentRoute && parentRoute.meta.authList) {
        parentRoute.meta.authList.push({
          id: menu.id,
          title: menu.authName,
          authMark: menu.authLabel
        });
      }
    }
  });
  
  // 第三步：构建树形结构
  menuList.forEach(menu => {
    if (menu.menuType === "button") return;
    
    const route = routeMap.get(menu.id)!;
    
    if (menu.pId === 0) {
      tree.push(route);
    } else {
      const parent = routeMap.get(menu.pId);
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
/**
 * 获取菜单类型标签颜色
 * @param row 菜单行数据
 * @returns 标签颜色类型
 */
const getMenuTypeTag = (
  row: AppRouteRecord
): 'primary' | 'success' | 'warning' | 'info' | 'danger' => {
  if (row.meta?.isAuthButton) return 'danger'
  if (row.children?.length) return 'info'
  if (row.meta?.link && row.meta?.isIframe) return 'success'
  if (row.path) return 'primary'
  if (row.meta?.link) return 'warning'
  return 'info'
}

/**
 * 获取菜单类型文本
 * @param row 菜单行数据
 * @returns 菜单类型文本
 */
const getMenuTypeText = (row: AppRouteRecord): string => {
  if (row.meta?.isAuthButton) return '按钮'
  if (row.children?.length) return '目录'
  if (row.meta?.link && row.meta?.isIframe) return '内嵌'
  if (row.path) return '菜单'
  if (row.meta?.link) return '外链'
  return '未知'
}

// 表格列配置
const { columnChecks, columns } = useTableColumns(() => [
{
    prop: 'meta.icon',
    label: '图标',
    formatter: (row: AppRouteRecord) => {
      return h(ArtSvgIcon, { icon: row.meta.icon })
    }
  },
  {
    prop: 'meta.title',
    label: '菜单名称',
    minWidth: 120,
    formatter: (row: AppRouteRecord) => formatMenuTitle(row.meta?.title)
  },
  {
    prop: 'type',
    label: '菜单类型',
    formatter: (row: AppRouteRecord) => {
      return h(ElTag, { type: getMenuTypeTag(row) }, () => getMenuTypeText(row))
    }
  },
  {
    prop: 'path',
    label: '路由',
    formatter: (row: AppRouteRecord) => {
      if (row.meta?.isAuthButton) return ''
      return row.meta?.link || row.path || ''
    }
  },
  {
    prop: 'meta.authList',
    label: '权限标识',
    formatter: (row: AppRouteRecord) => {
      if (row.meta?.isAuthButton) {
        return row.meta?.authMark || ''
      }
      if (!row.meta?.authList?.length) return ''
      return `${row.meta.authList.length} 个权限标识`
    }
  },
  {
    prop: 'status',
    label: '状态',
    formatter: () => h(ElTag, { type: 'success' }, () => '启用')
  },
  {
    prop: 'operation',
    label: '操作',
    width: 180,
    align: 'right',
    formatter: (row: AppRouteRecord) => {
      const buttonStyle = { style: 'text-align: right' }
      
      if (row.meta?.isAuthButton) {
        return h('div', buttonStyle, [
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
const tableData = ref<AppRouteRecord[]>([])

/**
 * 重置搜索条件
 */
const handleReset = (): void => {
  Object.assign(formFilters, { ...initialSearchState })
  Object.assign(appliedFilters, { ...initialSearchState })
  getMenuList()
}

/**
 * 执行搜索
 */
const handleSearch = (): void => {
  Object.assign(appliedFilters, { ...formFilters })
  getMenuList()
}

/**
 * 刷新菜单列表
 */
const handleRefresh = (): void => {
  getMenuList()
}

/**
 * 深度克隆对象
 * @param obj 要克隆的对象
 * @returns 克隆后的对象
 */
const deepClone = <T,>(obj: T): T => {
  if (obj === null || typeof obj !== 'object') return obj
  if (obj instanceof Date) return new Date(obj) as T
  if (Array.isArray(obj)) return obj.map((item) => deepClone(item)) as T

  const cloned = {} as T
  for (const key in obj) {
    if (Object.prototype.hasOwnProperty.call(obj, key)) {
      cloned[key] = deepClone(obj[key])
    }
  }
  return cloned
}

/**
 * 将权限列表转换为子节点
 * @param items 菜单项数组
 * @returns 转换后的菜单项数组
 */
const convertAuthListToChildren = (items: AppRouteRecord[]): AppRouteRecord[] => {
  return items.map((item) => {
    const clonedItem = deepClone(item)

    if (clonedItem.children?.length) {
      clonedItem.children = convertAuthListToChildren(clonedItem.children)
    }

    if (item.meta?.authList?.length) {
      const authChildren: AppRouteRecord[] = item.meta.authList.map(
        (auth: { title: string; authMark: string; id: number }) => ({
          id: auth.id,
          path: `${item.path}_auth_${auth.authMark}`,
          name: `${String(item.name)}_auth_${auth.authMark}`,
          meta: {
            title: auth.title,
            authMark: auth.authMark,
            isAuthButton: true,
            parentPath: item.path
          }
        })
      )

      clonedItem.children = clonedItem.children?.length
        ? [...clonedItem.children, ...authChildren]
        : authChildren
    }

    return clonedItem
  })
}

/**
 * 搜索菜单
 * @param items 菜单项数组
 * @returns 搜索结果数组
 */
const searchMenu = (items: AppRouteRecord[]): AppRouteRecord[] => {
  const results: AppRouteRecord[] = []

  for (const item of items) {
    const searchName = appliedFilters.name?.toLowerCase().trim() || ''
    const searchRoute = appliedFilters.route?.toLowerCase().trim() || ''
    const menuTitle = formatMenuTitle(item.meta?.title || '').toLowerCase()
    const menuPath = (item.path || '').toLowerCase()
    const nameMatch = !searchName || menuTitle.includes(searchName)
    const routeMatch = !searchRoute || menuPath.includes(searchRoute)

    if (item.children?.length) {
      const matchedChildren = searchMenu(item.children)
      if (matchedChildren.length > 0) {
        const clonedItem = deepClone(item)
        clonedItem.children = matchedChildren
        results.push(clonedItem)
        continue
      }
    }

    if (nameMatch && routeMatch) {
      results.push(deepClone(item))
    }
  }

  return results
}

// 过滤后的表格数据
const filteredTableData = computed(() => {
  const searchedData = searchMenu(tableData.value)
  return convertAuthListToChildren(searchedData)
})

/**
 * 添加菜单
 */
const handleCreate = (): void => {
  dialogType.value = 'create'
  // editData.value = null
  id.value = 0
  lockMenuType.value = true
  dialogVisible.value = true
}


const handleChildren = (row: AppRouteRecord):  void =>{
  dialogType.value = 'create'
  lockMenuType.value = false
  dialogVisible.value = true
  isChild.value = true
  id.value = row.id!
}


/**
 * 编辑菜单
 * @param row 菜单行数据
 */
const handleEdit = (row: AppRouteRecord): void => {
  console.log(row)
  dialogType.value = 'edit'
  id.value = row.id!
  lockMenuType.value = true
  dialogVisible.value = true
}


/**
 * 菜单表单数据类型
 */
interface MenuFormData {
  name: string
  path: string
  component?: string
  icon?: string
  roles?: string[]
  sort?: number
  [key: string]: any
}

/**
 * 提交表单数据
 * @param formData 表单数据
 */
const handleSubmit = (formData: MenuFormData): void => {
  console.log('提交数据:', formData)
  // TODO: 调用API保存数据
  getMenuList()
}

/**
 * 删除菜单
 */
const handleDelete = async (row: AppRouteRecord): Promise<void> => {
  try {
    await ElMessageBox.confirm('确定要删除该菜单吗？删除后无法恢复', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    await fetchMenuDelete({ids:[row.id!]})
    ElMessage.success('删除成功')
    getMenuList()
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error('删除失败')
    }
  }
}

/**
 * 切换展开/收起所有菜单
 */
const toggleExpand = (): void => {
  isExpanded.value = !isExpanded.value
  nextTick(() => {
    if (tableRef.value?.elTableRef && filteredTableData.value) {
      const processRows = (rows: AppRouteRecord[]) => {
        rows.forEach((row) => {
          if (row.children?.length) {
            tableRef.value.elTableRef.toggleRowExpansion(row, isExpanded.value)
            processRows(row.children)
          }
        })
      }
      processRows(filteredTableData.value)
    }
  })
}
</script>
