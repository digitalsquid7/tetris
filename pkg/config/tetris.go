package config

type Tetris struct {
	EnableSounds bool
}

func LoadTetris() *Tetris {
	return &Tetris{
		EnableSounds: false,
	}
}
