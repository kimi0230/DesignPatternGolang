### Facade 外觀模式：
> 替子系統裡的一堆介面定義一套統一的高階介面, 讓子系統更容易使用

為子系統中的一組介面提供一個一致的界面，此模式定義了一個高層介面，
這個介面使得這一子系統更加容易使用（投資：基金，股票，房產）

##### 外觀 vs. [轉接器](https://github.com/kimi0230/DesignPatternGolang/tree/master/Adapter)

外觀定義的是一個新的介面, 為現存系統停工一個更為方便的存取介面;
轉接則是復用一個原有的介面, 轉接器是使兩個已有的介面協同工作
外觀模式用來轉接整個子系統, 轉接用來轉接物件


##### 外觀  vs. [Mediator 中介者模式](https://github.com/kimi0230/DesignPatternGolang/tree/master/Mediator)、

外觀模式：外觀中保存了一堆對象,這些對像或者是組成某個子系統的,
將其封裝在外觀對像中,給客戶端一種只有一個對象的感覺,
中介者模式:
每個對像都保存一份中介者對象,
在和其他對象交互時,通過中介者來完成,
外觀模式是一結構型模式,
中介者模式是一行為性模式

![UML](https://github.com/kimi0230/DesignPatternGolang/blob/master/UML/Facade.png?raw=true)