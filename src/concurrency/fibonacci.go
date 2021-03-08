package concurrency

const MaxFibNumber = 30

type fibResult struct {
	n      int
	fibNum int
}

func Fibonacci(n int) map[int]int {
	result := make(map[int]int, n)
	for i := 0; i <= n; i++ {
		result[i] = computeFibSeq(i)
	}
	return result
}

func FibonacciConcurrent(n int) map[int]int {
	result := make(map[int]int, n)
	fibJobCh := make(chan int, n)
	fibResultCh := make(chan fibResult, n)

	go worker(fibJobCh, fibResultCh)
	go worker(fibJobCh, fibResultCh)
	go worker(fibJobCh, fibResultCh)

	for i := 0; i <= n; i++ {
		fibJobCh <- i
	}
	close(fibJobCh)

	for j := 0; j <= n; j++ {
		fn := <-fibResultCh
		result[fn.n] = fn.fibNum
	}
	return result
}

func computeFibSeq(n int) int {
	if n <= 1 {
		return n
	}
	return computeFibSeq(n-1) + computeFibSeq(n-2)
}

func worker(fibJobCh <-chan int, fibResultCh chan<- fibResult) {
	for n := range fibJobCh {
		fibResultCh <- fibResult{n, computeFibSeq(n)}
	}
}

var fibSeq40 = map[int]int{
	0:  0,
	1:  1,
	2:  1,
	3:  2,
	4:  3,
	5:  5,
	6:  8,
	7:  13,
	8:  21,
	9:  34,
	10: 55,
	11: 89,
	12: 144,
	13: 233,
	14: 377,
	15: 610,
	16: 987,
	17: 1597,
	18: 2584,
	19: 4181,
	20: 6765,
	21: 10946,
	22: 17711,
	23: 28657,
	24: 46368,
	25: 75025,
	26: 121393,
	27: 196418,
	28: 317811,
	29: 514229,
	30: 832040,
	//31:1346269,
	//32:2178309,
	//33:3524578,
	//34:5702887,
	//35:9227465,
	//36:14930352,
	//37:24157817,
	//38:39088169,
	//39:63245986,
	//40:102334155,
	//41:165580141,
	//42:267914296,
	//43:433494437,
	//44:701408733,
	//45:1134903170,
}
