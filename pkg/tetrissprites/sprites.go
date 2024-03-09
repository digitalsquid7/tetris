package tetrissprites

import (
	"fmt"
	"image"
	_ "image/png"
	"os"

	"github.com/gopxl/pixel"
)

type Name string

const (
	BlueBlock       Name = "Blue Block"
	CyanBlock       Name = "Cyan Block"
	GreenBlock      Name = "Green Block"
	PinkBlock       Name = "Pink Block"
	PurpleBlock     Name = "Purple Block"
	RedBlock        Name = "Red Block"
	YellowBlock     Name = "Yellow Block"
	BlueTetromino   Name = "Blue Tetromino"
	CyanTetromino   Name = "Cyan Tetromino"
	GreenTetromino  Name = "Green Tetromino"
	PinkTetromino   Name = "Pink Tetromino"
	PurpleTetromino Name = "Purple Tetromino"
	RedTetromino    Name = "Red Tetromino"
	YellowTetromino Name = "Yellow Tetromino"
	Border          Name = "Border"
	Hold            Name = "Hold"
	Next            Name = "Next"
)

var fileByName = map[Name]string{
	BlueBlock:       "assets/images/blocks/blue_block/blue_block.png",
	CyanBlock:       "assets/images/blocks/cyan_block/cyan_block.png",
	GreenBlock:      "assets/images/blocks/green_block/green_block.png",
	PinkBlock:       "assets/images/blocks/pink_block/pink_block.png",
	PurpleBlock:     "assets/images/blocks/purple_block/purple_block.png",
	RedBlock:        "assets/images/blocks/red_block/red_block.png",
	YellowBlock:     "assets/images/blocks/yellow_block/yellow_block.png",
	BlueTetromino:   "assets/images/blocks/blue_block/blue_T1.png",
	CyanTetromino:   "assets/images/blocks/cyan_block/cyan_Z1.png",
	GreenTetromino:  "assets/images/blocks/green_block/green_l1.png",
	PinkTetromino:   "assets/images/blocks/pink_block/pink_J1.png",
	PurpleTetromino: "assets/images/blocks/purple_block/purple_I1.png",
	RedTetromino:    "assets/images/blocks/red_block/red_O1.png",
	YellowTetromino: "assets/images/blocks/yellow_block/yellow_S1.png",
	Border:          "assets/images/UI/border/border.png",
	Hold:            "assets/images/UI/hold/hold.png",
	Next:            "assets/images/UI/next/next.png",
}

type Sprites map[Name]*pixel.Sprite

func Load() (Sprites, error) {
	sprites := make(Sprites)
	for name, file := range fileByName {
		picture, err := loadPicture(file)
		if err != nil {
			return nil, fmt.Errorf("load image: %w", err)
		}

		sprites[name] = pixel.NewSprite(picture, picture.Bounds())
	}

	return sprites, nil
}

func loadPicture(path string) (pic pixel.Picture, err error) {
	var file *os.File
	file, err = os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open file: %w", err)
	}

	defer func() {
		closeErr := file.Close()
		if err != nil {
			err = closeErr
		}
	}()

	var img image.Image
	img, _, err = image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("decode image: %w", err)
	}

	return pixel.PictureDataFromImage(img), nil
}
