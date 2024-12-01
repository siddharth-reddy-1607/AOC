package solutions

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func abs(x int) int{
    if x > 0{
        return x
    }
    return -x
} 

func Day1(){
    file,err := os.Open("../input.txt") 
    defer file.Close()
    if err != nil{
        log.Fatalf("Error while opening input file :%v",err)
    }
    scanner := bufio.NewScanner(file)
    leftArr,rightArr := []int{},[]int{}
    hashmap := make(map[int]int)
    for scanner.Scan(){
        splittedString := strings.Split(scanner.Text(),"   ")
        left,_ := strconv.Atoi(splittedString[0])
        right,_ := strconv.Atoi(splittedString[1])
        leftArr = append(leftArr,left)
        rightArr = append(rightArr,right) 
        hashmap[right] += 1
    }
    slices.Sort(leftArr)
    slices.Sort(rightArr)
    part_1_answer := 0
    part_2_answer := 0
    for idx := range(len(leftArr)){
        part_1_answer += abs(leftArr[idx] - rightArr[idx])
        part_2_answer += (leftArr[idx] * hashmap[leftArr[idx]])
    }
    fmt.Printf("Part 1 Answer = %v, Part 2 Answer = %v\n",part_1_answer,part_2_answer)
}
