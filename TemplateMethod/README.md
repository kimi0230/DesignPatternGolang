### Template Method 範本方式模式: 由次類別決定如何實現一個演算法中的步驟
> 對於操作, 只先定義好演算法的輪廓, 某些步驟則留給子類別去填飽, 以便在不改變演算法整體構造的情況下讓子類別去精煉某些步驟
    
定義一個操作中的演算法的骨架, 而將一些步驟延遲到子類別中.
範本方式使得子類別可以不改變一個演算法的結構即可重定義該演算法的某些特定步驟

![UML](https://github.com/kimi0230/DesignPatternGolang/blob/master/UML/TemplateMethod.png?raw=true)