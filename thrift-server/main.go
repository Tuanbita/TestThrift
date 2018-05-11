package main

import (
    "fmt"
    "os"
    "time"

    "git.apache.org/thrift.git/lib/go/thrift"
    "example/TestThrift/gen-go/mythrift/demo"
    "bufio"
    "strings"
)

const (
    NetworkAddr = "127.0.0.1:9090"
)

type mythriftThrift struct{}

func (this *mythriftThrift) CallBack(callTime int64, name string, paramMap map[string]string) (r []string, err error) {

    var mess string
    fmt.Print("Nhap content: ")
    user := bufio.NewReader(os.Stdin)
    mess,_ = user.ReadString('\n')
    mess = strings.TrimSpace(mess)

    fmt.Println("-->from client Call:", time.Unix(callTime, 0).Format("2006-01-02 15:04:05"), name, paramMap)
    r = append(r, ""+mess+"  " +paramMap["a"]+"    value:"+paramMap["b"])

    return
}

func (this *mythriftThrift) Put(s *demo.Article) (err error) {
    fmt.Printf("Article--->id: %d\tTitle:%s\tContent:%s\tAuthor:%s\n", s.ID, s.Title, s.Content, s.Author)
    return nil
}
func (this *mythriftThrift) SendSMS()(r string, err error)  {

    //var sdt string
    //fmt.Print("Nhap sdt: ")
    //user := bufio.NewReader(os.Stdin)
    //sdt,_ = user.ReadString('\n')
    //sdt = strings.TrimSpace(sdt)
	//
    //var mess string
    //fmt.Print("Nhap content: ")
    //user = bufio.NewReader(os.Stdin)
    //mess,_ = user.ReadString('\n')
    //mess = strings.TrimSpace(mess)

    r = "1" + " " +"2"
    time.Sleep(time.Second*2)
    return
}

func main() {

    transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
    protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
    serverTransport, err := thrift.NewTServerSocket(NetworkAddr)
    if err != nil {
        fmt.Println("Error!", err)
        os.Exit(1)
    }

    handler := &mythriftThrift{}
    processor := demo.NewMyThriftProcessor(handler)

    server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
    fmt.Println("thrift server in", NetworkAddr)
    server.Serve()
}
