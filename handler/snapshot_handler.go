package handler

import "github.com/duonghds/ledis/ledis_global"

func handleSaveCommand(globalService ledis_global.GlobalService) string {
	err := globalService.Save()
	if err != nil {
		return err.Error()
	}
	return "Snapshotted"
}

func handleRestoreCommand(globalService ledis_global.GlobalService) string {
	err := globalService.Restore()
	if err != nil {
		return err.Error()
	}
	return "Restored"
}
