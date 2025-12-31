declare namespace Product {
    namespace Params {
        /** 角色搜索参数 */
        type Query = {
            page: number;
            limit: number;
            name?: string;
            gameId?: number | null;
            status?: number | null
        }
        interface Model {
            id?: number;
            name: string;
            gameId: number |null;
            price: number | null;
            rate: number;
            witkeyCount: number | null;
            purchaseLimit: number | null;
            salesCount: number | null;
            categoryId: number |null;
            unit: string | null;
            pic: string;
            description: string;
            content: string;
            status: number;
            // skuType:number;
            // specification:Array<{
            //     name: string;
            //     values: Array<string>;
            // }>;
            // sku:Array<{
            //     attrNames: string[]; // 属性名称
            //     values: string[];     // 对应的值
            //     attrs: Array<{
            //       attr:string,
            //       value:string
            //     }>;  
            //     code: string;          // 唯一标识
            //     price: number;        // 出售价格
            //     originalPrice: number; // 原价价格
            //     stock: number; // 库存
            // }>;
        }
    }
    namespace Response {
        type Info = {
            id: number;
            name: string;
            code: string;
            pic: string;
            game: string;
            category: string;
            price: number;
            rate: number;
            salesCount:number;
            status:number;
            createTime: string;
        }
        type List = Api.Common.PaginatedResponse<Info>

        
        type Edit = {
            id?: number;
            name: string;
            gameId: number |null;
            price: number | null;
            rate: number;
            witkeyCount: number | null;
            salesCount: number | null;
            categoryId: number |null;
            unit: string | null;
            purchaseLimit: number | null;
            pic: string;
            description: string;
            content: string;
            status: number;
        }
    }
}