package db

import (
	"UnionDesk/pkg/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

const getGroupByIDSQL = "SELECT * FROM groups WHERE id = $1 LIMIT 1;"

func GetGroup(conn *pgx.Conn, uuid uuid.UUID) (*model.Group, error) {
	return getRow[model.Group](conn, getGroupByIDSQL, uuid)
}

const getGroupByRoleIDSQL = "SELECT * FROM groups WHERE role_id = $1 LIMIT 1;"

func GetGroupByRoleID(conn *pgx.Conn, roleId string) (*model.Group, error) {
	return getRow[model.Group](conn, getGroupByRoleIDSQL, roleId)
}
