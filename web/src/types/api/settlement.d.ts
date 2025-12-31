declare namespace Settlement {
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
            orderCode:string;
            witkey:string;
            amount:number;
            status:number;
            createTime:string;
        }
        type List = Api.Common.PaginatedResponse<Info>

        type Detail = {
            id: number;
            code:string;
            order:{
                id: number;
                code:string;
                createTime:string
                user: {
                    id:number
                    name:string;
                    avatar:string;
                };
                product: {
                    id: number;
                    name: string;
                    pic: string;
                    quantity: number;
                    unitPrice: number;
                    unit: string;
                    game: string;
                    category: string;
                };
                totalAmount: number;
                discountAmount: number;
                actualAmount:number;
                payAmount: number;
                status: number;
                payStatus: number;
                payMode: number;
                payTime:string | null;
                startTime:string | null;
                finishTime:string | null;
                requirements:string | null;
            };
            manage: string;
            witkey:string;
            amount:number;
            commission:number;
            serviceCharge:number;
            status:number;
            reason:string;
            images:string[];
            createTime:string;
        }

    }
}