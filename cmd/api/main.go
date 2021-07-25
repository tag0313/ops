package main

import (
	"net/http"
	"ops/cmd/api/docs"
	"ops/pkg/daemon"
	"ops/pkg/model/consulreg"
	"strings"
	"time"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/unrolled/secure"
	"golang.org/x/crypto/acme/autocert"

	"ops/controller"
	"ops/pkg/logger"
	"ops/pkg/model"
	"ops/pkg/model/jwt"
	"ops/pkg/utils"
	"strconv"
)

// @title Opsnft API Interface
// @version 1.0
// @description 这是面向opsnft客户端的后端接口.
// @termsOfService http://swagger.io/terms/

// @contact.name Jack_Zhang
// @contact.url http://www.swagger.io/support
// @contact.email gongzhang67@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /api/v1.0
// @query.collection.format multi
// @schemes https

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationurl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationurl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

// @x-extension-openapi {"example": "value on a json format"}

//router used for unit test
func ginRouter(middleWare gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	// swagger
	docs.SwaggerInfo.Host = utils.GetConfigStr("tls_domain")
	url := ginSwagger.URL(utils.GetConfigStr("swagger_doc_url"))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r1 := router.Group("api/v1.0/login")
	{
		r1.POST("/message", controller.PostMessage)
		//解析从客户端发回的用户公钥跟随机码
		r1.POST("/token", controller.PostToken)
	}

	oCardStyle := router.Group("api/v1.0/style")
	{
		oCardStyle.GET("/:style-code", controller.GetBKGImage)
	}

	if middleWare != nil {
		router.Use(middleWare)
	}
	r2 := router.Group("api/v1.0/user")
	{
		r2.POST("/store", controller.StoreUserInfo)
		r2.PUT("/update", controller.UpdateUserInfo)
		r2.DELETE("/delete", controller.DeleteUserInfo)
		r2.POST("/query", controller.QueryUserInfo)
		r2.POST("/query/ops-account", controller.IsRepeatedOpsAccount)
		r2.POST("/private", controller.SetPrivate)
		r2.POST("/price", controller.SetPrice)
		r2.POST("/get-uid-by-account", controller.GetUidByOpsAccount)
	}

	r3 := router.Group("api/v1.0/oop")
	{
		r3.POST("/store", controller.StoreOop)
		r3.PUT("/update", controller.UpdateOop)
		r3.DELETE("/delete", controller.DeleteOop)
		r3.POST("/owner", controller.QueryOwnerOop)
		r3.POST("/other", controller.QueryOtherOop)
		r3.POST("/query", controller.QueryOop)
		r3.POST("/square", controller.SquareOop)
		r3.POST("/like", controller.LikeOop)
		r3.POST("/my-like-list", controller.MyLikeOop)
		r3.POST("/cancel-like", controller.CancelLikeOop)
	}

	r4 := router.Group("api/v1.0/follow")
	{
		r4.PUT("/following", controller.Following)
		r4.DELETE("/cancel", controller.CancelFollow)
		r4.POST("/oop", controller.QueryFollowOop)
		r4.POST("/following-list", controller.QueryFollowingList)
		r4.POST("/followed-list", controller.QueryFollowedList)
		r4.POST("/who-following-list", controller.WhoFollowingMe)
		r4.POST("/who-followed-list", controller.WhoFollowedMe)
	}

	r5 := router.Group("api/v1.0/contract")
	{
		r5.POST("/balance", controller.EthBalance)
		r5.POST("/get-eth-gas-fee", controller.GetGasFee)

		r5.POST("/ops/balance", controller.OpsBalance)
		//仅测试使用
		r5.POST("/ops/withdraw", controller.WithdrawOps)
		r5.POST("/ops/withdraw-fee", controller.WithDrawGasFee)
		r5.POST("/ops/subscribe-ops-recharge", controller.SubscribeOPSRecharge)
		r5.POST("/ops/info", controller.OpsInfo)

		//reading methods
		r5.POST("/nft1155/info", controller.NFT1155Info)
		r5.POST("/nft1155/balance-of", controller.NFT1155Balance)
		r5.POST("/nft1155/balance-of-batch", controller.NFT1155BalanceBatch)
		r5.POST("/nft1155/uri", controller.NFT1155URI)
		r5.POST("/nft1155/get-next-token-id", controller.NFT1155NextTokenID)
		r5.POST("/nft1155/is-approved-for-all", controller.NFT1155IsApprovedForAll)

		//writing methods
		// internal interfaces
		r5.POST("/nft1155/create", controller.NFT1155Create)
		r5.POST("/nft1155/create-batch", controller.NFT1155CreateBatch)
		r5.POST("/nft1155/create-batch-price", controller.NFT1155CreateBatchPrice)
		/*
			r5.PUT("/nft1155/set-base-meta-uri", controller.NFT1155SetBaseMetaURI)
			r5.POST("/nft1155/create", controller.NFT1155Create)
			r5.POST("/nft1155/create-batch", controller.NFT1155CreateBatch)
			r5.PUT("/nft1155/transferGovernorship", controller.NFT1155TransferGovernorship)
			r5.POST("/nft1155/mint", controller.NFT1155Mint)
			r5.POST("/nft1155/mint-batch", controller.NFT1155MintBatch)
			r5.PUT("/nft1155/set-creator", controller.NFT1155SetCreator)
			r5.PUT("/nft1155/set-id-uri", controller.NFT1155SetIdURI)
			r5.PUT("/nft1155/set-id-uri-batch", controller.NFT1155SetIdURIBatch)
		*/
	}

	r6 := router.Group("api/v1.0/search")
	{
		r6.POST("/id", controller.SearchID)
		r6.POST("/content", controller.SearchContent)
		r6.POST("/user", controller.SearchUser)
	}

	r7 := router.Group("api/v1.0/report")
	{
		r7.POST("/oop", controller.ReportOop)
		r7.POST("/user", controller.ReportUser)
		r7.POST("/times", controller.ReportTimes)
	}
	r8 := router.Group("api/v1.0/message")
	{
		r8.POST("/get-all", controller.GetAllMessage)
		r8.POST("/get", controller.GetMessage)
		r8.POST("/push", controller.PushMessage)
	}

	r9 := router.Group("api/v1.0/property")
	{
		r9.POST("/opspoint", controller.CheckOpsPoint)
		r9.POST("/deposit", controller.DepositOpspoint)
		r9.POST("/withdraw", controller.WithdrawOpspoint)
		subR9 := r9.Group("/ocard")
		subR9.POST("/query-mint", controller.QueryMintedOCard)
		subR9.POST("/query-ops", controller.QueryOCardOnOps)
		subR9.POST("/query-chain", controller.QueryOCardOnChain)
		subR9.POST("/mint-batch", controller.MintBatchCard)
		subR9.POST("/buy-ops", controller.BuyOCardOnOps)
		subR9.POST("/chain", controller.QueryOCardOnChain)
		//subR9.POST("/mint", controller.MintCard)
		subR9.POST("/transfer-batch", controller.TransferCards)

		subR9.POST("/query-transaction", controller.QueryTransaction)
		//subR9.POST("/charge-to-server", controller.ChargeToServer)
		subR9.POST("/query-transaction-fee", controller.QueryTransactionCardFee)
	}
	rlp := router.Group("api/v1.0/lp")
	{
		rlp.POST("/get-ops-usdt-apy", controller.GetOpsUsdtApy)
		rlp.POST("/get-ops-flux-apy", controller.GetOpsFluxApy)
		rlp.POST("/get-ops-price-usdt", controller.GetOpsPriceUsdt)
		rlp.POST("/get-ops-price-flux", controller.GetOpsPriceFlux)
		rlp.POST("/get-mining-pool-worth-usdt", controller.GetMiningPoolWorthUsdt)
		rlp.POST("/get-user-lp-info", controller.GetUserLPInfo)
	}

	return router
}

func main() {
	flag := daemon.NewCmdFlags()
	if err := flag.Parse(); err != nil {
		panic(err)
	}
	err := utils.LoadConfigFile(flag.ConfigFile)
	if err != nil {
		panic(err)
	}

	loggerOpt := logger.NewOpts(utils.GetConfigStr("log_path"))
	if gin.Mode() != gin.ReleaseMode {
		loggerOpt.Debug = true
	}

	logger.InitDefault(loggerOpt)
	defer logger.Sync()

	//Create service
	err = consulreg.InitMicro(utils.GetConfigStr("micro.addr"), "api")
	if err != nil {
		daemon.Exit(-1, err.Error())
	}
	controller.InitRpcCaller()

	router := ginRouter(JWTAuth())
	router.Use(ginLogger())

	if gin.Mode() == gin.ReleaseMode {
		logger.Info(gin.Version)
		m := autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist(utils.GetConfigStr("tls_domain")),
			Cache:      autocert.DirCache(utils.GetConfigStr("cert_folder")),
		}
		logger.Fatal(autotls.RunWithManager(router, &m))
	} else {
		port := utils.GetConfigInt("port")
		host := utils.GetConfigStr("tls_domain")
		var builder strings.Builder
		builder.WriteString(host)
		builder.WriteString(":")
		builder.WriteString(strconv.Itoa(port))
		router.Run(builder.String())
	}
}

func TlsHandler(port int) gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     ":" + strconv.Itoa(port),
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			return
		}
		c.Next()
	}
}

// JWTAuth 中间件，检查token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var result model.JSONResult
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusOK, result.NewError(utils.RECODE_TOKENERR))
			c.Abort()
			return
		}

		j := jwt.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == jwt.TokenExpired {
				c.JSON(http.StatusOK, result.NewError(utils.RECODE_TOKENEXPIRED))
				c.Abort()
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    err.Error(),
			})
			c.Abort()
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("claims", claims)
	}
}

func ginLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := utils.Now()
		// before request
		c.Next()
		// after request
		latency := time.Since(t)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUrl := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求ip
		clientIP := c.ClientIP()

		logger.APIAccess(reqMethod, clientIP, reqUrl, statusCode, c.Request.ContentLength,
			int64(c.Writer.Size()), t, latency)
	}

}
