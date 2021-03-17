# 設計模式 (Design Pattern)

基於 程杰-大話設計模式, 使用 Golang 來實現

## 生成模式 (Creational Patterns)
* [Factory Method 工廠模式：](https://github.com/kimi0230/DesignPatternGolang/tree/master/FactoryMethod) 定義可滋生成物件的介面, 但讓子類別去決定該具現出哪一種類別的物件. 此模式讓類別將具現化程序交給子類別去處置
* [Prototype 原型模式：](https://github.com/kimi0230/DesignPatternGolang/tree/master/Prototype) 制定可用原型個體生成的物件類型, 爾後只需複製此原型即可生成新物件

## 結構模式 (Structural Patterns)
* [Decorator 裝飾模式：](https://github.com/kimi0230/DesignPatternGolang/tree/master/Decorator) 將額外權責動態附加於物件身上, 不必衍生子類別即可彈性擴增功能
* [Proxy 代理模式：](https://github.com/kimi0230/DesignPatternGolang/tree/master/Proxy) 替其他物件預留代理者空位, 就此控制存取其他物件

## 行為模式 (Behavioral Patterns)
* [Strategy 策略模式：](https://github.com/kimi0230/DesignPatternGolang/tree/master/Strategy) 定義一整族演算法, 講每一個演算法封裝起來, 可互換使用, 更可在不影響外界的情況下個別抽換所以引用的演算法
* [Template Method 範本方式模式:](https://github.com/kimi0230/DesignPatternGolang/tree/master/TemplateMethod) 對於操作, 只先定義好演算法的輪廓, 某些步驟則留給子類別去填飽, 以便在不改變演算法整體構造的情況下讓子類別去精煉某些步驟