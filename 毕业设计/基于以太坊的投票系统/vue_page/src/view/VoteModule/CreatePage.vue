<template>
  <div id="bg">
    <h1>发起投票</h1>
    <el-form id="voteForm" label-position="right" size="large" ref="dynamicForm" label-width="90px">
      <el-form-item label="标题">
        <el-input style="width:500px" v-model="title" maxlength="30" show-word-limit/>
      </el-form-item>

      <el-form-item label="投票内容">
        <el-input style="width:500px" v-model="text" type="textarea" placeholder="输入投票内容" :rows="5" maxlength="1000" show-word-limit/>
      </el-form-item>

      <el-form-item
        v-for="(option,idx) in options"
        :key="idx"
        :label="`选项`+(idx+1)">
        <el-input v-model="options[idx]" style="width:500px" maxlength="30" show-word-limit/>
        <el-button type="primary" @click.prevent="removeOption(idx)">删除选项</el-button>
      </el-form-item>

      <el-form-item>
        <el-button type="primary" @click="addOption">添加选项</el-button>
        <el-button type="primary" @click="resetOptions">重置</el-button>
      </el-form-item>

      <div id="duration">
        <el-select style="text-align:left" placeholder="选择投票持续时间" v-model="duration">
          <el-option v-for="option in selectTime" :key="option" :label="option+`天`" :value="option"></el-option>
        </el-select>
      </div>

    </el-form>
    <div></div>
    <el-button type="primary" size="large" style="width: 100px;margin-top: 20px" @click="submit">提交</el-button>
    <el-button type="primary" size="large" style="width: 100px;margin-top: 20px" @click="toVote">返回</el-button>
  </div>
</template>

<style scoped>
#bg{
  min-width: 900px;
  max-height: 475px;
  width: 1000px;
  display: inline-block;
}
#voteForm{
  display: inline-block;
  text-align: center;
}
#duration{
  position: relative;
  left: -140px;
}
</style>

<script>
import axios from 'axios'
export default {
  data(){
    return{
      title:"",
      text:"",
      options:["",""],
      selectTime:[1,3,7],
      duration:null
    }
  },
  methods:{
    removeOption:function (idx) {
      if(this.options.length<=2){
        this.$message({
          message:"选项至少有两项",
          type:"error"
        });
      }else{
        this.options.splice(idx,1)
      }
    },
    addOption:function () {
      if(this.options.length>=5){
        this.$message({
          message:"选项最多五项",
          type:"warning"
        });
      }else{
        this.options.push("")
      }
    },
    resetOptions:function () {
      this.options=["",""]
    },
    submit:function (){
      let okSend=true
      if(this.title.length===0){
        this.$message({
          message:"标题为空",
          type:"error"
        });
        okSend=false
      }
      if(this.text.length===0){
        this.$message({
          message:"投票内容为空",
          type:"error"
        });
        okSend=false
      }
      for (let i = 0; i < this.options.length; i++) {
        if(this.options[i].length===0){
          this.$message({
            message:"选项"+(i+1)+"内容为空",
            type:"error"
          });
          okSend=false
        }
      }
      if(this.duration==null){
        this.$message({
          message:"未选择持续时间",
          type:"error"
        });
        okSend=false
      }

      if (okSend){
        axios.post("/vote/create",{title:this.title,text:this.text,options:this.options,duration:this.duration}).then((res)=>{
          if(res.data.msg===true){
            this.$message({
              message:"创建成功",
              type:"success"
            });
            this.$router.back();
          }else{
            this.$message({
              message:"创建失败",
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
      }
    },
    toVote:function () {
      this.$router.back();
    }
  },
}
</script>
