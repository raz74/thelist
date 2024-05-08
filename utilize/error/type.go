package error

import "errors"

var (
	RestaurantNotFoundError     = errors.New("Restaurant Not Found")
	RestaurantDontCreatedError  = errors.New("Restaurant Created Failed")
	SendNotificationFailedError = errors.New("Send Notification Failed")
)
