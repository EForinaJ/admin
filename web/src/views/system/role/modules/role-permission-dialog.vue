<template>
  <ElDialog
    v-model="visible"
    title="接口权限"
    width="520px"
    align-center
    class="el-dialog-border"
    @close="handleClose"
  >
    <ElScrollbar height="70vh">
      <ElTree
        ref="treeRef"
        :data="treeList"
        show-checkbox
        node-key="id"
        :default-expand-all="isExpandAll"
        :default-checked-keys="[1, 2, 3]"
        :props="defaultProps"
        @check="handleTreeCheck"
      >
        <template #default="{ data }">
          <div style="display: flex; align-items: center">
            <span>{{ defaultProps.label(data) }}</span>
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
  import {  fetchRoleEditPermissions, fetchRolePermissionAll, fetchRolePermissions } from '@/api/role';
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
   * 监听弹窗打开，初始化权限数据
   */
  watch(
    () => props.modelValue,
    (newVal) => {
      if (newVal && props.id) {
        // TODO: 根据角色加载对应的权限数据
        getList()
      }
    }
  )

   /** 菜单列表 */
   const treeList = ref<PermissionTree[]>([])
  /**
   * 获取菜单列表数据
   */
  const getList = async (): Promise<void> => {
    loading.value = true
    try {
      const res = await fetchRolePermissionAll()
      treeList.value = convertTree(res.list)
      getRolePermission()
    } catch (error) {
      throw error instanceof Error ? error : new Error('获取菜单失败')
    } finally {
      loading.value = false
    }
  }

  const getRolePermission = async () =>{
    try {
      const res = await fetchRolePermissions({id:props.id!})
      const tree = treeRef.value
      if (!tree) return
      tree.setCheckedKeys(res.permissiones)
    } catch (error) {
      throw error instanceof Error ? error : new Error('获取菜单失败')
    }
  }


    /**
   * 应用路由记录接口
   * 扩展 Vue Router 的路由记录类型
   */
  interface PermissionTree extends Omit<Permission.Response.Info,  'children'> {
    children?: PermissionTree[]
  }

  const convertTree = (list: Permission.Response.Info[]): PermissionTree[] => {
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
     /**
   * 树形组件配置
   */
  const defaultProps = {
    children: 'children',
    label: (data: any) =>  data.name
  }
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
    console.log(checkedKeys)
    await fetchRoleEditPermissions({id:props.id!,permissions:checkedKeys})
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
      const allKeys = getAllNodeKeys(treeList.value)
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
  const getAllNodeKeys = (nodes: PermissionTree[]): number[] => {
    const keys: number[] = []
    const traverse = (nodeList: PermissionTree[]): void => {
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
    const allKeys = getAllNodeKeys(treeList.value)

    isSelectAll.value = checkedKeys.length === allKeys.length && allKeys.length > 0
  }
</script>
