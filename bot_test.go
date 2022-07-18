package bot_test

import (
	"fmt"
	"testing"

	"github.com/xLanStar/bot"
	"github.com/xLanStar/bot/keys"
)

func test(s interface{}) {
	fmt.Println("test:", s)
}

func TestNewEngine(t *testing.T) {
	bot.Register(keys.ModAlt, keys.KeyA, test, "測試")
	bot.Run()
	fmt.Println()
}
