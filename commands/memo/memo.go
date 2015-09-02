package memo

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/fabioxgn/go-bot"
	_ "github.com/lib/pq"
)

const (
	insertMemoDesc = "메모를 추가합니다"
	viewMemoDesc   = "메모를 조회합니다"
	deleteMemoDesc = "메모를 삭제합니다"
)

const (
	insertMemoUsage = "!메모 헥사 유니스트 컴퓨터 동아리"
	viewMemoDesc    = "!메보 헥사"
	deleteMemoDesc  = "!메삭 헥사"
)

var (
	db *sql.DB = nil
)

func formatUsageError(msg string) error {
	return fmt.Errorf("> Usage: %s", msg)
}

func insertMemo(command *bot.Cmd) (msg string, err error) {
	msgs := strings.SplitN(command.Message, " ", 2)
	if len(msgs) < 2 {
		return "", formatUsageError(insertMemoUsage)
	}
	return fmt.Sprintf("> 메모가 추가되었습니다 [%s]", msgs[0])
}

func viewMemo(command *bot.Cmd) (msg string, err error) {
	msgs := strings.SplitN(command.Message, " ", 2)
	if len(msgs) < 2 {
		return "", formatUsageError(viewMemoUsage)
	}
	return fmt.Sprintf("> 메모가 추가되었습니다 [%s]", msgs[0])
}

func deleteMemo(command *bot.Cmd) (msg string, err error) {
	msgs := strings.SplitN(command.Message, " ", 2)
	if len(msgs) < 2 {
		return "", formatUsageError(deleteMemoUsage)
	}
	return fmt.Sprintf("> 메모가 삭제되었습니다 [%s]", msgs[0])
}

func init() {
	bot.RegisterCommand("메모", insertMemoDesc, insertMemoUsage, insertMemo)
	bot.RegisterCommand("ㅁㅁ", insertMemoDesc, insertMemoUsage, insertMemo)
	bot.RegisterCommand("메보", viewMemoDesc, viewMemoUsage, viewMemo)
	bot.RegisterCommand("ㅁㅂ", viewMemoDesc, viewMemoUsage, viewMemo)
	bot.RegisterCommand("메삭", deleteMemoDesc, deleteMemoUsage, deleteMemo)
	bot.RegisterCommand("ㅁㅅ", deleteMemoDesc, deleteMemoUsage, deleteMemo)
}
