package game

type Player struct {
    ID string
    X  int
    Y  int
    Score int
}

func MovePlayer(player *Player, dx, dy int) {
    player.X += dx
    player.Y += dy

    // Boundary checks
    if player.X < 0 {
        player.X = 0
    }
    if player.Y < 0 {
        player.Y = 0
    }
    if player.X > 800 {
        player.X = 800
    }
    if player.Y > 600 {
        player.Y = 600
    }
}

func CheckCollision(player *Player, ghostX, ghostY int) bool {
    return player.X == ghostX && player.Y == ghostY
}
