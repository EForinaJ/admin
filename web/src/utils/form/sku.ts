import { CartesianProduct } from "./cartesian";

interface SpecAttribute {
    name: string;
    values: string[];
}
  
interface SpecCombination {
  attrNames: string[]; // 属性名称
  values: string[];     // 对应的值
  attrs: Array<{
    attr:string,
    value:string
  }>;  
  code: string;          // 唯一标识
  price: number;        // 出售价格
  originalPrice: number; // 原价价格
  stock: number; // 库存
}
  
export class SpecCartesianProduct {
    /**
     * 为商品规格生成笛卡尔积
     */
    static generateSpecCombinations(attributes: SpecAttribute[]): SpecCombination[] {
      if (!attributes || attributes.length === 0) {
        return [];
      }
  
      // 提取所有属性的值数组
      const valueArrays = attributes.map(attr => attr.values);
      
      // 计算笛卡尔积
      const valueCombinations = CartesianProduct.compute(valueArrays);
      
      // 转换为规格组合对象
      return valueCombinations.map((values, index) => {
        const attributeNames = attributes.map(attr => attr.name);
        const code = this.generateCombinationKey(attributeNames, values);
        
        return {
          attrs: this.generateCombinationAttrs(attributeNames, values),
          attrNames: attributeNames,
          values,
          code,
          display: this.formatDisplay(attributeNames, values),
          price:0,
          originalPrice:0,
          stock:0,
        };
      });
    }

    /**
     * 生成组合对象
     */
    private static generateCombinationAttrs(attributes: string[], values: string[]): Array<{
      attr:string,
      value:string
    }> {
      return attributes.map((attr, index) => {
        return {
          attr:attr,
          value:values[index]
        }
      })
    }
  
    /**
     * 生成组合的唯一键
     */
    private static generateCombinationKey(attributes: string[], values: string[]): string {
      return attributes.map((attr, index) => 
        `${attr}:${values[index]}`
      ).join('|');
    }
  
    /**
     * 格式化显示文本
     */
    private static formatDisplay(attributes: string[], values: string[]): string {
      return attributes.map((attr, index) => 
        `${attr}:${values[index]}`
      ).join(' + ');
    }
  
    /**
     * 过滤有效的规格组合（可自定义过滤条件）
     */
    static filterValidCombinations(
      combinations: SpecCombination[], 
      filterFn?: (combo: SpecCombination) => boolean
    ): SpecCombination[] {
      const defaultFilter = (combo: SpecCombination) => {
        // 示例过滤：排除包含"红色"和"XL"的不合理组合
        return !(combo.values.includes('红色') && combo.values.includes('XL'));
      };
  
      return combinations.filter(filterFn || defaultFilter);
    }
}