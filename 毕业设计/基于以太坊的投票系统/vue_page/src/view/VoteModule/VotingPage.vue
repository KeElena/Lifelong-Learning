<template>
  <div id="bg">
    <div style="margin: 50px"></div>
    <div id="content">
      <el-descriptions :column="1"  border size="default">
        <el-descriptions-item width="0" label-align="center" label="发起人">{{content.sponsor}}</el-descriptions-item>
        <el-descriptions-item width="0" label-align="center" label="标题" >{{content.title}}</el-descriptions-item>
        <el-descriptions-item width="0" label-align="center" label="投票内容">{{content.text}}</el-descriptions-item>
        <el-descriptions-item width="0" label-align="center" label="过期时间">{{expire}}</el-descriptions-item>
        <el-descriptions-item width="0" label-align="center" label="投票选项">
          <div id="options">
            <el-radio-group size="large" v-model="choose">
              <el-radio v-for="(option,idx) in content.options" :label="idx" :key="idx" :disabled="btnStatus" border>{{option}}</el-radio>
            </el-radio-group>
          </div>
        </el-descriptions-item>
        <el-descriptions-item width="0" label-align="center" label="投票情况">
          <div  v-for="(num,idx) in content.result" :key="idx" >
            {{content.options[idx]}}
            <el-progress :text-inside="true" :stroke-width="20" status="success" :percentage="Math.floor(num/count*100)"></el-progress>
          </div>
        </el-descriptions-item>
      </el-descriptions>
    </div>

    <div>
      <el-button class="btn" type="primary" @click="toVote">返回</el-button>
      <el-button class="btn" type="primary" @click="getProve">我的投票</el-button>
      <el-button class="btn" type="primary" @click="submit" :disabled="btnStatus">提交</el-button>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  data(){
    return{
      content: {
        sponsor:"",
        title:"",
        text:"",
        expire:"",
        options:null,
        result:null
      },
      choose:null,
      prove:{
        idx:null,
        voted:true,
        time:null
      },
      btnStatus:true
    }
  },
  created() {
    //请求数据
    axios.post("/vote/votecontent",{addr:this.$route.params.addr}).then((res)=>{
      if(res.data.msg){
        //内容
        this.content=JSON.parse(res.data.content)
        //投票票据
        this.prove=JSON.parse(res.data.prove)
        this.btnStatus=this.prove.voted
      }else{
        this.$message({
          message:"系统错误",
          type:"error"
        });
      }
    }).catch(()=>{
      this.$message({
        message:"连接失败",
        type:"error"
      });
      this.$router.push("/login");
    })
  },
  methods:{
    toVote:function (){
      this.$router.push("/vote")
    },
    submit:function () {
      if (this.choose==null){
        this.$message({
          message:"无选项",
          type:"warning"
        })
        return
      }
      axios.defaults.withCredentials=true;
      axios.post("/vote/submit",{choose:this.choose,addr:this.$route.params.addr,title:this.content.title}).then((res)=>{
        this.btnStatus=true
        if(res.data.msg){
          this.$message({
            message:"投票成功",
            type:"success"
          })
        }else{
          this.$message({
            message:"投票失败",
            type:"error"
          })
        }
      }).catch(()=> {
        this.$message({
          message: "服务器处理错误",
          type: "error"
        })
        this.$router.push("/login");
      })
    },
    getProve:function () {
      axios.post("/vote/getprove",{addr:this.$route.params.addr}).then((res)=>{
        if(res.data.msg){
          this.prove=JSON.parse(res.data.prove)
        }else{
          this.$message({
            message:res.data.msg,
            type:"error"
          });
        }
      })
      if(this.prove.idx==null){
        this.$router.push("/login");
        return
      }
      if(this.prove.voted){
        this.$alert("<p>投票选项："+this.content.options[this.prove.idx]+"</p><p>投票状态："+this.prove.voted+"</p>"+"<p>投票时间："+this.proveTime+"</p>","我的投票",{
          confirmButtonText:"确定",
          dangerouslyUseHTMLString:true
        });
      }else{
        this.$alert("<p>投票状态："+this.prove.voted+"</p>","我的投票",{
          confirmButtonText:"确定",
          dangerouslyUseHTMLString:true
        });
      }
    }
  },
  computed:{
    count:function (){
      let count=0
      for(let i=0;i<this.content.result.length;i++){
        count+=this.content.result[i]
      }
      return count
    },
    expire:function () {
      let date=new Date(parseInt(this.content.expire))
      return date.toLocaleDateString().replace(/\//g,"-")+" "+date.toTimeString().substring(0,8)
    },
    proveTime:function () {
      let date=new Date(parseInt(this.prove.time))
      return date.toLocaleDateString().replace(/\//g,"-")+" "+date.toTimeString().substring(0,8)
    }
  }
}
</script>

<style scoped>
#bg{
  min-width: 1000px;
  display: inline-block;
}
#content{
  width:900px;
  padding: 20px;
  border: 2px dashed deepskyblue;
}
#options{
  margin-top: 20px;
}
.btn{
  width: 100px;
  margin: 20px;
}
</style>