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



