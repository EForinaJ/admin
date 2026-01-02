declare namespace User {
    namespace Params {
        /** 角色搜索参数 */
        type Query = Partial<
            Pick<'name' | 'status' | 'phone'> &
            Api.Common.CommonSearchParams
        >
        interface Model {
            id?: number | null,
            name: string | null,
            phone: string | null,
            password: string | null,
            address: string[],
            birthday: number,
            description: string | null,
            sex: number | null,
            avatar: string,
            status: number,
        }

        interface ChangeBalance {
            id?: number | null,
            mode: number | null;
            amount: number | null;
            remark: string | null;
        }
    }
    namespace Response {
       
        type Info = {
            id: number;
            name: string;
            phone: string;
            avatar: string;
            balance: number;
            sex: number;
            status: number;
            createTime: string;
        }
        type List = Api.Common.PaginatedResponse<Info>
        type Edit = {
            id?: number | null,
            name: string | null,
            phone: string | null,
            password: string | null,
            address: string[],
            birthday: number,
            description: string | null,
            sex: number | null,
            avatar: string,
            status: number,
        }
        type Detail = {
            id: number;
            name: string;
            phone: string;
            address: string[];
            birthday: string;
            sex: number;
            balance: number;
            marginDeposit: number;
            avatar: string;
            description: string;
            witkey:{
                game:string;
                title:string;
                album: Array<string>;
                rate: number;
                credentialType: number;
                credentialName: string;
                credentialNumber: string;
                credentialFrontImage: string;
                credentialAgainstImage: string;
                createTime:string;
            }
            level:{
                id:number;
                name:string;
                icon:string;
            }
            status: number;
            createTime: string;
        }

        type BalanceList = Api.Common.PaginatedResponse<{
            id: number;
            related:string;
            type:number;
            after:number;
            before:number;
            amount:number;
            mode:number;
            remark:string;
            createTime: string;
        }>
    }
}