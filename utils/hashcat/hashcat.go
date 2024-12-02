package hashcat

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"wails_hashcat/global"
	"wails_hashcat/model"
)

var (
	HT *HashcatTask
)

type HashcatTask struct {
	Ctx     context.Context
	cancel  context.CancelFunc
	CmdChan chan string
	Session int
}

func InitTask(session int) {
	if HT != nil {
		HT.StopTask()
	}
	HT = &HashcatTask{Session: session}
	HT.Ctx, HT.cancel = context.WithCancel(context.Background())
	HT.CmdChan = make(chan string)
}

func (h *HashcatTask) StopTask() {
	if h.Ctx != nil {
		h.cancel()
	}
}

func (h *HashcatTask) StartSession(commandStr string) {
	cmdAll := fmt.Sprintf(" --session %d  --status --status-json --status-timer 2  --quiet --force ", h.Session)
	commandStr += cmdAll
	//h.ExecCommand(commandStr)
	//查看是否已有记录
	showCommandStr := commandStr + " --show"
	h.ExecCommand(showCommandStr, 0)
	var ht model.HashCatTask
	global.DB.Model(&model.HashCatTask{}).Where("id = ?", h.Session).First(&ht)
	if ht.Result == "" {
		h.ExecCommand(commandStr, 1)
	}
}
func (h *HashcatTask) ReStartSession() {
	commandStr := fmt.Sprintf(" --session %d --restore", h.Session)
	h.ExecCommand(commandStr, 1)
}

func (h *HashcatTask) ExecCommand(commandStr string, showLabel int) {
	parts := strings.Fields(commandStr)
	cmd := exec.CommandContext(h.Ctx, "./hashcat", parts[0:]...)
	workHome, _ := os.Getwd()
	cmd.Dir = filepath.Join(workHome, "static/hashcat")
	global.SugarLogger.Debug("StartRun: ", cmd.String())
	stdoutPipe, _ := cmd.StdoutPipe()
	scanner := bufio.NewScanner(stdoutPipe)

	//处理回传数据
	go func() {
		for scanner.Scan() {
			global.SugarLogger.Debug("Scanner: ", scanner.Text())
			var hr HashcatResponse
			err := json.Unmarshal([]byte(scanner.Text()), &hr)
			if err != nil {
				continue
			}
			var ht model.HashCatTask
			global.DB.Model(&model.HashCatTask{}).Where("id = ?", h.Session).First(&ht)
			ht.CMD = commandStr
			ht.CurrentProgress = hr.Progress[0]
			ht.TotalProgress = hr.Progress[1]
			ht.GuessBasePercent = hr.Guess.GuessBasePercent
			ht.GuessModPercent = hr.Guess.GuessModPercent
			ht.Status = model.TaskStatusRunning
			global.DB.Save(&ht)
		}
	}()
	err := cmd.Start()
	if err != nil {
		global.SugarLogger.Error("Error starting command: ", err)
		return
	}
	err = cmd.Wait()
	global.SugarLogger.Info("==========================================================================================")
	//执行结束后处理
	select {
	case <-h.Ctx.Done():
		global.SugarLogger.Debug("Command execution canceled")
		global.DB.Model(&model.HashCatTask{}).Where("session = ?", h.Session).Update("status", model.TaskStatusStop)
	default:
		if err != nil && err.(*exec.ExitError).ExitCode() != 1 {
			if exitErr, ok := err.(*exec.ExitError); ok {
				if exitErr.ExitCode() == 1 {
					h.FinishWork(showLabel)
					return
				}
			}
			global.SugarLogger.Error(cmd.String(), ":   ", err)
			global.DB.Model(&model.HashCatTask{}).Where("session = ?", h.Session).Update("status", model.TaskStatusError)
		} else {
			h.FinishWork(showLabel)
		}
	}
}

func (h *HashcatTask) FinishWork(showLabel int) {
	global.SugarLogger.Debug("Command executed successfully")
	var ht model.HashCatTask
	global.DB.Model(&model.HashCatTask{}).Where("id = ?", h.Session).First(&ht)
	file, _ := os.ReadFile(ht.OutFilePath)
	content := string(file)
	content = strings.TrimSpace(content)
	global.SugarLogger.Debug("content: ", content)
	content = strings.Split(content, "\r\n")[0]
	contentL := strings.Split(content, ":")
	global.SugarLogger.Debug("contentL: ", contentL)
	if len(contentL) >= 2 {
		result := contentL[len(contentL)-1]
		global.DB.Model(&model.HashCatTask{}).Where("id = ?", h.Session).
			Update("status", model.TaskStatusFinish).
			Update("result", result)
	} else if showLabel == 0 {

	} else {
		global.DB.Model(&model.HashCatTask{}).Where("id = ?", h.Session).
			Update("status", model.TaskStatusFinish).
			Update("result", "无匹配")
	}
}
