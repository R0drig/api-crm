// database.go
package main

import (
    "gorm.io/gorm"
)

type Handler struct {
    db *gorm.DB
}

func newHandlers(db *gorm.DB) *Handler {
    return &Handler{db}
}
