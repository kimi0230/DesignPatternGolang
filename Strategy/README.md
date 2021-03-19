### Strategy 策略模式： 
> 定義一整組演算法, 講每一個演算法封裝起來, 可互換使用, 更可在不影響外界的情況下個別抽換所以引用的演算法

可以將互換的行為封裝起來, 並使用轉接方式決定要使用何者
用戶端實體化的事 CashContext 的物件, 調用的是 CashContext 的方法 GetResult
使具體的收費演算法測底與用戶端分離

![UML](https://github.com/kimi0230/DesignPatternGolang/blob/master/UML/Strategy.png?raw=true)