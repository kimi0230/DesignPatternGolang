### Proxy 代理模式： 
> 替其他物件預留代理者空位, 就此控制存取其他物件

將物件包裝起來, 以控制對此物件的存取
為其他物件提供一種代理, 以控制對這個物件的存取
代理和代理的對象接口一致，客户端不知道代理對象

##### 代理 vs. [外觀](https://github.com/kimi0230/DesignPatternGolang/tree/master/Facade)
    
代理物件代表一個單一物件, 而外觀物件代表一個子系統. 
代理的客戶無法直接存取目標物件, 
由代理提供的對單獨之目標物件的存取控制, 
而外觀的客戶物件可以直接存取子系統中的各個物件, 
但通常由外觀物件提供對子系統各元件功能的簡化的共同層次的調用介面

##### 代理 vs. [轉接器](https://github.com/kimi0230/DesignPatternGolang/tree/master/Adapter)

都屬於一種銜接性質的功能.
代理是一種原來物件的代表, 其他需要與這個物件打交道的操作都是和這個代表交涉
轉接器則不需要虛構出一個代表者, 只需要為應付特定使用目的, 將原來的類別進行一些組合

![UML](https://github.com/kimi0230/DesignPatternGolang/blob/master/UML/Proxy.png?raw=true)