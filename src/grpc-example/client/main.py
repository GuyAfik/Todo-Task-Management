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
                test=[users_pb2.Test(score=90, name="Mathematics"), users_pb2.Test(score=70, name="English")], 
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
