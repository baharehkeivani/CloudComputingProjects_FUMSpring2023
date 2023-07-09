package controller

import (
	"net/http"
	"strconv"

	"github.com/alireza-dehghan-nayeri/simple-go-rest-api/api/service"
	"github.com/alireza-dehghan-nayeri/simple-go-rest-api/models"
	"github.com/alireza-dehghan-nayeri/simple-go-rest-api/util"
	"github.com/gin-gonic/gin"
)

// TextController -> TextController
type TextController struct {
	service service.TextService
}

// NewTextController : NewTextController
func NewTextController(s service.TextService) TextController {
	return TextController{
		service: s,
	}
}

// GetTexts : GetTexts controller
func (p TextController) GetTexts(ctx *gin.Context) {
	var texts models.Text

	keyword := ctx.Query("keyword")

	println("-----controller")
	println(keyword)

	data, total, err := p.service.FindAll(texts, keyword)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to find questions")
		return
	}
	respArr := make([]map[string]interface{}, 0)

	for _, n := range *data {
		resp := n.ResponseMap()
		respArr = append(respArr, resp)
	}

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Text result set",
		Data: map[string]interface{}{
			"rows":       respArr,
			"total_rows": total,
		}})
}

// AddText : AddText controller
func (p *TextController) AddText(ctx *gin.Context) {
	var text models.Text
	ctx.ShouldBindJSON(&text)

	if text.Title == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Title is required")
		return
	}
	if text.Body == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Body is required")
		return
	}

	err := p.service.Save(text)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to create quote")
		return
	}

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Successfully Created text",
		Data:    &text})
}

// GetText : get text by id
func (p *TextController) GetText(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to int64
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "id invalid")
		return
	}
	var text models.Text
	text.ID = id
	foundText, err := p.service.Find(text)
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Error Finding Text")
		return
	}

	response := foundText.ResponseMap()

	c.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Result set of Text",
		Data:    &response})

}

// DeleteText : Deletes Text
func (p *TextController) DeleteText(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to uint64
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "id invalid")
		return
	}
	err = p.service.Delete(id)

	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Failed to delete Text")
		return
	}
	response := &util.Response{
		Success: true,
		Message: "Deleted Sucessfully"}
	c.JSON(http.StatusOK, response)
}

// UpdateText : get update by id
func (p TextController) UpdateText(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	var text models.Text
	text.ID = id

	textRecord, err := p.service.Find(text)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Text with given id not found")
		return
	}
	ctx.ShouldBindJSON(&textRecord)

	if textRecord.Title == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Title is required")
		return
	}
	if textRecord.Body == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Body is required")
		return
	}

	if err := p.service.Update(textRecord); err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to store Text")
		return
	}
	response := textRecord.ResponseMap()

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Successfully Updated Text",
		Data:    response,
	})
}
