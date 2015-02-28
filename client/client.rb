require_relative "proto/hellowold_services"

def main
  stub = Helloworld::Greeter::Stub.new('localhost:11111')
  user = ARGV.size > 0 ? ARGV[0] : 'world'
  message = stub.say_hello(Helloworld::HelloRequest.new(name: user)).message
  puts "Greeting: #{message}"
end

main
