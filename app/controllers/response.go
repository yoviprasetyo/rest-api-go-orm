package controllers

import (
	"fmt"
	"orm/app/helper"

	"github.com/gin-gonic/gin"
)

func responseAPI(response Response) {
	var (
		result = gin.H{
			"status": "success",
		}
	)
	if response.ErrorMessage != "" {
		result["status"] = "error"
		result["message"] = response.ErrorMessage
		response.Context.JSON(response.StatusCode, result)
		response.Context.Abort()
	} else {
		result["data"] = response.Data
		result["total"] = response.Total

		if response.UsePagination {
			result["per_page"] = response.PerPage
			result["from"] = ((response.Page - 1) * response.PerPage) + 1
			result["to"] = getTo(response)
			result["current_page"] = response.Page
			result["last_page"] = getLastPage(response)
			result["current_page_url"] = response.ResponseURL.CurrentPageURL
			result["first_page_url"] = response.ResponseURL.FirstPageURL
			result["last_page_url"] = response.ResponseURL.LastPageURL
			result["next_page_url"] = response.ResponseURL.NextPageURL
			result["prev_page_url"] = response.ResponseURL.PrevPageURL
		}
		response.Context.JSON(response.StatusCode, result)
	}
}

// Response struct.
type Response struct {
	Context       *gin.Context
	StatusCode    int
	Data          interface{}
	Total         int64
	Page          int
	PerPage       int
	Additional    interface{}
	ErrorMessage  string
	ResponseURL   ResponseURL
	UsePagination bool
}

// ResponseURL struct.
type ResponseURL struct {
	CurrentPageURL string `json:"current_page_url"`
	FirstPageURL   string `json:"first_page_url"`
	LastPageURL    string `json:"last_page_url"`
	NextPageURL    string `json:"next_page_url"`
	PrevPageURL    string `json:"prev_page_url"`
}

// SetCurrentPageURL method.
func (responseURL *ResponseURL) SetCurrentPageURL(path string) {
	responseURL.CurrentPageURL = helper.BaseURL(path)
}

// SetFirstPageURL method.
func (responseURL *ResponseURL) SetFirstPageURL(path string) {
	responseURL.FirstPageURL = helper.BaseURL(path)
	if responseURL.FirstPageURL == responseURL.CurrentPageURL {
		responseURL.FirstPageURL = ""
	}
}

// SetLastPageURL method.
func (responseURL *ResponseURL) SetLastPageURL(path string) {
	responseURL.LastPageURL = helper.BaseURL(path)
	if responseURL.LastPageURL == responseURL.CurrentPageURL {
		responseURL.LastPageURL = ""
	}
}

// SetNextPageURL method.
func (responseURL *ResponseURL) SetNextPageURL(path string) {
	responseURL.NextPageURL = helper.BaseURL(path)
	if responseURL.LastPageURL == "" {
		responseURL.NextPageURL = ""
	}
}

// SetPrevPageURL method.
func (responseURL *ResponseURL) SetPrevPageURL(path string) {
	responseURL.PrevPageURL = helper.BaseURL(path)
	if responseURL.FirstPageURL == "" {
		responseURL.PrevPageURL = ""
	}
}

// getTo total and per page.
func getTo(response Response) int64 {
	var to int64
	to = int64(((response.Page - 1) * response.PerPage) + response.PerPage)
	fmt.Println(response.Total, to)
	if response.Total < to {
		return response.Total
	}
	return to
}

// getLastPage method.
func getLastPage(response Response) int {
	total := response.Total
	perPage := response.PerPage
	modPage := total % int64(perPage)
	page := (total / int64(perPage))
	if modPage > 0 {
		return int(page) + 1
	}
	return int(page)
}
