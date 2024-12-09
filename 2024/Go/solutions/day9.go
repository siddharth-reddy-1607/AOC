package solutions

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getFilesAndHoles(fileLayoutWithFreeBlocks []string,files, holes *[][]int){
    for idx := 0; idx < len(fileLayoutWithFreeBlocks);{
        i := idx
        if fileLayoutWithFreeBlocks[idx] == "."{
            holeStart := idx
            holeSize := 0
            for i = idx; i < len(fileLayoutWithFreeBlocks); i++{
                if fileLayoutWithFreeBlocks[i] != "."{
                    break
                }
                holeSize += 1
            }
            idx = i
            (*holes) = append(*holes, []int{holeStart,holeSize})
            continue
        }
        fileStart := idx
        fileSize := 0
        for i = idx; i < len(fileLayoutWithFreeBlocks); i++{
            if fileLayoutWithFreeBlocks[i] != fileLayoutWithFreeBlocks[idx]{
                break
            }
            fileSize += 1
        }
        idx = i
        (*files) = append(*files, []int{fileStart,fileSize})
    }
}

func Day9(){
    file,err := os.Open("../input.txt") 
    if err != nil{
        log.Fatalf("Error while opening the file: %v",err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    line := ""
    for scanner.Scan(){
        line = scanner.Text()
    } 
    fileLayoutWithFreeBlocksP1:= []string{}
    fileIdx := 0
    for i := 0; i < len(line); i+= 2{
        if i == len(line){
            break
        }
        if i+1 == len(line){
            for _ = range(line[i] - '0'){
                fileLayoutWithFreeBlocksP1 = append(fileLayoutWithFreeBlocksP1, strconv.Itoa(fileIdx))
            }
            break
        }
        for _ = range(line[i] - '0'){
            fileLayoutWithFreeBlocksP1 = append(fileLayoutWithFreeBlocksP1, strconv.Itoa(fileIdx))
        }
        for _ = range(line[i+1] - '0'){
            fileLayoutWithFreeBlocksP1 = append(fileLayoutWithFreeBlocksP1, ".")
        }
        fileIdx += 1
    }
    fileLayoutWithFreeBlocksP2 := []string{}
    fileLayoutWithFreeBlocksP2 = append(fileLayoutWithFreeBlocksP2, fileLayoutWithFreeBlocksP1...)

    freeBlockPtr,diskBlockPtr := 0,len(fileLayoutWithFreeBlocksP1) - 1;
    for freeBlockPtr < diskBlockPtr{
        for fileLayoutWithFreeBlocksP1[freeBlockPtr] != "."{
            freeBlockPtr += 1
        }
        for fileLayoutWithFreeBlocksP1[diskBlockPtr] == "."{
            diskBlockPtr -= 1
        }
        if freeBlockPtr >= diskBlockPtr{
            break
        }
        fileLayoutWithFreeBlocksP1[freeBlockPtr] = fileLayoutWithFreeBlocksP1[diskBlockPtr]
        fileLayoutWithFreeBlocksP1[diskBlockPtr] = "."
        freeBlockPtr += 1
        diskBlockPtr -= 1
    }

    var part1 int64 = 0
    for i,numStr := range(fileLayoutWithFreeBlocksP1){
        if numStr == "."{
            break
        }
        num,_ := strconv.Atoi(numStr)
        part1 += (int64)(i*num)
    }

    holes := [][]int{}
    files := [][]int{}
    getFilesAndHoles(fileLayoutWithFreeBlocksP2,&files,&holes)

    for i := len(files) - 1; i >= 0; i--{
        fileNumber := fileLayoutWithFreeBlocksP2[files[i][0]]
        fileSize := files[i][1]
        for _,hole := range(holes){
            if hole[0] >= files[i][0]{
                break
            }
            if hole[1] >= fileSize{
                for idx := range(fileSize){
                    fileLayoutWithFreeBlocksP2[hole[0]+idx] = fileNumber
                }     
                for idx := range(fileSize){
                    fileLayoutWithFreeBlocksP2[files[i][0] + idx] = "."
                }
                hole[0] += fileSize
                hole[1] -= fileSize
                break
            }
        }
    }
    var part2 int64 = 0
    for i,numStr := range(fileLayoutWithFreeBlocksP2){
        num,_ := strconv.Atoi(numStr)
        part2 += (int64)(i*num)
    }
    fmt.Printf("Part 1 : %v, Part 2 : %v\n",part1,part2)
}
