import os,sys, subprocess
from tinyrpc.transports import ServerTransport
from tinyrpc.protocols.jsonrpc import JSONRPCProtocol
from tinyrpc.dispatch import public,RPCDispatcher
from tinyrpc.server import RPCServer


try:
    import urllib.parse as urlparse
except ImportError:
    import urllib as urlparse

class StdIOTransport(ServerTransport):

    def receive_message(self):
        return None, urlparse.unquote(sys.stdin.readline())

    def send_reply(self, context, reply):
        print(reply)

class PipeTransport(ServerTransport):


    def __init__(self,input, output):
        self.input = input
        self.output = output

    def receive_message(self):
        data = self.input.readline()
        print(">> {}".format( data))
        return None, urlparse.unquote(data)

    def send_reply(self, context, reply):
        print("<< {}".format( reply))
        self.output.write(reply)
        self.output.write("\n")

class StdIOHandler():
    def __init__(self):
        pass

    @public
    def ApproveTx(self,req):

        transaction = req.get('transaction')
        _from       = req.get('from')
        call_info   = req.get('call_info')
        meta        = req.get('meta')

        return {
            "approved" : False,

        }

    @public
    def ApproveSignData(self, req):

        return {"approved": False, "password" : None}

    @public
    def ApproveExport(self, req):

        return {"approved" : False}

    @public
    def ApproveImport(self, req):

        return { "approved" : False, "old_password": "", "new_password": ""}

    @public
    def ApproveListing(self, req):
 
        return {'accounts': []}

    @public
    def ApproveNewAccount(self, req):

        return {"approved": False,
                #"password": ""
                }

    @public
    def ShowError(self,message = {}):

        if 'text' in message.keys():
            sys.stderr.write("Error: {}\n".format( message['text']))
        return

    @public
    def ShowInfo(self,message = {}):


        if 'text' in message.keys():
            sys.stdout.write("Error: {}\n".format( message['text']))
        return

def main(args):
    cmd = ["clef", "--stdio-ui"]
    if len(args) > 0 and args[0] == "test":
        cmd.extend(["--stdio-ui-test"])
    print("cmd: {}".format(" ".join(cmd)))
    dispatcher = RPCDispatcher()
    dispatcher.register_instance(StdIOHandler(), '')
    p = subprocess.Popen(cmd, bufsize=1, universal_newlines=True, stdin=subprocess.PIPE, stdout=subprocess.PIPE)

    rpc_server = RPCServer(
        PipeTransport(p.stdout, p.stdin),
        JSONRPCProtocol(),
        dispatcher
    )
    rpc_server.serve_forever()

if __name__ == '__main__':
    main(sys.argv[1:])
