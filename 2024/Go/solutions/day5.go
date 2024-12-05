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

func makeOrderingValid(ordering []int,rules map[int]map[int]bool) []int{
    slices.SortFunc(ordering,func(a,b int) int{
        if rules[a][b]{
            return -1
        }
        return 1
    })
    return ordering
}

func checkValidOrdering(ordering []int,rules map[int]map[int]bool) bool{
    for i := range(len(ordering)){
        for j := i+1; j < len(ordering); j++{
            if !rules[ordering[i]][ordering[j]]{
                return false
            }
        }
    }
    return true
}

func Day5(){
    file,err := os.Open("../input.txt")
    if err != nil{
        log.Fatalf("Error while opening file: %v\n",err)
    }
    scanner := bufio.NewScanner(file)
    getOrderings := false
    rules := make(map[int]map[int]bool)
    orderings := [][]int{}
    part1,part2 := 0,0
    for scanner.Scan(){
        line := scanner.Text()
        if line == ""{
            getOrderings = true 
            continue
        }
        if getOrderings{
            ordering := []int{}
            for _,numStr := range(strings.Split(line,",")){
                num,_ := strconv.Atoi(numStr)
                ordering = append(ordering, num)
            }
            orderings = append(orderings, ordering)
        }else{
            xy := strings.Split(line,"|")
            x,_ := strconv.Atoi(xy[0])
            y,_ := strconv.Atoi(xy[1])
            if rules[x] == nil{
                rules[x] = make(map[int]bool)
            }
            rules[x][y] = true
        }
    } 
    for _,ordering := range(orderings){
        if checkValidOrdering(ordering,rules){
            part1 += ordering[len(ordering)/2]
        }else{
            ordering = makeOrderingValid(ordering,rules)
            if !checkValidOrdering(ordering,rules){
               fmt.Println("Ordering still invalid after 'making it valid'")  
            }
            part2 += ordering[len(ordering)/2]
        }
    }
    fmt.Printf("Part1 : %v, Part2 : %v",part1,part2)

}
