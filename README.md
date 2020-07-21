# crawler
crawler. Search pages and Downloads files

## 功能

仅用于周期性获取BGP数据 (食之无味，弃之可惜)

## 使用方法
```
# 简易版
./main http://archive.routeviews.org bgpdata 2003 03 UPDATES

# 集成版
python3 set.py
sh filelist.sh
```

数据存储格式为files/yearmonth 目录下存放,可以适应ETL流程

## 个人自评

Go写的小爬虫,会解析访问网址中所有a标签.通过后缀匹配下载数据到files文件夹中,日志存储到log
有点鸡肋...一时兴起的产物

## 推荐的方法

wget curl 或 xpath 都是更好的选择,不推荐引用这个项目hhh

## 报错解析
panic: runtime error: invalid memory address or nil pointer dereference [signal SIGSEGV: segmentatio

诊断原因触发了反爬机制，爬取速度过快，加入减速机制
