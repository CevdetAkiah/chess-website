import path from 'path'
import * as grpc from '@grpc/grpc-js'
import * as protoLoader from '@grpc/proto-loader'
import {ProtoGrpcType} from './proto/random'
import { ChatServiceHandlers } from './proto/randomPackage/ChatService'


const PORT = 8082
const PROTO_FILE = './proto/random.proto'

const packageDef = protoLoader.loadSync(path.resolve(__dirname, PROTO_FILE))
const grpcObj = (grpc.loadPackageDefinition(packageDef) as unknown) as ProtoGrpcType
const randomPackage = grpcObj.randomPackage

function main() {
  const server = getServer()

  server.bindAsync(`0.0.0.0:${PORT}`, grpc.ServerCredentials.createInsecure(),
  (err, port) => {
    if (err) {
      console.error(err)
      return
    }
    console.log(`Your server as started on port ${port}`)
    server.start()
  })
}

// const callObjByUsername = new Map<string, grpc.ServerDuplexStream<ChatRequest, ChatResponse>>()

function getServer() {
  const server = new grpc.Server()
  server.addService(randomPackage.ChatService.service, {
    // call = InitiateRequest, callback = InitiateResponse (see random.proto)
    ChatInitiate: (call, callback) =>{
      // simple validation
      const sessionName = call.request.name || '' // it's possible to not send a name so check for empty string
      const avatar = call.request.avatarUrl || ''
      if (!sessionName || !avatar) callback(new Error("Name and avatar required")) // send an error back to the client
      console.log("hey")
      callback(null, {id: Math.floor(Math.random() * 10000)}) // send back InitiateResponse with the promised id
      
    }
  } as ChatServiceHandlers)

  return server
}

main()