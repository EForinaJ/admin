declare namespace Distribute {
    namespace Params {
        type Query = {
            page: number;
            limit: number;
            name?: string;
            code?: string;
        }
        interface Model {
            witkeyId?: number| null;
            code: string | null;
        }
        interface Cancel {
            id?: number| null;
            reason: string | null;
        }
    }
    namespace Response {
       
        type Info = {
            id: number;
            code:string;
            manage: string;
            witkey:string;
            game:string;
            title:string;
            reason:string;
            isCancel:number;
            createTime:string;
        }
        type List = Api.Common.PaginatedResponse<Info>

    }
}