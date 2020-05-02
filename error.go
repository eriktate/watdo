package wrkhub

type ErrType int

const (
	ErrInvalid = iota
	ErrStoreFailure
)

type WrkhubErr struct {
	errType ErrType
	msg     string // a msg that should only be viewable internally
	safeMsg string // a msg that's safe to surface to users
}

func NewErr(errType ErrType, msg string) WrkhubErr {
	return WrkhubErr{
		errType: errType,
		msg:     msg,
	}
}

func (e WrkhubErr) WithSafeMsg(msg string) WrkhubErr {
	e.safeMsg = msg
	return e
}

func (e WrkhubErr) Error() string {
	return e.msg
}

func isErrType(errType ErrType, e error) bool {
	if err, ok := e.(WrkhubErr); ok {
		return err.errType == errType
	}

	return false
}

func IsErrInvalid(e error) bool {
	return isErrType(ErrInvalid, e)
}

func IsErrStoreFailure(e error) bool {
	return isErrType(ErrStoreFailure, e)
}

func (e WrkhubErr) ErrType() ErrType {
	return e.errType
}

func (e WrkhubErr) SafeMsg() string {
	return e.safeMsg
}
