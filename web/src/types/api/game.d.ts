declare namespace Game {
    namespace Params {
        type Query = {
            page: number;
            limit: number;
            name?: string;
        }
        interface Model {
            id?: number;
            name: string | null;
            pic: string | null;
            description: string | null;
        }
    }
    namespace Response {
       
        type Info = {
            id: number;
            name: string;
            game: string;
            pic: string;
            description: string;
            createTime: string;
        }
        type List = Api.Common.PaginatedResponse<Info>

        
        type Edit = {
            id?: number | null,
            name: string | null;
            gameId:number | null;
            pic: string | null;
            description: string | null;
        }
    }
}