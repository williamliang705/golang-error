# golang-error
原生 golang 所提供的 error 的內容，因為描述太過簡易，在系統運行及開發上可能會有無法快速定位問題的情況發生，以及 HTTP 或 GRPC 服務發生錯誤時狀態也未能對應，因此需要有一組定義 Error 的結構及功能來幫助我們解決。

## AppError
### Field
- code: 系統錯誤碼
- message: 描述錯誤訊息
- cause: 包裝原始發生的 error
- httpStatus: HTTP 對應狀態碼
- grpcStatus: GRPC 對應狀態碼
- stack: 錯誤堆疊

### Func
- Wrap: 包裝發生的 error
- Cause: 取得最出 error 原因
- Unify: 解析 Error 結構
- IsAppError: 判斷 error 為 appError 結構