declare namespace Order {
    namespace Params {
        type Query = {
            page: number;
            limit: number;
            code?: string;
            status?: number;
        }
        type AddDiscount = {
            id: number;
            amount?: number | null;
        }
        type Refund = {
            id: number;
            type: number;
            amount: number | null;
            reason: string | null;
        }
    }
    namespace Response {
       
        type Info = {
            id: number;
            code: string;
            user: string;
            product: {
                id: number;
                name: string;
                pic: string;
                quantity: number;
                unitPrice: number;
                game: string;
                category: string;
            };
            payAmount: number;
            status: number;
            payMode: number;
            createTime: string;
        }
        type List = Api.Common.PaginatedResponse<Info>

        type Detail = {
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
        }
    }
}