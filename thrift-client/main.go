package main

import (
    "fmt"
    "net"
    "os"

    "time"
    "git.apache.org/thrift.git/lib/go/thrift"
    "example/TestThrift/gen-go/mythrift/demo"
    "strings"
)

const (
    HOST = "127.0.0.1"
    PORT = "9090"
)

func main() {
   // startTime := currentTimeMillis()

    transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
    protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

    transport, err := thrift.NewTSocket(net.JoinHostPort(HOST, PORT))
    if err != nil {
        fmt.Fprintln(os.Stderr, "error resolving address:", err)
        os.Exit(1)
    }

    useTransport := transportFactory.GetTransport(transport)
    client := demo.NewMyThriftClientFactory(useTransport, protocolFactory)
    if err := transport.Open(); err != nil {
        fmt.Fprintln(os.Stderr, "Error opening socket to "+HOST+":"+PORT, " ", err)
        os.Exit(1)
    }
    defer transport.Close()

    //for i := 0; i < 10; i++ {
    //    paramMap := make(map[string]string)
    //    paramMap["a"] = "mythrift.demo"
    //    paramMap["b"] = "test" + strconv.Itoa(i+1)
	//
    //    fmt.Println("aaa")
    //    r1, _ := client.CallBack(time.Now().Unix(), "Go client", paramMap)
	//
    //    fmt.Println("Go client call->", r1)
	//
    //}
for {
    sendmess, _ := client.SendSMS()

    s := strings.Split(string(sendmess), " ")

    fmt.Println("sdt: ", s[0])
    fmt.Println("content: ", s[1])
}
    //model := demo.Article{1, "Send from Golang Thrift Client", "Hello World!", "Golang"}
    //client.Put(&model)
    //endTime := currentTimeMillis()
    //fmt.Printf("The call took:%d-%d=%d Millis \n", endTime, startTime, (endTime - startTime))
}

func currentTimeMillis() int64 {
    return time.Now().UnixNano() / 1000000
}
