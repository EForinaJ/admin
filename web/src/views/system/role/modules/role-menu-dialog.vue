<template>
  <ElDialog
    v-model="visible"
    title="菜单权限"
    width="520px"
    align-center
    class="el-dialog-border"
    @close="handleClose"
  >
    <ElScrollbar height="70vh">
      <ElTree
        ref="treeRef"
        :data="processedMenuList"
        show-checkbox
        node-key="id"
        :default-expand-all="isExpandAll"
        :default-checked-keys="[1, 2, 3]"
        :props="defaultProps"
        @check="handleTreeCheck"
      >
        <template #default="{ data }">
          <div style="display: flex; align-items: center">
            <span v-if="data.isAuth">
              {{ data.label }}
            </span>
            <span v-else>{{ defaultProps.label(data) }}</span>
          </div>
        </template>
      </ElTree>
    </ElScrollbar>
    <template #footer>
      <!-- <ElButton @click="outputSelectedData" style="margin-left: 8px">获取选中数据</ElButton> -->
      <ElButton @click="toggleExpandAll">{{ isExpandAll ? '全部收起' : '全部展开' }}</ElButton>
      <ElButton @click="toggleSelectAll" style="margin-left: 8px">{{
        isSelectAll ? '取消全选' : '全部选择'
      }}</ElButton>
      <ElButton type="primary" @click="handleSave">保存</ElButton>
    </template>
  </ElDialog>
</template>

<script setup lang="ts">
import { fetchRoleEditMenus, fetchRoleMenuAll, fetchRoleMenus } from '@/api/role';

import { AppRouteRecord } from '@/types';
import { formatMenuTitle } from '@/utils/router'


  interface Props {
    modelValue: boolean
    id?: number | null
  }

  interface Emits {
    (e: 'update:modelValue', value: boolean): void
    (e: 'submit'): void
  }

  const props = withDefaults(defineProps<Props>(), {
    modelValue: false,
    roleData: undefined
  })

  const emit = defineEmits<Emits>()
  // 状态管理
  const loading = ref(false)
  
  const treeRef = ref()
  const isExpandAll = ref(true)
  const isSelectAll = ref(false)

  /**
   * 弹窗显示状态双向绑定
   */
  const visible = computed({
    get: () => props.modelValue,
    set: (value) => emit('update:modelValue', value)
  })

  /**
   * 菜单节点类型
   */
  interface MenuNode {
    id?: number
    name?: string
    label?: string
    meta?: {
      title?: string
      authList?: Array<{
        id: number
        authMark: string
        title: string
        checked?: boolean
      }>
    }
    children?: MenuNode[]
    [key: string]: any
  }

  /**
   * 处理菜单数据，将 authList 转换为树形子节点
   * 递归处理菜单树，将权限列表展开为可选择的子节点
   */
  const processedMenuList = computed(() => {
    const processNode = (node: MenuNode): MenuNode => {
      const processed = { ...node }

      // 如果有 authList，将其转换为子节点
      if (node.meta?.authList?.length) {
        const authNodes = node.meta.authList.map((auth) => ({
          id: auth.id,
          name: `${node.name}_${auth.authMark}`,
          label: auth.title,
          authMark: auth.authMark,
          isAuth: true,
          checked: auth.checked || false
        }))

        processed.children = processed.children ? [...processed.children, ...authNodes] : authNodes
      }

      // 递归处理子节点
      if (processed.children) {
        processed.children = processed.children.map(processNode)
      }

      return processed
    }

    return (menuList.value as any[]).map(processNode)
  })

  /**
   * 树形组件配置
   */
  const defaultProps = {
    children: 'children',
    label: (data: any) => formatMenuTitle(data.meta?.title) || data.label || ''
  }

  /**
   * 监听弹窗打开，初始化权限数据
   */
  watch(
    () => props.modelValue,
    (newVal) => {
      if (newVal && props.id) {
        // TODO: 根据角色加载对应的权限数据
        console.log('设置权限:')
        getMenuList()
      }
    }
  )

   /** 菜单列表 */
   const menuList = ref<AppRouteRecord[]>([])
  /**
   * 获取菜单列表数据
   */
  const getMenuList = async (): Promise<void> => {
    loading.value = true
    try {
      const res = await fetchRoleMenuAll()
      menuList.value = convertMenuToRouteTree(res.list)
      getRoleMenus()
    } catch (error) {
      throw error instanceof Error ? error : new Error('获取菜单失败')
    } finally {
      loading.value = false
    }
  }

  const getRoleMenus = async () =>{
    try {
      const res = await fetchRoleMenus({id:props.id!})
      const tree = treeRef.value
      if (!tree) return
      tree.setCheckedKeys(res.menus)
    } catch (error) {
      throw error instanceof Error ? error : new Error('获取菜单失败')
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
   * 关闭弹窗并清空选中状态
   */
  const handleClose = () => {
    visible.value = false
    treeRef.value?.setCheckedKeys([])
  }

  /**
   * 保存权限配置
   */
  const handleSave = async () => {
    // TODO: 调用保存权限接口
    const tree = treeRef.value
    if (!tree) return
    const  checkedKeys = tree.getCheckedKeys()
    await fetchRoleEditMenus({id:props.id!,menus:checkedKeys})
    ElMessage.success('权限保存成功')
    emit('submit')
    handleClose()
  }

  /**
   * 切换全部展开/收起状态
   */
  const toggleExpandAll = () => {
    const tree = treeRef.value
    if (!tree) return

    const nodes = tree.store.nodesMap
    // 这里保留 any，因为 Element Plus 的内部节点类型较复杂
    Object.values(nodes).forEach((node: any) => {
      node.expanded = !isExpandAll.value
    })

    isExpandAll.value = !isExpandAll.value
  }

  /**
   * 切换全选/取消全选状态
   */
  const toggleSelectAll = () => {
    const tree = treeRef.value
    if (!tree) return

    if (!isSelectAll.value) {
      const allKeys = getAllNodeKeys(processedMenuList.value)
      tree.setCheckedKeys(allKeys)
    } else {
      tree.setCheckedKeys([])
    }

    isSelectAll.value = !isSelectAll.value
  }

  /**
   * 递归获取所有节点的 key
   * @param nodes 节点列表
   * @returns 所有节点的 key 数组
   */
  const getAllNodeKeys = (nodes: MenuNode[]): number[] => {
    const keys: number[] = []
    const traverse = (nodeList: MenuNode[]): void => {
      nodeList.forEach((node) => {
        if (node.id) keys.push(node.id)
        if (node.children?.length) traverse(node.children)
      })
    }
    traverse(nodes)
    return keys
  }

  /**
   * 处理树节点选中状态变化
   * 同步更新全选按钮状态
   */
  const handleTreeCheck = () => {
    const tree = treeRef.value
    if (!tree) return

    const checkedKeys = tree.getCheckedKeys()
    const allKeys = getAllNodeKeys(processedMenuList.value)

    isSelectAll.value = checkedKeys.length === allKeys.length && allKeys.length > 0
  }
</script>
