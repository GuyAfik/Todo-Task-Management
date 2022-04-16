# from client.users_pb2 import EyeColour
from unicodedata import name
import grpc
import users_pb2
import users_pb2_grpc

def main():
    try:
        channel = grpc.insecure_channel('localhost:3001')
        stub = users_pb2_grpc.UserServiceStub(channel)
        request = users_pb2.UserRequest(
                name="Guy", 
                test=[users_pb2.Test(score=90, name="Mathematics"), 
                users_pb2.Test(score=70, name="English")], 
                eyeColour=users_pb2.GREEN, 
                hasFacebook=True,
                height=1.75
            )
        response = stub.scoreAvg(request)
        print(response)
    except Exception as e:
        print(str(e))

if __name__ == "__main__":
    main()

# Name        string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
# 	Height      float32   `protobuf:"fixed32,2,opt,name=height,proto3" json:"height,omitempty"`
# 	HasFacebook bool      `protobuf:"varint,3,opt,name=hasFacebook,proto3" json:"hasFacebook,omitempty"`
# 	EyeColour   EyeColour `protobuf:"varint,4,opt,name=eyeColour,proto3,enum=example.users.EyeColour" json:"eyeColour,omitempty"`
# 	Test        []*Test   `protobuf:"bytes,5,rep,name=test,proto3" json:"test,omitempty"`