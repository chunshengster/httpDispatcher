- Why Httpdns
- 需求
    - dns外挂
        - 兼容任意支持 edns client subnet 扩展的 dos server（dnspod、baidu、ali）
    - cache
        - 自动更新 (未完整支持)
            - soa
            - domainnode
            - rr for region
    - region支持自定义，未必采用dns server在edns subnet中返回的 region
    - 支持数据库数据 （未完整支持）
- 设计
    - dns request
        - soa
        - ns
        - a
        - cname
    - cache
        - rbtree
        - radix32
- 性能
mac air 2013 乞丐版
![](https://github.com/chunshengster/httpDispatcher/blob/develop/images/bench.png)