# etcd

**一、etcd介绍**

* 分布式、高可用的一致性键值存储系统，提供可靠的分布式键值存储、配置共享和服务发现等功能
* 文档：[etcd](https://etcd.io/docs/v3.5/)

**二、etcd的特点**

* 简单：安装配置简单，而且提供了 HTTP API 进行交互，使用也很简单
* 键值对存储：将数据存储在分层组织的目录中，如同在标准文件系统中
* 监测变更：监测特定的键或目录以进行更改，并对值的更改做出反应
* 安全：支持 SSL 证书验证
* 快速：根据官方提供的 benchmark 数据，单实例支持每秒 2k+ 读操作
* 可靠：采用 raft 算法，实现分布式系统数据的可用性和一致性

**三、etcd使用场景**

* etcd一般用于处理控制数据
* 常见的使用场景：
  * 服务发现
  * 分布式锁
  * 分布式队列
  * 分布式通知和协调
  * 主备选举

**四、etcd与ZooKeeper**

* 服务发现的实现上，etcd使用的是节点租约，支持Group（多key），ZooKeeper使用的是临时节点，临时节点存在不少问题
* etcd支持稳定的watch，ZooKeeper使用的是简单的单次触发watch，在微服务环境下，通过调度系统的调度，一个服务随时可能会下线，也可能扩容增加新的服务节点，而调度系统是需要得到完整的节点历史记录，etcd可以存储数十万个历史变更
* etcd支持MVCC（多版本并发控制），因为有协同系统需要无锁操作
* etcd支持更大的数据规模，支持存储百万到千万级别的key
* etcd的性能比ZooKeeper更好，etcd3节点集群能实现每秒万级写操作，十万级的读操作

# etcd的部署

**一、docker-compose单节点一键部署**

```yaml
version: "3.8"
services:
  Etcd:
    image: bitnami/etcd:latest
    deploy:
      replicas: 1
#      restart_policy:
#        condition: on-failure
    environment:
      - ALLOW_NONE_AUTHENTICATION=yes
    privileged: true
    volumes:
      - ${PWD}/data:/bitnami/etcd/data
    ports:
      - "2379:2379"
      - "2380:2380"
```

* 常见问题："msg":"failed to start etcd","error":"cannot access data directory: open /bitnami/etcd/data/.touch: permission denied"
  * 解决方式：对挂载到主机的目录授权

```bash
mkdir ./data
chmod a+rwx ./data
docker-compose up -d
```

# etcd的crud操作

* 使用Restful形式的接口实现（类似Elasticsearch）

## 添加键值对

**一、添加键值对**

* 格式：`etcdctl put [key] [value] `
* key命名习惯：对于一种key常用`/`进行，用于隔离不同级别的key

```bash
etcdctl put /mykey myvalue
```

## 修改键值对

**一、修改键值对**

* 格式：`etcdctl put [key][new value]`
* 覆盖上一个版本的值
* 初始版本为1，每修改一次+1

```bash
etcdctl put /mykey new
```

## 查询键值对的值

**一、获取指定key的值**

* 格式：`etcdctl get [key]`

```bash
etcdctl get /mykey
```

**二、获取指定键值对并以十六进制形式返回值**

* 格式：`etcdctl get [key] --hex`

```bash
etcdctl get /mykey --hex
```

**三、获取key范围内所有存在的key和value**

* 格式：`etcdctl get [key1] [key2]`
* 区间：半开区间，左闭右开
* 输出key和value，对于不存在的键值对不输出

```bash
etcdctl get /mykey/0 /mykey/10
```

**四、获取匹配指定前缀的所有键值对的key和value**

* 格式：`etcdctl get --prefix [前缀] --limit [num]`
* `--limit`用于限制数量，默认不限制数量（即为不用加`--limit`）

```bash
etcdctl get --prefix /mykey --limit 2
```

**五、读取大于等于某个值的所有键值对的key和value**

* 格式：`etcdctl get --from-key [值]`
* 可以使用`--limit`限制条数

```bash
etcdctl get --from-key /mykey/0
```

## 删除键值对

**一、删除键值对**

* 格式：`etcdctl del [key]`

```bash
etcdctl del /mykey
```

## 查询版本

**一、获取键值对当前版本信息**

* 格式：`etcdctl get [key] -w=json`
* 以json形式输出

```bash
etcdctl get /mykey -w=json
```

| 字段       | 说明                     |
| ---------- | ------------------------ |
| cluster_id | 请求的etcd集群ID         |
| member_id  | 请求的etcd节点ID         |
| revision   | 全局数据版本号           |
| raft_term  | etcd当前raft主节点任期号 |

**二、访问键值对以前版本的信息**

* 格式：`etcdctl get --rev=[revision版本号] [key]`

```bash
etcdctl get --rev=95 mykey
```

# watch和lease

## watch监听

* 持续监听一个键值对，值发生变化时会输出最新的值并退出

**一、监听键值对**

* 格式：`etcdctl watch [key]`
* 输出修改键值对的操作、键名和新的值

```bash
etcdctl watch /key
```

## lease租约

* 类似redis的ttl机制，将etcd的键值对绑定到租约上实现存活周期的控制
* 租约过期后，所有绑定到该租约的键值对都删除

**一、创建租约**

* 格式：`etcdctl lease grant [second]`
* 返回租约ID

```bash
#创建30s的租约
etcdctl lease grant 30
```

**二、创建键值对并绑定到租约**

* 格式：`etcdctl put --lease=[租约ID] [key] [value]`
* 一个租约可以绑定多个键值对，到期后一起删除

```bash
etcdctl put --lease=[租约ID] key value
```

**三、撤销租约**

* 格式：`etcdctl lease revoke [租约ID]`

```bash
etcdctl lease revoke [租约ID]
```

**四、刷新租约**

* 格式：`etcdctl lease keep-alive [租约ID]`
* 一直刷新租约保持租约存活

```bash
etcdctl lease keep-alive [租约ID]
```

**五、查询租约剩余时间**

* 格式：`etcdctl lease timetolive [租约ID]`

```bash
etcdctl lease timetolive [租约ID]
```

**六、查询租约绑定的键值对**

* 格式：`etcdctl lease timetolive --keys [租约ID]`

# etcd权限管理

* etcd的权限管理是通过角色分配权限，使用用户绑定角色以获取权限

## 启动权限管理

**一、查看是否启动权限管理**

* 格式：`etcdctl auth status`
* `Authentication`值为false时表示没开启权限管理

**二、创建root用户**

* 格式：`etcdctl user add root`

**三、开启权限管理**

* 格式：`etcdctl auth enable`

## 用户

**一、创建用户**

* 格式：`etcdctl user add [用户名]`

**二、删除用户**

* 格式：`etcdctl user del [用户名]`

**三、修改密码**

* 格式：`etcdctl user passwd [用户名]`

**四、查看所有用户**

* 格式：`etcdctl user list`

**五、用户绑定角色**

* 格式：`etcdctl user grant-role [用户名] [角色名]`

**六、用户删除角色**

* 格式：`etcdctl user revoke-role [用户名] [角色名]`

**七、查看指定用户以及其绑定角色**

* 格式：`etcdctl user get [用户名]`

## 角色

**一、创建角色**

* 格式：`etcdctl role add [角色名]`

**二、给角色授权**

* 格式：`etcdctl role grant-permission [角色] [权限] [键名]`
* 添加`--prefix=true`，为角色授予键以某字符串开头的所有键值对的权限
* 格式：`etcdctl role grant-permission [角色] --prefix=true [权限] [前缀]`
* 权限：
  * readwrite：读写
  * read：读
  * write：写

**三、回收角色权限**

* 格式：`etcdctl role revoke-permission`

**四、删除角色**

* 格式：`etcdctl role delete [角色]`

## 用户查询或修改数据

**一、用户查询数据**

* 格式：`etcdctl get [key] --user=[用户]`

**二、用户修改数据**

* 格式·：`etcdctl get [key] --user=[用户]`







