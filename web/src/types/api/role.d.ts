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
            description: string;
            status: number;
        }
    }
    namespace Response {
       
        type Info = {
            id: number;
            name: string;
            code: string;
            description: string;
            status: number;
            createTime: string;
        }

        type List = Api.Common.PaginatedResponse<Info>

        type Edit = {
            id: number;
            name: string;
            code: string;
            description: string;
            status: number;
        }
    }
}