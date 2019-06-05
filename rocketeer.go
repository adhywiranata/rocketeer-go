package rocketeer

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

// HandlerFunc Function type definition
type HandlerFunc func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

// NoOpMiddleware basically does nothing
func NoOpMiddleware(next HandlerFunc) HandlerFunc {
	return HandlerFunc(func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		log.Println("=> middleware accessed without any operation")
		return next(request)
	})
}

// WarmerInterceptorMiddleware as the interceptor will ping a  warmed cold start
func WarmerInterceptorMiddleware(next HandlerFunc) HandlerFunc {
	type warmer struct {
		Warmer bool `json:"warmer"`
	}

	return HandlerFunc(func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		var w warmer

		// If unmarshal into the ping struct is successful and there was a value in ping, return out.
		if err := json.Unmarshal([]byte(request.Body), &w); err == nil && w.Warmer == true {
			log.Println("function warmed up!")
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusOK,
				Body:       "warmed",
			}, nil
		}
		return next(request)
	})
}
