package main
import (
    "fmt"
    "time")

func sum(a []int, c chan int) {
    sum := 0
    for _ , v := range a {
        sum += v
    }
    // チャネルに送信(ブロックされるのでraceしない
    c <- sum
}


func main() {
    a := []int{7, 2, 8, -9, 4, 0}
    // チャネルの作成
    c := make(chan int)
    // 配列の前半分
    go sum(a[:len(a)/2], c)
    // 後半
    go sum(a[len(a)/2:], c)
    // チャネルのそれぞれをx,yに代入
    x, y := <-c, <-c
    fmt.Println(x, y, x+y)
}