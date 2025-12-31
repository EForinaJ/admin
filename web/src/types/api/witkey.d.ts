declare namespace Witkey {
    namespace Params {
        /** 角色搜索参数 */
        type Query = Partial<
            Pick<'name' | 'phone'> &
            Api.Common.CommonSearchParams
        >
        interface Model {
            id?: number | null,
            titleId: number | null;
            gameId: number| null;
            userId: number| null;
            album: string[];
            rate: number;
        }

        interface ChangeCommission {
            id?: number | null,
            mode: number | null;
            amount: number | null;
            remark: string | null;
        }
    }
    namespace Response {
       
        type Info = {
            id: number;
            user: {
                name:string;
                avatar:string;
                phone:string;
            };
            game: string;
            title: string;
            commission: number;
            rate: number;
            createTime: string;
        }
        type List = Api.Common.PaginatedResponse<Info>

        
        type Edit = {
            id?: number | null,
            titleId: number | null;
            gameId: number| null;
            userId?: number| null;
            album: string[];
            rate: number;
        }
        type Detail = {
            id: number;
            user: {
                name:string;
                avatar:string;
                phone:string;
            };
            game: string;
            title: string;
            commission: number;
            rate: number;
            album: string[];
            createTime: string;
        }

        type CommissionList = Api.Common.PaginatedResponse<{
            id: number;
            related:string;
            type:number;
            after:number;
            before:number;
            amount:number;
            mode:number;
            remark:string;
            createTime: string;
        }>
    }
}