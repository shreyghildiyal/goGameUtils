package sprites

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type SpriteHandler struct {
	spriteMap map[string]*ebiten.Image
}

func GetSpriteHandler(folderPath string) (SpriteHandler, error) {
	sh := SpriteHandler{}
	sh.spriteMap = map[string]*ebiten.Image{}

	entries, err := os.ReadDir(folderPath)
	if err != nil {
		return sh, err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			fileName := entry.Name()

			err := loadImage(sh, folderPath, fileName)
			if err != nil {
				return sh, errors.New("unable to read file with name: " + fileName)
			}

		}
	}

	return sh, nil
}

func loadImage(sh SpriteHandler, folderPath, fileName string) error {
	parts := strings.Split(fileName, ".")
	parts = parts[:len(parts)-1]
	key := strings.Join(parts, ".")
	filePath := fmt.Sprintf("%s/%s", folderPath, fileName)
	img, _, err := ebitenutil.NewImageFromFile(filePath)
	if err != nil {
		return err
	}
	sh.spriteMap[key] = img
	return nil
}

func (sh *SpriteHandler) GetSprite(key string) (*ebiten.Image, error) {

	img, ok := sh.spriteMap[key]
	if !ok {
		return nil, errors.New("no such sprite")
	}

	return img, nil

}
