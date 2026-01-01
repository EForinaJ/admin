declare namespace Role {
    namespace Params {
        /** 角色搜索参数 */
        type Query = Partial<
            Pick<'name'> &
            Api.Common.CommonSearchParams
        >
        interface Model {
            id?: number; // 权限ID
            name: string;
            code: string;
            type: number;
            description: string;
        }
    }
    namespace Response {
       
        type Info = {
            id: number;
            name: string;
            code: string;
            description: string;
            type: number;
            createTime: string;
        }

        type List = Api.Common.PaginatedResponse<Info>

        type Edit = {
            id: number;
            name: string;
            code: string;
            type: number;
            description: string;
        }
    }
}