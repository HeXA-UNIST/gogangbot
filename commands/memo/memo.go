package memo

import (
	"bytes"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/HeXA-UNIST/gogangbot/store"
	"github.com/fabioxgn/go-bot"
)

const (
	insertMemoDesc = "메모를 추가합니다"
	viewMemoDesc   = "메모를 조회합니다"
	deleteMemoDesc = "메모를 하나 삭제합니다"
	clearMemoDesc  = "메모를 전부삭제 합니다"
)

const (
	insertMemoUsage = "key value"
	viewMemoUsage   = "key [, offset]"
	deleteMemoUsage = "key"
	clearMemoUsage  = "key"
)

var (
	db *sql.DB = nil
)

func formatUsageError(msg string) error {
	return fmt.Errorf("> Usage: %s", msg)
}

func insertMemo(command *bot.Cmd) (msg string, err error) {
	msgs := strings.SplitN(command.FullArg, " ", 2)
	if len(msgs) < 2 {
		return "", formatUsageError(insertMemoUsage)
	}

	db, err := store.Instance()
	if err != nil {
		return "", err
	}

	_, err = db.Exec("INSERT INTO `memo` (`key`, `value`, `creator`) VALUES (?, ?, ?)", msgs[0], msgs[1], command.Nick)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("> 메모가 추가되었습니다 [%s]", msgs[0]), nil
}

func viewMemo(command *bot.Cmd) (msg string, err error) {
	msgs := strings.SplitN(command.FullArg, " ", 2)
	if len(msgs) == 0 {
		return "", formatUsageError(viewMemoUsage)
	}

	offset := 0
	if len(msgs) == 2 {
		offset, err = strconv.Atoi(msgs[1])
		if err != nil {
			return "", formatUsageError(viewMemoUsage)
		}
	}

	db, err := store.Instance()
	if err != nil {
		return "", err
	}

	var n int
	err = db.QueryRow("SELECT count(*) FROM `memo` WHERE `key`=?", msgs[0]).Scan(&n)
	if err != nil {
		return "", err
	}

	rows, err := db.Query("SELECT `value` FROM `memo` WHERE `key`=? LIMIT ?, 10",
		msgs[0], offset)

	if err != nil {
		return "", err
	}
	defer rows.Close()

	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("> %s - %d개 찾음\n", msgs[0], n))

	for rows.Next() {
		var value string
		if err := rows.Scan(&value); err != nil {
			return "", err
		}
		buffer.WriteString(fmt.Sprintf("> * %s\n", value))
	}
	return buffer.String(), nil
}

func deleteMemo(command *bot.Cmd) (msg string, err error) {
	msgs := strings.SplitN(command.FullArg, " ", 1)
	if len(msgs) < 1 {
		return "", formatUsageError(deleteMemoUsage)
	}

	db, err := store.Instance()
	if err != nil {
		return "", err
	}

	_, err = db.Exec("DELETE FROM `memo` where `key`=? LIMIT 1", msgs[0])
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("> 메모가 삭제되었습니다 [%s]", msgs[0]), nil
}

func clearMemo(command *bot.Cmd) (msg string, err error) {
	msgs := strings.SplitN(command.FullArg, " ", 1)
	if len(msgs) < 1 {
		return "", formatUsageError(deleteMemoUsage)
	}

	db, err := store.Instance()
	if err != nil {
		return "", err
	}

	_, err = db.Exec("DELETE FROM `memo` where `key`=?", msgs[0])
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("> 메모가 모두 삭제되었습니다 [%s]", msgs[0]), nil
}

func init() {
	bot.RegisterCommand("메모", insertMemoDesc, insertMemoUsage, insertMemo)
	bot.RegisterCommand("ㅁㅁ", insertMemoDesc, insertMemoUsage, insertMemo)
	bot.RegisterCommand("메보", viewMemoDesc, viewMemoUsage, viewMemo)
	bot.RegisterCommand("ㅁㅂ", viewMemoDesc, viewMemoUsage, viewMemo)
	bot.RegisterCommand("메삭", deleteMemoDesc, deleteMemoUsage, deleteMemo)
	bot.RegisterCommand("ㅁㅅ", deleteMemoDesc, deleteMemoUsage, deleteMemo)
	// bot.RegisterCommand("메클", clearMemoDesc, clearMemoUsage, clearMemo)
	// bot.RegisterCommand("ㅁㅋ", clearMemoDesc, clearMemoUsage, clearMemo)
}
