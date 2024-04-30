package service

import (
	"context"
	"fmt"

	"github.com/akmal4410/gestapo/pkg/api/proto"
	"github.com/akmal4410/gestapo/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (admin *adminService) CreateCategory(ctx context.Context, req *proto.AddCategoryRequest) (*proto.Response, error) {
	err := validateAddCategoryRequest(req)
	if err != nil {
		admin.log.LogError("Error while validateAddCategoryRequest", err)
		return nil, status.Errorf(codes.InvalidArgument, utils.InvalidRequest)
	}

	res, err := admin.storage.CheckCategoryExist(req.GetCategoryName())
	if err != nil {
		admin.log.LogError("Error while CheckCategoryExist", err)
		return nil, status.Errorf(codes.Internal, utils.InternalServerError)
	}
	if res {
		err = fmt.Errorf("category already exist: %s", req.GetCategoryName())
		admin.log.LogError("Error ", err)
		return nil, status.Errorf(codes.AlreadyExists, err.Error())
	}

	err = admin.storage.AddCategory(req)
	if err != nil {
		admin.log.LogError("Error while InsertCategory", err)
		return nil, status.Errorf(codes.Internal, utils.InternalServerError)
	}
	response := &proto.Response{
		Code:    200,
		Status:  true,
		Message: "Category insterted Successfully",
	}
	return response, err
}

func (admin *adminService) GetCategories(ctx context.Context, in *proto.Request) (*proto.GetCategoryResponse, error) {
	res, err := admin.storage.GetCategories()
	if err != nil {
		admin.log.LogError("Error while GetCategories", err)
		return nil, status.Errorf(codes.Internal, utils.InternalServerError)
	}
	response := &proto.GetCategoryResponse{
		Code:    200,
		Status:  true,
		Message: "Categories fetched successfull",
		Data:    res,
	}
	return response, nil
}

func (admin *adminService) GetUsers(ctx context.Context, in *proto.Request) (*proto.GetUsersResponse, error) {
	res, err := admin.storage.GetUsers()
	if err != nil {
		admin.log.LogError("Error while GetUsers", err)
		return nil, status.Errorf(codes.Internal, utils.InternalServerError)
	}

	response := &proto.GetUsersResponse{
		Code:    200,
		Status:  true,
		Message: "Users fetched successfull",
		Data:    res,
	}
	return response, nil
}