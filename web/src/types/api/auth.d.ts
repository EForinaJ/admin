declare namespace Auth {
    namespace Params {
        /** 登录参数 */
        interface Login {
            phone: string
            password: string
        }
    }
    namespace Response {
        /** 登录参数 */
        interface Login {
            token: string
            refreshToken: string
        }
    }
}