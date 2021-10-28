/*
需求：
计算一个数字的各个位数之和，例如数字123，结果为1+2+3=6
随机生成数字进行计算
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Job struct {
	Id      int
	RandNum int
}

type Result struct {
	// 这里必须传对象实例
	job *Job
	sum int
}

// 创建工作池
// 参数1：开几个协程
func createPool(num int, jobChan chan *Job, resultChan chan *Result) {
	// 根据开协程个数，去运行
	for i := 0; i < num; i++ {
		go func(jobChan chan *Job, resultChan chan *Result) {
			// 执行运算
			// 遍历job管道所有数据，进行相加
			for job := range jobChan {
				// 随机数接过来
				r_num := job.RandNum
				// 随机数每一位相加
				// 定义返回值
				var sum int
				for r_num != 0 {
					tmp := r_num % 10
					sum += tmp
					r_num /= 10
				}
				// 想要的结果是Result
				r := &Result{
					job: job,
					sum: sum,
				}
				// 将运算结果传入管道
				resultChan <- r
			}
		}(jobChan, resultChan)
	}
}

func main() {
	// 需要两个管道
	jobChan := make(chan *Job, 128)
	resultChan := make(chan *Result, 128)

	// 循环创建job，输入到管道
	for id := 1; id <= 100; id++ {
		// 生成随机数
		r_num := rand.Int()
		job := &Job{
			Id:      id,
			RandNum: r_num,
		}
		jobChan <- job
	}

	/*
		go func() {
			for id := 1; id <= 100; id++ {
				// 生成随机数
				r_num := rand.Int()
				job := &Job{
					Id:      id,
					RandNum: r_num,
				}
				jobChan <- job
			}
		}()
	*/

	// 创建工作池
	createPool(5, jobChan, resultChan)
	// 开个打印的协程
	go func(resultChan chan *Result) {
		// 遍历结果管道并打印
		for result := range resultChan {
			fmt.Printf("job id:%v randnum:%v result:%d\n", result.job.Id, result.job.RandNum, result.sum)
		}
	}(resultChan)

	time.Sleep(time.Second)
}
