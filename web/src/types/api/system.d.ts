declare namespace System {
    namespace Params {
        interface Base {
            title: string;
            domain: string;
            logo: string;
            icon: string;
            description: string;
            icp: string;
        }
        interface User {
            avatar: string;
            levelId: number | null;
            checkInExperience: number | null;
        }
        interface Withdraw {
            symbol: null,
            minWithdraw: null,
            withDrawPercent: null,
        }
        interface File {
            engine: null,
            path: null,
            fileSize: null,
            fileType: any,
            imageSize: null,
            imageType: any,
        }
        //  小程序
        interface WechatMiniProgram {
            appId: string | nill,
            secret: string | nill,
        }
        //  微信支付
        interface WechatPay {
            mchId: string | null;
            mchKey: string | null;
            serialNo: string | null;
            apiclientCert: string | null;
            apiclientKey: string | null;
            pubKey: string | null;
            pubId: string | null;
        }
    }
    namespace Response {

        type Base =  {
            title: string;
            domain: string;
            logo: string;
            icon: string;
            description: string;
            icp: string;
        }

        type User =  {
            avatar: string;
            levelId: number | null;
            checkInExperience: number | null;
        }

        type Withdraw =  {
            symbol: null,
            minWithdraw: null,
            withDrawPercent: null,
        }

        type File =  {
            engine: null,
            path: null,
            fileSize: null,
            fileType: string[],
            imageSize: null,
            imageType: string[],
        }

        type WechatMiniProgram =  {
            appId: string | nill,
            secret: string | nill,
        }
        type WechatPay =  {
            mchId: string | null;
            mchKey: string | null;
            serialNo: string | null;
            apiclientCert: string | null;
            apiclientKey: string | null;
            pubKey: string | null;
            pubId: string | null;
        }
    }
}