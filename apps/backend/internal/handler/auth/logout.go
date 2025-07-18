package auth

import (
	"connectrpc.com/connect"
	"context"
	"panelium/backend/internal/middleware"
	"panelium/backend/internal/security/cookies"
	"panelium/backend/internal/security/session"
	"panelium/common/errors"
	"panelium/proto_gen_go"
)

func (s *AuthServiceHandler) Logout(
	ctx context.Context,
	req *connect.Request[proto_gen_go.Empty],
) (*connect.Response[proto_gen_go.SuccessMessage], error) {
	// TODO: make this only need either access or refresh token, not both or just access token

	sessionInfoData := ctx.Value("panelium_session_info")
	sessionInfo, ok := sessionInfoData.(*middleware.SessionInfo)
	if !ok || sessionInfo == nil || sessionInfo.SessionID == "" || sessionInfo.UserID == "" {
		return nil, errors.ConnectInvalidCredentials
	}

	err := session.DeleteSession(sessionInfo.SessionID)
	if err != nil {
		res := connect.NewResponse(&proto_gen_go.SuccessMessage{
			Success: false,
		})
		// TODO: log this error?
		return res, connect.NewError(connect.CodeInternal, errors.SessionDeletionFailed)
	}

	res := connect.NewResponse(&proto_gen_go.SuccessMessage{
		Success: true,
	})

	cookies.ClearJWTCookie(res.Header(), "refresh_jwt")
	cookies.ClearJWTCookie(res.Header(), "access_jwt")

	return res, nil
}
