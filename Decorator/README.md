### Decorator 裝飾模式： 
> 將額外權責動態附加於物件身上, 不必衍生子類別即可彈性擴增功能

將一個物件包裝起來, 以提供新的行為
將額外權責動態附加於物件身上, 不必衍生子類別即可彈性擴增功能
動態低給一個物件加入一些額外的職責, 就增加功能來說, 裝飾模式比產生子類別更為靈活
所有 decorator 在操作時需使用函數, 裝飾者順序很重要, 比如加密資料和過濾辭彙都可以是資料持久化前的裝飾功能


![UML](https://github.com/kimi0230/DesignPatternGolang/blob/master/UML/Decorator.png?raw=true)