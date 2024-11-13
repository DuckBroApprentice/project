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
