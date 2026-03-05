package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

func main1() {
	txtResult := make(chan string, 5)
	go func() { txtResult <- getTxt("res1.flysnow.org") }()
	go func() { txtResult <- getTxt("res2.flysnow.org") }()
	go func() { txtResult <- getTxt("res3.flysnow.org") }()
	go func() { txtResult <- getTxt("res4.flysnow.org") }()
	go func() { txtResult <- getTxt("res5.flysnow.org") }()
	// 5个协程并发向channel写入字符串，然后输出第一个
	println(<-txtResult)
}

func getTxt(host string) string {
	//省略网络访问逻辑，直接返回模拟结果
	//http.Get(host+"/1.txt")
	return host + "：模拟结果"
}

// 生产者：生成factor整数倍的序列
func Producer(factor int, out chan<- int) {
	for i := 0; ; i++ {
		out <- i * factor
	}
}

// 消费者
func Consumer(in <-chan int) {
	fmt.Println(cap(in))

	// range可以用来遍历消费的队列，这个队列可以是chan，就是管道，也就是说可能是无限循环的，只要管道还有数据
	for v := range in {
		fmt.Println(v)
	}
}
func main2() {
	ch := make(chan int, 64) // 成果队列
	go Producer(3, ch)       // 生成3的倍数的序列
	go Producer(5, ch)       // 生成5的倍数的序列
	go Consumer(ch)          // 消费生成的队列
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", <-sig)
}

func main3() {
	ch := make(chan int)
	go func() {
		for {
			select {
			case ch <- 0:
			case ch <- 1:
			}
		}
	}()
	for v := range ch {
		fmt.Println(v)
	}
}

// 返回生成自然数序列的通道: 2, 3, 4, ...
func GenerateNatural(ctx context.Context) chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			select {
			case <-ctx.Done():
				return
			case ch <- i:
				fmt.Println("GenerateNatural ch写入", i)
			}
		}
	}()
	return ch
}

// 通道过滤器：删除能被素数整除的数
func PrimeFilter(ctx context.Context, in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				select {
				case <-ctx.Done():
					return
				case out <- i:
					fmt.Println("PrimeFilter 写入", i)
				}
			}
		}
	}()
	return out
}

func main4() {
	// 通过Context控制后台Goroutine状态
	ctx, cancel := context.WithCancel(context.Background())
	ch := GenerateNatural(ctx) // 自然数序列：2, 3, 4, ...
	for i := 0; i < 100; i++ {
		prime := <-ch // 新出现的素数
		//fmt.Printf("%v: %v\n", i+1, prime)
		ch = PrimeFilter(ctx, ch, prime) // 基于新素数构造的过滤器
	}
	cancel()
}

func main5() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("ending")
		}
	}()

	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("finishing")
		}
	}()
	//取消的时候两个协程都能接收到信号然后退出
	time.Sleep(time.Second)
	cancel()
	time.Sleep(time.Second)

}

func main6() {
	cn := make(chan int, 2)
	//这个只支持写入的chan
	strChan := make(chan<- string)

	//只支持读的chan,现在是没有长度的，无法读出内容
	//stringsChan := make(<-chan string)
	//读不出来，会报错的，因为是无缓冲区无内容的
	//_ = <-stringsChan

	//var str ="aaa"
	strChan <- "aaa"

	go func() {
		time.Sleep(time.Second)
		cn <- 20
		cn <- 22
		cn <- 24
		//只有两个缓冲区，所以等读出去一个之后才会执行这块代码，所以是最后的输出语句
		fmt.Println("cn <- 20")
	}()
	time.Sleep(time.Second * 3)
	fmt.Println("before <-cn")
	i := <-cn
	fmt.Println("after <-cn")
	fmt.Println("case <- cn : ", i)
	time.Sleep(time.Second)

}

func main7() {
	var str = "{\n    \"kind\": \"Ingress\",\n    \"apiVersion\": \"networking.k8s.io/v1\",\n    \"metadata\": {\n        \"name\": \"blue-net-ui-dev\",\n        \"namespace\": \"mcsd-cluster\",\n        \"labels\": {\n            \"app\": \"blue-net-ui-dev\",\n            \"svc-namespace\": \"yth-mdgj\"\n        },\n        \"annotations\": {\n            \"field.cattle.io/publicEndpoints\": \"[{\\\"addresses\\\":[\\\"10.68.66.70\\\"],\\\"port\\\":80,\\\"protocol\\\":\\\"HTTP\\\",\\\"serviceName\\\":\\\"mcsd-cluster:ythmdgj-blue-net-ui-fat\\\",\\\"ingressName\\\":\\\"mcsd-cluster:blue-net-ui-dev\\\",\\\"hostname\\\":\\\"blue-net-ui.yth-mdgj.ft.ztosys.com\\\",\\\"path\\\":\\\"/\\\",\\\"allNodes\\\":false}]\",\n            \"nginx.ingress.kubernetes.io/proxy-body-size\": \"200m\"\n        },\n        \"managedFields\": [\n            {\n                \"manager\": \"Mozilla\",\n                \"operation\": \"Update\",\n                \"apiVersion\": \"networking.k8s.io/v1\",\n                \"time\": \"2023-03-02T05:39:07Z\",\n                \"fieldsType\": \"FieldsV1\",\n                \"fieldsV1\": {\n                    \"f:metadata\": {\n                        \"f:annotations\": {\n                            \".\": {},\n                            \"f:kubesphere.io/creator\": {},\n                            \"f:nginx.ingress.kubernetes.io/proxy-body-size\": {}\n                        },\n                        \"f:labels\": {\n                            \".\": {},\n                            \"f:app\": {},\n                            \"f:svc-namespace\": {}\n                        }\n                    },\n                    \"f:spec\": {\n                        \"f:ingressClassName\": {},\n                        \"f:rules\": {}\n                    }\n                }\n            },\n            {\n                \"manager\": \"agent\",\n                \"operation\": \"Update\",\n                \"apiVersion\": \"networking.k8s.io/v1\",\n                \"time\": \"2023-03-02T05:39:11Z\",\n                \"fieldsType\": \"FieldsV1\",\n                \"fieldsV1\": {\n                    \"f:metadata\": {\n                        \"f:annotations\": {\n                            \"f:field.cattle.io/publicEndpoints\": {}\n                        }\n                    }\n                }\n            },\n            {\n                \"manager\": \"nginx-ingress-controller\",\n                \"operation\": \"Update\",\n                \"apiVersion\": \"networking.k8s.io/v1\",\n                \"time\": \"2023-03-02T05:39:58Z\",\n                \"fieldsType\": \"FieldsV1\",\n                \"fieldsV1\": {\n                    \"f:status\": {\n                        \"f:loadBalancer\": {\n                            \"f:ingress\": {}\n                        }\n                    }\n                },\n                \"subresource\": \"status\"\n            }\n        ]\n    },\n    \"spec\": {\n        \"ingressClassName\": \"nginx\",\n        \"rules\": [\n            {\n                \"host\": \"blue-net-ui.yth-mdgj.ft.ztosys.com\",\n                \"http\": {\n                    \"paths\": [\n                        {\n                            \"path\": \"/\",\n                            \"pathType\": \"ImplementationSpecific\",\n                            \"backend\": {\n                                \"service\": {\n                                    \"name\": \"ythmdgj-blue-net-ui-fat\",\n                                    \"port\": {\n                                        \"number\": 8080\n                                    }\n                                }\n                            }\n                        }\n                    ]\n                }\n            }\n        ]\n    }\n}"
	//var result []byte
	bytes, _ := json.Marshal(&str)
	fmt.Println(bytes)

}

func main8() {

	//开启多个协程，直到所有的resp都拿到了才结束，对于resp进行组装

	//写一个reactor模型，起100个协程，接收到事件之后丢入 chan进行处理，然后便利chan

	var wg = sync.WaitGroup{}
	wg.Add(1)

	for i := 0; i < 10000; i++ {
		IntChan <- i
	}

	// create a cancel channel
	cancelChan := make(chan struct{})

	// start the goroutine passing it the cancel channel
	go worker(cancelChan)

	wg.Wait()
	// to cancel the worker, close the cancel channel

	//for {
	//	select {}
	//}
	//close(cancelChan)

}

var IntChan = make(chan int, 10000)

func worker(cancelChan <-chan struct{}) {
	for {
		select {
		case <-cancelChan:
			return

		case num := <-IntChan:

			go process(num)
		}
	}
}

func process(num int) {
	time.Sleep(time.Second / 10)
	fmt.Println("工作队列输出值为: ", num)
}

func main9() {

	m := make(map[string]string, 10)

	m["aaa"] = "bbb"

	//https://zhuanlan.zhihu.com/p/495998623 go map数据结构
	// go map底层是hmap，也就是维护的bmap桶结构，每个桶8对K V 以及K的hash高8位，
	// 如果桶满了，就会通过overflow 去找到下一个桶，也即溢出桶；

	var num int32 = 0

	for i := 0; i < 100; i++ {
		atomic.AddInt32(&num, 1)
	}

}

var count int32

// 方式2
func add2(wg *sync.WaitGroup) {
	defer wg.Done()
	//现场安全的add，具有原子性
	atomic.AddInt32(&count, 1)
	//count++
}

// 修改方式1
func add(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		//把该数字真实值与获取到的数字值进行比较，如果两个一样，则把真实值+1，并返回true，否则不变，且返回false
		if atomic.CompareAndSwapInt32(&count, count, count+1) {
			break
		}
	}
}

func main88() {
	wg := sync.WaitGroup{}
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go add(&wg)
	}
	wg.Wait()
	fmt.Println(count)
}
