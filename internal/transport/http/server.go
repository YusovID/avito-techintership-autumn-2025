package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/YusovID/pr-reviewer-service/internal/apperrors"
	"github.com/YusovID/pr-reviewer-service/internal/service"
	"github.com/YusovID/pr-reviewer-service/pkg/api"
	"github.com/YusovID/pr-reviewer-service/pkg/logger/sl"
)

type Server struct {
	log         *slog.Logger
	teamService service.TeamService
	userService service.UserService
}

func NewServer(log *slog.Logger, ts service.TeamService, us service.UserService) *Server {
	return &Server{
		log:         log,
		teamService: ts,
		userService: us,
	}
}

func (s *Server) Routes() http.Handler {
	baseRouter := api.Handler(s)

	return s.logRequest(baseRouter)
}

func (s *Server) PostTeamAdd(w http.ResponseWriter, r *http.Request) {
	const op = "internal.transport.http.PostTeamAdd"
	log := s.log.With(slog.String("op", op))

	var req api.Team
	if err := s.decode(r, &req); err != nil {
		log.Error("failed to decode request body", sl.Err(err))
		s.respondError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	team, err := s.teamService.CreateTeam(r.Context(), req)
	if err != nil {
		log.Error("failed to create team", sl.Err(err))
		if errors.Is(err, apperrors.ErrAlreadyExists) {
			s.respondAPIError(w, http.StatusConflict, api.TEAMEXISTS, "team name already exists")
			return
		}
		s.respondError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	s.respond(w, http.StatusCreated, team)
}

func (s *Server) GetTeamGet(w http.ResponseWriter, r *http.Request, params api.GetTeamGetParams) {
	const op = "internal.transport.http.GetTeamGet"
	log := s.log.With(slog.String("op", op), slog.String("team_name", params.TeamName))

	team, err := s.teamService.GetTeam(r.Context(), params.TeamName)
	if err != nil {
		log.Error("failed to get team", sl.Err(err))
		if errors.Is(err, apperrors.ErrNotFound) {
			s.respondAPIError(w, http.StatusNotFound, api.NOTFOUND, "team not found")
			return
		}
		s.respondError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	s.respond(w, http.StatusOK, team)
}

func (s *Server) PostUsersSetIsActive(w http.ResponseWriter, r *http.Request) {
	const op = "internal.transport.http.PostUsersSetIsActive"
	log := s.log.With(slog.String("op", op))

	var req api.PostUsersSetIsActiveJSONBody
	if err := s.decode(r, &req); err != nil {
		log.Error("failed to decode request", sl.Err(err))
		s.respondError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	user, err := s.userService.SetIsActive(r.Context(), req.UserId, req.IsActive)
	if err != nil {
		log.Error("failed to set user active status", sl.Err(err))
		if errors.Is(err, apperrors.ErrNotFound) {
			s.respondAPIError(w, http.StatusNotFound, api.NOTFOUND, "user not found")
			return
		}
		s.respondError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	s.respond(w, http.StatusOK, user)
}

func (s *Server) PostPullRequestCreate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (s *Server) PostPullRequestMerge(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (s *Server) PostPullRequestReassign(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (s *Server) GetUsersGetReview(w http.ResponseWriter, r *http.Request, params api.GetUsersGetReviewParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (s *Server) respond(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			s.log.Error("failed to encode response", sl.Err(err))
		}
	}
}

func (s *Server) respondError(w http.ResponseWriter, code int, message string) {
	s.respond(w, code, map[string]string{"error": message})
}

func (s *Server) respondAPIError(w http.ResponseWriter, code int, apiCode api.ErrorResponseErrorCode, message string) {
	errResp := api.ErrorResponse{
		Error: struct {
			Code    api.ErrorResponseErrorCode `json:"code"`
			Message string                     `json:"message"`
		}{
			Code:    apiCode,
			Message: message,
		},
	}
	s.respond(w, code, errResp)
}

func (s *Server) decode(r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return fmt.Errorf("decode json: %w", err)
	}
	return nil
}
