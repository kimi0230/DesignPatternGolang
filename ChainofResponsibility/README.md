### Chain of Responsibility 職責鏈模式:
> 讓多個物件都有機會處理某一訊息, 以降低訊息發送者和接收者之間的耦合關係. 它將接收者物件串連起來, 讓訊息流經其中, 直到被處理了為止

使多個物件都有機會處理請求, 從而避免請求的發送者和接收者之間的耦合關係. 將這個物件連成一條鏈, 並沿著這條鏈傳遞該請求, 直到有一個物件處理它為止


![UML](https://github.com/kimi0230/DesignPatternGolang/blob/master/UML/ChainofResponsibility.png?raw=true)
