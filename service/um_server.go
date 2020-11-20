package service

import (
	"context"
	"fmt"

	"github.com/Sheerley/pluggabl/codes"
	"github.com/Sheerley/pluggabl/db"
	"github.com/Sheerley/pluggabl/plog"

	"github.com/Sheerley/pluggabl/pb"
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
	defer plog.ContextStatus(ctx)

	credentials := req.GetCredentials()
	plog.Verbosef("recieved Create User request for user %v", credentials.UserId)

	if len(credentials.UserId) > 0 {
		// checking if user exists
		err = db.UserExists(ctx, credentials)
		if err != nil {
			if err == (&codes.RecordExists{}) {
				err = fmt.Errorf("user already exists: %v", err)
			}

			respBody := pb.Response{
				ReturnMessage: err.Error(),
				ReturnCode:    pb.Response_error,
			}

			rsp = &pb.ActionUserResponse{
				Response: &respBody,
			}

			plog.Errorf("%v", err)

			return
		}

		err = db.CreateUser(ctx, credentials)
		if err != nil {
			err = fmt.Errorf("unable to create user: %v", err)

			respBody := pb.Response{
				ReturnMessage: err.Error(),
				ReturnCode:    pb.Response_error,
			}

			rsp = &pb.ActionUserResponse{
				Response: &respBody,
			}

			plog.Errorf("%v", err)

			return
		}

		respBody := pb.Response{
			ReturnMessage: "User created succesfully",
			ReturnCode:    pb.Response_ok,
		}

		rsp = &pb.ActionUserResponse{
			Response: &respBody,
		}

		plog.Errorf("%v", err)

		return
	}

	respBody := pb.Response{
		ReturnMessage: "Username is too short",
		ReturnCode:    pb.Response_error,
	}

	rsp = &pb.ActionUserResponse{
		Response: &respBody,
	}

	err = fmt.Errorf("username length is equal to or smaller than 0")

	plog.Errorf("%v", err)

	return
}

// ModifyUser is a unary RPC to modify an existing user
func (srv *UserManagementServer) ModifyUser(ctx context.Context, req *pb.ModifyUserRequest) (rsp *pb.ActionUserResponse, err error) {
	defer plog.ContextStatus(ctx)

	oldCredentials := req.GetOldCredentials()
	plog.Verbosef("recieved Modify User request for user %v", oldCredentials.UserId)

	credentials := req.GetNewCredentials()

	if len(oldCredentials.UserId) > 0 {
		// checking if user exists
		err = db.UserExists(ctx, oldCredentials)
		if err != nil {
			err = fmt.Errorf("user does not exists: %v", err)

			respBody := pb.Response{
				ReturnMessage: err.Error(),
				ReturnCode:    pb.Response_error,
			}

			rsp = &pb.ActionUserResponse{
				Response: &respBody,
			}

			plog.Errorf("%v", err)

			return
		}

		err = db.ModifyUser(ctx, credentials, oldCredentials)
		if err != nil {
			err = fmt.Errorf("Unable to delete user: %v", err)

			respBody := pb.Response{
				ReturnMessage: err.Error(),
				ReturnCode:    pb.Response_error,
			}

			rsp = &pb.ActionUserResponse{
				Response: &respBody,
			}

			plog.Errorf("%v", err)

			return
		}

		respBody := pb.Response{
			ReturnMessage: "User deleted succesfully",
			ReturnCode:    pb.Response_ok,
		}

		rsp = &pb.ActionUserResponse{
			Response: &respBody,
		}

		plog.Errorf("%v", err)

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

	plog.Errorf("%v", err)

	return
}

// DeleteUser is a unary RPC to delete an existing user
func (srv *UserManagementServer) DeleteUser(ctx context.Context, req *pb.ActionUserRequest) (rsp *pb.ActionUserResponse, err error) {
	defer plog.ContextStatus(ctx)

	credentials := req.GetCredentials()
	plog.Verbosef("recieved Delete User request for user %v", credentials.UserId)

	if len(credentials.UserId) > 0 {
		// checking if user exists

		err = db.UserExists(ctx, credentials)
		if err != nil {
			err = fmt.Errorf("user does not exists: %v", err)

			respBody := pb.Response{
				ReturnMessage: err.Error(),
				ReturnCode:    pb.Response_error,
			}

			rsp = &pb.ActionUserResponse{
				Response: &respBody,
			}

			plog.Errorf("%v", err)

			return
		}

		err = db.DeleteUser(ctx, credentials)
		if err != nil {
			err = fmt.Errorf("Unable to delete user: %v", err)

			respBody := pb.Response{
				ReturnMessage: err.Error(),
				ReturnCode:    pb.Response_error,
			}

			rsp = &pb.ActionUserResponse{
				Response: &respBody,
			}

			plog.Errorf("%v", err)

			return
		}

		respBody := pb.Response{
			ReturnMessage: "User deleted succesfully",
			ReturnCode:    pb.Response_ok,
		}

		rsp = &pb.ActionUserResponse{
			Response: &respBody,
		}

		plog.Errorf("%v", err)

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

	plog.Errorf("%v", err)

	return
}
