package kuroctxfiber

//type DefaultHandleError struct {
//	Code    int
//	Message string
//}
//
//func (f *defaultFiber) HandlerError(structs any) {
//	defaultError := structs
//	if defaultError == nil {
//		defaultError = &DefaultHandleError{Code: 500, Message: ""}
//	}
//
//	CFG.ErrorHandler = func(ctx *fiber.Ctx, err error) error {
//		return ctx.Status(fiber.StatusInternalServerError).JSON(defaultError)
//	}
//}
