# 介绍
这是一个go语言写的 beanstalkd 工作者程序，使用CGo连接虹软人脸识别 linux C++ SDK。

# 待完成部分
1. 现在做成docker镜像时有1.3G大小，大部分是opencv以及各种依赖环境。除非改为静态链接到opencv，才能把镜像缩小。