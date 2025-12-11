package errs

import (
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HTTPStatus(err error) (int, string) {
	if err == nil {
		return http.StatusOK, ""
	}

	st, ok := status.FromError(err)
	if !ok {
		// Если ошибка не от gRPC (например, ошибка маршалинга или баз данных, не обернутая в status)
		// Возвращаем 500 и стандартный текст, чтобы не светить кишки наружу
		return http.StatusInternalServerError, "internal server error"
	}

	switch st.Code() {
	case codes.OK:
		return http.StatusOK, ""
	case codes.NotFound:
		return http.StatusNotFound, st.Message()
	case codes.InvalidArgument:
		return http.StatusBadRequest, st.Message()
	case codes.Unauthenticated:
		return http.StatusUnauthorized, st.Message()
	case codes.PermissionDenied:
		return http.StatusForbidden, st.Message()
	case codes.AlreadyExists:
		return http.StatusConflict, st.Message()
	case codes.Internal:
		return http.StatusInternalServerError, "internal server error"
	case codes.Unavailable:
		return http.StatusServiceUnavailable, st.Message()
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout, st.Message()
	case codes.ResourceExhausted:
		return http.StatusTooManyRequests, st.Message()
	case codes.Aborted:
		return http.StatusConflict, st.Message()
	case codes.Unimplemented:
		return http.StatusNotImplemented, st.Message()
	default:
		return http.StatusInternalServerError, "internal server error"
	}
}
