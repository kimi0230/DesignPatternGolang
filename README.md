# 設計模式 (Design Pattern)

基於 程杰-大話設計模式 和 Design Patterns: Elements of Reusable Object-Oriented Software. 
使用 Golang 來實現

## 生成模式 (Creational Patterns)
牽涉到將物件實體化。這類模式都提供一個方法，將客戶從所需要實體化的物件中鬆綁出來
* [Factory Method 工廠模式：](https://github.com/kimi0230/DesignPatternGolang/tree/master/FactoryMethod) 定義可滋生成物件的介面, 但讓子類別去決定該具現出哪一種類別的物件. 此模式讓類別將具現化程序交給子類別去處置
* [Prototype 原型模式：](https://github.com/kimi0230/DesignPatternGolang/tree/master/Prototype) 制定可用原型個體生成的物件類型, 爾後只需複製此原型即可生成新物件
* [Builder 建造者模式：](https://github.com/kimi0230/DesignPatternGolang/tree/master/Builder) 從複雜物件的佈局中抽取出生成程序, 以便用同一個生成程序製造各種不同的物件佈局
* [Abstract Factory 抽象工廠模式：](https://github.com/kimi0230/DesignPatternGolang/tree/master/AbstractFactory) 以同一個介面來建立一整組相關或相依的物件, 不需要點名個物件真正所屬的具象類別

## 結構模式 (Structural Patterns)
讓你合成類別或物件到大型的結構
* [Decorator 裝飾模式：](https://github.com/kimi0230/DesignPatternGolang/tree/master/Decorator) 將額外權責動態附加於物件身上, 不必衍生子類別即可彈性擴增功能
* [Proxy 代理模式：](https://github.com/kimi0230/DesignPatternGolang/tree/master/Proxy) 替其他物件預留代理者空位, 就此控制存取其他物件
* [Facade 外觀模式：](https://github.com/kimi0230/DesignPatternGolang/tree/master/Facade) 替子系統裡的一堆介面定義一套統一的高階介面, 讓子系統更容易使用

## 行為模式 (Behavioral Patterns)
模述類別和物件如何互動，以及各自的責任
* [Strategy 策略模式：](https://github.com/kimi0230/DesignPatternGolang/tree/master/Strategy) 定義一整組演算法, 講每一個演算法封裝起來, 可互換使用, 更可在不影響外界的情況下個別抽換所以引用的演算法
* [Template Method 範本方式模式:](https://github.com/kimi0230/DesignPatternGolang/tree/master/TemplateMethod) 對於操作, 只先定義好演算法的輪廓, 某些步驟則留給子類別去填飽, 以便在不改變演算法整體構造的情況下讓子類別去精煉某些步驟
* [Observer 觀察者模式:](https://github.com/kimi0230/DesignPatternGolang/tree/master/Observer) 定義一對多的物件依存關係, 讓物件狀態一有變動, 就自動通知其他相依物件做該做的更新動作