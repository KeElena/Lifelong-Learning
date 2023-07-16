//SPDX-License-Identifier:MIT
pragma solidity ^0.8.17;

contract VoteSystem{
    //----------------数据结构-------------------//
	struct Vote{
		//是否投过
		bool voted;
		//投给谁,序号
		uint8 idx;
        //投票处理时间
        uint64 time;
	}
    //投票发起人
    address Sponsor;
    //投票选项
    string[] VotingOptions;
    //投票计数情况
    uint32[] voting;
    //初始化结束时间，建议秒为单位
    uint64 expireTime;
    //投票仓库，用于获取投票和保存投票票据
    mapping(address=>Vote) VoteStore;
    //-------------初始化合约----------------//
    constructor(){
        //发起人为部署合约的地址
        Sponsor=msg.sender;
    }
    //-------------票内容的初始化-------------//
    //初始化投票信息
    function initVotingOptions(string[] memory contents,uint64 et)external returns(bool){
        require(msg.sender==Sponsor);
        //要求合约没有初始化
        require(expireTime==0);
        //要求et大于区块当前时间
        require(et>block.timestamp);
        expireTime=et;
        for(uint8 i=0;i<contents.length;i++){
            VotingOptions.push(contents[i]);
            voting.push(0);
        }
        return true;
    }
    //-------------投票----------------------//
    function vote(uint8 idx)external returns(bool){
        //票仓中取出票据，没有参与投票的默认为初始值
        Vote storage ballot=VoteStore[msg.sender];
        //过滤过期且无效的选票
        require(block.timestamp<expireTime);
        //要求票没用过
        require(!ballot.voted);
        //修改票据信息
        ballot.idx=idx;
        ballot.voted=true;
        ballot.time=uint64(block.timestamp);
        //计数
        voting[idx]+=1;
        return true;
    }
    //--------------相关的get方法------------//
        //获取投票选项
    function getVotingInfo()external view returns(string[] memory,uint32[] memory,uint64){
        return (VotingOptions,voting,expireTime);
    }
        //获取票据信息
    function getVote()external view returns(Vote memory){
        return VoteStore[msg.sender];
    }
}