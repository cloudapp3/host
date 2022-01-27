echo "欢迎使用host-telegram探针"
echo "从github下载host"
wget https://github.com/cloudapp3/host/releases/download/v1.0.0.0/host_linux_amd64
chmod 777 host_linux_amd64
echo "请输入token 获取 token方法请参考 https://github.com/cloudapp3/host#readme"
read token
nohup ./host_linux_amd64 -t $token &
echo "运行中 如有任何问题 请到telegram群组 https://t.me/VPSTG 反馈交流"