### Observer 觀察者模式:
> 定義一對多的物件依存關係, 讓物件狀態一有變動, 就自動通知其他相依物件做該做的更新動作

定義了一種一對多的依賴關係, 讓多個觀察者物件同時監聽某一個主題物件. 
這主題物件在狀態發生變化時,會通知所有觀察者物件, 
使他們能夠自動更新自己

![UML](https://github.com/kimi0230/DesignPatternGolang/blob/master/UML/Observer.png?raw=true)