// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"wktline-server/internal/dao/internal"
)

// internalTaskAttachmentDao is internal type for wrapping internal DAO implements.
type internalTaskAttachmentDao = *internal.TaskAttachmentDao

// taskAttachmentDao is the data access object for table task_attachment.
// You can define custom methods on it to extend its functionality as you wish.
type taskAttachmentDao struct {
	internalTaskAttachmentDao
}

var (
	// TaskAttachment is globally public accessible object for table task_attachment operations.
	TaskAttachment = taskAttachmentDao{
		internal.NewTaskAttachmentDao(),
	}
)

// Fill with you ideas below.
