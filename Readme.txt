�뽫���ļ��зŵ���Ҫ�ϴ���json�ļ�·���£����߻��Զ����ϲ�Ŀ¼ͳ��json�ļ���

conf.txt�ļ����ͣ� ������������ϸ��մ˸�ʽ�����ã�����������ġ�
sendpath1=/data/game/basketball_0/newserver/data/  #��Ϊ�ϴ�����������Ŀ¼�����ϸ��մ˸�ʽ�������ļ�Ŀ¼  ���Բ��������д�ո�֮���
sendpath2=/data/game/basketball_1/newserver/data/
sendpath3=/data/game/basketball_2/newserver/data/
ip=192.168.1.219                                   #��Ϊ������ip
f_port=21                                          #��Ϊ������ftp�˿� ���ڷ����ļ�
s_port=22                                          #��Ϊ������ssh�˿� ���ڷ�������
user=root                                          #��Ϊ�������û��� ���������� ����ֱ��ʹ��root�û�
password=123456                                    #��Ϊ����������
cmd1=                                              #��һ��ִ��Զ�̷��������� ���Բ���
cmd2=systemctl restart supervisord                 #�ڶ���ִ��Զ�̷�����������  ������Ϊ������Ϸ����






##########################################################
�vsftpd:
1:yum install vsftpd -y
2:vim /etc/vsftpd/ftpusers
  ��root  ע�͵� ��Ϊ #root
3:vim /etc/vsftpd/user_list
   ��root  ע�͵� ��Ϊ #root
   
4:vim /etc/vsftpd/vsftpd.conf
   ��anonymous_enable=YES ��Ϊanonymous_enable=NO
   
5:systemctl start vsftpd #����ftp����