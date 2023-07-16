package vote

import (
	"BridgeModule/model"
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
)

type Vote struct {
	Sponsor      string `json:"sponsor"`
	Title        string `json:"title"`
	Text         string `json:"text"`
	ContractAddr string `json:"contractAddr"`
}

type UserInfo struct {
	Username   string `json:"username"`
	Ethaddr    string `json:"ethaddr"`
	Ethbalance int64  `json:"ethbalance"`
}

type JoinVote struct {
	Title        string `json:"title"`
	ContractAddr string `json:"contractAddr"`
}

type ContentVote struct {
	Sponsor string   `json:"sponsor"`
	Title   string   `json:"title"`
	Text    string   `json:"text"`
	Options []string `json:"options"`
	Result  []int    `json:"result"`
}

// GetUserInfo 获取用户信息卡
func GetUserInfo(c *gin.Context) {
	var userinfo UserInfo
	var err error
	//获取session数据
	session := sessions.Default(c)
	userinfo.Username = session.Get("username").(string)
	userinfo.Ethaddr = session.Get("ethaddr").(string)
	if len(userinfo.Ethaddr) != 0 {
		userinfo.Ethbalance, err = model.GetBalance(userinfo.Ethaddr)
	} else {
		id := session.Get("id").(int64)
		userinfo.Ethaddr, err = model.GetEthAddr(id)
		userinfo.Ethbalance = 0
	}
	if err != nil {
		log.Println(err)
		c.JSON(200, gin.H{"msg": false})
		return
	}
	//结构体序列化
	data, _ := json.Marshal(userinfo)
	c.JSON(200, gin.H{"msg": true, "userinfo": string(data)})
}

// GetETH 获取以太币
func GetETH(c *gin.Context) {
	var err error
	//获取session数据
	session := sessions.Default(c)
	to := session.Get("ethaddr").(string)
	//获取以太币
	err = model.GetETH(to)
	if err != nil {
		c.JSON(200, gin.H{"msg": false})
	}
	c.JSON(200, gin.H{"msg": true})
}

// FreshBalance 刷新以太币
func FreshBalance(c *gin.Context) {
	//获取session数据
	session := sessions.Default(c)
	addr := session.Get("ethaddr").(string)
	if len(addr) == 0 {
		id := session.Get("id").(int64)
		addr, _ = model.GetEthAddr(id)
		if len(addr) != 0 {
			ks, _ := model.GetKeystoreFromMysql(id)
			session.Set("keystore", ks)
			session.Set("ethaddr", addr)
			_ = session.Save()
			balance, err := model.GetBalance(addr)
			if err != nil {
				c.JSON(200, gin.H{"msg": false})
			}
			c.JSON(200, gin.H{"msg": true, "balance": balance, "ethaddr": addr})
		} else {
			c.JSON(200, gin.H{"msg": true, "balance": 0})
		}
		return
	}
	balance, err := model.GetBalance(addr)
	if err != nil {
		c.JSON(200, gin.H{"msg": false})
	}
	c.JSON(200, gin.H{"msg": true, "balance": balance})
}

// CreateVote 发起投票
func CreateVote(c *gin.Context) {
	var err error
	//获取session数据
	session := sessions.Default(c)
	password := session.Get("password").(string)
	keystore := session.Get("keystore").(string)
	uname := session.Get("username").(string)
	//检查session
	if len(password) == 0 {
		c.JSON(200, gin.H{"msg": false})
		return
	}
	//检查keystore
	if len(keystore) == 0 {
		id := session.Get("id").(int64)
		keystore, err = model.GetKeystoreFromMysql(id)
		if err != nil {
			log.Println(err)
			c.JSON(200, gin.H{"msg": false})
		}
	}
	//接收数据
	var content model.VoteContent
	err = c.Bind(&content)
	content.Sponsor = uname
	if err != nil {
		log.Println(err)
		c.JSON(200, gin.H{"msg": false})
		return
	}
	//创建投票并初始化内容
	var contractAddr string
	contractAddr, err = model.CreateVote(password, keystore, content.Options, uint64(content.Duration))
	if err != nil {
		log.Println(err)
		c.JSON(200, gin.H{"msg": false})
		return
	}
	content.ContractAddr = contractAddr
	//存储投票内容
	err = model.SetVoteContent(&content)
	if err != nil {
		log.Println(err)
		c.JSON(200, gin.H{"msg": false})
		return
	}
	c.JSON(200, gin.H{"msg": true})
}

// GetVoteList 获取投票列表
func GetVoteList(c *gin.Context) {
	list, err := model.GetVoteList()
	data, err := json.Marshal(list)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"msg": false})
		return
	}
	c.JSON(200, gin.H{"msg": true, "content": string(data)})
}

// GetVoteContent 获取投票内容
func GetVoteContent(c *gin.Context) {
	var err error
	//获取session数据
	session := sessions.Default(c)
	ethaddr := session.Get("ethaddr").(string)
	//获取数据
	uploadData, err := c.GetRawData()
	if err != nil {
		log.Println(err)
		c.JSON(200, gin.H{"msg": false})
		return
	}
	//反序列化
	var m map[string]interface{}
	_ = json.Unmarshal(uploadData, &m)
	content, err := model.GetDetailContent(m["addr"].(string))
	prove, err := model.GetVoteProve(ethaddr, m["addr"].(string))
	if err != nil {
		log.Println(err)
		c.JSON(200, gin.H{"msg": false})
		return
	}
	contentData, _ := json.Marshal(content)
	proveData, _ := json.Marshal(prove)
	if err != nil {
		log.Println(err)
		c.JSON(200, gin.H{"msg": false})
		return
	}
	c.JSON(200, gin.H{"msg": true, "content": string(contentData), "prove": string(proveData)})
}

// Submit 提交投票信息
func Submit(c *gin.Context) {
	var err error
	//获取session数据
	session := sessions.Default(c)
	id := session.Get("id").(int64)
	password := session.Get("password").(string)
	keystore := session.Get("keystore").(string)
	//检查session
	if len(password) == 0 {
		c.JSON(200, gin.H{"msg": false})
		return
	}
	//检查keystore
	if len(keystore) == 0 {
		id := session.Get("id").(int64)
		keystore, err = model.GetKeystoreFromMysql(id)
		if err != nil {
			log.Println(err)
			c.JSON(200, gin.H{"msg": false})
		}
	}
	//获取数据
	uploadData, err := c.GetRawData()
	if err != nil {
		log.Println(err)
		c.JSON(200, gin.H{"msg": false})
		return
	}
	//反序列化
	var m map[string]interface{}
	_ = json.Unmarshal(uploadData, &m)
	idx := m["choose"].(float64)
	err = model.Vote(password, keystore, m["addr"].(string), uint32(idx))
	if err != nil {
		log.Println(err)
		c.JSON(200, gin.H{"msg": false})
	}
	_ = model.SetUserHistory(id, m["addr"].(string), m["title"].(string))
	c.JSON(200, gin.H{"msg": true})
}

func GetProve(c *gin.Context) {
	var err error
	//获取session数据
	session := sessions.Default(c)
	ethaddr := session.Get("ethaddr").(string)
	//获取数据
	uploadData, err := c.GetRawData()
	if err != nil {
		log.Println(err)
		c.JSON(200, gin.H{"msg": false})
		return
	}
	//反序列化
	var m map[string]interface{}
	_ = json.Unmarshal(uploadData, &m)
	//获取数据
	prove, err := model.GetVoteProve(ethaddr, m["addr"].(string))
	if err != nil {
		log.Println(err)
		c.JSON(200, gin.H{"msg": false})
		return
	}
	//返回数据
	resData, err := json.Marshal(prove)
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(200, gin.H{"msg": true, "prove": string(resData)})
}

// GetHistory 获取参与历史
func GetHistory(c *gin.Context) {
	session := sessions.Default(c)
	id := session.Get("id").(int64)
	VoteList, err := model.GetHistory(id)
	if err != nil {
		log.Println(err)
		c.JSON(200, gin.H{"msg": false})
		return
	}
	data, _ := json.Marshal(VoteList)
	c.JSON(200, gin.H{"list": string(data)})
}
