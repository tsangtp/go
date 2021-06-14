package main

import (
    "math/rand"
    "fmt"
    "time"
)

var (
    name string
    num int
    timestamp int
    guess int
    id int
    playnum int
    players [] string
)

func main() {
    rand.Seed(time.Now().Unix())
    num := rand.Intn(501)
    // fmt.Println(num)
    fmt.Println("Welcome to the game!")
    fmt.Println("This game is guessing number.(1 to 500)")
    fmt.Println("How many player to play.")
    for true{
        fmt.Println("Please enter the number of player.(At least 1, up to 10)")
        fmt.Scanln(&playnum)
        if playnum >= 1 && playnum <= 10{
            fmt.Printf("Now, We have %d player to play.\n" , playnum)
            break
        }else{
            fmt.Println("Input error!Please try again.")
        }
    }
    var player = 0
    for player < playnum{
        fmt.Printf("Please enter player%d name.\n" , player+1)
        fmt.Scanln(&name)
        players = append(players, name)
        player ++
    }
    var playersid map[string]int
    playersid = make(map[string]int)
    for i , c := range players{

        id := rand.Intn(1000)
        playersid[c] = id
        i++
    }
    for a := range playersid{
        fmt.Println(a , "ID is" , playersid[a])
    }
    for i , x := range players{
        fmt.Printf("The round %d is %s to play.\n" ,i+1 , x)
    }
    for true{
        fmt.Println("Please set the rounds.(Maximum 10 rounds per person)")
        fmt.Scanln(&timestamp)
        if timestamp <= 10  && timestamp != 0{
            timestamp *= len(players)
            fmt.Printf("You set the %d rounds.\n" , timestamp)
            fmt.Printf("You have %d rounds to guess.(each people)\n" , timestamp)
            break
        }else{
            fmt.Println("Input error!Please try again.")
            fmt.Println("The number of rounds you enter must be able to divide the number of players. ")
        }
    }
    max := 500
    min := 1
    i := 0
    for true {

        if i == len(players){
            i = 0
        }
        fmt.Printf("It is %s turn." , players[i])
        fmt.Printf("Please enter the number.(%d to %d)\n" , min , max)
        fmt.Scanln(&guess)
        if guess > min && guess <max{
            if (guess%num == 0) {
                fmt.Printf("%s win!The answer is %d.\n" , players[i] , num )
                break
            }else if (guess > num){
                fmt.Println("Answer is smaller.")
                max = guess
                i++
            }else{
                fmt.Println("Answer is bigger.")
                min = guess
                i++
            }
            timestamp -= 1
            if timestamp == 0{
                fmt.Println("Time up!")
                fmt.Printf("You lose!The number is %d." , num)
                break
            }
        }else{
            fmt.Println("Input error!Please try again.")
            fmt.Printf("Please enter %d to %d." , min , max)
        }

    }

}
