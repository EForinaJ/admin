declare namespace Menu {
    namespace Params {
        interface Model {
            id: number
            pId: number
            name: string
            path: string
            label: string
            component: string
            icon: string
            isEnable: boolean
            sort: number
            isMenu: boolean
            keepAlive: boolean
            isHide: boolean
            isHideTab: boolean
            link: string
            isIframe: boolean
            showBadge: boolean
            showTextBadge: string
            fixedTab: boolean
            activePath: string
            isFullPage: boolean
            authName: string
            authLabel: string
            authIcon: string
            authSort: number
        }
    }
    namespace Response {
       
        type List = {
            id: number;
            pId: number;
            menuType: string;
            name: string;
            path: string;
            label: string;
            component: string;
            icon: string;
            isEnable: boolean;
            sort: number;
            isMenu: bool;
            keepAlive: bool;
            isHide: bool;
            isHideTab: bool;
            link: string;
            isIframe: bool;
            showBadge: bool;
            showTextBadge: bool;
            fixedTab: bool;
            activePath: string;
            isFullPage: bool;
            authName: string;
            authLabel: string;
            authIcon: string;
            authSort: number;
            children: List[];
        }
        type Edit = {
            id: number;
            pId: number;
            menuType: 'menu' | 'button';
            name: string;
            path: string;
            label: string;
            component: string;
            icon: string;
            isEnable: boolean;
            sort: number;
            isMenu: bool;
            keepAlive: bool;
            isHide: bool;
            isHideTab: bool;
            link: string;
            isIframe: bool;
            showBadge: bool;
            showTextBadge: bool;
            fixedTab: bool;
            activePath: string;
            isFullPage: bool;
            authName: string;
            authLabel: string;
            authIcon: string;
            authSort: number;
            children: List[];
        }
    }
}