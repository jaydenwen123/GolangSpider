# 从docker到istio之三 - kubernetes集群编排 #

## 前言 ##

容器化，云原生越演越烈，新概念非常之多。信息爆炸的同时，带来层层迷雾。我尝试从扩容出发理解其脉路，经过实践探索，整理形成一个入门教程，包括下面四篇文章。

* [从docker到istio之一 - 使用Docker将应用容器化]( https://link.juejin.im?target=https%3A%2F%2Fgame404.github.io%2Fpost%2Fdocker2istio-docker%2F )
* [从docker到istio之二 - 使用compose部署应用]( https://link.juejin.im?target=https%3A%2F%2Fgame404.github.io%2Fpost%2Fdocker2istio-compose%2F )
* 从docker到istio之三 - kubernetes编排应用
* 从docker到istio之四 - istio管理应用

这是第三篇，kubernetes编排应用。

## kubernetes ##

> 
> 
> 
> Kubernetes是一个开源的，用于管理云平台中多个主机上的容器化的应用，Kubernetes的目标是让部署容器化的应用简单并且高效（powerful）,Kubernetes提供了应用部署，规划，更新，维护的一种机制。
> 
> 
> 

> 
> 
> 
> Kubernetes在希腊语中意思是船长或领航员，这也恰好与它在容器集群管理中的作用吻合，即作为装载了集装箱（Container）的众多货船的指挥者，负担着全局调度和运行监控的职责。因为Kubernetes在k和s之间有8个字母，所以又简称k8s
> 
> 
> 

快速体验k8s，可以使用Docker for mac中集成的k8s。

![DockerForMac](https://user-gold-cdn.xitu.io/2019/6/5/16b285c518d3efd2?imageView2/0/w/1280/h/960/ignore-error/1)

启动k8s后，等待其初始化完成，然后 ` docker ps` 可以看到k8s启动了一系列的容器:

` CONTAINER ID IMAGE COMMAND CREATED STATUS PORTS NAMES 17a693617137 docker/kube-compose-controller "/compose-controller…" 3 days ago Up 3 days k8s_compose_compose-74649b4db6-szsqz_docker_4f5997b7-5c47-11e9-95b9-025000000001_0 a9b666b48815 docker/kube-compose-api-server "/api-server --kubec…" 3 days ago Up 3 days k8s_compose_compose-api-5d754cdd89-ncwrq_docker_131b4d65-04e7-11e9-837c-025000000001_0 f4b05eefc73a 6f7f2dc7fab5 "/sidecar --v=2 --lo…" 3 days ago Up 3 days k8s_sidecar_kube-dns-86f4d74b45-zh6qc_kube-system_f669bc59-04e6-11e9-837c-025000000001_0 867f8f040258 c2ce1ffb51ed "/dnsmasq-nanny -v=2…" 3 days ago Up 3 days k8s_dnsmasq_kube-dns-86f4d74b45-zh6qc_kube-system_f669bc59-04e6-11e9-837c-025000000001_0 17f26a6e91d2 80cc5ea4b547 "/kube-dns --domain=…" 3 days ago Up 3 days k8s_kubedns_kube-dns-86f4d74b45-zh6qc_kube-system_f669bc59-04e6-11e9-837c-025000000001_0 ... 复制代码`

` kubectl version` 查看集群版本:

` Client Version: version.Info{Major: "1" , Minor: "10" , GitVersion: "v1.10.11" , GitCommit: "637c7e288581ee40ab4ca210618a89a555b6e7e9" , GitTreeState: "clean" , BuildDate: "2018-11-26T14:38:32Z" , GoVersion: "go1.9.3" , Compiler: "gc" , Platform: "darwin/amd64" } Server Version: version.Info{Major: "1" , Minor: "10" , GitVersion: "v1.10.11" , GitCommit: "637c7e288581ee40ab4ca210618a89a555b6e7e9" , GitTreeState: "clean" , BuildDate: "2018-11-26T14:25:46Z" , GoVersion: "go1.9.3" , Compiler: "gc" , Platform: "linux/amd64" } 复制代码`

` kubectl get nodes` 查看k8s集群节点:

` NAME STATUS ROLES AGE VERSION docker-for-desktop Ready master 123d v1.10.11 复制代码`

` kubectl get service` 查看k8s默认启动的服务:

` NAME TYPE CLUSTER-IP EXTERNAL-IP PORT(S) AGE kubernetes ClusterIP 10.96.0.1 <none> 443/TCP 123d 复制代码`

## 部署应用及测试 ##

### 编写应用部署文件 ###

#### 1. flaskapp文件 ` k8s/flaskapp.yaml` ####

` apiVersion: v1 kind: Service metadata: name: flaskapp spec: ports: - port: 5000 selector: name: flaskapp --- apiVersion: extensions/v1beta1 kind: Deployment metadata: name: flaskapp spec: replicas: 1 template: metadata: labels: name: flaskapp spec: containers: - image: flaskapp:0.0.2 name: flaskapp ports: - containerPort: 5000 复制代码`

了解这个部署文件，需要先大概了解一下k8s的运作方式。k8s通过api server提供restful接口，用于集群交互。每一个部署对象，都有 ` apiVersion` ， ` kind` , ` metadata` , ` spec` 这几个关键字。

* 定义了Service和Deployment2个类型的对象。Service表示k8s对外提供的服务，Deployment表示某个service的部署方式。
* Service对象的ports描述了服务端口，这个是集群内部网络的端口。
* Service对象的selector描述了服务如何选择对于的部署，采用标签 **name: flaskapp** ,这是一种解耦合的依赖关系。
* Deployment的replicas描述了容器的副本个数，下文会演示如何扩充。
* Deployment的containers描述了镜像名称，服务端口等。

#### 2. redis服务文件 ` k8s/redis.yaml` ####

` apiVersion: v1 kind: Service metadata: name: redis spec: ports: - port: 6379 selector: name: redis --- apiVersion: extensions/v1beta1 kind: Deployment metadata: name: redis spec: replicas: 1 template: metadata: labels: name: redis spec: containers: - image: redis:4-alpine3.8 name: redis ports: - containerPort: 6379 复制代码`

redis的部署文件和flaskapp的部署文件类似。

#### 3. nginx服务文件 ` k8s/nginx.yaml` ####

` kind: ConfigMap apiVersion: v1 metadata: name: nginx-config data: default.conf: | upstream flaskapp { server flaskapp:5000; } server { listen 80; server_name localhost; root /usr/share/nginx/html; location / { proxy_pass http://flaskapp; proxy_set_header X-Real-IP $remote_addr ; proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for ; proxy_set_header Host $host ; proxy_redirect off; } } --- apiVersion: v1 kind: Service metadata: name: nginx spec: ports: - port: 80 selector: name: nginx type : NodePort --- apiVersion: extensions/v1beta1 kind: Deployment metadata: name: nginx spec: replicas: 1 template: metadata: labels: name: nginx spec: containers: - image: nginx:1.15.8-alpine name: nginx ports: - containerPort: 80 volumeMounts: - name: nginx-config-volume mountPath: /etc/nginx/conf.d/default.conf subPath: default.conf volumes: - name: nginx-config-volume configMap: name: nginx-config 复制代码`

nginx的部署文件，变化在:

* 多出了ConfigMap对象，这个对象主要定义了nginx.conf文件，其内容和 ` nginx\default.conf` 一致。
* nginx的container中mount了一个configmap对象作为nginx的配置文件。

### 部署应用到集群 ###

使用 ` kubectl apply -f k8s` 命令将编写yaml文件提交到k8s集群，集群会自动根据yaml文件的声明，进行部署。

` service "flaskapp" created deployment.extensions "flaskapp" created configmap "nginx-config" created service "nginx" created deployment.extensions "nginx" created service "redis" created deployment.extensions "redis" created 复制代码`
> 
> 
> 
> 
> 这里的 ` kubectl apply -f k8s` 表示将k8s目录下的文件都提交给k8s集群。当然，也可以逐个文件提交 ` kubectl
> apply -f k8s/redis.yaml` 。
> 
> 

### 访问应用 ###

先 ` kubectl get service` 检查一下k8s内的服务:

` NAME TYPE CLUSTER-IP EXTERNAL-IP PORT(S) AGE flaskapp ClusterIP 10.110.202.47 <none> 5000/TCP 31s kubernetes ClusterIP 10.96.0.1 <none> 443/TCP 123d nginx NodePort 10.100.233.149 <none> 80:30457/TCP 31s redis ClusterIP 10.106.55.214 <none> 6379/TCP 31s 复制代码`

注意nginx服务部分的PORTS为 **80:30457/TCP** ,这表示将容器的80端口暴露到本机网络的30457端口，和我们之前的docker启动时候的 ` -p 80:80` 参数类似。

服务是由Pod提供的，继续检查一下pods的状况 ` kubectl get pods` :

` NAME READY STATUS RESTARTS AGE flaskapp-6c4fccdf99-v6w2v 1/1 Running 0 2m nginx-85fb469b96-lr982 1/1 Running 0 2m redis-5b44bb8d97-wwmll 1/1 Running 0 2m 复制代码`

当然，也可以直接查看docker的容器 ` docker ps` :

` ➜ docker2istio docker ps CONTAINER ID IMAGE COMMAND CREATED STATUS PORTS NAMES ad7377ae7196 ae70b17240ec "docker-entrypoint.s…" About an hour ago Up About an hour k8s_redis_redis-5b44bb8d97-wwmll_default_2907f4a3-6639-11e9-b8cb-025000000001_0 c01108b49076 1a61773c4c07 "python flaskapp.py" About an hour ago Up About an hour k8s_flaskapp_flaskapp-6c4fccdf99-xcmwb_default_28fbe1b1-6639-11e9-b8cb-025000000001_0 11d1fa3f182b 315798907716 "nginx -g 'daemon of…" About an hour ago Up About an hour k8s_nginx_nginx-85fb469b96-lr982_default_28fbdeee-6639-11e9-b8cb-025000000001_0 c28032a4b068 k8s.gcr.io/pause-amd64:3.1 "/pause" About an hour ago Up About an hour k8s_POD_redis-5b44bb8d97-wwmll_default_2907f4a3-6639-11e9-b8cb-025000000001_0 7091657acfbc k8s.gcr.io/pause-amd64:3.1 "/pause" About an hour ago Up About an hour k8s_POD_flaskapp-6c4fccdf99-xcmwb_default_28fbe1b1-6639-11e9-b8cb-025000000001_0 97007670c247 k8s.gcr.io/pause-amd64:3.1 "/pause" About an hour ago Up About an hour k8s_POD_nginx-85fb469b96-lr982_default_28fbdeee-6639-11e9-b8cb-025000000001_0 ... 复制代码`
> 
> 
> 
> 
> **!!!注意:** pod并不等同于docker的容器，Pod才是k8s操作的最小单元。简单的说，一个Pod可能包含多个容器，从yaml文件中 **containers:
> **这个关键字可以看出。仔细观察 ` docker ps` 的输出，可以发现每个pod除了用户自定义的容器外，还有镜像为**
> k8s.gcr.io/pause-amd64:3.1** 的系统容器。
> 
> 

最后使用 ` curl http://127.0.0.1:30457` 访问服务

` Hello World by 10.1.0.21 from 192.168.65.3 ! 该页面已被访问 1 次。 复制代码`

### 扩容 ###

k8s集群下，扩容非常简单

` ➜ docker2istio kubectl edit deployment/flaskapp deployment.extensions "flaskapp" edited 复制代码`

修改其中的** replicas: 3 **。

> 
> 
> 
> 也可以修改 ` k8s\flaskapp.yaml` 中的值，然后 ` kubectl apply -f k8s\flaskapp.yaml`
> 
> 

> 
> 
> 
> 另外，如果镜像有更新，也是采用修改flaskapp.yaml文件然后apply的方式。
> 
> 

` kubectl get pods -o wide` 检查扩容结果, 这里使用了 ` -o wide` ,可以显示更多信息

` NAME READY STATUS RESTARTS AGE IP NODE flaskapp-6c4fccdf99-9xsjl 1/1 Running 0 3m 10.1.0.23 docker-for-desktop flaskapp-6c4fccdf99-xcmwb 1/1 Running 0 1h 10.1.0.21 docker-for-desktop flaskapp-6c4fccdf99-zp8mk 1/1 Running 0 3m 10.1.0.24 docker-for-desktop nginx-85fb469b96-lr982 1/1 Running 0 1h 10.1.0.19 docker-for-desktop redis-5b44bb8d97-wwmll 1/1 Running 0 1h 10.1.0.22 docker-for-desktop 复制代码`

多次访问服务:

` ➜ docker2istio curl http://127.0.0.1:30457 Hello World by 10.1.0.21 from 192.168.65.3 ! 该页面已被访问 2 次。 ➜ docker2istio curl http://127.0.0.1:30457 Hello World by 10.1.0.23 from 192.168.65.3 ! 该页面已被访问 3 次。 ➜ docker2istio curl http://127.0.0.1:30457 Hello World by 10.1.0.24 from 192.168.65.3 ! 该页面已被访问 4 次。 ➜ docker2istio curl http://127.0.0.1:30457 复制代码`

结合前面看到的flaskapp的IP，可以比较清晰的看到请求会自动负载到不同的Pod。

### 清理 ###

k8s下的容器清理也非常简单, 使用 ` kubectl delete -f k8s` :

` service "flaskapp" deleted deployment.extensions "flaskapp" deleted configmap "nginx-config" deleted service "nginx" deleted deployment.extensions "nginx" deleted service "redis" deleted deployment.extensions "redis" deleted 复制代码`

## 容器编排 ##

实际上，k8s集群在多集群情况下，会自动将Pod调度到合适的节点，这就是容器编排的概念。这种能力，主要有2个方式。

### 节点标签 ###

我们的k8s演示集群节点情况如下:

` [tyhall51@192-168-10-21 k8s]$ kubectl get nodes NAME STATUS ROLES AGE VERSION 192-168-10-14 Ready <none> 13d v1.14.0 192-168-10-18 Ready <none> 130d v1.14.0 192-168-10-21 Ready master 131d v1.14.0 复制代码`

部署示例应用到k8s演示集群:

` [tyhall51@192-168-10-21 docker2istio]$ kubectl apply -f k8s -n docker2istio service/flaskapp created deployment.extensions/flaskapp created configmap/nginx-config created service/nginx created deployment.extensions/nginx created service/redis created deployment.extensions/redis created 复制代码`
> 
> 
> 
> 
> **!!!注意** 为了不和别的服务发生名称冲突，这里部署时候使用了 ` -n docker2istio` 参数，创建了一个独立的名称空间。名称空间可以使用
> ` kubectl create namespace docker2istio` 命令创建。
> 
> 

查看名称空间下的服务:

` [tyhall51@192-168-10-21 docker2istio]$ kubectl get service -n docker2istio NAME TYPE CLUSTER-IP EXTERNAL-IP PORT(S) AGE flaskapp ClusterIP 10.101.127.107 <none> 5000/TCP 47s nginx NodePort 10.103.147.187 <none> 80:30387/TCP 46s redis ClusterIP 10.106.162.13 <none> 6379/TCP 46s 复制代码`

查看名称空间下的pod:

` [tyhall51@192-168-10-21 docker2istio]$ kubectl get pods -o wide -n docker2istio NAME READY STATUS RESTARTS AGE IP NODE NOMINATED NODE READINESS GATES flaskapp-589c4cdf86-sftr9 1/1 Running 0 81s 10.244.2.30 192-168-10-14 <none> <none> nginx-55b87f44ff-b4x88 1/1 Running 0 81s 10.244.2.31 192-168-10-14 <none> <none> redis-7 fc 7 fc 64fb-2nzjq 1/1 Running 0 81s 10.244.1.195 192-168-10-18 <none> <none> 复制代码`

参考前文，修改副本数量参数 **replicas** ，对flaskapp进行扩容:

` [tyhall51@192-168-10-21 docker2istio]$ kubectl get pods -o wide -n docker2istio NAME READY STATUS RESTARTS AGE IP NODE NOMINATED NODE READINESS GATES flaskapp-589c4cdf86-8jzwx 1/1 Running 0 4s 10.244.1.197 192-168-10-18 <none> <none> flaskapp-589c4cdf86-sftr9 1/1 Running 0 3m10s 10.244.2.30 192-168-10-14 <none> <none> flaskapp-589c4cdf86-tz98x 1/1 Running 0 4s 10.244.1.196 192-168-10-18 <none> <none> nginx-55b87f44ff-b4x88 1/1 Running 0 3m10s 10.244.2.31 192-168-10-14 <none> <none> redis-7 fc 7 fc 64fb-2nzjq 1/1 Running 0 3m10s 10.244.1.195 192-168-10-18 <none> <none> 复制代码`

这里就可以看到，扩容完成后，flaskapp的3个pod会自动调度到 **192-168-10-18** 和 **192-168-10-18** 2个业务节点。

192-168-10-14节点的磁盘使用的是高速ssd，io性能会更好一些，我们希望redis能够调度到该节点。

首先，给192-168-10-14节点打上 ` storage=ssd` 的标签:

` [tyhall51@192-168-10-21 docker2istio]$ kubectl label nodes 192-168-10-14 storage=ssd node/192-168-10-14 labeled 复制代码`

检查标签是否正常标记:

` [tyhall51@192-168-10-21 docker2istio]$ kubectl get nodes --show-labels | grep ssd 192-168-10-14 Ready <none> 13d v1.14.0 beta.kubernetes.io/arch=amd64,beta.kubernetes.io/os=linux,kubernetes.io/arch=amd64,kubernetes.io/hostname=192-168-10-14,kubernetes.io/os=linux,storage=ssd 复制代码`

然后修改 ` k8s/redis.yaml` ,增加 ` nodeSelector` 数值，其值为 ` storage: ssd` , 修改完成的deployment如下:

` apiVersion: extensions/v1beta1 kind: Deployment metadata: name: redis spec: replicas: 1 template: metadata: labels: name: redis spec: containers: - image: redis:4-alpine3.8 name: redis ports: - containerPort: 6379 nodeSelector: storage: ssd 复制代码`

使用 ` kubectl apply -f k8s/redis.yaml -n docker2istio` 应用修改。查看docker2istio的pod分布情况:

` [tyhall51@192-168-10-21 docker2istio]$ kubectl get pods -o wide -n docker2istio NAME READY STATUS RESTARTS AGE IP NODE NOMINATED NODE READINESS GATES flaskapp-589c4cdf86-8jzwx 1/1 Running 0 11m 10.244.1.197 192-168-10-18 <none> <none> flaskapp-589c4cdf86-sftr9 1/1 Running 0 14m 10.244.2.30 192-168-10-14 <none> <none> flaskapp-589c4cdf86-tz98x 1/1 Running 0 11m 10.244.1.196 192-168-10-18 <none> <none> nginx-55b87f44ff-b4x88 1/1 Running 0 14m 10.244.2.31 192-168-10-14 <none> <none> redis-66f66896b6-7666t 1/1 Running 0 4s 10.244.2.35 192-168-10-14 <none> <none> 复制代码`

可见redis节点重新被调度到192-168-10-14节点，表现出了节点标签的亲和力。

### 节点污点 ###

在k8s演示集群中192-168-10-21是master节点，默认不会调度业务pod，这种能力是采用节点污点实现的。 取消192-168-10-21调度污点:

` kubectl taint node 192-168-10-21 node-role.kubernetes.io/master:NoSchedule- 复制代码`

然后扩容flaskapp的副本数到6个，观察pod分布情况:

` [tyhall51@192-168-10-21 docker2istio]$ kubectl get pods -o wide -n docker2istio NAME READY STATUS RESTARTS AGE IP NODE NOMINATED NODE READINESS GATES flaskapp-589c4cdf86-8jzwx 1/1 Running 0 20m 10.244.1.197 192-168-10-18 <none> <none> flaskapp-589c4cdf86-92rm5 1/1 Running 0 5s 10.244.2.36 192-168-10-14 <none> <none> flaskapp-589c4cdf86-bfhs8 1/1 Running 0 5s 10.244.0.26 192-168-10-21 <none> <none> flaskapp-589c4cdf86-sftr9 1/1 Running 0 23m 10.244.2.30 192-168-10-14 <none> <none> flaskapp-589c4cdf86-srv25 1/1 Running 0 5s 10.244.0.25 192-168-10-21 <none> <none> flaskapp-589c4cdf86-tz98x 1/1 Running 0 20m 10.244.1.196 192-168-10-18 <none> <none> nginx-55b87f44ff-b4x88 1/1 Running 0 23m 10.244.2.31 192-168-10-14 <none> <none> redis-66f66896b6-7666t 1/1 Running 0 9m30s 10.244.2.35 192-168-10-14 <none> <none> 复制代码`

这里可以看到，有2个pod被调到到192-168-10-21节点了。

重新设置污点:

` [tyhall51@192-168-10-21 docker2istio]$ kubectl taint node 192-168-10-21 node-role.kubernetes.io/master=:NoSchedule node/192-168-10-21 tainted 复制代码`

删除在192-168-10-21上的2个pod:

` kubectl delete pod/flaskapp-589c4cdf86-bfhs8 -n docker2istio kubectl delete pod/flaskapp-589c4cdf86-srv25 -n docker2istio 复制代码`

观察pod分布情况:

` [tyhall51@192-168-10-21 docker2istio]$ kubectl get pods -o wide -n docker2istio NAME READY STATUS RESTARTS AGE IP NODE NOMINATED NODE READINESS GATES flaskapp-589c4cdf86-8jzwx 1/1 Running 0 25m 10.244.1.197 192-168-10-18 <none> <none> flaskapp-589c4cdf86-92rm5 1/1 Running 0 4m40s 10.244.2.36 192-168-10-14 <none> <none> flaskapp-589c4cdf86-fp5w4 1/1 Running 0 73s 10.244.2.37 192-168-10-14 <none> <none> flaskapp-589c4cdf86-lv2ch 1/1 Running 0 73s 10.244.1.199 192-168-10-18 <none> <none> flaskapp-589c4cdf86-p9kb6 1/1 Running 0 7s 10.244.2.38 192-168-10-14 <none> <none> flaskapp-589c4cdf86-sftr9 1/1 Running 0 28m 10.244.2.30 192-168-10-14 <none> <none> nginx-55b87f44ff-b4x88 1/1 Running 0 28m 10.244.2.31 192-168-10-14 <none> <none> redis-66f66896b6-7666t 1/1 Running 0 14m 10.244.2.35 192-168-10-14 <none> <none> 复制代码`

可以看到删除后的pod，在192-168-10-18和192-168-10-14这2个业务节点上重建了。

## 总结 ##

k8s相对于compose：

* 管理规模扩大，由单机到集群。
* 扩容更方便了，可以无缝扩容。
* 部署策略更完善，可以对容器进行 **编排** 。

## 相关组件 ##

* Etcd

> 
> 
> 
> etcd 是一个分布式键值对存储，设计用来可靠而快速的保存关键数据并提供访问。通过分布式锁，leader选举和写屏障(write
> barriers)来实现可靠的分布式协作。etcd集群是为高可用，持久性数据存储和检索而准备。k8s中使用etcd作为集群信息存储。
> 
> 

* Efk

> 
> 
> 
> EFK (Elasticsearch + Fluentd + Kibana) 是kubernetes官方推荐的日志收集方案
> 
> 

* helm

> 
> 
> 
> Helm helps you manage Kubernetes applications — Helm Charts help you
> define, install, and upgrade even the most complex Kubernetes application.
> 
> 
> 

* Rock

> 
> 
> 
> File, Block, and Object Storage Services for your Cloud-Native
> Environments
> 
> 

...

...