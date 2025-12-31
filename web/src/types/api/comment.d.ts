declare namespace Comment {
    namespace Params {
        type Query = {
            page: number;
            limit: number;
            status?: number;
            code?: string;
        }
        interface Apply {
            id?: number| null;
            status: number;
        }
    }
    namespace Response {
       
        type Info = {
            id: number;
            user:string;
            product:{
                id:number;
                pic:string;
                name:string;
                category:string;
            };
            content:string;
            images:string[];
            rate:number;
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