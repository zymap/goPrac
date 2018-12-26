package main

import (
	"fmt"
	"strconv"
)

func main() {
	//sig :
	// = "30" +
	//	"44" +
	//	"02" +
	//	"20" +
	//	"4655d1e61e06765e64993504fc0c5845555ad7a9380a1cf32c3878f38ab0d760" +
	//	"02" +
	//	"20" +
	//	"0c8afe6cad94f3f8403529158f4e0663d7fad14c3b2515e2e69cf42ee4a01ae5"

	sig := "3045022100d83c96e2656d8c91bf508c4dda68e13f6ea55cfd728e9f55d841d9e32d9325d302201673c42ba6b6546bda1fa0e072c5a423cb02d156406c8a5b59310aa86cab4af701"
	fmt.Println(sig[:2])
	fmt.Println(sig[2:4])
	fmt.Println(sig[4:6])
	fmt.Println(sig[6:8])
	length, _ := strconv.ParseInt(sig[6:8], 16, 64)
	fmt.Println("======", length)
	end := 8 + length*2
	fmt.Println(sig[8:end])
	fmt.Println(sig[end : end+2])
	fmt.Println(sig[end+2 : end+4])
	pos := end + 4
	length, _ = strconv.ParseInt(sig[end+2:end+4], 16, 32)
	fmt.Println("======", length)
	end = end + 4 + length*2
	fmt.Println(sig[pos:end])
	fmt.Println(sig[end:len(sig)])
	fmt.Println("=======================")
	sig = "3045022100cbe66a3d59d39f615034848680f0dc2bc6f278b93cc96f275b183381733be47b022029f651ee9ee856242b47b071f725fd1499ec28d2d7832bbb86a5fb6c4f73481201"
	fmt.Println(sig[:2])
	fmt.Println(sig[2:4])
	fmt.Println(sig[4:6])
	fmt.Println(sig[6:8])
	length, _ = strconv.ParseInt(sig[6:8], 16, 64)
	fmt.Println("======", length)
	end = 8 + length*2
	fmt.Println(sig[8:end])
	fmt.Println(sig[end : end+2])
	fmt.Println(sig[end+2 : end+4])
	pos = end + 4
	length, _ = strconv.ParseInt(sig[end+2:end+4], 16, 32)
	fmt.Println("======", length)
	end = end + 4 + length*2
	fmt.Println(sig[pos:end])
	fmt.Println(sig[end:len(sig)])
}
