// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"wktline-server/internal/dao/internal"
)

// internalProjectGroupDao is internal type for wrapping internal DAO implements.
type internalProjectGroupDao = *internal.ProjectGroupDao

// projectGroupDao is the data access object for table model_project.
// You can define custom methods on it to extend its functionality as you wish.
type projectGroupDao struct {
	internalProjectGroupDao
}

var (
	// ProjectGroup is globally public accessible object for table model_project operations.
	ProjectGroup = projectGroupDao{
		internal.NewProjectGroupDao(),
	}
)

// Fill with you ideas below.