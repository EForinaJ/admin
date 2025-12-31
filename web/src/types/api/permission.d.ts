declare namespace Permission {
    namespace Params {
        interface Model {
            id?: number; // 权限ID
            pId: number; // 父Id
            permission: string; // 权限标识
            name: string; // 权限名称
            description: string; // 权限描述
        }
    }
    namespace Response {
       
        type Info = {
            id: number; // 权限ID
            pId: number; // 父Id
            permission: string; // 权限标识
            name: string; // 权限名称
            description: string; // 权限描述
            createTime?: string | null; // 创建时间，假设 gtime.Time 在序列化后会变成 ISO8601 格式的字符串
        }
        type Edit = {
            id: number; // 权限ID
            pId: number; // 父Id
            permission: string; // 权限标识
            name: string; // 权限名称
            description: string; // 权限描述
        }
    }
}