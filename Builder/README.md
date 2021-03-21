### Builder 建造者模式:
> 從複雜物件的佈局中抽取出生成程序, 以便用同一個生成程序製造各種不同的物件佈局

將一個複雜物件的構建與他的表示分離, 使得同樣的構建過程可以建立不同的表示

建造者的建造流程是在指揮者(Director)中，指揮者在用戶通知他現在具體的建造者是誰後，
建造出對應的產品，建造者中實現了產品的建造細節

![UML](https://github.com/kimi0230/DesignPatternGolang/blob/master/UML/Builder.png?raw=true)
來源: 程杰-大話設計模式