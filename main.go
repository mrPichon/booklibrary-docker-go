package main

import (
	"isrmino/dockerapp/logging"
	"isrmino/dockerapp/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var config *util.Config

func init() {
	logging.SetUpLogging()
	myconfig, err := util.LoadConfig(".")
	if err != nil {
		logging.AppLog.WriteLogsError("cannot load config:", map[string]interface{}{"source": config, "err": err})
		log.Fatal("cannot load config: ", err)
	}

	config = &myconfig
	// load book collection
	// util.LoadCollection("./static/")
	util.InitializeRedis()
}

func main() {
	logging.AppLog.WriteLogsInfo("Application running in environment:",
		map[string]interface{}{"runtime_setup": viper.GetString("RUNTIME_SETUP"), "app_port": viper.GetInt("APP_PORT")})

	router := gin.Default()

	router.Static("/static/", "./static")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/collection", getCollection)
	/* router.GET("/collection", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "library.html", gin.H{
			"books": util.BookCollection.BookList,
		})
	}) */

	router.Run(viper.GetString("SERVER_ADDRESS") + ":" + viper.GetString("APP_PORT"))
	// router.Run(":8080")
}

func getCollection(ctx *gin.Context) {
	val := util.GetBookList()
	ctx.HTML(http.StatusOK, "library.html", gin.H{
		"books": val.BookList,
	})
}
