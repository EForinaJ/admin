declare namespace Aftersales {
    namespace Params {
        type Query = {
            page: number;
            limit: number;
            code?: string;
            orderCode?: string;
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
            product:{
                id:number;
                name:string;
                category:string;
                pic:string
            };
            user:string;
            amount:number;
            status:number;
            createTime:string;
        }
        type List = Api.Common.PaginatedResponse<Info>

        type Detail = {
            id: number;
            code:string;
            manage: string;
            amount:number;
            type:number;
            reason:string;
            status:number;
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
            reject:string;
            createTime:string;
        }

    }
}