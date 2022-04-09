package controllers

import (
	"famous-quote/dtos"
	"famous-quote/errors"
	"famous-quote/logger"
	"famous-quote/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type QuotesController struct {
	Base
	quotesService services.QuotesService
}

// @Description create quotes
// @Tags Quotes
// @Accept json
// @Produce json
// @Success 200 {object} dtos.GetQuotesResponse
// @Failure 401 {object} errors.AppError
// @Failure 500 {object} errors.AppError
// @Router /v1/quotes [get]
func (h *QuotesController) Get(c *gin.Context) {
	var (
		res *dtos.GetQuotesResponse
		err error
	)

	if res, err = h.quotesService.Get(c.Request.Context()); err != nil {
		h.HandleError(c, err)
		return
	}
	logger.Infof("him om")
	h.JSON(c, res)
}

// @Description create quotes
// @Tags Quotes
// @Accept json
// @Produce json
// @Param request	body dtos.CreateQuotesRequest	true "request"
// @Success 200 {object} dtos.CreateQuotesResponse
// @Failure 401 {object} errors.AppError
// @Failure 500 {object} errors.AppError
// @Router /v1/quotes [post]
func (h *QuotesController) Create(c *gin.Context) {
	var (
		req dtos.CreateQuotesRequest
		res *dtos.CreateQuotesResponse
		err error
	)

	if err = c.ShouldBindJSON(&req); err != nil {
		h.HandleError(c, errors.New(errors.ErrInvalidRequest, err.Error()))
		return
	}

	if res, err = h.quotesService.Create(c.Request.Context(), &req); err != nil {
		h.HandleError(c, err)
		return
	}

	h.JSON(c, res)
}

// @Description create quotes
// @Tags Quotes
// @Accept json
// @Produce json
// @Param id	path int64	true "quotes id"
// @Param values	body dtos.LikeQuotesRequest	true "request"
// @Success 200 {object} dtos.LikeQuotesResponse
// @Failure 401 {object} errors.AppError
// @Failure 500 {object} errors.AppError
// @Router /v1/quotes/like/:id [post]
func (h *QuotesController) Like(c *gin.Context) {
	var (
		req dtos.LikeQuotesRequest
		res *dtos.LikeQuotesResponse
		err error
	)

	if req.ID, err = strconv.ParseInt(c.Param("id"), 10, 64); err != nil {
		h.HandleError(c, err)
		return
	}

	if err = c.ShouldBindJSON(&req); err != nil {
		h.HandleError(c, err)
		return
	}

	if res, err = h.quotesService.Like(c.Request.Context(), &req); err != nil {
		h.HandleError(c, err)
		return
	}

	h.JSON(c, res)
}

func NewQuotesController(quotesService services.QuotesService) *QuotesController {
	return &QuotesController{
		quotesService: quotesService,
	}
}
