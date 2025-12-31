declare namespace Withdraw {
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
            witkey:string;
            type:number;
            amount:number;
            serviceFee:number;
            settledAmount:number;
            status:number;
            createTime:string;

        }
        type List = Api.Common.PaginatedResponse<Info>

        type Detail = {
            id: number;
            code:string;
            manage: string;
            witkey:string;
            amount:number;
            settledAmount:number;
            serviceFee:number;
            type:number;
            number:string;
            name:string;
            status:number;
            reason:string;
            createTime:string;
        }

    }
}