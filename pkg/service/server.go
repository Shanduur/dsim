package service

import (
	"context"
	"fmt"

	"github.com/Sheerley/pluggabl/pkg/db"
	"github.com/Sheerley/pluggabl/pkg/plog"

	"github.com/Sheerley/pluggabl/pkg/pb"
)

// UserManagementServer struct is implementation of the gRPC server
type UserManagementServer struct {
}

// NewUserManagementServer function initializes new server for the gRPC
func NewUserManagementServer() *UserManagementServer {
	return &UserManagementServer{}
}

// CreateUser is a unary RPC to create a new user
func (srv *UserManagementServer) CreateUser(ctx context.Context, req *pb.ActionUserRequest) (rsp *pb.ActionUserResponse, err error) {
	credentials := req.GetCredentials()
	plog.Messagef("recieved Create User request for user %v", credentials.UserId)

	if len(credentials.UserId) > 0 {
		// checking if user exists
		err = db.UserExists(credentials)
		if err != nil {
			err = fmt.Errorf("User already exists: %v", err)

			respBody := pb.Response{
				ReturnMessage: err.Error(),
				ReturnCode:    pb.Response_error,
			}

			rsp = &pb.ActionUserResponse{
				Response: &respBody,
			}

			return
		}

		err = db.CreateUser(credentials)
		if err != nil {
			err = fmt.Errorf("Unable to create user: %v", err)

			respBody := pb.Response{
				ReturnMessage: err.Error(),
				ReturnCode:    pb.Response_error,
			}

			rsp = &pb.ActionUserResponse{
				Response: &respBody,
			}

			return
		}

		respBody := pb.Response{
			ReturnMessage: "User created succesfully",
			ReturnCode:    pb.Response_ok,
		}

		rsp = &pb.ActionUserResponse{
			Response: &respBody,
		}

		return

	}

	respBody := pb.Response{
		ReturnMessage: "Username is too short",
		ReturnCode:    pb.Response_error,
	}

	rsp = &pb.ActionUserResponse{
		Response: &respBody,
	}

	err = fmt.Errorf("username length is equal or smaller than 0")

	return
}

// DeleteUser is a unary RPC to delete an existing user
func (srv *UserManagementServer) DeleteUser(ctx context.Context, req *pb.ActionUserRequest) (rsp *pb.ActionUserResponse, err error) {
	credentials := req.GetCredentials()
	plog.Messagef("recieved Delete User request for user %v", credentials.UserId)

	if len(credentials.UserId) > 0 {
		// checking if user exists
		err = db.UserExists(credentials)
		if err != nil {
			err = fmt.Errorf("User does not exists: %v", err)

			respBody := pb.Response{
				ReturnMessage: err.Error(),
				ReturnCode:    pb.Response_error,
			}

			rsp = &pb.ActionUserResponse{
				Response: &respBody,
			}

			return
		}

		err = db.DeleteUser(credentials)
		if err != nil {
			err = fmt.Errorf("Unable to delete user: %v", err)

			respBody := pb.Response{
				ReturnMessage: err.Error(),
				ReturnCode:    pb.Response_error,
			}

			rsp = &pb.ActionUserResponse{
				Response: &respBody,
			}

			return
		}

		respBody := pb.Response{
			ReturnMessage: "User deleted succesfully",
			ReturnCode:    pb.Response_ok,
		}

		rsp = &pb.ActionUserResponse{
			Response: &respBody,
		}

		return
	}

	respBody := pb.Response{
		ReturnMessage: "Username is too short",
		ReturnCode:    pb.Response_error,
	}

	rsp = &pb.ActionUserResponse{
		Response: &respBody,
	}

	err = fmt.Errorf("username length is equal or smaller than 0")

	return
}
