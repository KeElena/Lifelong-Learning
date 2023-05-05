<template>
  <div>
    <h1 id="title">基于区块链的投票系统</h1>
    <div>
      <router-view ></router-view>
    </div>
    <div>
      <h3 id="maker">made by 关兆源</h3>
    </div>
  </div>
</template>

<style>
#title{
  margin-top: 200px;
  font-weight: bold;
}
#maker{
  position: fixed;
  bottom: 0;
  text-align: center;
  display: inline;
}
</style>
