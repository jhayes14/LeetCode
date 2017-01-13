/* Number Complement
   Given a positive integer, output its complement number. The complement strategy is to flip the bits of its binary representation.
*/

package main

import(
    "fmt"
    //m "math"
    "strconv"
    "strings"
)


func NumCompl( x int64) int64{
    var binx string
    binx = strconv.FormatInt(x, 2)
    //fmt.Println(binx)
    xcompl := ""
    for i := 0; i < len(binx); i++{
        if binx[i:i+1] == "0"{
            xcompl = xcompl + "1"
        } else if binx[i:i+1] == "1"{
            xcompl = xcompl + "0"
        }
    }
    //fmt.Println(xcompl)
    com, _ := strconv.ParseInt(xcompl, 2, 64)
    return com
}

func main(){
    x := int64(2)
    fmt.Println(NumCompl(x))
}
