package global

import "github.com/go-programming-tour-book/blog-service/pkg/setting"

//對最初預估的三個區段進行了設定並宣告了全域變數
//註：全域變數的初始化會隨著應用程式的不斷演進而改變
var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
)
