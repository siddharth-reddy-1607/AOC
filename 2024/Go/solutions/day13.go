package solutions

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getMinTokensSolvingLinearEquation(a []int, b[]int,prize []int) int64{
    var prizeX int64 = 10000000000000 + int64(prize[0])
    var prizeY int64 = 10000000000000 +int64(prize[1])
    var x1 int64 = int64(a[0])
    var y1 int64 = int64(a[1])
    var x2 int64 = int64(b[0])
    var y2 int64 = int64(b[1])
    if (y2*prizeX - x2*prizeY )%(y2*x1 - x2*y1) == 0 && (x1*prizeY - y1*prizeX)%(y2*x1 - x2*y1) == 0 {
        return 3 * ((y2*prizeX - x2*prizeY )/(y2*x1 - x2*y1) ) + 1 * ((x1*prizeY - y1*prizeX)/(y2*x1 - x2*y1) )
    }
    return 0
}
func getMinTokensBruteForce(a []int,b []int,prize []int) int{
        minTokens := (1 << 31) - 1
        found := false
        for i := 1; i <= 100; i++{
            for j := 1; j <= 100; j++{
                x := (a[0] * i) + (b[0] * j)
                y := (a[1] * i) + (b[1] * j)
                if x == prize[0] && y == prize[1]{
                    minTokens = min(minTokens,i * 3 + j)
                    found = true
                }
            }
        }
        if found{
            return minTokens
        }
    return 0
}

func Day13(){
    file,err := os.Open("../input.txt")
    if err != nil{
        log.Fatalf("Error while opening the file: %v\n",err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    part1 := 0
    var part2 int64 = 0
    buttonA := [][]int{}
    buttonB := [][]int{}
    prizes := [][]int{}
    for (scanner.Scan()){
        splitByColon := strings.Split(scanner.Text(),":")
        if len(splitByColon) == 1{
            continue
        }
        if splitByColon[0][len(splitByColon[0]) - 1] == 'A'{
            splitByComma := strings.Split(splitByColon[1],",")
            x,_ := strconv.Atoi(strings.Split(strings.Trim(splitByComma[0]," "),"+")[1])
            y,_ := strconv.Atoi(strings.Split(strings.Trim(splitByComma[1]," "),"+")[1])
            buttonA = append(buttonA, []int{x,y})
        }else if splitByColon[0][len(splitByColon[0]) - 1] == 'B'{
            splitByComma := strings.Split(splitByColon[1],",")
            x,_ := strconv.Atoi(strings.Split(strings.Trim(splitByComma[0]," "),"+")[1])
            y,_ := strconv.Atoi(strings.Split(strings.Trim(splitByComma[1]," "),"+")[1])
            buttonB = append(buttonB, []int{x,y})
        }else{
            splitByComma := strings.Split(splitByColon[1],",")
            x,_ := strconv.Atoi(strings.Split(strings.Trim(splitByComma[0]," "),"=")[1])
            y,_ := strconv.Atoi(strings.Split(strings.Trim(splitByComma[1]," "),"=")[1])
            prizes = append(prizes, []int{x,y})
        }
    }
    for idx,prize := range(prizes){
        part1 += getMinTokensBruteForce(buttonA[idx],buttonB[idx],prize)
        part2 += getMinTokensSolvingLinearEquation(buttonA[idx],buttonB[idx],prize)
    }
    fmt.Printf("Part 1: %v,Part 2: %v\n",part1,part2)
}
