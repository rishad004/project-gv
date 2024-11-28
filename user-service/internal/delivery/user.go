package delivery

import (
	"context"
	"log"

	pb "github.com/rishad004/Gv_protofiles/user"
	"github.com/rishad004/project-gv/user-service/internal/domain"
)

func (h *UserHandler) SignUp(c context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {

	log.Println("Signup................................................................")

	user := domain.Users{
		Username: req.Username,
		Email:    req.Email,
		Phone:    req.Phone,
		Hashed:   req.Password,
		Gender:   req.Gender,
		Verified: false,
	}

	link, err := h.svc.SignUp(user)

	if err != nil {
		return nil, err
	}

	return &pb.SignUpResponse{Link: link, Message: "User signed up successfully!"}, nil
}

func (h *UserHandler) EmailVerification(c context.Context, req *pb.VerificationRequest) (*pb.VerificationResponse, error) {

	log.Println("EmailVerify................................................................")

	if err := h.svc.EmailVerification(req.Key); err != nil {
		return nil, err
	}
	return &pb.VerificationResponse{Message: "Email verified successfully!"}, nil
}

func (h *UserHandler) Login(c context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {

	log.Println("Login................................................................")

	token, err := h.svc.Login(req.Username, req.Password)

	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{Token: token, Message: "Logged in successfully!"}, nil
}

func (h *UserHandler) Profile(c context.Context, req *pb.Verification) (*pb.ProfileResponse, error) {

	log.Println("ProfileView................................................................")

	user, err := h.svc.Profile(int(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.ProfileResponse{Id: int32(user.ID), Username: user.Username, Email: user.Email, Phone: user.Phone, Gender: user.Gender}, nil
}

func (h *UserHandler) ProfileEditing(c context.Context, req *pb.EditRequest) (*pb.EditResponse, error) {

	log.Println("ProfileEdit................................................................")

	edits := domain.Users{
		Username: req.Username,
		Phone:    req.Phone,
		Gender:   req.Gender,
	}
	edits.ID = uint(req.Id)

	if err := h.svc.ProfileEditing(edits); err != nil {
		return nil, err
	}
	return &pb.EditResponse{Message: "Profile edited successfully!"}, nil
}

func (h *UserHandler) Following(c context.Context, req *pb.SubscribeRequest) (*pb.Empty, error) {

	log.Println("following")

	if err := h.svc.Following(int(req.Id), int(req.Userid)); err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}
