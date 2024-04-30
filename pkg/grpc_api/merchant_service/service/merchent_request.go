package service

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/akmal4410/gestapo/pkg/api/proto"
	"github.com/akmal4410/gestapo/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (merchant *merchantService) GetProfile(ctx context.Context, req *proto.GetMerchantProfileRequest) (*proto.GetMerchantProfileResponse, error) {
	if req.GetUserId() == "" {
		merchant.log.LogError("Error while Getting user id")
		return nil, status.Errorf(codes.InvalidArgument, utils.InvalidRequest)
	}
	res, err := merchant.storage.CheckDataExist("user_data", "id", req.GetUserId())
	if err != nil {
		merchant.log.LogError("Error while CheckUserExist", err)
		return nil, status.Errorf(codes.Internal, utils.InternalServerError)
	}
	if !res {
		err = fmt.Errorf("account does'nt exist using %s", req.GetUserId())
		merchant.log.LogError(err)
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	userData, err := merchant.storage.GetProfile(req.UserId)
	if err != nil {
		if err == sql.ErrNoRows {
			merchant.log.LogError("Error while GetProfile", err)
			return nil, status.Errorf(codes.NotFound, utils.NotFound)
		}
		merchant.log.LogError("Error while GetProfile", err)
		return nil, status.Errorf(codes.Internal, utils.InternalServerError)
	}

	if userData.ProfileImage != nil {
		url, err := merchant.s3.GetPreSignedURL(*userData.ProfileImage)
		if err != nil {
			merchant.log.LogError("Error while GetPreSignedURL", err)
			return nil, status.Errorf(codes.Internal, utils.InternalServerError)
		}
		userData.ProfileImage = &url
	}
	respone := &proto.GetMerchantProfileResponse{
		Code:    http.StatusOK,
		Status:  true,
		Message: "Profile fetched sucessfull",
		Data:    userData,
	}
	return respone, nil
}

// func (merchant *merchantService) EditProfile(ctx context.Context, req *proto.EditMerchantRequest) (*proto.Response, error) {
// 	const (
// 		thirtyTwoMB      = 32 << 20
// 		maxFileCount int = 1
// 	)

// 	// Extract the JSON data from the form
// 	jsonData := r.FormValue("data")
// 	reader := io.Reader(strings.NewReader(jsonData))

// 	err := helpers.ValidateBody(reader, req)
// 	if err != nil {
// 		merchant.log.LogError("Error while ValidateBody", err)
// 		helpers.ErrorJson(http.StatusBadRequest, InvalidBody)
// 		return
// 	}

// 	err = r.ParseMultipartForm(thirtyTwoMB)
// 	if err != nil {
// 		merchant.log.LogError("Unable to parse form", err.Error())
// 		helpers.ErrorJson(http.StatusBadRequest, StatusBadRequest)
// 		return
// 	}

// 	payload := r.Context().Value(utils.AuthorizationPayloadKey).(*token.AccessPayload)

// 	files := r.MultipartForm.File["files"]
// 	if len(files) > maxFileCount {
// 		merchant.log.LogError("Too many files uploaded", "Max allowed: %d", maxFileCount)
// 		errMsg := fmt.Sprintf("too many files uploaded. Max allowed: %s", strconv.Itoa(maxFileCount))
// 		helpers.ErrorJson(http.StatusBadRequest, errMsg)
// 		return
// 	}

// 	var uploadedFileKeys []string
// 	for _, fileHeader := range files {
// 		file, err := fileHeader.Open()
// 		if err != nil {
// 			merchant.log.LogError("Unable to open file", err)
// 			helpers.ErrorJson(http.StatusInternalServerError, "Unable to open file")
// 			return
// 		}
// 		defer file.Close()

// 		folderPath := "profile/" + payload.UserID + "/"
// 		fileURL, err := merchant.s3.UploadFileToS3(file, folderPath, fileHeader.Filename)
// 		if err != nil {
// 			merchant.log.LogError("Error uploading file to S3", err)
// 			helpers.ErrorJson(http.StatusInternalServerError, "Error uploading file to S3")
// 			return
// 		}

// 		merchant.log.LogInfo("File uploaded to S3 successfully", "FileURL:", fileURL)
// 		uploadedFileKeys = append(uploadedFileKeys, fileURL)
// 	}

// 	if len(uploadedFileKeys) != 0 {
// 		req.ProfileImage = uploadedFileKeys[0]
// 	}
// 	err = merchant.storage.UpdateProfile(payload.UserID, req)
// 	if err != nil {
// 		merchant.log.LogError("Error while UpdateProfile", err)
// 		helpers.ErrorJson(http.StatusInternalServerError, InternalServerError)
// 		return
// 	}

// 	helpers.WriteJSON(w, http.StatusOK, "User updated successfully")
// }
