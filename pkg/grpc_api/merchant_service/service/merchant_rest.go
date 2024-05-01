package service

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/akmal4410/gestapo/pkg/grpc_api/merchant_service/db/entity"
	"github.com/akmal4410/gestapo/pkg/helpers"
	"github.com/akmal4410/gestapo/pkg/helpers/token"
	"github.com/akmal4410/gestapo/pkg/utils"
)

func (handler *MerchantService) EditProfile(w http.ResponseWriter, r *http.Request) {
	const (
		thirtyTwoMB      = 32 << 20
		maxFileCount int = 1
	)

	// Extract the JSON data from the form
	jsonData := r.FormValue("data")
	reader := io.Reader(strings.NewReader(jsonData))

	req := new(entity.EditMerchantReq)
	err := helpers.ValidateBody(reader, req)
	if err != nil {
		handler.Log.LogError("Error while ValidateBody", err)
		helpers.ErrorJson(http.StatusBadRequest, utils.InvalidRequest)
		return
	}

	err = r.ParseMultipartForm(thirtyTwoMB)
	if err != nil {
		handler.Log.LogError("Unable to parse form", err.Error())
		helpers.ErrorJson(http.StatusBadRequest, utils.InvalidRequest)
		return
	}

	payload := r.Context().Value(utils.AuthorizationPayloadKey).(*token.AccessPayload)

	files := r.MultipartForm.File["files"]
	if len(files) > maxFileCount {
		handler.Log.LogError("Too many files uploaded", "Max allowed: %d", maxFileCount)
		errMsg := fmt.Sprintf("too many files uploaded. Max allowed: %s", strconv.Itoa(maxFileCount))
		helpers.ErrorJson(http.StatusBadRequest, errMsg)
		return
	}

	var uploadedFileKeys []string
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			handler.Log.LogError("Unable to open file", err)
			helpers.ErrorJson(http.StatusInternalServerError, "Unable to open file")
			return
		}
		defer file.Close()

		folderPath := "profile/" + payload.UserID + "/"
		fileURL, err := handler.s3.UploadFileToS3(file, folderPath, fileHeader.Filename)
		if err != nil {
			handler.Log.LogError("Error uploading file to S3", err)
			helpers.ErrorJson(http.StatusInternalServerError, "Error uploading file to S3")
			return
		}

		handler.Log.LogInfo("File uploaded to S3 successfully", "FileURL:", fileURL)
		uploadedFileKeys = append(uploadedFileKeys, fileURL)
	}

	if len(uploadedFileKeys) != 0 {
		req.ProfileImage = uploadedFileKeys[0]
	}
	err = handler.storage.UpdateProfile(payload.UserID, req)
	if err != nil {
		handler.Log.LogError("Error while UpdateProfile", err)
		helpers.ErrorJson(http.StatusInternalServerError, utils.InternalServerError)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, "User updated successfully")
}
