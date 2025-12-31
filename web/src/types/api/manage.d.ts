declare namespace Manage {
    namespace Params {
        /** 角色搜索参数 */
        type Query = Partial<
            Pick<'name' | 'status'> &
            Api.Common.CommonSearchParams
        >
        interface Model {
            id?: number;
            name: string;
            wechat: string;
            password: string;
            phone: string;
            sex: number;
            avatar: string;
            roles: number[];
            status: Status ;
        }
    }
    namespace Response {
       
        type Info = {
            id: number;
            name: string;
            email: string;
            phone: string;
            sex: number;
            avatar: string;
            roles: string[];
            loginIp: string;
            loginTime?: string;
            remark: string;
            status: number;
            createTime: string;
        }

        type List = Api.Common.PaginatedResponse<Info>

        type Edit = {
            id?: number;
            name: string;
            wechat: string;
            description: string;
            password: string;
            phone: string;
            sex: number;
            avatar: string;
            roles: number[];
            status: Status ;
        }
    }
}