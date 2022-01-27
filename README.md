### [VPS信息监控](https://github.com/cloudapp3/host)
1. 特性
    - 定时获取Linux-VPS信息
    - 支持通过TelegramBot查询[@loveforpeace_bot](https://t.me/loveforpeace_bot)
2. Linux宿主信息
    - CPU使用率
    - 硬盘空间
    - 内存可用
    - 流量
    - 在线时间

### 使用流程
1. telegram 打开[@loveforpeace_bot](https://t.me/loveforpeace_bot)
2. start开启bot  
3. 输入/token 获取token 记录下token
4. 64位linux请登录VPS执行    
 ~~~
 wget -O - https://github.com/cloudapp3/host/blob/master/run_linux_x64.sh | bash
 ~~~~
5. 其他系统执行文件请自行到release页面下载 
### todo
1. web dashboard
2. 指标监控
3. 邮件 && Telegram通知

### 技术栈
1. Go + rpcx

### 产品使用反馈
- [Telegram交流群](https://t.me/VPSTG)