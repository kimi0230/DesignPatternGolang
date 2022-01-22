# 設計模式 (Design Pattern)
[![Build latest tag](https://github.com/kimi0230/DesignPatternGolang/actions/workflows/releace.yml/badge.svg)](https://github.com/kimi0230/DesignPatternGolang/actions/workflows/releace.yml) [![Build my gitbook and deploy to gh-pages](https://github.com/kimi0230/DesignPatternGolang/actions/workflows/build.yml/badge.svg)](https://github.com/kimi0230/DesignPatternGolang/actions/workflows/build.yml) [![pages-build-deployment](https://github.com/kimi0230/DesignPatternGolang/actions/workflows/pages/pages-build-deployment/badge.svg)](https://github.com/kimi0230/DesignPatternGolang/actions/workflows/pages/pages-build-deployment)

基於 程杰-大話設計模式 和 Design Patterns: Elements of Reusable Object-Oriented Software. 
使用 Golang 來實現

https://kimi0230.github.io/DesignPatternGolang/

---

## 生成模式 (Creational Patterns)
> 牽涉到將物件實體化。這類模式都提供一個方法，將客戶從所需要實體化的物件中鬆綁出來

### [Factory Method 工廠模式：](https://github.com/kimi0230/DesignPatternGolang/tree/master/FactoryMethod) 
定義可滋生成物件的介面, 但讓子類別去決定該具現出哪一種類別的物件. 此模式讓類別將具現化程序交給子類別去處置

### [Prototype 原型模式：](https://github.com/kimi0230/DesignPatternGolang/tree/master/Prototype) 
制定可用原型個體生成的物件類型, 爾後只需複製此原型即可生成新物件

### [Builder 建造者模式：](https://github.com/kimi0230/DesignPatternGolang/tree/master/Builder) 
從複雜物件的佈局中抽取出生成程序, 以便用同一個生成程序製造各種不同的物件佈局

### [Abstract Factory 抽象工廠模式：](https://github.com/kimi0230/DesignPatternGolang/tree/master/AbstractFactory) 
以同一個介面來建立一整組相關或相依的物件, 不需要點名個物件真正所屬的具象類別

### [Singleton 獨體模式：](https://github.com/kimi0230/DesignPatternGolang/tree/master/Singleton) 
確保類別只會有一個物件實體存在, 並提供單一存取窗口

---

## 結構模式 (Structural Patterns)
> 讓你合成類別或物件到大型的結構

### [Decorator 裝飾模式：](https://github.com/kimi0230/DesignPatternGolang/tree/master/Decorator) 
將額外權責動態附加於物件身上, 不必衍生子類別即可彈性擴增功能

### [Proxy 代理模式：](https://github.com/kimi0230/DesignPatternGolang/tree/master/Proxy) 
替其他物件預留代理者空位, 就此控制存取其他物件

### [Facade 外觀模式：](https://github.com/kimi0230/DesignPatternGolang/tree/master/Facade) 
替子系統裡的一堆介面定義一套統一的高階介面, 讓子系統更容易使用

### [Adapter 轉接器模式：](https://github.com/kimi0230/DesignPatternGolang/tree/master/Adapter) 
將類別的介面轉換成外界所預期的另一種介面, 讓原先囿於見面不相容問題而無法協力合作的類別能夠兜在一起用

### [Composite 組合模式：](https://github.com/kimi0230/DesignPatternGolang/tree/master/Composite) 
將物件組織成樹狀結構, 部分-全體 層級關係, 讓外界以一致性的方式對待個別物件和整體物件

### [Bridge 橋接模式：](https://github.com/kimi0230/DesignPatternGolang/tree/master/Bridge) 
將實作體系與抽象體系分離開來, 讓兩者能各自更動各自演進

### [Flyweight 享元模式：](https://github.com/kimi0230/DesignPatternGolang/tree/master/Flyweight) 
以共享機制有效地支援一大堆小規模的物件

---
## 行為模式 (Behavioral Patterns)
> 模述類別和物件如何互動，以及各自的責任

### [Strategy 策略模式：](https://github.com/kimi0230/DesignPatternGolang/tree/master/Strategy) 
定義一整組演算法, 講每一個演算法封裝起來, 可互換使用, 更可在不影響外界的情況下個別抽換所以引用的演算法

### [Template Method 範本方式模式:](https://github.com/kimi0230/DesignPatternGolang/tree/master/TemplateMethod) 
對於操作, 只先定義好演算法的輪廓, 某些步驟則留給子類別去填飽, 以便在不改變演算法整體構造的情況下讓子類別去精煉某些步驟

### [Observer 觀察者模式:](https://github.com/kimi0230/DesignPatternGolang/tree/master/Observer) 
定義一對多的物件依存關係, 讓物件狀態一有變動, 就自動通知其他相依物件做該做的更新動作

### [State 狀態模式:](https://github.com/kimi0230/DesignPatternGolang/tree/master/State) 
讓物件的外顯行為隨內部狀態的改變而改變, 彷彿連類別也變了似的

### [Memento 備忘錄模式:](https://github.com/kimi0230/DesignPatternGolang/tree/master/Memento) 
在不違反封裝性的前提下, 捕捉物件的內部狀態並存在外面,以便日後回復至此一狀態

### [Iterator 迭代器模式:](https://github.com/kimi0230/DesignPatternGolang/tree/master/Iterator) 
無需知曉聚合物件的內部細節, 即可依序存取內含的每一個元素

### [Command 命令模式:](https://github.com/kimi0230/DesignPatternGolang/tree/master/Command) 
將訊息封裝成物件, 以便能用各種不同訊息, 暫佇, 紀律, 復原等方式加以參數化處理

### [Chain of Responsibility 職責鏈模式:](https://github.com/kimi0230/DesignPatternGolang/tree/master/ChainofResponsibility) 
讓多個物件都有機會處理某一訊息, 以降低訊息發送者和接收者之間的耦合關係. 它將接收者物件串連起來, 讓訊息流經其中,直到被處理了為止

### [Mediator 仲介者模式:](https://github.com/kimi0230/DesignPatternGolang/tree/master/Mediator) 
定義可將一群物件互動方式封裝起來的物件. 因為物件彼此不直接相互指涉, 所以耦合性低, 容易逐一變更互動關係

### [Interpreter 解譯器模式:](https://github.com/kimi0230/DesignPatternGolang/tree/master/Interpreter) 
針對標的語言定義文法, 以及可解讀這個語句的解譯器

### [Visitor 訪問者模式:](https://github.com/kimi0230/DesignPatternGolang/tree/master/Visitor) 
定義能逐一施行於物件結構裡各個元素的操作, 讓你不必修改作用對象的類別介面, 就能定義新的操作
