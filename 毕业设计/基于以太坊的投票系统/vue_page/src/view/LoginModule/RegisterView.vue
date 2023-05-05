<template>
  <div id="border">
    <div style="margin: 20px">
      <h1>注册</h1>
    </div>
    <el-form label-position="right" label-width="90px" size="large">
      <el-form-item label="用户名">
        <el-input style="width:300px" placeholder="输入用户名" v-model="uname"/>
      </el-form-item>
      <el-form-item label="手机号">
        <el-input style="width:300px" placeholder="输入手机号" v-model="phone" onkeyup="this.value=this.value.replace(/\D/g,'')"/>
      </el-form-item>
      <el-form-item label="密码">
        <el-input style="width:300px" placeholder="输入密码" v-model="password" show-password/>
      </el-form-item>
      <el-form-item label="身份证号码">
        <el-input style="width:300px" placeholder="输入身份证号码" v-model="nuid" />
      </el-form-item>
    </el-form>
    <div>
      <el-button type="primary" size="large" style="margin-bottom: 20px;width: 100px" @click="toLogin">返回</el-button>
      <el-button type="primary" size="large" style="margin-bottom: 20px;width: 100px" @click="register">注册</el-button>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
export default {
  data(){
    return{
      uname:"",
      password:"",
      nuid:"",
      phone:""
    }
  },
  methods:{
    register:function () {
      let ok=true
      if(this.uname.length>15){
        this.$message({
          message:"用户名超过15个字符",
          type:"error"
        });
        ok=false
      }
      if(this.password.length>15){
        this.$message({
          message:"密码超过15个字符",
          type:"error"
        });
        ok=false
      }
      if(this.phone.length!==11){
        this.$message({
          message:"手机号长度错误",
          type:"error"
        });
        ok=false
      }
      if(this.nuid.length!==18){
        this.$message({
          message:"身份证长度错误",
          type:"error"
        });
        ok=false
      }
      if (!ok){
        return
      }
      axios.defaults.headers["Content-Type"]="application/json";
      axios.post("/account/register",{uname:this.uname,password:this.password,nuid:this.nuid,phone:this.phone}).then((res)=>{
        if(res.data.msg===true){
          this.$router.push("/vote")
        }
      }).catch(()=>{
        this.$message({
          message:"注册失败",
          type:"error"
        });
      })
    },
    toLogin:function (){
      this.$router.back()
    }
  }
}
</script>

<style scoped>
#border{
  width: 500px;
  border: 2px dashed deepskyblue;
  border-radius: 10px;
  font-weight: bolder;
  display: inline-block;
  text-align: center;
}
</style>