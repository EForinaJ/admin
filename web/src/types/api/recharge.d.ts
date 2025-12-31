declare namespace Recharge {
    namespace Params {
        type Query = {
            page: number;
            limit: number;
            name?: string;
            code?: string;
        }
        interface Apply {
            id?: number| null;
            status: number;
            reason: string | null;
        }
    }
    namespace Response {
       
        type Info = {
            id: number;
            code:string;
            user:string;
            payMode:number;
            amount:number;
            status:number;
            createTime:string;

        }
        type List = Api.Common.PaginatedResponse<Info>
    }
}