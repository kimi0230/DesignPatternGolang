### Adapter 轉接器模式:
> 將類別的介面轉換成外界所預期的另一種介面, 讓原先囿於見面不相容問題而無法協力合作的類別能夠兜在一起用

將一個類別的介面轉換成客戶希望的另外一個介面, Adapter 模式始原本由於介面不相容而不能一起工作的那些類別可以一起工作

與 Proxy 代理模式相似
* [Proxy 代理模式:](https://github.com/kimi0230/DesignPatternGolang/tree/master/Proxy)
    使用介面, 客戶端不知道代理對象
* [ Adapter 轉接器模式:](https://github.com/kimi0230/DesignPatternGolang/tree/master/Adapter)
    可以是具體或抽象類別,或是介面, 對象的接口和客户端想要的不一樣, Adapter將接口封装一下, 改成客戶端想要的接口

![UML](https://github.com/kimi0230/DesignPatternGolang/blob/master/UML/Adapter.png?raw=true)