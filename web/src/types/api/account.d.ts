declare namespace Account {
    namespace Params {
        interface Model {
            name: string | null;
            email: string | null;
            address: string[];
            sex: number | null;
            avatar: string | null;
            description: string | null;
            tags: string[];
        }
        interface ChangePass {
            oldPass: string | null;
            newPass: string | null;
            confirmPass: string | null;
        }
    }
    namespace Response {
        /** 登录参数 */
        type Info = {
            // permission: string[]
            id: number
            name: string| null;
            avatar?: string| null;
            email: string | null;
            address: string[];
            roles: string[];
            sex: number | null;
            description: string| null;
            tags: string[];
            loginIp?: string| null;
        }
    }
}