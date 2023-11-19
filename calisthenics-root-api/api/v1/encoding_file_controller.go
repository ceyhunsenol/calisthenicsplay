package v1

import (
	"calisthenics-root-api/data"
	"calisthenics-root-api/service"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type EncodingFileController struct {
	encodingFileService service.IEncodingFileService
	cacheRequestService service.ICacheRequestService
}

func NewEncodingFileController(
	encodingFileService service.IEncodingFileService,
	cacheRequestService service.ICacheRequestService,
) *EncodingFileController {
	return &EncodingFileController{
		encodingFileService: encodingFileService,
		cacheRequestService: cacheRequestService,
	}
}

func (ef *EncodingFileController) InitEncodingFileRoutes(g *echo.Group) {
	g.POST("", ef.SaveFiles)
	g.GET("/encoding/:encodingID", ef.GetEncodingFilesByEncodingID)
	g.DELETE("/encoding/:encodingID", ef.DeleteEncodingFilesByEncodingID)
	g.GET("/:id", ef.GetEncodingFile)
	g.DELETE("/:id", ef.DeleteEncodingFile)
}

func (ef *EncodingFileController) SaveFiles(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: "Invalid form data"})
	}
	files := form.File["files"]
	encodingID := c.FormValue("encodingID")
	iv := c.FormValue("iv")
	ext := c.FormValue("ext")
	if encodingID == "" {
		return c.JSON(http.StatusBadRequest, &MessageResource{Message: "encodingID required"})
	}
	if iv == "" {
		return c.JSON(http.StatusBadRequest, &MessageResource{Message: "iv required"})
	}
	if ext == "" {
		return c.JSON(http.StatusBadRequest, &MessageResource{Message: "extInf required"})
	}
	ivValues := strings.Split(iv, ",")
	if len(ivValues) != len(files) {
		return c.JSON(http.StatusBadRequest, &MessageResource{Message: "iv and file lists must be equal in length"})
	}

	extValues := strings.Split(ext, ",")
	if len(extValues) != len(files) {
		return c.JSON(http.StatusBadRequest, &MessageResource{Message: "ext and file lists must be equal in length"})
	}
	if files != nil {
		rootPath, pathError := os.Getwd()
		if pathError != nil {
			return c.JSON(http.StatusInternalServerError, &MessageResource{Message: "Unknown error"})
		}
		cdnDir := filepath.Join(rootPath, "../calisthenics-cdn-api/medias")
		for _, file := range files {
			if filepath.Ext(file.Filename) != ".ts" {
				return c.JSON(http.StatusInternalServerError, &MessageResource{Message: "Just ts file"})
			}
		}
		for _, file := range files {
			src, fileError := file.Open()
			if fileError != nil {
				return c.JSON(http.StatusInternalServerError, &MessageResource{Message: "Unable to open the file"})
			}
			defer src.Close()

			filePath := filepath.Join(cdnDir, file.Filename)

			dst, fileError := os.Create(filePath)
			if fileError != nil {
				return c.JSON(http.StatusInternalServerError, &MessageResource{Message: "Unable to create the file"})
			}
			defer dst.Close()

			if _, err = io.Copy(dst, src); err != nil {
				return c.JSON(http.StatusInternalServerError, &MessageResource{Message: "Unable to copy the file"})
			}
		}
	}

	floatExt := make([]float64, 0)
	for _, value := range extValues {
		ex, _ := strconv.ParseFloat(value, 64)
		floatExt = append(floatExt, ex)
	}

	savedFiles := make([]data.EncodingFile, 0)
	for i, file := range files {
		encodingFile := data.EncodingFile{
			FileName:   file.Filename,
			EncodingID: encodingID,
			IV:         ivValues[i],
			Ext:        floatExt[i],
		}
		savedFile, fileError := ef.encodingFileService.Save(encodingFile)
		if fileError != nil {
			return c.JSON(http.StatusInternalServerError, &MessageResource{Message: "Unable to saved the file"})
		}
		savedFiles = append(savedFiles, *savedFile)
	}

	// content apiye cache icin request atiliyor
	serviceError := ef.cacheRequestService.HLSRefreshRequest(encodingID)
	if serviceError != nil && serviceError.Message != "Request error" {
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: serviceError.Message})
	}

	encodingFileResources := make([]EncodingFileResource, 0)
	for _, file := range savedFiles {
		encodingFileResources = append(encodingFileResources, EncodingFileResource{
			ID:         file.ID,
			EncodingID: file.EncodingID,
			FileName:   file.FileName,
			IV:         file.IV,
		})
	}

	return c.JSON(http.StatusOK, encodingFileResources)
}

func (ef *EncodingFileController) GetEncodingFilesByEncodingID(c echo.Context) error {
	encodingID := c.Param("encodingID")
	encodings, err := ef.encodingFileService.GetAllEncodingFilesByEncodingID(encodingID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: "Encoding files could not be got"})
	}

	encodingFileResources := make([]EncodingFileResource, 0)
	for _, encoding := range encodings {
		encodingFileResources = append(encodingFileResources, EncodingFileResource{
			ID:         encoding.ID,
			EncodingID: encoding.EncodingID,
			FileName:   encoding.FileName,
			IV:         encoding.IV,
		})
	}

	return c.JSON(http.StatusOK, encodingFileResources)
}

func (ef *EncodingFileController) GetEncodingFile(c echo.Context) error {
	id := c.Param("id")
	encoding, err := ef.encodingFileService.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: "Encoding could not be got"})
	}

	encodingFileResource := EncodingFileResource{
		ID:         encoding.ID,
		FileName:   encoding.FileName,
		EncodingID: encoding.EncodingID,
		IV:         encoding.IV,
	}

	return c.JSON(http.StatusOK, encodingFileResource)
}

func (ef *EncodingFileController) DeleteEncodingFile(c echo.Context) error {
	id := c.Param("id")
	encodingFile, err := ef.encodingFileService.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &MessageResource{Message: "Encoding not found"})
	}
	err = ef.encodingFileService.Delete(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: "Encoding file could not be deleted"})
	}

	// content apiye cache icin request atiliyor
	serviceError := ef.cacheRequestService.HLSRefreshRequest(encodingFile.EncodingID)
	if serviceError != nil && serviceError.Message != "Request error" {
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: serviceError.Message})
	}
	return c.JSON(http.StatusNoContent, nil)
}

func (ef *EncodingFileController) DeleteEncodingFilesByEncodingID(c echo.Context) error {
	id := c.Param("encodingID")
	err := ef.encodingFileService.DeleteEncodingFilesByEncodingID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: "Encoding files could not be deleted"})
	}

	// content apiye cache icin request atiliyor
	serviceError := ef.cacheRequestService.HLSRefreshRequest(id)
	if serviceError != nil && serviceError.Message != "Request error" {
		return c.JSON(http.StatusInternalServerError, &MessageResource{Message: serviceError.Message})
	}
	return c.JSON(http.StatusNoContent, nil)
}
