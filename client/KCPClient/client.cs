﻿using System;
using System.Net.Sockets;
using System.Net.Sockets.Kcp;
using System.Net;
using System.Threading;
using System.Text;
using KCPNET;
using System.Threading.Tasks;
using Pb;





namespace KCPClient
{
    class Client
    {
        static KCPNet<ClientSession> clientSession;
        static Task<bool> checkTask;
        
        static void Main(string[] args)
        {
            string ip = "127.0.0.1";
            
            clientSession = new KCPNet<ClientSession>();
            clientSession.StartClient(ip,7777);
            checkTask =  clientSession.ConnectServer(200,5000);
            Task.Run(ConnectServer);
        
        
        Console.ReadKey();
        
        }
        private static int counter = 0;
        static async void ConnectServer() {
            while(true){
                await Task.Delay(3000);

                if(checkTask != null &&checkTask.IsCompleted){
                    if(checkTask.Result){
                        checkTask = null;
                        await Task.Run(SendPingMessage);
                    }
                    else{
                         ++counter;
                        if(counter > 4) {
                            Console.WriteLine(string.Format("Connect Failed {0} Times,Check Your Network Connection.", counter));
                            checkTask = null;
                            break;
                        }
                        else {
                            Console.WriteLine(string.Format("Connect Faild {0} Times.Retry...", counter));
                            checkTask = clientSession.ConnectServer(200, 5000);
                        }
                        
                    }

                }
            }
        }


        static async void SendPingMessage(){
            while(true){
                await Task.Delay(5000);
                if(clientSession != null && clientSession.session != null){
                    var mes = new PbMessage{
                        Name = "ping_test",     
                        Cmd = PbMessage.Types.CMD.Login                   
                    };
                    clientSession.session.SendMessage(mes);
                    Console.WriteLine("发送的信息： ");
                    Console.WriteLine(mes.Cmd);
                }
                else
                    break;

            }
        }

    }
                   
}
