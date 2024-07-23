package router

import (
	v1 "xwya/api/v1"

	"github.com/gin-gonic/gin"
)

func initDictionariesRouter(r *gin.RouterGroup) {
	dictionariesRouter := r.Group("/dictionaries")
	{
		dictionariesRouter.POST("generateDictionary", v1.GenerateDictionary)
		dictionariesRouter.POST("deleteDictionary", v1.DeleteDictionary)
		dictionariesRouter.POST("getDictionaryList", v1.GetDictionaryList)
		dictionariesRouter.GET("getDictionaryInfo", v1.GetDictionaryInfo)
		dictionariesRouter.GET("cs", v1.Cheshi)
	}
}
