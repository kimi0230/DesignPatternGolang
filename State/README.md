### State 狀態模式:
> 讓物件的外顯行為隨內部狀態的改變而改變, 彷彿連類別也變了似的

當一個物件的內在狀態改變時, 允許改變其行為,
這個物件看起來像是改變了其類別

UML與策略模式相似
* 策略模式: 
    是用在對多個做同樣事情（統一接口）的類對象的選擇上
* 狀態模式：
    將對某個事情的處理過程抽象成接口和實現類的形式，
    由context保存一份state，在state實現類處理事情時，修改狀態傳遞給context， 由context繼續傳遞到下一個狀態處理中，

![UML](https://github.com/kimi0230/DesignPatternGolang/blob/master/UML/State.png?raw=true)