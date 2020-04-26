# crawler
crawler. Search pages and Downloads files

用于周期性获取BGP数据

使用方法
```
./main http://archive.routeviews.org bgpdata 2003 03 UPDATES
```

数据存储格式为files/yearmonth 目录下存放,可以适应ETL流程

报错:
panic: runtime error: invalid memory address or nil pointer dereference [signal SIGSEGV: segmentatio

诊断原因触发了反爬机制，爬取速度过快，加入减速机制