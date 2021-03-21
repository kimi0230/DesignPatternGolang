### Factory Method 工廠方法模式： 
> 由次類別決定要建立的具象類別為何者
	
定義一個用於創建對象的接口，讓子類決定實例化哪一個類
是一種管理物件創建的模式，隨著輸入的參數不同，簡單工廠會回傳不同的物件
使用者取得物件的時候只要傳入正確的參數，不需要去理解這個物件
是一個使一個類的實例化延遲到其子類

所有應用簡單工廠的地方, 都可以老綠用反射技術來去除 switch 或 if 解除分支判斷帶來的耦合

> Ex: 計算機

![UML](https://github.com/kimi0230/DesignPatternGolang/blob/master/UML/Factory.png?raw=true)