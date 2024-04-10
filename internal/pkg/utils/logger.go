package utils

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

const (
	DeliveryLayer   = "deliveryLayer"
	UsecaseLayer    = "usecaseLayer"
	RepositoryLayer = "repositoryLayer"
)

func LogInfo(logger *zap.Logger, requestId string, layer string, methodName string, message string) {
	logger.Info(
		fmt.Sprintf("REQUEST %s. INFO: %v", requestId, message),
		zap.String("layer", layer),
		zap.String("method", methodName),
		zap.String("requestId", requestId),
	)
}

func LogError(logger *zap.Logger, requestId string, layer string, methodName string, err error) {
	logger.Error(
		fmt.Sprintf("REQUEST %s. ERROR: %v", requestId, err.Error()),
		zap.String("layer", layer),
		zap.String("method", methodName),
		zap.String("requestId", requestId),
	)
}

func LogErrorResponse(logger *zap.Logger, requestId string, layer string, methodName string, err error, status int) {
	logger.Error(
		fmt.Sprintf("REQUEST %s. ERROR: %v", requestId, err.Error()),
		zap.String("layer", layer),
		zap.String("method", methodName),
		zap.String("requestId", requestId),
		zap.Int("responseStatus", status),
	)
}

func LogSucces(logger *zap.Logger, requestId string, layer string, methodName string) {
	logger.Info(
		fmt.Sprintf("REQUEST %s. OK", requestId),
		zap.String("layer", layer),
		zap.String("method", methodName),
		zap.String("requestId", requestId),
	)
}

func LogSuccesResponse(logger *zap.Logger, requestId string, layer string, methodName string) {
	logger.Info(
		fmt.Sprintf("REQUEST %s. OK", requestId),
		zap.String("layer", layer),
		zap.String("method", methodName),
		zap.String("requestId", requestId),
		zap.Int("responseStatus", http.StatusOK),
	)
}
