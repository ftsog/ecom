package handlers

import (
	"github.com/ftsog/ecom/models"
	"github.com/ftsog/ecom/utils"
)

type Handler struct {
	Db     *models.Model
	Logger *utils.Logger
}
