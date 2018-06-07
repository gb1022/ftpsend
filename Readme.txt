请将此文件夹放到需要上传的json文件路径下，工具会自动向上层目录统计json文件。

conf.txt文件解释： 各个配置项都请严格按照此格式来配置，不能随意更改。
sendpath1=/data/game/basketball_0/newserver/data/  #此为上传到服务器的目录，请严格按照此格式来配置文件目录  可以不填，但不能写空格之类的
sendpath2=/data/game/basketball_1/newserver/data/
sendpath3=/data/game/basketball_2/newserver/data/
ip=192.168.1.219                                   #此为服务器ip
f_port=21                                          #此为服务器ftp端口 用于发送文件
s_port=22                                          #此为服务器ssh端口 用于发送命令
user=root                                          #此为服务器用户名 由于是内网 所以直接使用root用户
password=123456                                    #此为服务器密码
cmd1=                                              #第一个执行远程服务器命令 可以不填
cmd2=systemctl restart supervisord                 #第二个执行远程服务器的命令  此命令为重启游戏服务






##########################################################
搭建vsftpd:
1:yum install vsftpd -y
2:vim /etc/vsftpd/ftpusers
  将root  注释掉 改为 #root
3:vim /etc/vsftpd/user_list
   将root  注释掉 改为 #root
   
4:vim /etc/vsftpd/vsftpd.conf
   将anonymous_enable=YES 改为anonymous_enable=NO
   
5:systemctl start vsftpd #启动ftp服务