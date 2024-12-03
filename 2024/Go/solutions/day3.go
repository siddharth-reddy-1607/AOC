package solutions

import (
	"fmt"
	"log"
	"os"
    "errors"
)

func parseDo(data []byte, length int,idxPtr *int) (bool,error){
    idx := *idxPtr
    defer func(){
            *idxPtr = idx
    }()
    do := "do()"
    dont := "don't()"
    if idx+6 <= length && string(data[idx:idx+7]) == dont{
        return false,nil
    }
    if idx+3 <= length && string(data[idx:idx+4]) == do {
        return true,nil
    }
    return false,errors.New("Didn't find do/don't")
}

func parseMult(data []byte,length int,idxPtr *int) (int){
    idx := *idxPtr
    defer func(){
            *idxPtr = idx
    }()
    for idx < length{
        if (idx + 3 >= length) || data[idx+1] != 'u' || data[idx+2] != 'l' || data[idx+3] != '('{
            idx += 1
            break
        }
        curIdx := idx + 4
        num1 := 0
        mult := 10
        commaFound := false
        digitFound := false
        for curIdx < length{
            if data[curIdx] == ','{
                commaFound = true
                curIdx += 1
                break
            }
            if data[curIdx] < '0' || data[curIdx] > '9'{
                break
            }
            digitFound = true
            num1 = (num1*mult) + int(data[curIdx] - '0')
            curIdx += 1
        }
        if !commaFound || !digitFound || num1 > 999{
            idx = curIdx
            break
        }
        num2 := 0
        bracketFound := false
        digitFound = false
        for curIdx < length{
            if data[curIdx] == ')'{
                bracketFound = true
                curIdx += 1
                break
            }
            if data[curIdx] < '0' || data[curIdx] > '9'{
                break
            }
            digitFound = true
            num2 = (num2*mult) + int(data[curIdx] - '0')
            curIdx += 1
        }
        if !bracketFound || !digitFound || num2 > 999{
            idx = curIdx
            break
        }
        idx = curIdx
        return num1*num2
    }
    return 0
}

func Day3(){
    data,err := os.ReadFile("../input.txt")
    if err != nil{
        log.Fatalf("Error reading data: %v\n",err)
    }
    part_1_answer,part_2_answer := 0,0
    length := len(data)
    idx := 0
    val := 0
    do := true
    for idx < length{
        if data[idx] != 'm' && data[idx] != 'd'{
            idx += 1
            continue
        }
        if data[idx] == 'm'{
            val  = parseMult(data,length,&idx)
            part_1_answer += val
            if do{
                part_2_answer += val
            }
        }else{
            doOrDont,err := parseDo(data,length,&idx) 
            idx += 1
            if err != nil{
                continue
            }
            do = doOrDont
        }
    }
    fmt.Printf("Part 1: %v,Part 2: %v",part_1_answer,part_2_answer)

}
