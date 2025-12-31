// 更健壮的笛卡尔积实现
export class CartesianProduct {
    /**
     * 计算多个数组的笛卡尔积
     */
    static compute<T>(arrays: T[][]): T[][] {
      if (!arrays || arrays.length === 0) {
        return [[]];
      }
  
      // 验证输入
      arrays.forEach((arr, index) => {
        if (!Array.isArray(arr)) {
          throw new Error(`参数 ${index} 不是数组`);
        }
      });
  
      // 过滤空数组
      const nonEmptyArrays = arrays.filter(arr => arr.length > 0);
      if (nonEmptyArrays.length === 0) {
        return [[]];
      }
  
      return nonEmptyArrays.reduce<T[][]>(
        (acc, current) => {
          const result: T[][] = [];
          acc.forEach(prev => {
            current.forEach(item => {
              result.push([...prev, item]);
            });
          });
          return result;
        },
        [[]]
      );
    }
  
    /**
     * 计算笛卡尔积的数量（不实际生成组合）
     */
    static count<T>(arrays: T[][]): number {
      if (arrays.length === 0) return 0;
      return arrays.reduce((total, arr) => total * Math.max(arr.length, 1), 1);
    }
  
    /**
     * 异步生成笛卡尔积（处理大数据集）
     */
    static async *computeAsync<T>(arrays: T[][]): AsyncGenerator<T[]> {
      if (arrays.length === 0) {
        yield [];
        return;
      }
  
      const nonEmptyArrays = arrays.filter(arr => arr.length > 0);
      if (nonEmptyArrays.length === 0) {
        yield [];
        return;
      }
  
      let result: T[][] = [[]];
      
      for (const array of nonEmptyArrays) {
        const temp: T[][] = [];
        
        for (const res of result) {
          for (const item of array) {
            // 每生成一定数量就 yield，避免阻塞
            if (temp.length % 1000 === 0) {
              await new Promise(resolve => setTimeout(resolve, 0));
            }
            temp.push([...res, item]);
          }
        }
        
        result = temp;
      }
  
      for (const combination of result) {
        yield combination;
      }
    }
}