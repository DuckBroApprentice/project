一.入口建立HTTP服務
　解析
　　1.gin.Default

func Default() *Engine {
	debugPrintWARNINGDefault()
//會檢查Go版本是否達到gin的最低要求，再偵錯記錄檔[WARNING]Creating an Engine instance with the Logger and Recovery middleware already attached.的輸出
	engine := New()
	engine.Use(Logger(), Recovery())
	return engine
}

　　　建立預設的Engine實例，在初始化階段引用Logger和Recovery中介軟體，保證應用程式最基本的運作。
　　　Logger:輸出請求記錄檔，並標準化記錄檔的格式。
　　　Recovery:例外捕捉，即針對每次請求處理進行Recovery處理，防止因為出現panic導致服務當機，同時將例外記錄檔的格式標準化。
　　　
　　2.gin.New
　　　是最重要的方法，會對Engine實例執行初始化動作並傳回，在gin中承擔了「主軸」的作用。(介紹參數...)
　　3.gin.GET
　　　註冊路由。
　　　解析報錯 (main.mainfunc1(3 handlers))數量為何為3 ->r.Get("/ping"...)+Logger+Recovery
　　4.gin.Run
　　　解析位址，再呼叫http.ListenAndServe，將Engine實例作為Handler註冊進去，啟動服務，開始對外提供HTTP服務。
二.專案設計
New Skill:標準目錄建立、資料庫設計、資料模型撰寫、介面方法的設計、介面處理方法、啟動連線
	1.目錄結構
		configs:設定檔。
		docs:文件集合。
		global:全域變數。
		internal:內部模組。
			dao:資料存取層，所有與資料相關的操作都會在dao層進行，如MySQL,Elasticsearch等。
			middleware:HTTP中介軟體。
			model:模型層，用於儲存model物件。
			routers:路由相關的邏輯。
			service:專案核心業務邏輯。
		pkg:專案相關的模組套件。
		storage:專案產生的暫存檔案。
		scripts:各種建置、安裝、分析等操作的指令稿。
		thrid_party:協力廠商的資源工具，如Swagger UI。

	※熟悉code要寫在哪裡
	2.資料庫
		MySQL:
		建立資料庫(MySQL)
			建立本專案的資料庫blog service，並將它的預設編碼設定為utf8mb4。
		建立標籤表
			建立標籤表，表欄位主要為標籤的名稱、狀態和公共欄位。
		建立文章表
		建立文章標籤連結表
		Go:
		internal/model建立model.go,tag.go,article.go,article_tag.go

		路由
		RESTful API
		GET:讀取和檢索動作
		POST:新增和新增動作
		PUT:更新動作，用於更新一個完整的資源，要求為冪等
		PATCH:更新動作，用於更新某一個資源的組成部分
		DELETE:刪除動作
		Go:
		設置標籤及文章路由:apiv1.方法("<<路徑>>ex:/tags")
		Go:
		internal/routers/api/v1建立tag.go,article.go
		分別設置標籤方法及文章Handler方法 func(t Tag) 方法(c *gin.Context){}
		Go:
		路由管理--註冊
		internal/routers中router.go
		Go:
		撰寫入口main.go
3.公共元件
//本例採用檔案設定作為選型(go get viper)
	錯誤標準化
	pkg/errcode: common_code.go;errcode.go
	common_code.go:預先定義項目中的部分公共錯誤碼，以便引導和規定大家使用
	errcode.go:撰寫常用的一些錯誤處理公共方法
	預設設定
	configs:config.yaml
	Server:服務設定，設定gin的執行模式、預設的HTTP監聽通訊埠，允許讀取和寫入的最大持續時間
	App:應用設定，設定預設每頁數量、所允許的最大每頁數量，以及預設的應用記錄檔儲存路徑
	Database:資料庫設定，主要是連接實例所必需的基礎參數

	設定管理

	pkg/setting:setting.go
	對讀取設定的行為進行封裝
	NewSetting方法，用於初始化本專案設定的基礎屬性，即設定設定檔的名稱為config、設定類型
	為yaml，並且設定其設定路徑為相對路徑configs/，以確保在專案目錄下能夠成功啟動撰寫元件。
	
	pkg/setting:section.go
	撰寫讀取區段設定的設定方法
	
	global:setting.go:
	套件全域變數
	將設定資訊和應用程式連結
	對最初預估的三個區段(ServerSetting、AppSetting、DatabaseSetting)進行了設定
	並宣告了全域變數

	main.go:
	初始化設定讀取
	func init : 把設定檔內容對映到應用設定結構中的作用
4.資料庫連接
	internal/model:model.go
	新增NewDBEngine方法
	針對建立DB實例的NewDBEngine方法，同時增加了gorm開放原始碼函數庫的引用和
	MySQL驅動函數庫github.com/jinzhu/gorm/dialects/mysql的初始化
	註：不同類型的DBType需要引用不同的驅動函數庫

	global:db.go
	套件全域變數

	main.go
	初始化
	新增setupDBEngine方法的初始化

5.記錄檔標準化
	pkg/logger:logger.go
	深入了解func clone()的作用：
	建立一個新的Logger實例，同時保留原有Logger的大部分配置，但又不影響原始的實例。
		1.配置的繼承與擴展： 當我們想在原有 Logger 的基礎上添加新的字段或修改部分配置
		  時，就可以通過 clone() 來創建一個新的 Logger，然後對新實例進行修改。這樣，我
		  們就可以在不同的函數或模塊中使用不同的 Logger 實例，而不會相互影響。
		2.避免修改原始對象： 在函數調用中，我們通常希望避免修改傳入的參數。通過clone()
		  ，我們可以創建一個副本，然後對副本進行修改，這樣就能保證原始的 Logger 不會被
		  意外改變。
	題外話：
		在高並發的環境下，使用 clone() 方法時需要考慮線程安全的問題:
			在高併發環境下，多個 goroutine可能同時訪問並修改同一個變量。
			如果這些 goroutine 在訪問和修改變量的過程中沒有適當的同步機制，
			就可能導致數據競爭，進而引發不可預測的錯誤。
		可能的原因:
		1.共享的原始對象： 如果多個 goroutine 共享同一個原始 Logger 實例，
		  那麼它們對這個實例的修改就會影響到其他 goroutine。
		2.複製的時機： 如果在複製過程中，原始 Logger 的狀態發生了變化，
		  那麼複製出來的副本可能就不再是原始狀態的準確反映。
		3.字段的類型： 如果 Logger 結構體中包含了可變的 slice 或 map 等類型，
		  那麼即使進行了深拷貝，也可能存在線程安全問題。
		解決方法:
		1.避免共享原始對象： 每个 goroutine 都應該有自己獨立的 Logger 實例，避免共享。
		2.使用同步機制： 如果必須共享原始對象，可以使用互斥鎖 (mutex) 或
		  讀寫鎖 (RWLock) 等同步機制來保護對共享資源的訪問。
		3.使用不可變對象： 如果可能，將 Logger 結構體中的字段設計為不可變的，
		  這樣就可以避免多個 goroutine 對其進行修改。
		4.使用线程安全的容器： 如果 Logger 結構體中包含了可變的集合類型，可以使用
		   Go 語言提供的线程安全的容器，例如 sync.Map。
	global:setting.go
	記錄檔標準化:全域變數
	main.go
	初始化
