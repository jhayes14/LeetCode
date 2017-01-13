/* Hamming Distance
   The Hamming distance between two integers is the number of positions at which the corresponding bits are different.
   Given two integers x and y, calculate the Hamming distance.
*/

package main

import(
    "fmt"
    //m "math"
    "strconv"
    "strings"
)

func leftPad2Len(s string, padStr string, overallLen int) string {
	var padCountInt int
	padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = strings.Repeat(padStr, padCountInt) + s
	return retStr[(len(retStr) - overallLen):]
}

func max(x, y int) int{
    if x > y{
        return x
    }
    return y
}

func HammDist( x, y int64) int{
    // find Hamming Distance
    var binx string
    var biny string
    binx = strconv.FormatInt(x, 2)
    biny = strconv.FormatInt(y, 2)
    lenbinx := len(binx)
    lenbiny := len(biny)
    if lenbinx > lenbiny{
        biny = leftPad2Len("", "0", (lenbinx-lenbiny)) + biny
    } else if lenbiny > lenbinx {
        binx = leftPad2Len("", "0", (lenbiny-lenbinx)) + binx
    }
    max_l := max(len(binx), len(biny))
    var count int
    for i := 0; i < max_l; i++ {
        binx_char := binx[i:i+1]
        biny_char := biny[i:i+1]
        if binx_char != biny_char{
            count += 1
        }
    }
    return count
}


func main(){
    x := int64(2)
    y := int64(3)
    fmt.Println(HammDist(x,y))
}
