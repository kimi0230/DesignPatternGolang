### 簡單工廠 vs. 工廠方法
> 定義可滋生成物件的介面, 但讓子類別去決定該具現出哪一種類別的物件. 此模式讓類別將具現化程序交給子類別去處置
#### 簡單工廠
最大優點在於工廠類別中包含了必要的邏輯判斷, 根據用戶端的選擇條件動態實體化相關的類別
對用戶端來說, 去除了與具體產品的依賴
但是違背了 開放封閉原則

#### 工廠模式
用戶端需決定實體化哪一個工廠來實現運算類別, 選擇判斷的問題還是存在
工廠方法把簡單工廠的內部邏輯判斷移到了用戶端程式碼來進行
如果要加功能, 本來是改工廠類型, 現在是修改用戶端
是簡單工廠模式的進一步抽象.
缺點是每加一個產品就需要加一個產品工廠的類別, 增加額外的開發量