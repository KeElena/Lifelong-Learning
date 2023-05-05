<template>
  <div id="login">
    <div style="margin: 20px"></div>
    <el-form label-position="right" label-width="80px" size="large">
      <el-form-item label="手机号">
        <el-input style="width:300px" placeholder="输入手机号码" v-model="phone" onkeyup="this.value=this.value.replace(/\D/g,'')"/>
      </el-form-item>
      <el-form-item label="密码">
        <el-input style="width:300px" placeholder="输入密码" v-model="password" show-password/>
      </el-form-item>
    </el-form>
    <div style="margin: 20px"></div>
    <div>
      <el-button type="primary" size="large" style="width: 100px;margin-bottom: 20px;" @click="login">登录</el-button>
      <div style="width: 20px;display: inline-block"></div>
      <el-button type="primary" size="large" style="width: 100px;margin-bottom: 20px;" @click="toRegister()">注册</el-button>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
export default {
  data(){
    return{
      phone:"",
      password:""
    }
  },
  methods:{
    toRegister:function (){
      this.$router.push("/register")
    },
    login:function (){
      axios.defaults.headers["Content-Type"]="application/json";
      axios.post("/account/login",{phone:this.phone,password:this.password}).then((res)=>{
        if (res.data.msg===true){
          this.$router.push("/vote")
        }else{
          this.$message({
            message:"手机号或密码输入错误",
            type:"error"
          });
        }
      }).catch((res)=>{
        this.$message({
          message:res.data.msg,
          type:"error"
        });
      })
    }
  }
}
</script>

<style scoped>
#login{
  width: 450px;
  font-weight: bolder;
  display: inline-block;
  text-align: center;
  border: 2px dashed deepskyblue;
  border-radius: 10px;
}
</style>