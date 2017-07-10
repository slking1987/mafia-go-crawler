# MAFIA-Go-Crawler
基于golang编写的爬虫程序   
v0.1 Present by slking1987
### 0.架构
input chan->module/crawler 使用http get获取页面数据->process chan   
process chan->module/processor 处理数据，提取url->input chan，提取目标数据->output chan   
output chan->module/output 输出到console/file/mysql/es等
### 1.代码
####核心模块
module/crawler 数据收集   
module/processor 数据提取   
module/output 数据输出   
####辅助模块
common 通用   
resource 资源   
module/monitor 监控   
module/settings 配置  
### 2.TODO
sub urls 结束条件, 可以考虑爬2层即可   
爬取内容暂时还不确定从何入手