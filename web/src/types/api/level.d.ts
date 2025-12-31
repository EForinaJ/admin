declare namespace Level {
    namespace Params {
        /** 角色搜索参数 */
        type Query = Partial<
            Pick<'name'> &
            Api.Common.CommonSearchParams
        >
        interface Model {
            id?: number;
            name: string | null;
            level: number | null;
            icon: string | null;
            experience: number | null;
        }
    }
    namespace Response {
       
        type Info = {
            id: number;
            name: string;
            level: number;
            icon: string;
            experience: number;
            benefits: string;
            createTime: string;
        }
        type List = Api.Common.PaginatedResponse<Info>

        
        type Edit = {
            id?: number;
            name: string | null;
            level: number | null;
            icon: string | null;
            experience: number | null;
        }
    }
}