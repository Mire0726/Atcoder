package main

import (
    "fmt"
		"bufio"
		"os"
		"strings"
		"sort"
		"strconv"
		)
    
func main() {
    var N int
    fmt.Scan(&N)
    
    scanner := bufio.NewScanner(os.Stdin)
    
    m := map[int]int{}
    
    for i:=0;i<N;i++{
    scanner.Scan()
    inputs := strings.Split(scanner.Text(), "")
    var cnt int
      for j:=0;j<N;j++{
        if inputs[j]=="o"{
         cnt++
      }
      }
    m[i+1] = cnt
    }
   type kv struct {
        Key   int
        Value int
    }

    var ss []kv
    for k, v := range m {
        ss = append(ss, kv{k, v})
    }

    sort.Slice(ss, func(i, j int) bool {
      if ss[i].Value == ss[j].Value {
            return ss[i].Key < ss[j].Key
        }
        return ss[i].Value > ss[j].Value
    })
    
    arr:=make([]string,0)
    
    for _, kv := range ss {
        s:=strconv.Itoa(kv.Key)
        arr = append(arr,s)
    }
    result:=strings.Join(arr," ")
    fmt.Println(result)
}
  