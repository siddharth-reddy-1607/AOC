package solutions

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
)

func getStoneAfterBlink(stone string) []string{
    stonesAfterBlink := []string{};
    if stone == "0"{
        stonesAfterBlink = append(stonesAfterBlink, "1")
    }else if len(stone)%2 == 0{
        leftInt,_ := strconv.ParseInt(stone[:len(stone)/2],10,64)  
        left := strconv.FormatInt(leftInt,10)
        rightInt,_ := strconv.ParseInt(stone[len(stone)/2:],10,64)  
        right := strconv.FormatInt(rightInt,10)
        stonesAfterBlink = append(stonesAfterBlink, left)
        stonesAfterBlink = append(stonesAfterBlink, right)
    }else{
        num,_ := strconv.ParseInt(stone,10,64)
        num = num*2024
        stonesAfterBlink = append(stonesAfterBlink, strconv.FormatInt(num,10))
    }
    return stonesAfterBlink
}

func getNumberOfStonesAfterKBlinks(k int,hashmap map[string]int64) int64{
    for _ = range(k){
        newHashmap := make(map[string]int64)
        for stone,freq := range(hashmap){
            for _,val := range(getStoneAfterBlink(stone)){
                newHashmap[val] += freq
            }
        }
        hashmap = newHashmap
    }
    var ans int64 = 0
    for _,val := range(hashmap){
        ans += val
    }
    return ans
}
func Day11(){
    file,err := os.Open("../input.txt")
    if err != nil{
        log.Fatalf("Error while opening file : %v\n",err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    stones := []string{}
    for scanner.Scan(){
        stones = strings.Split(scanner.Text()," ")
    }
    hashmap := make(map[string]int64)
    for _,stone := range(stones){
        hashmap[stone] += 1
    }
    part1 := getNumberOfStonesAfterKBlinks(25,hashmap)
    part2 := getNumberOfStonesAfterKBlinks(75,hashmap)
    fmt.Printf("Part 1:%v, Part 2:%v\n",part1,part2)
}

