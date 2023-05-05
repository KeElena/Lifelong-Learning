package router

import (
	"BridgeModule/controller/account"
	"BridgeModule/controller/vote"
	"github.com/gin-gonic/gin"
)

func InitRouter(server *gin.Engine) {
	AccountGroup := server.Group("/account")
	{
		AccountGroup.POST("/register", account.Register)
		AccountGroup.POST("/login", account.Login)
	}
	VoteGrounp := server.Group("/vote")
	{
		VoteGrounp.POST("/userinfo", vote.GetUserInfo)
		VoteGrounp.POST("/freshbalance", vote.FreshBalance)
		VoteGrounp.POST("/votelist", vote.GetVoteList)
		VoteGrounp.POST("/history", vote.GetHistory)
		VoteGrounp.POST("/create", vote.CreateVote)
		VoteGrounp.POST("/getbalance", vote.GetETH)
		VoteGrounp.POST("/votecontent", vote.GetVoteContent)
		VoteGrounp.POST("/submit", vote.Submit)
		VoteGrounp.POST("/getprove", vote.GetProve)
	}
}
