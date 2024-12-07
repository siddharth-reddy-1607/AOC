package solutions

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func concat(a int64,b int64) int64{
    aStr,bStr := strconv.FormatInt(a,10),strconv.FormatInt(b,10)
    ab,_ := strconv.ParseInt(aStr+bStr,10,64)
    return ab
}

func isPossibleConcat(idx int,a,b int64,nums []int64,target int64) bool{
    if idx == len(nums){
        return a+b == target || a*b == target || concat(a,b) == target
    }
    return isPossibleConcat(idx+1,a*b,nums[idx],nums,target) || isPossibleConcat(idx+1,a+b,nums[idx],nums,target) || isPossibleConcat(idx+1,concat(a,b),nums[idx],nums,target)
}

func isPossible(idx int,a,b int64,nums []int64,target int64) bool{
    if idx == len(nums){
        return a+b == target || a*b == target
    }
    return isPossible(idx+1,a*b,nums[idx],nums,target) || isPossible(idx+1,a+b,nums[idx],nums,target)
}

func Day7(){
    var part1 int64 = 0
    var part2 int64 = 0
    file,err := os.Open("../input.txt")
    if err != nil{
        log.Fatalf("Error while opening the file: %v",err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan(){
        splitByColon := strings.Split(scanner.Text(),":")
        target,_ := strconv.ParseInt(splitByColon[0],10,64)
        numbers := []int64{}
        for _,numStr := range(strings.Split(strings.Trim(splitByColon[1]," ")," ")){
            num,_ := strconv.ParseInt(numStr,10,64)
            numbers = append(numbers, num)
        }
        if isPossible(2,numbers[0],numbers[1],numbers,target){
            part1 += target
        }
        if isPossibleConcat(2,numbers[0],numbers[1],numbers,target){
            part2 += target
        }
    }
    fmt.Printf("Part 1: %v,Part 2: %v",part1,part2)
}
