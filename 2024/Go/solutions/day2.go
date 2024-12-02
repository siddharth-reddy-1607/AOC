package solutions

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func checkIncreasing(nums []int) bool{
    for i := range(len(nums) - 1){
        if nums[i] >= nums[i+1] || nums[i+1] - nums[i] > 3{
            return false
        }
    }
    return true
}

func checkDecreasing(nums []int) bool{
    for i := range(len(nums) - 1){
        if nums[i] <= nums[i+1] || nums[i] - nums[i+1] > 3{
            return false
        }
    }
    return true
}

func Day2(){
    file,err := os.Open("../input.txt")
    if err != nil{
        log.Panicf("Error while opening file :%v",err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    part_1_answer,part_2_answer := 0,0
    for scanner.Scan(){
        line := scanner.Text()
        reportStr := strings.Split(line," ")
        reportInt := []int{}
        for _,str := range(reportStr){
            num,_ := strconv.Atoi(str)
            reportInt = append(reportInt, num)
        }
        fmt.Println(reportInt)
        if len(reportInt) == 1{
            part_1_answer += 1
            part_1_answer += 1
            continue
        }
        safe := false
        if checkIncreasing(reportInt){
            fmt.Println("Report is safely increasing")
            safe = true
        }else if checkDecreasing(reportInt){
            fmt.Println("Report is safely decreasing")
            safe = true
        }
        if safe{
            part_1_answer += 1
            part_2_answer += 1
        }else{
            //Assume increasing aftering removing one number
            removedIdx := -1
            for idx := range(len(reportInt) - 1){
                if reportInt[idx] >= reportInt[idx+1] || reportInt[idx+1] - reportInt[idx] > 3{
                    removedIdx = idx
                    break
                }
            }
            //Remove idx
            newArr := []int{}
            for idx,num := range(reportInt){
                if idx == removedIdx{
                    continue
                }
                newArr = append(newArr, num)
            }
            if checkIncreasing(newArr){
                safe = true
            }
            newArr = []int{}
            //Remove idx + 1
            for idx,num := range(reportInt){
                if idx == removedIdx+1{
                    continue
                }
                newArr = append(newArr, num)
            }
            if checkIncreasing(newArr){
                safe = true
            }
            //Assume decreasing aftering removing one number
            removedIdx = -1
            for idx := range(len(reportInt) - 1){
                if reportInt[idx] <= reportInt[idx+1] || reportInt[idx] - reportInt[idx+1] > 3{
                    removedIdx = idx
                    break
                }
            }
            //Remove idx
            newArr = []int{}
            for idx,num := range(reportInt){
                if idx == removedIdx{
                    continue
                }
                newArr = append(newArr, num)
            }
            if checkDecreasing(newArr){
                safe = true
            }
            newArr = []int{}
            //Remove idx + 1
            for idx,num := range(reportInt){
                if idx == removedIdx+1{
                    continue
                }
                newArr = append(newArr, num)
            }
            if checkDecreasing(newArr){
                safe = true
            }
            if safe{
                part_2_answer += 1
            }
        } 
    }
    fmt.Printf("Part 1 = %v, Part 2 = %v\n",part_1_answer,part_2_answer)
}
