package v1

import "github.com/gin-gonic/gin"

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) Get(c *gin.Context) {}

//@Summary 取得多個標籤
//@Produce json
//@Param name query string false "標籤名稱" maxlength(100)
//@Param state query int fasle "狀態" Enums(0,1) default(1)
//@Param page query int fasle "頁碼"
//@Param page_size query int false "每頁數量"
//@Success 200 {object} model.Tag "成功"
//@Failure 400 {object} errcode.Error "請求錯誤"
//@Failure 500 {object} errcode.Error "內部錯誤"
//Router /api/vi/tags [get]
func (t Tag) List(c *gin.Context) {}

//@Summary 新增標籤
//@Produce json
//@Param name body string true "標籤名稱" minlength(3) maxlength(100)
//@Param state body int false "狀態" Enums(0,1) default(1)
//@param created_by body string true "建立者" minlength(3) maxlength(100)
//@Success 200 {object} model.Tag "成功"
//@Failure 400 {object} errcode.Error "請求錯誤"
//@Failure 500 {object} errcode.Error "內部錯誤"
//@Router /api/vi/tags [post]
func (t Tag) Create(c *gin.Context) {}

//@Summary 更新標籤
//@Produce json
//@Param id path int true "標籤ID"
//@Param name body string false "標籤名稱" minlength(3) maxlength(100)
//@Param state body int false "狀態" Enums(0,1) default(1)
//@Param modified_by body string true "修改者" minlength(3) maxlength(100)
//@Success 200 {array} model.Tag "成功"
//@Failure 400 {object} errcode.Error "請求錯誤"
//@Failure 500 {object} errcode.Error "內部錯誤"
//@Router /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {}

//@Summayr 刪除標籤
//@Produce json
//@Param id path int true "標籤ID"
//@Success 200 {string} string "成功"
//@Failure 400 {object} errcode.Error "請求錯誤"
//@Failure 500 {object} errcode.Error "內部錯誤"
//@Router /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {}
