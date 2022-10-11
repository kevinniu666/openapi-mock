package main
import(
	"github.com/gin-gonic/gin"
	"net/http"
)
func main() {
	app := gin.Default()
	app.POST("/api/v2/study/create/url", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "请求/api/v2/study/create/url成功",
		})
	})
	app.POST("/api/v2/mutual/project/notice", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "请求/api/v2/mutual/project/notice成功",
		})
	})
	app.POST("/api/v2/study/shorttime/list/total", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "请求/api/v2/study/shorttime/list/total成功",
		})
	})
	app.POST("/api/v2/mutual/project/callback", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "请求/api/v2/mutual/project/callback成功",
		})
	})
	app.Run(":8888")

}
