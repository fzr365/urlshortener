## 1.技术架构：GO+MySQL+Redis+Docker
## 2.项目简介：短链接生成器是将一个长链接转为短链接
## 项目难点
## 这是一个读多写少的项目。
## 难点1： GET请求需要服务时延低，响应速度快，使重定向的用户没有痛感。如果请求都要访问数据库，涉及到磁盘IO会增加响应时间，同时大量的请求会给数据库很多压力。所以需要使用redis进行缓存。
## 难点2： 短URL的id如何生成，短id是可能重复的，需要使用重试机制提高成功率。
## 持续更新中，欢迎star