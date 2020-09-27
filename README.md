# iptvProxy
代理iptv的源

## 在普通linux Server安装启动
### arm环境 生成docker镜像
docker build -t iptvproxy .  -f Arm64_Dockerfile
### 普通x86环境 生成docker镜像
docker build -t iptvproxy .  -f Dockerfile
### 部署docker镜像 暴露访问端口为19000
docker run -d --name iptvproxy --privileged --restart always -p 0.0.0.0:19000:19000 iptvproxy



## Install kintohub
1、注册kintohub账号

2、创建免费的service,选择web app

3、import URL，输入本项目地址，connect

4、Port端口改成443

5、Dockerfile Name 改为 hintohub_Dockerfile

6、Deploy,等待结果，成功后，会在log窗口输出访问的路径
