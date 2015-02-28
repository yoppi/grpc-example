require_relative "../proto/hellowold_services"

class GreeterServer < Helloworld::Greeter::Service
  def say_hello(hello_req, _unused_call)
    Helloworld::HelloReply.new(message: "Hello, #{hello_req.name}")
  end
end

def main
  server = GRPC::RpcServer.new
  server.add_http2_port("0.0.0.0:11111")
  server.handle(GreeterServer)
  server.run
end

main
