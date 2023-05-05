<template>
  <div style="margin: 100px"></div>
  <div id="voteBG">
    <div class="vote" id="content">
      <el-table stripe :data="contents" style="width: 550px">
        <el-table-column prop="sponsor" label="发起人" width="100px" show-overflow-tooltip="true"></el-table-column>
        <el-table-column prop="title" label="投票标题" width="350px" show-overflow-tooltip="true"></el-table-column>
        <el-table-column label="查看详细">
          <template #default="scope">
            <el-button type="text" @click="check(scope.$index)" :disabled="btOneStatus(scope.$index)">查看</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <div id="card">
      <el-descriptions :column="1"  border size="default">
        <el-descriptions-item label="用户名" >{{userinfo.username}}</el-descriptions-item>
        <el-descriptions-item label="以太坊地址">{{userinfo.ethaddr}}</el-descriptions-item>
        <el-descriptions-item label="以太币数量">{{userinfo.ethbalance}}</el-descriptions-item>
      </el-descriptions>
    </div>
    <div id="btn">
      <el-button type="primary" @click="getBalance">获取以太币</el-button>
      <el-button type="primary" @click="toCreate">发起投票</el-button>
      <el-button type="primary" @click="freshBalance">刷新</el-button>
    </div>

    <div id="history" class="vote">
      <h3 style="text-align: left;margin: 0;">最近参与的投票:</h3>
      <el-table stripe :data="history" style="width: 100%">
        <el-table-column prop="title" label="投票标题" width="200px" show-overflow-tooltip="true"></el-table-column>
        <el-table-column prop="vote_time" label="参与时间" width="150px" show-overflow-tooltip="true"></el-table-column>
        <el-table-column label="查看情况" >
          <template #default="scope">
            <el-button type="text" @click="checkHistory(scope.$index)" :disabled="btTwoStatus(scope.$index)">查看</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
export default {
  data(){
    return{
      userinfo:{
        username:"          ",
        ethaddr:"            ",
        ethbalance:null
      },
      contents:null,
      history:null,
    }
  },
  created() {
    axios.defaults.withCredentials=true;
    axios.post("/vote/userinfo").then((res)=>{
      if(res.data.msg){
        this.userinfo=JSON.parse(res.data.userinfo)
        this.userinfo.ethbalance=(this.userinfo.ethbalance/1000000000000000000).toFixed(3)+"eth"
      }else{
        this.$message({
          message:"用户信息请求失败",
          type:"error"
        });
      }
    }).catch(()=>{
      this.$message({
        message:"请求失败",
        type:"error"
      });
      this.$router.push("/login");
    })

    axios.post("/vote/votelist").then((res)=>{
      if(res.data.msg===true){
        this.contents=JSON.parse(res.data.content)
      }else{
        this.$message({
          message:"请求失败",
          type:"error"
        });
      }
    }).catch(()=>{
      this.$message({
        message:"请求失败",
        type:"error"
      });
    })

    axios.post("/vote/history").then((res)=>{
      this.history=JSON.parse(res.data.list)
    }).catch(()=>{
      this.$message({
        message:"请求失败",
        type:"error"
      });
    })
  },
  methods:{
    check:function (idx){
      this.$router.push("/vote/voting/"+this.contents[idx].contractAddr)
    },
    toCreate:function () {
      this.$router.push("/vote/create")
    },
    getBalance:function () {
      axios.defaults.withCredentials=true;
      axios.post("/vote/getbalance").then((res)=>{
        if(res.data.msg){
          this.$message({
            message:"获取成功，以太币刷新有延迟",
            type:"success"
          });
        }else{
          this.$message({
            message:"获取失败",
            type:"error"
          });
        }
      }).catch(()=>{
        this.$message({
          message:"连接失败",
          type:"error"
        });
      })
    },
    freshBalance:function () {
      axios.defaults.withCredentials=true;
      axios.post("/vote/freshbalance").then((res)=>{
        if(res.data.msg){
          this.$message({
            message:"以太币刷新成功",
            type:"success"
          });
          this.userinfo.ethbalance=(res.data.balance/1000000000000000000).toFixed(3)+"eth"
        }
      }).catch(()=>{
        this.$message({
          message:"连接失败",
          type:"error"
        });
      })
    },
    btOneStatus:function (idx) {
      return this.contents[idx].contractAddr === "";
    },
    btTwoStatus:function (idx) {
      return this.history[idx].contract_addr === "";
    },
    checkHistory:function (idx) {
      this.$router.push("/vote/voting/"+this.history[idx].contract_addr)
    }
  }
}
</script>
<style scoped>
#voteBG{
  min-width: 900px;
  max-height: 475px;
  width: 1000px;
  display: inline-block;
}

#content{
  width: 550px;
  padding:20px;
  border: 2px dashed deepskyblue;
  border-radius: 10px;
  position: relative;
  left:-100px;
}
#card{
  width: 450px;
  padding:20px;
  border: 2px dashed deepskyblue;
  border-radius: 10px;
  position: relative;
  right: -520px;
  top:-473px;
}
#btn{
  position: relative;
  width: 494px;
  right: -397px;
  top:-460px;
}
#history{
  width: 450px;
  padding:20px;
  border: 2px dashed deepskyblue;
  border-radius: 10px;
  position: relative;
  right: -520px;
  top:-450px;
}
</style>