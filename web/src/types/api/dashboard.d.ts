declare namespace Dashboard {
    namespace Response {
        type Detail = {
            userCount:{
                totalCount:number;
                todayCount:number;
            },
            orderCount:{
                totalCount:number;
                todayCount:number;
            },
            salesAmount:{
                amountTotal:number;
                amountToday:number;
            },
            profitAmount:{
                amountTotal:number;
                amountToday:number;
            },
            orderStatistic:{
                weekdays:string[];
                count:number[];
            },
            userStatistic:{
                days:string[];
                count:number[];
            },
            pendingServiceOrder:{
                totalCount:number;
                todayCount:number;
            },
            applySettlement:{
                totalCount:number;
                todayCount:number;
            },
            applyWithdraw:{
                totalCount:number;
                todayCount:number;
            },
            applyComment:{
                totalCount:number;
                todayCount:number;
            }
        }

    }
}