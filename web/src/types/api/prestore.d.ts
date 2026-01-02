declare namespace Prestore {
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
        interface Moadl {
            userId: number | null;
            amount: number | null;
            bonusAmount: number | null;
        }
    }
    namespace Response {
       
        type Info = {
            id: number;
            user:string;
            amount:number;
            bonusAmount:number;
            status:number;
            createTime:string;

        }
        type List = Api.Common.PaginatedResponse<Info>

        type Detail = {
            id: number;
            code:string;
            manage: string;
            user:string;
            amount:number;
            bonusAmount:number;
            status:number;
            reason:string;
            createTime:string;
        }

    }
}