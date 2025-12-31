declare namespace Site {
    namespace Params {
       
    }
    namespace Response {
        /** 登录参数 */
        type Info =  {
            title: string;       // 网站标题
            domain: string;      // 网站域名
            logo: string;        // 网站Logo图片地址
            icon: string;        // 网站图标地址
            description: string; // 网站描述
            address: string;     // 网站地址
            icp: string;         // ICP备案号
            copyright: string;   // 版权信息
            minFileSize : number,
            bigFileSize: number,
            mediaType: Array<string>,
            symbol: string;   // 货币类型
        }
    }
}