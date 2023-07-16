<template>
<div>
  <div>
    <p>
      帐号：<input type="text" v-model="account.account">
    </p>
    <p>
      密码：<input type="password" v-model="account.password">
    </p>
    <button @click="login">登录</button>
  </div>
  <div>
    <button @click="getMsg">发送含jwt信息的请求</button>
  </div>
</div>
</template>

<script setup>
import {reactive} from "vue";
import axios from "axios";
let account=reactive({account:null,password:null})
function login() {
  axios.post("/login",{account:account.account,password:account.password}).then((res)=>{
    localStorage.setItem("token",res.data.token)
    alert(res.data.msg)
  }).catch(()=>{
    alert("连接错误")
  })
}
function getMsg(){
  axios.interceptors.request.use(config=>{
    let token=localStorage.getItem('token')
    if (token){
      config.headers.Authorization=token
    }
    return config
  })
  axios.get("/get").then((res)=>{
    alert(res.data.msg)
  }).catch(()=>{
    alert("连接错误")
  })
}
</script>

<style>

</style>
