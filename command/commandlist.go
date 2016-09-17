package command


func ProcessCommandInput(input string) {
    switch input {
    case north:
        //move player up one tile
        fmt.Println("player move north\n")
        movePlayer(0, 1)
    case south:
        //move player down one tile
        fmt.Println("player move south\n")
        movePlayer(0,-1)
    case east:
        //move player right one tile
        fmt.Println("player move east\n")
        movePlayer(1,0)
    case west:
        //move player left one tile
        fmt.Println("player move west\n")
        movePlayer(-1,0)






    default:
        fmt.Println("default exec\n")
    }
}

func movePlayer(changeX int, changeY int) {

}